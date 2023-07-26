package main

import (
	"os"

	"github.com/joho/godotenv"
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

	HOST = os.Getenv("SMTP_HOST")
	PORT = os.Getenv("SMTP_PORT")
	USERNAME = os.Getenv("SMTP_USERNAME")
	PASSWORD = os.Getenv("SMTP_PASSWORD")

}
