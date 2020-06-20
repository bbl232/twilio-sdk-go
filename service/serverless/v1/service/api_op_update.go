// This is an autogenerated file. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateServiceInput struct {
	FriendlyName       *string `form:"FriendlyName,omitempty"`
	IncludeCredentials *bool   `form:"IncludeCredentials,omitempty"`
	UiEditable         *bool   `form:"UiEditable,omitempty"`
}

type UpdateServiceOutput struct {
	Sid                string     `json:"sid"`
	AccountSid         string     `json:"account_sid"`
	FriendlyName       string     `json:"friendly_name"`
	UniqueName         string     `json:"unique_name"`
	IncludeCredentials bool       `json:"include_credentials"`
	UiEditable         bool       `json:"ui_editable"`
	DateCreated        time.Time  `json:"date_created"`
	DateUpdated        *time.Time `json:"date_updated,omitempty"`
	URL                string     `json:"url"`
}

func (c Client) Update(input *UpdateServiceInput) (*UpdateServiceOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateServiceInput) (*UpdateServiceOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Services/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	output := &UpdateServiceOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
