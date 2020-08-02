package machaao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//MachaaoAPIToken Get MachaaoAPIToken from https://portal.messengerx.io
var MachaaoAPIToken string = os.Getenv("MachaaoAPIToken")

//WitAPIToken Get WitAPIToken from https://wit.ai
var WitAPIToken string = os.Getenv("WitAPIToken")

//MachaaoBaseURL for dev, use https://ganglia-dev.machaao.com
var MachaaoBaseURL string = os.Getenv("MachaaoBaseURL")

//MessageHandler This function handles messages
//MessageHandler Input parameters (http.ResponseWriter, *http.Request)
type MessageHandler func(http.ResponseWriter, *http.Request)

//Server Starts server at given PORT. WebHook is machaao_hook
//Server input message handler type function(http.ResponseWriter, *http.Request)
func Server(handler MessageHandler) {
	port := GetPort()

	if WitAPIToken == "" {
		log.Fatalln("Wit API Token not initialised.")
	}
	if MachaaoAPIToken == "" {
		log.Fatalln("Machaao API Token not initialised.")
	}
	if MachaaoBaseURL == "" {
		log.Fatalln("Machaao Base URL not initialised.")
	}

	//API handler function
	http.HandleFunc("/machaao_hook", handler)

	//Go http server
	log.Println("[-] Listening on...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}

}

//SendPostReq Send post request
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

//SendGetReq Send get request
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

//GetPort Set PORT as env var or leave it to use 4747
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4747"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
