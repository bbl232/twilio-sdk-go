// This is an autogenerated file. DO NOT MODIFY
package service

import (
	"context"
	"net/http"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

func (c Client) Delete() error {
	return c.DeleteWithContext(context.Background())
}

func (c Client) DeleteWithContext(context context.Context) error {
	op := client.Operation{
		Method: http.MethodDelete,
		URI:    "/Services/{sid}",
		PathParams: map[string]string{
			"sid": c.sid,
		},
	}

	if err := c.client.Send(context, op, nil, nil); err != nil {
		return err
	}
	return nil
}
