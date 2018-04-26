package githubdonor

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"

	"github.com/lexfrei/lolnet"
)

var _ lolnet.Donor = &githubDonor{}

type githubDonor struct {
	cleint http.Client
}

var reNet = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\/\d+)`)
var reIP = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

type ips []net.IP
type nets []net.IPNet

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func (nets nets) isIPInNets(ip net.IP) bool {
	for _, i := range nets {
		if i.Contains(ip) {
			return true
		}
	}
	return false
}

// NewGitHubDonor gives you github donor object
func NewGitHubDonor(c http.Client) githubDonor {
	return githubDonor{c}
}

func (githubDonor) Get() (*lolnet.Blood, error) {
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

	rawNets := unique(reNet.FindAllString(bodyString, -1))
	rawIPs := unique(reIP.FindAllString(bodyString, -1))

	var networks nets
	var ipaddresses ips

	for _, i := range rawNets {
		address, network, err := net.ParseCIDR(i)
		if err != nil {
			return nil, err
		}
		if networks.isIPInNets(address) {
			continue
		} else {
			networks = append(networks, *network)
		}
	}

	for _, i := range rawIPs {
		ipaddress := net.ParseIP(i)
		if networks.isIPInNets(ipaddress) {
			continue
		} else {
			ipaddresses = append(ipaddresses, ipaddress)
		}
	}
	return &lolnet.Blood{
		IPs:  ipaddresses,
		Nets: networks,
	}, nil
}
