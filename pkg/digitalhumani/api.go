package digitalhumani

import (
	"context"
	"encoding/json"
	"net/http"
)

type DigitialHumani struct {
	url          string
	apiKey       string
	enterpriseId string
	HttpClient   *http.Client
}

func New(url string, apiKey string, enterpriseId string) *DigitialHumani {
	return &DigitialHumani{
		url:          url,
		apiKey:       apiKey,
		enterpriseId: enterpriseId,
		HttpClient:   http.DefaultClient,
	}
}

func (digitalhumani *DigitialHumani) doAction(ctx context.Context, req *http.Request, dst interface{}) error {
	req.Header.Add("X-Api-Key", digitalhumani.apiKey)

	res, err := digitalhumani.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close() //nolint

	dec := json.NewDecoder(res.Body)
	return dec.Decode(&dst)
}
