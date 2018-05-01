package main

import (
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/lexfrei/lolnet"
	"github.com/lexfrei/lolnet/doctor"
	"github.com/lexfrei/lolnet/donor/web"
	"github.com/lexfrei/lolnet/recipient/stdout"
)

func isOk(err error) {
	if err != nil {
		panic(err)
	}
}

type config struct {
	Donor     string `required:"true"`
	Recipient string `required:"true"`
}

func main() {
	var conf config
	var donor lolnet.Donor
	isOk(envconfig.Process("lolnet", &conf))
	donor, err := webdonor.NewWebDonor(http.Client{})
	isOk(err)
	blood, err := donor.Get()
	isOk(err)
	doc := doctor.NewDoctor()
	cleanBlood, err := doc.Diagnose(blood)
	isOk(err)
	recipient := lolout.NewStdOut(true)
	isOk(recipient.Add(cleanBlood))
}
