# hostile

### Usage
```go
import "github.com/nilslice/hostile"
```

hostile is a package for multi-host routing. When you need to respond to
requests made with different host names use hostile. Each host has its
own mux that registers unique routes.

```go
h := NewHostHandler()

host1 := h.AddHost("host1.com")
host2 := h.AddHost("host2.com")

host1.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("host1 has own '/' route and handler"))
})

host2.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
    res.Write([]byte("host2 has own '/' route and handler"))
})

err := http.ListenAndServe(":8000", h)
if err != nil {
    fmt.Println(err)
}
```

#### type HostHandler struct

    HostHandler implements http.Handler interface and contains a map of host
    names whose values are a *http.ServeMux to register per-host routes.

#### func NewHostHandler() *HostHandler
    NewHostHandler allocates and initializes a map to hold eligible hosts
    and the ServeMux to register per-host routes. Returns a pointer to
    HostHandler which implements http.Handler and is expected by
    http.ListenAndServe or http.Handle and the like.

#### func (h *HostHandler) AddHost(host string) *http.ServeMux
    AddHost creates a new ServeMux per added host and sets in in the
    eligibleHosts map within the HostHandler. Each host has its own mux so
    you can assign routes to individial hosts.
    
    Example:
    ```go
	newHost := h.AddHost("newhost.com")
	newHost.HandleFunc("/route", handler)
    ```

#### func (h HostHandler) IsEligible(host string) (string, bool)
    IsEligible is a helper function which tests host eligibility and returns
    the host provided and a bool indicating status.

#### func (h *HostHandler) ServeHTTP(res http.ResponseWriter, req *http.Request)
    ServeHTTP implements http.Handler for HostHandler. First it checks if we
    have a request from an eligible host, then passes on the request and
    response writer to the handler of a the host if it is found.


