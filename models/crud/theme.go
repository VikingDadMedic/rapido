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

// Object to communicate with table theme
type Theme struct {
	ID      int `db:"id" json:"id"`
	ThemeID int `db:"theme_id" json:"theme_id"`
	UserID  int `db:"user_id" json:"user_id"`
}

// Selects all row from table menu item
func (obj *Theme) SelectAll(id int) (output []Theme, err error) {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return
	}

	err = tx.Select("*").
		From("theme").
		Where(dbx.HashExp{"user_id": id}).
		All(&output)
	if err == nil {
		tx.Commit()
	}

	return
}

// Selects single row from table theme
func (obj *Theme) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("theme").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
	}

	return err
}

// Inserts single row in table theme
func (obj *Theme) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	_, err = tx.Insert("theme",
		dbx.Params{
			"theme_id": obj.ThemeID,
			"user_id":  obj.UserID}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table theme
func (obj *Theme) Update(newTheme string) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	_, err = tx.Update("theme",
		dbx.Params{
			"theme_id": newTheme},
		dbx.HashExp{"id": obj.ID, "user_id": obj.UserID}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj Theme) IsInitial() bool {
	if reflect.DeepEqual(obj, Theme{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *Theme) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
