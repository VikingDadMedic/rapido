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
	"time"

	"framagit.org/InfoLibre/rapido/models/email"
	"framagit.org/InfoLibre/rapido/models/library"
	"framagit.org/InfoLibre/rapido/models/settings"

	"github.com/gin-gonic/gin"
)

/**
 * @api {POST} /api/forget_password Password reset procedure
 * @apiDescription Password reset procedure
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiExample {json} Example usage:
 * {
 *		"email_address": "test@mail.com"
 * }
 * @apiSuccess {String} message Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": "account__password_reset"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} email_address Email address
 */

// parameter for forget password process
type ForgetPasswordParam struct {
	EmailAddress string `json:"email_address" binding:"required"`
}

// Ask email address for password reset
func ForgetPassword(c *gin.Context) {
	var param ForgetPasswordParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	var obj Account
	obj.Account.Select("email_address", param.EmailAddress)
	if obj.Account.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "account__email_address_not_found"})
		return
	}
	if obj.Account.ReminderTime.Int64 > time.Now().Unix()-3600 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "account__request_for_password_reset"})
		return
	}

	obj.Account.ReminderTime.Int64, obj.Account.ReminderTime.Valid = time.Now().Unix(), true
	obj.Account.ReminderToken.String, obj.Account.ReminderToken.Valid = library.CreateHash(fmt.Sprintf("%v%v%v", obj.Account.ReminderTime.Int64, obj.Account.Pseudonym, obj.Account.Password)), true
	if err := obj.Account.Update(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_updating_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "account__password_reset_link_sent"})

	// Sends an email with a password reset link
	data := map[string]string{
		"forgot_password_link": fmt.Sprintf("%v/forgot_password?token=%v&id=%v", strings.TrimRight(c.Request.Referer(), "/"), obj.Account.ReminderToken.String, obj.Account.ID),
		"server_name":          strings.TrimRight(c.Request.Referer(), "/"),
		"user_email":           param.EmailAddress,
	}
	obj.Account.Select("id", 1)
	dest := map[string]string{
		"to":  param.EmailAddress,
		"bcc": obj.Account.EmailAddress,
	}
	go email.SendMail(
		settings.ForgetPasswordEmail(),
		fmt.Sprintf("%v - %v", data["server_name"], "account__password_reset"), // Email subject
		dest,
		data,
	)
}
