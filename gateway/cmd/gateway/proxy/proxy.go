package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

type proxy struct {
	remote       *url.URL
	reverseProxy *httputil.ReverseProxy
}

func New(urlString string) *proxy {
	remote, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(remote)

	return &proxy{
		remote:       remote,
		reverseProxy: reverseProxy,
	}
}

func (p *proxy) HandleRequest(w http.ResponseWriter, r *http.Request) {
	r.URL.Host = p.remote.Host
	r.URL.Scheme = p.remote.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = p.remote.Host

	r.URL.Path = mux.Vars(r)["rest"]
	p.reverseProxy.ServeHTTP(w, r)
}
