/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package page

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"framagit.org/InfoLibre/rapido/models/crud"
	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

/**
 * @api {POST} /api/page Create a new page
 * @apiDescription Create a new page
 * @apiGroup Page
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * {
 *		"published": 0,
 *		"link": "test",
 *		"name": "Home",
 *		"title": "Rapido Home Page",
 *		"keywords": "Rapido,home,page",
 *		"descriptions": null,
 *		"json_settings": null,
 * }
 * @apiSuccess {Int} id Page ID
 * @apiSuccess {Int} published Publication status
 * @apiSuccess {String} link URL link
 * @apiSuccess {String} name Page name
 * @apiSuccess {String} title Page title
 * @apiSuccess {String} keywords Page keywords
 * @apiSuccess {String} descriptions Page description
 * @apiSuccess {[]Object} json_settings Page settings
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 201 Created
 * {
 *		"id": 1,
 *		"published": 0,
 *		"link": "test",
 *		"name": "Home",
 *		"title": "Rapido Home Page",
 *		"keywords": "Rapido,home,page",
 *		"descriptions": null,
 *		"json_settings": null,
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} [published] Publication status
 * @apiParam {String} [link] URL link
 * @apiParam {String} name Page Name
 * @apiParam {String} [title] Page title
 * @apiParam {String} [keywords] Page keywords
 * @apiParam {String} [descriptions] Page description
 * @apiParam {[]Object} [json_settings] Page settings
 */

// Creates a new page
func Add(c *gin.Context) {
	var param crud.Page
	if err := c.ShouldBindJSON(&param); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}
	if param.Published == 1 {
		param.Published = 0
	}
	if param.Title == "" {
		param.Title = param.Name
	}
	if param.Link == "" {
		param.Link = strings.ToLower(param.Name)
	}
	// Uses custom link encoding function to replace special characters
	param.Link = CustomLinkEncoding(param.Link)
	var obj Page
	obj.Page.Select("link", param.Link)
	if !obj.Page.IsInitial() {
		c.AbortWithStatusJSON(http.StatusConflict, map[string]interface{}{"message": "create_page__link_already_exists"})
		return
	}
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	if err := param.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_insertion_failed", "error": err.Error()})
		return
	}

	go addInitContent(payload.ID, param.ID)
	// Adds page creator as page author
	obj.PageCollaboration.AccountID, obj.PageCollaboration.PageID = payload.ID, param.ID
	if err := obj.PageCollaboration.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_insertion_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, param)
}

func addInitContent(accountID, pageID int) {
	var obj Page
	obj.Content.Name.String, obj.Content.Name.Valid = fmt.Sprintf("Initial content for page %v", pageID), true
	obj.Content.Content = ``
	if err := obj.Content.Insert(); err != nil {
		log.Println(err)
	}

	obj.PageContent.PageID = pageID
	obj.PageContent.ContentID.Int64, obj.PageContent.ContentID.Valid = int64(obj.Content.ID), true
	obj.PageContent.Location = "main"
	obj.PageContent.Column, obj.PageContent.Position = 1, 1
	if err := obj.PageContent.Insert(); err != nil {
		log.Println(err)
	}

	AddVersion(accountID, 0, pageID)
}

// Encodes link and replaces all special characters
func CustomLinkEncoding(str string) string {
	// Rules of characters replacement, only single characters can be replaced
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

	// Creates a random id using xid lib
	guid := xid.New()
	StringID := guid.String()
	// Concatenates the converted string and the random id
	convertedString = result + "_" + StringID

	// Returns the correctly formatted link
	return convertedString
}
