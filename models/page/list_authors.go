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
	"strconv"

	"framagit.org/InfoLibre/rapido/models/crud"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/page/:id Page Details
 * @apiDescription Page details
 * @apiGroup Page
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiSuccess {Int} id Page ID
 * @apiSuccess {Int} publiched Publication status
 * @apiSuccess {String} link URL link
 * @apiSuccess {String} name Page name
 * @apiSuccess {String} title Page title
 * @apiSuccess {String} keywords Page Keywords
 * @apiSuccess {String} descriptions Page description
 * @apiSuccess {[]Object} json_settings Page settings
 * @apiSuccess {[]String} authors Authors email address
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *  	"id": 2,
 *  	"published": 0,
 *  	"link": "test",
 *  	"name": "Home",
 *  	"title": "Rapido Home Page",
 *  	"keywords": "Rapido,home,page",
 *  	"descriptions": null,
 *  	"json_settings": null,
 *		"authors": ["email@example.com"]
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets list of page authors
func ListAuthors(c *gin.Context) {
	id := c.Param("id")
	pageID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err1.Error()})
		return
	}
	var obj Page
	if err := obj.Page.Select("id", pageID); err != nil || obj.Page.IsInitial() {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	authors := make([]crud.PageCollaboration, 0)
	authors = append(authors, obj.PageCollaboration.SelectBy("page_id", pageID)...)
	if len(authors) > 0 {
		uid := make([]interface{}, 0)
		for index := range authors {
			if authors[index].AccountID > 0 {
				uid = append(uid, authors[index].AccountID)
			} else {
				// Append unactivated accounts
				obj.Page.Authors = append(obj.Page.Authors, authors[index].Email.String)
			}
		}
		accounts := obj.Account.SelectBy("id", uid)
		for index := range accounts {
			// Append activated accounts
			obj.Page.Authors = append(obj.Page.Authors, accounts[index].EmailAddress)
		}
	}

	c.JSON(http.StatusOK, obj.Page)
}
