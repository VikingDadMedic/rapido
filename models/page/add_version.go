/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package page

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"time"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/library"
)

// Places manifest extension data
type manifestObj struct {
	ContentSettings []map[string]interface{} `json:"content_settings,omitempty"`
}

// Generates versioning page at each change
func AddVersion(userID int, contentID int, pageID int) {
	SaveRestorePoint(CreateRestorePoint(userID, contentID, pageID))
}

// Creates a page version at each saving
func CreateRestorePoint(userID int, contentID int, pageID int) (output []crud.PageVersion) {
	var obj Page
	if contentID > 0 {
		data := getPageByContent(contentID)
		for index := range data {
			obj.PageVersion.PageID = data[index].Page.ID
			obj.PageVersion.Data = data[index]
			obj.PageVersion.CreatedOn = time.Now().Unix()
			output = append(output, obj.PageVersion)
		}
	} else if pageID > 0 {
		obj.PageVersion.PageID = pageID
		obj.PageVersion.Data.PageContent = getPageByID(pageID)
		obj.Page.Select("id", pageID)
		obj.PageVersion.Data.Page = obj.Page
		obj.PageVersion.CreatedOn = time.Now().Unix()
		output = append(output, obj.PageVersion)
	}
	return
}

// Saves versioning page at each change
func SaveRestorePoint(input []crud.PageVersion) {
	for index := range input {
		if err := input[index].Insert(); err != nil {
			log.Println(err)
		} else {
			input[index].Update()
		}
	}
}

func getPageByContent(contentID int) []crud.Data {
	var obj Page
	data := make([]crud.Data, 0)
	pageContent := obj.PageContent.SelectBy("content_id", contentID)
	if len(pageContent) > 0 {
		ids := make([]interface{}, 0)
		for index := range pageContent {
			ids = append(ids, pageContent[index].PageID)
		}
		library.UniqueInterface(&ids)

		page := obj.Page.SelectBy("id", ids)
		for index := range page {
			row := crud.Data{}
			row.PageContent = getPageByID(page[index].ID)
			row.Page = page[index]
			data = append(data, row)
		}
	}

	return data
}

func getPageByID(pageID int) []crud.PageContent {
	var obj Page
	pageContent := obj.PageContent.SelectBy("page_id", pageID)
	if len(pageContent) == 0 {
		return nil
	}

	ids := make([]interface{}, 0)
	for index := range pageContent {
		if pageContent[index].ContentID.Int64 > 0 {
			ids = append(ids, pageContent[index].ContentID.Int64)
		}
	}

	contents := obj.Content.SelectBy("id", ids)
	sort.Sort(crud.PageByPosition(pageContent))
	for index := range pageContent {
		pageContent[index].ID = 0
		for ico := range contents {
			if contents[ico].ID == int(pageContent[index].ContentID.Int64) {
				contents[ico].ID = 0
				pageContent[index].Content = contents[ico]
				break
			}
		}

		if pageContent[index].JSONSettings.String != "" {
			jsonSetting, _ := library.Jsonify(pageContent[index].JSONSettings.String)
			pageContent[index].Settings = jsonSetting
		}

		if pageContent[index].Extension.String != "" {
			manifest := manifestObj{}
			raw, err := ioutil.ReadFile(fmt.Sprintf("./extension/%v.json", pageContent[index].Extension.String))
			if err == nil {
				json.Unmarshal(raw, &manifest)
				if len(manifest.ContentSettings) > 0 {
					for ics := range manifest.ContentSettings {
						exist := false
						if len(pageContent[index].Settings) > 0 {
							for ist := range pageContent[index].Settings {
								if pageContent[index].Settings[ist]["name"] == manifest.ContentSettings[ics]["name"] {
									exist = true
									break
								}
							}
							if !exist {
								pageContent[index].Settings = append(pageContent[index].Settings, manifest.ContentSettings[ics])
							}
						} else {
							pageContent[index].Settings = manifest.ContentSettings
						}
					}
				}
			} else {
				log.Println(err)
			}
		}
	}

	return pageContent
}
