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

// AuditJob is an object representing the database table.
type AuditJob struct {
	ID            int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GroupID       null.Int64  `boil:"group_id" json:"group_id,omitempty" toml:"group_id" yaml:"group_id,omitempty"`
	Operation     null.Int16  `boil:"operation" json:"operation,omitempty" toml:"operation" yaml:"operation,omitempty"`
	Trigger       null.Int16  `boil:"trigger" json:"trigger,omitempty" toml:"trigger" yaml:"trigger,omitempty"`
	RequesterID   null.Int64  `boil:"requester_id" json:"requester_id,omitempty" toml:"requester_id" yaml:"requester_id,omitempty"`
	RequesterName null.String `boil:"requester_name" json:"requester_name,omitempty" toml:"requester_name" yaml:"requester_name,omitempty"`
	CreatedAt     null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`

	R *auditJobR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L auditJobL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuditJobColumns = struct {
	ID            string
	GroupID       string
	Operation     string
	Trigger       string
	RequesterID   string
	RequesterName string
	CreatedAt     string
}{
	ID:            "id",
	GroupID:       "group_id",
	Operation:     "operation",
	Trigger:       "trigger",
	RequesterID:   "requester_id",
	RequesterName: "requester_name",
	CreatedAt:     "created_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_Int16 struct{ field string }

func (w whereHelpernull_Int16) EQ(x null.Int16) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int16) NEQ(x null.Int16) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int16) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int16) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int16) LT(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int16) LTE(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int16) GT(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int16) GTE(x null.Int16) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AuditJobWhere = struct {
	ID            whereHelperint
	GroupID       whereHelpernull_Int64
	Operation     whereHelpernull_Int16
	Trigger       whereHelpernull_Int16
	RequesterID   whereHelpernull_Int64
	RequesterName whereHelpernull_String
	CreatedAt     whereHelpernull_Time
}{
	ID:            whereHelperint{field: `id`},
	GroupID:       whereHelpernull_Int64{field: `group_id`},
	Operation:     whereHelpernull_Int16{field: `operation`},
	Trigger:       whereHelpernull_Int16{field: `trigger`},
	RequesterID:   whereHelpernull_Int64{field: `requester_id`},
	RequesterName: whereHelpernull_String{field: `requester_name`},
	CreatedAt:     whereHelpernull_Time{field: `created_at`},
}

// AuditJobRels is where relationship names are stored.
var AuditJobRels = struct {
}{}

// auditJobR is where relationships are stored.
type auditJobR struct {
}

// NewStruct creates a new relationship struct
func (*auditJobR) NewStruct() *auditJobR {
	return &auditJobR{}
}

// auditJobL is where Load methods for each relationship are stored.
type auditJobL struct{}

var (
	auditJobColumns               = []string{"id", "group_id", "operation", "trigger", "requester_id", "requester_name", "created_at"}
	auditJobColumnsWithoutDefault = []string{"group_id", "operation", "trigger", "requester_id", "requester_name"}
	auditJobColumnsWithDefault    = []string{"id", "created_at"}
	auditJobPrimaryKeyColumns     = []string{"id"}
)

type (
	// AuditJobSlice is an alias for a slice of pointers to AuditJob.
	// This should generally be used opposed to []AuditJob.
	AuditJobSlice []*AuditJob
	// AuditJobHook is the signature for custom AuditJob hook methods
	AuditJobHook func(context.Context, boil.ContextExecutor, *AuditJob) error

	auditJobQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	auditJobType                 = reflect.TypeOf(&AuditJob{})
	auditJobMapping              = queries.MakeStructMapping(auditJobType)
	auditJobPrimaryKeyMapping, _ = queries.BindMapping(auditJobType, auditJobMapping, auditJobPrimaryKeyColumns)
	auditJobInsertCacheMut       sync.RWMutex
	auditJobInsertCache          = make(map[string]insertCache)
	auditJobUpdateCacheMut       sync.RWMutex
	auditJobUpdateCache          = make(map[string]updateCache)
	auditJobUpsertCacheMut       sync.RWMutex
	auditJobUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var auditJobBeforeInsertHooks []AuditJobHook
var auditJobBeforeUpdateHooks []AuditJobHook
var auditJobBeforeDeleteHooks []AuditJobHook
var auditJobBeforeUpsertHooks []AuditJobHook

var auditJobAfterInsertHooks []AuditJobHook
var auditJobAfterSelectHooks []AuditJobHook
var auditJobAfterUpdateHooks []AuditJobHook
var auditJobAfterDeleteHooks []AuditJobHook
var auditJobAfterUpsertHooks []AuditJobHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuditJob) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditJobBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuditJob) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditJobBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuditJob) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditJobBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuditJob) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditJobBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuditJob) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditJobAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuditJob) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditJobAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuditJob) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditJobAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuditJob) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditJobAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuditJob) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range auditJobAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuditJobHook registers your hook function for all future operations.
func AddAuditJobHook(hookPoint boil.HookPoint, auditJobHook AuditJobHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		auditJobBeforeInsertHooks = append(auditJobBeforeInsertHooks, auditJobHook)
	case boil.BeforeUpdateHook:
		auditJobBeforeUpdateHooks = append(auditJobBeforeUpdateHooks, auditJobHook)
	case boil.BeforeDeleteHook:
		auditJobBeforeDeleteHooks = append(auditJobBeforeDeleteHooks, auditJobHook)
	case boil.BeforeUpsertHook:
		auditJobBeforeUpsertHooks = append(auditJobBeforeUpsertHooks, auditJobHook)
	case boil.AfterInsertHook:
		auditJobAfterInsertHooks = append(auditJobAfterInsertHooks, auditJobHook)
	case boil.AfterSelectHook:
		auditJobAfterSelectHooks = append(auditJobAfterSelectHooks, auditJobHook)
	case boil.AfterUpdateHook:
		auditJobAfterUpdateHooks = append(auditJobAfterUpdateHooks, auditJobHook)
	case boil.AfterDeleteHook:
		auditJobAfterDeleteHooks = append(auditJobAfterDeleteHooks, auditJobHook)
	case boil.AfterUpsertHook:
		auditJobAfterUpsertHooks = append(auditJobAfterUpsertHooks, auditJobHook)
	}
}

// One returns a single auditJob record from the query.
func (q auditJobQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AuditJob, error) {
	o := &AuditJob{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for audit_jobs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AuditJob records from the query.
func (q auditJobQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuditJobSlice, error) {
	var o []*AuditJob

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuditJob slice")
	}

	if len(auditJobAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AuditJob records in the query.
func (q auditJobQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count audit_jobs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q auditJobQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if audit_jobs exists")
	}

	return count > 0, nil
}

// AuditJobs retrieves all the records using an executor.
func AuditJobs(mods ...qm.QueryMod) auditJobQuery {
	mods = append(mods, qm.From("\"audit_jobs\""))
	return auditJobQuery{NewQuery(mods...)}
}

// FindAuditJob retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuditJob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*AuditJob, error) {
	auditJobObj := &AuditJob{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"audit_jobs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, auditJobObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from audit_jobs")
	}

	return auditJobObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AuditJob) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no audit_jobs provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(auditJobColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	auditJobInsertCacheMut.RLock()
	cache, cached := auditJobInsertCache[key]
	auditJobInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			auditJobColumns,
			auditJobColumnsWithDefault,
			auditJobColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(auditJobType, auditJobMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(auditJobType, auditJobMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"audit_jobs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"audit_jobs\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into audit_jobs")
	}

	if !cached {
		auditJobInsertCacheMut.Lock()
		auditJobInsertCache[key] = cache
		auditJobInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AuditJob.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AuditJob) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	auditJobUpdateCacheMut.RLock()
	cache, cached := auditJobUpdateCache[key]
	auditJobUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			auditJobColumns,
			auditJobPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update audit_jobs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"audit_jobs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, auditJobPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(auditJobType, auditJobMapping, append(wl, auditJobPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update audit_jobs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for audit_jobs")
	}

	if !cached {
		auditJobUpdateCacheMut.Lock()
		auditJobUpdateCache[key] = cache
		auditJobUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q auditJobQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for audit_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for audit_jobs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuditJobSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), auditJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"audit_jobs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, auditJobPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in auditJob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all auditJob")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AuditJob) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no audit_jobs provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(auditJobColumnsWithDefault, o)

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

	auditJobUpsertCacheMut.RLock()
	cache, cached := auditJobUpsertCache[key]
	auditJobUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			auditJobColumns,
			auditJobColumnsWithDefault,
			auditJobColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			auditJobColumns,
			auditJobPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert audit_jobs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(auditJobPrimaryKeyColumns))
			copy(conflict, auditJobPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"audit_jobs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(auditJobType, auditJobMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(auditJobType, auditJobMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert audit_jobs")
	}

	if !cached {
		auditJobUpsertCacheMut.Lock()
		auditJobUpsertCache[key] = cache
		auditJobUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AuditJob record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuditJob) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AuditJob provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), auditJobPrimaryKeyMapping)
	sql := "DELETE FROM \"audit_jobs\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from audit_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for audit_jobs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q auditJobQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no auditJobQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from audit_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for audit_jobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuditJobSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AuditJob slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(auditJobBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), auditJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"audit_jobs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, auditJobPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from auditJob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for audit_jobs")
	}

	if len(auditJobAfterDeleteHooks) != 0 {
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
func (o *AuditJob) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuditJob(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuditJobSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuditJobSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), auditJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"audit_jobs\".* FROM \"audit_jobs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, auditJobPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuditJobSlice")
	}

	*o = slice

	return nil
}

// AuditJobExists checks if the AuditJob row exists.
func AuditJobExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"audit_jobs\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if audit_jobs exists")
	}

	return exists, nil
}
