package bootstrap

import (
	"fmt"
	"net/http"
	"time"
)

// StartServer spins up the production HTTP listener.
func StartServer(port int, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		_ = srv.ListenAndServe()
	}()

	return srv
}
