package main

import (
	"net/http"
)

func getIssues() {
	req, err := http.NewRequest("GET", "http://jira.com/rest/api/", nil)
}
