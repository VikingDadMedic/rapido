/*
Rapido is software to make a website. Rapido is fast, easy to use and respectful of your privacy.
Copyright (C) 2018 Azzam A.I <azzamai91@gmail.com>, Widiyaksa A <widiyaksa@gmail.com>, David VANTYGHEM <david.vantyghem@laposte.net>
Rapido is a complete rewrite of razorCMS from Paul SMITH (https://github.com/smiffy6969).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package email

import (
	"bytes"
	"html/template"
	"log"
	"time"

	"framagit.org/InfoLibre/rapido/models/settings"

	"gopkg.in/gomail.v2"
)

// Creates layout and sends it as email
func SendMail(body, subject string, dest, data map[string]string) {
	t := template.New("some template") // Create a template.
	t, _ = t.Parse(body)               // Parse template file.
	output := bytes.Buffer{}
	t.Execute(&output, data)
	mail := gomail.NewMessage()
	mail.SetHeader("Subject", subject)
	mail.SetHeader("From", settings.EmailSender())
	mail.SetHeader("To", dest["to"])
	if val, ok := dest["cc"]; ok {
		mail.SetHeader("Cc", val)
	}
	if val, ok := dest["bcc"]; ok {
		mail.SetHeader("Bcc", val)
	}
	mail.SetBody("text/html", output.String())

	sender := gomail.NewDialer(
		settings.SMTPHost(),
		settings.SMTPPort(),
		settings.SMTPUser(),
		settings.SMTPPassword())

	counter := 0
	for {
		if err := sender.DialAndSend(mail); err == nil || counter == 5 {
			break
		} else {
			log.Println(err)
		}
		time.Sleep(5 * time.Second)
		counter++
	}
}
