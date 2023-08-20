package repository

import (
	"Golang-Gin-Gonic/configuration"
	"Golang-Gin-Gonic/model"
	"bytes"
	"fmt"
	"text/template"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

var sender = "andhikarizki00000@gmail.com"

func SendEmail(ctx *gin.Context, account model.Credential) error {
	html := generateHtml(account.Username)

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", account.Email)
	m.SetHeader("Subject", "Account Verification")
	m.SetBody("text/html", html)

	if err := configuration.Dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func generateHtml(username string) string {
	var htmlBody bytes.Buffer
	t, err := template.ParseFiles("../Golang-Gin-Gonic/configuration/verification.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(&htmlBody, struct {
		Username string
	}{
		Username: username,
	})

	return htmlBody.String()
}
