package main

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

var HOST string
var PORT string
var USERNAME string
var PASSWORD string

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	//  SMTP Config
	HOST = os.Getenv("SMTP_HOST")
	PORT = os.Getenv("SMTP_PORT")
	USERNAME = os.Getenv("SMTP_USERNAME")
	PASSWORD = os.Getenv("SMTP_PASSWORD")

}

// Should return error. Because we going to run this in unit test
func Send() error {
	// Initialize Gomail
	mailer := gomail.NewMessage()

	// Sender
	from := "rneko@dev.test"
	to := "user@cvz.id"

	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Development Purpose!")
	mailer.SetBody("text/html", "Hello this mail from your Developer!")

	// Convert string number to int
	parsedPort, _ := strconv.ParseInt(PORT, 0, 8)
	intPort := int(parsedPort)

	dialer := gomail.NewDialer(HOST, intPort, USERNAME, PASSWORD)

	err := dialer.DialAndSend(mailer)

	if err != nil {
		return err
	}

	return nil
}
