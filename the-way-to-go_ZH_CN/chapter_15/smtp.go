package main

import (
	"bytes"
	"log"
	"net/smtp"
)

func main() {
	client, err := smtp.Dial("mail.example.com:25")
	if err != nil {
		log.Fatal(err)
	}
	client.Mail("sender@example.org")
	client.Rcpt("recipient@example.net")
	wc, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("This is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
