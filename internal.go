package schoology

import (
	"io/ioutil"
	"net/url"
)

// Internal functions used to make authenticated HTTP requests easier

// func buildURL(endpoint string) string {
// 	return fmt.Sprintf("%s/%s", BaseURL, endpoint)
// }

func getURL(url string, formValues url.Values) (string, int, error) {
	resp, err := oauthClient.GetContext(ctx, emptyCreds, url, formValues)
	if err != nil {
		return "", 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, err
	}

	return string(body), resp.StatusCode, nil
}
