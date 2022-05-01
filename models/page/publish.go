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
	"log"
	"net/http"
	"strconv"

	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {PUT} /api/page/:archive_id/publish Publish page
 * @apiDescription Publish page
 * @apiGroup Page
 * @apiVersion 0.1.0
 * @apiPermission Admin
 * @apiUse AuthorizeHeader
 * @apiSuccess {String} message Success Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": "settings_pages__page_published"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Publishes page
func Publish(c *gin.Context) {
	id := c.Param("id")
	pageID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})
		return
	}
	var obj Page
	// Access not granted if logged user isn't administrator AND page has author(s) AND user isn't one of these author(s)
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.AccessLevel != jwt.AdminAL && obj.PageCollaboration.HasAuthor(pageID) && !obj.PageCollaboration.IsAuthor(pageID, payload.ID) {
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "settings_pages__cannot_publish_page"})
		return
	}
	if err := obj.Page.Select("id", pageID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "page_not_found", "error": err.Error()})
		return
	}

	// Publishes page
	obj.Page.Published = 1
	if err := obj.Page.Update(); err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "settings_pages__page_published"})
}
