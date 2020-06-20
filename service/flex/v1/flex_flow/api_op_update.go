// This is an autogenerated file. DO NOT MODIFY
package flex_flow

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateFlexFlowInput struct {
	FriendlyName                 *string `form:"FriendlyName,omitempty"`
	ChatServiceSid               *string `form:"ChatServiceSid,omitempty"`
	ChannelType                  *string `form:"ChannelType,omitempty"`
	ContactIdentity              *string `form:"ContactIdentity,omitempty"`
	Enabled                      *bool   `form:"Enabled,omitempty"`
	IntegrationType              *string `form:"IntegrationType,omitempty"`
	IntegrationFlowSid           *string `form:"Integration.FlowSid,omitempty"`
	IntegrationUrl               *string `form:"Integration.Url,omitempty"`
	IntegrationWorkspaceSid      *string `form:"Integration.WorkspaceSid,omitempty"`
	IntegrationChannel           *string `form:"Integration.Channel,omitempty"`
	IntegrationTimeout           *int    `form:"Integration.Timeout,omitempty"`
	IntegrationPriority          *int    `form:"Integration.Priority,omitempty"`
	IntegrationCreationOnMessage *string `form:"Integration.CreationOnMessage,omitempty"`
	IntegrationRetryCount        *int    `form:"Integration.RetryCount,omitempty"`
	LongLived                    *bool   `form:"LongLived,omitempty"`
	JanitorEnabled               *bool   `form:"JanitorEnabled,omitempty"`
}

type UpdateFlexFlowOutputIntegration struct {
	FlowSid           *string `json:"flow_sid,omitempty"`
	Url               *string `json:"url,omitempty"`
	WorkspaceSid      *string `json:"workspace_sid,omitempty"`
	Channel           *string `json:"channel,omitempty"`
	Timeout           *int    `json:"timeout,omitempty"`
	Priority          *int    `json:"priority,omitempty"`
	CreationOnMessage *string `json:"creation_on_message,omitempty"`
	RetryCount        *int    `json:"retry_count,omitempty"`
}

type UpdateFlexFlowOutput struct {
	Sid             string                           `json:"sid"`
	AccountSid      string                           `json:"account_sid"`
	FriendlyName    string                           `json:"friendly_name"`
	ChatServiceSid  string                           `json:"chat_service_sid"`
	ChannelType     string                           `json:"channel_type"`
	ContactIdentity *string                          `json:"contact_identity,omitempty"`
	Enabled         bool                             `json:"enabled"`
	IntegrationType *string                          `json:"integration_type,omitempty"`
	Integration     *UpdateFlexFlowOutputIntegration `json:"integration,omitempty"`
	LongLived       *bool                            `json:"long_lived,omitempty"`
	JanitorEnabled  *bool                            `json:"janitor_enabled,omitempty"`
	DateCreated     time.Time                        `json:"date_created"`
	DateUpdated     *time.Time                       `json:"date_updated,omitempty"`
	URL             string                           `json:"url"`
}

func (c Client) Update(input *UpdateFlexFlowInput) (*UpdateFlexFlowOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateFlexFlowInput) (*UpdateFlexFlowOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/FlexFlows/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	output := &UpdateFlexFlowOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
