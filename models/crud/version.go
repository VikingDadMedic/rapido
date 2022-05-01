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
	"io/ioutil"
	"reflect"
)

// Object to communicate with version file
type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

const versionPath = "version.json"

// Selects single row from version file
func (obj *Version) Select() error {
	raw, err := ioutil.ReadFile(versionPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, obj)
}

// Updates single row in version file
func (obj Version) Update() error {
	jsonBlob, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(versionPath, jsonBlob, 0644)
}

// Checks if this object is empty or not
func (obj Version) IsInitial() bool {
	if reflect.DeepEqual(obj, Version{}) {
		return true
	}
	return false
}

// Sets to empty struct
func (obj *Version) clear() {
	p := reflect.ValueOf(obj).Elem()
	p.Set(reflect.Zero(p.Type()))
}
