// This is an autogenerated file. DO NOT MODIFY
package phone_numbers

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client     *client.Client
	serviceSid string
}

func New(client *client.Client, serviceSid string) *Client {
	return &Client{
		client:     client,
		serviceSid: serviceSid,
	}
}
