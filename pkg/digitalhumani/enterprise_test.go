package digitalhumani

import (
	"context"
	"testing"

	"github.com/AidanFogarty/go-digital-humani/pkg/testutil"
)

const (
	successfulGetEnterprise = `
	{
		"created": "2018-12-12T09:07:00.725Z",
		"updated": "2019-03-30T15:03:42.095Z",
		"id": "11111111",
		"name": "Test - Cable Company ABC",
		"contact": {
		  "name": "Jane Doe"
		}
	}
	`
	successfulGetEnterpriseTreeCount = `
	{
		"count": 57
	}
	`
	successfulGetEnterpriseMonthTreeCount = `
	{
		"count": 17
	}
	`
)

func TestDigitialHumani_GetEnterprise(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     *Enterprise
		wantErr  bool
	}{
		{
			name:     "Successful GetEnterprise 200",
			args:     args{context.TODO()},
			response: successfulGetEnterprise,
			want:     &Enterprise{Created: "2018-12-12T09:07:00.725Z", Updated: "2019-03-30T15:03:42.095Z", ID: "11111111", Name: "Test - Cable Company ABC", Contact: Contact{Name: "Jane Doe"}},
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

			enterprise, err := digitalhumani.GetEnterprise(tt.args.ctx)

			testutil.Equals(t, tt.want, enterprise)

			if (err != nil) != tt.wantErr {
				t.Errorf("DigitialHumani.GetEnterprise() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDigitialHumani_GetEnterpriseTreeCount(t *testing.T) {
	type args struct {
		ctx       context.Context
		dateRange *dateRange
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     *EnterpriseTreeCount
		wantErr  bool
	}{
		{
			name:     "Successful GetEnterpriseTreeCount 200",
			args:     args{context.TODO(), &dateRange{StartDate: "2022-01-01", EndDate: "2022-12-31"}},
			response: successfulGetEnterpriseTreeCount,
			want:     &EnterpriseTreeCount{Count: 57},
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

			enterprise, err := digitalhumani.GetEnterpriseTreeCount(tt.args.ctx, tt.args.dateRange)

			testutil.Equals(t, tt.want, enterprise)

			if (err != nil) != tt.wantErr {
				t.Errorf("DigitialHumani.GetEnterpriseTreeCount() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDigitialHumani_GetEnterpriseMonthTreeCount(t *testing.T) {
	type args struct {
		ctx   context.Context
		month string
	}
	tests := []struct {
		name     string
		args     args
		response string
		want     *EnterpriseTreeCount
		wantErr  bool
	}{
		{
			name:     "Successful GetEnterpriseMonthTreeCount 200",
			args:     args{context.TODO(), "2020-02"},
			response: successfulGetEnterpriseMonthTreeCount,
			want:     &EnterpriseTreeCount{Count: 17},
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

			enterprise, err := digitalhumani.GetEnterpriseMonthTreeCount(tt.args.ctx, tt.args.month)

			testutil.Equals(t, tt.want, enterprise)

			if (err != nil) != tt.wantErr {
				t.Errorf("DigitialHumani.GetEnterpriseMonthTreeCount() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
