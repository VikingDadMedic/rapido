/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package file

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Object to store files list
type BackupFileList struct {
	Name      string `json:"name,omitempty"`
	CreatedOn string `json:"created_on,omitempty"`
}

/**
 * @api {GET} /api/file/backup_list Backup List
 * @apiDescription Backup List
 * @apiGroup File
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccess {String} name Image Name
 * @apiSuccess {Timestamp} created_on Data Modified Time
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * [
 *		{
 *			"name": "upgrade_backup.zip",
 *			"created_on": "2018-01-02 11:20:31"
 *		}
 * ]
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets list of backup files
func BackupList(c *gin.Context) {
	var wd, _ = os.Getwd()
	var backupPath = strings.Join([]string{wd, "storage", "backup"}, string(os.PathSeparator))
	temp, _ := ioutil.ReadDir(backupPath)
	list := make([]BackupFileList, 0)
	for index := range temp {
		if !temp[index].IsDir() {
			if strings.HasPrefix(strings.ToLower(temp[index].Name()), "upgrade_backup_") && strings.HasSuffix(strings.ToLower(temp[index].Name()), ".zip") {
				list = append(list, BackupFileList{temp[index].Name(), temp[index].ModTime().Format("2006-01-02 15:04:05")})
			}
		}
	}

	if len(list) == 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, list)
}
