package main

import (
	"net/http"

	"github.com/lexfrei/lolnet/doctor"
	"github.com/lexfrei/lolnet/donor"
	"github.com/lexfrei/lolnet/recipient"
)

func isOk(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	donor := githubdonor.NewGitHubDonor(http.Client{})
	blood, err := donor.Get()
	isOk(err)
	doc := doctor.NewDoctor()
	isOk(doc.Diagnose(blood))
	recipient := lolout.Newlolout()
	isOk(recipient.Add(blood))
}
