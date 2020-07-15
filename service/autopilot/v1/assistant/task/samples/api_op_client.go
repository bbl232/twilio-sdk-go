// This is an autogenerated file. DO NOT MODIFY
package samples

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	assistantSid string
	taskSid      string
}

type ClientProperties struct {
	AssistantSid string
	TaskSid      string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		taskSid:      properties.TaskSid,
	}
}
