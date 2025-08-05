package helpers

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

func SendMail(to, subject, body string) error {
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	pwd := os.Getenv("SMTP_PWD")
	email := os.Getenv("DEFAULT_EMAIL")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Error parsing PORT:", err)
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(host, port, username, pwd)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending email: %v", err)
		return err
	}
	return nil
}
