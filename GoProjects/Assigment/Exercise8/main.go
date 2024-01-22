package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	email("Success Login")
}

func email(condition string) {
	from := "fromEmailAddress@gmail.com"
	password := "Password key"

	toList := []string{"toEmailAddress@gmail.com"}

	host := "smtp.gmail.com"

	port := "587"

	var msg string

	if condition == "Unauthorized Access" {
		msg = "Unauthorised Access!!!"
	} else if condition == "Success Login" {
		msg = "Success Login!!!"
	} else if condition == "Failure Login" {
		msg = "Failure Login!!!"
	} else {
		msg = "Please provide a valid input!!!"
	}

	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully sent mail to all user in toList")
}
