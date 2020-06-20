// This is an autogenerated file. DO NOT MODIFY
package service

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_numbers"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/session"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/sessions"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_code"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_codes"
)

type Client struct {
	client       *client.Client
	sid          string
	PhoneNumbers *phone_numbers.Client
	PhoneNumber  func(string) *phone_number.Client
	ShortCodes   *short_codes.Client
	ShortCode    func(string) *short_code.Client
	Sessions     *sessions.Client
	Session      func(string) *session.Client
}

func New(client *client.Client, sid string) *Client {
	return &Client{
		client:       client,
		sid:          sid,
		PhoneNumbers: phone_numbers.New(client, sid),
		PhoneNumber:  func(phoneNumberSid string) *phone_number.Client { return phone_number.New(client, sid, phoneNumberSid) },
		ShortCodes:   short_codes.New(client, sid),
		ShortCode:    func(shortCodeSid string) *short_code.Client { return short_code.New(client, sid, shortCodeSid) },
		Sessions:     sessions.New(client, sid),
		Session:      func(sessionSid string) *session.Client { return session.New(client, sid, sessionSid) },
	}
}
