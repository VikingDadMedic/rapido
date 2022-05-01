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
	"strconv"

	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {DELETE} /api/content/:id Remove Content
 * @apiDescription Remove Content
 * @apiGroup Content
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 204 No Content
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Removes content data
func Remove(c *gin.Context) {
	id := c.Param("id")
	contentID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})
		return
	}
	var obj Content
	obj.Content.Select("id", contentID)
	if obj.Content.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_contents__content_not_found"})
		return
	}

	tmpPC := obj.PageContent.SelectBy("content_id", contentID)
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
					c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "settings_contents__cannot_delete_content"})
					return
				}
			}
		}
	}

	if err := obj.Content.Remove("id", contentID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_deletion_failed", "error": err.Error()})
		return
	}

	if err := obj.PageContent.Remove("content_id", contentID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_deletion_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
