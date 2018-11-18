package schoology

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/garyburd/go-oauth/oauth"
)

var (
	oauthClient oauth.Client
	ctx         context.Context
	emptyCreds  = &oauth.Credentials{}
)

//SetCredentials Set the consumer key and secret to use for authenticating with the API
func SetCredentials(key, secret string) {
	client := &http.Client{}
	ctx = context.WithValue(context.Background(), oauth.HTTPClient, client)
	oauthClient = oauth.Client{
		Credentials: oauth.Credentials{
			Secret: secret,
			Token:  key,
		},
	}

}

//Get Makes an authenticated GET request and returns the raw http response
func Get(url string, form url.Values) (*http.Response, error) {
	return oauthClient.GetContext(ctx, emptyCreds, BaseURL+url, form)
}

func GetSchoolBuildings(id string) ([]Building, int, error) {
	var buildings struct {
		Buildings []Building `json:"building"`
	}

	resp, code, err := getURL(fmt.Sprintf("%s/schools/%s/buildings", BaseURL, id), nil)
	if err != nil {
		return []Building{}, code, err
	}

	err = json.Unmarshal([]byte(resp), &buildings)
	if err != nil {
		return []Building{}, 0, err
	}

	return buildings.Buildings, code, nil
}

//GetSchool Get school by ID
func GetSchool(id string) (School, int, error) {
	resp, code, err := getURL(fmt.Sprintf("%s/schools/%s", BaseURL, id), nil)
	if err != nil {
		return School{}, code, err
	}

	var s School
	err = json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return School{}, 0, err
	}

	return s, code, nil
}
