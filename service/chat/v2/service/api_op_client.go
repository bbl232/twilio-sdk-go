// This is an autogenerated file. DO NOT MODIFY
package service

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/binding"
)

type Client struct {
	client  *client.Client
	sid     string
	Binding func(string) *binding.Client
}

func New(client *client.Client, sid string) *Client {
	return &Client{
		client:  client,
		sid:     sid,
		Binding: func(bindingSid string) *binding.Client { return binding.New(client, sid, bindingSid) },
	}
}
