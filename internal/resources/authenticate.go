package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type AuthenticateResource struct{}

func (ar AuthenticateResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{code}", ar.authenticate)
	return r
}

func (ar AuthenticateResource) authenticate(w http.ResponseWriter, r *http.Request) {}
