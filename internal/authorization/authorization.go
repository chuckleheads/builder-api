package authorization

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"net/http"

	"github.com/chuckleheads/builder-api/pkg/models"
)

func AuthorizeSession(r *http.Request) {

}

func CreateSessionOauth(r *http.Request, db *sql.DB, oauth_token string, account models.Account, provider string) (*models.Account, error) {
	account.Upsert(r.Context(), db, true, []string{"id", "name", "email"}, boil.Whitelist("name"), boil.Infer())
	return models.Accounts(qm.Where("name = ?", account.Name)).One(r.Context(), db)
}
