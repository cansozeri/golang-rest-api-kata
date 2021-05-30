package http

import (
	"github.com/urfave/negroni"
	router "golang-rest-api-kata/internal/http"
)

func MapRecordRoutes(r router.Router, h *RecordHandlers, n *negroni.Negroni) {
	r.POST("/records/search", n.With(
		negroni.Wrap(h.SearchRecords()),
	))
}
