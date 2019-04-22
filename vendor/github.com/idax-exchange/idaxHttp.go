package idax

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Idax HTTP request tool function
//
// Functions only support GET/POST
// The call function automatically parses the request structure and adds signatures
func (irc *IdaxRestConn) IdaxHttp(method string, httpUrl string, reqStruct interface{}) ([]byte, Error) {
	// Initialize Error objects
	var err Error
	// Initialize send string
	var sendJsonStr string
	// Judging whether the structure is empty
	if reqStruct != nil {
		// Determine the request mode POST/GET
		if method == "POST" {
			// Signature Processing
			sendJsonStr = AddSignToJsonStr(reqStruct, irc.Key, irc.Secret)
		} else {
			// Structural body rotation URLCode
			httpUrl += "?" + ToUrlParam(StructToMap(reqStruct))
		}
	}
	fmt.Println("URL:", httpUrl)
	// Create a request object
	req, _ := http.NewRequest(method, httpUrl, strings.NewReader(sendJsonStr))
	// Set no cache
	req.Header.Add("cache-control", "no-cache")
	// Setting content type application/json
	req.Header.Add("content-type", "application/json;charset=utf-8")
	// Send requests
	res, _ := http.DefaultClient.Do(req)
	// Close Body
	defer res.Body.Close()
	// Get response IO
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("res:", string(body))
	// Handling error messages
	json.Unmarshal(body, &err)

	return body, err
}
