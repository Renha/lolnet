package doctor

import (
	"github.com/lexfrei/lolnet"
)

var _ lolnet.Defibrillator = &rawBlood{}

type rawBlood struct {
}

func (rawBlood) Diagnose(blood *lolnet.Blood) error {

	return nil
}
