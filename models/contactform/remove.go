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

	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {DELETE} /api/ContactForm/:id Remove ContactForm
 * @apiDescription Remove ContactForm
 * @apiGroup ContactForm
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 204 No ContactForm
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Removes contact form data
func Remove(c *gin.Context) {
	id := c.Param("id")
	ContactFormID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})
		return
	}

	var obj ContactForm
	obj.ContactForm.Select("id", ContactFormID)
	if obj.ContactForm.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_contents__contact_form_not_found"})
		return
	}

	tmpPC := obj.PageContent.SelectBy("id", ContactFormID)
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

	if err := obj.ContactForm.Remove("id", ContactFormID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_deletion_failed", "error": err.Error()})
		return
	}

	if err := obj.PageContent.Remove("id", ContactFormID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_deletion_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
