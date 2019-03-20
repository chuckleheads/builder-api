// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// OriginProject is an object representing the database table.
type OriginProject struct {
	ID                int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	OriginName        null.String `boil:"origin_name" json:"origin_name,omitempty" toml:"origin_name" yaml:"origin_name,omitempty"`
	PackageName       null.String `boil:"package_name" json:"package_name,omitempty" toml:"package_name" yaml:"package_name,omitempty"`
	Name              null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	PlanPath          null.String `boil:"plan_path" json:"plan_path,omitempty" toml:"plan_path" yaml:"plan_path,omitempty"`
	OwnerID           null.Int64  `boil:"owner_id" json:"owner_id,omitempty" toml:"owner_id" yaml:"owner_id,omitempty"`
	VCSType           null.String `boil:"vcs_type" json:"vcs_type,omitempty" toml:"vcs_type" yaml:"vcs_type,omitempty"`
	VCSData           null.String `boil:"vcs_data" json:"vcs_data,omitempty" toml:"vcs_data" yaml:"vcs_data,omitempty"`
	CreatedAt         null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt         null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	VCSInstallationID null.Int64  `boil:"vcs_installation_id" json:"vcs_installation_id,omitempty" toml:"vcs_installation_id" yaml:"vcs_installation_id,omitempty"`
	Visibility        string      `boil:"visibility" json:"visibility" toml:"visibility" yaml:"visibility"`
	AutoBuild         bool        `boil:"auto_build" json:"auto_build" toml:"auto_build" yaml:"auto_build"`

	R *originProjectR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L originProjectL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OriginProjectColumns = struct {
	ID                string
	OriginName        string
	PackageName       string
	Name              string
	PlanPath          string
	OwnerID           string
	VCSType           string
	VCSData           string
	CreatedAt         string
	UpdatedAt         string
	VCSInstallationID string
	Visibility        string
	AutoBuild         string
}{
	ID:                "id",
	OriginName:        "origin_name",
	PackageName:       "package_name",
	Name:              "name",
	PlanPath:          "plan_path",
	OwnerID:           "owner_id",
	VCSType:           "vcs_type",
	VCSData:           "vcs_data",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	VCSInstallationID: "vcs_installation_id",
	Visibility:        "visibility",
	AutoBuild:         "auto_build",
}

// Generated where

var OriginProjectWhere = struct {
	ID                whereHelperint64
	OriginName        whereHelpernull_String
	PackageName       whereHelpernull_String
	Name              whereHelpernull_String
	PlanPath          whereHelpernull_String
	OwnerID           whereHelpernull_Int64
	VCSType           whereHelpernull_String
	VCSData           whereHelpernull_String
	CreatedAt         whereHelpernull_Time
	UpdatedAt         whereHelpernull_Time
	VCSInstallationID whereHelpernull_Int64
	Visibility        whereHelperstring
	AutoBuild         whereHelperbool
}{
	ID:                whereHelperint64{field: `id`},
	OriginName:        whereHelpernull_String{field: `origin_name`},
	PackageName:       whereHelpernull_String{field: `package_name`},
	Name:              whereHelpernull_String{field: `name`},
	PlanPath:          whereHelpernull_String{field: `plan_path`},
	OwnerID:           whereHelpernull_Int64{field: `owner_id`},
	VCSType:           whereHelpernull_String{field: `vcs_type`},
	VCSData:           whereHelpernull_String{field: `vcs_data`},
	CreatedAt:         whereHelpernull_Time{field: `created_at`},
	UpdatedAt:         whereHelpernull_Time{field: `updated_at`},
	VCSInstallationID: whereHelpernull_Int64{field: `vcs_installation_id`},
	Visibility:        whereHelperstring{field: `visibility`},
	AutoBuild:         whereHelperbool{field: `auto_build`},
}

// OriginProjectRels is where relationship names are stored.
var OriginProjectRels = struct {
	Origin                           string
	ProjectOriginProjectIntegrations string
}{
	Origin:                           "Origin",
	ProjectOriginProjectIntegrations: "ProjectOriginProjectIntegrations",
}

// originProjectR is where relationships are stored.
type originProjectR struct {
	Origin                           *Origin
	ProjectOriginProjectIntegrations OriginProjectIntegrationSlice
}

// NewStruct creates a new relationship struct
func (*originProjectR) NewStruct() *originProjectR {
	return &originProjectR{}
}

// originProjectL is where Load methods for each relationship are stored.
type originProjectL struct{}

var (
	originProjectColumns               = []string{"id", "origin_name", "package_name", "name", "plan_path", "owner_id", "vcs_type", "vcs_data", "created_at", "updated_at", "vcs_installation_id", "visibility", "auto_build"}
	originProjectColumnsWithoutDefault = []string{"origin_name", "package_name", "name", "plan_path", "owner_id", "vcs_type", "vcs_data", "vcs_installation_id"}
	originProjectColumnsWithDefault    = []string{"id", "created_at", "updated_at", "visibility", "auto_build"}
	originProjectPrimaryKeyColumns     = []string{"id"}
)

type (
	// OriginProjectSlice is an alias for a slice of pointers to OriginProject.
	// This should generally be used opposed to []OriginProject.
	OriginProjectSlice []*OriginProject
	// OriginProjectHook is the signature for custom OriginProject hook methods
	OriginProjectHook func(context.Context, boil.ContextExecutor, *OriginProject) error

	originProjectQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	originProjectType                 = reflect.TypeOf(&OriginProject{})
	originProjectMapping              = queries.MakeStructMapping(originProjectType)
	originProjectPrimaryKeyMapping, _ = queries.BindMapping(originProjectType, originProjectMapping, originProjectPrimaryKeyColumns)
	originProjectInsertCacheMut       sync.RWMutex
	originProjectInsertCache          = make(map[string]insertCache)
	originProjectUpdateCacheMut       sync.RWMutex
	originProjectUpdateCache          = make(map[string]updateCache)
	originProjectUpsertCacheMut       sync.RWMutex
	originProjectUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var originProjectBeforeInsertHooks []OriginProjectHook
var originProjectBeforeUpdateHooks []OriginProjectHook
var originProjectBeforeDeleteHooks []OriginProjectHook
var originProjectBeforeUpsertHooks []OriginProjectHook

var originProjectAfterInsertHooks []OriginProjectHook
var originProjectAfterSelectHooks []OriginProjectHook
var originProjectAfterUpdateHooks []OriginProjectHook
var originProjectAfterDeleteHooks []OriginProjectHook
var originProjectAfterUpsertHooks []OriginProjectHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OriginProject) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originProjectBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OriginProject) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originProjectBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OriginProject) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originProjectBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OriginProject) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originProjectBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OriginProject) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originProjectAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OriginProject) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originProjectAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OriginProject) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originProjectAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OriginProject) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originProjectAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OriginProject) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originProjectAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOriginProjectHook registers your hook function for all future operations.
func AddOriginProjectHook(hookPoint boil.HookPoint, originProjectHook OriginProjectHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		originProjectBeforeInsertHooks = append(originProjectBeforeInsertHooks, originProjectHook)
	case boil.BeforeUpdateHook:
		originProjectBeforeUpdateHooks = append(originProjectBeforeUpdateHooks, originProjectHook)
	case boil.BeforeDeleteHook:
		originProjectBeforeDeleteHooks = append(originProjectBeforeDeleteHooks, originProjectHook)
	case boil.BeforeUpsertHook:
		originProjectBeforeUpsertHooks = append(originProjectBeforeUpsertHooks, originProjectHook)
	case boil.AfterInsertHook:
		originProjectAfterInsertHooks = append(originProjectAfterInsertHooks, originProjectHook)
	case boil.AfterSelectHook:
		originProjectAfterSelectHooks = append(originProjectAfterSelectHooks, originProjectHook)
	case boil.AfterUpdateHook:
		originProjectAfterUpdateHooks = append(originProjectAfterUpdateHooks, originProjectHook)
	case boil.AfterDeleteHook:
		originProjectAfterDeleteHooks = append(originProjectAfterDeleteHooks, originProjectHook)
	case boil.AfterUpsertHook:
		originProjectAfterUpsertHooks = append(originProjectAfterUpsertHooks, originProjectHook)
	}
}

// One returns a single originProject record from the query.
func (q originProjectQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OriginProject, error) {
	o := &OriginProject{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for origin_projects")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OriginProject records from the query.
func (q originProjectQuery) All(ctx context.Context, exec boil.ContextExecutor) (OriginProjectSlice, error) {
	var o []*OriginProject

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OriginProject slice")
	}

	if len(originProjectAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OriginProject records in the query.
func (q originProjectQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count origin_projects rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q originProjectQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if origin_projects exists")
	}

	return count > 0, nil
}

// Origin pointed to by the foreign key.
func (o *OriginProject) Origin(mods ...qm.QueryMod) originQuery {
	queryMods := []qm.QueryMod{
		qm.Where("name=?", o.OriginName),
	}

	queryMods = append(queryMods, mods...)

	query := Origins(queryMods...)
	queries.SetFrom(query.Query, "\"origins\"")

	return query
}

// ProjectOriginProjectIntegrations retrieves all the origin_project_integration's OriginProjectIntegrations with an executor via project_id column.
func (o *OriginProject) ProjectOriginProjectIntegrations(mods ...qm.QueryMod) originProjectIntegrationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"origin_project_integrations\".\"project_id\"=?", o.ID),
	)

	query := OriginProjectIntegrations(queryMods...)
	queries.SetFrom(query.Query, "\"origin_project_integrations\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"origin_project_integrations\".*"})
	}

	return query
}

// LoadOrigin allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (originProjectL) LoadOrigin(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOriginProject interface{}, mods queries.Applicator) error {
	var slice []*OriginProject
	var object *OriginProject

	if singular {
		object = maybeOriginProject.(*OriginProject)
	} else {
		slice = *maybeOriginProject.(*[]*OriginProject)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &originProjectR{}
		}
		if !queries.IsNil(object.OriginName) {
			args = append(args, object.OriginName)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &originProjectR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.OriginName) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.OriginName) {
				args = append(args, obj.OriginName)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`origins`), qm.WhereIn(`name in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Origin")
	}

	var resultSlice []*Origin
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Origin")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for origins")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for origins")
	}

	if len(originProjectAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Origin = foreign
		if foreign.R == nil {
			foreign.R = &originR{}
		}
		foreign.R.OriginOPR = append(foreign.R.OriginOPR, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.OriginName, foreign.Name) {
				local.R.Origin = foreign
				if foreign.R == nil {
					foreign.R = &originR{}
				}
				foreign.R.OriginOPR = append(foreign.R.OriginOPR, local)
				break
			}
		}
	}

	return nil
}

// LoadProjectOriginProjectIntegrations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (originProjectL) LoadProjectOriginProjectIntegrations(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOriginProject interface{}, mods queries.Applicator) error {
	var slice []*OriginProject
	var object *OriginProject

	if singular {
		object = maybeOriginProject.(*OriginProject)
	} else {
		slice = *maybeOriginProject.(*[]*OriginProject)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &originProjectR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &originProjectR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`origin_project_integrations`), qm.WhereIn(`project_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load origin_project_integrations")
	}

	var resultSlice []*OriginProjectIntegration
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice origin_project_integrations")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on origin_project_integrations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for origin_project_integrations")
	}

	if len(originProjectIntegrationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ProjectOriginProjectIntegrations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &originProjectIntegrationR{}
			}
			foreign.R.Project = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ProjectID {
				local.R.ProjectOriginProjectIntegrations = append(local.R.ProjectOriginProjectIntegrations, foreign)
				if foreign.R == nil {
					foreign.R = &originProjectIntegrationR{}
				}
				foreign.R.Project = local
				break
			}
		}
	}

	return nil
}

// SetOrigin of the originProject to the related item.
// Sets o.R.Origin to related.
// Adds o to related.R.OriginOPR.
func (o *OriginProject) SetOrigin(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Origin) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"origin_projects\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"origin_name"}),
		strmangle.WhereClause("\"", "\"", 2, originProjectPrimaryKeyColumns),
	)
	values := []interface{}{related.Name, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.OriginName, related.Name)
	if o.R == nil {
		o.R = &originProjectR{
			Origin: related,
		}
	} else {
		o.R.Origin = related
	}

	if related.R == nil {
		related.R = &originR{
			OriginOPR: OriginProjectSlice{o},
		}
	} else {
		related.R.OriginOPR = append(related.R.OriginOPR, o)
	}

	return nil
}

// RemoveOrigin relationship.
// Sets o.R.Origin to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *OriginProject) RemoveOrigin(ctx context.Context, exec boil.ContextExecutor, related *Origin) error {
	var err error

	queries.SetScanner(&o.OriginName, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("origin_name")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Origin = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.OriginOPR {
		if queries.Equal(o.OriginName, ri.OriginName) {
			continue
		}

		ln := len(related.R.OriginOPR)
		if ln > 1 && i < ln-1 {
			related.R.OriginOPR[i] = related.R.OriginOPR[ln-1]
		}
		related.R.OriginOPR = related.R.OriginOPR[:ln-1]
		break
	}
	return nil
}

// AddProjectOriginProjectIntegrations adds the given related objects to the existing relationships
// of the origin_project, optionally inserting them as new records.
// Appends related to o.R.ProjectOriginProjectIntegrations.
// Sets related.R.Project appropriately.
func (o *OriginProject) AddProjectOriginProjectIntegrations(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*OriginProjectIntegration) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProjectID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"origin_project_integrations\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"project_id"}),
				strmangle.WhereClause("\"", "\"", 2, originProjectIntegrationPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProjectID = o.ID
		}
	}

	if o.R == nil {
		o.R = &originProjectR{
			ProjectOriginProjectIntegrations: related,
		}
	} else {
		o.R.ProjectOriginProjectIntegrations = append(o.R.ProjectOriginProjectIntegrations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &originProjectIntegrationR{
				Project: o,
			}
		} else {
			rel.R.Project = o
		}
	}
	return nil
}

// OriginProjects retrieves all the records using an executor.
func OriginProjects(mods ...qm.QueryMod) originProjectQuery {
	mods = append(mods, qm.From("\"origin_projects\""))
	return originProjectQuery{NewQuery(mods...)}
}

// FindOriginProject retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOriginProject(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*OriginProject, error) {
	originProjectObj := &OriginProject{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"origin_projects\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, originProjectObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from origin_projects")
	}

	return originProjectObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OriginProject) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no origin_projects provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(originProjectColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	originProjectInsertCacheMut.RLock()
	cache, cached := originProjectInsertCache[key]
	originProjectInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			originProjectColumns,
			originProjectColumnsWithDefault,
			originProjectColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(originProjectType, originProjectMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(originProjectType, originProjectMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"origin_projects\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"origin_projects\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into origin_projects")
	}

	if !cached {
		originProjectInsertCacheMut.Lock()
		originProjectInsertCache[key] = cache
		originProjectInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OriginProject.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OriginProject) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	originProjectUpdateCacheMut.RLock()
	cache, cached := originProjectUpdateCache[key]
	originProjectUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			originProjectColumns,
			originProjectPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update origin_projects, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"origin_projects\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, originProjectPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(originProjectType, originProjectMapping, append(wl, originProjectPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update origin_projects row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for origin_projects")
	}

	if !cached {
		originProjectUpdateCacheMut.Lock()
		originProjectUpdateCache[key] = cache
		originProjectUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q originProjectQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for origin_projects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for origin_projects")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OriginProjectSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), originProjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"origin_projects\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, originProjectPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in originProject slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all originProject")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OriginProject) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no origin_projects provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(originProjectColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	originProjectUpsertCacheMut.RLock()
	cache, cached := originProjectUpsertCache[key]
	originProjectUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			originProjectColumns,
			originProjectColumnsWithDefault,
			originProjectColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			originProjectColumns,
			originProjectPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert origin_projects, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(originProjectPrimaryKeyColumns))
			copy(conflict, originProjectPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"origin_projects\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(originProjectType, originProjectMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(originProjectType, originProjectMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert origin_projects")
	}

	if !cached {
		originProjectUpsertCacheMut.Lock()
		originProjectUpsertCache[key] = cache
		originProjectUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OriginProject record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OriginProject) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OriginProject provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), originProjectPrimaryKeyMapping)
	sql := "DELETE FROM \"origin_projects\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from origin_projects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for origin_projects")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q originProjectQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no originProjectQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from origin_projects")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for origin_projects")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OriginProjectSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OriginProject slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(originProjectBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), originProjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"origin_projects\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, originProjectPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from originProject slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for origin_projects")
	}

	if len(originProjectAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OriginProject) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOriginProject(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OriginProjectSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OriginProjectSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), originProjectPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"origin_projects\".* FROM \"origin_projects\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, originProjectPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OriginProjectSlice")
	}

	*o = slice

	return nil
}

// OriginProjectExists checks if the OriginProject row exists.
func OriginProjectExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"origin_projects\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if origin_projects exists")
	}

	return exists, nil
}