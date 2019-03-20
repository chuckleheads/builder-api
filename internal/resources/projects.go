package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ProjectsResource struct{}

func (pr ProjectsResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/projects", func(r chi.Router) {
		r.Post("/", pr.createproject)
		r.Route("/{origin}", func(r chi.Router) {
			r.Get("/", pr.getProjects)
			r.Route("/{name}", func(r chi.Router) {
				r.Get("/", pr.getProject)
				r.Put("/", pr.updateProject)
				r.Delete("/", pr.deleteProject)
				r.Get("/jobs", pr.getJobs)
				r.Route("/integrations/{integration}/default", func(r chi.Router) {
					r.Get("/", pr.getIntegration)
					r.Put("/", pr.createIntegration)
					r.Delete("/", pr.deleteIntegration)
				})
				r.Patch("/{visibility}", pr.togglePrivacy)
			})
		})
	})
	return r
}
func (pr ProjectsResource) createproject(w http.ResponseWriter, r *http.Request)     {}
func (pr ProjectsResource) getProjects(w http.ResponseWriter, r *http.Request)       {}
func (pr ProjectsResource) getProject(w http.ResponseWriter, r *http.Request)        {}
func (pr ProjectsResource) updateProject(w http.ResponseWriter, r *http.Request)     {}
func (pr ProjectsResource) deleteProject(w http.ResponseWriter, r *http.Request)     {}
func (pr ProjectsResource) getJobs(w http.ResponseWriter, r *http.Request)           {}
func (pr ProjectsResource) getIntegration(w http.ResponseWriter, r *http.Request)    {}
func (pr ProjectsResource) createIntegration(w http.ResponseWriter, r *http.Request) {}
func (pr ProjectsResource) deleteIntegration(w http.ResponseWriter, r *http.Request) {}
func (pr ProjectsResource) togglePrivacy(w http.ResponseWriter, r *http.Request)     {}
