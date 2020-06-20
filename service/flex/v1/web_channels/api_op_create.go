// This is an autogenerated file. DO NOT MODIFY
package web_channels

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateWebChannelInput struct {
	ChatFriendlyName     string  `validate:"required" form:"ChatFriendlyName"`
	ChatUniqueName       *string `form:"ChatUniqueName,omitempty"`
	CustomerFriendlyName string  `validate:"required" form:"CustomerFriendlyName"`
	FlexFlowSid          string  `validate:"required" form:"FlexFlowSid"`
	Identity             string  `validate:"required" form:"Identity"`
	PreEngagementData    *string `form:"PreEngagementData,omitempty"`
}

type CreateWebChannelOutput struct {
	Sid         string     `json:"sid"`
	AccountSid  string     `json:"account_sid"`
	FlexFlowSid string     `json:"flex_flow_sid"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	URL         string     `json:"url"`
}

func (c Client) Create(input *CreateWebChannelInput) (*CreateWebChannelOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateWebChannelInput) (*CreateWebChannelOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/WebChannels",
		ContentType: client.URLEncoded,
	}

	output := &CreateWebChannelOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
