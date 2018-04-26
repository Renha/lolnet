package doctor

import (
	"github.com/lexfrei/lolnet"
)

var _ lolnet.Doctor = &rawBlood{}

type rawBlood struct {
}

func NewRawBlood() rawBlood {
	return rawBlood{}
}

func (rawBlood) Diagnose(blood *lolnet.Blood) error {

	return nil
}
