package resources

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/chuckleheads/habitat/components/core/ident"
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
func (pr PkgsResource) uploadPackage(w http.ResponseWriter, r *http.Request) {
	forcedParam := r.URL.Query().Get("forced")
	isForced := false
	if forcedParam != "" || forcedParam == "true" {
		isForced = true
	}
	fmt.Println(isForced)
	pkgIdent, err := ident.New(chi.URLParam(r, "origin"), chi.URLParam(r, "pkg"), chi.URLParam(r, "version"), chi.URLParam(r, "release"))
	fmt.Println(pkgIdent)
	if err != nil {
		panic(err)
	}
	body, err := r.GetBody()
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("poopy"),
		Key:    aws.String("doodles"),
		Body:   body,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
}
func (pr PkgsResource) getPackage(w http.ResponseWriter, r *http.Request)              {}
func (pr PkgsResource) downloadPackage(w http.ResponseWriter, r *http.Request)         {}
func (pr PkgsResource) getPackageChannels(w http.ResponseWriter, r *http.Request)      {}
func (pr PkgsResource) packageVisibilityToggle(w http.ResponseWriter, r *http.Request) {}
