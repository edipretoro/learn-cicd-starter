package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey testApi")
	fmt.Printf("%v", headers)
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("problem when getting the ApiKey: %v", err)
	}
	want := "testApi"
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
