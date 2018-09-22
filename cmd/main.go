package main

import (
	"net/http"

	"github.com/lexfrei/lolnet/doctor"
	"github.com/lexfrei/lolnet/donor/web"
	"github.com/lexfrei/lolnet/recipient/stdout"
)

func isOk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	donor, err := webdonor.NewWebDonor(http.Client{},
		"https://raw.githubusercontent.com/zapret-info/z-i/master/dump.csv")
	isOk(err)
	blood, err := donor.Get()
	isOk(err)
	doc := doctor.NewDoctor()
	cleanBlood, err := doc.Diagnose(blood)
	isOk(err)
	recipient := lolout.NewStdOut(false)
	isOk(recipient.Add(cleanBlood))
}
