package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type jiraClient struct {
	UserName          string
	Pass              string
	Url               string
	AuthorizationCode string
}

func (jClient *jiraClient) callJiraAPI(apiName string, method string, urlParams, postparams map[string]string) []byte {
	if len(jClient.AuthorizationCode) == 0 {
		jClient.AuthorizationCode = authorizationCode(jClient.UserName, jClient.Pass)
	}

	resource := "/rest/api/latest/" + apiName
	data := url.Values{}

	//set url parameters
	for key, value := range urlParams {
		data.Set(key, value)
	}

	u, _ := url.ParseRequestURI(jClient.Url)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/?name=foo&surname=bar"

	client := &http.Client{}

	postparamsByte, err := json.Marshal(postparams)
	if err != nil {
		panic("json parse error")
	}
	postParamsReader := bytes.NewReader(postparamsByte)
	r, _ := http.NewRequest(method, urlStr, postParamsReader)
	r.Header.Add("Authorization", "Basic "+jClient.AuthorizationCode)
	r.Header.Add("Content-Type", "Application/json")
	r.Header.Add("Accept", "Application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, _ := client.Do(r)
	respByte, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return respByte
}

func (jClient *jiraClient) getIssues() {
	params := make(map[string]string)
	params["jql"] = "assignee=\"" + jClient.UserName + "\" and status = \"open\""
	responseByte := jClient.callJiraAPI("latest/search", "GET", params, nil)
	var responseInterface interface{}
	err := json.Unmarshal(responseByte, &responseInterface)
	if err != nil {
		fmt.Println("Some error occured")
	} else {
		fmt.Println(responseInterface)
	}
}

func (jClient *jiraClient) logHours(issue string, message string, duration string) {
	postParams := make(map[string]string)
	postParams["timeSpent"] = duration
	postParams["comment"] = message
	responseByte := jClient.callJiraAPI("issue/"+issue+"/worklog", "POST", nil, postParams)
	var responseInterface map[string]interface{}
	err := json.Unmarshal(responseByte, &responseInterface)
	if err != nil {
		fmt.Println("Some error occured")
	} else {
		id, _ := responseInterface["id"].(string)
		fmt.Println("Successfully created the worklog to impress your manager. \nWorklog id = " + id)
	}

}
