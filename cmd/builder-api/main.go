package main

//go:generate sqlboiler --wipe --output "../../pkg/models" -c ../../sqlboiler.toml psql

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/chuckleheads/builder-api/internal/resources"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db, err := sql.Open("postgres", "dbname=builder user=elliott.davis sslmode=disable")
	if err != nil {
		panic(err)
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Mount("/authenticate", resources.AuthenticateResource{}.Routes())
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

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	http.ListenAndServe(":3333", r)
}
