/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/version System version
 * @apiDescription System version
 * @apiGroup Version
 * @apiVersion 0.1.1
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccess {Int} major Major Number
 * @apiSuccess {Int} minor Minor Number
 * @apiSuccess {Int} patch Patch Number
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"major": 3,
 *		"minor": 4,
 *		"patch": 6
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets system status
func Get(c *gin.Context) {

	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__access_forbidden"})
		return
	}
	var obj Version
	if err := obj.Version.Select(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_selection_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, obj.Version)
}
