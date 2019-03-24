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

// AuditPackage is an object representing the database table.
type AuditPackage struct {
	ID            int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Operation     null.String `boil:"operation" json:"operation,omitempty" toml:"operation" yaml:"operation,omitempty"`
	Trigger       null.String `boil:"trigger" json:"trigger,omitempty" toml:"trigger" yaml:"trigger,omitempty"`
	RequesterID   null.Int64  `boil:"requester_id" json:"requester_id,omitempty" toml:"requester_id" yaml:"requester_id,omitempty"`
	RequesterName null.String `boil:"requester_name" json:"requester_name,omitempty" toml:"requester_name" yaml:"requester_name,omitempty"`
	CreatedAt     null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	Origin        null.String `boil:"origin" json:"origin,omitempty" toml:"origin" yaml:"origin,omitempty"`
	Channel       null.String `boil:"channel" json:"channel,omitempty" toml:"channel" yaml:"channel,omitempty"`
	PackageIdent  null.String `boil:"package_ident" json:"package_ident,omitempty" toml:"package_ident" yaml:"package_ident,omitempty"`

	R *auditPackageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L auditPackageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuditPackageColumns = struct {
	ID            string
	Operation     string
	Trigger       string
	RequesterID   string
	RequesterName string
	CreatedAt     string
	Origin        string
	Channel       string
	PackageIdent  string
}{
	ID:            "id",
	Operation:     "operation",
	Trigger:       "trigger",
	RequesterID:   "requester_id",
	RequesterName: "requester_name",
	CreatedAt:     "created_at",
	Origin:        "origin",
	Channel:       "channel",
	PackageIdent:  "package_ident",
}

// Generated where

var AuditPackageWhere = struct {
	ID            whereHelperint
	Operation     whereHelpernull_String
	Trigger       whereHelpernull_String
	RequesterID   whereHelpernull_Int64
	RequesterName whereHelpernull_String
	CreatedAt     whereHelpernull_Time
	Origin        whereHelpernull_String
	Channel       whereHelpernull_String
	PackageIdent  whereHelpernull_String
}{
	ID:            whereHelperint{field: `id`},
	Operation:     whereHelpernull_String{field: `operation`},
	Trigger:       whereHelpernull_String{field: `trigger`},
	RequesterID:   whereHelpernull_Int64{field: `requester_id`},
	RequesterName: whereHelpernull_String{field: `requester_name`},
	CreatedAt:     whereHelpernull_Time{field: `created_at`},
	Origin:        whereHelpernull_String{field: `origin`},
	Channel:       whereHelpernull_String{field: `channel`},
	PackageIdent:  whereHelpernull_String{field: `package_ident`},
}

// AuditPackageRels is where relationship names are stored.
var AuditPackageRels = struct {
	OriginName string
}{
	OriginName: "OriginName",
}

// auditPackageR is where relationships are stored.
type auditPackageR struct {
	OriginName *Origin
}

// NewStruct creates a new relationship struct
func (*auditPackageR) NewStruct() *auditPackageR {
	return &auditPackageR{}
}

// auditPackageL is where Load methods for each relationship are stored.
type auditPackageL struct{}

var (
	auditPackageColumns               = []string{"id", "operation", "trigger", "requester_id", "requester_name", "created_at", "origin", "channel", "package_ident"}
	auditPackageColumnsWithoutDefault = []string{"operation", "trigger", "requester_id", "requester_name", "origin", "channel", "package_ident"}
	auditPackageColumnsWithDefault    = []string{"id", "created_at"}
	auditPackagePrimaryKeyColumns     = []string{"id"}
)

type (
	// AuditPackageSlice is an alias for a slice of pointers to AuditPackage.
	// This should generally be used opposed to []AuditPackage.
	AuditPackageSlice []*AuditPackage
	// AuditPackageHook is the signature for custom AuditPackage hook methods
	AuditPackageHook func(context.Context, boil.ContextExecutor, *AuditPackage) error

	auditPackageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	auditPackageType                 = reflect.TypeOf(&AuditPackage{})
	auditPackageMapping              = queries.MakeStructMapping(auditPackageType)
	auditPackagePrimaryKeyMapping, _ = queries.BindMapping(auditPackageType, auditPackageMapping, auditPackagePrimaryKeyColumns)
	auditPackageInsertCacheMut       sync.RWMutex
	auditPackageInsertCache          = make(map[string]insertCache)
	auditPackageUpdateCacheMut       sync.RWMutex
	auditPackageUpdateCache          = make(map[string]updateCache)
	auditPackageUpsertCacheMut       sync.RWMutex
	auditPackageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var auditPackageBeforeInsertHooks []AuditPackageHook
var auditPackageBeforeUpdateHooks []AuditPackageHook
var auditPackageBeforeDeleteHooks []AuditPackageHook
var auditPackageBeforeUpsertHooks []AuditPackageHook

var auditPackageAfterInsertHooks []AuditPackageHook
var auditPackageAfterSelectHooks []AuditPackageHook
var auditPackageAfterUpdateHooks []AuditPackageHook
var auditPackageAfterDeleteHooks []AuditPackageHook
var auditPackageAfterUpsertHooks []AuditPackageHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuditPackage) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditPackageBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuditPackage) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditPackageBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuditPackage) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditPackageBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuditPackage) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditPackageBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuditPackage) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditPackageAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuditPackage) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditPackageAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuditPackage) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditPackageAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuditPackage) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditPackageAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuditPackage) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditPackageAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuditPackageHook registers your hook function for all future operations.
func AddAuditPackageHook(hookPoint boil.HookPoint, auditPackageHook AuditPackageHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		auditPackageBeforeInsertHooks = append(auditPackageBeforeInsertHooks, auditPackageHook)
	case boil.BeforeUpdateHook:
		auditPackageBeforeUpdateHooks = append(auditPackageBeforeUpdateHooks, auditPackageHook)
	case boil.BeforeDeleteHook:
		auditPackageBeforeDeleteHooks = append(auditPackageBeforeDeleteHooks, auditPackageHook)
	case boil.BeforeUpsertHook:
		auditPackageBeforeUpsertHooks = append(auditPackageBeforeUpsertHooks, auditPackageHook)
	case boil.AfterInsertHook:
		auditPackageAfterInsertHooks = append(auditPackageAfterInsertHooks, auditPackageHook)
	case boil.AfterSelectHook:
		auditPackageAfterSelectHooks = append(auditPackageAfterSelectHooks, auditPackageHook)
	case boil.AfterUpdateHook:
		auditPackageAfterUpdateHooks = append(auditPackageAfterUpdateHooks, auditPackageHook)
	case boil.AfterDeleteHook:
		auditPackageAfterDeleteHooks = append(auditPackageAfterDeleteHooks, auditPackageHook)
	case boil.AfterUpsertHook:
		auditPackageAfterUpsertHooks = append(auditPackageAfterUpsertHooks, auditPackageHook)
	}
}

// One returns a single auditPackage record from the query.
func (q auditPackageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AuditPackage, error) {
	o := &AuditPackage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for audit_package")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AuditPackage records from the query.
func (q auditPackageQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuditPackageSlice, error) {
	var o []*AuditPackage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuditPackage slice")
	}

	if len(auditPackageAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AuditPackage records in the query.
func (q auditPackageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count audit_package rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q auditPackageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if audit_package exists")
	}

	return count > 0, nil
}

// OriginName pointed to by the foreign key.
func (o *AuditPackage) OriginName(mods ...qm.QueryMod) originQuery {
	queryMods := []qm.QueryMod{
		qm.Where("name=?", o.Origin),
	}

	queryMods = append(queryMods, mods...)

	query := Origins(queryMods...)
	queries.SetFrom(query.Query, "\"origins\"")

	return query
}

// LoadOriginName allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (auditPackageL) LoadOriginName(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAuditPackage interface{}, mods queries.Applicator) error {
	var slice []*AuditPackage
	var object *AuditPackage

	if singular {
		object = maybeAuditPackage.(*AuditPackage)
	} else {
		slice = *maybeAuditPackage.(*[]*AuditPackage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &auditPackageR{}
		}
		if !queries.IsNil(object.Origin) {
			args = append(args, object.Origin)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &auditPackageR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.Origin) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.Origin) {
				args = append(args, obj.Origin)
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

	if len(auditPackageAfterSelectHooks) != 0 {
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
		object.R.OriginName = foreign
		if foreign.R == nil {
			foreign.R = &originR{}
		}
		foreign.R.OriginAP = append(foreign.R.OriginAP, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.Origin, foreign.Name) {
				local.R.OriginName = foreign
				if foreign.R == nil {
					foreign.R = &originR{}
				}
				foreign.R.OriginAP = append(foreign.R.OriginAP, local)
				break
			}
		}
	}

	return nil
}

// SetOriginName of the auditPackage to the related item.
// Sets o.R.OriginName to related.
// Adds o to related.R.OriginAP.
func (o *AuditPackage) SetOriginName(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Origin) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"audit_package\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"origin"}),
		strmangle.WhereClause("\"", "\"", 2, auditPackagePrimaryKeyColumns),
	)
	values := []interface{}{related.Name, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.Origin, related.Name)
	if o.R == nil {
		o.R = &auditPackageR{
			OriginName: related,
		}
	} else {
		o.R.OriginName = related
	}

	if related.R == nil {
		related.R = &originR{
			OriginAP: AuditPackageSlice{o},
		}
	} else {
		related.R.OriginAP = append(related.R.OriginAP, o)
	}

	return nil
}

// RemoveOriginName relationship.
// Sets o.R.OriginName to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *AuditPackage) RemoveOriginName(ctx context.Context, exec boil.ContextExecutor, related *Origin) error {
	var err error

	queries.SetScanner(&o.Origin, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("origin")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.OriginName = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.OriginAP {
		if queries.Equal(o.Origin, ri.Origin) {
			continue
		}

		ln := len(related.R.OriginAP)
		if ln > 1 && i < ln-1 {
			related.R.OriginAP[i] = related.R.OriginAP[ln-1]
		}
		related.R.OriginAP = related.R.OriginAP[:ln-1]
		break
	}
	return nil
}

// AuditPackages retrieves all the records using an executor.
func AuditPackages(mods ...qm.QueryMod) auditPackageQuery {
	mods = append(mods, qm.From("\"audit_package\""))
	return auditPackageQuery{NewQuery(mods...)}
}

// FindAuditPackage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuditPackage(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*AuditPackage, error) {
	auditPackageObj := &AuditPackage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"audit_package\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, auditPackageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from audit_package")
	}

	return auditPackageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AuditPackage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no audit_package provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(auditPackageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	auditPackageInsertCacheMut.RLock()
	cache, cached := auditPackageInsertCache[key]
	auditPackageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			auditPackageColumns,
			auditPackageColumnsWithDefault,
			auditPackageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(auditPackageType, auditPackageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(auditPackageType, auditPackageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"audit_package\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"audit_package\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into audit_package")
	}

	if !cached {
		auditPackageInsertCacheMut.Lock()
		auditPackageInsertCache[key] = cache
		auditPackageInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AuditPackage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AuditPackage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	auditPackageUpdateCacheMut.RLock()
	cache, cached := auditPackageUpdateCache[key]
	auditPackageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			auditPackageColumns,
			auditPackagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update audit_package, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"audit_package\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, auditPackagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(auditPackageType, auditPackageMapping, append(wl, auditPackagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update audit_package row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for audit_package")
	}

	if !cached {
		auditPackageUpdateCacheMut.Lock()
		auditPackageUpdateCache[key] = cache
		auditPackageUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q auditPackageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for audit_package")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for audit_package")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuditPackageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), auditPackagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"audit_package\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, auditPackagePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in auditPackage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all auditPackage")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AuditPackage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no audit_package provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(auditPackageColumnsWithDefault, o)

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

	auditPackageUpsertCacheMut.RLock()
	cache, cached := auditPackageUpsertCache[key]
	auditPackageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			auditPackageColumns,
			auditPackageColumnsWithDefault,
			auditPackageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			auditPackageColumns,
			auditPackagePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert audit_package, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(auditPackagePrimaryKeyColumns))
			copy(conflict, auditPackagePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"audit_package\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(auditPackageType, auditPackageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(auditPackageType, auditPackageMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert audit_package")
	}

	if !cached {
		auditPackageUpsertCacheMut.Lock()
		auditPackageUpsertCache[key] = cache
		auditPackageUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AuditPackage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuditPackage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AuditPackage provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), auditPackagePrimaryKeyMapping)
	sql := "DELETE FROM \"audit_package\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from audit_package")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for audit_package")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q auditPackageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no auditPackageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from audit_package")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for audit_package")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuditPackageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AuditPackage slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(auditPackageBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), auditPackagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"audit_package\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, auditPackagePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from auditPackage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for audit_package")
	}

	if len(auditPackageAfterDeleteHooks) != 0 {
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
func (o *AuditPackage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuditPackage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuditPackageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuditPackageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), auditPackagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"audit_package\".* FROM \"audit_package\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, auditPackagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuditPackageSlice")
	}

	*o = slice

	return nil
}

// AuditPackageExists checks if the AuditPackage row exists.
func AuditPackageExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"audit_package\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if audit_package exists")
	}

	return exists, nil
}
