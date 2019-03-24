package resources

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/chuckleheads/builder-api/pkg/models"
)

type OriginResource struct {
	db *sql.DB
}

func NewOriginResource(db *sql.DB) *OriginResource {
	return &OriginResource{db}
}

func (or OriginResource) Routes() chi.Router {
	r := chi.NewRouter()

	// 	r.Get("/{origin}/pkgs", or.listUniquePackages)
	r.Post("/", or.createOrigin)
	r.Route("/{origin}", func(r chi.Router) {
		r.Get("/", or.getOrigin)
		r.Put("/", or.updateOrigin)
		r.Delete("/", or.deleteOrigin)
		r.Route("/users", func(r chi.Router) {
			r.Get("/", or.getOriginMembers)
			r.Delete("/{user}", or.deleteOriginMember)
			r.Post("/{username}/invitations", or.inviteToOrigin)
		})
		r.Route("/invitations", func(r chi.Router) {
			r.Get("/", or.listOriginInvitations)
			r.Put("/{invitationId}", or.acceptInvitation)
			r.Delete("/{invitationId}", or.rescindInvitation)
			r.Put("/{invitationId}/ignore", or.ignoreInvitation)
		})
		r.Route("/keys", func(r chi.Router) {
			r.Post("/", or.createKeys)
			r.Get("/", or.listOriginKeys)
			r.Get("/latest", or.downloadLatestOriginKey)
			r.Post("/{revision}", or.uploadOriginKey)
			r.Get("/{revision}", or.downloadOriginKey)
		})
		r.Route("/secret", func(r chi.Router) {
			r.Get("/", or.listOriginSecrets)
			r.Post("/", or.createOriginSecret)
			r.Delete("/{secret}", or.deleteOriginSecret)
		})
		r.Route("/integrations", func(r chi.Router) {
			r.Get("/", or.fetchOriginIntegrations)
			r.Route("/{integration}", func(r chi.Router) {
				r.Get("/names", or.fetchOriginIntegrationNames)
				r.Get("/{name}", or.getOriginIntegration)
				r.Delete("/{name}", or.deleteOriginIntegration)
				r.Put("/{name}", or.createOriginIntegrationAsync)
			})
		})
		r.Route("/secret_keys", func(r chi.Router) {
			r.Get("/latest", or.downloadLatestOriginSecretKey)
			r.Post("/{revision}", or.uploadOriginSecretKey)
		})
		r.Get("/encryption_key", or.downloadLatestOriginEncryptionKey)
	})
	return r
}

// func (or OriginResource) listUniquePackages(w http.ResponseWriter, r *http.Request) {}
func (or OriginResource) createOrigin(w http.ResponseWriter, r *http.Request) {
	data := &OriginRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	// Get account_id from users session
	_, claims, _ := jwtauth.FromContext(r.Context())
	data.OwnerID = int64(claims["account_id"].(float64))

	if err := data.Insert(r.Context(), or.db, boil.Infer()); err != nil {
		// TODO - check for other errors here
		render.Render(w, r, ErrConflict(err))
		return
	}

	render.Status(r, http.StatusCreated)
}
func (or OriginResource) getOrigin(w http.ResponseWriter, r *http.Request) {
	origin, err := models.FindOrigin(r.Context(), or.db, chi.URLParam(r, "origin"))
	if err != nil {
		render.Render(w, r, ErrDatabase(err))
		return
	}
	render.JSON(w, r, origin)
}
func (or OriginResource) updateOrigin(w http.ResponseWriter, r *http.Request) {
	origin, err := models.FindOrigin(r.Context(), or.db, chi.URLParam(r, "origin"))
	data := &OriginRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	origin.DefaultPackageVisibility = data.DefaultPackageVisibility
	_, err = origin.Update(r.Context(), or.db, boil.Infer())
	if err != nil {
		render.Render(w, r, ErrConflict(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
func (or OriginResource) deleteOrigin(w http.ResponseWriter, r *http.Request)                      {}
func (or OriginResource) getOriginMembers(w http.ResponseWriter, r *http.Request)                  {}
func (or OriginResource) deleteOriginMember(w http.ResponseWriter, r *http.Request)                {}
func (or OriginResource) inviteToOrigin(w http.ResponseWriter, r *http.Request)                    {}
func (or OriginResource) listOriginInvitations(w http.ResponseWriter, r *http.Request)             {}
func (or OriginResource) acceptInvitation(w http.ResponseWriter, r *http.Request)                  {}
func (or OriginResource) rescindInvitation(w http.ResponseWriter, r *http.Request)                 {}
func (or OriginResource) ignoreInvitation(w http.ResponseWriter, r *http.Request)                  {}
func (or OriginResource) createKeys(w http.ResponseWriter, r *http.Request)                        {}
func (or OriginResource) listOriginKeys(w http.ResponseWriter, r *http.Request)                    {}
func (or OriginResource) downloadLatestOriginKey(w http.ResponseWriter, r *http.Request)           {}
func (or OriginResource) uploadOriginKey(w http.ResponseWriter, r *http.Request)                   {}
func (or OriginResource) downloadOriginKey(w http.ResponseWriter, r *http.Request)                 {}
func (or OriginResource) listOriginSecrets(w http.ResponseWriter, r *http.Request)                 {}
func (or OriginResource) createOriginSecret(w http.ResponseWriter, r *http.Request)                {}
func (or OriginResource) deleteOriginSecret(w http.ResponseWriter, r *http.Request)                {}
func (or OriginResource) fetchOriginIntegrations(w http.ResponseWriter, r *http.Request)           {}
func (or OriginResource) fetchOriginIntegrationNames(w http.ResponseWriter, r *http.Request)       {}
func (or OriginResource) getOriginIntegration(w http.ResponseWriter, r *http.Request)              {}
func (or OriginResource) deleteOriginIntegration(w http.ResponseWriter, r *http.Request)           {}
func (or OriginResource) createOriginIntegrationAsync(w http.ResponseWriter, r *http.Request)      {}
func (or OriginResource) downloadLatestOriginSecretKey(w http.ResponseWriter, r *http.Request)     {}
func (or OriginResource) uploadOriginSecretKey(w http.ResponseWriter, r *http.Request)             {}
func (or OriginResource) downloadLatestOriginEncryptionKey(w http.ResponseWriter, r *http.Request) {}

type OriginRequest struct {
	*models.Origin
}

func (a *OriginRequest) Bind(r *http.Request) error {
	if a.Origin == nil {
		return errors.New("missing required Origin fields")
	}
	// Disallow changing an origins name
	if r.Method == "PUT" {
		a.Name = ""
	}
	return nil
}
