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

// Object to communicate with table page_author
type PageCollaboration struct {
	ID        int                `db:"id" json:"id,omitempty"`
	PageID    int                `db:"page_id" json:"page_id,omitempty"`
	AccountID int                `db:"account_id" json:"account_id,omitempty"`
	Email     library.NullString `db:"email" json:"email,omitempty"`
}

// Selects first line with page_id and account_id from table page_author
func (obj *PageCollaboration) Select(pageID, accountID int) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	obj.clear()
	err = tx.Select("*").
		From("page_author").
		Where(dbx.HashExp{"page_id": pageID, "account_id": accountID}).
		One(obj)
	if err == nil {
		tx.Commit()
	}

	return err
}

// Tests if logged in person is a page author
func (obj *PageCollaboration) IsAuthor(pageID, accountID int) bool {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return false
	}

	obj.clear()
	err = tx.Select("*").
		From("page_author").
		Where(dbx.HashExp{"page_id": pageID, "account_id": accountID}).
		One(obj)
	if err == nil {
		tx.Commit()
	}

	if obj.IsInitial() {
		return false
	}

	return true
}

// Tests if page has one or more authors with activated account
func (obj *PageCollaboration) HasAuthor(pageID int) bool {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return false
	}

	obj.clear()
	err = tx.Select("*").
		From("page_author").
		Where(dbx.And(dbx.HashExp{"page_id": pageID}, dbx.NewExp("account_id>0"))).
		One(obj)
	if err == nil {
		tx.Commit()
	}

	if obj.IsInitial() {
		return false
	}

	return true
}

// Checks if data with page_id and email exists in table page_author
func (obj *PageCollaboration) Exist(pageID int, email string) bool {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return false
	}

	obj.clear()
	err = tx.Select("*").
		From("page_author").
		Where(dbx.HashExp{"page_id": pageID, "email": email}).
		One(obj)
	if err == nil {
		tx.Commit()
	}

	if obj.IsInitial() {
		return false
	}

	return true
}

// Selects by column from table page_author
func (obj *PageCollaboration) SelectBy(column string, value ...interface{}) (output []PageCollaboration) {
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
		From("page_author").
		Where(exp).
		All(&output)
	if err == nil {
		tx.Commit()
	} else {
		log.Println(err)
	}

	return
}

// Inserts single row in table page_author
func (obj *PageCollaboration) Insert() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	_, err = tx.Insert("page_author",
		dbx.Params{
			"page_id":    obj.PageID,
			"account_id": obj.AccountID,
			"email":      obj.Email}).
		Execute()
	if err == nil {
		tx.Select("last_insert_rowid() as id").One(obj)
		tx.Commit()
	}

	return err
}

// Updates single row in table page_author
func (obj *PageCollaboration) Update() error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	_, err = tx.Update("page_author",
		dbx.Params{
			"page_id":    obj.PageID,
			"account_id": obj.AccountID,
			"email":      obj.Email},
		dbx.HashExp{"id": obj.ID}).
		Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Deletes single row in table page_author
func (obj *PageCollaboration) Remove(column string, value interface{}) error {
	tx, err := library.DB.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	_, err = tx.Delete("page_author", dbx.HashExp{column: value}).Execute()
	if err == nil {
		tx.Commit()
	}

	return err
}

// Checks if an object is empty or not
func (obj PageCollaboration) IsInitial() bool {
	if reflect.DeepEqual(obj, PageCollaboration{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *PageCollaboration) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
