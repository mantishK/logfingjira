package main

import (
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

func (jClient *jiraClient) callJiraAPI(apiName string, mothod string, urlParams map[string]string) {
	if len(jClient.AuthorizationCode) == 0 {
		jClient.AuthorizationCode = authorizationCode(jClient.UserName, jClient.Pass)
	}

	resource := "/rest/api/" + apiName
	data := url.Values{}

	//set url parameters
	for key, value := range urlParams {
		data.Set(key, value)
	}

	u, _ := url.ParseRequestURI(jClient.Url)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u) // "https://api.com/user/?name=foo&surname=bar"
	fmt.Println(u)

	client := &http.Client{}
	r, _ := http.NewRequest(mothod, urlStr, nil)
	r.Header.Add("Authorization", "Basic "+jClient.AuthorizationCode)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
	robots, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s", robots)
}

func (jClient *jiraClient) getIssues() {
	params := make(map[string]string)
	params["jql"] = "assignee=\"" + jClient.UserName + "\" and status = \"open\""
	jClient.callJiraAPI("latest/search", "GET", params)
}

func (jiraClient *jiraClient) logHours(issue string, message string, duration int) {
	params := make(map[string]string)
	params["jql"] = "assignee=\"" + jClient.UserName + "\" and status = \"open\""
	jClient.callJiraAPI("issue/"+issue+"/worklog", "POST", params)
}
