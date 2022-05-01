/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package page

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {POST} /api/page/:id/copy Copy Page
 * @apiDescription Copy Page
 * @apiGroup Page
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * {
 *  	"published": 0,
 *  	"link": "test",
 *  	"name": "Home",
 *  	"title": "Rapido Home Page",
 *  	"keywords": "Rapido,home,page",
 *  	"descriptions": null,
 *  	"json_settings": null,
 * }
 * @apiSuccess {Int} id Page ID
 * @apiSuccess {Int} published Publication status
 * @apiSuccess {String} link URL link
 * @apiSuccess {String} name Page name
 * @apiSuccess {String} title Page title
 * @apiSuccess {String} keywords Page keywords
 * @apiSuccess {String} descriptions Page description
 * @apiSuccess {[]Object} json_settings Page settings
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 201 Created
 * {
 *  	"id": 2,
 *  	"published": 0,
 *  	"link": "test",
 *  	"name": "Home",
 *  	"title": "Rapido Home Page",
 *  	"keywords": "Rapido,home,page",
 *  	"descriptions": null,
 *  	"json_settings": null,
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} [published] Publication status
 * @apiParam {String} [link] URL link
 * @apiParam {String} name Page name
 * @apiParam {String} [title] Page title
 * @apiParam {String} [keywords] Page keywords
 * @apiParam {String} [descriptions] Page description
 * @apiParam {[]Object} [json_settings] Page settings
 */

// Copy is used to duplicate page
func Copy(c *gin.Context) {
	var param crud.Page
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	id := c.Param("id")
	pageID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err.Error()})
		return
	}

	if param.Published == 1 {
		param.Published = 0
	}
	if param.Title == "" {
		param.Title = param.Name
	}
	param.Link = strings.ToLower(param.Name)
	// Use custom link encoding function to replace all special characters
	param.Link = CustomLinkEncoding(param.Link)
	var obj Page
	obj.Page.Select("link", param.Link)
	if !obj.Page.IsInitial() {
		c.AbortWithStatusJSON(http.StatusConflict, map[string]interface{}{"message": "create_page__link_already_exists"})
		return
	}

	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if err = param.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_insertion_failed", "error": err.Error()})
		return
	}
	host := c.Request.Host

	pageContent := obj.PageContent.SelectBy("page_id", pageID)
	for index := range pageContent {
		obj.Content.Select("id", pageContent[index].ContentID)
		if !obj.Content.IsInitial() {
			obj.Content.Insert()
			go saveMedia(host, "", obj.Content.Content)
			pageContent[index].ContentID.Int64, pageContent[index].ContentID.Valid = int64(obj.Content.ID), true
		}
		pageContent[index].PageID = param.ID
		pageContent[index].Insert()
	}

	go AddVersion(payload.ID, 0, param.ID)

	c.JSON(http.StatusCreated, param)
}

// Grabs all url media and saves them in db
func saveMedia(host, oldContent, newContent string) {
	oldURL := getURLList(host, oldContent)
	newURL := getURLList(host, newContent)
	urlList := make(map[string]int)
	for ixo := range oldURL {
		urlList[oldURL[ixo]]--
	}
	for ixn := range newURL {
		urlList[newURL[ixn]]++
	}
	var obj Page
	for key := range urlList {
		if urlList[key] != 0 {
			if err := obj.File.Select("url", key); err == nil {
				obj.File.Used += urlList[key]
				if err = obj.File.Update(); err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func getURLList(host, content string) []string {
	regexpPattern := fmt.Sprintf(`(%v\/storage\/files\/)[a-z0-9]+\/.+`, host)
	result := regexp.MustCompile(regexpPattern)
	return result.FindAllString(content, -1)
}
