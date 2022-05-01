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
	"net/http"
	"sort"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/library"
	"framagit.org/InfoLibre/rapido/models/settings"

	"github.com/gin-gonic/gin"
)

// Page content parameter
type PageContentParam struct {
	ID   int    `form:"id"`
	Link string `form:"link"`
}

// Places manifest extension data
type manifestObj struct {
	ContentSettings []map[string]interface{} `json:"content_settings,omitempty"`
}

/**
 * @api {GET} /api/content/page Content On Page
 * @apiDescription Content On Page
 * @apiGroup Content
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiSuccess {Int} id PageContent ID
 * @apiSuccess {Int} page_id Page ID
 * @apiSuccess {Int} content_id Content ID
 * @apiSuccess {String} extension Extension Name
 * @apiSuccess {String} location Location
 * @apiSuccess {Int} column Column
 * @apiSuccess {Int} position Position
 * @apiSuccess {[]Object} json_settings Page Setting
 * @apiSuccess {Object} content Content Object
 * @apiSuccess {Int} content.id Content ID
 * @apiSuccess {String} content.name Content Name
 * @apiSuccess {String} content.content Content
 * @apiSuccess {[]Object} content.json_settings Content Setting
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
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
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} [id] Page ID
 * @apiParam {String} [link] Page Link
 */

// Gets content data on page
func Get(c *gin.Context) {
	var param PageContentParam
	var pageID int
	if err := c.ShouldBind(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}

	var obj Content
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

	pageContent := obj.PageContent.SelectBy("page_id", pageID)
	if len(pageContent) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{"message": "page_not_found"})
		return
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
		for ico := range contents {
			if contents[ico].ID == int(pageContent[index].ContentID.Int64) {
				pageContent[index].Content = contents[ico]
				break
			}
		}

		if pageContent[index].JSONSettings.String != "" {
			jsonSetting, _ := library.Jsonify(pageContent[index].JSONSettings.String)
			pageContent[index].Settings = jsonSetting
		}

	}
	c.JSON(http.StatusOK, pageContent)
}
