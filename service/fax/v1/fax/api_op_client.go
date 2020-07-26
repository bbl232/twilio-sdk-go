// This is an autogenerated file. DO NOT MODIFY
package fax

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/fax/v1/fax/media"
)

type Client struct {
	client *client.Client

	sid string

	// Sub client to manage media resources
	Media func(string) *media.Client
}

// The properties required to manage the fax resources
type ClientProperties struct {
	Sid string
}

// Create a new instance of the client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Media: func(mediaSid string) *media.Client {
			return media.New(client, media.ClientProperties{
				FaxSid: properties.Sid,
				Sid:    mediaSid,
			})
		},
	}
}
