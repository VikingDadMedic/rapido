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
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"framagit.org/InfoLibre/rapido/models/settings"

	"github.com/alecthomas/template"
	"github.com/gin-gonic/gin"
	gomail "gopkg.in/gomail.v2"
)

// Handles HTTP request and uploaded file(s) to send mail
func Action(c *gin.Context) {
	// Parses the Multipart Data sent from Vue
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"message": "incorrect_settings_parameter", "error": err.Error()})
		return
	}

	var attachments []string
	for i := 0; i < 6; i++ {
		_, h, err := c.Request.FormFile("file[" + fmt.Sprint(i) + "]")

		if err != nil {
			fmt.Print(err.Error())
			continue
		}

		// Saves the attachments in temporary folder
		c.SaveUploadedFile(h, "storage/tmp/"+h.Filename)

		// Populates attachment array
		attachments = appendAttachments()

	}
	// Sends the mail
	sendMail(c.Request.PostFormValue("message"), c.Request.PostFormValue("subject"), c.Request.PostFormValue("to"), attachments, c.Request.PostFormValue("email"))
	clearTmp()

}

// Reads tmp folder for uploaded files and appends filepath to attachments[]
func appendAttachments() []string {
	file, err := os.Open("storage/tmp")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	var attachments []string
	i := 0
	for _, name := range list {
		attachments = append(attachments, "storage/tmp/"+name)
		i++

	}
	return attachments
}

// Creates layout and sends it as email
func sendMail(body, subject string, dest string, attachment []string, email string) {
	t := template.New("some template") // Create a template.
	t, _ = t.Parse(body)               // Parse template file.
	mail := gomail.NewMessage()
	mail.SetHeader("Subject", subject)
	mail.SetHeader("From", settings.EmailSender())
	mail.SetHeader("To", dest)
	mail.SetBody("text/html", "<h1>You got a message from Rapido !</h1> </br>"+"<p>"+email+" sent you this message</p>"+"</br>"+body)

	if len(attachment) > 0 {

		for i := range attachment {
			mail.Attach(attachment[i])
		}

	}
	sender := gomail.NewDialer(
		settings.SMTPHost(),
		settings.SMTPPort(),
		settings.SMTPUser(),
		settings.SMTPPassword(),
	)

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

// Clears tmp folder
func clearTmp() {
	// Opens tmp directory
	file, err := os.Open("storage/tmp")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()
	// Reads files names
	list, _ := file.Readdirnames(0)
	// Loops and removes files
	for _, name := range list {
		os.Remove("storage/tmp/" + name)
	}
}
