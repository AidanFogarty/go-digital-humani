package digitalhumani

import (
	"context"
	"fmt"
	"net/http"
)

type Enterprise struct {
	Created string  `json:"created"`
	Updated string  `json:"updated"`
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Contact Contact `json:"Contact"`
}

type Contact struct {
	Name string `json:"name"`
}

type EnterpriseTreeCount struct {
	Count int `json:"count"`
}

type dateRange struct {
	StartDate string
	EndDate   string
}

func (digitalhumani *DigitialHumani) GetEnterprise(ctx context.Context) (*Enterprise, error) {
	url := fmt.Sprintf("%s/enterprise/%s", digitalhumani.url, digitalhumani.enterpriseId)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	enterprise := &Enterprise{}
	err = digitalhumani.doAction(ctx, req, enterprise)
	if err != nil {
		return nil, err
	}

	return enterprise, nil
}

func (digitalhumani *DigitialHumani) GetEnterpriseTreeCount(ctx context.Context, dateRange *dateRange) (*EnterpriseTreeCount, error) {
	url := fmt.Sprintf("%s/enterprise/%s/treeCount", digitalhumani.url, digitalhumani.enterpriseId)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if dateRange != nil {
		queryParams := req.URL.Query()
		queryParams.Add("startDate", dateRange.StartDate)
		queryParams.Add("endDate", dateRange.EndDate)
		req.URL.RawQuery = queryParams.Encode()
	}

	fmt.Println(req.URL.RawQuery)

	enterpriseTreeCount := &EnterpriseTreeCount{}
	err = digitalhumani.doAction(ctx, req, enterpriseTreeCount)
	if err != nil {
		return nil, err
	}

	return enterpriseTreeCount, nil
}

func (digitalhumani *DigitialHumani) GetEnterpriseMonthTreeCount(ctx context.Context, month string) (*EnterpriseTreeCount, error) {
	url := fmt.Sprintf("%s/enterprise/%s/treeCount/%s", digitalhumani.url, digitalhumani.enterpriseId, month)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	enterpriseTreeCount := &EnterpriseTreeCount{}
	err = digitalhumani.doAction(ctx, req, enterpriseTreeCount)
	if err != nil {
		return nil, err
	}

	return enterpriseTreeCount, nil
}

func NewDateRange(start string, end string) *dateRange {
	return &dateRange{
		StartDate: start,
		EndDate:   end,
	}
}
