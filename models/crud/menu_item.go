/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package crud

import (
	"encoding/json"
	"reflect"

	"framagit.org/InfoLibre/rapido/models/library"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Object to communicate with table menu_item
type MenuItem struct {
	ID           int                      `db:"id" json:"id,omitempty"`
	MenuID       int                      `db:"menu_id" json:"menu_id,omitempty"`
	Position     int                      `db:"position" json:"position,omitempty"`
	Level        int                      `db:"level" json:"level,omitempty"`
	PageID       library.NullInt64        `db:"page_id" json:"page_id,omitempty"`
	LinkLabel    library.NullString       `db:"link_label" json:"link_label,omitempty"`
	LinkURL      library.NullString       `db:"link_url" json:"link_url,omitempty"`
	LinkTarget   library.NullString       `db:"link_target" json:"link_target,omitempty"`
	JSONSettings library.NullString       `db:"json_settings" json:"-"`
	AccessLevel  library.NullInt64        `db:"access_level" json:"access_level,omitempty"`
	UserID       int                      `db:"user_id" json:"user_id,omitempty"`
	Settings     []map[string]interface{} `db:"-" json:"json_settings,omitempty"`
	Page         Page                     `db:"-" json:"page,omitempty"`
	SubMenu      []*MenuItem              `db:"-" json:"sub_menu,omitempty"`
}

// Sorts Menu Item object
type MenuByPosition []MenuItem

func (a MenuByPosition) Len() int           { return len(a) }
func (a MenuByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MenuByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

// Unmarshals JSON settings
func (obj *MenuItem) unmarshalJSON() {
	if obj.JSONSettings.Valid {
		obj.Settings, _ = library.Jsonify(obj.JSONSettings.String)
	}
}

// Marshals JSON settings
func (obj *MenuItem) marshalJSON() {
	jsonByte, err := json.Marshal(obj.Settings)
	if err == nil {
		obj.JSONSettings.String = string(jsonByte)
		obj.JSONSettings.Valid = true
	}
}

// Selects single row from table menu_item
func (obj *MenuItem) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("menu_item").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
		obj.unmarshalJSON()
	}

	return err
}

// Selects all row from table menu item
func (obj *MenuItem) SelectAll() (output []MenuItem, err error) {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return
	}

	err = tx.Select("*").
		From("menu_item").
		All(&output)
	if err == nil {
		tx.Commit()
		for index := range output {
			output[index].unmarshalJSON()
		}
	}

	return
}

// Inserts single row in table menu_item
func (obj *MenuItem) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	_, err = tx.Insert("menu_item",
		dbx.Params{
			"menu_id":       obj.MenuID,
			"position":      obj.Position,
			"level":         obj.Level,
			"page_id":       obj.PageID,
			"link_label":    obj.LinkLabel,
			"link_url":      obj.LinkURL,
			"link_target":   obj.LinkTarget,
			"json_settings": obj.JSONSettings,
			"access_level":  obj.AccessLevel,
			"user_id":       obj.UserID}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table menu_item
func (obj *MenuItem) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	_, err = tx.Update("menu_item",
		dbx.Params{
			"menu_id":       obj.MenuID,
			"position":      obj.Position,
			"level":         obj.Level,
			"page_id":       obj.PageID,
			"link_label":    obj.LinkLabel,
			"link_url":      obj.LinkURL,
			"link_target":   obj.LinkTarget,
			"json_settings": obj.JSONSettings,
			"access_level":  obj.AccessLevel},
		dbx.HashExp{"id": obj.ID, "user_id": obj.UserID}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Deletes single row in table menu_item
func (obj *MenuItem) Remove(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	_, err = tx.Delete("menu_item", dbx.HashExp{column: value}).Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj MenuItem) IsInitial() bool {
	if reflect.DeepEqual(obj, MenuItem{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *MenuItem) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
