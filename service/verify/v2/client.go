package v2

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Verify client is used to manage resources for Twilio Verify
// See https://www.twilio.com/docs/verify for more details
type Verify struct {
	client   *client.Client
	Service  func(string) *service.Client
	Services *services.Client
}

// Used for testing purposes only
func (s Verify) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Verify {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "verify"
	config.APIVersion = "v2"

	return NewWithClient(client.New(sess, config))
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Verify {
	return &Verify{
		client: client,
		Service: func(sid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: sid,
			})
		},
		Services: services.New(client),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Verify {
	return New(session.New(creds))
}