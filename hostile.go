package hostile

import (
	"fmt"
	"net/http"
)

// HostHandler implements http.Handler interface and contains a map of host names
// whose values are a *http.ServeMux to register per-host routes.
type HostHandler struct {
	eligibleHosts map[string]*http.ServeMux
}

// NewHostHandler allocates and initializes a map to hold eligible hosts and the
// ServeMux to register per-host routes. Returns a pointer to HostHandler which
// implements http.Handler and is expected by http.ListenAndServe or http.Handle
// and the like.
func NewHostHandler() *HostHandler {
	h := &HostHandler{
		eligibleHosts: make(map[string]*http.ServeMux),
	}

	return h
}

// IsEligible is a helper function which tests host eligibility and returns the
// host provided and a bool indicating status.
func (h HostHandler) IsEligible(host string) (string, bool) {
	if _, ok := h.eligibleHosts[host]; ok {
		return host, true
	}
	return host, false
}

// AddHost creates a new ServeMux per added host and sets in in the eligibleHosts
// map within the HostHandler. Each host has its own mux so you can assign routes
// to individial hosts.
// Example:
// 	newHost := h.AddHost("newhost.com")
// 	newHost.HandleFunc("/route", handler)
func (h *HostHandler) AddHost(host string) *http.ServeMux {
	mux := http.NewServeMux()
	h.eligibleHosts[host] = mux
	return mux
}

// ServeHTTP implements http.Handler for HostHandler. First it checks if we
// have a request from an eligible host, then passes on the request and response
// writer to the handler of a the host if it is found.
func (h *HostHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if host, ok := h.IsEligible(req.Host); ok {
		fmt.Println("request to", host)
		h.eligibleHosts[host].ServeHTTP(res, req)
	} else {
		http.Error(res, "Forbidden", http.StatusForbidden)
	}
}
