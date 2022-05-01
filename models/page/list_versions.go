/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package page

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/page/:page_id/version Page Version List
 * @apiDescription Page Version List
 * @apiGroup Page
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiUse AuthorizeHeader
 * @apiSuccess {Int} id Unique ID Version
 * @apiSuccess {Int} page_id Page ID
 * @apiSuccess {Int} version Page Version Number
 * @apiSuccess {[]Object} data Page Data
 * @apiSuccess {Int} data.id PageContent ID
 * @apiSuccess {Int} data.page_id Page ID
 * @apiSuccess {Int} data.content_id Content ID
 * @apiSuccess {String} data.extension Extension Name
 * @apiSuccess {String} data.location Location
 * @apiSuccess {Int} data.column Column
 * @apiSuccess {Int} data.position Position
 * @apiSuccess {[]Object} data.json_settings Page Setting
 * @apiSuccess {Object} data.content Content Object
 * @apiSuccess {Int} data.content.id Content ID
 * @apiSuccess {String} data.content.name Content Name
 * @apiSuccess {String} data.content.content Content
 * @apiSuccess {[]Object} data.content.json_settings Content Setting
 * @apiSuccess {Int} created_on Version Created Timestamp
 * @apiSuccess {Int} current Current Version Flag
 * @apiSuccess {String} name Name Of User
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * [
 *		{
 *			"id": 1,
 *			"page_id": 1,
 *			"version": 1,
 *			"data":
 *			[
 *				{
 *					"id": 1,
 *					"page_id": 1,
 *					"content_id": 1,
 *					"extension": "test",
 *					"location": "header",
 *					"column": 1,
 *					"position": 1,
 *					"json_settings": [{}],
 *					"content":
 *					{
 *						"id": 1,
 *						"name": "header",
 *						"content": "<h1>Rapido</h1><p>Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.</p>",
 *					}
 *				}
 * 			],
 *			"created_on": "1517677505",
 *			"current": 1,
 *			"name": "admin"
 *		}
 * ]
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets list of page versions
func ListVersions(c *gin.Context) {
	id := c.Param("id")
	pageID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err.Error()})
		return
	}
	var obj Page
	pageVersion := obj.PageVersion.SelectBy("page_id", pageID)
	if len(pageVersion) == 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	ids := make([]interface{}, 0)
	/* for index := range pageVersion {
		ids = append(ids, pageVersion[index].CreatedBy)
	} */
	account := obj.Account.SelectBy("id", ids)
	for ip := range pageVersion {
		for iu := range account {
			// if pageVersion[ip].CreatedBy == account[iu].ID {
			pageVersion[ip].Name = account[iu].Pseudonym
			break
			// }
		}
	}

	c.JSON(http.StatusOK, pageVersion)
}
