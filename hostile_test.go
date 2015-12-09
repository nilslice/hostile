package hostile

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	mockHost = "myhost.dev"
	badHost  = "nohost.dev"
)

func TestNewHostHandler(t *testing.T) {
	h := NewHostHandler()
	if len(h.eligibleHosts) > 0 {
		t.Errorf("incorrect initialization for NewHostHandler (%d eligibleHosts)", len(h.eligibleHosts))
	}
}

func TestHostEligibility(t *testing.T) {
	h := NewHostHandler()
	_ = h.AddHost(mockHost)

	if _, ok := h.IsEligible(mockHost); !ok {
		t.Errorf("Eligible host (%s) should be ok, but failed: %s", mockHost, ok)
	}

	if _, ok := h.IsEligible(badHost); ok {
		t.Errorf("Bad host (%s) should be !ok, but passed: %s", badHost, ok)
	}
}

func TestAddHost(t *testing.T) {
	h := NewHostHandler()
	_ = h.AddHost(mockHost)
	numHosts := len(h.eligibleHosts)

	if numHosts != 1 {
		t.Errorf("Eligible hosts has unexpected length:", numHosts)
	}
}

func TestHTTPHandlerImplementation(t *testing.T) {
	h := NewHostHandler()
	server := httptest.NewServer(h)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusForbidden {
		t.Errorf("Unexpected http statusCode. expects 403, got:", resp.StatusCode)
	}

}
