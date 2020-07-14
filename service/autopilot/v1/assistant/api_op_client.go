// This is an autogenerated file. DO NOT MODIFY
package assistant

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/defaults"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/style_sheet"
)

type Client struct {
	client *client.Client

	sid string

	Defaults   func() *defaults.Client
	StyleSheet func() *style_sheet.Client
}

type ClientProperties struct {
	Sid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Defaults: func() *defaults.Client {
			return defaults.New(client, defaults.ClientProperties{
				AssistantSid: properties.Sid,
			})
		},
		StyleSheet: func() *style_sheet.Client {
			return style_sheet.New(client, style_sheet.ClientProperties{
				AssistantSid: properties.Sid,
			})
		},
	}
}