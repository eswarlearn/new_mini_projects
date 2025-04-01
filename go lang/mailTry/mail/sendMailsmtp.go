package mail

import (
	"fmt"
	"net/smtp"
)

func SendMailsmtp(email string) { // Function name starts with uppercase to be accessible in main
	auth := smtp.PlainAuth(
		"",
		"me4717187@gmail.com",
		"bsne aoju xzap egze", // Replace with actual password or app password
		"smtp.gmail.com",
	)

	// Email content
	msg := "Subject: Proposal\n\nLove you Asha Kutty ❤️"

	// Sending email
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"me4717187@gmail.com", // The sender's email
		[]string{email},       // Recipient(s)
		[]byte(msg),
	)

	// Handling success and error cases
	if err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
