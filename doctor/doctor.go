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

func (nets nets) Len() int {
	return len(nets)
}

func (nets nets) Swap(i, j int) {
	nets[i], nets[j] = nets[j], nets[i]
}

func (nets nets) Less(i, j int) bool {
	if len(nets[i].Mask) != len(nets[j].Mask) {
		return len(nets[i].Mask) < len(nets[j].Mask)
	} else {
		for octNum := 0; octNum < len(nets[i].Mask); octNum++ {
			if nets[i].Mask[octNum] != nets[j].Mask[octNum] {
				return nets[i].Mask[octNum] < nets[j].Mask[octNum]
			}
		}
		for octNum := 0; octNum < len(nets[i].IP); octNum++ {
			if nets[i].IP[octNum] != nets[j].IP[octNum] {
				return nets[i].IP[octNum] < nets[j].IP[octNum]
			}
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

type ips []net.IP

func (ips ips) Len() int {
	return len(ips)
}

func (ips ips) Swap(i, j int) {
	ips[i], ips[j] = ips[j], ips[i]
}

func (ips ips) Less(i, j int) bool {
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
	var cleanIPs ips
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
	sort.Sort(networks)

	return &lolnet.Blood{
		IPs:  cleanIPs,
		Nets: networks,
	}, nil
}
