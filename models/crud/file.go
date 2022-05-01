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
	"log"
	"reflect"

	"framagit.org/InfoLibre/rapido/models/library"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Object to communicate with table file
type File struct {
	URL  string `db:"url" json:"url,omitempty"`
	Used int    `db:"used" json:"used,omitempty"`
}

// Selects single row from table file
func (obj *File) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("file").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
	}

	return err
}

// Selects all row from table file
func (obj *File) SelectAll() (output []File) {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		log.Println(err)
		return
	}

	err = tx.Select("*").
		From("file").
		All(&output)
	if err == nil {
		tx.Commit()
	} else {
		log.Println(err)
	}

	return
}

// Selects by column from table file
func (obj *File) SelectBy(column string, value ...interface{}) (output []File) {
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
		From("file").
		Where(exp).
		All(&output)
	if err == nil {
		tx.Commit()
	} else {
		log.Println(err)
	}

	return
}

// Inserts single row in table file
func (obj *File) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	_, err = tx.Insert("file",
		dbx.Params{
			"url":  obj.URL,
			"used": obj.Used}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table file
func (obj *File) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	_, err = tx.Update("file",
		dbx.Params{"used": obj.Used},
		dbx.HashExp{"url": obj.URL}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Deletes single row in table file
func (obj *File) Remove(column string, value ...interface{}) error {
	var exp dbx.Expression
	if len(value) > 1 {
		exp = dbx.In(column, value)
	} else {
		exp = dbx.HashExp{column: value[0]}
	}

	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	_, err = tx.Delete("file", exp).Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj File) IsInitial() bool {
	if reflect.DeepEqual(obj, File{}) {
		return true
	}
	return false
}

// Sets empty struct
func (obj *File) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
