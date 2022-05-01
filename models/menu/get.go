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
	"sort"

	"framagit.org/InfoLibre/rapido/models/crud"
	//"framagit.org/InfoLibre/rapido/models/library"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/menu_item Menu Item
 * @apiDescription Menu Item
 * @apiGroup Menu
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiSuccess {Int} id Menu ID
 * @apiSuccess {String} name Menu Name
 * @apiSuccess {[]Object} item Menu Item
 * @apiSuccess {Int} item.id Menu Item ID
 * @apiSuccess {Int} item.menu_id Menu ID
 * @apiSuccess {Int} item.position Position
 * @apiSuccess {Int} item.level Level
 * @apiSuccess {Int} item.page_id Page ID
 * @apiSuccess {String} item.link_label Link Label
 * @apiSuccess {String} item.link_url Link URL
 * @apiSuccess {String} item.link_target Link Target
 * @apiSuccess {Object} item.page Item Page
 * @apiSuccess {Int} item.page.id Page ID
 * @apiSuccess {Int} item.page.active Page Active Status
 * @apiSuccess {String} item.page.name Page Name
 * @apiSuccess {String} item.page.link Page Link
 * @apiSuccess {[]Object} item.sub_menu Sub Menu Item
 * @apiSuccess {Int} item.sub_menu.id Menu Item ID
 * @apiSuccess {Int} item.sub_menu.menu_id Menu ID
 * @apiSuccess {Int} item.sub_menu.position Position
 * @apiSuccess {Int} item.sub_menu.level Level
 * @apiSuccess {Int} item.sub_menu.page_id Page ID
 * @apiSuccess {String} item.sub_menu.link_label Link Label
 * @apiSuccess {String} item.sub_menu.link_url Link URL
 * @apiSuccess {String} item.sub_menu.link_target Link Target
 * @apiSuccess {Object} item.sub_menu.page Item Page
 * @apiSuccess {Int} item.sub_menu.page.id Page ID
 * @apiSuccess {Int} item.sub_menu.page.active Page Active Status
 * @apiSuccess {String} item.sub_menu.page.name Page Name
 * @apiSuccess {String} item.sub_menu.page.link Page Link
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
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
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets menu data
func Get(c *gin.Context) {
	// var payload library.Payload
	// cookie, err := c.Request.Cookie("token")
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusNoContent, nil)
	// 	return
	// }

	// payload, err = library.Extract(fmt.Sprintf("Bearer %v", cookie.Value))
	//payload, _ := library.Extract(c.GetHeader("Authorization"))
	var obj Menu
	menuItems, err1 := obj.MenuItem.SelectAll()
	if err1 != nil || len(menuItems) == 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	menuIDs := make([]interface{}, 0)
	pageIDs := make([]interface{}, 0)
	for index := range menuItems {
		if menuItems[index].MenuID > 0 {
			menuIDs = append(menuIDs, menuItems[index].MenuID)
		}
		if menuItems[index].PageID.Int64 > 0 {
			pageIDs = append(pageIDs, menuItems[index].PageID.Int64)
		}
	}
	menu := obj.Menu.SelectBy("id", menuIDs)
	page := obj.Page.SelectBy("id", pageIDs)
	// cookie, err := c.Request.Cookie("token")
	// if err == nil {
	// 	payload, err = library.Extract(fmt.Sprintf("Bearer %v", cookie.Value))
	// 	if err == nil {
	// 		pageCollaboration = obj.PageCollaboration.SelectBy("user_id", payload.ID)
	// 	}
	// }
	sort.Sort(crud.MenuByPosition(menuItems))

	for im := range menu {
		for imi := range menuItems {
			if menu[im].ID != menuItems[imi].MenuID {
				continue
			}

			if menuItems[imi].PageID.Int64 > 0 {
				for ip := range page {
					if menuItems[imi].PageID.Int64 == int64(page[ip].ID) {
						// if page[ip].AccessLevel.Int64 == 1 && payload.AccessLevel <= 1 {
						// 	if len(pageCollaboration) > 0 {
						// 		for ipc := range pageCollaboration {
						// 			if pageCollaboration[ipc].PageID == page[ip].ID {
						// 				menuItems[imi].Page = page[ip]
						// 				break
						// 			}
						// 		}
						// 	}
						// } else {
						menuItems[imi].Page = page[ip]
						// }
						break
					}
				}
			}
			if menuItems[imi].Level == 1 {
				menu[im].Item = append(menu[im].Item, menuItems[imi])
			} else {
				parent := 0
				if len(menu[im].Item) > 0 {
					parent = len(menu[im].Item) - 1
				}
				menu[im].Item[parent].SubMenu = append(menu[im].Item[parent].SubMenu, &menuItems[imi])
			}
		}
	}

	c.JSON(http.StatusOK, menu)
}
