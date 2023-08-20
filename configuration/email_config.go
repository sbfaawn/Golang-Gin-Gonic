package configuration

import (
	properties_reader "Golang-Gin-Gonic/properties/reader"

	"gopkg.in/gomail.v2"
)

var Dialer *gomail.Dialer
var sender = "andhikarizki00000@gmail.com"

func init() {
	appPass := properties_reader.Config.AppPassword

	Dialer = gomail.NewDialer("smtp.gmail.com", 587, sender, appPass)
}
