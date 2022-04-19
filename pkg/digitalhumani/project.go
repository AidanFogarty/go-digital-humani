package digitalhumani

import (
	"context"
	"fmt"
	"net/http"
)

// ProjectDetailed represents a detailed view of a DigitalHumani reforestation projects.
type ProjectDetailed struct {
	Created                           string `json:"created"`
	Updated                           string `json:"updated"`
	ID                                string `json:"id"`
	Name                              string `json:"name"`
	Description                       string `json:"description"`
	ReforestationCompanyNameEn        string `json:"reforestationCompanyName_en"`
	ReforestationCompanyNameFr        string `json:"reforestationCompanyName_fr"`
	ReforestationCompanyAddressEn     string `json:"reforestationCompanyAddress_en"`
	ReforestationCompanyAddressFr     string `json:"reforestationCompanyAddress_fr"`
	ReforestationCompanyWebsiteEn     string `json:"reforestationCompanyWebsite_en"`
	ReforestationCompanyWebsiteFr     string `json:"reforestationCompanyWebsite_fr"`
	ReforestationProjectCountryEn     string `json:"reforestationProjectCountry_en"`
	ReforestationProjectCountryFr     string `json:"reforestationProjectCountry_fr"`
	ReforestationProjectDescriptionEn string `json:"reforestationProjectDescription_en"`
	ReforestationProjectDescriptionFr string `json:"reforestationProjectDescription_fr"`
	ReforestationProjectImageURLEn    string `json:"reforestationProjectImageURL_en"`
	ReforestationProjectImageURLFr    string `json:"reforestationProjectImageURL_fr"`
	ReforestationProjectStateEn       string `json:"reforestationProjectState_en"`
	ReforestationProjectStateFr       string `json:"reforestationProjectState_fr"`
	ReforestationProjectWebsiteEn     string `json:"reforestationProjectWebsite_en"`
	ReforestationProjectWebsiteFr     string `json:"reforestationProjectWebsite_fr"`
}

// Project represents reforestation project with a reduced number of attributes. For more information on a given project,
// User GetProject method.
type Project struct {
	ID                                string `json:"id"`
	Name                              string `json:"name"`
	ReforestationProjectDescriptionEn string `json:"reforestationProjectDescription_en"`
	ReforestationProjectStateEn       string `json:"reforestationProjectState_en"`
	ReforestationProjectCountryEn     string `json:"reforestationProjectCountry_en"`
	ReforestationProjectWebsiteEn     string `json:"reforestationProjectWebsite_en"`
	ReforestationCompanyNameEn        string `json:"reforestationCompanyName_en"`
}

// GetProject method allows you to retrieve the details of a single reforestation project.
// Example of an id: 93333333 (Project Ids are 8 digits long).
func (digitalhumani *DigitialHumani) GetProject(ctx context.Context, id string) (*ProjectDetailed, error) {
	url := fmt.Sprintf("%s/project/%s", digitalhumani.url, id)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	project := &ProjectDetailed{}
	err = digitalhumani.doAction(req, project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// GetAllProjects method retrieves all the reforestation projects available.
func (digitalhumani *DigitialHumani) GetAllProjects(ctx context.Context) ([]Project, error) {
	url := fmt.Sprintf("%s/project", digitalhumani.url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	projects := []Project{}
	err = digitalhumani.doAction(req, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
