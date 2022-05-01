/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package library

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Creates hash from string
func CreateHash(input string) string {
	hasher, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	fmt.Println("error", err)
	return string(hasher)
}

// Compares input password and hash to see if they match
func CheckPasswordHash(input string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
	return err == nil
}
