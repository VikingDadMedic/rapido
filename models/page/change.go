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
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {PUT} /api/page/:id Change Page
 * @apiDescription Change Page
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
 * HTTP/1.1 200 Ok
 * {
 *  	"id": 1,
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
 * @apiParam {Int} published Publication status
 * @apiParam {String} link URL link
 * @apiParam {String} name Page name
 * @apiParam {String} title Page title
 * @apiParam {String} [keywords] Page keywords
 * @apiParam {String} [descriptions] Page description
 * @apiParam {[]Object} [json_settings] Page settings
 */

// Saves page changes
func Change(c *gin.Context) {
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
	var obj Page
	obj.Page.Select("id", pageID)
	if obj.Page.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "page_not_found"})
		return
	}
	// Access not granted if logged user isn't administrator AND page has author(s) AND user isn't one of these author(s)
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.AccessLevel != jwt.AdminAL && obj.PageCollaboration.HasAuthor(pageID) && !obj.PageCollaboration.IsAuthor(pageID, payload.ID) {
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "edit_page__cannot_edit_page"})
		return
	}
	param.ID = pageID
	if param.Title == "" {
		param.Title = param.Name
	}
	if param.Link == "" {
		param.Link = strings.ToLower(param.Name)
	}
	param.Link = url.PathEscape(param.Link)

	if err = param.Update(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_updating_failed", "error": err.Error()})
		return
	}
	go AddVersion(payload.ID, 0, param.ID)

	c.JSON(http.StatusOK, param)
}
