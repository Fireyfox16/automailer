Before starting this program you should change the conf.json file

The ServerLink should be the link you use to access the mail server
The ServerPort should be the port your server uses to send and receive mails
The SenderEmail should be the email you want to send the file from
The UserName should be the username you use to log into the mail server
The UserPassword should be the password you use to log into the mail server

Example:
If you were to use gmail, you would have to add two-step verification. Afterwards in your google settings go to Security > Two-step verification > App Passwords. There you should create a new app password, this password you copy to UserPassword. Your SenderEmail and UserName are both the gmail you use. As for the ServerLink it is smtp.gmail.com and the ServerPort should be one of 25, 465, 587.