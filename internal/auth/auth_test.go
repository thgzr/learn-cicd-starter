package auth

import (
	"net/http"
	"reflect"
	"testing"
)

// GetAPIKey -
func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "ApiKey sometoken")
	got, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "sometoken"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	req, _ = http.NewRequest("GET", "/", nil)
	// req.Header.Del("Authorization")

	got, err = GetAPIKey(req.Header)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if got != "" {
		t.Fatalf("expected empty token, got %q", got)
	}
	if err.Error() != "no authorization header included" {
		t.Fatalf("unexpected error: %v", err)
	}
}
