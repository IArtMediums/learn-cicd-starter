package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		authHeader http.Header
		wantOut    string
		wantErr    bool
	}{
		"Valid test 1":                         {authHeader: getAuthHeader("ApiKey 12345"), wantOut: "12345", wantErr: false},
		"Valid test 2":                         {authHeader: getAuthHeader("ApiKey 09187093124"), wantOut: "09187093124", wantErr: false},
		"No authorization header":              {authHeader: getAuthHeader(""), wantOut: "", wantErr: true},
		"Invalid formating (many spaces)":      {authHeader: getAuthHeader("ApiKey 24421 1234 5643"), wantOut: "", wantErr: true},
		"Invalid formating (no ApiKey prefix)": {authHeader: getAuthHeader("1294807"), wantOut: "", wantErr: true},
		"Invalid formating (no spaces)":        {authHeader: getAuthHeader("ApiKey13245"), wantOut: "", wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.authHeader)

			if (err != nil) && !tc.wantErr {
				t.Fatalf("want error= %v, err= %v", tc.wantErr, err)
			}

			diff := cmp.Diff(tc.wantOut, got)

			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func getAuthHeader(value string) http.Header {
	h := http.Header{}
	if value == "" {
		return h
	}

	h.Add("Authorization", value)
	return h
}
