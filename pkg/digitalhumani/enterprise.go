package digitalhumani

import (
	"context"
	"fmt"
	"net/http"
)

// Enterprise represents the details of DigitalHumani enterprise.
type Enterprise struct {
	Created string  `json:"created"`
	Updated string  `json:"updated"`
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Contact Contact `json:"Contact"`
}

// Contact represents the contact name for an enterprise.
type Contact struct {
	Name string `json:"name"`
}

// EnterpriseTreeCount represents the number trees planted by an enterprise.
type EnterpriseTreeCount struct {
	Count int `json:"count"`
}

type dateRange struct {
	StartDate string
	EndDate   string
}

// GetEnterprise method allows you to retrieve the details of your enterprise.
func (digitalhumani *DigitialHumani) GetEnterprise(ctx context.Context) (*Enterprise, error) {
	url := fmt.Sprintf("%s/enterprise/%s", digitalhumani.url, digitalhumani.enterpriseID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	enterprise := &Enterprise{}
	err = digitalhumani.doAction(req, enterprise)
	if err != nil {
		return nil, err
	}

	return enterprise, nil
}

// GetEnterpriseTreeCount method allows you to retrieve the number of trees planted by an enterprise for any range of date.
func (digitalhumani *DigitialHumani) GetEnterpriseTreeCount(ctx context.Context, dateRange *dateRange) (*EnterpriseTreeCount, error) {
	url := fmt.Sprintf("%s/enterprise/%s/treeCount", digitalhumani.url, digitalhumani.enterpriseID)

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
	err = digitalhumani.doAction(req, enterpriseTreeCount)
	if err != nil {
		return nil, err
	}

	return enterpriseTreeCount, nil
}

// GetEnterpriseMonthTreeCount method allows you to retrieve the number of trees planted by an enterprise for a specific month.
func (digitalhumani *DigitialHumani) GetEnterpriseMonthTreeCount(ctx context.Context, month string) (*EnterpriseTreeCount, error) {
	url := fmt.Sprintf("%s/enterprise/%s/treeCount/%s", digitalhumani.url, digitalhumani.enterpriseID, month)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	enterpriseTreeCount := &EnterpriseTreeCount{}
	err = digitalhumani.doAction(req, enterpriseTreeCount)
	if err != nil {
		return nil, err
	}

	return enterpriseTreeCount, nil
}

// NewDateRange returns a struct holding a start and a end date. Useful for the GetEnterpriseTreeCount method.
func NewDateRange(start string, end string) *dateRange { //nolint
	return &dateRange{
		StartDate: start,
		EndDate:   end,
	}
}
