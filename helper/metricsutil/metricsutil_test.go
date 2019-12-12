package metricsutil

import (
	"github.com/hashicorp/vault/sdk/logical"
	"testing"
)

func TestFormatFromRequest(t *testing.T) {
	testCases := []struct {
		original *logical.Request
		expected string
	}{
		{
			original: &logical.Request{Headers: map[string][]string{
				"Accept": {
					"application/vnd.google.protobuf",
					"schema=\"prometheus/telemetry\"",
				},
			}},
			expected: "prometheus",
		},
		{
			original: &logical.Request{Headers: map[string][]string{
				"Accept": {
					"schema=\"prometheus\"",
				},
			}},
			expected: "",
		},
		{
			original: &logical.Request{Headers: map[string][]string{
				"Accept": {
					"application/openmetrics-text",
				},
			}},
			expected: "prometheus",
		},
	}

	for _, tCase := range testCases {
		if metricsType := FormatFromRequest(tCase.original); metricsType != tCase.expected {
			t.Fatalf("expected %s but got %s", tCase.expected, metricsType)
		}
	}
}