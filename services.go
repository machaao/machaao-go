package machaao

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	witai "github.com/wit-ai/wit-go"
)

// Send a subscription message announcement to a tag or a timezone
func SendAnnouncement(body interface{}) (*http.Response, error) {
	var apiURL string = "v1/messages/announce"

	return SendPostReq(apiURL, body)
}

// Send a standard message to a single user id or multiple user id
func SendMessage(body interface{}) (*http.Response, error) {
	var apiURL string = "v1/messages/send"

	return SendPostReq(apiURL, body)
}

// Insert or update content for your bot
func AddContent(body interface{}) (*http.Response, error) {
	var apiURL string = "v1/content"

	return SendPostReq(apiURL, body)
}

// Tag or Un-tag a specific userId
func TagUser(userID string, body interface{}) (*http.Response, error) {
	var apiURL string = "v1/users/tag/" + userID

	return SendPostReq(apiURL, body)
}

// Get basic profile of the user
func GetUserProfile(userID string) (*http.Response, error) {
	var apiURL string = "v1/users/" + userID

	return SendGetReq(apiURL)
}

//GetUserTag Get all tags for a specific userID
func GetUserTag(userID string) (*http.Response, error) {
	var apiURL string = "v1/users/tags/" + userID

	return SendGetReq(apiURL)
}

//SearchContent Search content on your bot
func SearchContent(query string) (*http.Response, error) {
	var url string = MachaaoBaseURL + "v1/content/search/"

	req, err1 := http.NewRequest("GET", url, nil)

	if err1 != nil {
		log.Println(err1)
		return nil, err1
	}

	//Sets required headers for MessengerX.io API
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_token", MachaaoAPIToken)
	req.Header.Set("q", query)

	client := &http.Client{}
	resp, err2 := client.Do(req)

	if err2 != nil {
		log.Println(err2)
		return nil, err2
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	bodyf, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(bodyf))

	return resp, nil
}

// Get content for your bot via an unique slug
func SearchContentViaSlug(slug string) (*http.Response, error) {

	var urlAPI string = MachaaoBaseURL + "v1/content/" + slug

	return SendGetReq(urlAPI)
}

//WitAIResponse Get reponse for the input message
func WitAIResponse(message string) (*witai.MessageResponse, error) {
	client := witai.NewClient(WitAPIToken)
	// Use client.SetHTTPClient() to set custom http.Client

	msg, err := client.Parse(&witai.MessageRequest{
		Query: message,
	})

	return msg, err
}
