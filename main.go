package main

import (
	"github.com/sirupsen/logrus"

	"github.com/ageeknamedslickback/go-hubspot-integration/hapikey"
	"github.com/ageeknamedslickback/go-hubspot-integration/oauth"
)

func main() {
	err := hapikey.CreateContact()
	if err != nil {
		logrus.Print(err)
	}

	logrus.Printf("LIST CONTACTS ...\n")
	err = hapikey.ListContacts()
	if err != nil {
		logrus.Print(err)
	}

	logrus.Print("OAUTH ... \n")
	err = oauth.ListContacts()
	if err != nil {
		logrus.Print(err)
	}
}
