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

func testOriginPublicEncryptionKeys(t *testing.T) {
	t.Parallel()

	query := OriginPublicEncryptionKeys()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOriginPublicEncryptionKeysDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
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

	count, err := OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginPublicEncryptionKeysQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OriginPublicEncryptionKeys().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginPublicEncryptionKeysSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginPublicEncryptionKeySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginPublicEncryptionKeysExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OriginPublicEncryptionKeyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OriginPublicEncryptionKey exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OriginPublicEncryptionKeyExists to return true, but got false.")
	}
}

func testOriginPublicEncryptionKeysFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	originPublicEncryptionKeyFound, err := FindOriginPublicEncryptionKey(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if originPublicEncryptionKeyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOriginPublicEncryptionKeysBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OriginPublicEncryptionKeys().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOriginPublicEncryptionKeysOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OriginPublicEncryptionKeys().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOriginPublicEncryptionKeysAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	originPublicEncryptionKeyOne := &OriginPublicEncryptionKey{}
	originPublicEncryptionKeyTwo := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, originPublicEncryptionKeyOne, originPublicEncryptionKeyDBTypes, false, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}
	if err = randomize.Struct(seed, originPublicEncryptionKeyTwo, originPublicEncryptionKeyDBTypes, false, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originPublicEncryptionKeyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originPublicEncryptionKeyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginPublicEncryptionKeys().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOriginPublicEncryptionKeysCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	originPublicEncryptionKeyOne := &OriginPublicEncryptionKey{}
	originPublicEncryptionKeyTwo := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, originPublicEncryptionKeyOne, originPublicEncryptionKeyDBTypes, false, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}
	if err = randomize.Struct(seed, originPublicEncryptionKeyTwo, originPublicEncryptionKeyDBTypes, false, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originPublicEncryptionKeyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originPublicEncryptionKeyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func originPublicEncryptionKeyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPublicEncryptionKey) error {
	*o = OriginPublicEncryptionKey{}
	return nil
}

func originPublicEncryptionKeyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPublicEncryptionKey) error {
	*o = OriginPublicEncryptionKey{}
	return nil
}

func originPublicEncryptionKeyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OriginPublicEncryptionKey) error {
	*o = OriginPublicEncryptionKey{}
	return nil
}

func originPublicEncryptionKeyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginPublicEncryptionKey) error {
	*o = OriginPublicEncryptionKey{}
	return nil
}

func originPublicEncryptionKeyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginPublicEncryptionKey) error {
	*o = OriginPublicEncryptionKey{}
	return nil
}

func originPublicEncryptionKeyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginPublicEncryptionKey) error {
	*o = OriginPublicEncryptionKey{}
	return nil
}

func originPublicEncryptionKeyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginPublicEncryptionKey) error {
	*o = OriginPublicEncryptionKey{}
	return nil
}

func originPublicEncryptionKeyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPublicEncryptionKey) error {
	*o = OriginPublicEncryptionKey{}
	return nil
}

func originPublicEncryptionKeyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPublicEncryptionKey) error {
	*o = OriginPublicEncryptionKey{}
	return nil
}

func testOriginPublicEncryptionKeysHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OriginPublicEncryptionKey{}
	o := &OriginPublicEncryptionKey{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey object: %s", err)
	}

	AddOriginPublicEncryptionKeyHook(boil.BeforeInsertHook, originPublicEncryptionKeyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	originPublicEncryptionKeyBeforeInsertHooks = []OriginPublicEncryptionKeyHook{}

	AddOriginPublicEncryptionKeyHook(boil.AfterInsertHook, originPublicEncryptionKeyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	originPublicEncryptionKeyAfterInsertHooks = []OriginPublicEncryptionKeyHook{}

	AddOriginPublicEncryptionKeyHook(boil.AfterSelectHook, originPublicEncryptionKeyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	originPublicEncryptionKeyAfterSelectHooks = []OriginPublicEncryptionKeyHook{}

	AddOriginPublicEncryptionKeyHook(boil.BeforeUpdateHook, originPublicEncryptionKeyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	originPublicEncryptionKeyBeforeUpdateHooks = []OriginPublicEncryptionKeyHook{}

	AddOriginPublicEncryptionKeyHook(boil.AfterUpdateHook, originPublicEncryptionKeyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	originPublicEncryptionKeyAfterUpdateHooks = []OriginPublicEncryptionKeyHook{}

	AddOriginPublicEncryptionKeyHook(boil.BeforeDeleteHook, originPublicEncryptionKeyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	originPublicEncryptionKeyBeforeDeleteHooks = []OriginPublicEncryptionKeyHook{}

	AddOriginPublicEncryptionKeyHook(boil.AfterDeleteHook, originPublicEncryptionKeyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	originPublicEncryptionKeyAfterDeleteHooks = []OriginPublicEncryptionKeyHook{}

	AddOriginPublicEncryptionKeyHook(boil.BeforeUpsertHook, originPublicEncryptionKeyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	originPublicEncryptionKeyBeforeUpsertHooks = []OriginPublicEncryptionKeyHook{}

	AddOriginPublicEncryptionKeyHook(boil.AfterUpsertHook, originPublicEncryptionKeyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	originPublicEncryptionKeyAfterUpsertHooks = []OriginPublicEncryptionKeyHook{}
}

func testOriginPublicEncryptionKeysInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginPublicEncryptionKeysInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(originPublicEncryptionKeyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginPublicEncryptionKeyToOneOriginUsingOriginName(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OriginPublicEncryptionKey
	var foreign Origin

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, originDBTypes, false, originColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Origin struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.Origin, foreign.Name)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.OriginName().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.Name, foreign.Name) {
		t.Errorf("want: %v, got %v", foreign.Name, check.Name)
	}

	slice := OriginPublicEncryptionKeySlice{&local}
	if err = local.L.LoadOriginName(ctx, tx, false, (*[]*OriginPublicEncryptionKey)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OriginName == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OriginName = nil
	if err = local.L.LoadOriginName(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OriginName == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOriginPublicEncryptionKeyToOneSetOpOriginUsingOriginName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginPublicEncryptionKey
	var b, c Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originPublicEncryptionKeyDBTypes, false, strmangle.SetComplement(originPublicEncryptionKeyPrimaryKeyColumns, originPublicEncryptionKeyColumnsWithoutDefault)...); err != nil {
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
		err = a.SetOriginName(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OriginName != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OriginOPUEK[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.Origin, x.Name) {
			t.Error("foreign key was wrong value", a.Origin)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Origin))
		reflect.Indirect(reflect.ValueOf(&a.Origin)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.Origin, x.Name) {
			t.Error("foreign key was wrong value", a.Origin, x.Name)
		}
	}
}

func testOriginPublicEncryptionKeyToOneRemoveOpOriginUsingOriginName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginPublicEncryptionKey
	var b Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originPublicEncryptionKeyDBTypes, false, strmangle.SetComplement(originPublicEncryptionKeyPrimaryKeyColumns, originPublicEncryptionKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, originDBTypes, false, strmangle.SetComplement(originPrimaryKeyColumns, originColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetOriginName(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveOriginName(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.OriginName().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.OriginName != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.Origin) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.OriginOPUEK) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOriginPublicEncryptionKeysReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
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

func testOriginPublicEncryptionKeysReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginPublicEncryptionKeySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOriginPublicEncryptionKeysSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginPublicEncryptionKeys().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	originPublicEncryptionKeyDBTypes = map[string]string{`ID`: `bigint`, `OwnerID`: `bigint`, `Name`: `text`, `Revision`: `text`, `FullName`: `text`, `Body`: `bytea`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `Origin`: `text`}
	_                                = bytes.MinRead
)

func testOriginPublicEncryptionKeysUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(originPublicEncryptionKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(originPublicEncryptionKeyColumns) == len(originPublicEncryptionKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOriginPublicEncryptionKeysSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(originPublicEncryptionKeyColumns) == len(originPublicEncryptionKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originPublicEncryptionKeyDBTypes, true, originPublicEncryptionKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(originPublicEncryptionKeyColumns, originPublicEncryptionKeyPrimaryKeyColumns) {
		fields = originPublicEncryptionKeyColumns
	} else {
		fields = strmangle.SetComplement(
			originPublicEncryptionKeyColumns,
			originPublicEncryptionKeyPrimaryKeyColumns,
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

	slice := OriginPublicEncryptionKeySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOriginPublicEncryptionKeysUpsert(t *testing.T) {
	t.Parallel()

	if len(originPublicEncryptionKeyColumns) == len(originPublicEncryptionKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OriginPublicEncryptionKey{}
	if err = randomize.Struct(seed, &o, originPublicEncryptionKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginPublicEncryptionKey: %s", err)
	}

	count, err := OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, originPublicEncryptionKeyDBTypes, false, originPublicEncryptionKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginPublicEncryptionKey struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginPublicEncryptionKey: %s", err)
	}

	count, err = OriginPublicEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
