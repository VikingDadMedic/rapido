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

// Object to communicate with table page
type Page struct {
	ID           int                      `db:"id" json:"id,omitempty"`
	Published    int                      `db:"published" json:"published,omitempty"`
	Name         string                   `db:"name" json:"name,omitempty" binding:"required"`
	Title        string                   `db:"title" json:"title,omitempty"`
	Link         string                   `db:"link" json:"link,omitempty"`
	Keywords     library.NullString       `db:"keywords" json:"keywords,omitempty"`
	Description  library.NullString       `db:"description" json:"description,omitempty"`
	JSONSettings library.NullString       `db:"json_settings" json:"-"`
	Settings     []map[string]interface{} `db:"-" json:"json_settings,omitempty"`
	Authors      []string                 `db:"-" json:"authors,omitempty"`
}

// Unmarshals JSON settings
func (obj *Page) unmarshalJSON() {
	if obj.JSONSettings.Valid {
		obj.Settings, _ = library.Jsonify(obj.JSONSettings.String)
	}
}

// Marshals JSON settings
func (obj *Page) marshalJSON() {
	jsonByte, err := json.Marshal(obj.Settings)
	if err == nil {
		obj.JSONSettings.String = string(jsonByte)
		obj.JSONSettings.Valid = true
	}
}

// Selects single row from table page
func (obj *Page) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("page").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
		obj.unmarshalJSON()
	}

	return err
}

// Selects by column from table page
func (obj *Page) SelectBy(column string, value ...interface{}) (output []Page) {
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
		From("page").
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

// Selects all row from table page
func (obj *Page) SelectAll() (output []Page, err error) {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return
	}

	err = tx.Select("*").
		From("page").
		All(&output)
	if err == nil {
		tx.Commit()
		for index := range output {
			output[index].unmarshalJSON()
		}
	}

	return
}

// Inserts single row in table page
func (obj *Page) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	_, err = tx.Insert("page",
		dbx.Params{
			"published":     obj.Published,
			"name":          obj.Name,
			"title":         obj.Title,
			"link":          obj.Link,
			"keywords":      obj.Keywords,
			"description":   obj.Description,
			"json_settings": obj.JSONSettings}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table page
func (obj *Page) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	_, err = tx.Update("page",
		dbx.Params{
			"published":     obj.Published,
			"name":          obj.Name,
			"title":         obj.Title,
			"link":          obj.Link,
			"keywords":      obj.Keywords,
			"description":   obj.Description,
			"json_settings": obj.JSONSettings},
		dbx.HashExp{"id": obj.ID}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Deletes single row in table page
func (obj *Page) Remove(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	_, err = tx.Delete("page", dbx.HashExp{column: value}).Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj Page) IsInitial() bool {
	if reflect.DeepEqual(obj, Page{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *Page) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
