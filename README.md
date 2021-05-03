# Contact

This project is used to handle contact form from a hugp web site. 
It acts as a server handling the POST request from the web site
and send an email specified in the settings.

# How it works

Build the application and copy template.html in the same folder as
where the binary will run.

Also in the same folder, created a file settings.yaml with the 
following structure:

```yaml
// The email address used for the smtp server and where the email
// will be sent.
Email: "yourname@server.com" 

// The password used for the smtp server
Password: "yourpassword"

// The email address where the email will be sent.
SendTo: "yourname@server.com"

// The SMTP server to use.
SMTP: "smtp.gmail.com"

// The port number for the SMTP server.
Port: "587"

// The interface and port number for the server to listen.
ServerInterface: ":3001"
```