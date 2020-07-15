// This is an autogenerated file. DO NOT MODIFY
package field

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	assistantSid string
	sid          string
	taskSid      string
}

type ClientProperties struct {
	AssistantSid string
	Sid          string
	TaskSid      string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		sid:          properties.Sid,
		taskSid:      properties.TaskSid,
	}
}
