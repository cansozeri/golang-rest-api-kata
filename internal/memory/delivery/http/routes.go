package http

import (
	"github.com/urfave/negroni"
	router "golang-rest-api-kata/internal/http"
)

func MapMemoryRoutes(r router.Router, h *MemoryHandlers, n *negroni.Negroni) {
	r.GET("/in-memory", n.With(
		negroni.Wrap(h.GetInMemory()),
	))

	r.POST("/in-memory", n.With(
		negroni.Wrap(h.CreateInMemory()),
	))
}
