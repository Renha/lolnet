package lolout

import (
	"fmt"

	"github.com/lexfrei/lolnet"
)

var _ lolnet.Recipient = &lolout{}

type lolout struct {
	short bool
}

// NewStdOut provides new recipient object
func NewStdOut(short bool) *lolout {
	return &lolout{
		short: short,
	}
}

func (lolout) Remove() error {
	return nil
}

func (lo lolout) Add(bl *lolnet.Blood) error {
	if lo.short {
		fmt.Println(bl.String())
	} else {
		var out string
		for _, i := range bl.IPs {
			out += fmt.Sprintf("%s\n", i)
		}
		for _, o := range bl.Nets {
			out += fmt.Sprintf("%s\n", o.String())
		}
		fmt.Println(out)
	}
	return nil
}

func (lo lolout) Update(bl *lolnet.Blood) error {
	if err := lo.Remove(); err != nil {
		return err
	}
	return lo.Add(bl)
}
