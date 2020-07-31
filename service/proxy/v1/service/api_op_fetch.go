// This is an autogenerated file. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type FetchServiceResponse struct {
	AccountSid              string     `json:"account_sid"`
	CallbackURL             *string    `json:"callback_url,omitempty"`
	ChatInstanceSid         *string    `json:"chat_instance_sid,omitempty"`
	ChatServiceSid          string     `json:"chat_service_sid"`
	DateCreated             time.Time  `json:"date_created"`
	DateUpdated             *time.Time `json:"date_updated,omitempty"`
	DefaultTtl              *int       `json:"default_ttl,omitempty"`
	GeoMatchLevel           *string    `json:"geo_match_level,omitempty"`
	InterceptCallbackURL    *string    `json:"intercept_callback_url,omitempty"`
	NumberSelectionBehavior *string    `json:"number_selection_behavior,omitempty"`
	OutOfSessionCallbackURL *string    `json:"out_of_session_callback_url,omitempty"`
	Sid                     string     `json:"sid"`
	URL                     string     `json:"url"`
	UniqueName              string     `json:"unique_name"`
}

func (c Client) Fetch() (*FetchServiceResponse, error) {
	return c.FetchWithContext(context.Background())
}

func (c Client) FetchWithContext(context context.Context) (*FetchServiceResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	response := &FetchServiceResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}