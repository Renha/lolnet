package lolnet

import (
	"fmt"
	"net"
)

// Blood contains some IPs and subnets
type Blood struct {
	Nets []net.IPNet
	IPs  []net.IP
}

// Donor provides list provider
type Donor interface {
	Get() (*Blood, error)
}

// Recipient implements reciver funcs
type Recipient interface {
	Remove() error
	Add(*Blood) error
	Update(*Blood) error
}

// Defibrillator filters blood
type Defibrillator interface {
	Diagnose(*Blood) error
}

func (bl Blood) String() string {
	return fmt.Sprintf("subnets:\t%d\nips:\t\t%d\n", len(bl.Nets), len(bl.IPs))
}
