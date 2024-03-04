package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/wneessen/go-mail"
)

var conf Configuration

type Configuration struct {
	SenderEmail  string
	UserName     string
	UserPassword string
}

func init() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func main() {
	var filePath string
	var recipientEmail string

	fmt.Println("Please enter the path to the file you want to send (e.g. C:\\Program Files\\Go\\README.md):")
	fmt.Scanln(&filePath)
	fmt.Println("Please enter the email address you want to send the file to (e.g. example@example.com):")
	fmt.Scanln(&recipientEmail)

	makeMailMsg(recipientEmail, filePath)
	c := makeClient()
	sendEmail(c)
}

func makeMailMsg(recipientEmail string, filePath string) {
	m := mail.NewMsg()
	m.From(conf.SenderEmail)
	m.To(recipientEmail)
	m.Subject("File")
	m.AttachFile(filePath)

}

func makeClient() mail.Client {
	mail.NewClient("smtp.example.com", mail.WithPort(25), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(conf.UserName), mail.WithPassword(conf.UserPassword))
	return mail.Client{}
}

func sendEmail(c mail.Client) {
	if err := c.DialAndSend(mail.NewMsg()); err != nil {
		log.Fatalf("Failed to send email: %s", err)
	}
}
