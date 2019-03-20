package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ProfileResource struct{}

func (pr ProfileResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", pr.getAccount)
	r.Patch("/", pr.updateAccount)
	r.Route("/access-tokens", func(r chi.Router) {
		r.Get("/", pr.getAccessTokens)
		r.Post("/", pr.generateAccessToken)
		r.Delete("/{id}", pr.revokeAccessToken)
	})
	return r
}
func (pr ProfileResource) getAccount(w http.ResponseWriter, r *http.Request)          {}
func (pr ProfileResource) updateAccount(w http.ResponseWriter, r *http.Request)       {}
func (pr ProfileResource) getAccessTokens(w http.ResponseWriter, r *http.Request)     {}
func (pr ProfileResource) generateAccessToken(w http.ResponseWriter, r *http.Request) {}
func (pr ProfileResource) revokeAccessToken(w http.ResponseWriter, r *http.Request)   {}
