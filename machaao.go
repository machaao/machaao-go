package machaao

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Get MachaaoAPIToken from https://portal.messengerx.io
var MachaaoAPIToken string = os.Getenv("MachaaoAPIToken")

// Get WitAPIToken from https://wit.ai
var WitAPIToken string = os.Getenv("WitAPIToken")

// for dev, use https://ganglia-dev.machaao.com
var MachaaoBaseURL string = os.Getenv("MachaaoBaseURL")

// This function handles messages
// Input parameters (http.ResponseWriter, *http.Request)
type MessageHandler func(http.ResponseWriter, *http.Request)

// Starts server at given PORT. WebHook is machaao_hook
// input message handler type function(http.ResponseWriter, *http.Request)
func Server(handler MessageHandler) {
	port := GetPort()

	if MachaaoAPIToken == "" {
		log.Fatalln("[ERROR] Machaao API Token not initialised.")
	}
	if MachaaoBaseURL == "" {
		log.Fatalln("[ERROR] Machaao Base URL not initialised.")
	}
	if WitAPIToken == "" {
		log.Println("[WARNING] Wit API Token not initialised.")
	}

	//API handler function
	http.HandleFunc("/machaao_hook", handler)

	//Go http server
	log.Println("[-] Listening on...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}

// Sends post request to MessengerX.io API
func SendPostReq(apiURL string, body interface{}) (*http.Response, error) {

	var url string = MachaaoBaseURL + "/" + apiURL

	//Body converted to json bytes from interface.
	jsonBody, _ := json.Marshal(body)

	//Post request sent to MessengerX.io API
	req, err1 := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	//Sets required headers for MessengerX.io API
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_token", MachaaoAPIToken)

	if err1 != nil {
		log.Println(err1)
		return nil, err1
	}

	client := &http.Client{}
	resp, err2 := client.Do(req)

	return resp, err2
}

// Sends get request to MessengerX.io API
func SendGetReq(apiURL string) (*http.Response, error) {

	var url string = MachaaoBaseURL + "/" + apiURL

	req, err1 := http.NewRequest("GET", url, nil)

	if err1 != nil {
		log.Println(err1)
		return nil, err1
	}

	//Sets required headers for MessengerX.io API
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_token", MachaaoAPIToken)

	client := &http.Client{}
	resp, err2 := client.Do(req)

	return resp, err2
}

// Set PORT as env var or leave it to use 4747
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
