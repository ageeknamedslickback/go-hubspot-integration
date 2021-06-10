package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	hubspotTokenURL = "https://api.hubapi.com/oauth/v1/token"
	grantType       = "HUBSPOT_GRANT_TYPE"
	clientID        = "HUBSPOT_CLIENT_ID"
	clientSecret    = "HUBSPOT_CLIENT_SECRET"
	code            = "HUBSPOT_AUTHORIZATION_CODE"
	hubSpotURL      = "https://api.hubapi.com/crm/v3/"
)

type AccessTokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
}

// GetOauthAccessToken returns an access token given an oauth authorization code
func GetOauthAccessToken() (*AccessTokenResponse, error) {
	// TODO: How to handle Authorization Code expiration
	// TODO: How to handle UI account confirmtation
	// TODO: How to handle Access Token expiration
	oauthData := url.Values{}
	oauthData.Set("grant_type", os.Getenv(grantType))
	oauthData.Set("client_id", os.Getenv(clientID))
	oauthData.Set("client_secret", os.Getenv(clientSecret))
	oauthData.Set("redirect_uri", "https://github.com")
	oauthData.Set("code", os.Getenv(code))

	request, err := http.NewRequest(
		http.MethodPost,
		hubspotTokenURL,
		strings.NewReader(oauthData.Encode()),
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var accessToken AccessTokenResponse
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}

func ListContacts() error {
	request, err := http.NewRequest(
		http.MethodGet,
		hubSpotURL+"objects/contacts",
		nil,
	)
	if err != nil {
		return err
	}

	accessToken, err := GetOauthAccessToken()
	if err != nil {
		return err
	}

	bearerToken := fmt.Sprintf("Bearer %s", accessToken.AccessToken)
	request.Header.Set("accept", "application/json")
	request.Header.Set("authorization", bearerToken)

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
