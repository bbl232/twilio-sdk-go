// This is an autogenerated file. DO NOT MODIFY
package service

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateServiceInput struct {
	FriendlyName                             *string   `form:"FriendlyName,omitempty"`
	DefaultServiceRoleSid                    *string   `form:"DefaultServiceRoleSid,omitempty"`
	DefaultChannelRoleSid                    *string   `form:"DefaultChannelRoleSid,omitempty"`
	DefaultChannelCreatorRoleSid             *string   `form:"DefaultChannelCreatorRoleSid,omitempty"`
	ReadStatusEnabled                        *bool     `form:"ReadStatusEnabled,omitempty"`
	TypingIndicatorTimeout                   *int      `form:"TypingIndicatorTimeout,omitempty"`
	ConsumptionReportInterval                *int      `form:"ConsumptionReportInterval,omitempty"`
	NotificationsNewMessageEnabled           *bool     `form:"Notifications.NewMessage.Enabled,omitempty"`
	NotificationsNewMessageTemplate          *string   `form:"Notifications.NewMessage.Template,omitempty"`
	NotificationsNewMessageSound             *string   `form:"Notifications.NewMessage.Sound,omitempty"`
	NotificationsNewMessageBadgeCountEnabled *bool     `form:"Notifications.NewMessage.BadgeCountEnabled,omitempty"`
	NotificationsAddedToChannelEnabled       *bool     `form:"Notifications.AddedToChannel.Enabled,omitempty"`
	NotificationsAddedToChannelTemplate      *string   `form:"Notifications.AddedToChannel.Template,omitempty"`
	NotificationsAddedToChannelSound         *string   `form:"Notifications.AddedToChannel.Sound,omitempty"`
	NotificationsRemovedToChannelEnabled     *bool     `form:"Notifications.RemovedToChannel.Enabled,omitempty"`
	NotificationsRemovedToChannelTemplate    *string   `form:"Notifications.RemovedToChannel.Template,omitempty"`
	NotificationsRemovedToChannelSound       *string   `form:"Notifications.RemovedToChannel.Sound,omitempty"`
	NotificationsInvitedToChannelEnabled     *bool     `form:"Notifications.InvitedToChannel.Enabled,omitempty"`
	NotificationsInvitedToChannelTemplate    *string   `form:"Notifications.InvitedToChannel.Template,omitempty"`
	NotificationsInvitedToChannelSound       *string   `form:"Notifications.InvitedToChannel.Sound,omitempty"`
	PreWebhookUrl                            *string   `form:"PreWebhookUrl,omitempty"`
	PreWebhookRetryCount                     *int      `form:"PreWebhookRetryCount,omitempty"`
	PostWebhookUrl                           *string   `form:"PostWebhookUrl,omitempty"`
	PostWebhookRetryCount                    *int      `form:"PostWebhookRetryCount,omitempty"`
	WebhookMethod                            *string   `form:"WebhookMethod,omitempty"`
	WebhookFilters                           *[]string `form:"WebhookFilters,omitempty"`
	LimitsChannelMembers                     *int      `form:"Limits.ChannelMembers,omitempty"`
	LimitsUserChannels                       *int      `form:"Limits.UserChannels,omitempty"`
	MediaCompatibilityMessage                *int      `form:"Media.CompatibilityMessage,omitempty"`
	NotificationsLogEnabled                  *bool     `form:"Notifications.LogEnabled,omitempty"`
}

type UpdateServiceOutput struct {
	Sid                          string                 `json:"sid"`
	AccountSid                   string                 `json:"account_sid"`
	ConsumptionReportInterval    int                    `json:"consumption_report_interval"`
	DefaultChannelCreatorRoleSid string                 `json:"default_channel_creator_role_sid"`
	DefaultChannelRoleSid        string                 `json:"default_channel_role_sid"`
	DefaultServiceRoleSid        string                 `json:"default_service_role_sid"`
	FriendlyName                 string                 `json:"friendly_name"`
	Limits                       map[string]interface{} `json:"limits"`
	Media                        map[string]interface{} `json:"media"`
	Notifications                map[string]interface{} `json:"notifications"`
	PostWebhookRetryCount        *int                   `json:"post_webhook_retry_count,omitempty"`
	PostWebhookUrl               *string                `json:"post_webhook_url,omitempty"`
	PreWebhookRetryCount         *int                   `json:"pre_webhook_retry_count,omitempty"`
	PreWebhookUrl                *string                `json:"pre_webhook_url,omitempty"`
	ReachabilityEnabled          bool                   `json:"reachability_enabled"`
	ReadStatusEnabled            bool                   `json:"read_status_enabled"`
	TypingIndicatorTimeout       int                    `json:"typing_indicator_timeout"`
	WebhookFilters               *[]string              `json:"webhook_filters,omitempty"`
	WebhookMethod                *string                `json:"webhook_method,omitempty"`
	DateCreated                  time.Time              `json:"date_created"`
	DateUpdated                  *time.Time             `json:"date_updated,omitempty"`
	URL                          string                 `json:"url"`
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
