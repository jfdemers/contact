# Contact

This project is used to handle contact form from a hugp web site. 
It acts as a server handling the POST request from the web site
and send an email specified in the settings.

# How to run as a daemon

Build the application with go build.
Copy the executable to /usr/local/bin
Copy template.html to /etc/contact
In /etc/contact, create a file settings.yaml with the 
following structure:

```yaml
# The email address used for the smtp server and where the email
# will be sent.
Email: "yourname@server.com" 

# The password used for the smtp server
Password: "yourpassword"

# The email address where the email will be sent.
SendTo: "yourname@server.com"

# The SMTP server to use.
SMTP: "smtp.gmail.com"

# The port number for the SMTP server.
Port: "587"

# The interface and port number for the server to listen.
ServerInterface: ":3001"
```

Copy contact.service to /lib/systemd/system/contact.service

Run the following to enable the service:
sudo systemctl daemon-reload
sudo systemctl enable contact.service
