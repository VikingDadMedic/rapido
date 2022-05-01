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

// Object to communicate with table page_content
type PageContent struct {
	ID            int                      `db:"id" json:"id,omitempty"`
	PageID        int                      `db:"page_id" json:"page_id,omitempty" binding:"required"`
	ContentID     library.NullInt64        `db:"content_id" json:"content_id,omitempty"`
	ContactFormID library.NullInt64        `db:"contactform_id" json:"contactform_id,omitempty"`
	Extension     library.NullString       `db:"extension" json:"extension,omitempty"`
	Location      string                   `db:"location" json:"location,omitempty" binding:"required"`
	Column        int                      `db:"column" json:"column,omitempty" binding:"required"`
	Position      int                      `db:"position" json:"position,omitempty" binding:"required"`
	JSONSettings  library.NullString       `db:"json_settings" json:"-"`
	Settings      []map[string]interface{} `db:"-" json:"json_settings,omitempty"`
	Content       Content                  `db:"-" json:"content,omitempty"`
	ContactForm   ContactForm              `db:"-" json:"contactform,omitempty"`
}

// Sort PageContent object
type PageByPosition []PageContent

func (a PageByPosition) Len() int           { return len(a) }
func (a PageByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PageByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

// Unmarshals JSON settings
func (obj *PageContent) unmarshalJSON() {
	if obj.JSONSettings.Valid {
		obj.Settings, _ = library.Jsonify(obj.JSONSettings.String)
	}
}

// Marshals JSON settings
func (obj *PageContent) marshalJSON() {
	jsonByte, err := json.Marshal(obj.Settings)
	if err == nil {
		obj.JSONSettings.String = string(jsonByte)
		obj.JSONSettings.Valid = true
	}
}

// Selects single row from table page_content
func (obj *PageContent) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("page_content").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
		obj.unmarshalJSON()
	}

	return err
}

// Selects by column from table page_content
func (obj *PageContent) SelectBy(column string, value ...interface{}) (output []PageContent) {
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
		From("page_content").
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

// Inserts single row in table page_content
func (obj *PageContent) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	_, err = tx.Insert("page_content",
		dbx.Params{
			"page_id":       obj.PageID,
			"content_id":    obj.ContentID,
			"extension":     obj.Extension,
			"location":      obj.Location,
			"column":        obj.Column,
			"position":      obj.Position,
			"json_settings": obj.JSONSettings}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table page_content
func (obj *PageContent) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	_, err = tx.Update("page_content",
		dbx.Params{
			"page_id":       obj.PageID,
			"content_id":    obj.ContentID,
			"extension":     obj.Extension,
			"location":      obj.Location,
			"column":        obj.Column,
			"position":      obj.Position,
			"json_settings": obj.JSONSettings},
		dbx.HashExp{"id": obj.ID}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Deletes single row in table page_content
func (obj *PageContent) Remove(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	_, err = tx.Delete("page_content", dbx.HashExp{column: value}).Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj PageContent) IsInitial() bool {
	if reflect.DeepEqual(obj, PageContent{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *PageContent) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
