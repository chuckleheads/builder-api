package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type RdepsResource struct{}

func (rr RdepsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{origin}/{name}", func(r chi.Router) {
		r.Get("/", rr.getRdeps)
		r.Get("/group", rr.getRdepsGroup)
	})
	return r
}

func (rr RdepsResource) getRdeps(w http.ResponseWriter, r *http.Request)      {}
func (rr RdepsResource) getRdepsGroup(w http.ResponseWriter, r *http.Request) {}
