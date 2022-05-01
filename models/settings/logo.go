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
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/logo Get Logo And Favicon URL
 * @apiDescription Get Logo And Favicon URL
 * @apiGroup Setting
 * @apiVersion 0.1.0
 * @apiPermission none
 * @apiSuccess {String} favicon_url Favicon URL
 * @apiSuccess {String} logo_url Logo URL
 * @apiSuccess {String} name Name Of Website
 * @apiSuccess {Int} max_size_upload Maximum Limit For Uploading File in MB
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 * 		"favicon_url": "",
 *		"logo_url": "",
 *		"name": "my website",
 *		"max_size_upload": 50
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets logo URL
func Logo(c *gin.Context) {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	result := map[string]interface{}{
		"favicon_url":     obj.Settings.FaviconURL,
		"logo_url":        obj.Settings.LogoURL,
		"name":            obj.Settings.Name,
		"max_size_upload": obj.Settings.MaxSizeUpload,
	}

	c.JSON(http.StatusOK, result)
}
