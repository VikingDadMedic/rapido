/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package contactform

import (
	"fmt"
	"net/http"

	"framagit.org/InfoLibre/rapido/models/crud"
	// "framagit.org/InfoLibre/rapido/models/jwt" // TODO : uncomment after adding access test

	"github.com/gin-gonic/gin"
)

/**
 * @api {POST} /api/contactform Add new ContactForm
 * @apiDescription Adds contact form
 * @apiGroup ContactForm
 * @apiVersion 0.1.0
 * @apiPermission Authorize
 * @apiUse AuthorizeHeader
 * @apiExample {json} Example usage:
 * [
 *		{
 *			"id": 1,
 *			"extension": "test",
 *			"json_settings": [{
 *			to :"test"
 *			layout:"test"
 *							}],
 *
 *		}
 * ]
 * @apiSuccess {String} message Success Message
 * @apiSuccessExample {json} Success-Response:
 * HTTP/1.1 200 Ok
 * {
 *		"message": ""
 * }
 * @apiUse Error40X
 * @apiUse Error50X
 * @apiParam {Int} [id] ContactForm ID
 * @apiParam {String} [extension] Extension Name
 * @apiParam {[]Object} [json_settings] Page Setting
 */

// Adds contact form data on page
func Add(c *gin.Context) {
	var param crud.ContactForm
	if err := c.ShouldBindJSON(&param); err != nil {
		fmt.Println(&param.JSONSettings)
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}

	if param.Extension == "" {
		param.Extension = "contact-form"
	}

	// payload, _ := jwt.Extract(c.GetHeader("Authorization")) // TODO : add access test

	if err := param.Insert(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{"message": "data_insertion_failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, param)
}
