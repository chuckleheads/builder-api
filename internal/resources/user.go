package resources

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/chuckleheads/builder-api/pkg/models"
)

type UserResource struct {
	db *sql.DB
}

func NewUserResource(db *sql.DB) *UserResource {
	return &UserResource{db}
}

func (ur UserResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/invitations", ur.getInvitations)
	r.Get("/origins", ur.getOrigins)
	return r
}
func (ur UserResource) getInvitations(w http.ResponseWriter, r *http.Request) {}
func (ur UserResource) getOrigins(w http.ResponseWriter, r *http.Request) {
	boil.DebugMode = true

	_, claims, _ := jwtauth.FromContext(r.Context())
	origins, err := models.Origins(qm.Where("owner_id = ?", claims["account_id"].(float64))).All(r.Context(), ur.db)
	if err != nil {
		render.Render(w, r, ErrDatabase(err))
		return
	}
	render.JSON(w, r, origins)
}
