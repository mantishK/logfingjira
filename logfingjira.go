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
	username := flag.String("u", "", "Jira username")
	pass := flag.String("p", "", "Jira password")
	companyName := flag.String("c", "", "Jira company name")

	message := flag.String("m", "", "log message")
	issue := flag.String("i", "", "log issue key/id")
	duration := flag.String("d", "", "log duration")

	flag.Parse()
	envUsername := os.Getenv("JUNAME")
	envCompanyName := os.Getenv("JCOMPNAME")

	if len(*message) == 0 || len(*issue) == 0 || len(*duration) == 0 {
		fmt.Println("-m -i and -d flags must be set")
		return
	}

	if len(*username) == 0 {
		if len(envUsername) == 0 {
			username = fetchUsername()
		} else {
			*username = envUsername
		}
	}
	if len(*pass) == 0 {
		pass = fetchPass()
	}
	if len(*companyName) == 0 {
		if len(envCompanyName) != 0 {
			companyName = &envCompanyName
		} else {
			companyName = fetchCompanyName()
		}
	}
	jiraClientObj := jiraClient{UserName: *username, Pass: *pass, Url: "https://" + *companyName + ".atlassian.net/"}
	jiraClientObj.logHours(*issue, *message, *duration)
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

func fetchCompanyName() *string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter Company name (as needed for the API): ")
	companyName, _ := reader.ReadString('\n')
	companyName = strings.TrimSpace(companyName)
	return &companyName
}

func fetchUsername() *string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	return &username
}
