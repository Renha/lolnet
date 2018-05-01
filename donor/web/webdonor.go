package webdonor

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"

	"github.com/kelseyhightower/envconfig"
	"github.com/lexfrei/lolnet"
)

var _ lolnet.Donor = &webDonor{}

type webDonor struct {
	client http.Client
	url    url.URL
}

type ips []net.IP

type config struct {
	webLink string `required:"true"`
}

// NewWebDonor gives you github donor object
func NewWebDonor(c http.Client) (*webDonor, error) {
	var wconf config
	err := envconfig.Process("lolnet", &wconf)
	if err != nil {
		return nil, err
	}
	link, err := url.Parse(wconf.webLink)
	if err != nil {
		return nil, err
	}
	return &webDonor{client: c, url: *link}, nil
}

func (wd webDonor) Get() (*string, error) {
	resp, err := wd.client.Get(wd.url.String())
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
