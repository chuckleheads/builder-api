package resources

import (
	"github.com/go-chi/render"
	"golang.org/x/oauth2"
	"net/http"

	"github.com/go-chi/chi"
)

type AuthenticateResource struct {
	oauth *oauth2.Config
}

func NewAuthenticateResource(oauth *oauth2.Config) AuthenticateResource {
	return AuthenticateResource{oauth}
}

func (ar AuthenticateResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/{code}", ar.authenticate)
	return r
}

func (ar AuthenticateResource) authenticate(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	r.Context()
	token, err := ar.oauth.Exchange(r.Context(), code)
	if err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	render.PlainText(w, r, token.AccessToken)
}
