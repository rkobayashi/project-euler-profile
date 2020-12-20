package profsvg

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetXMLResponseCode(t *testing.T) {
	testcases := []struct {
		statusCode int
		wantError  bool
	}{
		{http.StatusOK, false},
		{http.StatusNotFound, true},
	}

	for _, tc := range testcases {
		_, err := getXML("", func(url string) (*http.Response, error) {
			body := strings.NewReader("")
			return &http.Response{
				StatusCode: tc.statusCode,
				Body:       ioutil.NopCloser(body),
			}, nil
		})

		if (err != nil) != tc.wantError {
			t.Errorf("code =%v, err = %v, wantError = %v", tc.statusCode, err, tc.wantError)
		}
	}
}
