/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package content

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/content Content list
 * @apiDescription Content list
 * @apiGroup Content
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiSuccess {Int} id Content ID
 * @apiSuccess {Int} name Content name
 * @apiSuccess {String} content Content
 * @apiSuccess {[]Object} json_settings Content settings
 * @apiSuccess {String} pages Page name
 * @apiSuccess {Int} pages.id Page ID
 * @apiSuccess {Int} pages.published Publication status
 * @apiSuccess {String} pages.link URL link
 * @apiSuccess {String} pages.name Page name
 * @apiSuccess {String} pages.title Page title
 * @apiSuccess {String} pages.keywords Page Keywords
 * @apiSuccess {String} pages.descriptions Page description
 * @apiSuccess {[]Object} pages.json_settings Page settings
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * [
 *		{
 *			"id": 1,
 *			"name": "header intro",
 *			"content": "<h1>Rapido</h1><p>Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.</p>",
 *			"json_settings": [{}],
 *			"pages":
 *			[
 *				{
 *					"id": 1,
 *					"published": 0,
 *					"name": "Home",
 *					"link": "",
 *					"title": "Rapido Home Page",
 *					"keywords": "Rapido,home,page",
 *					"descriptions": null,
 *				}
 *			]
 *		}
 * ]
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets content list
func List(c *gin.Context) {
	var obj Content
	content := obj.Content.SelectAll()
	if len(content) == 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	contentIDs := make([]interface{}, 0)
	for index := range content {
		contentIDs = append(contentIDs, content[index].ID)
	}
	pageContent := obj.PageContent.SelectBy("content_id", contentIDs)
	if len(pageContent) > 0 {
		pageIDs := make([]interface{}, 0)
		for index := range pageContent {
			pageIDs = append(pageIDs, pageContent[index].PageID)
		}
		page := obj.Page.SelectBy("id", pageIDs)

		for ipc := range pageContent {
			for ico := range content {
				if pageContent[ipc].ContentID.Int64 == int64(content[ico].ID) {
					for ipa := range page {
						if pageContent[ipc].PageID == page[ipa].ID {
							content[ico].Pages = append(content[ico].Pages, page[ipa])
							break
						}
					}
					break
				}
			}
		}
	}

	c.JSON(http.StatusOK, content)
}
