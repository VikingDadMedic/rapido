/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package contactform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/library"
	"framagit.org/InfoLibre/rapido/models/settings"

	"github.com/gin-gonic/gin"
)

// PageContactFormParam is used as page ContactForm parameter
type PageContactFormParam struct {
	ID   int    `form:"id"`
	Link string `form:"link"`
}

// manifestObj is used to place manifest extension data
type manifestObj struct {
	ContactFormSettings []map[string]interface{} `json:"contactform_settings,omitempty"`
}

/**
 * @api {GET} /api/ContactForm/page ContactForm On Page
 * @apiDescription ContactForm On Page
 * @apiGroup ContactForm
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiSuccess {Int} id PageContactForm ID
 * @apiSuccess {Int} page_id Page ID
 * @apiSuccess {Int} ContactForm_id ContactForm ID
 * @apiSuccess {String} extension Extension Name
 * @apiSuccess {String} location Location
 * @apiSuccess {Int} column Column
 * @apiSuccess {Int} position Position
 * @apiSuccess {[]Object} json_settings Page Setting
 * @apiSuccess {Object} ContactForm ContactForm Object
 * @apiSuccess {Int} ContactForm.id ContactForm ID
 * @apiSuccess {String} ContactForm.name ContactForm Name
 * @apiSuccess {String} ContactForm.ContactForm ContactForm
 * @apiSuccess {[]Object} ContactForm.json_settings ContactForm Setting
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * [
 *		{
 *			"id": 1,
 *			"page_id": 1,
 *			"ContactForm_id": 1,
 *			"extension": "test",
 *			"location": "header",
 *			"column": 1,
 *			"position": 1,
 *			"json_settings": [{}],
 *			"ContactForm": {
 *			    "id": 1,
 *			    "name": "header",
 *			    "ContactForm": "<h1>Rapido</h1><p>Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.</p>",
 *			}
 *		}
 * ]
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} [id] Page ID
 * @apiParam {String} [link] Page Link
 */

// Gets contact form data on page
func Get(c *gin.Context) {
	var param PageContactFormParam
	if err := c.ShouldBind(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}

	var obj ContactForm
	var pageID int
	if param.Link != "" {
		obj.Page.Select("link", param.Link)
		if obj.Page.IsInitial() {
			c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{"message": "incorrect_settings_parameter"})
			return
		}
		pageID = obj.Page.ID
	} else if param.ID > 0 {
		pageID = param.ID
		obj.Page.Select("id", param.Link)
		if obj.Page.IsInitial() {
			c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{"message": "incorrect_settings_parameter"})
			return
		}
	} else {
		pageID = settings.HomePage()
	}

	pageContactForm := obj.PageContent.SelectBy("page_id", pageID)
	if len(pageContactForm) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{"message": "page_not_found"})
		return
	}

	ids := make([]interface{}, 0)
	for index := range pageContactForm {
		if pageContactForm[index].ContactFormID.Int64 > 0 {
			ids = append(ids, pageContactForm[index].ContactFormID.Int64)
		}
	}
	ContactForms := obj.ContactForm.SelectBy("id", ids)

	sort.Sort(crud.PageByPosition(pageContactForm))

	for index := range pageContactForm {
		for ico := range ContactForms {
			if ContactForms[ico].ID == int(pageContactForm[index].ContactFormID.Int64) {
				pageContactForm[index].ContactForm = ContactForms[ico]
				break
			}
		}

		if pageContactForm[index].JSONSettings.String != "" {
			jsonSetting, _ := library.Jsonify(pageContactForm[index].JSONSettings.String)
			pageContactForm[index].Settings = jsonSetting
		}

		if pageContactForm[index].Extension.String != "" {
			manifest := manifestObj{}
			manifests := library.SearchFile(fmt.Sprintf("./client/src/extensions/%v", pageContactForm[index].Extension.String), ".manifest", "json")
			raw, err1 := ioutil.ReadFile(manifests[0])
			if err1 == nil {
				json.Unmarshal(raw, &manifest)
				if len(manifest.ContactFormSettings) > 0 {
					for ics := range manifest.ContactFormSettings {
						exist := false
						if len(pageContactForm[index].Settings) > 0 {
							for ist := range pageContactForm[index].Settings {
								if pageContactForm[index].Settings[ist]["name"] == manifest.ContactFormSettings[ics]["name"] {
									exist = true
									break
								}
							}
							if !exist {
								pageContactForm[index].Settings = append(pageContactForm[index].Settings, manifest.ContactFormSettings[ics])
							}
						} else {
							pageContactForm[index].Settings = manifest.ContactFormSettings
						}
					}
				}
			} else {
				log.Println(err1)
			}
		}
	}

	c.JSON(http.StatusOK, pageContactForm)
}
