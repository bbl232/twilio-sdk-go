package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/notify/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/service/bindings"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/google/uuid"
)

var notifySession *v1.Notify

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	notifySession = twilio.NewWithCredentials(creds).Notify.V1
}

func main() {
	resp, err := notifySession.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Bindings.
		Create(&bindings.CreateBindingInput{
			Identity:    uuid.New().String(),
			BindingType: "sms",
			Address:     "+10123456789",
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("SID: %s", resp.Sid)
}
