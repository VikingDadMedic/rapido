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
	"log"
	"reflect"

	"framagit.org/InfoLibre/rapido/models/library"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Object to communicate with table menu
type Menu struct {
	ID           int                      `db:"id" json:"id,omitempty"`
	Name         string                   `db:"name" json:"name,omitempty"`
	JSONSettings library.NullString       `db:"json_settings" json:"-"`
	AccessLevel  library.NullInt64        `db:"access_level" json:"access_level,omitempty"`
	Settings     []map[string]interface{} `db:"-" json:"json_settings,omitempty"`
	Item         []MenuItem               `db:"-" json:"item,omitempty"`
}

// Unmarshals JSON settings
func (obj *Menu) unmarshalJSON() {
	if obj.JSONSettings.Valid {
		obj.Settings, _ = library.Jsonify(obj.JSONSettings.String)
	}
}

// Marshals JSON settings
func (obj *Menu) marshalJSON() {
	jsonByte, err := json.Marshal(obj.Settings)
	if err == nil {
		obj.JSONSettings.String = string(jsonByte)
		obj.JSONSettings.Valid = true
	}
}

// Selects single row from table menu
func (obj *Menu) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("menu").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
		obj.unmarshalJSON()
	}

	return err
}

// Selects all row from table menu
func (obj *Menu) SelectAll() (output []Menu, err error) {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return
	}

	err = tx.Select("*").
		From("menu").
		All(&output)
	if err == nil {
		tx.Commit()
		for index := range output {
			output[index].unmarshalJSON()
		}
	}

	return
}

// Selects by column from table content
func (obj *Menu) SelectBy(column string, value ...interface{}) (output []Menu) {
	var exp dbx.Expression
	if len(value) > 1 {
		exp = dbx.In(column, value)
	} else {
		exp = dbx.HashExp{column: value[0]}
	}

	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		log.Println(err)
		return
	}

	err = tx.Select("*").
		From("menu").
		Where(exp).
		All(&output)
	if err == nil {
		tx.Commit()
		for index := range output {
			output[index].unmarshalJSON()
		}
	} else {
		log.Println(err)
	}

	return
}

// Inserts single row in table menu
func (obj *Menu) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	_, err = tx.Insert("menu",
		dbx.Params{
			"name":          obj.Name,
			"json_settings": obj.JSONSettings,
			"access_level":  obj.AccessLevel}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table menu
func (obj *Menu) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	_, err = tx.Update("menu",
		dbx.Params{
			"name":          obj.Name,
			"json_settings": obj.JSONSettings,
			"access_level":  obj.AccessLevel},
		dbx.HashExp{"id": obj.ID}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Deletes single row in table menu
func (obj *Menu) Remove(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	_, err = tx.Delete("menu", dbx.HashExp{column: value}).Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj Menu) IsInitial() bool {
	if reflect.DeepEqual(obj, Menu{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *Menu) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
