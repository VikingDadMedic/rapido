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
	"net/http"

	"framagit.org/InfoLibre/rapido/models/library"

	"github.com/gin-gonic/gin"
)

// param when renewing password process
type NewPass struct {
	ID            int    `json:"id" binding:"required"`
	ReminderToken string `json:"reminder_token" binding:"required"`
	NewPassword   string `json:"new_password" binding:"required"`
}

/**
 * @api {POST} /api/new_password Changes account password after clicking on "Reset password" link
 * @apiDescription Changes account password after clicking on "Reset password" link
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiExample {json} Example usage:
 * {
 *  	"id": 1,
 *  	"reminder_token": "abcdefg123456asdfghjbcd",
 *  	"new_password": "new_pass"
 * }
 * @apiSuccess {String} message Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": "settings_accounts__password_updated"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} id Account ID
 * @apiParam {String} reminder_token Reminder token
 */

// Changes account password after clicking on "Reset password" link
func RenewPass(c *gin.Context) {
	var param NewPass
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	var obj Account
	obj.Account.Select("id", param.ID)
	if obj.Account.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "account__not_found"})
		return
	}
	if obj.Account.ReminderToken.String != param.ReminderToken {
		c.AbortWithStatusJSON(http.StatusConflict, map[string]interface{}{"message": "account__invalid_reset_token"})
		return
	}

	obj.Account.ReminderToken.Valid, obj.Account.ReminderToken.String = false, ""
	obj.Account.Password = library.CreateHash(param.NewPassword)
	if err := obj.Account.Update(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_updating_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "settings_accounts__password_updated"})
}
