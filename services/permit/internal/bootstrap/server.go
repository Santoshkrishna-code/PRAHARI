package bootstrap

import (
	"fmt"
	"net/http"
	"time"
)

// StartServer binds HTTP listener configurations.
func StartServer(port int, handler http.Handler) *http.Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		_ = srv.ListenAndServe()
	}()

	return srv
}
