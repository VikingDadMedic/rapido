/*
framagit.org/InfoLibre/rapido is software to make a website. framagit.org/InfoLibre/rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>, Félix PORTIER <f.portierdev@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package models

import (
	"net/http"

	"framagit.org/InfoLibre/rapido/models/account"
	"framagit.org/InfoLibre/rapido/models/contactform"
	"framagit.org/InfoLibre/rapido/models/content"
	"framagit.org/InfoLibre/rapido/models/email"
	"framagit.org/InfoLibre/rapido/models/file"
	"framagit.org/InfoLibre/rapido/models/jwt"
	"framagit.org/InfoLibre/rapido/models/menu"
	"framagit.org/InfoLibre/rapido/models/page"
	"framagit.org/InfoLibre/rapido/models/settings"
	"framagit.org/InfoLibre/rapido/models/theme"
	"framagit.org/InfoLibre/rapido/models/version"

	"github.com/gin-gonic/gin"
)

// Sets up routing blueprint for app
func BlueprintApp(router *gin.Engine) {
	router.MaxMultipartMemory = settings.MaxSizeUpload() << 20
	router.NoRoute(func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "client/dist/index.html")
	})

	router.Use(corsMiddleware())

	router.Static("/storage/files", "./storage/files")
	router.Static("/static/", "./client/dist/static/")
	router.Static("/lang/", "./lang/")

	public := router.Group("/api")
	{
		public.GET("/page/:id", page.ListAuthors)
		public.GET("/published", page.ListPublished)
		public.GET("/archived", page.ListArchived)
		public.GET("/page/:id/version", page.ListVersions)
		public.POST("/page/:id/author", page.AddAuthor)
		public.POST("/page/:id/share", page.Share)

		public.GET("/content/page", content.Get)
		public.GET("/content", content.List)

		public.GET("/contactform/page", contactform.Get)
		public.GET("/contactform", contactform.List)

		public.POST("/email", email.Action)

		public.GET("/menu_item", menu.Get)
		public.GET("/menu", menu.List)

		public.GET("/theme", theme.Get)

		public.GET("/activate_account", account.Activate)
		public.POST("/register", account.Register)
		public.POST("/forget_password", account.ForgetPassword)
		public.POST("/new_password", account.RenewPass)
		public.POST("/login", account.Login)

		public.GET("/setting", settings.Get)
		public.GET("/logo", settings.Logo)
	}

	authorize := public.Group("")
	{
		authorize.Use(jwt.Auth(settings.SecretKey()))

		authorize.POST("/page", page.Add)
		authorize.POST("/page/:id/copy", page.Copy)
		authorize.PUT("/page/:id", page.Change)
		authorize.PUT("/page/:id/archive", page.Archive)
		authorize.PUT("/page/:id/publish", page.Publish)
		authorize.PUT("/page/:id/revert", page.RevertVersion)
		authorize.DELETE("/page/:id/author/:email", page.RemoveAuthor)
		authorize.DELETE("/page/:id/page", page.DeletePage)

		authorize.POST("/content", content.Add)
		authorize.PUT("/content/:id", content.Change)
		authorize.DELETE("/content/:id", content.Remove)

		authorize.POST("/contactform", contactform.Add)
		authorize.PUT("/contactform/:id", contactform.Change)
		authorize.DELETE("/contactform/:id", contactform.Remove)

		authorize.GET("/file/media", file.FilesList)
		authorize.POST("/file/media", file.AddMedia)
		authorize.DELETE("/file/media", file.RemoveMedia)

		authorize.PUT("/theme", theme.Change)

		authorize.GET("/account", account.List)
		authorize.GET("/check_token", account.CheckToken)
		authorize.PUT("/account/:id", account.Change)
		authorize.PUT("/change_password/:id", account.ChangePassword)
		authorize.DELETE("/account/:id", account.Remove)

		authorize.POST("/menu", menu.Change)

		authorize.GET("/version", version.Get)
	}

	admin := authorize.Group("")
	{
		admin.Use(jwt.RBAC(jwt.AdminAL))
		admin.POST("/account", account.Add)


		admin.GET("/listallpages", page.ListAll)
		admin.PUT("/setting", settings.Change)

		// admin.GET("/file/backup", backup.Backup)
		// admin.GET("/file/backup_list", backup.BackupList)
		// admin.POST("/system/revert", backup.Revert)
	}

}
