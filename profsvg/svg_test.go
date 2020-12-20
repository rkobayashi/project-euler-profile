package profsvg

import (
	"net/url"
	"testing"
)

func TestParseQuery(t *testing.T) {
	testcases := []struct {
		input, expected string
		wantError       bool
	}{
		{"username=rkobayashi", "rkobayashi", false},
		{"", "", true},
	}

	for _, tc := range testcases {
		query, _ := url.ParseQuery(tc.input)
		actual, err := parseUserNameFromQuery(query)

		if (err != nil) != tc.wantError {
			t.Errorf("err = %v, wantError = %v", err, tc.wantError)
		}

		if actual != tc.expected {
			t.Errorf("expected = %s, but actual = %s, query=%v",
				tc.expected, actual, query)
		}
	}
}
