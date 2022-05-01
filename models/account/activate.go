/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package account

import (
	"fmt"
	"net/http"
	"strings"

	"framagit.org/InfoLibre/rapido/models/email"
	"framagit.org/InfoLibre/rapido/models/settings"

	"github.com/gin-gonic/gin"
)

// param when activating account process
type ActivateAccount struct {
	ID            int    `form:"id" binding:"required"`
	ActivateToken string `form:"activate_token" binding:"required"`
}

/**
 * @api {GET} /api/activate_account Activates created account
 * @apiDescription Activates created account
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiExample {form} Example usage:
 * {
 *  	"id": 1,
 *  	"active_token": "abcdefg123456asdfghjbcd"
 * }
 * @apiSuccess {String} message Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": "account__account_activated"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} id Account ID
 * @apiParam {String} active_token Active token
 */

// Activates created account
func Activate(c *gin.Context) {
	var param ActivateAccount
	if err := c.ShouldBind(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	var obj Account
	obj.Account.Select("id", param.ID)
	if obj.Account.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "account__not_found"})
		return
	}
	if obj.Account.ActivateToken.String != param.ActivateToken {
		c.AbortWithStatusJSON(http.StatusConflict, map[string]interface{}{"message": "account__invalid_activation_token"})
		return
	}

	obj.Account.ActivateToken.Valid, obj.Account.ActivateToken.String = false, ""
	obj.Account.Active = 1
	if err := obj.Account.Update(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_updating_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "account__account_activated"})

	data := map[string]string{
		"email":       obj.Account.EmailAddress,
		"server_name": strings.TrimRight(c.Request.Referer(), "/"),
	}
	obj.Account.Select("id", 1)
	go email.SendMail(
		settings.ActivatedEmail(),
		fmt.Sprintf("%v - %v", data["server_name"], "account__account_activated"), // Email subject
		map[string]string{
			"to": obj.Account.EmailAddress,
		},
		data,
	)
}
