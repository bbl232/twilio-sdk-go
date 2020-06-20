// This is an autogenerated file. DO NOT MODIFY
package workflow

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateWorkflowInput struct {
	FriendlyName                  *string `form:"FriendlyName,omitempty"`
	Configuration                 *string `form:"Configuration,omitempty"`
	AssignmentCallbackURL         *string `form:"AssignmentCallbackUrl,omitempty"`
	FallbackAssignmentCallbackURL *string `form:"fallbackAssignmentCallbackUrl,omitempty"`
	TaskReservationTimeout        *int    `form:"TaskReservationTimeout,omitempty"`
	ReEvaluateTasks               *bool   `form:"ReEvaluateTasks,omitempty"`
}

type UpdateWorkflowOutput struct {
	Sid                           string      `json:"sid"`
	AccountSid                    string      `json:"account_sid"`
	WorkspaceSid                  string      `json:"workspace_sid"`
	FriendlyName                  string      `json:"friendly_name"`
	FallbackAssignmentCallbackURL *string     `json:"fallback_assignment_callback_url,omitempty"`
	AssignmentCallbackURL         *string     `json:"assignment_callback_url,omitempty"`
	TaskReservationTimeout        int         `json:"task_reservation_timeout"`
	DocumentContentType           string      `json:"document_content_type"`
	Configuration                 interface{} `json:"configuration"`
	DateCreated                   time.Time   `json:"date_created"`
	DateUpdated                   *time.Time  `json:"date_updated,omitempty"`
	URL                           string      `json:"url"`
}

func (c Client) Update(input *UpdateWorkflowInput) (*UpdateWorkflowOutput, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateWorkflowInput) (*UpdateWorkflowOutput, error) {
	op := client.Operation{
		HTTPMethod:  http.MethodPost,
		HTTPPath:    "/Workspaces/{workspaceSid}/Workflows/{sid}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"workspaceSid": c.workspaceSid,
			"sid":          c.sid,
		},
	}

	output := &UpdateWorkflowOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
