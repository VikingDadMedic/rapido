/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package page

import (
	"net/http"

//	"framagit.org/InfoLibre/rapido/models/crud"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/page List of published pages
 * @apiDescription List of published pages
 * @apiGroup Page
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiSuccess {Int} id Page ID
 * @apiSuccess {Int} published Publication status
 * @apiSuccess {String} link URL link
 * @apiSuccess {String} name Page name
 * @apiSuccess {String} title Page title
 * @apiSuccess {String} keywords Page keywords
 * @apiSuccess {String} descriptions Page description
 * @apiSuccess {[]Object} json_settings Page settings
 * @apiSuccess {[]String} authors Authors email address
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * [
 *		{
 *			"id": 2,
 *			"published": 1,
 *			"link": "test",
 *			"name": "Home",
 *			"title": "Rapido Home Page",
 *			"keywords": "Rapido,home,page",
 *			"descriptions": null,
 *			"json_settings": null,
 *			"authors": ["email@example.com"]
 *		}
 * ]
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets list of archived pages
func ListAll(c *gin.Context) {
	var obj Page
	output, err1 := obj.Page.SelectAll()
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})

		return
	}

	c.JSON(http.StatusOK, output)
}
