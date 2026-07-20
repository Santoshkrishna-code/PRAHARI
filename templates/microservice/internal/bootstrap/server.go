package bootstrap

import (
	"fmt"
	"net/http"
	"time"

	prahariLogger "prahari/shared/logger"
)

// StartServer spins up the production-grade HTTP listener.
func StartServer(port int, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		// Start listener
		_ = srv.ListenAndServe()
	}()

	return srv
}
