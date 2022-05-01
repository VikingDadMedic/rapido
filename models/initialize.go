/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

// initialize.go is used to setup application configuration

package models

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware to handle cors
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Creates storing directories
func InitPath() {
	wd, _ := os.Getwd()
	rootStorage := strings.Join([]string{wd, "storage"}, string(os.PathSeparator))
	if _, err := os.Stat(rootStorage); os.IsNotExist(err) {
		err = os.Mkdir(rootStorage, os.ModePerm)
	}

	// Creates directories to store uploaded documents
	rootFiles := strings.Join([]string{wd, "storage", "files"}, string(os.PathSeparator))
	if _, err := os.Stat(rootFiles); os.IsNotExist(err) {
		err = os.Mkdir(rootFiles, os.ModePerm)
	}
	documentPath := strings.Join([]string{wd, "storage", "files", "document"}, string(os.PathSeparator))
	if _, err := os.Stat(documentPath); os.IsNotExist(err) {
		err = os.Mkdir(documentPath, os.ModePerm)
	}
	audioPath := strings.Join([]string{wd, "storage", "files", "audio"}, string(os.PathSeparator))
	if _, err := os.Stat(audioPath); os.IsNotExist(err) {
		err = os.Mkdir(audioPath, os.ModePerm)
	}
	videoPath := strings.Join([]string{wd, "storage", "files", "video"}, string(os.PathSeparator))
	if _, err := os.Stat(videoPath); os.IsNotExist(err) {
		err = os.Mkdir(videoPath, os.ModePerm)
	}
	imagePath := strings.Join([]string{wd, "storage", "files", "image"}, string(os.PathSeparator))
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		err = os.Mkdir(imagePath, os.ModePerm)
	}
	otherPath := strings.Join([]string{wd, "storage", "files", "other"}, string(os.PathSeparator))
	if _, err := os.Stat(otherPath); os.IsNotExist(err) {
		err = os.Mkdir(otherPath, os.ModePerm)
	}

	// Creates directory to store application backups
	rootBackup := strings.Join([]string{wd, "storage", "backup"}, string(os.PathSeparator))
	if _, err := os.Stat(rootBackup); os.IsNotExist(err) {
		err = os.Mkdir(rootBackup, os.ModePerm)
	}

	// Creates directory to store temporary emails attachments
	rootTmp := strings.Join([]string{wd, "storage", "tmp"}, string(os.PathSeparator))
	if _, err := os.Stat(rootTmp); os.IsNotExist(err) {
		err = os.Mkdir(rootTmp, os.ModePerm)
	}
}

/**
 * @apiDefine Authorize
 * Can be access after login
 */

/**
 * @apiDefine Admin
 * Admin user access only
 */

/**
 * @apiDefine Public
 * Public access
 */

/**
 * @apiDefine AuthorizeHeader
 * @apiHeader {String} Authorization Access token.
 * @apiHeaderExample {json} Header-Example:
 * {
 *		"Authorization": "Bearer ajsnfs9d8nv98njdfsiyvnowo8efasjfisudb82e.djsidbfidu9vudsbvwd.snf98sdvub9sduhw9ufwsf"
 * }
 */

/**
 * @apiDefine Error40X Error 40X
 * @apiError (Error40X) {String} message Error Message
 * @apiErrorExample {json} Error-Response:
 * HTTP/1.1 40X
 * {
 *		"message": "Error message"
 * }
 */

/**
 * @apiDefine Error50X Error 50X
 * @apiError (Error50X) {String} message Error Message
 * @apiErrorExample {json} Error-Response:
 * HTTP/1.1 50X
 * {
 *		"message": "Error message"
 * }
 */
