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

// ShareParam is used as param to share page
type ShareParam struct {
	Email     string `json:"email,omitempty" binding:"required,max=100"`
	YourEmail string `json:"your_email,omitempty" binding:"omitempty,max=100"`
}

/**
 * @api {POST} /api/page/:id/share Share Page
 * @apiDescription Share Page
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
 *		"message": "message__page_shared"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} email Author email
 */

// Shares page
func Share(c *gin.Context) {
	var param ShareParam
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
	if param.YourEmail == "" && c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err1.Error()})
		return
	}
	var obj Page
	obj.Page.Select("id", pageID)
	if obj.Page.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "page_not_found"})
		return
	}
	// Authentication token
	if c.GetHeader("Authorization") != "" {
		payload, err2 := jwt.Extract(c.GetHeader("Authorization"))
		if err2 == nil {
			obj.Account.Select("id", payload.ID)
			param.YourEmail = obj.Account.EmailAddress
		}
	}
	// Maps metadata to use in template
	data := map[string]string{
		"server_name": strings.TrimRight(c.Request.Referer(), "/"),
		"page_link":   fmt.Sprintf("%v%v", strings.TrimRight(c.Request.Referer(), "/"), obj.Page.Link),
		"user_email":  param.YourEmail,
	}
	// Maps destination email address
	dest := map[string]string{
		"to": param.Email,
	}
	// Calls send email (body = database -> settings / share_page_email)
	go email.SendMail(
		settings.SharePageEmail(),
		fmt.Sprintf(data["server_name"]+"-%v ", param.YourEmail), // Prints server name + sender email as subject
		dest,
		data,
	)

	c.JSON(http.StatusCreated, map[string]interface{}{"message": "share_page__page_shared"})
}
