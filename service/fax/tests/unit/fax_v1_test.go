package tests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jarcoal/httpmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/RJPearson94/twilio-sdk-go/service/fax"
	faxResource "github.com/RJPearson94/twilio-sdk-go/service/fax/v1/fax"
	"github.com/RJPearson94/twilio-sdk-go/service/fax/v1/faxes"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Fax V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	faxSession := fax.NewWithCredentials(creds).V1

	httpmock.ActivateNonDefault(faxSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given I have a faxes client", func() {
		faxesClient := faxSession.Faxes

		Describe("When the fax resource is successfully created", func() {
			createInput := &faxes.CreateFaxInput{
				To:       "+1987654321",
				MediaURL: "http://localhost/media",
			}

			httpmock.RegisterResponder("POST", "https://fax.twilio.com/v1/Faxes",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/faxResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := faxesClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create fax response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("v1"))
				Expect(resp.Direction).To(Equal("outbound"))
				Expect(resp.From).To(Equal("+1123456789"))
				Expect(resp.MediaURL).To(Equal(utils.String("http://localhost/media")))
				Expect(resp.MediaSid).To(BeNil())
				Expect(resp.NumPages).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(BeNil())
				Expect(resp.Quality).To(Equal("superfine"))
				Expect(resp.Status).To(Equal("queued"))
				Expect(resp.To).To(Equal("+1987654321"))
				Expect(resp.Duration).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the fax request does not contain a to", func() {
			createInput := &faxes.CreateFaxInput{
				MediaURL: "http://localhost/media",
			}

			resp, err := faxesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create fax response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the fax request does not contain a media url", func() {
			createInput := &faxes.CreateFaxInput{
				To: "+1987654321",
			}

			resp, err := faxesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the create fax response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the create faxes api returns a 500 response", func() {
			createInput := &faxes.CreateFaxInput{
				To:       "+1987654321",
				MediaURL: "http://localhost/media",
			}

			httpmock.RegisterResponder("POST", "https://fax.twilio.com/v1/Faxes",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := faxesClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create fax response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a fax sid", func() {
		faxClient := faxSession.Fax("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the fax resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/faxResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := faxClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get fax resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("v1"))
				Expect(resp.Direction).To(Equal("outbound"))
				Expect(resp.From).To(Equal("+1123456789"))
				Expect(resp.MediaURL).To(Equal(utils.String("http://localhost/media")))
				Expect(resp.MediaSid).To(BeNil())
				Expect(resp.NumPages).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(BeNil())
				Expect(resp.Quality).To(Equal("superfine"))
				Expect(resp.Status).To(Equal("queued"))
				Expect(resp.To).To(Equal("+1987654321"))
				Expect(resp.Duration).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the fax resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://fax.twilio.com/v1/Faxes/FX71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := faxSession.Fax("FX71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get fax response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the fax resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateFaxResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &faxResource.UpdateFaxInput{
				Status: utils.String("cancelled"),
			}

			resp, err := faxClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update fax response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.APIVersion).To(Equal("v1"))
				Expect(resp.Direction).To(Equal("outbound"))
				Expect(resp.From).To(Equal("+1123456789"))
				Expect(resp.MediaURL).To(Equal(utils.String("http://localhost/media")))
				Expect(resp.MediaSid).To(BeNil())
				Expect(resp.NumPages).To(BeNil())
				Expect(resp.Price).To(BeNil())
				Expect(resp.PriceUnit).To(BeNil())
				Expect(resp.Quality).To(Equal("superfine"))
				Expect(resp.Status).To(Equal("canceled"))
				Expect(resp.To).To(Equal("+1987654321"))
				Expect(resp.Duration).To(BeNil())
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2020-06-20T20:55:24Z"))
				Expect(resp.URL).To(Equal("https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update fax resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://fax.twilio.com/v1/Faxes/FX71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &faxResource.UpdateFaxInput{
				Status: utils.String("cancelled"),
			}

			resp, err := faxSession.Fax("FX71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update fax response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the fax resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := faxClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the fax resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://fax.twilio.com/v1/Faxes/FX71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := faxSession.Fax("FX71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a media sid", func() {
		mediaClient := faxSession.Fax("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Media("MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the media resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/mediaResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := mediaClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get media resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.FaxSid).To(Equal("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.ContentType).To(Equal("application/pdf"))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2020-06-20T20:50:24Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the media resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/ME71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := faxSession.Fax("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Media("ME71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get media response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the media resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/MEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := mediaClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the media resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://fax.twilio.com/v1/Faxes/FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Media/ME71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := faxSession.Fax("FXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Media("ME71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

})

func ExpectInvalidInputError(err error) {
	ExpectErrorToNotBeATwilioError(err)
	Expect(err.Error()).To(Equal("Invalid input supplied"))
}

func ExpectNotFoundError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))

	code := 20404
	Expect(twilioErr.Code).To(Equal(&code))
	Expect(twilioErr.Message).To(Equal("The requested resource /Faxes/FX71 was not found"))

	moreInfo := "https://www.twilio.com/docs/errors/20404"
	Expect(twilioErr.MoreInfo).To(Equal(&moreInfo))
	Expect(twilioErr.Status).To(Equal(404))
}

func ExpectInternalServerError(err error) {
	Expect(err).ToNot(BeNil())
	twilioErr, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(true))
	Expect(twilioErr.Code).To(BeNil())
	Expect(twilioErr.Message).To(Equal("An error occurred"))
	Expect(twilioErr.MoreInfo).To(BeNil())
	Expect(twilioErr.Status).To(Equal(500))
}

func ExpectErrorToNotBeATwilioError(err error) {
	Expect(err).ToNot(BeNil())
	_, ok := err.(*utils.TwilioError)
	Expect(ok).To(Equal(false))
}