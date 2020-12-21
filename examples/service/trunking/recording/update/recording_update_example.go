package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/trunking/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/recording"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var trunkingSession *v1.Trunking

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	trunkingSession = twilio.NewWithCredentials(creds).Trunking.V1
}

func main() {
	resp, err := trunkingSession.
		Trunk("TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Recording().
		Update(&recording.UpdateRecordingInput{
			Trim: utils.String("trim-silence"),
		})

	if err != nil {
		log.Panicf("%s", err.Error())
	}

	log.Printf("Trim: %s", resp.Trim)
	log.Printf("Mode: %s", resp.Mode)
}