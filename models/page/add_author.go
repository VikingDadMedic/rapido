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
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"framagit.org/InfoLibre/rapido/models/email"
	"framagit.org/InfoLibre/rapido/models/jwt"
	"framagit.org/InfoLibre/rapido/models/settings"

	"github.com/gin-gonic/gin"
)

// Used as param to add a page author
type AuthorParam struct {
	Email string `json:"email,omitempty" binding:"required,max=100"`
}

/**
 * @api {POST} /api/page/:id/author Add Page Author
 * @apiDescription Add Page Author
 * @apiGroup Page
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * {
 *		"email": "sample@email.com"
 * }
 * @apiSuccess {String} message Success Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 201 Created
 * {
 *		"message": "settings_pages__author_added"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} email Author email
 */

// Adds a page author
func AddAuthor(c *gin.Context) {
	var param AuthorParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
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
	// Access not granted if logged in person isn't administrator AND page has author(s) AND logged in person isn't one of these author(s)
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.AccessLevel != jwt.AdminAL && obj.PageCollaboration.HasAuthor(pageID) && !obj.PageCollaboration.IsAuthor(pageID, payload.ID) {
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "settings_pages__cannot_add_author"})
		return
	}
	if settings.PageAuthor() == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_pages__no_author"})
		return
	}
	// Tests if author is already added and if he didn't activate his account
	if obj.PageCollaboration.Exist(pageID, param.Email) {
		c.AbortWithStatusJSON(http.StatusConflict, map[string]interface{}{"message": "settings_pages__author_already_added"})
		return
	}

	obj.Account.Select("email_address", strings.ToLower(param.Email))
	if obj.Account.IsInitial() {
		obj.PageCollaboration.Email.String, obj.PageCollaboration.Email.Valid = strings.ToLower(param.Email), true
		obj.PageCollaboration.PageID = pageID
	} else {
		obj.PageCollaboration.AccountID, obj.PageCollaboration.PageID = obj.Account.ID, pageID
	}
	if err := obj.PageCollaboration.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_insertion_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{"message": "settings_pages__author_added"})

	// Sends a confirmation email
	data := map[string]string{
		"server_name": strings.TrimRight(c.Request.Referer(), "/"),
		"host":        fmt.Sprintf("%v", strings.TrimRight(c.Request.Referer(), "/")),
		"user_email":  param.Email,
	}
	dest := map[string]string{
		"to": param.Email,
	}
	go email.SendMail(
		settings.PageAuthor(),
		fmt.Sprintf("%v-Author Page Invite", data["server_name"]),
		dest,
		data,
	)
}
