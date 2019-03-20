// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testAccounts(t *testing.T) {
	t.Parallel()

	query := Accounts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAccountsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Accounts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccountSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AccountExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Account exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AccountExists to return true, but got false.")
	}
}

func testAccountsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	accountFound, err := FindAccount(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if accountFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAccountsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Accounts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAccountsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Accounts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAccountsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountOne := &Account{}
	accountTwo := &Account{}
	if err = randomize.Struct(seed, accountOne, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTwo, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = accountOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accountTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Accounts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAccountsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	accountOne := &Account{}
	accountTwo := &Account{}
	if err = randomize.Struct(seed, accountOne, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTwo, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = accountOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = accountTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func accountBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func accountAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Account) error {
	*o = Account{}
	return nil
}

func testAccountsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Account{}
	o := &Account{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, accountDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Account object: %s", err)
	}

	AddAccountHook(boil.BeforeInsertHook, accountBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	accountBeforeInsertHooks = []AccountHook{}

	AddAccountHook(boil.AfterInsertHook, accountAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	accountAfterInsertHooks = []AccountHook{}

	AddAccountHook(boil.AfterSelectHook, accountAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	accountAfterSelectHooks = []AccountHook{}

	AddAccountHook(boil.BeforeUpdateHook, accountBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	accountBeforeUpdateHooks = []AccountHook{}

	AddAccountHook(boil.AfterUpdateHook, accountAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	accountAfterUpdateHooks = []AccountHook{}

	AddAccountHook(boil.BeforeDeleteHook, accountBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	accountBeforeDeleteHooks = []AccountHook{}

	AddAccountHook(boil.AfterDeleteHook, accountAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	accountAfterDeleteHooks = []AccountHook{}

	AddAccountHook(boil.BeforeUpsertHook, accountBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	accountBeforeUpsertHooks = []AccountHook{}

	AddAccountHook(boil.AfterUpsertHook, accountAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	accountAfterUpsertHooks = []AccountHook{}
}

func testAccountsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(accountColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAccountsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AccountSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAccountsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Accounts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	accountDBTypes = map[string]string{`ID`: `bigint`, `Name`: `text`, `Email`: `text`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_              = bytes.MinRead
)

func testAccountsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(accountColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accountDBTypes, true, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAccountsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(accountColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Account{}
	if err = randomize.Struct(seed, o, accountDBTypes, true, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, accountDBTypes, true, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(accountColumns, accountPrimaryKeyColumns) {
		fields = accountColumns
	} else {
		fields = strmangle.SetComplement(
			accountColumns,
			accountPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := AccountSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAccountsUpsert(t *testing.T) {
	t.Parallel()

	if len(accountColumns) == len(accountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Account{}
	if err = randomize.Struct(seed, &o, accountDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Account: %s", err)
	}

	count, err := Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, accountDBTypes, false, accountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Account: %s", err)
	}

	count, err = Accounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
