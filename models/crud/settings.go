/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>, Louis LAUGIER <l.laugier@protonmail.com>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package crud

import (
	"reflect"

	"framagit.org/InfoLibre/rapido/models/library"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Object to handle settings variables
type Settings struct {
	ID                  int    `db:"id" json:"id,omitempty"`
	ShowLogo            bool   `db:"show_logo" json:"show_logo"`
	ShowMenus           bool   `db:"show_menus" json:"show_menus"`
	ShowLocalesMenu     bool   `db:"show_locales_menu" json:"show_locales_menu"`
	SMTPHost            string `db:"smtp_host" json:"smtp_host"`
	SMTPPort            int    `db:"smtp_port" json:"smtp_port"`
	SMTPUser            string `db:"smtp_user" json:"smtp_user"`
	SMTPPassword        string `db:"smtp_password" json:"smtp_password"`
	EmailSender         string `db:"email_sender" json:"email_sender"`
	ActivationEmail     string `db:"activation_email" json:"activation_email"`
	ActivatedEmail      string `db:"activated_email" json:"activated_email"`
	ForgetPasswordEmail string `db:"forget_password_email" json:"forget_password_email"`
	PageAuthor          string `db:"page_author" json:"page_author"`
	SharePageEmail      string `db:"share_page_email" json:"share_page_email"`
	Name                string `db:"name" json:"name"`
	HomePage            int    `db:"home_page" json:"home_page"`
	BannedAttemps       int  `db:"banned_attemps" json:"banned_attemps"`
	LockTime            int  `db:"lock_time" json:"lock_time"`
	AccessAttemps       int    `db:"access_attemps" json:"access_attemps"`
	AccessExpired       int  `db:"access_expired" json:"access_expired"`
	SecretKey           string `db:"secret_key" json:"secret_key"`
	MaxSizeUpload       int  `db:"max_size_upload" json:"max_size_upload"`
	FaviconURL          string `db:"favicon_url" json:"favicon_url"`
	LogoURL             string `db:"logo_url" json:"logo_url"`
}

// Selects single row from table settings
func (obj *Settings) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	obj.clear()
	err = tx.Select("*").
		From("settings").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
	}

	return err

}

// Selects all row from table settings
func (obj *Settings) SelectAll() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}
	obj.clear()
	err = tx.Select("*").
		From("settings").
		One(obj)
	if err == nil {
		tx.Commit()
	}

	return err

}

// Update single row in table settings
func (obj *Settings) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	_, err = tx.Update("settings",
		dbx.Params{
			"show_logo":             obj.ShowLogo,
			"show_menus":            obj.ShowMenus,
			"show_locales_menu":     obj.ShowLocalesMenu,
			"smtp_host":             obj.SMTPHost,
			"smtp_port":             obj.SMTPPort,
			"smtp_user":             obj.SMTPUser,
			"smtp_password":         obj.SMTPPassword,
			"email_sender":          obj.EmailSender,
			"activation_email":      obj.ActivationEmail,
			"activated_email":       obj.ActivatedEmail,
			"forget_password_email": obj.ForgetPasswordEmail,
			"page_author":           obj.PageAuthor,
			"share_page_email":      obj.SharePageEmail,
			"name":                  obj.Name,
			"home_page":             obj.HomePage,
			"banned_attemps":        obj.BannedAttemps,
			"lock_time":             obj.LockTime,
			"access_attemps":        obj.AccessAttemps,
			"access_expired":        obj.AccessExpired,
			"secret_key":            obj.SecretKey,
			"max_size_upload":       obj.MaxSizeUpload,
			"favicon_url":           obj.FaviconURL,
			"logo_url":              obj.LogoURL,
		},
		dbx.HashExp{}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err

}

// Checks if an object is empty or not
func (obj Settings) IsInitial() bool {
	if reflect.DeepEqual(obj, Settings{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *Settings) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
