package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ExtResource struct{}

func (er ExtResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/installations/{install_id}/repos/{repo_id}/contents/{path}", er.repoFileContent)
	r.Post("/ext/integrations/{registry_type}/credentials/validate", er.validateRegistryCredentials)
	return r
}

func (er ExtResource) repoFileContent(w http.ResponseWriter, r *http.Request)             {}
func (er ExtResource) validateRegistryCredentials(w http.ResponseWriter, r *http.Request) {}
