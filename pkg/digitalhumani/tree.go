package digitalhumani

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserTreeCount struct {
	EnterpriseID string `json:"enterpriseId"`
	User         string `json:"user"`
	Count        int    `json:"count"`
}

type PlantTreeRequest struct {
	UUID         string `json:"uuid"`
	Created      string `json:"created"`
	TreeCount    int    `json:"treeCount"`
	EnterpriseID string `json:"enterpriseId"`
	ProjectID    string `json:"projectId"`
	User         string `json:"user"`
}

type newPlantTreeReq struct {
	EnterpriseID string `json:"enterpriseId"`
	ProjectID    string `json:"projectId"`
	User         string `json:"user"`
	TreeCount    int    `json:"treeCount"`
}

func (digitalhumani *DigitialHumani) GetTreeCount(ctx context.Context, userID string) (*UserTreeCount, error) {
	url := fmt.Sprintf("%s/tree", digitalhumani.url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	queryParams := req.URL.Query()
	queryParams.Add("enterpriseId", digitalhumani.enterpriseId)
	queryParams.Add("user", userID)
	req.URL.RawQuery = queryParams.Encode()

	userTreeCount := &UserTreeCount{}
	err = digitalhumani.doAction(ctx, req, userTreeCount)
	if err != nil {
		return nil, err
	}

	return userTreeCount, nil
}

func (digitalhumani *DigitialHumani) PlantTree(ctx context.Context, projectID string, userID string, treeCount int) (*PlantTreeRequest, error) {
	url := fmt.Sprintf("%s/tree", digitalhumani.url)

	newPlantTreeReq := &newPlantTreeReq{
		EnterpriseID: digitalhumani.enterpriseId,
		ProjectID:    projectID,
		User:         userID,
		TreeCount:    treeCount,
	}
	body, _ := json.Marshal(newPlantTreeReq)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	plantTreeRequest := &PlantTreeRequest{}
	err = digitalhumani.doAction(ctx, req, plantTreeRequest)
	if err != nil {
		return nil, err
	}

	return plantTreeRequest, nil
}

func (digitalhumani *DigitialHumani) GetTree(ctx context.Context, uuid string) (*PlantTreeRequest, error) {
	url := fmt.Sprintf("%s/tree/%s", digitalhumani.url, uuid)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	plantTreeRequest := &PlantTreeRequest{}
	err = digitalhumani.doAction(ctx, req, plantTreeRequest)
	if err != nil {
		return nil, err
	}

	return plantTreeRequest, nil
}
