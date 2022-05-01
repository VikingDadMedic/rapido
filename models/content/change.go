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
	"strconv"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"
	"framagit.org/InfoLibre/rapido/models/page"

	"github.com/gin-gonic/gin"
)

/**
 * @api {PUT} /api/content/:id Change Content
 * @apiDescription Change Content
 * @apiGroup Content
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * {
 *  	"name": "Title",
 *  	"content": "<h1>Rapido</h1>",
 *  	"json_settings": null,
 * }
 * @apiSuccess {Int} id Content ID
 * @apiSuccess {String} name Content Name
 * @apiSuccess {String} content Content
 * @apiSuccess {[]Object} json_settings Content Setting
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *  	"id": 1,
 *  	"name": "Title",
 *  	"content": "<h1>Rapido</h1>",
 *  	"json_settings": null,
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} name Content Name
 * @apiParam {String} content Content
 * @apiParam {[]Object} [json_settings] Content Setting
 */

// Changes content
func Change(c *gin.Context) {
	var param crud.Content
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	id := c.Param("id")
	contentID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})
		return
	}
	var obj Content
	obj.Content.Select("id", contentID)
	if obj.Content.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_contents__content_not_found"})
		return
	}

	tmpPC := obj.PageContent.SelectBy("content_id", contentID)
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if len(tmpPC) > 0 {
		pageIDs := make([]interface{}, 0)
		for ipc := range tmpPC {
			pageIDs = append(pageIDs, tmpPC[ipc].PageID)
		}
		if len(pageIDs) > 0 {
			tmpPg := obj.Page.SelectBy("id", pageIDs)
			for ipg := range tmpPg {
				obj.PageCollaboration.Select(tmpPg[ipg].ID, payload.ID)
				if obj.PageCollaboration.IsInitial() {
					c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "settings_contents__cannot_edit_content"})
					return
				}
			}
		}
	}
	obj.Content.Select("id", contentID)
	param.ID = contentID
	err1 = param.Update()
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_updating_failed", "error": err1.Error()})
		return
	}

	go saveMedia(c.Request.Host, obj.Content.Content, param.Content)
	go page.AddVersion(payload.ID, param.ID, 0)

	c.JSON(http.StatusOK, param)
}
