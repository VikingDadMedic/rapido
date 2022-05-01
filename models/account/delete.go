/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package account

import (
	"net/http"
	"strconv"

	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

/**
 * @api {DELETE} /api/account/:id Deletes account
 * @apiDescription Deletes account
 * @apiGroup Account
 * @apiVersion 0.1.0
 * @apiPermission Admin
 * @apiUse AuthorizeHeader
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 204 No Content
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Deletes account
func Remove(c *gin.Context) {
	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__access_forbidden"})
		return
	}
	id := c.Param("id")
	accountID, err1 := strconv.Atoi(id)
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "id_is_not_a_number", "error": err1.Error()})
		return
	}
	var obj Account
	obj.Account.Select("id", accountID)
	if obj.Account.IsInitial() {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "account__not_found"})
		return
	}
	// Only owner of an account can delete his account. Administrator(s) can delete user(s) account too.
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.ID != accountID && (obj.Account.AccessLevel != jwt.UserAL || (obj.Account.AccessLevel == jwt.UserAL && payload.AccessLevel != jwt.AdminAL)) {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__cannot_modify_account"})
		return
	}
	// super-administrator account can't be deleted.
	if accountID == 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__cannot_self_terminate"})
		return
	}

	// First, remove this author from all pages
	pages := obj.PageCollaboration.SelectBy("account_id", accountID)
	if len(pages) > 0 {
		for index := range pages {
			if err := obj.PageCollaboration.Remove("account_id", pages[index].AccountID); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_deletion_failed", "error": err.Error()})
				return
			}
		}
	}
	// Then, deletes this account
	if err := obj.Account.Remove("id", accountID); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_deletion_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
