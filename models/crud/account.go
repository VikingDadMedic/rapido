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

// Object to communicate with table account
type Account struct {
	ID            int                `db:"id" json:"id,omitempty"`
	Password      string             `db:"password" json:"password,omitempty"`
	Active        int                `db:"active" json:"active,omitempty"`
	Pseudonym     string             `db:"pseudonym" json:"pseudonym,omitempty"`
	EmailAddress  string             `db:"email_address" json:"email_address,omitempty"  binding:"required"`
	LastLoggedIn  library.NullInt64  `db:"last_logged_in" json:"last_logged_in,omitempty"`
	AccessLevel   int                `db:"access_level" json:"access_level,omitempty"`
	LastAccessed  library.NullInt64  `db:"last_accessed" json:"last_accessed,omitempty"`
	FailedAttemps library.NullInt64  `db:"failed_attempts" json:"failed_attempts,omitempty"`
	LockUntil     library.NullInt64  `db:"lock_until" json:"lock_until,omitempty"`
	IPAddress     library.NullString `db:"ip_address" json:"ip_address,omitempty"`
	ActivateToken library.NullString `db:"activate_token" json:"activate_token,omitempty"`
	ReminderToken library.NullString `db:"reminder_token" json:"reminder_token,omitempty"`
	ReminderTime  library.NullInt64  `db:"reminder_time" json:"reminder_time,omitempty"`
}

// Selects single row from table account
func (obj *Account) Select(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("account").
		Where(dbx.HashExp{column: value}).
		One(obj)
	if err == nil {
		tx.Commit()
	}

	return err
}

// Selects all row from table account
func (obj *Account) SelectAll() (output []Account, err error) {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return
	}

	err = tx.Select("*").
		From("account").
		All(&output)
	if err == nil {
		tx.Commit()
		for index := range output {
			output[index].Password = ""
		}
	}

	return
}

// Selects by column from table account
func (obj *Account) SelectBy(column string, value ...interface{}) (output []Account) {
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
		From("account").
		Where(exp).
		All(&output)
	if err == nil {
		tx.Commit()
	} else {
		log.Println(err)
	}

	return
}

// Inserts single row in table account
func (obj *Account) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	_, err = tx.Insert("account",
		dbx.Params{
			"password":        obj.Password,
			"active":          obj.Active,
			"pseudonym":       obj.Pseudonym,
			"email_address":   obj.EmailAddress,
			"last_logged_in":  obj.LastLoggedIn,
			"access_level":    obj.AccessLevel,
			"last_accessed":   obj.LastAccessed,
			"failed_attempts": obj.FailedAttemps,
			"lock_until":      obj.LockUntil,
			"ip_address":      obj.IPAddress,
			"activate_token":  obj.ActivateToken,
			"reminder_token":  obj.ReminderToken,
			"reminder_time":   obj.ReminderTime}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table account
func (obj *Account) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	_, err = tx.Update("account",
		dbx.Params{
			"password":        obj.Password,
			"active":          obj.Active,
			"pseudonym":       obj.Pseudonym,
			"email_address":   obj.EmailAddress,
			"last_logged_in":  obj.LastLoggedIn,
			"access_level":    obj.AccessLevel,
			"last_accessed":   obj.LastAccessed,
			"failed_attempts": obj.FailedAttemps,
			"lock_until":      obj.LockUntil,
			"ip_address":      obj.IPAddress,
			"activate_token":  obj.ActivateToken,
			"reminder_token":  obj.ReminderToken,
			"reminder_time":   obj.ReminderTime},
		dbx.HashExp{"id": obj.ID}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Deletes single row in table account
func (obj *Account) Remove(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	_, err = tx.Delete("account", dbx.HashExp{column: value}).Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj Account) IsInitial() bool {
	if reflect.DeepEqual(obj, Account{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *Account) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
