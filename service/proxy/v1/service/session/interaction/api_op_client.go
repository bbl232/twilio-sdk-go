// This is an autogenerated file. DO NOT MODIFY
package interaction

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client     *client.Client
	serviceSid string
	sessionSid string
	sid        string
}

func New(client *client.Client, serviceSid string, sessionSid string, sid string) *Client {
	return &Client{
		client:     client,
		serviceSid: serviceSid,
		sessionSid: sessionSid,
		sid:        sid,
	}
}