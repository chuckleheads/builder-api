package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type NotifyResource struct{}

func (nr NotifyResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", nr.notify)
	return r
}

func (nr NotifyResource) notify(w http.ResponseWriter, r *http.Request) {}
