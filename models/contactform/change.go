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
	"strconv"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {PUT} /api/ContactForm/:id Change ContactForm
 * @apiDescription Change ContactForm
 * @apiGroup ContactForm
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * {
 *  	"name": "Title",
 *  	"ContactForm": "<h1>Rapido</h1>",
 *  	"json_settings": null,
 * }
 * @apiSuccess {Int} id ContactForm ID
 * @apiSuccess {String} name ContactForm Name
 * @apiSuccess {String} ContactForm ContactForm
 * @apiSuccess {[]Object} json_settings ContactForm Setting
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *  	"id": 1,
 *  	"name": "Title",
 *  	"ContactForm": "<h1>Rapido</h1>",
 *  	"json_settings": null,
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} name ContactForm Name
 * @apiParam {String} ContactForm ContactForm
 * @apiParam {[]Object} [json_settings] ContactForm Setting
 */

// Changes contact form
func Change(c *gin.Context) {
	var obj ContactForm
	var param crud.ContactForm
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}

	id := c.Param("id")
	ContactFormID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})
		return
	}

	obj.ContactForm.Select("id", ContactFormID)
	if obj.ContactForm.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_contents__contact_form_not_found"})
		return
	}

	tmpPC := obj.PageContent.SelectBy("contactform_id", ContactFormID)
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if len(tmpPC) > 0 {
		pageIDs := make([]interface{}, 0)
		for ipc := range tmpPC {
			pageIDs = append(pageIDs, tmpPC[ipc].PageID)
		}
		if len(pageIDs) > 0 {
			tmpPg := obj.Page.SelectBy("id", pageIDs)
			for ipg := range tmpPg {
				obj.PageCollaboration.Select(tmpPg[ipg].ID, payload.ID)
				if obj.PageCollaboration.IsInitial() {
					c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "settings_contents__cannot_edit_contact_form"})
					return
				}
			}
		}
	}

	obj.ContactForm.Select("id", ContactFormID)
	param.ID = ContactFormID

	if err := param.Update(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_updating_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, param)
}
