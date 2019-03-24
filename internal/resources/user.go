package resources

import (
	"net/http"

	"github.com/go-chi/chi"
)

type UserResource struct{}

func (ur UserResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/invitations", ur.getInvitations)
	r.Get("/origins", ur.getOrigins)
	return r
}
func (ur UserResource) getInvitations(w http.ResponseWriter, r *http.Request) {}
func (ur UserResource) getOrigins(w http.ResponseWriter, r *http.Request) {

}
