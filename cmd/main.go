package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
)

func isOK(err error) {
	if err != nil {
		panic(err)
	}
}

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

var reNet = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\/\d+)`)
var reIP = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

type ips []net.IP
type nets []net.IPNet

func (nets nets) isIPInNets(ip net.IP) bool {
	for _, i := range nets {
		if i.Contains(ip) {
			return true
		}
	}
	return false
}

func main() {
	resp, err := http.Get("https://raw.githubusercontent.com/zapret-info/z-i/master/dump.csv")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(errors.New("Can't"))
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)

	rawNets := unique(reNet.FindAllString(bodyString, -1))
	rawIPs := unique(reIP.FindAllString(bodyString, -1))

	var networks nets
	var ipaddresses ips

	for _, i := range rawNets {
		address, network, err := net.ParseCIDR(i)
		isOK(err)
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

	fmt.Printf("subnets:\t%d\nips:\t\t%d\n", len(networks), len(ipaddresses))
}
