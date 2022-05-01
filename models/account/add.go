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
	"framagit.org/InfoLibre/rapido/models/library"

	"github.com/gin-gonic/gin"
)

/**
 * @api {POST} /api/account Creates an account
 * @apiDescription Creates an account
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Admin
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * {
 *		"password": "abcdefghij"
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
 * @apiSuccess {Int} failed_attempts Failed log in attempts
 * @apiSuccess {Int} lock_until Lock until timestamp
 * @apiSuccess {String} ip_address IP address
 * @apiSuccess {String} activate_token Activate token
 * @apiSuccess {String} reminder_token Reminder token
 * @apiSuccess {Int} reminder_time Reminder timestamp
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 201 Created
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
 * @apiParam {String} password Password
 * @apiParam {String} pseudonym Pseudonym
 * @apiParam {String} email_address Email address
 * @apiParam {Int} access_level Access level
 */

// Creates an account
func Add(c *gin.Context) {
	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__access_forbidden"})
		return
	}
	// Only administrator can create a account
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.AccessLevel != jwt.AdminAL {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__cannot_create_account"})
		return
	}
	var param crud.Account
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	var obj Account
	obj.Account.Select("email_address", param.EmailAddress)
	if !obj.Account.IsInitial() {
		c.AbortWithStatusJSON(http.StatusConflict, map[string]interface{}{"message": "settings_accounts__address_already_registered"})
		return
	}

	// New account is automatically activated
	param.Password = library.CreateHash(param.Password)
	if param.Active == 0 {
		param.Active = 1
	}
	if err := param.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_insertion_failed", "error": err.Error()})
		return
	}
	// Removes email address and adds account id in table page_author, for each page user or administrator is author
	obj.Account.Select("email_address", param.EmailAddress)
	if obj.Account.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "data_insertion_failed"})
		return
	}

	go pageAuthorChange(obj.Account.ID, param.EmailAddress)

	c.JSON(http.StatusCreated, param)
}
