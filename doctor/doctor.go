package doctor

import (
	"net"
	"regexp"

	"github.com/lexfrei/lolnet"
)

// IP and Subnet regexps
var reNet = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\/\d+)`)
var reIP = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

// Check interface
var _ lolnet.Doctor = &doctor{}

// Implement interface
type doctor struct{}

// Mask ipnet
type nets []net.IPNet

func (nets nets) isIPInNets(ip net.IP) bool {
	for _, i := range nets {
		if i.Contains(ip) {
			return true
		}
	}
	return false
}

// NewDoctor gives you doctor object
func NewDoctor() doctor {
	return doctor{}
}

// Dedupe func
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

func (doctor) Diagnose(blood *string) (*lolnet.Blood, error) {
	var cleanIPs []net.IP
	rawNets := unique(reNet.FindAllString(*blood, -1))

	var networks nets
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

	for _, ip := range unique(reIP.FindAllString(*blood, -1)) {
		ipaddress := net.ParseIP(ip)
		if networks.isIPInNets(ipaddress) {
			continue
		} else {
			cleanIPs = append(cleanIPs, ipaddress)
		}
	}

	return &lolnet.Blood{
		IPs:  cleanIPs,
		Nets: networks,
	}, nil
}
