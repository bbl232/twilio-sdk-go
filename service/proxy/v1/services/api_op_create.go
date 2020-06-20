// This is an autogenerated file. DO NOT MODIFY
package services

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreateServiceInput struct {
	UniqueName              string  `validate:"required" form:"UniqueName"`
	DefaultTtl              *int    `form:"DefaultTtl,omitempty"`
	CallbackUrl             *string `form:"CallbackUrl,omitempty"`
	GeoMatchLevel           *string `form:"GeoMatchLevel,omitempty"`
	NumberSelectionBehavior *string `form:"NumberSelectionBehavior,omitempty"`
	InterceptCallbackUrl    *string `form:"InterceptCallbackUrl,omitempty"`
	OutOfSessionCallbackUrl *string `form:"OutOfSessionCallbackUrl,omitempty"`
	ChatInstanceSid         *string `form:"ChatInstanceSid,omitempty"`
}

type CreateServiceOutput struct {
	Sid                     string     `json:"sid"`
	AccountSid              string     `json:"account_sid"`
	ChatInstanceSid         *string    `json:"chat_instance_sid,omitempty"`
	ChatServiceSid          string     `json:"chat_service_sid"`
	UniqueName              string     `json:"unique_name"`
	DefaultTtl              *int       `json:"default_ttl,omitempty"`
	CallbackUrl             *string    `json:"callback_url,omitempty"`
	GeoMatchLevel           *string    `json:"geo_match_level,omitempty"`
	NumberSelectionBehavior *string    `json:"number_selection_behavior,omitempty"`
	InterceptCallbackUrl    *string    `json:"intercept_callback_url,omitempty"`
	OutOfSessionCallbackUrl *string    `json:"out_of_session_callback_url,omitempty"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	URL                     string     `json:"url"`
}

func (c Client) Create(input *CreateServiceInput) (*CreateServiceOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreateServiceInput) (*CreateServiceOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Services",
		ContentType: client.URLEncoded,
	}

	output := &CreateServiceOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
