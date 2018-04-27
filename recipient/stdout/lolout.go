package lolout

import (
	"fmt"

	"github.com/lexfrei/lolnet"
)

var _ lolnet.Recipient = &lolout{}

type lolout struct {
}

// Newlolout provides new recipient object
func Newlolout() lolout {
	return struct{}{}
}

func (lolout) Remove() error {
	return nil
}

func (lolout) Add(bl *lolnet.Blood) error {
	fmt.Println(bl.String())
	return nil
}

func (l lolout) Update(bl *lolnet.Blood) error {
	l.Remove()
	fmt.Println(bl.String())
	return nil
}
