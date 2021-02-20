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

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/video"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/recordings"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/room"
	roomRecording "github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/recording"
	roomRecordings "github.com/RJPearson94/twilio-sdk-go/service/video/v1/room/recordings"
	"github.com/RJPearson94/twilio-sdk-go/service/video/v1/rooms"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

var _ = Describe("Video V1", func() {
	creds, err := credentials.New(credentials.Account{
		Sid:       "ACxxx",
		AuthToken: "Test",
	})
	if err != nil {
		log.Panicf("%s", err)
	}

	videoSession := video.New(session.New(creds), &client.Config{
		RetryAttempts: utils.Int(0),
	}).V1

	httpmock.ActivateNonDefault(videoSession.GetClient().GetRestyClient().GetClient())
	defer httpmock.DeactivateAndReset()

	Describe("Given I have a rooms client", func() {
		roomsClient := videoSession.Rooms

		Describe("When the room resource is successfully created", func() {
			createInput := &rooms.CreateRoomInput{}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(201, resp)
				},
			)

			resp, err := roomsClient.Create(createInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the create room response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.VideoCodecs).To(Equal(&[]string{"VP8", "H264"}))
				Expect(resp.MaxParticipants).To(Equal(50))
				Expect(resp.RecordParticipantsOnConnect).To(Equal(false))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.MaxConcurrentPublishedTracks).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.Type).To(Equal("group"))
				Expect(resp.MediaRegion).To(Equal(utils.String("us1")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the create rooms api returns a 500 response", func() {
			createInput := &rooms.CreateRoomInput{}

			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := roomsClient.Create(createInput)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the create room response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the page of rooms are successfully retrieved", func() {
			pageOptions := &rooms.RoomsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := roomsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the rooms page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://video.twilio.com/v1/Rooms?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://video.twilio.com/v1/Rooms?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("rooms"))

				rooms := resp.Rooms
				Expect(rooms).ToNot(BeNil())
				Expect(len(rooms)).To(Equal(1))

				Expect(rooms[0].Sid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(rooms[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(rooms[0].UniqueName).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(rooms[0].Status).To(Equal("in-progress"))
				Expect(rooms[0].VideoCodecs).To(Equal(&[]string{"VP8", "H264"}))
				Expect(rooms[0].MaxParticipants).To(Equal(50))
				Expect(rooms[0].RecordParticipantsOnConnect).To(Equal(false))
				Expect(rooms[0].EndTime).To(BeNil())
				Expect(rooms[0].Duration).To(BeNil())
				Expect(rooms[0].MaxConcurrentPublishedTracks).To(BeNil())
				Expect(rooms[0].StatusCallbackMethod).To(BeNil())
				Expect(rooms[0].StatusCallback).To(BeNil())
				Expect(rooms[0].Type).To(Equal("group"))
				Expect(rooms[0].MediaRegion).To(Equal(utils.String("us1")))
				Expect(rooms[0].DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(rooms[0].DateUpdated).To(BeNil())
				Expect(rooms[0].URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of rooms api returns a 500 response", func() {
			pageOptions := &rooms.RoomsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := roomsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the rooms page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated rooms are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := roomsClient.NewRoomsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated rooms current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated rooms results should be returned", func() {
				Expect(len(paginator.Rooms)).To(Equal(3))
			})
		})

		Describe("When the rooms api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := roomsClient.NewRoomsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated rooms current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a room sid", func() {
		roomClient := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the room resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := roomClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get room resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("in-progress"))
				Expect(resp.VideoCodecs).To(Equal(&[]string{"VP8", "H264"}))
				Expect(resp.MaxParticipants).To(Equal(50))
				Expect(resp.RecordParticipantsOnConnect).To(Equal(false))
				Expect(resp.EndTime).To(BeNil())
				Expect(resp.Duration).To(BeNil())
				Expect(resp.MaxConcurrentPublishedTracks).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.Type).To(Equal("group"))
				Expect(resp.MediaRegion).To(Equal(utils.String("us1")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated).To(BeNil())
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the room resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.Room("RM71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get room response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the room resource is successfully updated", func() {
			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/updateRoomResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			updateInput := &room.UpdateRoomInput{
				Status: "completed",
			}

			resp, err := roomClient.Update(updateInput)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the update room response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.UniqueName).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.VideoCodecs).To(Equal(&[]string{"VP8", "H264"}))
				Expect(resp.MaxParticipants).To(Equal(50))
				Expect(resp.RecordParticipantsOnConnect).To(Equal(false))
				Expect(resp.EndTime.Format(time.RFC3339)).To(Equal("2021-02-20T10:05:00Z"))
				Expect(resp.Duration).To(Equal(utils.Int(320)))
				Expect(resp.MaxConcurrentPublishedTracks).To(BeNil())
				Expect(resp.StatusCallbackMethod).To(BeNil())
				Expect(resp.StatusCallback).To(BeNil())
				Expect(resp.Type).To(Equal("group"))
				Expect(resp.MediaRegion).To(Equal(utils.String("us1")))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.DateUpdated.Format(time.RFC3339)).To(Equal("2021-02-20T10:05:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the update room request does not contain a status", func() {
			updateInput := &room.UpdateRoomInput{}

			resp, err := roomClient.Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectInvalidInputError(err)
			})

			It("Then the update service response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the update room resource api returns a 404", func() {
			httpmock.RegisterResponder("POST", "https://video.twilio.com/v1/Rooms/RM71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			updateInput := &room.UpdateRoomInput{
				Status: "completed",
			}

			resp, err := videoSession.Room("RM71").Update(updateInput)
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the update room response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})
	})

	Describe("Given I have a recordings client", func() {
		recordingsClient := videoSession.Recordings

		Describe("When the page of recordings are successfully retrieved", func() {
			pageOptions := &recordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the recordings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://video.twilio.com/v1/Recordings?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://video.twilio.com/v1/Recordings?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("recordings"))

				recordingGroupingSidsResponse := recordings.PageRecordingGroupingSidsResponse{
					RoomSid:        "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					ParticipantSid: "PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				}

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].Sid).To(Equal("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Status).To(Equal("completed"))
				Expect(recordings[0].GroupingSids).To(Equal(recordingGroupingSidsResponse))
				Expect(recordings[0].ContainerFormat).To(Equal("mka"))
				Expect(recordings[0].TrackName).To(Equal("test"))
				Expect(recordings[0].Offset).To(Equal(171092213859))
				Expect(recordings[0].Codec).To(Equal("opus"))
				Expect(recordings[0].SourceSid).To(Equal("MTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Duration).To(Equal(15))
				Expect(recordings[0].Type).To(Equal("audio"))
				Expect(recordings[0].RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Size).To(Equal(4234))
				Expect(recordings[0].DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(recordings[0].URL).To(Equal("https://video.twilio.com/v1/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of recordings api returns a 500 response", func() {
			pageOptions := &recordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := recordingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the recordings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated recordings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated recordings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated recordings results should be returned", func() {
				Expect(len(paginator.Recordings)).To(Equal(3))
			})
		})

		Describe("When the recordings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := recordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated recordings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a recording sid", func() {
		recordingClient := videoSession.Recording("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the recording resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/recordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := recordingClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get recording resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.GroupingSids).To(Equal(recording.FetchRecordingGroupingSidsResponse{
					RoomSid:        "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					ParticipantSid: "PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				}))
				Expect(resp.ContainerFormat).To(Equal("mka"))
				Expect(resp.TrackName).To(Equal("test"))
				Expect(resp.Offset).To(Equal(171092213859))
				Expect(resp.Codec).To(Equal("opus"))
				Expect(resp.SourceSid).To(Equal("MTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Duration).To(Equal(15))
				Expect(resp.Type).To(Equal("audio"))
				Expect(resp.RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Size).To(Equal(4234))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the recording resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Recordings/RT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.Recording("RT71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the recording resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := recordingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the recording resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Recordings/RT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := videoSession.Recording("RT71").Delete()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})
		})
	})

	Describe("Given I have a room recordings client", func() {
		roomRecordingsClient := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recordings

		Describe("When the page of room recordings are successfully retrieved", func() {
			pageOptions := &roomRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingsPageResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := roomRecordingsClient.Page(pageOptions)
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the room recordings page response should be returned", func() {
				Expect(resp).ToNot(BeNil())

				meta := resp.Meta
				Expect(meta).ToNot(BeNil())
				Expect(meta.Page).To(Equal(0))
				Expect(meta.PageSize).To(Equal(50))
				Expect(meta.FirstPageURL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?PageSize=50&Page=0"))
				Expect(meta.PreviousPageURL).To(BeNil())
				Expect(meta.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?PageSize=50&Page=0"))
				Expect(meta.NextPageURL).To(BeNil())
				Expect(meta.Key).To(Equal("recordings"))

				recordingGroupingSidsResponse := roomRecordings.PageRecordingGroupingSidsResponse{
					RoomSid:        "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					ParticipantSid: "PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				}

				recordings := resp.Recordings
				Expect(recordings).ToNot(BeNil())
				Expect(len(recordings)).To(Equal(1))

				Expect(recordings[0].Sid).To(Equal("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Status).To(Equal("completed"))
				Expect(recordings[0].GroupingSids).To(Equal(recordingGroupingSidsResponse))
				Expect(recordings[0].ContainerFormat).To(Equal("mka"))
				Expect(recordings[0].TrackName).To(Equal("test"))
				Expect(recordings[0].Offset).To(Equal(171092213859))
				Expect(recordings[0].Codec).To(Equal("opus"))
				Expect(recordings[0].SourceSid).To(Equal("MTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Duration).To(Equal(15))
				Expect(recordings[0].Type).To(Equal("audio"))
				Expect(recordings[0].RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(recordings[0].Size).To(Equal(4234))
				Expect(recordings[0].DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(recordings[0].URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the page of room recordings api returns a 500 response", func() {
			pageOptions := &roomRecordings.RecordingsPageOptions{
				PageSize: utils.Int(50),
				Page:     utils.Int(0),
			}

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?Page=0&PageSize=50",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			resp, err := roomRecordingsClient.Page(pageOptions)
			It("Then an error should be returned", func() {
				ExpectInternalServerError(err)
			})

			It("Then the room recordings page response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the paginated room recordings are successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingsPaginatorPage1Response.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			counter := 0
			paginator := roomRecordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then no error should be returned", func() {
				Expect(paginator.Error()).To(BeNil())
			})

			It("Then the paginated room recordings current page should be returned", func() {
				Expect(paginator.CurrentPage()).ToNot(BeNil())
			})

			It("Then the paginated room recordings results should be returned", func() {
				Expect(len(paginator.Recordings)).To(Equal(3))
			})
		})

		Describe("When the room recordings api returns a 500 response when making a paginated request", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingsPaginatorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings?Page=1&PageSize=50&PageToken=abc1234",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/internalServerErrorResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(500, resp)
				},
			)

			counter := 0
			paginator := roomRecordingsClient.NewRecordingsPaginator()

			for paginator.Next() {
				counter++

				if counter > 2 {
					Fail("Too many paginated requests have been made")
				}
			}

			It("Then an error should be returned", func() {
				ExpectInternalServerError(paginator.Error())
			})

			It("Then the paginated room recordings current page should be nil", func() {
				Expect(paginator.CurrentPage()).To(BeNil())
			})
		})
	})

	Describe("Given I have a room recording sid", func() {
		roomRecordingClient := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")

		Describe("When the room recording resource is successfully retrieved", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/roomRecordingResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(200, resp)
				},
			)

			resp, err := roomRecordingClient.Fetch()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})

			It("Then the get room recording resource response should be returned", func() {
				Expect(resp).ToNot(BeNil())
				Expect(resp.Sid).To(Equal("RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.AccountSid).To(Equal("ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Status).To(Equal("completed"))
				Expect(resp.GroupingSids).To(Equal(roomRecording.FetchRecordingGroupingSidsResponse{
					RoomSid:        "RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					ParticipantSid: "PAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
				}))
				Expect(resp.ContainerFormat).To(Equal("mka"))
				Expect(resp.TrackName).To(Equal("test"))
				Expect(resp.Offset).To(Equal(171092213859))
				Expect(resp.Codec).To(Equal("opus"))
				Expect(resp.SourceSid).To(Equal("MTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Duration).To(Equal(15))
				Expect(resp.Type).To(Equal("audio"))
				Expect(resp.RoomSid).To(Equal("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
				Expect(resp.Size).To(Equal(4234))
				Expect(resp.DateCreated.Format(time.RFC3339)).To(Equal("2021-02-20T10:00:00Z"))
				Expect(resp.URL).To(Equal("https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
			})
		})

		Describe("When the room recording resource api returns a 404", func() {
			httpmock.RegisterResponder("GET", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			resp, err := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RT71").Fetch()
			It("Then an error should be returned", func() {
				ExpectNotFoundError(err)
			})

			It("Then the get room recording response should be nil", func() {
				Expect(resp).To(BeNil())
			})
		})

		Describe("When the room recording resource is successfully deleted", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RTXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", httpmock.NewStringResponder(204, ""))

			err := roomRecordingClient.Delete()
			It("Then no error should be returned", func() {
				Expect(err).To(BeNil())
			})
		})

		Describe("When the room recording resource api returns a 404", func() {
			httpmock.RegisterResponder("DELETE", "https://video.twilio.com/v1/Rooms/RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Recordings/RT71",
				func(req *http.Request) (*http.Response, error) {
					fixture, _ := ioutil.ReadFile("testdata/notFoundResponse.json")
					resp := make(map[string]interface{})
					json.Unmarshal(fixture, &resp)
					return httpmock.NewJsonResponse(404, resp)
				},
			)

			err := videoSession.Room("RMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX").Recording("RT71").Delete()
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
	Expect(twilioErr.Message).To(Equal("The requested resource /Rooms/RM71 was not found"))

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
