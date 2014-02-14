package main

import (
	"bufio"
	"code.google.com/p/go.crypto/ssh/terminal"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	username := flag.String("u", "", "help message for flagname")
	pass := flag.String("p", "", "help message for flagname")
	flag.Parse()
	if len(*username) == 0 {
		username = fetchUsername()
	}
	if len(*pass) == 0 {
		pass = fetchPass()
	}
	authorizationCode := authorizationCode(*username, *pass)
	fmt.Println(authorizationCode)
	jiraClientObj := jiraClient{UserName: *username, Pass: *pass, Url: "https://ymedialabs.atlassian.net/"}
	jiraClientObj.getIssues()
}

func authorizationCode(username, pass string) string {
	unamePassByte := []byte(username + ":" + pass)
	unamePass := base64.StdEncoding.EncodeToString(unamePassByte)
	return unamePass
}

func fetchPass() *string {
	fmt.Print("Enter Password: ")
	passwordByte, _ := terminal.ReadPassword(0)
	password := string(passwordByte)
	password = strings.TrimSpace(password)

	return &password
}

func fetchUsername() *string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	return &username
}
