// This is an autogenerated file. DO NOT MODIFY
package flow

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/execution"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/executions"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/revision"
	"github.com/RJPearson94/twilio-sdk-go/service/studio/v2/flow/test_users"
)

type Client struct {
	client *client.Client

	sid string

	Revision   func(int) *revision.Client
	TestUsers  func() *test_users.Client
	Executions *executions.Client
	Execution  func(string) *execution.Client
}

type ClientProperties struct {
	Sid string
}

func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,

		Revision: func(revisionNumber int) *revision.Client {
			return revision.New(client, revision.ClientProperties{
				FlowSid:        properties.Sid,
				RevisionNumber: revisionNumber,
			})
		},
		TestUsers: func() *test_users.Client {
			return test_users.New(client, test_users.ClientProperties{
				FlowSid: properties.Sid,
			})
		},
		Executions: executions.New(client, executions.ClientProperties{
			FlowSid: properties.Sid,
		}),
		Execution: func(executionSid string) *execution.Client {
			return execution.New(client, execution.ClientProperties{
				FlowSid: properties.Sid,
				Sid:     executionSid,
			})
		},
	}
}
