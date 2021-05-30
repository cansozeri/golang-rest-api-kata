package router

import (
	"context"
	"net/http"
)

type Router interface {
	GET(uri string, handler http.Handler)
	POST(uri string, handler http.Handler)
	SERVE() error
	PREFIX(path string)
	SHUTDOWN(ctx context.Context) error
}
