// This is an autogenerated file. DO NOT MODIFY
package channels

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateChannelInput struct {
	FriendlyName *string    `form:"FriendlyName,omitempty"`
	UniqueName   *string    `form:"UniqueName,omitempty"`
	Attributes   *string    `form:"Attributes,omitempty"`
	Type         *string    `form:"Type,omitempty"`
	DateCreated  *time.Time `form:"DateCreated,omitempty"`
	DateUpdated  *time.Time `form:"DateUpdated,omitempty"`
	CreatedBy    *string    `form:"CreatedBy,omitempty"`
}

type CreateChannelOutput struct {
	Sid           string     `json:"sid"`
	AccountSid    string     `json:"account_sid"`
	ServiceSid    string     `json:"service_sid"`
	FriendlyName  *string    `json:"friendly_name,omitempty"`
	UniqueName    *string    `json:"unique_name,omitempty"`
	Attributes    *string    `json:"attributes,omitempty"`
	Type          string     `json:"type"`
	CreatedBy     string     `json:"created_by"`
	MembersCount  int        `json:"members_count"`
	MessagesCount int        `json:"messages_count"`
	DateCreated   time.Time  `json:"date_created"`
	DateUpdated   *time.Time `json:"date_updated,omitempty"`
	URL           string     `json:"url"`
}

func (c Client) Create(input *CreateChannelInput) (*CreateChannelOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateChannelInput) (*CreateChannelOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Services/{serviceSid}/Channels",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	output := &CreateChannelOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}