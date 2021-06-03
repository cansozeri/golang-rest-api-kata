package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type chiRouter struct {
	server     *http.Server
	dispatcher *chi.Mux
}

func NewChiRouter(server *http.Server) Router {
	return &chiRouter{
		server:     server,
		dispatcher: chi.NewRouter(),
	}
}

func (c *chiRouter) GET(uri string, handler http.Handler) {
	c.dispatcher.Method("GET", uri, handler)
}

func (c *chiRouter) POST(uri string, handler http.Handler) {
	c.dispatcher.Method("POST", uri, handler)
}

func (c *chiRouter) SERVE() error {
	c.server.Handler = c.dispatcher
	err := c.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (c *chiRouter) SHUTDOWN(ctx context.Context) error {
	err := c.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *chiRouter) PREFIX(path string) {
	c.dispatcher.Mount(path, c.dispatcher)
}
