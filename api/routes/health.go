package routes

import (
	"net/http"
	"strings"

	"github.com/secnex/api-server/api/res"
)

func (router *Router) RouteHealth(w http.ResponseWriter, r *http.Request) {
	methods := []string{http.MethodGet}
	switch r.Method {
	case http.MethodGet:
		router.GETHealth(w, r)
	default:
		w.Header().Set("Allow", strings.Join(methods, ", "))
		w.WriteHeader(http.StatusMethodNotAllowed)
		result := res.ResultError{
			Code:    http.StatusMethodNotAllowed,
			Message: http.StatusText(http.StatusMethodNotAllowed),
			Error:   "Method not allowed",
		}
		w.Write([]byte(result.String()))
	}
}

// New Route: GET /health
func (router *Router) GETHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := res.ResultHealth{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Status:  "OK",
	}
	w.Write([]byte(result.String()))
}
