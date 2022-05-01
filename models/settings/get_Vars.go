/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package settings

func SMTPHost() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.SMTPHost
}

func SMTPPort() int {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.SMTPPort
}

func SMTPUser() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.SMTPUser
}

func SMTPPassword() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.SMTPPassword
}

func EmailSender() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.EmailSender
}

func Name() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.Name
}

func HomePage() int {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.HomePage
}

func BannedAttemps() int64 {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	v := int64(obj.Settings.BannedAttemps)
	return v
}

func LockTime() int64 {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	v := int64(obj.Settings.LockTime)
	return v
}

func AccessAttemps() int {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.AccessAttemps
}

func AccessExpired() int64 {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	v := int64(obj.Settings.AccessExpired)
	return v
}

func SecretKey() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.SecretKey
}

func MaxSizeUpload() int64 {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	v := int64(obj.Settings.MaxSizeUpload)
	return v
}

func FaviconURL() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.FaviconURL
}

func LogoURL() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.LogoURL
}

func ActivationEmail() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.ActivationEmail
}

func ActivatedEmail() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.ActivatedEmail
}

func ForgetPasswordEmail() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.ForgetPasswordEmail
}

func PageAuthor() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.PageAuthor
}

func SharePageEmail() string {
	var obj Settings
	if err := obj.Settings.SelectAll(); err != nil || obj.Settings.IsInitial() {
	}
	return obj.Settings.SharePageEmail
}
