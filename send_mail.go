package main

// Author: Andreas Pfister <andi.dopamin@gmail.com>

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

// Mail reprents the actual mail settings.
type Mail struct {
	senderID string
	toIds    []string
	subject  string
	body     string
}

// SMTPServer represents the settings of the smtp server
type SMTPServer struct {
	host string
	port string
}

// ServerName builds the servername
func (s *SMTPServer) ServerName() string {
	return s.host + ":" + s.port
}

// BuildMessage builds the message sent through the mail server.
// It sets the sender, receiver, subject and the body of the message.
func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.senderID)
	if len(mail.toIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.toIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.subject)
	message += "\r\n" + mail.body

	return message
}

func main() {

	senderIDPtr := flag.String("sender", "", "The sender of mail (Required).")
	toIdsPtr := flag.String("receiver", "", "Receiver of the mail (Required). \r\n Multiple receivers separated by semicolon. \r\n Mutiple receivers must be within quotes.")
	subjectPtr := flag.String("subject", "", "The subject of the mail")
	bodyPtr := flag.String("message", "", "The message of the mail")
	authPasswortPtr := flag.String("password", "", "The password to authenticate to the smtp server (Required).")
	smtpHostname := flag.String("smtp", "", "The hostname of the smtp server (Required).")
	smtpPort := flag.String("port", "", "The port of the smtp server (Required).")

	flag.Parse()

	if *senderIDPtr == "" ||
		*toIdsPtr == "" ||
		*authPasswortPtr == "" ||
		*smtpHostname == "" ||
		*smtpPort == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	mail := Mail{}
	mail.senderID = *senderIDPtr
	mail.toIds = strings.Split(*toIdsPtr, ";")
	mail.subject = *subjectPtr
	mail.body = *bodyPtr

	messageBody := mail.BuildMessage()

	smtpServer := SMTPServer{host: *smtpHostname, port: *smtpPort}

	log.Println(smtpServer.host)

	auth := smtp.PlainAuth("", mail.senderID, *authPasswortPtr, smtpServer.host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.senderID); err != nil {
		log.Panic(err)
	}
	for _, k := range mail.toIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")

}
