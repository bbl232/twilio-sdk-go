// This is an autogenerated file. DO NOT MODIFY
package phone_numbers

import (
	"context"
	"net/http"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
)

type CreatePhoneNumberInput struct {
	Sid         *string `form:"Sid,omitempty"`
	PhoneNumber *string `form:"PhoneNumber,omitempty"`
	IsReserved  *bool   `form:"IsReserved,omitempty"`
}

type CreatePhoneNumberResponseCapabilities struct {
	SmsInbound               *bool `json:"sms_inbound,omitempty"`
	SmsOutbound              *bool `json:"sms_outbound,omitempty"`
	RestrictionSmsDomestic   *bool `json:"restriction_sms_domestic,omitempty"`
	RestrictionVoiceDomestic *bool `json:"restriction_voice_domestic,omitempty"`
	VoiceOutbound            *bool `json:"voice_outbound,omitempty"`
	VoiceInbound             *bool `json:"voice_inbound,omitempty"`
	FaxInbound               *bool `json:"fax_inbound,omitempty"`
	FaxOutbound              *bool `json:"fax_outbound,omitempty"`
	RestrictionFaxDomestic   *bool `json:"restriction_fax_domestic,omitempty"`
	RestrictionMmsDomestic   *bool `json:"restriction_mms_domestic,omitempty"`
	MmsOutbound              *bool `json:"mms_outbound,omitempty"`
	MmsInbound               *bool `json:"mms_inbound,omitempty"`
	SipTrunking              *bool `json:"sip_trunking,omitempty"`
}

type CreatePhoneNumberOutput struct {
	Sid          string                                 `json:"sid"`
	AccountSid   string                                 `json:"account_sid"`
	ServiceSid   string                                 `json:"service_sid"`
	PhoneNumber  *string                                `json:"phone_number,omitempty"`
	FriendlyName *string                                `json:"friendly_name,omitempty"`
	IsoCountry   *string                                `json:"iso_country,omitempty"`
	Capabilities *CreatePhoneNumberResponseCapabilities `json:"capabilities,omitempty"`
	IsReserved   *bool                                  `json:"is_reserved,omitempty"`
	InUse        *int                                   `json:"in_use,omitempty"`
	DateCreated  time.Time                              `json:"date_created"`
	DateUpdated  *time.Time                             `json:"date_updated,omitempty"`
	URL          string                                 `json:"url"`
}

func (c Client) Create(input *CreatePhoneNumberInput) (*CreatePhoneNumberOutput, error) {
	return c.CreateWithContext(context.Background(), input)
}

func (c Client) CreateWithContext(context context.Context, input *CreatePhoneNumberInput) (*CreatePhoneNumberOutput, error) {
	op := client.Operation{
		Method:      http.MethodPost,
		URI:         "/Services/{serviceSid}/PhoneNumbers",
		ContentType: client.URLEncoded,
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
		},
	}

	output := &CreatePhoneNumberOutput{}
	if err := c.client.Send(context, op, input, output); err != nil {
		return nil, err
	}
	return output, nil
}
