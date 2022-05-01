/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2020 David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package theme

import (
	"net/http"

	"framagit.org/InfoLibre/rapido/models/jwt"

	"github.com/gin-gonic/gin"
)

// Gets theme data
func Get(c *gin.Context) {
	var obj Theme
	payload, _ := jwt.Extract(c.GetHeader("Authorization"))
	themes, err := obj.Theme.SelectAll(payload.ID)
	if err != nil || len(themes) == 0 {
		c.AbortWithStatusJSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, themes)
}
