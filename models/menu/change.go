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
	"log"
	"net/http"
	"net/url"
	"sort"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {POST} /api/menu Change Menu Item
 * @apiDescription Change Menu Item
 * @apiGroup Menu
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
 * 					"menu_id": 1,
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
 * 					"sub_menu":
 * 					[
 * 						{
 * 							"id": 4,
 * 							"menu_id": 1,
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
 *		"message": "navigation_menu__navigation_menu_changes_saved"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} id Menu ID
 * @apiParam {String} name Menu Name
 * @apiParam {[]Object} item Menu Item
 * @apiParam {Int} item.id Menu Item ID
 * @apiParam {Int} item.menu_id Menu ID (1 for header, 2 for footer, 3 for main)
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
 * @apiParam {[]Object} [item.sub_menu] Sub Menu Item
 * @apiParam {Int} [item.sub_menu.id] Menu Item ID
 * @apiParam {Int} [item.sub_menu.menu_id] Menu ID
 * @apiParam {Int} [item.sub_menu.position] Position
 * @apiParam {Int} [item.sub_menu.level] Level
 * @apiParam {Int} [item.sub_menu.page_id] Page ID
 * @apiParam {String} [item.sub_menu.link_label] Link Label
 * @apiParam {String} [item.sub_menu.link_url] Link URL
 * @apiParam {String} [item.sub_menu.link_target] Link Target
 * @apiParam {Object} [item.sub_menu.page] Item Page
 * @apiParam {Int} [item.sub_menu.page.id] Page ID
 * @apiParam {Int} [item.sub_menu.page.active] Page Active Status
 * @apiParam {String} [item.sub_menu.page.name] Page Name
 * @apiParam {String} [item.sub_menu.page.link] Page Link
 */

// Changes menu composition
func Change(c *gin.Context) {
	var param []*crud.Menu
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	var obj Menu
	currentMenu, err := obj.MenuItem.SelectAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "navigation_menu__menu_selection_failed", "error": err.Error()})
		return
	}

	sort.Sort(crud.MenuByPosition(currentMenu))

	newMenu := make([]crud.MenuItem, 0)
	for im := range param {
		position := 1
		for imi := range param[im].Item {
			param[im].Item[imi].Position = position
			if param[im].Item[imi].LinkURL.String != "" {
				tempURL, _ := url.Parse(param[im].Item[imi].LinkURL.String)
				param[im].Item[imi].LinkURL.String = tempURL.String()
				if param[im].Item[imi].LinkLabel.String == "" {
					param[im].Item[imi].LinkLabel.String, param[im].Item[imi].LinkLabel.Valid = param[im].Item[imi].LinkURL.String, true
				}
			}
			newMenu = append(newMenu, param[im].Item[imi])
			position++
			if len(param[im].Item[imi].SubMenu) > 0 {
				for ism := range param[im].Item[imi].SubMenu {
					temp := crud.MenuItem{}
					temp.ID = param[im].Item[imi].SubMenu[ism].ID
					temp.MenuID = param[im].Item[imi].SubMenu[ism].MenuID
					temp.Position = position
					temp.Level = param[im].Item[imi].SubMenu[ism].Level
					temp.PageID = param[im].Item[imi].SubMenu[ism].PageID
					temp.LinkLabel = param[im].Item[imi].SubMenu[ism].LinkLabel
					temp.LinkURL = param[im].Item[imi].SubMenu[ism].LinkURL
					if temp.LinkURL.String != "" {
						tempsURL, _ := url.Parse(temp.LinkURL.String)
						temp.LinkURL.String = tempsURL.String()
						if param[im].Item[imi].SubMenu[ism].LinkLabel.String == "" {
							temp.LinkLabel.String, temp.LinkLabel.Valid = temp.LinkURL.String, true
						}
					}
					temp.LinkTarget = param[im].Item[imi].SubMenu[ism].LinkTarget
					temp.JSONSettings = param[im].Item[imi].SubMenu[ism].JSONSettings
					temp.AccessLevel = param[im].Item[imi].SubMenu[ism].AccessLevel
					temp.Settings = param[im].Item[imi].SubMenu[ism].Settings
					temp.Page = param[im].Item[imi].SubMenu[ism].Page
					newMenu = append(newMenu, temp)
					position++
				}
			}
		}
	}

	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	for index := range newMenu {
		newMenu[index].UserID = payload.ID
		if newMenu[index].ID > 0 {
			if err := newMenu[index].Update(); err != nil {
				log.Println(err)
			}
		} else {
			if err := newMenu[index].Insert(); err != nil {
				log.Println(err)
			}
		}
	}

	for ic := range currentMenu {
		exist := false
		for in := range newMenu {
			if currentMenu[ic].ID == newMenu[in].ID {
				exist = true
				break
			}
		}

		if !exist {
			if err := obj.MenuItem.Remove("id", currentMenu[ic].ID); err != nil {
				log.Println(err)
			}
		}
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "navigation_menu__navigation_menu_changes_saved"})
}
