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

func testOriginSecretKeys(t *testing.T) {
	t.Parallel()

	query := OriginSecretKeys()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOriginSecretKeysDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
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

	count, err := OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginSecretKeysQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OriginSecretKeys().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginSecretKeysSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginSecretKeySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginSecretKeysExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OriginSecretKeyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OriginSecretKey exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OriginSecretKeyExists to return true, but got false.")
	}
}

func testOriginSecretKeysFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	originSecretKeyFound, err := FindOriginSecretKey(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if originSecretKeyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOriginSecretKeysBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OriginSecretKeys().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOriginSecretKeysOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OriginSecretKeys().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOriginSecretKeysAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	originSecretKeyOne := &OriginSecretKey{}
	originSecretKeyTwo := &OriginSecretKey{}
	if err = randomize.Struct(seed, originSecretKeyOne, originSecretKeyDBTypes, false, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}
	if err = randomize.Struct(seed, originSecretKeyTwo, originSecretKeyDBTypes, false, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originSecretKeyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originSecretKeyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginSecretKeys().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOriginSecretKeysCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	originSecretKeyOne := &OriginSecretKey{}
	originSecretKeyTwo := &OriginSecretKey{}
	if err = randomize.Struct(seed, originSecretKeyOne, originSecretKeyDBTypes, false, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}
	if err = randomize.Struct(seed, originSecretKeyTwo, originSecretKeyDBTypes, false, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originSecretKeyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originSecretKeyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func originSecretKeyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginSecretKey) error {
	*o = OriginSecretKey{}
	return nil
}

func originSecretKeyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginSecretKey) error {
	*o = OriginSecretKey{}
	return nil
}

func originSecretKeyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OriginSecretKey) error {
	*o = OriginSecretKey{}
	return nil
}

func originSecretKeyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginSecretKey) error {
	*o = OriginSecretKey{}
	return nil
}

func originSecretKeyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginSecretKey) error {
	*o = OriginSecretKey{}
	return nil
}

func originSecretKeyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginSecretKey) error {
	*o = OriginSecretKey{}
	return nil
}

func originSecretKeyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginSecretKey) error {
	*o = OriginSecretKey{}
	return nil
}

func originSecretKeyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginSecretKey) error {
	*o = OriginSecretKey{}
	return nil
}

func originSecretKeyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginSecretKey) error {
	*o = OriginSecretKey{}
	return nil
}

func testOriginSecretKeysHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OriginSecretKey{}
	o := &OriginSecretKey{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey object: %s", err)
	}

	AddOriginSecretKeyHook(boil.BeforeInsertHook, originSecretKeyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	originSecretKeyBeforeInsertHooks = []OriginSecretKeyHook{}

	AddOriginSecretKeyHook(boil.AfterInsertHook, originSecretKeyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	originSecretKeyAfterInsertHooks = []OriginSecretKeyHook{}

	AddOriginSecretKeyHook(boil.AfterSelectHook, originSecretKeyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	originSecretKeyAfterSelectHooks = []OriginSecretKeyHook{}

	AddOriginSecretKeyHook(boil.BeforeUpdateHook, originSecretKeyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	originSecretKeyBeforeUpdateHooks = []OriginSecretKeyHook{}

	AddOriginSecretKeyHook(boil.AfterUpdateHook, originSecretKeyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	originSecretKeyAfterUpdateHooks = []OriginSecretKeyHook{}

	AddOriginSecretKeyHook(boil.BeforeDeleteHook, originSecretKeyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	originSecretKeyBeforeDeleteHooks = []OriginSecretKeyHook{}

	AddOriginSecretKeyHook(boil.AfterDeleteHook, originSecretKeyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	originSecretKeyAfterDeleteHooks = []OriginSecretKeyHook{}

	AddOriginSecretKeyHook(boil.BeforeUpsertHook, originSecretKeyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	originSecretKeyBeforeUpsertHooks = []OriginSecretKeyHook{}

	AddOriginSecretKeyHook(boil.AfterUpsertHook, originSecretKeyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	originSecretKeyAfterUpsertHooks = []OriginSecretKeyHook{}
}

func testOriginSecretKeysInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginSecretKeysInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(originSecretKeyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginSecretKeyToOneOriginUsingOrigin(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OriginSecretKey
	var foreign Origin

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, originDBTypes, false, originColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Origin struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.OriginName, foreign.Name)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Origin().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.Name, foreign.Name) {
		t.Errorf("want: %v, got %v", foreign.Name, check.Name)
	}

	slice := OriginSecretKeySlice{&local}
	if err = local.L.LoadOrigin(ctx, tx, false, (*[]*OriginSecretKey)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Origin == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Origin = nil
	if err = local.L.LoadOrigin(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Origin == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOriginSecretKeyToOneSetOpOriginUsingOrigin(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginSecretKey
	var b, c Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originSecretKeyDBTypes, false, strmangle.SetComplement(originSecretKeyPrimaryKeyColumns, originSecretKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, originDBTypes, false, strmangle.SetComplement(originPrimaryKeyColumns, originColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, originDBTypes, false, strmangle.SetComplement(originPrimaryKeyColumns, originColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Origin{&b, &c} {
		err = a.SetOrigin(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Origin != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OriginOSK[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.OriginName, x.Name) {
			t.Error("foreign key was wrong value", a.OriginName)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OriginName))
		reflect.Indirect(reflect.ValueOf(&a.OriginName)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.OriginName, x.Name) {
			t.Error("foreign key was wrong value", a.OriginName, x.Name)
		}
	}
}

func testOriginSecretKeyToOneRemoveOpOriginUsingOrigin(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginSecretKey
	var b Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originSecretKeyDBTypes, false, strmangle.SetComplement(originSecretKeyPrimaryKeyColumns, originSecretKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, originDBTypes, false, strmangle.SetComplement(originPrimaryKeyColumns, originColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetOrigin(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveOrigin(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Origin().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Origin != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.OriginName) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.OriginOSK) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOriginSecretKeysReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
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

func testOriginSecretKeysReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginSecretKeySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOriginSecretKeysSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginSecretKeys().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	originSecretKeyDBTypes = map[string]string{`ID`: `bigint`, `OwnerID`: `bigint`, `Name`: `text`, `Revision`: `text`, `FullName`: `text`, `Body`: `bytea`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `OriginName`: `text`}
	_                      = bytes.MinRead
)

func testOriginSecretKeysUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(originSecretKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(originSecretKeyColumns) == len(originSecretKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOriginSecretKeysSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(originSecretKeyColumns) == len(originSecretKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginSecretKey{}
	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originSecretKeyDBTypes, true, originSecretKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(originSecretKeyColumns, originSecretKeyPrimaryKeyColumns) {
		fields = originSecretKeyColumns
	} else {
		fields = strmangle.SetComplement(
			originSecretKeyColumns,
			originSecretKeyPrimaryKeyColumns,
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

	slice := OriginSecretKeySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOriginSecretKeysUpsert(t *testing.T) {
	t.Parallel()

	if len(originSecretKeyColumns) == len(originSecretKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OriginSecretKey{}
	if err = randomize.Struct(seed, &o, originSecretKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginSecretKey: %s", err)
	}

	count, err := OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, originSecretKeyDBTypes, false, originSecretKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginSecretKey struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginSecretKey: %s", err)
	}

	count, err = OriginSecretKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
