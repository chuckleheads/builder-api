package cmd

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/spf13/cobra"

	"github.com/chuckleheads/builder-api/internal/datastore"
	"github.com/chuckleheads/builder-api/internal/oauth"
	"github.com/chuckleheads/builder-api/internal/resources"
)

var tokenAuth *jwtauth.JWTAuth

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start builder api",
	Run: func(cmd *cobra.Command, args []string) {
		setup()
	},
}

func debug() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For testing purposes, we generate and print
	_, tokenString, _ := tokenAuth.Encode(jwt.MapClaims{"account_id": 123, "extern_id": 456, "oauth_token": "poopydoodles"})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func setup() {
	config, err := ConfigFromViper()
	if err != nil {
		panic(err.Error())
	}
	debug()
	addr := fmt.Sprintf(":%d", config.Port)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	db := datastore.New(&config.Datastore)
	oauthConfig := oauth.New(&config.Oauth)

	r.Mount("/authenticate", resources.NewAuthenticateResource(oauthConfig).Routes())
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Mount("/depot/channels", resources.ChannelResource{}.Routes())
		r.Mount("/ext", resources.ExtResource{}.Routes())
		r.Mount("/jobs", resources.JobsResource{}.Routes())
		r.Mount("/notify", resources.NotifyResource{}.Routes())
		r.Mount("/depot/origins", resources.NewOriginResource(db).Routes())
		r.Mount("/depot/pkgs", resources.PkgsResource{}.Routes())
		r.Mount("/profile", resources.ProfileResource{}.Routes())
		r.Mount("/projects", resources.ProjectsResource{}.Routes())
		r.Mount("/rdeps", resources.RdepsResource{}.Routes())
		r.Mount("/user", resources.UserResource{}.Routes())
	})

	// walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	// 	route = strings.Replace(route, "/*/", "/", -1)
	// 	fmt.Printf("%s %s\n", method, route)
	// 	return nil
	// }

	// if err := chi.Walk(r, walkFunc); err != nil {
	// 	fmt.Printf("Logging err: %s\n", err.Error())
	// }

	log.Printf("HTTP Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
