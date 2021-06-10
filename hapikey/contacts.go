package hapikey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	APIKey     = "HUBSPOT_API_KEY"
	hubSpotURL = "https://api.hubapi.com/crm/v3/%s?%s"
)

// ListContacts gets a list of all your Hubspot CRM contacts
func ListContacts() error {
	contactsReqData := url.Values{}
	contactsReqData.Set("hapikey", os.Getenv(APIKey))

	listContactsURL := fmt.Sprintf(
		hubSpotURL,
		"objects/contacts",
		contactsReqData.Encode(),
	)
	request, err := http.NewRequest(
		http.MethodGet,
		listContactsURL,
		nil,
	)
	if err != nil {
		return err
	}
	request.Header.Set("accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	logrus.Print(string(body))
	return nil
}

// CreateContact creates a new Hubspot CRM contact
func CreateContact() error {
	contactsReqData := url.Values{}
	contactsReqData.Set("hapikey", os.Getenv(APIKey))

	createContactsURL := fmt.Sprintf(
		hubSpotURL,
		"objects/contacts",
		contactsReqData.Encode(),
	)
	contact := map[string]interface{}{
		"properties": map[string]string{
			"company":   "Biglytics",
			"email":     "bcooper@biglytics.net",
			"firstname": "Bryan",
			"lastname":  "Cooper",
			"phone":     "(877) 929-0687",
			"website":   "biglytics.net",
		},
	}
	bs, err := json.Marshal(contact)
	if err != nil {
		return err
	}
	contactBs := bytes.NewBuffer(bs)

	request, err := http.NewRequest(
		http.MethodPost,
		createContactsURL,
		contactBs,
	)
	if err != nil {
		return err
	}
	request.Header.Set("accept", "application/json")
	request.Header.Set("content-type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	logrus.Print(string(body))
	return nil
}
