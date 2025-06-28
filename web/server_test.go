package web

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

// helper to create a test server with the same routes as ArtServer
func newTestServer() *httptest.Server {
	// Ensure we run from repository root so relative template paths work.
	_ = os.Chdir("..")

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/decoder", postDecoder)
	mux.HandleFunc("/decoder/", postDecoder)
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("styles"))))
	return httptest.NewServer(mux)
}

func TestGetRootReturnsOK(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatalf("GET / failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}
}

func TestPostDecoderMalformedReturnsBadRequest(t *testing.T) {
	ts := newTestServer()
	defer ts.Close()

	invalid := "[1 foo" // missing closing bracket
	resp, err := http.Post(ts.URL+"/decoder", "text/plain", strings.NewReader(invalid))
	if err != nil {
		t.Fatalf("POST /decoder failed: %v", err)
	}
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", resp.StatusCode)
	}
}
