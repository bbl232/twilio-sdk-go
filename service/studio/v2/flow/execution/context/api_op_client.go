// This is an autogenerated file. DO NOT MODIFY
package context

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	flowSid      string
	executionSid string
}

type ClientProperties struct {
	FlowSid      string
	ExecutionSid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		flowSid:      properties.FlowSid,
		executionSid: properties.ExecutionSid,
	}
}
