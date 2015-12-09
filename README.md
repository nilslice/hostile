# hostile
    import "github.com/nilslice/hostile"

    hostile is a package for multi-host routing. When you need to respond to
    requests made with different host names use hostile. Each host has its
    own mux that registers unique routes.

```
type HostHandler struct {
    // contains filtered or unexported fields
}
```
    HostHandler implements http.Handler interface and contains a map of host
    names whose values are a *http.ServeMux to register per-host routes.

### func NewHostHandler() *HostHandler
    NewHostHandler allocates and initializes a map to hold eligible hosts
    and the ServeMux to register per-host routes. Returns a pointer to
    HostHandler which implements http.Handler and is expected by
    http.ListenAndServe or http.Handle and the like.

### func (h *HostHandler) AddHost(host string) *http.ServeMux
    Example:

	newHost := h.AddHost("newhost.com")
	newHost.HandleFunc("/route", handler)

### func (h HostHandler) IsEligible(host string) (string, bool)
    IsEligible is a helper function which tests host eligibility and returns
    the host provided and a bool indicating status.

### func (h *HostHandler) ServeHTTP(res http.ResponseWriter, req *http.Request)
    ServeHTTP implements http.Handler for HostHandler. First it checks if we
    have a request from an eligible host, then passes on the request and
    response writer to the handler of a the host if it is found.


