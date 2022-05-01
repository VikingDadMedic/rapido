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
	"time"

	"framagit.org/InfoLibre/rapido/models/library"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Object to communicate with table page_version
type PageVersion struct {
	ID         int    `db:"id" json:"id,omitempty"`
	PageID     int    `db:"page_id" json:"page_id,omitempty"`
	Version    int    `db:"-" json:"version,omitepty"`
	DataString string `db:"data" json:"-"`
	Data       Data   `db:"-" json:"data,omitempty"`
	CreatedOn  int64  `db:"created_on" json:"created_on,omitempty"`
	CreatedBy  int    `db:"created_by" json:"created_by,omitempty"`
	Current    int    `db:"current" json:"current,omitempty"`
	Name       string `db:"-" json:"name"`
}

// Object to store and get version data
type Data struct {
	Page        Page          `db:"-" json:"page,omitempty"`
	PageContent []PageContent `db:"-" json:"page_content,omitempty"`
}

// Unmarshals JSON settings
func (obj *PageVersion) unmarshalJSON() {
	json.Unmarshal([]byte(obj.DataString), &obj.Data)
}

// Marshals JSON settings
func (obj *PageVersion) marshalJSON() {
	jsonByte, err := json.Marshal(obj.Data)
	if err == nil {
		obj.DataString = string(jsonByte)
	}
}

// Selects single row from table page_version
func (obj *PageVersion) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("page_version").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
		obj.unmarshalJSON()
	}

	return err
}

// Selects by column from table page_version
func (obj *PageVersion) SelectBy(column string, value ...interface{}) (output []PageVersion) {
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
		From("page_version").
		Where(exp).
		All(&output)
	if err == nil {
		tx.Commit()
		for index := range output {
			output[index].Version = index + 1
			output[index].unmarshalJSON()
		}
	} else {
		log.Println(err)
	}

	return
}

// Inserts single row in table page_version
func (obj *PageVersion) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	obj.CreatedOn = time.Now().Unix()
	_, err = tx.Insert("page_version",
		dbx.Params{
			"page_id":    obj.PageID,
			"data":       obj.DataString,
			"current":    1,
			"created_on": obj.CreatedOn,
			"created_by": obj.CreatedBy,
		}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table page_version
func (obj *PageVersion) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.marshalJSON()
	_, err = tx.Update("page_version",
		dbx.Params{
			"current": 0,
		},
		dbx.HashExp{"page_id": obj.PageID}).
		Execute()
	if err == nil {
		_, err = tx.Update("page_version",
			dbx.Params{
				"current": 1,
			},
			dbx.HashExp{"id": obj.ID}).
			Execute()
		if err == nil {
			tx.Commit()
		}
	}

	return err
}

// Deletes single row in table page_version
func (obj *PageVersion) Remove(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	_, err = tx.Delete("page_version", dbx.HashExp{column: value}).Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj PageVersion) IsInitial() bool {
	if reflect.DeepEqual(obj, PageVersion{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *PageVersion) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
