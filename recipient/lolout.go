package lolout

import (
	"log"

	"github.com/lexfrei/lolnet"
)

var _ lolnet.Recipient = &lolout{}

type lolout struct {
}

func Newlolout() lolout {
	return struct{}{}
}

func (lolout) Remove() error {
	return nil
}

func (lolout) Add(bl *lolnet.Blood) error {
	log.Println(bl.String())
	return nil
}

func (lolout) Update(bl *lolnet.Blood) error {
	log.Println(bl.String())
	return nil
}
