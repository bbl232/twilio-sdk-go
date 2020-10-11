// Package service contains auto-generated files. DO NOT MODIFY
package service

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/binding"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/bindings"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/configuration"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/user"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/users"
)

// Client for managing a specific service resource
// See https://www.twilio.com/docs/conversations/api/service-resource for more details
type Client struct {
	client *client.Client

	sid string

	Binding       func(string) *binding.Client
	Bindings      *bindings.Client
	Configuration func() *configuration.Client
	User          func(string) *user.Client
	Users         *users.Client
}

// ClientProperties are the properties required to manage the service resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the service client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Binding: func(bindingSid string) *binding.Client {
			return binding.New(client, binding.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        bindingSid,
			})
		},
		Bindings: bindings.New(client, bindings.ClientProperties{
			ServiceSid: properties.Sid,
		}),
		Configuration: func() *configuration.Client {
			return configuration.New(client, configuration.ClientProperties{
				ServiceSid: properties.Sid,
			})
		},
		User: func(userSid string) *user.Client {
			return user.New(client, user.ClientProperties{
				ServiceSid: properties.Sid,
				Sid:        userSid,
			})
		},
		Users: users.New(client, users.ClientProperties{
			ServiceSid: properties.Sid,
		}),
	}
}
