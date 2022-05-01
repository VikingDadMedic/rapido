/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package menu

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/menu Menu List
 * @apiDescription Menu List
 * @apiGroup Menu
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiSuccess {Int} id Menu ID
 * @apiSuccess {String} name Menu Name
 * @apiSuccess {Int} access_level Access Level Menu
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * [
 *		{
 *			"id": 1,
 *			"name": "header",
 *			"access_level": null
 *		}
 * ]
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets all menu
func List(c *gin.Context) {
	var obj Menu
	output, err := obj.Menu.SelectAll()
	if err != nil || len(output) == 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, output)
}
