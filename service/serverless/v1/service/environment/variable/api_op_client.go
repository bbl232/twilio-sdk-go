// This is an autogenerated file. DO NOT MODIFY
package variable

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client         *client.Client
	environmentSid string
	serviceSid     string
	sid            string
}

func New(client *client.Client, environmentSid string, serviceSid string, sid string) *Client {
	return &Client{
		client:         client,
		environmentSid: environmentSid,
		serviceSid:     serviceSid,
		sid:            sid,
	}
}
