package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"text/template"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	Email      string `yaml:"Email"`
	Password   string `yaml:"Password"`
	SendTo     string `yaml:"SendTo"`
	SMTP       string `yaml:"SMTP"`
	Port       string `yaml:"Port"`
	ServerPort string `yaml:"ServerInterface"`
}

func handler(settings *Settings) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("courriel")
		message := r.FormValue("message")
		subject := r.FormValue("sujet")

		w.Header().Set("Content Type", "application/text")
		w.Header().Set("Cache-Control", "no-store")

		if len(subject) > 0 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Ok"))
		} else {
			from := settings.Email
			password := settings.Password

			to := []string{settings.Email}
			smtpHost := settings.SMTP
			smtpPort := settings.Port

			t, err := template.ParseFiles("template.html")
			if err != nil {
				log.Println("Error processing template.")
			}

			var body bytes.Buffer

			mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
			body.Write([]byte(fmt.Sprintf("Subject: Received message from web site. \n%s\n\n", mimeHeaders)))

			t.Execute(&body, struct {
				EMail   string
				Message string
			}{
				EMail:   email,
				Message: message,
			})

			auth := smtp.PlainAuth("", from, password, smtpHost)

			err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())

			if err != nil {
				log.Printf("Error sending email: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Error"))
			} else {
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}

func startServer(settings *Settings) (*http.Server, error) {
	router := mux.NewRouter()
	server := &http.Server{}

	router.HandleFunc("/", handler(settings)).Methods("POST")

	server.Addr = settings.ServerPort
	loggingHandler := handlers.LoggingHandler(os.Stdout, router)

	server.Handler = loggingHandler

	return server, server.ListenAndServe()
}

func getSettings() (*Settings, error) {
	data, err := ioutil.ReadFile("./settings.yaml")
	if err != nil {
		return nil, err
	}

	settings := Settings{}
	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}

func main() {
	println("Starting contact form server.")

	settings, err := getSettings()
	if err != nil {
		log.Panicf("Error loading settings.")
	}

	_, err = startServer(settings)

	if err != nil {
		log.Panicf("Can't start server. Error: %v", err)
	}
}
