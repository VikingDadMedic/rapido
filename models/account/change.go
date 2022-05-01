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
	"strconv"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"
	"framagit.org/InfoLibre/rapido/models/library"

	"github.com/gin-gonic/gin"
)

/**
 * @api {PUT} /api/account/:id Changes account settings
 * @apiDescription Changes account settings
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * {
 *		"password": "ab23fecd32efw34efghe4g5tvij"
 *		"pseudonym": "test01",
 *		"email_address": "test@mail.com",
 *		"access_level": 1
 * }
 * @apiSuccess {Int} id Account ID
 * @apiSuccess {String} password Hash password
 * @apiSuccess {Int} active Active flag
 * @apiSuccess {String} pseudonym Pseudonym
 * @apiSuccess {String} email_address Email address
 * @apiSuccess {Int} last_logged_in Last login timestamp
 * @apiSuccess {Int} access_level Access level
 * @apiSuccess {Int} last_accessed Last access timestamp
 * @apiSuccess {Int} failed_attempts Failed login attempts
 * @apiSuccess {Int} lock_until Lock until timestamp
 * @apiSuccess {String} ip_address IP address
 * @apiSuccess {String} activate_token Activate token
 * @apiSuccess {String} reminder_token Reminder token
 * @apiSuccess {Int} reminder_time Reminder timestamp
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"id": 1,
 *		"password": "ab23fecd32efw34efghe4g5tvij"
 *		"active": 1,
 *		"pseudonym": "test01",
 *		"email_address": "test@mail.com",
 *		"last_logged_in": null,
 *		"access_level": 1,
 *		"last_accessed": null,
 *		"failed_attempts": null,
 *		"lock_until": null,
 *		"ip_address": null,
 *		"activate_token": null,
 *		"reminder_token": null,
 *		"reminder_time": null
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} [password] Password
 * @apiParam {String} pseudonym Pseudonym
 * @apiParam {String} email_address Email address
 * @apiParam {Int} access_level Access level
 */

// Changes account settings
func Change(c *gin.Context) {
	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__access_forbidden"})
		return
	}
	var param crud.Account
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	id := c.Param("id")
	accountID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})
		return
	}
	var obj Account
	obj.Account.Select("id", accountID)
	if obj.Account.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "account__not_found"})
		return
	}
	// Only owner of an account can change his account settings. Administrator(s) can change user(s) account settings too.
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.ID != accountID && (obj.Account.AccessLevel != jwt.UserAL || (obj.Account.AccessLevel == jwt.UserAL && payload.AccessLevel != jwt.AdminAL)) {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__cannot_modify_account"})
		return
	}
	// Only administrator can change an account into an administrator account.
	if payload.AccessLevel != jwt.AdminAL && param.AccessLevel == 10 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__cannot_become_administrator"})
		return
	}
	// A super-administrator account can't be changed into a user account.
	if accountID == 1 && param.AccessLevel == 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__superadmin_account_cannot_be_set_as_user"})
		return
	}
	obj.Account.Select("email_address", param.EmailAddress)
	if obj.Account.IsInitial() {
		obj.Account.Select("id", accountID)
		if obj.Account.IsInitial() {
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "account__email_address_not_found"})
			return
		}
	}
	if obj.Account.ID != accountID {
		c.AbortWithStatusJSON(http.StatusConflict, map[string]interface{}{"message": "settings_accounts__address_already_registered"})
		return
	}

	if param.Password != "" {
		param.Password = library.CreateHash(param.Password)
	} else {
		param.Password = obj.Account.Password
	}
	param.ID = accountID
	param.Active, param.LastLoggedIn, param.LastAccessed = obj.Account.Active, obj.Account.LastLoggedIn, obj.Account.LastAccessed
	param.FailedAttemps, param.LockUntil, param.IPAddress = obj.Account.FailedAttemps, obj.Account.LockUntil, obj.Account.IPAddress
	param.ActivateToken, param.ReminderToken, param.ReminderTime = obj.Account.ActivateToken, obj.Account.ReminderToken, obj.Account.ReminderTime
	if err := param.Update(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_updating_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, param)
}
