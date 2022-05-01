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
	"strings"
	"time"

	"framagit.org/InfoLibre/rapido/models/library"

	"github.com/gin-gonic/gin"
)

/**
 * @api {GET} /api/file/backup Backup Project
 * @apiDescription Backup Project
 * @apiGroup File
 * @apiVersion 0.1.0
 * @apiPermission Admin
 * @apiUse AuthorizeHeader
 * @apiSuccess {String} message Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *  	"message": "upgrade_backup.zip"
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Backups this project to zip
func Backup(c *gin.Context) {
	var wd, _ = os.Getwd()
	var backupPath = strings.Join([]string{wd, "storage", "backup"}, string(os.PathSeparator))
	backupPath, err := BackupSystem()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "settings_system__backup_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": backupPath})
}

// Backups system to zip
func BackupSystem() (string, error) {
	var wd, _ = os.Getwd()
	var backupPath = strings.Join([]string{wd, "storage", "backup"}, string(os.PathSeparator))
	fileName := strings.Join([]string{"upgrade_backup_", fmt.Sprint(time.Now().Unix()), ".zip"}, "")
	backPath := strings.Join([]string{backupPath, fileName}, string(os.PathSeparator))
	tempBackupPath := strings.Replace(backPath, strings.Join([]string{"Rapido", "storage", "backup", ""}, string(os.PathSeparator)), "", -1)
	zip := new(library.ZipFile)
	err := zip.Create(tempBackupPath)
	if err == nil {
		err = zip.AddAll(wd, false)
		zip.Close()
		os.Rename(tempBackupPath, backPath)
	}
	return backPath, err
}
