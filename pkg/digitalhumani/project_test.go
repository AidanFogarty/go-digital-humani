package digitalhumani

import (
	"context"
	"testing"

	"github.com/AidanFogarty/go-digital-humani/pkg/testutil"
)

const (
	successfulGetProject = `
	{
		"reforestationCompanyName_fr": "WeForest",
		"reforestationProjectImageURL_en": "https://www.weforest.org/sites/IMG_20190423_132725_0.jpg",
		"reforestationCompanyName_en": "WeForest",
		"reforestationProjectCountry_en": "India",
		"reforestationCompanyAddress_en": "Ogentroostlaan 15, 3090 Overijse, Belgium",
		"created": "2018-12-12T09:05:00.725Z",
		"reforestationProjectWebsite_en": "https://www.weforest.org/project/india-khasi-hills",
		"name": "Khasi Hills in India, WeForest",
		"reforestationProjectWebsite_fr": "https://www.weforest.org/project/india-khasi-hills",
		"reforestationProjectCountry_fr": "Inde",
		"updated": "2019-05-19T19:24:10.761Z",
		"reforestationProjectDescription_fr": "Projet de reforestation aux Khasi Hills en Inde",
		"reforestationProjectDescription_en": "Reforestation project in the Khasi Hills in India",
		"reforestationCompanyWebsite_fr": "https://www.weforest.org/",
		"reforestationCompanyWebsite_en": "https://www.weforest.org/",
		"reforestationCompanyAddress_fr": "Ogentroostlaan 15, 3090 Overijse, Belgique",
		"description": "Khasi Hills in India, WeForest",
		"id": "96666666",
		"reforestationProjectImageURL_fr": "https://www.weforest.org/sites/default/IMG_20190423_132725_0.jpg"
	}
	`
	successfulGetAllProjects = `
	[
		{
			"id": "91111111",
			"name": "California OneTreePlanted",
			"reforestationProjectDescription_en": "Reforestation project in California, United States",
			"reforestationProjectState_en": "California",
			"reforestationProjectCountry_en": "United States",
			"reforestationProjectWebsite_en": "https://onetreeplanted.org/united-states/products/california-forests",
			"reforestationCompanyName_en": "OneTreePlanted"
		},
		{
			"id": "96666666",
			"name": "Khasi Hills in India, WeForest",
			"reforestationProjectDescription_en": "Reforestation project in the Khasi Hills in India",
			"reforestationProjectState_en": "Khasi Hills",
			"reforestationProjectCountry_en": "India",
			"reforestationProjectWebsite_en": "https://www.weforest.org/project/india-khasi-hills",
			"reforestationCompanyName_en": "WeForest"
		}
	]`
)

func TestDigitialHumani_GetProject(t *testing.T) {
	type args struct {
		ctx       context.Context
		projectID string
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     *ProjectDetailed
		wantErr  bool
	}{
		{
			name:     "Successful GetProject 200",
			args:     args{context.TODO(), "96666666"},
			response: successfulGetProject,
			want: &ProjectDetailed{Created: "2018-12-12T09:05:00.725Z", Updated: "2019-05-19T19:24:10.761Z", ID: "96666666", Name: "Khasi Hills in India, WeForest", Description: "Khasi Hills in India, WeForest",
				ReforestationCompanyNameEn: "WeForest", ReforestationCompanyNameFr: "WeForest", ReforestationCompanyAddressEn: "Ogentroostlaan 15, 3090 Overijse, Belgium", ReforestationCompanyAddressFr: "Ogentroostlaan 15, 3090 Overijse, Belgique",
				ReforestationCompanyWebsiteEn: "https://www.weforest.org/", ReforestationCompanyWebsiteFr: "https://www.weforest.org/", ReforestationProjectCountryEn: "India", ReforestationProjectCountryFr: "Inde",
				ReforestationProjectDescriptionEn: "Reforestation project in the Khasi Hills in India", ReforestationProjectDescriptionFr: "Projet de reforestation aux Khasi Hills en Inde", ReforestationProjectImageURLEn: "https://www.weforest.org/sites/IMG_20190423_132725_0.jpg",
				ReforestationProjectImageURLFr: "https://www.weforest.org/sites/default/IMG_20190423_132725_0.jpg", ReforestationProjectStateEn: "", ReforestationProjectStateFr: "",
				ReforestationProjectWebsiteEn: "https://www.weforest.org/project/india-khasi-hills", ReforestationProjectWebsiteFr: "https://www.weforest.org/project/india-khasi-hills"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := newTestServer(tt.response)

			digitalhumani := DigitialHumani{
				url:          srv.URL,
				HTTPClient:   srv.Client(),
				apiKey:       "testkey",
				enterpriseID: "testID",
			}

			project, err := digitalhumani.GetProject(tt.args.ctx, tt.args.projectID)

			testutil.Equals(t, tt.want, project)

			if (err != nil) != tt.wantErr {
				t.Errorf("DigitialHumani.GetProject() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDigitialHumani_GetAllProjects(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     []Project
		wantErr  bool
	}{
		{
			name:     "Successful GetAllProjects 200",
			args:     args{context.TODO()},
			response: successfulGetAllProjects,
			want: []Project{
				{ID: "91111111", Name: "California OneTreePlanted", ReforestationProjectDescriptionEn: "Reforestation project in California, United States", ReforestationProjectStateEn: "California", ReforestationProjectCountryEn: "United States", ReforestationProjectWebsiteEn: "https://onetreeplanted.org/united-states/products/california-forests", ReforestationCompanyNameEn: "OneTreePlanted"},
				{ID: "96666666", Name: "Khasi Hills in India, WeForest", ReforestationProjectDescriptionEn: "Reforestation project in the Khasi Hills in India", ReforestationProjectStateEn: "Khasi Hills", ReforestationProjectCountryEn: "India", ReforestationProjectWebsiteEn: "https://www.weforest.org/project/india-khasi-hills", ReforestationCompanyNameEn: "WeForest"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := newTestServer(tt.response)

			digitalhumani := DigitialHumani{
				url:          srv.URL,
				HTTPClient:   srv.Client(),
				apiKey:       "testkey",
				enterpriseID: "testID",
			}

			projects, err := digitalhumani.GetAllProjects(tt.args.ctx)

			testutil.Equals(t, tt.want, projects)

			if (err != nil) != tt.wantErr {
				t.Errorf("DigitialHumani.GetAllProjects() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
