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
	"log"
	"net/http"
	"strings"
	"time"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/email"
	"framagit.org/InfoLibre/rapido/models/library"
	"framagit.org/InfoLibre/rapido/models/settings"

	"github.com/gin-gonic/gin"
)

/**
 * @api {POST} /api/register Registers to create a new account
 * @apiDescription Registers to create a new account
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiExample {json} Example usage:
 * {
 *		"password": "abcdefghij"
 *		"pseudonym": "test01",
 *		"email_address": "test@mail.com"
 * }
 * @apiSuccess {String} message Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": "account__account_created"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} password Password
 * @apiParam {String} pseudonym Pseudonym
 * @apiParam {String} email_address Email address
 */

// Registers to create a new account
func Register(c *gin.Context) {
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

	if param.Password != "" {
		param.Password = library.CreateHash(param.Password)
	}
	param.ActivateToken.String, param.ActivateToken.Valid = library.CreateHash(fmt.Sprintf("%v%v%v", time.Now().Unix(), param.Pseudonym, param.Password)), true
	param.AccessLevel = 1
	if err := param.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_insertion_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "account__account_created"})

	data := map[string]string{
		"activation_link": fmt.Sprintf("%v/login?token=%v&id=%v", strings.TrimRight(c.Request.Referer(), "/"), param.ActivateToken.String, param.ID),
		"server_name":     strings.TrimRight(c.Request.Referer(), "/"),
	}
	obj.Account.Select("id", 1)
	dest := map[string]string{
		"to":  param.EmailAddress,
		"bcc": obj.Account.EmailAddress,
	}
	go email.SendMail(
		settings.ActivationEmail(),
		fmt.Sprintf("%v - %v", data["server_name"], "account__account_activation"), // Email subject
		dest,
		data,
	)
	go pageAuthorChange(param.ID, param.EmailAddress)
}

// Removes email address and adds account id in table page_author, for each page user or administrator is author
func pageAuthorChange(id int, email string) {
	var obj Account
	pageCollaboration := obj.PageCollaboration.SelectBy("email", email)
	for index := range pageCollaboration {
		pageCollaboration[index].AccountID = id
		pageCollaboration[index].Email.String, pageCollaboration[index].Email.Valid = "", false
		if err := pageCollaboration[index].Update(); err != nil {
			log.Println(err)
		}
	}
}
