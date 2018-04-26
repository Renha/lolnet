package doctor

import (
	"net"

	"github.com/lexfrei/lolnet"
)

var _ lolnet.Doctor = &doctor{}

type doctor struct {
}

// Newdoctor gives you doctor object
func NewDoctor() doctor {
	return doctor{}
}

func (doctor) Diagnose(blood *lolnet.Blood) error {
	keys := make(map[string]struct{})

	IPList := []net.IP{}
	netList := []net.IPNet{}

	for _, entry := range blood.IPs {
		if _, value := keys[entry.String()]; !value {
			keys[entry.String()] = struct{}{}
			IPList = append(IPList, entry)
		}
	}
	blood.IPs = IPList

	for _, nentry := range blood.Nets {
		if _, nvalue := keys[nentry.String()]; !nvalue {
			keys[nentry.String()] = struct{}{}
			netList = append(netList, nentry)
		}
	}
	return nil
}
