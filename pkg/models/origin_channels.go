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

// OriginChannel is an object representing the database table.
type OriginChannel struct {
	ID         int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	OwnerID    null.Int64  `boil:"owner_id" json:"owner_id,omitempty" toml:"owner_id" yaml:"owner_id,omitempty"`
	Name       null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	CreatedAt  null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt  null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	OriginName null.String `boil:"origin_name" json:"origin_name,omitempty" toml:"origin_name" yaml:"origin_name,omitempty"`

	R *originChannelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L originChannelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OriginChannelColumns = struct {
	ID         string
	OwnerID    string
	Name       string
	CreatedAt  string
	UpdatedAt  string
	OriginName string
}{
	ID:         "id",
	OwnerID:    "owner_id",
	Name:       "name",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	OriginName: "origin_name",
}

// Generated where

var OriginChannelWhere = struct {
	ID         whereHelperint64
	OwnerID    whereHelpernull_Int64
	Name       whereHelpernull_String
	CreatedAt  whereHelpernull_Time
	UpdatedAt  whereHelpernull_Time
	OriginName whereHelpernull_String
}{
	ID:         whereHelperint64{field: `id`},
	OwnerID:    whereHelpernull_Int64{field: `owner_id`},
	Name:       whereHelpernull_String{field: `name`},
	CreatedAt:  whereHelpernull_Time{field: `created_at`},
	UpdatedAt:  whereHelpernull_Time{field: `updated_at`},
	OriginName: whereHelpernull_String{field: `origin_name`},
}

// OriginChannelRels is where relationship names are stored.
var OriginChannelRels = struct {
	Origin                       string
	ChannelOriginChannelPackages string
}{
	Origin:                       "Origin",
	ChannelOriginChannelPackages: "ChannelOriginChannelPackages",
}

// originChannelR is where relationships are stored.
type originChannelR struct {
	Origin                       *Origin
	ChannelOriginChannelPackages OriginChannelPackageSlice
}

// NewStruct creates a new relationship struct
func (*originChannelR) NewStruct() *originChannelR {
	return &originChannelR{}
}

// originChannelL is where Load methods for each relationship are stored.
type originChannelL struct{}

var (
	originChannelColumns               = []string{"id", "owner_id", "name", "created_at", "updated_at", "origin_name"}
	originChannelColumnsWithoutDefault = []string{"owner_id", "name", "origin_name"}
	originChannelColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	originChannelPrimaryKeyColumns     = []string{"id"}
)

type (
	// OriginChannelSlice is an alias for a slice of pointers to OriginChannel.
	// This should generally be used opposed to []OriginChannel.
	OriginChannelSlice []*OriginChannel
	// OriginChannelHook is the signature for custom OriginChannel hook methods
	OriginChannelHook func(context.Context, boil.ContextExecutor, *OriginChannel) error

	originChannelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	originChannelType                 = reflect.TypeOf(&OriginChannel{})
	originChannelMapping              = queries.MakeStructMapping(originChannelType)
	originChannelPrimaryKeyMapping, _ = queries.BindMapping(originChannelType, originChannelMapping, originChannelPrimaryKeyColumns)
	originChannelInsertCacheMut       sync.RWMutex
	originChannelInsertCache          = make(map[string]insertCache)
	originChannelUpdateCacheMut       sync.RWMutex
	originChannelUpdateCache          = make(map[string]updateCache)
	originChannelUpsertCacheMut       sync.RWMutex
	originChannelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var originChannelBeforeInsertHooks []OriginChannelHook
var originChannelBeforeUpdateHooks []OriginChannelHook
var originChannelBeforeDeleteHooks []OriginChannelHook
var originChannelBeforeUpsertHooks []OriginChannelHook

var originChannelAfterInsertHooks []OriginChannelHook
var originChannelAfterSelectHooks []OriginChannelHook
var originChannelAfterUpdateHooks []OriginChannelHook
var originChannelAfterDeleteHooks []OriginChannelHook
var originChannelAfterUpsertHooks []OriginChannelHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OriginChannel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originChannelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OriginChannel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originChannelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OriginChannel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originChannelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OriginChannel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originChannelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OriginChannel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originChannelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OriginChannel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originChannelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OriginChannel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originChannelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OriginChannel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originChannelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OriginChannel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range originChannelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOriginChannelHook registers your hook function for all future operations.
func AddOriginChannelHook(hookPoint boil.HookPoint, originChannelHook OriginChannelHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		originChannelBeforeInsertHooks = append(originChannelBeforeInsertHooks, originChannelHook)
	case boil.BeforeUpdateHook:
		originChannelBeforeUpdateHooks = append(originChannelBeforeUpdateHooks, originChannelHook)
	case boil.BeforeDeleteHook:
		originChannelBeforeDeleteHooks = append(originChannelBeforeDeleteHooks, originChannelHook)
	case boil.BeforeUpsertHook:
		originChannelBeforeUpsertHooks = append(originChannelBeforeUpsertHooks, originChannelHook)
	case boil.AfterInsertHook:
		originChannelAfterInsertHooks = append(originChannelAfterInsertHooks, originChannelHook)
	case boil.AfterSelectHook:
		originChannelAfterSelectHooks = append(originChannelAfterSelectHooks, originChannelHook)
	case boil.AfterUpdateHook:
		originChannelAfterUpdateHooks = append(originChannelAfterUpdateHooks, originChannelHook)
	case boil.AfterDeleteHook:
		originChannelAfterDeleteHooks = append(originChannelAfterDeleteHooks, originChannelHook)
	case boil.AfterUpsertHook:
		originChannelAfterUpsertHooks = append(originChannelAfterUpsertHooks, originChannelHook)
	}
}

// One returns a single originChannel record from the query.
func (q originChannelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OriginChannel, error) {
	o := &OriginChannel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for origin_channels")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OriginChannel records from the query.
func (q originChannelQuery) All(ctx context.Context, exec boil.ContextExecutor) (OriginChannelSlice, error) {
	var o []*OriginChannel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OriginChannel slice")
	}

	if len(originChannelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OriginChannel records in the query.
func (q originChannelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count origin_channels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q originChannelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if origin_channels exists")
	}

	return count > 0, nil
}

// Origin pointed to by the foreign key.
func (o *OriginChannel) Origin(mods ...qm.QueryMod) originQuery {
	queryMods := []qm.QueryMod{
		qm.Where("name=?", o.OriginName),
	}

	queryMods = append(queryMods, mods...)

	query := Origins(queryMods...)
	queries.SetFrom(query.Query, "\"origins\"")

	return query
}

// ChannelOriginChannelPackages retrieves all the origin_channel_package's OriginChannelPackages with an executor via channel_id column.
func (o *OriginChannel) ChannelOriginChannelPackages(mods ...qm.QueryMod) originChannelPackageQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"origin_channel_packages\".\"channel_id\"=?", o.ID),
	)

	query := OriginChannelPackages(queryMods...)
	queries.SetFrom(query.Query, "\"origin_channel_packages\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"origin_channel_packages\".*"})
	}

	return query
}

// LoadOrigin allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (originChannelL) LoadOrigin(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOriginChannel interface{}, mods queries.Applicator) error {
	var slice []*OriginChannel
	var object *OriginChannel

	if singular {
		object = maybeOriginChannel.(*OriginChannel)
	} else {
		slice = *maybeOriginChannel.(*[]*OriginChannel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &originChannelR{}
		}
		if !queries.IsNil(object.OriginName) {
			args = append(args, object.OriginName)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &originChannelR{}
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

	if len(originChannelAfterSelectHooks) != 0 {
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
		foreign.R.OriginOC = append(foreign.R.OriginOC, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.OriginName, foreign.Name) {
				local.R.Origin = foreign
				if foreign.R == nil {
					foreign.R = &originR{}
				}
				foreign.R.OriginOC = append(foreign.R.OriginOC, local)
				break
			}
		}
	}

	return nil
}

// LoadChannelOriginChannelPackages allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (originChannelL) LoadChannelOriginChannelPackages(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOriginChannel interface{}, mods queries.Applicator) error {
	var slice []*OriginChannel
	var object *OriginChannel

	if singular {
		object = maybeOriginChannel.(*OriginChannel)
	} else {
		slice = *maybeOriginChannel.(*[]*OriginChannel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &originChannelR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &originChannelR{}
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

	query := NewQuery(qm.From(`origin_channel_packages`), qm.WhereIn(`channel_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load origin_channel_packages")
	}

	var resultSlice []*OriginChannelPackage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice origin_channel_packages")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on origin_channel_packages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for origin_channel_packages")
	}

	if len(originChannelPackageAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ChannelOriginChannelPackages = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &originChannelPackageR{}
			}
			foreign.R.Channel = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChannelID {
				local.R.ChannelOriginChannelPackages = append(local.R.ChannelOriginChannelPackages, foreign)
				if foreign.R == nil {
					foreign.R = &originChannelPackageR{}
				}
				foreign.R.Channel = local
				break
			}
		}
	}

	return nil
}

// SetOrigin of the originChannel to the related item.
// Sets o.R.Origin to related.
// Adds o to related.R.OriginOC.
func (o *OriginChannel) SetOrigin(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Origin) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"origin_channels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"origin_name"}),
		strmangle.WhereClause("\"", "\"", 2, originChannelPrimaryKeyColumns),
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
		o.R = &originChannelR{
			Origin: related,
		}
	} else {
		o.R.Origin = related
	}

	if related.R == nil {
		related.R = &originR{
			OriginOC: OriginChannelSlice{o},
		}
	} else {
		related.R.OriginOC = append(related.R.OriginOC, o)
	}

	return nil
}

// RemoveOrigin relationship.
// Sets o.R.Origin to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *OriginChannel) RemoveOrigin(ctx context.Context, exec boil.ContextExecutor, related *Origin) error {
	var err error

	queries.SetScanner(&o.OriginName, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("origin_name")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Origin = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.OriginOC {
		if queries.Equal(o.OriginName, ri.OriginName) {
			continue
		}

		ln := len(related.R.OriginOC)
		if ln > 1 && i < ln-1 {
			related.R.OriginOC[i] = related.R.OriginOC[ln-1]
		}
		related.R.OriginOC = related.R.OriginOC[:ln-1]
		break
	}
	return nil
}

// AddChannelOriginChannelPackages adds the given related objects to the existing relationships
// of the origin_channel, optionally inserting them as new records.
// Appends related to o.R.ChannelOriginChannelPackages.
// Sets related.R.Channel appropriately.
func (o *OriginChannel) AddChannelOriginChannelPackages(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*OriginChannelPackage) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChannelID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"origin_channel_packages\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"channel_id"}),
				strmangle.WhereClause("\"", "\"", 2, originChannelPackagePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ChannelID, rel.PackageID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ChannelID = o.ID
		}
	}

	if o.R == nil {
		o.R = &originChannelR{
			ChannelOriginChannelPackages: related,
		}
	} else {
		o.R.ChannelOriginChannelPackages = append(o.R.ChannelOriginChannelPackages, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &originChannelPackageR{
				Channel: o,
			}
		} else {
			rel.R.Channel = o
		}
	}
	return nil
}

// OriginChannels retrieves all the records using an executor.
func OriginChannels(mods ...qm.QueryMod) originChannelQuery {
	mods = append(mods, qm.From("\"origin_channels\""))
	return originChannelQuery{NewQuery(mods...)}
}

// FindOriginChannel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOriginChannel(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*OriginChannel, error) {
	originChannelObj := &OriginChannel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"origin_channels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, originChannelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from origin_channels")
	}

	return originChannelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OriginChannel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no origin_channels provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(originChannelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	originChannelInsertCacheMut.RLock()
	cache, cached := originChannelInsertCache[key]
	originChannelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			originChannelColumns,
			originChannelColumnsWithDefault,
			originChannelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(originChannelType, originChannelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(originChannelType, originChannelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"origin_channels\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"origin_channels\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into origin_channels")
	}

	if !cached {
		originChannelInsertCacheMut.Lock()
		originChannelInsertCache[key] = cache
		originChannelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OriginChannel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OriginChannel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	originChannelUpdateCacheMut.RLock()
	cache, cached := originChannelUpdateCache[key]
	originChannelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			originChannelColumns,
			originChannelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update origin_channels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"origin_channels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, originChannelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(originChannelType, originChannelMapping, append(wl, originChannelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update origin_channels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for origin_channels")
	}

	if !cached {
		originChannelUpdateCacheMut.Lock()
		originChannelUpdateCache[key] = cache
		originChannelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q originChannelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for origin_channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for origin_channels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OriginChannelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), originChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"origin_channels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, originChannelPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in originChannel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all originChannel")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OriginChannel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no origin_channels provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(originChannelColumnsWithDefault, o)

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

	originChannelUpsertCacheMut.RLock()
	cache, cached := originChannelUpsertCache[key]
	originChannelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			originChannelColumns,
			originChannelColumnsWithDefault,
			originChannelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			originChannelColumns,
			originChannelPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert origin_channels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(originChannelPrimaryKeyColumns))
			copy(conflict, originChannelPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"origin_channels\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(originChannelType, originChannelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(originChannelType, originChannelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert origin_channels")
	}

	if !cached {
		originChannelUpsertCacheMut.Lock()
		originChannelUpsertCache[key] = cache
		originChannelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OriginChannel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OriginChannel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OriginChannel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), originChannelPrimaryKeyMapping)
	sql := "DELETE FROM \"origin_channels\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from origin_channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for origin_channels")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q originChannelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no originChannelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from origin_channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for origin_channels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OriginChannelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OriginChannel slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(originChannelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), originChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"origin_channels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, originChannelPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from originChannel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for origin_channels")
	}

	if len(originChannelAfterDeleteHooks) != 0 {
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
func (o *OriginChannel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOriginChannel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OriginChannelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OriginChannelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), originChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"origin_channels\".* FROM \"origin_channels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, originChannelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OriginChannelSlice")
	}

	*o = slice

	return nil
}

// OriginChannelExists checks if the OriginChannel row exists.
func OriginChannelExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"origin_channels\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if origin_channels exists")
	}

	return exists, nil
}
