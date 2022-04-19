package digitalhumani

import (
	"encoding/json"
	"log"
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
func New(apiKey string, enterpriseID string, env string) *DigitialHumani {
	return &DigitialHumani{
		url:          getEnv(env),
		apiKey:       apiKey,
		enterpriseID: enterpriseID,
		HTTPClient:   http.DefaultClient,
	}
}

func getEnv(env string) string {
	switch env {
	case "production":
		return "https://api.digitalhumani.com"
	case "sandbox":
		return "https://api.sandbox.digitalhumani.com"
	default:
		log.Fatalf("unknown env, must be either 'production' or 'sandbox'")
		return ""
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
