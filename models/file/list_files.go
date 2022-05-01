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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Param on this file
type MediaParam struct {
	Show string `form:"show"`
}

// Object to store files list
type MediaList struct {
	Name string `json:"name,omitempty" form:"name" binding:"required"`
	Type string `json:"type,omitempty" form:"type" binding:"required"`
	URL  string `json:"url,omitempty" form:"url"`
}

/**
 * @api {GET} /api/file/file Lists files
 * @apiDescription Lists files
 * @apiGroup File
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccess {String} name File name
 * @apiSuccess {String} type File type
 * @apiSuccess {String} url File URL
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * [
 *		{
 *			"name": "image_a.jpeg",
 *			"type": "image",
 *			"url": "http://rapidohandle.com/storage/file/image/image_a.jpeg"
 *		}
 * ]
 * @apiUse Error40X
 * @apiUse Error50X
 */

// Gets list of uploaded files
func FilesList(c *gin.Context) {
	var wd, _ = os.Getwd()
	var rootFiles = strings.Join([]string{wd, "storage", "files"}, string(os.PathSeparator))
	var obj File
	var param MediaParam
	if err := c.ShouldBindQuery(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}

	files_list := obj.File.SelectAll()

	root, _ := ioutil.ReadDir(rootFiles)
	list := make([]MediaList, 0)
	for index := range root {
		sub, _ := ioutil.ReadDir(strings.Join([]string{rootFiles, root[index].Name()}, string(os.PathSeparator)))
		for ix := range sub {
			if !sub[ix].IsDir() {
				url := fmt.Sprintf("%v/storage/files/%v/%v", c.Request.Host, root[index].Name(), sub[ix].Name())
				for ixm := range files_list {
					if files_list[ixm].URL == url {
						show := false
						switch param.Show {
						case "used":
							if files_list[ixm].Used > 0 {
								show = true
							}
						case "unused":
							if files_list[ixm].Used == 0 {
								show = true
							}
						case "all":
							show = true
						case "none":
							show = false
						default:
							show = true
						}

						if show {
							list = append(list, MediaList{sub[ix].Name(), root[index].Name(), url})
						}
						break
					}
				}
			}
		}
	}

	if len(list) == 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, list)
}
