package webdonor

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/lexfrei/lolnet"
)

var _ lolnet.Donor = &githubDonor{}

type githubDonor struct {
	cleint http.Client
}

type ips []net.IP

// NewGitHubDonor gives you github donor object
func NewGitHubDonor(c http.Client) githubDonor {
	return githubDonor{c}
}

func (githubDonor) Get() (*string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/zapret-info/z-i/master/dump.csv")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Can't get list from GitHub")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)

	return &bodyString, nil
}
