package main

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"code.google.com/p/go-smtpd/smtpd"
)

type env struct {
        *smtpd.BasicEnvelope
}

func (e *env) AddRecipient(rcpt smtpd.MailAddress) error {
        if strings.HasPrefix(rcpt.Email(), "bad@") {
                return errors.New("we don't send email to bad@")
	}
        return e.BasicEnvelope.AddRecipient(rcpt)
}

func onNewMail(c smtpd.Connection, from smtpd.MailAddress) (smtpd.Envelope, error) {
	log.Printf("ajas: new mail from %q", from)
        return &env{new(smtpd.BasicEnvelope)}, nil
}

func main() {
	go http.ListenAndServe(":8081", nil)

	s := &smtpd.Server{
		Addr:      ":2500",
		OnNewMail: onNewMail,
        }
        err := s.ListenAndServe()
        if err != nil {
                log.Fatalf("ListenAndServe: %v", err)
	}
}
