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
	"log"
	"net/http"
	"strconv"

	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {PUT} /api/page/:version_id/revert Revert Page
 * @apiDescription Revert Page
 * @apiGroup Page
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccess {String} message Success Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": "revert_page_success"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Reverts page versions
func RevertVersion(c *gin.Context) {
	id := c.Param("id")
	versionID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})
		return
	}
	var obj Page
	if err := obj.PageVersion.Select("id", versionID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "versions__version_not_found", "error": err.Error()})
		return
	}
	obj.Page.Select("id", obj.PageVersion.PageID)
	if obj.Page.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "page_not_found"})
		return
	}
	// Forbidden access if logged user isn't administrator nor page author and if page has author
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.AccessLevel != jwt.AdminAL && obj.PageCollaboration.IsAuthor(obj.PageVersion.PageID, payload.ID) && obj.PageCollaboration.HasAuthor(obj.PageVersion.PageID) {
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{"message": "versions__cannot_revert"})
		return
	}

	ids := make([]int64, 0)
	pageContent := obj.PageContent.SelectBy("page_id", obj.PageVersion.PageID)
	if len(pageContent) > 0 {
		for index := range pageContent {
			ids = append(ids, pageContent[index].ContentID.Int64)
		}
		obj.PageContent.Remove("page_id", obj.PageVersion.PageID)
	}
	if err := obj.PageVersion.Data.Page.Update(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_updating_failed", "error": err.Error()})
		return
	}
	for index := range obj.PageVersion.Data.PageContent {
		exist := false
		for i := 0; i < len(ids); i++ {
			if pageContent[index].ContentID.Int64 == ids[i] {
				exist = true
				break
			}
		}
		if exist {
			obj.PageVersion.Data.PageContent[index].Content.ID = int(pageContent[index].ContentID.Int64)
			obj.Content.Select("id", obj.PageVersion.Data.PageContent[index].Content.ID)
			if err := obj.PageVersion.Data.PageContent[index].Content.Update(); err != nil {
				log.Println(err)
			}
			go saveMedia(c.Request.Host, obj.Content.Content, obj.PageVersion.Data.PageContent[index].Content.Content)
		}
		if err := obj.PageVersion.Data.PageContent[index].Insert(); err != nil {
			log.Println(err)
		}
	}
	go obj.PageVersion.Update()

	c.JSON(http.StatusOK, nil)
}
