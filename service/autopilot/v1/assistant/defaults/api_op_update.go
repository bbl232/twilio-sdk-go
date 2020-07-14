// This is an autogenerated file. DO NOT MODIFY
package defaults

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateDefaultInput struct {
	Defaults *string `form:"Defaults,omitempty"`
}

type UpdateDefaultResponse struct {
	AccountSid   string      `json:"account_sid"`
	AssistantSid string      `json:"assistant_sid"`
	Data         interface{} `json:"data"`
	URL          string      `json:"url"`
}

func (c Client) Update(input *UpdateDefaultInput) (*UpdateDefaultResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateDefaultInput) (*UpdateDefaultResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Assistants/{assistantSid}/Defaults",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"assistantSid": c.assistantSid,
		},
	}

	response := &UpdateDefaultResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}