package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

type proxy struct {
	remote       *url.URL
	reverseProxy *httputil.ReverseProxy
	path         string
}

func New(urlString, path string) *proxy {
	remote, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(remote)

	return &proxy{
		remote:       remote,
		reverseProxy: reverseProxy,
		path:         path,
	}
}

func (p *proxy) HandleRequest(w http.ResponseWriter, r *http.Request) {
	r.URL.Host = p.remote.Host
	r.URL.Scheme = p.remote.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = p.remote.Host

	rest := mux.Vars(r)["rest"]
	path := fmt.Sprintf("/%s", p.path)
	if rest != "" {
		path = fmt.Sprintf("/%s/%s", p.path, rest)
	}
	r.URL.Path = path

	p.reverseProxy.ServeHTTP(w, r)
}
