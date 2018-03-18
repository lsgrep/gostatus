package addon

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var client = &http.Client{}
var githubNotificationsURL = "https://api.github.com/notifications"
var gitAuthRegex = `https\://([a-zA-Z0-9]+)\:x\-oauth\-basic@github\.com`

type githubNotification struct {
	username string
	token    string
}

// ReadGithubToken reads github personal access token from ~/.git-credentianls
// https://git-scm.com/docs/git-credential-store#_storage_format
func ReadGithubToken() string {
	bs, err := ioutil.ReadFile(fmt.Sprintf("%s/.git-credentials", os.Getenv("HOME")))
	if err != nil {
		panic(err)
	}

	re, err := regexp.Compile(gitAuthRegex)
	if err != nil {
		panic(err)
	}
	result := re.FindStringSubmatch(string(bs))
	if len(result) == 0 {
		return ""
	}
	return result[1]
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (gn *githubNotification) Update() string {
	token := ReadGithubToken()
	if token == "" {
		return ""
	}

	request, err := http.NewRequest("GET", githubNotificationsURL, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", "Basic "+basicAuth(gn.username, token))
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if string(bs) != "[]" {
		return "New Messages"
	}

	// TODO this is way too primitive now.
	return ""
}

func NewGithubNotificationsAddon(username string) *Addon {
	gn := &githubNotification{username: username, token: ReadGithubToken()}
	return &Addon{
		UpdateIntervalMs: 1000 * 30,
		Icon:             "\uf09b",
		Updater:          gn,
	}
}
