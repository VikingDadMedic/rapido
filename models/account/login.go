/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>,Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package account

import (
	"net/http"
	"time"

	"framagit.org/InfoLibre/rapido/models/jwt"
	"framagit.org/InfoLibre/rapido/models/library"
	"framagit.org/InfoLibre/rapido/models/settings"

	"github.com/gin-gonic/gin"
)

// LoginParam used as parameter for login process
type LoginParam struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Remember bool   `json:"remember"`
}

/**
 * @api {POST} /api/login Logs into application
 * @apiDescription Logs into application
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiExample {json} Example usage:
 * {
 *		"username": "user@mail.com",
 *		"password": "abcdefghij",
 *		"remember": true
 * }
 * @apiSuccess {String} token Access token
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 202 Accepted
 * {
 *		"token": "ab23fecd32efw34efghe4g5tvij.ajfsdhods8sd2323.dsjfhsd9fbwgvwf27fevwcx"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} username User name
 * @apiParam {String} password Password
 * @apiParam {Bool} remember Remember flag
 */

// Logs into application
func Login(c *gin.Context) {
	var param LoginParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	var obj Account
	if settings.BannedAttemps() > 0 {
		obj.Banned.Select("ip_address", c.ClientIP())
		if !obj.Banned.IsInitial() {
			if obj.Banned.UserAgent.String == c.Request.UserAgent() {
				c.AbortWithStatusJSON(http.StatusNotAcceptable, map[string]interface{}{"message": "account__ip_banned"})
				return
			}
		}
	}
	obj.Account.Select("email_address", param.UserName)
	if obj.Account.IsInitial() {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, map[string]interface{}{"message": "account__email_address_not_found"})
		return
	}
	if obj.Account.LockUntil.Valid && obj.Account.LockUntil.Int64 > time.Now().Unix() {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, map[string]interface{}{"message": "account__account_temporarily_blocked", "time_left": obj.Account.LockUntil.Int64 - time.Now().Unix()})
		return
	}
	if obj.Account.Active == 0 {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, map[string]interface{}{"message": "account__account_not_activated"})
		return
	}
	if library.CheckPasswordHash(param.Password, obj.Account.Password) != true {
		obj.Account.FailedAttemps.Int64++
		obj.Account.FailedAttemps.Valid = true
		if obj.Account.FailedAttemps.Int64%settings.BannedAttemps() == 0 {
//		if obj.Account.FailedAttemps%settings.BannedAttemps() == 0 {
			obj.Account.LockUntil.Int64, obj.Account.LockUntil.Valid = time.Now().Unix()+settings.LockTime(), true
			obj.Account.Update()
		}
//		if obj.Account.FailedAttemps >= settings.BannedAttemps() {
		if obj.Account.FailedAttemps.Int64 >= settings.BannedAttemps() {
			obj.Banned.IPAddress = c.ClientIP()
			obj.Banned.UserAgent.String, obj.Banned.UserAgent.Valid = c.Request.UserAgent(), true
			obj.Banned.Insert()
		}
		c.AbortWithStatusJSON(http.StatusNotAcceptable, map[string]interface{}{"message": "account__wrong_password"})
		return
	}

	lastLogged := time.Now().Unix()
	obj.Account.LastLoggedIn.Int64, obj.Account.LastLoggedIn.Valid = lastLogged, true
	obj.Account.LastAccessed.Int64, obj.Account.LastAccessed.Valid = lastLogged, true
	obj.Account.IPAddress.String, obj.Account.IPAddress.Valid = c.ClientIP(), true
	obj.Account.Update()
	expired := int64(0)
//	expired := 0
	if param.Remember {
		expired = 168
	} else {
		expired = settings.AccessExpired()
	}
	token, err1 := jwt.GenerateToken(map[string]interface{}{
		"id":             obj.Account.ID,
		"pseudonym":      obj.Account.Pseudonym,
		"email_address":  obj.Account.EmailAddress,
		"last_logged_in": lastLogged,
		"access_level":   obj.AccessLevel,
	}, time.Duration(expired), c.Request.Host)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "account__token_creation_failed", "error": err1.Error()})
		return
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{"token": token})
}
