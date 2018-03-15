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

func GotNewNotifications(username string) bool {
	token := ReadGithubToken()
	if token == "" {
		return false
	}

	request, err := http.NewRequest("GET", githubNotificationsURL, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("Authorization", "Basic "+basicAuth(username, token))
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// TODO this is way too primitive now.
	return string(bs) != "[]"
}

func NewGithubNotificationsAddon(username string) *Addon {
	return &Addon{
		UpdateIntervalMs: 1000 * 30,
		UpdateFn: func(a *Addon) {
			var defaultBlock = &Block{FullText: ""}
			hasNew := GotNewNotifications(username)
			if hasNew {
				defaultBlock.FullText = "\uf09b  New Messages"
				defaultBlock.Color = "#00ff00"
			}
			a.LastData = defaultBlock
		},
	}
}
