package doctor

import (
	"net"
	"regexp"
	"sort"

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

type IPs []net.IP

func (ips IPs) Len() int {
	return len(ips)
}

func (ips IPs) Swap(i, j int) {
	ips[i], ips[j] = ips[j], ips[i]
}

func (ips IPs) Less(i, j int) bool {
	if len(ips[i]) != len(ips[j]) {
		return len(ips[i]) < len(ips[j])
	}
	for octNum := 0; octNum < len(ips[i]); octNum++ {
		if ips[i][octNum] != ips[j][octNum] {
			return ips[i][octNum] < ips[j][octNum]
		}
	}
	return false
}

func (doctor) Diagnose(blood *string) (*lolnet.Blood, error) {
	var cleanIPs IPs
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

	sort.Sort(cleanIPs)

	return &lolnet.Blood{
		IPs:  cleanIPs,
		Nets: networks,
	}, nil
}
