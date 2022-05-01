/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package theme

import (
	"log"
	"net/http"

	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {POST} /api/Theme Change Theme Item
 * @apiDescription Change Theme Item
 * @apiGroup Theme
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * [
 * 		{
 * 			"id": 1,
 * 			"name": "header",
 * 			"item":
 * 			[
 * 				{
 * 					"id": 1,
 * 					"Theme_id": 1,
 * 					"position": 1,
 * 					"level": 1,
 * 					"page_id": 1,
 * 					"link_label": "",
 * 					"link_url": "",
 * 					"link_target": "",
 * 					"page":
 * 					{
 * 						"id": 1,
 * 						"active": 1,
 * 						"name": "Home",
 * 						"link": "/index"
 * 					},
 * 					"sub_Theme":
 * 					[
 * 						{
 * 							"id": 4,
 * 							"Theme_id": 1,
 * 							"position": 1,
 * 							"level": 2,
 * 							"page_id": 1,
 * 							"link_label": "",
 * 							"link_url": "",
 * 							"link_target": "",
 * 							"page":
 * 							{
 * 								"id": 1,
 * 								"active": 1,
 * 								"name": "Home",
 * 								"link": "/index"
 * 							}
 * 						}
 * 					]
 * 				}
 * 			]
 * 		}
 * ]
 * @apiSuccess {String} message Success Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": "settings_features__theme_changed"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} id Theme ID
 * @apiParam {String} name Theme Name
 * @apiParam {[]Object} item Theme Item
 * @apiParam {Int} item.id Theme Item ID
 * @apiParam {Int} item.Theme_id Theme ID (1 for header, 2 for footer, 3 for main)
 * @apiParam {Int} item.position Position
 * @apiParam {Int} item.level Level (0 for root, 1 for child)
 * @apiParam {Int} item.page_id Page ID
 * @apiParam {String} item.link_label Link Label
 * @apiParam {String} item.link_url Link URL
 * @apiParam {String} item.link_target Link Target
 * @apiParam {Object} item.page Item Page
 * @apiParam {Int} item.page.id Page ID
 * @apiParam {Int} item.page.active Page Active Status
 * @apiParam {String} item.page.name Page Name
 * @apiParam {String} item.page.link Page Link
 * @apiParam {[]Object} [item.sub_Theme] Sub Theme Item
 * @apiParam {Int} [item.sub_Theme.id] Theme Item ID
 * @apiParam {Int} [item.sub_Theme.Theme_id] Theme ID
 * @apiParam {Int} [item.sub_Theme.position] Position
 * @apiParam {Int} [item.sub_Theme.level] Level
 * @apiParam {Int} [item.sub_Theme.page_id] Page ID
 * @apiParam {String} [item.sub_Theme.link_label] Link Label
 * @apiParam {String} [item.sub_Theme.link_url] Link URL
 * @apiParam {String} [item.sub_Theme.link_target] Link Target
 * @apiParam {Object} [item.sub_Theme.page] Item Page
 * @apiParam {Int} [item.sub_Theme.page.id] Page ID
 * @apiParam {Int} [item.sub_Theme.page.active] Page Active Status
 * @apiParam {String} [item.sub_Theme.page.name] Page Name
 * @apiParam {String} [item.sub_Theme.page.link] Page Link
 */

// Changes theme
func Change(c *gin.Context) {
	var obj Theme
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	currentTheme, err1 := obj.Theme.SelectAll(payload.ID)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "settings_features__theme_failed", "error": err1.Error()})
		return
	}
	if err := currentTheme[0].Update(c.Query("new")); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "settings_features__theme_changed"})
}
