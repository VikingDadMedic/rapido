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
	"net/http"
	"strconv"

	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {DELETE} /api/page/:id/author/:email Remove Page Authors
 * @apiDescription Remove Page Authors
 * @apiGroup Page
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 204 No Content
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Removes an author from a page
func RemoveAuthor(c *gin.Context) {
	id := c.Param("id")
	pageID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})
		return
	}
	var obj Page
	obj.Page.Select("id", pageID)
	if obj.Page.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "page_not_found"})
		return
	}
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	// Tests if logged in person is administrator
	if payload.AccessLevel != jwt.AdminAL {
		// Tests if logged in person is a page author
		obj.PageCollaboration.Select(pageID, payload.ID)
		if obj.PageCollaboration.IsInitial() {
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "settings_pages__cannot_remove_author"})
			return
		}
	}
	registered := true
	email := c.Param("email")
	if err := obj.Account.Select("email_address", email); err != nil {
		registered = false
	}
	authors := obj.PageCollaboration.SelectBy("page_id", pageID)
	authorID := 0
	for index := range authors {
		if registered {
			if authors[index].AccountID == obj.Account.ID {
				authorID = authors[index].ID
			}
		} else {
			if authors[index].Email.Valid && authors[index].Email.String == email {
				authorID = authors[index].ID
			}
		}
	}
	if authorID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, map[string]interface{}{"message": "account__not_found"})
		return
	}

	if err := obj.PageCollaboration.Remove("id", authorID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_deletion_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
