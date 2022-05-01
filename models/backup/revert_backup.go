/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package backup

import (
	"fmt"
	"net/http"
	"os"

	"framagit.org/InfoLibre/rapido/models/jwt"
	"framagit.org/InfoLibre/rapido/models/library"

	"github.com/gin-gonic/gin"
)

/**
 * @api {POST} /api/backup/revert Reverts to a system backup
 * @apiDescription Reverts to a system backup
 * @apiGroup Backup
 * @apiVersion 0.1.0
 * @apiPermission Admin
 * @apiUse AuthorizeHeader
 * @apiSuccess {String} message Success Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": "Revert to upgrade_backup_1520491283.zip success"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} filename Target File To Revert
 */

// Reverts system backup
type FileParam struct {
	Filename string `db:"filename" json:"filename" binding:"required"`
}

// Reverts to a system backup
func Revert(c *gin.Context) {
	if c.GetHeader("Authorization") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_accounts__access_forbidden"})
		return
	}
	// Only administrator can reverts to a system backup
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if payload.AccessLevel != jwt.AdminAL {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_system__cannot_revert"})
		return
	}
	var param FileParam
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	source, _ := os.Getwd()
	sep := string(os.PathSeparator)
	filePath := fmt.Sprintf("%v%vstorage%vbackup%v%v", source, sep, sep, sep, sep, param.Filename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_files__file_not_found"})
		return
	}

	if err := library.Unzip(filePath, source); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "file_extraction_failed", "error": err.Error()})
		return
	}

	defer library.Restart()

	c.JSON(http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("%v %v", "settings_system__system_and_data_restored", param.Filename)})
}
