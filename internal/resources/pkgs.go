package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type PkgsResource struct{}

func (pr PkgsResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/{origin}", pr.getPackagesForOrigin)
	r.Get("/search/{query}", pr.searchPackges)
	r.Route("/schedule", func(r chi.Router) {
		r.Get("/{groupId}", pr.getSchedule)
		r.Get("/{origin}/status", pr.getOriginScheduleStatus)
		r.Post("/{origin}/{pkg}", pr.scheduleJobGroup)
	})
	r.Route("/{origin}/{pkg}", func(r chi.Router) {
		r.Get("/latest", pr.getLatestPackageForOriginPackage)
		r.Get("/versions", pr.listPackageVersions)
		r.Route("/{version}", func(r chi.Router) {
			r.Get("/", pr.getPackagesForOriginPackageVersion)
			r.Get("/latest", pr.getLatestPackageForOriginPackageVersion)
			r.Route("/{release}", func(r chi.Router) {
				r.Post("/", pr.uploadPackage)
				r.Get("/", pr.getPackage)
				r.Get("/download", pr.downloadPackage)
				r.Get("/channels", pr.getPackageChannels)
				r.Patch("/{visibility}", pr.packageVisibilityToggle)
			})
		})
	})
	return r
}

func (pr PkgsResource) getPackagesForOrigin(w http.ResponseWriter, r *http.Request)               {}
func (pr PkgsResource) searchPackges(w http.ResponseWriter, r *http.Request)                      {}
func (pr PkgsResource) getSchedule(w http.ResponseWriter, r *http.Request)                        {}
func (pr PkgsResource) getOriginScheduleStatus(w http.ResponseWriter, r *http.Request)            {}
func (pr PkgsResource) scheduleJobGroup(w http.ResponseWriter, r *http.Request)                   {}
func (pr PkgsResource) getLatestPackageForOriginPackage(w http.ResponseWriter, r *http.Request)   {}
func (pr PkgsResource) listPackageVersions(w http.ResponseWriter, r *http.Request)                {}
func (pr PkgsResource) getPackagesForOriginPackageVersion(w http.ResponseWriter, r *http.Request) {}
func (pr PkgsResource) getLatestPackageForOriginPackageVersion(w http.ResponseWriter, r *http.Request) {
}
func (pr PkgsResource) uploadPackage(w http.ResponseWriter, r *http.Request)           {}
func (pr PkgsResource) getPackage(w http.ResponseWriter, r *http.Request)              {}
func (pr PkgsResource) downloadPackage(w http.ResponseWriter, r *http.Request)         {}
func (pr PkgsResource) getPackageChannels(w http.ResponseWriter, r *http.Request)      {}
func (pr PkgsResource) packageVisibilityToggle(w http.ResponseWriter, r *http.Request) {}
