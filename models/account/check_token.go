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
	"time"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/check_token Checks access token
 * @apiDescription Checks access token
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccess {Int} iat Time created token
 * @apiSuccess {Int} exp Time expired token
 * @apiSuccess {String} iss Issuer token
 * @apiSuccess {String} sub Subject token
 * @apiSuccess {Int} id Account ID
 * @apiSuccess {String} pseudonym Pseudonym
 * @apiSuccess {String} email_address Email address
 * @apiSuccess {Int} last_logged_in Last login timestamp
 * @apiSuccess {Int} access_level Access Level
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"iat": 12345678,
 *		"exp": 12345923,
 *		"iss": "rapidohandle.com",
 *		"sub": "access granted for 24 hour(s)",
 *		"id": 1,
 *		"pseudonym": "admin",
 *		"email_address": "admin@admin",
 *		"last_logged_in": 12345678,
 *		"access_level": 10
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Checks access token
func CheckToken(c *gin.Context) {
	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__access_forbidden"})
		return
	}
	var obj Account
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	obj.Account.Select("id", payload.ID)
	if !obj.Account.IsInitial() {
		go func(account crud.Account) {
			account.LastAccessed.Int64, account.LastAccessed.Valid = time.Now().Unix(), true
			account.Update()
		}(obj.Account)
	}

	c.JSON(http.StatusOK, payload)
}
