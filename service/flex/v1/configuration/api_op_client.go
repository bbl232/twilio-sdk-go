// This is an autogenerated file. DO NOT MODIFY
package configuration

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client
}

func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}
