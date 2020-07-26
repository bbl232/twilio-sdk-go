// This is an autogenerated file. DO NOT MODIFY
package revision

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	flowSid        string
	revisionNumber int
}

// The properties required to manage the revision resources
type ClientProperties struct {
	FlowSid        string
	RevisionNumber int
}

// Create a new instance of the client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		flowSid:        properties.FlowSid,
		revisionNumber: properties.RevisionNumber,
	}
}
