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

// OriginMember is an object representing the database table.
type OriginMember struct {
	ID         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	OriginName null.String `boil:"origin_name" json:"origin_name,omitempty" toml:"origin_name" yaml:"origin_name,omitempty"`
	AccountID  int64       `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	CreatedAt  null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt  null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *originMemberR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L originMemberL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OriginMemberColumns = struct {
	ID         string
	OriginName string
	AccountID  string
	CreatedAt  string
	UpdatedAt  string
}{
	ID:         "id",
	OriginName: "origin_name",
	AccountID:  "account_id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// Generated where

var OriginMemberWhere = struct {
	ID         whereHelperint
	OriginName whereHelpernull_String
	AccountID  whereHelperint64
	CreatedAt  whereHelpernull_Time
	UpdatedAt  whereHelpernull_Time
}{
	ID:         whereHelperint{field: `id`},
	OriginName: whereHelpernull_String{field: `origin_name`},
	AccountID:  whereHelperint64{field: `account_id`},
	CreatedAt:  whereHelpernull_Time{field: `created_at`},
	UpdatedAt:  whereHelpernull_Time{field: `updated_at`},
}

// OriginMemberRels is where relationship names are stored.
var OriginMemberRels = struct {
	Origin string
}{
	Origin: "Origin",
}

// originMemberR is where relationships are stored.
type originMemberR struct {
	Origin *Origin
}

// NewStruct creates a new relationship struct
func (*originMemberR) NewStruct() *originMemberR {
	return &originMemberR{}
}

// originMemberL is where Load methods for each relationship are stored.
type originMemberL struct{}

var (
	originMemberColumns               = []string{"id", "origin_name", "account_id", "created_at", "updated_at"}
	originMemberColumnsWithoutDefault = []string{"origin_name", "account_id"}
	originMemberColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	originMemberPrimaryKeyColumns     = []string{"id"}
)

type (
	// OriginMemberSlice is an alias for a slice of pointers to OriginMember.
	// This should generally be used opposed to []OriginMember.
	OriginMemberSlice []*OriginMember
	// OriginMemberHook is the signature for custom OriginMember hook methods
	OriginMemberHook func(context.Context, boil.ContextExecutor, *OriginMember) error

	originMemberQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	originMemberType                 = reflect.TypeOf(&OriginMember{})
	originMemberMapping              = queries.MakeStructMapping(originMemberType)
	originMemberPrimaryKeyMapping, _ = queries.BindMapping(originMemberType, originMemberMapping, originMemberPrimaryKeyColumns)
	originMemberInsertCacheMut       sync.RWMutex
	originMemberInsertCache          = make(map[string]insertCache)
	originMemberUpdateCacheMut       sync.RWMutex
	originMemberUpdateCache          = make(map[string]updateCache)
	originMemberUpsertCacheMut       sync.RWMutex
	originMemberUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var originMemberBeforeInsertHooks []OriginMemberHook
var originMemberBeforeUpdateHooks []OriginMemberHook
var originMemberBeforeDeleteHooks []OriginMemberHook
var originMemberBeforeUpsertHooks []OriginMemberHook

var originMemberAfterInsertHooks []OriginMemberHook
var originMemberAfterSelectHooks []OriginMemberHook
var originMemberAfterUpdateHooks []OriginMemberHook
var originMemberAfterDeleteHooks []OriginMemberHook
var originMemberAfterUpsertHooks []OriginMemberHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OriginMember) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originMemberBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OriginMember) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originMemberBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OriginMember) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originMemberBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OriginMember) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originMemberBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OriginMember) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originMemberAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OriginMember) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originMemberAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OriginMember) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originMemberAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OriginMember) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originMemberAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OriginMember) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originMemberAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOriginMemberHook registers your hook function for all future operations.
func AddOriginMemberHook(hookPoint boil.HookPoint, originMemberHook OriginMemberHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		originMemberBeforeInsertHooks = append(originMemberBeforeInsertHooks, originMemberHook)
	case boil.BeforeUpdateHook:
		originMemberBeforeUpdateHooks = append(originMemberBeforeUpdateHooks, originMemberHook)
	case boil.BeforeDeleteHook:
		originMemberBeforeDeleteHooks = append(originMemberBeforeDeleteHooks, originMemberHook)
	case boil.BeforeUpsertHook:
		originMemberBeforeUpsertHooks = append(originMemberBeforeUpsertHooks, originMemberHook)
	case boil.AfterInsertHook:
		originMemberAfterInsertHooks = append(originMemberAfterInsertHooks, originMemberHook)
	case boil.AfterSelectHook:
		originMemberAfterSelectHooks = append(originMemberAfterSelectHooks, originMemberHook)
	case boil.AfterUpdateHook:
		originMemberAfterUpdateHooks = append(originMemberAfterUpdateHooks, originMemberHook)
	case boil.AfterDeleteHook:
		originMemberAfterDeleteHooks = append(originMemberAfterDeleteHooks, originMemberHook)
	case boil.AfterUpsertHook:
		originMemberAfterUpsertHooks = append(originMemberAfterUpsertHooks, originMemberHook)
	}
}

// One returns a single originMember record from the query.
func (q originMemberQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OriginMember, error) {
	o := &OriginMember{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for origin_members")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OriginMember records from the query.
func (q originMemberQuery) All(ctx context.Context, exec boil.ContextExecutor) (OriginMemberSlice, error) {
	var o []*OriginMember

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OriginMember slice")
	}

	if len(originMemberAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OriginMember records in the query.
func (q originMemberQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count origin_members rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q originMemberQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if origin_members exists")
	}

	return count > 0, nil
}

// Origin pointed to by the foreign key.
func (o *OriginMember) Origin(mods ...qm.QueryMod) originQuery {
	queryMods := []qm.QueryMod{
		qm.Where("name=?", o.OriginName),
	}

	queryMods = append(queryMods, mods...)

	query := Origins(queryMods...)
	queries.SetFrom(query.Query, "\"origins\"")

	return query
}

// LoadOrigin allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (originMemberL) LoadOrigin(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOriginMember interface{}, mods queries.Applicator) error {
	var slice []*OriginMember
	var object *OriginMember

	if singular {
		object = maybeOriginMember.(*OriginMember)
	} else {
		slice = *maybeOriginMember.(*[]*OriginMember)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &originMemberR{}
		}
		if !queries.IsNil(object.OriginName) {
			args = append(args, object.OriginName)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &originMemberR{}
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

	if len(originMemberAfterSelectHooks) != 0 {
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
		foreign.R.OriginOM = append(foreign.R.OriginOM, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.OriginName, foreign.Name) {
				local.R.Origin = foreign
				if foreign.R == nil {
					foreign.R = &originR{}
				}
				foreign.R.OriginOM = append(foreign.R.OriginOM, local)
				break
			}
		}
	}

	return nil
}

// SetOrigin of the originMember to the related item.
// Sets o.R.Origin to related.
// Adds o to related.R.OriginOM.
func (o *OriginMember) SetOrigin(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Origin) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"origin_members\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"origin_name"}),
		strmangle.WhereClause("\"", "\"", 2, originMemberPrimaryKeyColumns),
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
		o.R = &originMemberR{
			Origin: related,
		}
	} else {
		o.R.Origin = related
	}

	if related.R == nil {
		related.R = &originR{
			OriginOM: OriginMemberSlice{o},
		}
	} else {
		related.R.OriginOM = append(related.R.OriginOM, o)
	}

	return nil
}

// RemoveOrigin relationship.
// Sets o.R.Origin to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *OriginMember) RemoveOrigin(ctx context.Context, exec boil.ContextExecutor, related *Origin) error {
	var err error

	queries.SetScanner(&o.OriginName, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("origin_name")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Origin = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.OriginOM {
		if queries.Equal(o.OriginName, ri.OriginName) {
			continue
		}

		ln := len(related.R.OriginOM)
		if ln > 1 && i < ln-1 {
			related.R.OriginOM[i] = related.R.OriginOM[ln-1]
		}
		related.R.OriginOM = related.R.OriginOM[:ln-1]
		break
	}
	return nil
}

// OriginMembers retrieves all the records using an executor.
func OriginMembers(mods ...qm.QueryMod) originMemberQuery {
	mods = append(mods, qm.From("\"origin_members\""))
	return originMemberQuery{NewQuery(mods...)}
}

// FindOriginMember retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOriginMember(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*OriginMember, error) {
	originMemberObj := &OriginMember{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"origin_members\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, originMemberObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from origin_members")
	}

	return originMemberObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OriginMember) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no origin_members provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(originMemberColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	originMemberInsertCacheMut.RLock()
	cache, cached := originMemberInsertCache[key]
	originMemberInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			originMemberColumns,
			originMemberColumnsWithDefault,
			originMemberColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(originMemberType, originMemberMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(originMemberType, originMemberMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"origin_members\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"origin_members\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into origin_members")
	}

	if !cached {
		originMemberInsertCacheMut.Lock()
		originMemberInsertCache[key] = cache
		originMemberInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OriginMember.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OriginMember) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	originMemberUpdateCacheMut.RLock()
	cache, cached := originMemberUpdateCache[key]
	originMemberUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			originMemberColumns,
			originMemberPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update origin_members, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"origin_members\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, originMemberPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(originMemberType, originMemberMapping, append(wl, originMemberPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update origin_members row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for origin_members")
	}

	if !cached {
		originMemberUpdateCacheMut.Lock()
		originMemberUpdateCache[key] = cache
		originMemberUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q originMemberQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for origin_members")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for origin_members")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OriginMemberSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), originMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"origin_members\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, originMemberPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in originMember slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all originMember")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OriginMember) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no origin_members provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(originMemberColumnsWithDefault, o)

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

	originMemberUpsertCacheMut.RLock()
	cache, cached := originMemberUpsertCache[key]
	originMemberUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			originMemberColumns,
			originMemberColumnsWithDefault,
			originMemberColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			originMemberColumns,
			originMemberPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert origin_members, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(originMemberPrimaryKeyColumns))
			copy(conflict, originMemberPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"origin_members\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(originMemberType, originMemberMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(originMemberType, originMemberMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert origin_members")
	}

	if !cached {
		originMemberUpsertCacheMut.Lock()
		originMemberUpsertCache[key] = cache
		originMemberUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OriginMember record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OriginMember) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OriginMember provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), originMemberPrimaryKeyMapping)
	sql := "DELETE FROM \"origin_members\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from origin_members")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for origin_members")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q originMemberQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no originMemberQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from origin_members")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for origin_members")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OriginMemberSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OriginMember slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(originMemberBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), originMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"origin_members\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, originMemberPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from originMember slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for origin_members")
	}

	if len(originMemberAfterDeleteHooks) != 0 {
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
func (o *OriginMember) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOriginMember(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OriginMemberSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OriginMemberSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), originMemberPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"origin_members\".* FROM \"origin_members\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, originMemberPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OriginMemberSlice")
	}

	*o = slice

	return nil
}

// OriginMemberExists checks if the OriginMember row exists.
func OriginMemberExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"origin_members\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if origin_members exists")
	}

	return exists, nil
}