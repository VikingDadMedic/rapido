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
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
 * @api {DELETE} /api/file/media Deletes file
 * @apiDescription Deletes file
 * @apiGroup File
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 204 No Content
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {String} name File name
 * @ApiParam {String} type File type
 */

// Deletes file from server
func RemoveMedia(c *gin.Context) {
	var param MediaList
	if err := c.ShouldBindQuery(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	var wd, _ = os.Getwd()
	var rootFiles = strings.Join([]string{wd, "storage", "files"}, string(os.PathSeparator))
	name := param.Name
	types := param.Type
	path := strings.Join([]string{rootFiles, types, name}, string(os.PathSeparator))
	url := strings.Join([]string{c.Request.Host, "storage", "files", types, name}, "/")
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, map[string]interface{}{"message": "settings_files__file_not_found"})
		return
	}
	var obj File
	obj.File.Select("url", url)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_files__file_path_not_found", "error": err.Error()})
		return
	}
	if obj.File.Used != 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "settings_files__file_in_use"})
		return
	}

	obj.File.Remove("url", url)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "settings_files__file_path_deletion_failed", "error": err.Error()})
		return
	}
	if err = os.Remove(path); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "settings_files__file_deletion_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
