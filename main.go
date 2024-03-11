package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/wneessen/go-mail"
)

var conf Configuration

type Configuration struct {
	ServerLink   string
	ServerPort   int
	SenderEmail  string
	UserName     string
	UserPassword string
}

func init() {
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatalf("Error opening conf.json: %s", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatalf("Error decoding conf.json: %s", err)
	}
}

func main() {
	var filePath string
	var recipientEmail string

	fmt.Println("Please enter the path to the file you want to send (e.g. C:\\Program Files\\Go\\README.md):")
	filePath = getInput()
	if filePath == "" {
		log.Fatal("File path cannot be empty.")
	}

	fmt.Println("Please enter the email address you want to send the file to (e.g. example@example.com):")
	recipientEmail = getInput()
	if recipientEmail == "" {
		log.Fatal("Recipient email cannot be empty.")
	}

	msg := makeMailMsg(recipientEmail, filePath)
	c := makeClient()
	sendEmail(c, msg)
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %s", err)
	}
	input = strings.TrimSpace(input) // Remove leading and trailing whitespace
	return input
}

func makeMailMsg(recipientEmail string, filePath string) *mail.Msg {
	m := mail.NewMsg()
	if err := m.From(conf.SenderEmail); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(recipientEmail); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}
	m.Subject("File")
	m.AttachFile(filePath)

	return m
}

func makeClient() mail.Client {
	c, err := mail.NewClient(conf.ServerLink, mail.WithPort(conf.ServerPort), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(conf.UserName), mail.WithPassword(conf.UserPassword))
	if err != nil {
		log.Fatalf("Error creating SMTP client: %s", err)
	}
	return *c
}

func sendEmail(c mail.Client, msg *mail.Msg) {
	if err := c.DialAndSend(msg); err != nil {
		log.Fatalf("Failed to send email: %s", err)
	}
}
