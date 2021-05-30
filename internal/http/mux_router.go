package router

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct {
	server     *http.Server
	dispatcher *mux.Router
}

func NewMuxRouter(server *http.Server) Router {
	return &muxRouter{
		server:     server,
		dispatcher: mux.NewRouter(),
	}
}

func (m *muxRouter) GET(uri string, handler http.Handler) {
	m.dispatcher.Handle(uri, handler).Methods(http.MethodGet)
}

func (m *muxRouter) POST(uri string, handler http.Handler) {
	m.dispatcher.Handle(uri, handler).Methods(http.MethodPost)
}

func (m *muxRouter) SERVE() error {
	m.server.Handler = m.dispatcher
	err := m.server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (m *muxRouter) SHUTDOWN(ctx context.Context) error {
	err := m.server.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (m *muxRouter) PREFIX(path string) {
	m.dispatcher = m.dispatcher.PathPrefix(path).Subrouter().StrictSlash(true)
}
