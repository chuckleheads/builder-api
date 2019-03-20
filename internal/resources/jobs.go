package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type JobsResource struct{}

func (jr JobsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", jr.getJob)
		r.Get("/log", jr.getJobLog)
	})
	r.Route("/group/{id}", func(r chi.Router) {
		r.Post("/promote/{channel}", jr.promoteJobGroup)
		r.Post("/demote/{channel}", jr.demoteJobGroup)
		r.Post("/cancel", jr.cancelJobGroup)
	})
	return r
}

func (jr JobsResource) getJob(w http.ResponseWriter, r *http.Request)          {}
func (jr JobsResource) getJobLog(w http.ResponseWriter, r *http.Request)       {}
func (jr JobsResource) promoteJobGroup(w http.ResponseWriter, r *http.Request) {}
func (jr JobsResource) demoteJobGroup(w http.ResponseWriter, r *http.Request)  {}
func (jr JobsResource) cancelJobGroup(w http.ResponseWriter, r *http.Request)  {}
