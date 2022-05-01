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
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"framagit.org/InfoLibre/rapido/models/settings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

/**
 * @api {POST} /api/file/media Adds file
 * @apiDescription Adds file
 * @apiGroup File
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiSuccess {String} name File name
 * @apiSuccess {String} url File URL
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * [
 *		{
 *			"name": "image_a.jpeg",
 *			"url": "http://rapidohandle.com/storage/file/image/image_a.jpeg"
 *		}
 * ]
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {[]Blob} files File.
 */

// Uploads files to server
func AddMedia(c *gin.Context) {
	var wd, _ = os.Getwd()
	var rootFiles = strings.Join([]string{wd, "storage", "files"}, string(os.PathSeparator))
	form, err1 := c.MultipartForm()
	if err1 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err1.Error()})
		return
	}
	var obj File
	files := form.File["files"]
	errorList := make([]map[string]interface{}, 0)
	successList := make([]map[string]interface{}, 0)
	for _, file := range files {
		if file.Size > (settings.MaxSizeUpload() << 20) {
			errorList = append(errorList, map[string]interface{}{"message": "settings_files__exceeds_allowed_maximum_size", "error": fmt.Sprintf("File %v", file.Filename)})
		} else {
			fileType, err2 := findFileType(file)
			if err2 != nil {
				errorList = append(errorList, map[string]interface{}{"message": "settings_files__file_type_detection_failed", "error": fmt.Sprintf("File %v - Error %v", file.Filename, err2.Error())})
			} else {
				tempName := EncodeFilename(file.Filename)
				if err := c.SaveUploadedFile(file, strings.Join([]string{rootFiles, fileType, tempName}, string(os.PathSeparator))); err != nil {
					errorList = append(errorList, map[string]interface{}{"message": "settings_files__upload_failed", "error": fmt.Sprintf("File %v - Error %v", file.Filename, err.Error())})
				} else {
					obj.File.URL = fmt.Sprintf("%v/storage/files/%v/%v", c.Request.Host, fileType, tempName)
					successList = append(successList, map[string]interface{}{"name": tempName, "url": obj.File.URL})
					obj.File.Insert()
				}
			}
		}
	}

	if len(successList) > 0 {
		if len(errorList) > 0 {
			c.AbortWithStatusJSON(http.StatusMultiStatus, map[string]interface{}{"success": successList, "message": errorList})
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{"success": successList, "message": "settings_files__file_uploaded"})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, map[string]interface{}{"message": errorList})
	}
	return
}

func findFileType(file *multipart.FileHeader) (fileType string, err error) {
	buff, err := file.Open()
	head := make([]byte, 512)
	_, err = buff.Read(head)
	// See https://github.com/gabriel-vasile/mimetype/blob/master/supported_mimes.md
	contentType := mimetype.Detect(head).String()
	if err == nil {
		if strings.Contains(contentType, "image/") {
			fileType = "image"
		} else if strings.Contains(contentType, "audio/") {
			fileType = "audio"
		} else if strings.Contains(contentType, "video/") {
			fileType = "video"
		} else if strings.Contains(contentType, "application/vnd.oasis.opendocument") ||
			strings.Contains(contentType, "application/vnd.openxmlformats-officedocument") ||
			contentType == "application/vnd.ms-excel" || contentType == "application/vnd.ms-publisher" || contentType == "application/vnd.ms-powerpoint" ||
			contentType == "application/msword" || contentType == "application/x-msaccess" ||
			contentType == "application/pdf" || contentType == "application/vnd.fdf" ||
			contentType == "application/postscript" || contentType == "text/rtf" ||
			strings.Contains(contentType, "text/csv") || strings.Contains(contentType, "text/html") || strings.Contains(contentType, "text/plain") || strings.Contains(contentType, "text/xml") ||
			contentType == "application/epub+zip" {
			fileType = "document"
		} else {
			fileType = "other"
		}
	}

	return fileType, err
}

// Replaces all special characters in filename and adds a random number at the beginning of the filename
func EncodeFilename(str string) string {

	// Rules of chars to replace, a regex is defined below to replace remaining special characters
	mapRule := func(r rune) rune {
		switch {
		case r == ' ':
			return '_'
		case r == 'á' || r == 'à' || r == 'â' || r == 'ã' || r == 'ä' || r == 'å':
			return 'a'
		case r == 'Á' || r == 'À' || r == 'Â' || r == 'Ã' || r == 'Ä' || r == 'Å':
			return 'A'
		case r == 'ç':
			return 'c'
		case r == 'Ç':
			return 'C'
		case r == 'é' || r == 'è' || r == 'ê' || r == 'ẽ' || r == 'ë':
			return 'e'
		case r == 'É' || r == 'È' || r == 'Ê' || r == 'Ẽ' || r == 'Ë' || r == '€':
			return 'E'
		case r == 'í' || r == 'ì' || r == 'î' || r == 'ĩ' || r == 'ï':
			return 'i'
		case r == 'Í' || r == 'Ì' || r == 'Î' || r == 'Ĩ' || r == 'Ï':
			return 'I'
		case r == 'µ':
			return 'm'
		case r == 'ñ':
			return 'n'
		case r == 'Ñ':
			return 'N'
		case r == 'ó' || r == 'ò' || r == 'ô' || r == 'õ' || r == 'ö' || r == 'ø':
			return 'o'
		case r == 'Ó' || r == 'Ò' || r == 'Ô' || r == 'Õ' || r == 'Ö' || r == 'Ø':
			return 'O'
		case r == 'ú' || r == 'ù' || r == 'û' || r == 'ũ' || r == 'ü' || r == 'ů':
			return 'u'
		case r == 'Ú' || r == 'Ù' || r == 'Û' || r == 'Ũ' || r == 'Ü' || r == 'Ů':
			return 'U'
		case r == 'ý' || r == 'ỳ' || r == 'ŷ' || r == 'ỹ' || r == 'ÿ' || r == 'ẙ':
			return 'y'
		case r == 'Ý' || r == 'Ỳ' || r == 'Ŷ' || r == 'Ỹ' || r == 'Ÿ':
			return 'Y'
		}

		return r
	}

	// Uses the map to replace all special characters
	convertedString := strings.Map(mapRule, str)

	// Defines a regular expression that only accepts a-z A-Z 0-9 . _ - ~ characters (see RFC3986 section 2.3)
	regexpPattern, err := regexp.Compile("[^a-zA-Z0-9._\\-~]+")
	if err != nil {
		log.Fatal(err)
	}
	// Deletes all special characters
	result := regexpPattern.ReplaceAllString(convertedString, "")

	// Creates a random id using xid library
	guid := xid.New()

	StringID := guid.String()

	// Adds the random id before the converted string extension or at the end of the converted string
	extension := filepath.Ext(result)
	if extension != "" {
		filenameWithoutExt := strings.TrimSuffix(result, extension)
		result = filenameWithoutExt + "_" + StringID + extension
	} else {
		result = result + "_" + StringID
	}

	// Returns the correctly formatted link
	return result
}
