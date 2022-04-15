package digitalhumani

import (
	"context"
	"fmt"
	"net/http"
)

type Project struct {
	Created                            string `json:"created"`
	Updated                            string `json:"updated"`
	ID                                 string `json:"id"`
	Name                               string `json:"name"`
	Description                        string `json:"description"`
	ReforestationCompanyName_en        string `json:"reforestationCompanyName_en"`
	ReforestationCompanyName_fr        string `json:"reforestationCompanyName_fr"`
	ReforestationCompanyAddress_en     string `json:"reforestationCompanyAddress_en"`
	ReforestationCompanyAddress_fr     string `json:"reforestationCompanyAddress_fr"`
	ReforestationCompanyWebsite_en     string `json:"reforestationCompanyWebsite_en"`
	ReforestationCompanyWebsite_fr     string `json:"reforestationCompanyWebsite_fr"`
	ReforestationProjectCountry_en     string `json:"reforestationProjectCountry_en"`
	ReforestationProjectCountry_fr     string `json:"reforestationProjectCountry_fr"`
	ReforestationProjectDescription_en string `json:"reforestationProjectDescription_en"`
	ReforestationProjectDescription_fr string `json:"reforestationProjectDescription_fr"`
	ReforestationProjectImageURL_en    string `json:"reforestationProjectImageURL_en"`
	ReforestationProjectImageURL_fr    string `json:"reforestationProjectImageURL_fr"`
	ReforestationProjectState_en       string `json:"reforestationProjectState_en"`
	ReforestationProjectState_fr       string `json:"reforestationProjectState_fr"`
	ReforestationProjectWebsite_en     string `json:"reforestationProjectWebsite_en"`
	ReforestationProjectWebsite_fr     string `json:"reforestationProjectWebsite_fr"`
}

type Projects []struct {
	ID                                 string `json:"id"`
	Name                               string `json:"name"`
	ReforestationProjectDescription_en string `json:"reforestationProjectDescription_en"`
	ReforestationProjectState_en       string `json:"reforestationProjectState_en"`
	ReforestationProjectCountry_en     string `json:"reforestationProjectCountry_en"`
	ReforestationProjectWebsite_en     string `json:"reforestationProjectWebsite_en"`
	ReforestationCompanyName_en        string `json:"reforestationCompanyName_en"`
}

func (digitalhumani *DigitialHumani) GetProject(ctx context.Context, id string) (*Project, error) {
	url := fmt.Sprintf("%s/project/%s", digitalhumani.url, id)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	project := &Project{}
	err = digitalhumani.doAction(ctx, req, project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (digitalhumani *DigitialHumani) GetAllProjects(ctx context.Context) (*Projects, error) {
	url := fmt.Sprintf("%s/project", digitalhumani.url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	projects := &Projects{}
	err = digitalhumani.doAction(ctx, req, projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
