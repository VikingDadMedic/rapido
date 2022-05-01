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

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/account Gets accounts list
 * @apiDescription Gets accounts list
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Admin
 * @apiUse AuthorizeHeader
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
 * [
 * 		{
 *			"id": 1,
 *			"password": "ab23fecd32efw34efghe4g5tvij"
 *			"active": 1,
 *			"pseudonym": "test01",
 *			"email_address": "test@mail.com",
 *			"last_logged_in": null,
 *			"access_level": 1,
 *			"last_accessed": null,
 *			"failed_attempts": null,
 *			"lock_until": null,
 *			"ip_address": null,
 *			"activate_token": null,
 *			"reminder_token": null,
 *			"reminder_time": null
 * 		}
 * ]
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets accounts list
func List(c *gin.Context) {
	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__access_forbidden"})
		return
	}

	// Only owner of an account can see his account, administrator(s) can see users accounts too.
	var obj Account
	output := make([]crud.Account, 0)
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.AccessLevel == jwt.AdminAL {
		output = append(output, obj.Account.SelectBy("access_level", jwt.UserAL)...)
		output = append(output, obj.Account.SelectBy("id", payload.ID)...)
	} else if payload.AccessLevel == jwt.UserAL {
		output = append(output, obj.Account.SelectBy("id", payload.ID)...)
	} else {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}
	if len(output) == 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, output)
}
