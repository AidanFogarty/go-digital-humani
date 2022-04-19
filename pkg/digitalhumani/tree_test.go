package digitalhumani

import (
	"context"
	"testing"

	"github.com/AidanFogarty/go-digital-humani/pkg/testutil"
)

const (
	successfulGetTreeCount = `
	{
		"enterpriseId": "testID",
		"user": "test_user_1",
		"count": 4
	}
	`
	sucessfulPlantTree = `
	{
		"uuid": "eef9f369-9ae0-45b8-ab07-10650f53a71e",
		"created": "2019-05-17T00:36:25.797Z",
		"treeCount": 1,
		"enterpriseId": "testID",
		"projectId": "93333333",
		"user": "email@test.com"
	}
	`
	successfulGetTree = `
	{
		"projectId": "81818182",
		"created": "2020-04-11T18:01:40.441Z",
		"uuid": "bcd35c97-d66c-412e-89ae-ecbac0f629ac",
		"user": "1",
		"treeCount": 1,
		"enterpriseId": "33333333"
	}
	`
)

func TestDigitialHumani_GetTreeCount(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     *UserTreeCount
		wantErr  bool
	}{
		{
			name:     "Successful GetTreeCount 200",
			args:     args{context.TODO(), "test_user_1"},
			response: successfulGetTreeCount,
			want:     &UserTreeCount{EnterpriseID: "testID", User: "test_user_1", Count: 4},
			wantErr:  false,
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

			treeCount, err := digitalhumani.GetTreeCount(tt.args.ctx, tt.args.userID)

			testutil.Equals(t, tt.want, treeCount)

			if (err != nil) != tt.wantErr {
				t.Errorf("DigitialHumani.GetTreeCount() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDigitialHumani_PlantTree(t *testing.T) {
	type args struct {
		ctx       context.Context
		projectID string
		userID    string
		treeCount int
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     *PlantTreeRequest
		wantErr  bool
	}{
		{
			name:     "Successful PlantTree 200",
			args:     args{context.TODO(), "93333333", "test_user_1", 1},
			response: sucessfulPlantTree,
			want:     &PlantTreeRequest{UUID: "eef9f369-9ae0-45b8-ab07-10650f53a71e", Created: "2019-05-17T00:36:25.797Z", TreeCount: 1, EnterpriseID: "testID", ProjectID: "93333333", User: "email@test.com"},
			wantErr:  false,
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

			plantTree, err := digitalhumani.PlantTree(tt.args.ctx, tt.args.projectID, tt.args.userID, tt.args.treeCount)

			testutil.Equals(t, tt.want, plantTree)

			if (err != nil) != tt.wantErr {
				t.Errorf("DigitialHumani.PlantTree() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDigitialHumani_GetTree(t *testing.T) {
	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     *PlantTreeRequest
		wantErr  bool
	}{
		{
			name:     "Successful GetTree 200",
			args:     args{context.TODO(), "bcd35c97-d66c-412e-89ae-ecbac0f629ac"},
			response: successfulGetTree,
			want:     &PlantTreeRequest{UUID: "bcd35c97-d66c-412e-89ae-ecbac0f629ac", Created: "2020-04-11T18:01:40.441Z", TreeCount: 1, EnterpriseID: "33333333", ProjectID: "81818182", User: "1"},
			wantErr:  false,
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

			getTree, err := digitalhumani.GetTree(tt.args.ctx, tt.args.uuid)

			testutil.Equals(t, tt.want, getTree)

			if (err != nil) != tt.wantErr {
				t.Errorf("DigitialHumani.GetTree() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
