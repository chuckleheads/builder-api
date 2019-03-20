package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type ChannelResource struct{}

func (cr ChannelResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{origin}", func(r chi.Router) {
		r.Get("/", cr.getChannels)
		r.Route("/{channel}", func(r chi.Router) {
			r.Post("/", cr.createChannel)
			r.Delete("/", cr.deleteChannel)
			r.Route("/pkgs", func(r chi.Router) {
				r.Get("/", cr.getPackagesForOriginChannel)
				r.Route("/{pkg}", func(r chi.Router) {
					r.Get("/", cr.getPackagesForOriginChannelPackage)
					r.Get("/latest", cr.getLatestPackageForOriginChannelPackage)
					r.Route("/{version}", func(r chi.Router) {
						r.Get("/", cr.getPackagesForOriginChannelPackageVersion)
						r.Get("/latest", cr.getLatestPackageForOriginChannelPackageVersion)
						r.Route("/{release}", func(r chi.Router) {
							r.Get("/", cr.getPackageFullyQualified)
							r.Put("/promote", cr.promotePackage)
							r.Put("/demote", cr.demotePackage)
						})
					})
				})
			})
		})

	})

	return r
}

func (cr ChannelResource) getChannels(w http.ResponseWriter, r *http.Request)                        {}
func (cr ChannelResource) createChannel(w http.ResponseWriter, r *http.Request)                      {}
func (cr ChannelResource) deleteChannel(w http.ResponseWriter, r *http.Request)                      {}
func (cr ChannelResource) getPackagesForOriginChannel(w http.ResponseWriter, r *http.Request)        {}
func (cr ChannelResource) getPackagesForOriginChannelPackage(w http.ResponseWriter, r *http.Request) {}
func (cr ChannelResource) getLatestPackageForOriginChannelPackage(w http.ResponseWriter, r *http.Request) {
}
func (cr ChannelResource) getPackagesForOriginChannelPackageVersion(w http.ResponseWriter, r *http.Request) {
}
func (cr ChannelResource) getLatestPackageForOriginChannelPackageVersion(w http.ResponseWriter, r *http.Request) {
}
func (cr ChannelResource) getPackageFullyQualified(w http.ResponseWriter, r *http.Request) {}
func (cr ChannelResource) promotePackage(w http.ResponseWriter, r *http.Request)           {}
func (cr ChannelResource) demotePackage(w http.ResponseWriter, r *http.Request)            {}
