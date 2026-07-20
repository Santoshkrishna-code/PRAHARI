package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"prahari/edge/gateway/internal/domain/route"
)

// Forward proxies incoming requests to target upstreams, rewriting headers.
func Forward(w http.ResponseWriter, r *http.Request, rt *route.Route) error {
	targetURL, err := url.Parse(rt.Upstream)
	if err != nil {
		return fmt.Errorf("failed to parse upstream URL: %w", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Ingress path rewrite logic
	if rt.StripPrefix {
		patternPrefix := strings.TrimSuffix(rt.Path, "*")
		r.URL.Path = "/" + strings.TrimPrefix(r.URL.Path, patternPrefix)
	}

	// Update host headers to match upstream server expectation
	r.Host = targetURL.Host
	r.URL.Host = targetURL.Host
	r.URL.Scheme = targetURL.Scheme

	proxy.ServeHTTP(w, r)
	return nil
}
