/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package content

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"
	"framagit.org/InfoLibre/rapido/models/page"

	"github.com/gin-gonic/gin"
	"github.com/kennygrant/sanitize"
)

/**
 * @api {POST} /api/content Change Content On Page
 * @apiDescription Change Content On Page
 * @apiGroup Content
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * [
 *		{
 *			"id": 1,
 *			"page_id": 1,
 *			"content_id": 1,
 *			"extension": "test",
 *			"location": "header",
 *			"column": 1,
 *			"position": 1,
 *			"json_settings": [{}],
 *			"content": {
 *			    "id": 1,
 *			    "name": "header",
 *			    "content": "<h1>Rapido</h1><p>Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.</p>",
 *			}
 *		}
 * ]
 * @apiSuccess {String} message Success Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": "edit_page__page_changes_saved"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} [id] PageContent ID
 * @apiParam {Int} page_id Page ID
 * @apiParam {Int} [content_id] Content ID
 * @apiParam {String} [extension] Extension Name
 * @apiParam {String} location Location
 * @apiParam {Int} column Column
 * @apiParam {Int} position Position
 * @apiParam {[]Object} [json_settings] Page Setting
 * @apiParam {Object} content Content Object
 * @apiParam {Int} [content.id] Content ID
 * @apiParam {String} [content.name] Content Name
 * @apiParam {String} content.content Content
 * @apiParam {[]Object} [content.json_settings] Content Setting
 */

// Adds content data on page
func Add(c *gin.Context) {
	var param []*crud.PageContent
	if err1 := c.ShouldBindJSON(&param); err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err1.Error()})
		return
	}
	if len(param) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter"})
		return
	}
	var obj Content
	obj.Page.Select("id", param[0].PageID)
	if obj.Page.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "page_not_found"})
		return
	}

	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.AccessLevel != jwt.AdminAL {
		obj.PageCollaboration.Select(param[0].PageID, payload.ID)
		if obj.PageCollaboration.IsInitial() {
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "edit_page__cannot_edit_page"})
			return
		}
	}

	for index := range param {
		if param[index].Content.Content != "" {
			if param[index].Content.Name.String == "" {
				param[index].Content.Name.String = sanitize.HTML(strings.Replace(param[index].Content.Content, "><", "> <", -1))
				if len(param[index].Content.Name.String) > 30 {
					param[index].Content.Name.String = param[index].Content.Name.String[0:30]
				} else {
					param[index].Content.Name.String = param[index].Content.Name.String[0:len(param[index].Content.Name.String)]
				}
				param[index].Content.Name.Valid = true
			}
		}

		if param[index].ContentID.Int64 > 0 {
			if param[index].Extension.String == "" {
				obj.Content.Select("id", int(param[index].ContentID.Int64))
				err2 := param[index].Content.Update()
				if err2 != nil {
					log.Println(err2)
				}
				go saveMedia(c.Request.Host, obj.Content.Content, param[index].Content.Content)
			}

			err3 := param[index].Update()
			if err3 != nil {
				log.Println(err3)
			}
		} else {
			if param[index].Extension.String == "" && param[index].Content.Content != "" {
				err4 := param[index].Content.Insert()
				if err4 != nil {
					log.Println(err4)
				}
				go saveMedia(c.Request.Host, "", param[index].Content.Content)
			}

			if param[index].ID == 0 {
				param[index].ContentID.Int64, param[index].ContentID.Valid = int64(param[index].Content.ID), true
				err5 := param[index].Insert()
				if err5 != nil {
					log.Println(err5)
				}
			}
		}
	}

	current := obj.PageContent.SelectBy("page_id", param[0].PageID)
	for ic := range current {
		exist := false
		for ip := range param {
			if current[ic].ID == param[ip].ID {
				exist = true
				break
			}
		}
		if !exist {
			err6 := obj.PageContent.Remove("id", current[ic].ID)
			if err6 != nil {
				log.Println(err6)
			}
		}
	}

	go page.AddVersion(payload.ID, 0, param[0].PageID)

	c.JSON(http.StatusOK, map[string]interface{}{"message": "edit_page__page_changes_saved"})
}

// saveMedia is used to grab all url media and save them to db
func saveMedia(host, oldContent, newContent string) {
	oldURL := getURLList(host, oldContent)
	newURL := getURLList(host, newContent)
	if len(oldURL) == 0 && len(newURL) == 0 {
		return
	}
	urlList := make(map[string]int)
	for ixo := range oldURL {
		urlList[oldURL[ixo]]--
	}

	for ixn := range newURL {
		urlList[newURL[ixn]]++
	}

	var obj Content
	for key := range urlList {
		if urlList[key] != 0 {
			if err := obj.File.Select("url", key); err == nil {
				obj.File.Used += urlList[key]
				if err = obj.File.Update(); err != nil {
					log.Println(err)
				}
			} else {
				log.Println(err)
			}
		}
	}
}

func getURLList(host, content string) []string {
	regexpPattern := fmt.Sprintf(`%v\/storage\/files\/+\w*\/\S*\.[a-zA-Z0-9]*`, host)
	result := regexp.MustCompile(regexpPattern)
	return result.FindAllString(content, -1)
}
