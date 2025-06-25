package main

import (
	"flag"
	"fmt"
	"net/smtp"
)

func main() {
	// Define command line flags
	smtpHost := flag.String("host", "email-smtp.us-east-1.amazonaws.com", "SMTP server host")
	smtpPort := flag.String("port", "587", "SMTP server port")
	username := flag.String("username", "", "SMTP username")
	password := flag.String("password", "", "SMTP password")
	from := flag.String("from", "", "From email address")
	to := flag.String("to", "", "To email address")
	subject := flag.String("subject", "Test Email", "Email subject")
	body := flag.String("body", "Hello,\n\nThis is a test email sent from a Go program.\n", "Email body")

	// Parse the flags
	flag.Parse()

	// Validate required parameters
	if *username == "" || *password == "" || *from == "" || *to == "" {
		fmt.Println("Error: username, password, from, and to are required parameters")
		flag.Usage()
		return
	}

	// Convert to slice for the 'to' field
	toSlice := []string{*to}

	// Compose the message
	message := []byte(
		"From: " + *from + "\r\n" +
			"To: " + *to + "\r\n" +
			"Subject: " + *subject + "\r\n" +
			"\r\n" +
			*body + "\r\n")

	// Set up authentication information
	auth := smtp.PlainAuth("", *username, *password, *smtpHost)

	// Send the email
	err := smtp.SendMail(*smtpHost+":"+*smtpPort, auth, *from, toSlice, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
	fmt.Println("Email sent successfully!")
}
