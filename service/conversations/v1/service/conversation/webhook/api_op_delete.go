// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

// Delete removes a webhook resource from the account
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#delete-a-conversationscopedwebhook-resource for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

// DeleteWithContext removes a webhook resource from the account
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource#delete-a-conversationscopedwebhook-resource for more details
func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{serviceSid}/Conversations/{conversationSid}/Webhooks/{sid}",
		PathParams: map[string]string{
			"serviceSid":      c.serviceSid,
			"conversationSid": c.conversationSid,
			"sid":             c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
