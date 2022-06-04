package helloworld

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloGet(t *testing.T) {
	payload := strings.NewReader("")
	req := httptest.NewRequest("GET", "/", payload)

	rr := httptest.NewRecorder()
	HelloGet(rr, req)

	if got, want := rr.Body.String(), "Hello, World!"; got != want {
		t.Errorf("HelloWord = %q, want %q", got, want)
	}
}
