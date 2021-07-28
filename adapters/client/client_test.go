package client

import (
	"testing"
)

func TestClientSettings(t *testing.T) {
	tt := []struct {
		test string
		err  error
	}{
		{"Without protocol", NewClient("", "8090").Run()},
		{"Without address", NewClient("tcp", "").Run()},
	}

	for _, tc := range tt {
		t.Run(tc.test, func(t *testing.T) {
			if tc.err == nil {
				t.Error("it should be an invalid client")
			}
		})
	}
}
