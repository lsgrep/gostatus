package addon

import (
	"encoding/base64"
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

var httpCli = &http.Client{}
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
		log.Error(err)
		return ""
	}

	re := regexp.MustCompile(gitAuthRegex)
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

func (gn *githubNotification) Update() *Block {
	token := ReadGithubToken()
	if token == "" {
		return nil
	}

	request, err := http.NewRequest("GET", githubNotificationsURL, nil)
	if err != nil {
		log.Error(err)
		return nil
	}
	request.Header.Add("Authorization", "Basic "+basicAuth(gn.username, token))
	response, err := httpCli.Do(request)
	if err != nil {
		log.Error(err)
		return nil
	}

	bs, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return nil
	}

	if string(bs) != "[]" {
		fullTxt := fmt.Sprintf(" %s  %s", IconGithub, "New Messages")
		return &Block{FullText: fullTxt, Color: ColorLime}
	}

	// TODO this is way too primitive now.
	return nil
}

func NewGithubNotificationsAddon(username string) *Addon {
	gn := &githubNotification{username: username, token: ReadGithubToken()}
	return &Addon{
		UpdateInterval: 30 * time.Second,
		Updater:        gn,
	}
}
