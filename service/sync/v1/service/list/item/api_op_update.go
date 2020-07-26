// This is an autogenerated file. DO NOT MODIFY
package item

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type UpdateListItemInput struct {
	CollectionTtl *int    `form:"CollectionTtl,omitempty"`
	Data          *string `form:"Data,omitempty"`
	ItemTtl       *int    `form:"ItemTtl,omitempty"`
	Ttl           *int    `form:"Ttl,omitempty"`
}

type UpdateListItemResponse struct {
	AccountSid  string                 `json:"account_sid"`
	CreatedBy   string                 `json:"created_by"`
	Data        map[string]interface{} `json:"data"`
	DateCreated time.Time              `json:"date_created"`
	DateExpires *time.Time             `json:"date_expires,omitempty"`
	DateUpdated *time.Time             `json:"date_updated,omitempty"`
	Index       int                    `json:"index"`
	ListSid     string                 `json:"list_sid"`
	Revision    string                 `json:"revision"`
	ServiceSid  string                 `json:"service_Sid"`
	URL         string                 `json:"url"`
}

func (c Client) Update(input *UpdateListItemInput) (*UpdateListItemResponse, error) {
	return c.UpdateWithContext(context.Background(), input)
}

func (c Client) UpdateWithContext(context context.Context, input *UpdateListItemInput) (*UpdateListItemResponse, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/Lists/{listSid}/Items/{index}",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"listSid":    c.listSid,
			"index":      strconv.Itoa(c.index),
		},
	}

	response := &UpdateListItemResponse{}
	if err := c.client.Send(context, op, input, response); err != nil {
		return nil, err
	}
	return response, nil
}