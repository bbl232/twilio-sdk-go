// This is an autogenerated file. DO NOT MODIFY
package versions

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client      *client.Client
	functionSid string
	serviceSid  string
}

func New(client *client.Client, functionSid string, serviceSid string) *Client {
	return &Client{
		client:      client,
		functionSid: functionSid,
		serviceSid:  serviceSid,
	}
}
