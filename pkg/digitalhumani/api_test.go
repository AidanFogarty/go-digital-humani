package digitalhumani

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AidanFogarty/go-digital-humani/pkg/testutil"
)

func TestNew(t *testing.T) {
	type args struct {
		url          string
		apiKey       string
		enterpriseID string
	}
	tests := []struct {
		name string
		args args
		want *DigitialHumani
	}{
		{
			name: "Successful Test New Sandbox",
			args: args{"testkey", "testID", "sandbox"},
			want: &DigitialHumani{"https://api.sandbox.digitalhumani.com", "testkey", "testID", http.DefaultClient},
		},
		{
			name: "Successful Test New Prod",
			args: args{"testkey", "testID", "production"},
			want: &DigitialHumani{"https://api.digitalhumani.com", "testkey", "testID", http.DefaultClient},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutil.Equals(t, tt.want, New(tt.args.url, tt.args.apiKey, tt.args.enterpriseID))
		})
	}
}

func newTestServer(response string) *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(response))
	})
	return httptest.NewServer(handler)
}
