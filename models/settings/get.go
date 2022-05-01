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

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/setting Setting Parameter
 * @apiDescription Setting Parameter
 * @apiGroup Setting
 * @apiVersion 0.1.0
 * @apiPermission Admin
 * @apiSuccess {String} smtp_host SMTP Host For Sending Email
 * @apiSuccess {Int} smtp_port SMTP Host For Sending Email
 * @apiSuccess {String} smtp_user SMTP User For Sending Email
 * @apiSuccess {String} smtp_password SMTP Password For Sending Email
 * @apiSuccess {String} email_sender Email Sender For Sending Email
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
 */

// Gets settings param
func Get(c *gin.Context) {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
		fmt.Print(err)
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, obj.Settings)
}
