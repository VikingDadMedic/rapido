/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package contactform

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/contactform contactform List
 * @apiDescription contactform List
 * @apiGroup contactform
 * @apiVersion 0.1.0
 * @apiPermission Public
 * @apiSuccess {Int} id Contact form ID
 * @apiSuccess {Int} name Contact form name
 * @apiSuccess {String} contactform Contact form
 * @apiSuccess {[]Object} json_settings Contact form settings
 * @apiSuccess {String} pages Page name
 * @apiSuccess {Int} pages.id Page ID
 * @apiSuccess {Int} pages.published Publication status
 * @apiSuccess {String} pages.link URL link
 * @apiSuccess {String} pages.name Page name
 * @apiSuccess {String} pages.title Page title
 * @apiSuccess {String} pages.keywords Page keywords
 * @apiSuccess {String} pages.descriptions Page description
 * @apiSuccess {[]Object} pages.json_settings Page settings
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * [
 *		{
 *			"id": 1,
 *			"name": "header intro",
 *			"contactform": "<h1>Rapido</h1><p>Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.</p>",
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

// Gets list of contact forms
func List(c *gin.Context) {
	var obj ContactForm
	content := obj.ContactForm.SelectAll()
	if len(content) == 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	contactformIDs := make([]interface{}, 0)
	for index := range content {
		contactformIDs = append(contactformIDs, content[index].ID)
	}
	pagecontent := obj.PageContent.SelectBy("id", contactformIDs)
	if len(pagecontent) > 0 {
		pageIDs := make([]interface{}, 0)
		for index := range pagecontent {
			pageIDs = append(pageIDs, pagecontent[index].PageID)
		}
		page := obj.Page.SelectBy("id", pageIDs)

		for ipc := range pagecontent {
			for ico := range content {
				if pagecontent[ipc].ContactFormID.Int64 == int64(content[ico].ID) {
					for ipa := range page {
						if pagecontent[ipc].PageID == page[ipa].ID {
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
