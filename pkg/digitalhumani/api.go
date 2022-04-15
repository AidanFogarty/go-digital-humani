package digitalhumani

import (
	"encoding/json"
	"net/http"
)

// DigitialHumani represents the DigitialHumani Client for interacting with the API.
type DigitialHumani struct {
	url          string
	apiKey       string
	enterpriseID string
	HTTPClient   *http.Client
}

// New returns a DigitalHumani client, with a default http client.
func New(url string, apiKey string, enterpriseID string) *DigitialHumani {
	return &DigitialHumani{
		url:          url,
		apiKey:       apiKey,
		enterpriseID: enterpriseID,
		HTTPClient:   http.DefaultClient,
	}
}

func (digitalhumani *DigitialHumani) doAction(req *http.Request, dst interface{}) error {
	req.Header.Add("X-Api-Key", digitalhumani.apiKey)

	res, err := digitalhumani.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close() //nolint

	dec := json.NewDecoder(res.Body)
	return dec.Decode(&dst)
}
