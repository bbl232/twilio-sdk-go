package main

import (
	"log"
	"os"

	"github.com/RJPearson94/twilio-sdk-go"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/sync/v1"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

var syncSession *v1.Sync

func init() {
	creds, err := credentials.New(credentials.Account{
		Sid:       os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	if err != nil {
		log.Panicf("%s", err.Error())
	}

	syncSession = twilio.NewWithCredentials(creds).Sync.V1
}

func main() {
	paginator := syncSession.
		Service("ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		SyncMap("MPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").
		Items.
		NewSyncMapItemsPaginator()

	for paginator.Next() {
		currentPage := paginator.CurrentPage()
		log.Printf("%v sync map item(s) found on page %v", len(currentPage.SyncMapItems), currentPage.Meta.Page)
	}

	if paginator.Error() != nil {
		log.Panicf("%s", paginator.Error())
	}

	log.Printf("Total number of sync map item(s) found: %v", len(paginator.SyncMapItems))
}
