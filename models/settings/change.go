/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package settings

import (
	"fmt"
	"net/http"

	"framagit.org/InfoLibre/rapido/models/crud"

	"github.com/gin-gonic/gin"
)

/**
 * @api {PUT} /api/setting Change Setting
 * @apiDescription Change Setting
 * @apiGroup Setting
 * @apiVersion 0.1.0
 * @apiPermission Admin
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * {
 * 		"smtp_host": "-",
 * 		"smtp_port": 1,
 * 		"smtp_user": "-",
 * 		"smtp_password": "-",
 * 		"email_sender": "-",
 * 		"activation_email": "<html><head><title>Activate Account</title></head><body></body></html>",
 * 		"forget_password_email": "<html><head><title>Password Reset</title></head><body></body></html>",
 * 		"name": "My Site",
 * 		"home_page": 1,
 * 		"banned_attemps": 250,
 * 		"lock_time": 1,
 * 		"access_attemps": 5,
 * 		"access_expired": 86400,
 * 		"secret_key": "HRqchYI51YsPEA6E9H0ITwBM4B5F5Nnf",
 * 		"max_size_upload": 50,
 * 		"favicon_url": "",
 *		"logo_url": ""
 * }
 * @apiSuccess {String} smtp_host SMTP Host For Sending Email
 * @apiSuccess {Int} smtp_port SMTP Host For Sending Email
 * @apiSuccess {String} smtp_user SMTP User For Sending Email
 * @apiSuccess {String} smtp_password SMTP Password For Sending Email
 * @apiSuccess {String} email_sender Email Sender For Sending Email
 * @apiSuccess {String} activation_email Activation Email Template
 * @apiSuccess {String} forget_password_email Forget Password Email Template
 * @apiSuccess {String} name Site Name
 * @apiSuccess {Int} home_page Home Page ID
 * @apiSuccess {Int} banned_attemps Login Failure Attemps Before Banned
 * @apiSuccess {Int} lock_time Seconds To Lockout For After Above Failures Detected
 * @apiSuccess {Int} access_attemps Attempts Are Allowed Before Lockout
 * @apiSuccess {Int} access_expired Amount Of Seconds Of Token
 * @apiSuccess {String} secret_key Secret Key For Generate Token
 * @apiSuccess {Int} max_size_upload Maximum Limit For Uploading File in MB
 * @apiSuccess {String} favicon_url Favicon URL
 * @apiSuccess {String} logo_url Logo URL
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 * 		"smtp_host": "-",
 * 		"smtp_port": 1,
 * 		"smtp_user": "-",
 * 		"smtp_password": "-",
 * 		"email_sender": "-",
 * 		"name": "My Site",
 * 		"home_page": 1,
 * 		"banned_attemps": 250,
 * 		"lock_time": 1,
 * 		"access_attemps": 5,
 * 		"access_expired": 86400,
 * 		"secret_key": "HRqchYI51YsPEA6E9H0ITwBM4B5F5Nnf",
 * 		"max_size_upload": 50,
 * 		"favicon_url": "",
 *		"logo_url": ""
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} [smtp_host] SMTP Host For Sending Email
 * @apiParam {Int} [smtp_port] SMTP Host For Sending Email
 * @apiParam {String} [smtp_user] SMTP User For Sending Email
 * @apiParam {String} [smtp_password] SMTP Password For Sending Email
 * @apiParam {String} [email_sender] Email Sender For Sending Email
 * @apiParam {String} [name] Site Name
 * @apiParam {Int} [home_page] Home Page ID
 * @apiParam {Int} [banned_attemps] Login Failure Attemps Before Banned
 * @apiParam {Int} [lock_time] Seconds To Lockout For After Above Failures Detected
 * @apiParam {Int} [access_attemps] Attempts Are Allowed Before Lockout
 * @apiParam {Int} [access_expired] Amount Of Seconds Of Token
 * @apiParam {String} [secret_key] Secret Key For Generate Token
 * @apiParam {Int} [max_size_upload] Maximum Limit For Uploading File in MB
 * @apiParam {String} [favicon_url] Favicon URL
 * @apiParam {String} [logo_url] Logo URL
 */

// Changes page
func Change(c *gin.Context) {
	var param crud.Settings
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}

	if err := param.Update(); err != nil {
		fmt.Print(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_updating_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, param)
}
