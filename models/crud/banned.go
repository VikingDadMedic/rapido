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
	"reflect"

	"framagit.org/InfoLibre/rapido/models/library"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// Object to communicate with table banned
type Banned struct {
	ID        int                `db:"id" json:"id,omitempty"`
	IPAddress string             `db:"ip_address" json:"ip_address,omitempty"`
	UserAgent library.NullString `db:"user_agent" json:"user_agent,omitempty"`
}

// Selects single row from table banned
func (obj *Banned) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("banned").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
	}

	return err
}

// Inserts single row in table banned
func (obj *Banned) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = tx.Insert("banned",
		dbx.Params{
			"ip_address": obj.IPAddress,
			"user_agent": obj.UserAgent}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table banned
func (obj *Banned) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = tx.Update("banned",
		dbx.Params{
			"ip_address": obj.IPAddress,
			"user_agent": obj.UserAgent},
		dbx.HashExp{"id": obj.ID}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Deletes single row in table banned
func (obj *Banned) Remove(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = tx.Delete("banned", dbx.HashExp{column: value}).Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj Banned) IsInitial() bool {
	if reflect.DeepEqual(obj, Banned{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *Banned) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
