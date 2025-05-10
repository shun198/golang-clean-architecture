package emails

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail() {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_ADMIN"))
	m.SetHeader("To", "to@example.com")
	m.SetHeader("Subject", "ようこそ")
	m.SetBody("text/html", "ようこそ!")
	m.AddAlternative("text/plain", "ようこそ!")

	host := os.Getenv("MAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")

	d := gomail.NewDialer(host, port, username, password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
