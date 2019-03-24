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

func testOriginPrivateEncryptionKeys(t *testing.T) {
	t.Parallel()

	query := OriginPrivateEncryptionKeys()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOriginPrivateEncryptionKeysDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
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

	count, err := OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginPrivateEncryptionKeysQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OriginPrivateEncryptionKeys().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginPrivateEncryptionKeysSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginPrivateEncryptionKeySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginPrivateEncryptionKeysExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OriginPrivateEncryptionKeyExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OriginPrivateEncryptionKey exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OriginPrivateEncryptionKeyExists to return true, but got false.")
	}
}

func testOriginPrivateEncryptionKeysFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	originPrivateEncryptionKeyFound, err := FindOriginPrivateEncryptionKey(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if originPrivateEncryptionKeyFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOriginPrivateEncryptionKeysBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OriginPrivateEncryptionKeys().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOriginPrivateEncryptionKeysOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OriginPrivateEncryptionKeys().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOriginPrivateEncryptionKeysAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	originPrivateEncryptionKeyOne := &OriginPrivateEncryptionKey{}
	originPrivateEncryptionKeyTwo := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, originPrivateEncryptionKeyOne, originPrivateEncryptionKeyDBTypes, false, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}
	if err = randomize.Struct(seed, originPrivateEncryptionKeyTwo, originPrivateEncryptionKeyDBTypes, false, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originPrivateEncryptionKeyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originPrivateEncryptionKeyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginPrivateEncryptionKeys().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOriginPrivateEncryptionKeysCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	originPrivateEncryptionKeyOne := &OriginPrivateEncryptionKey{}
	originPrivateEncryptionKeyTwo := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, originPrivateEncryptionKeyOne, originPrivateEncryptionKeyDBTypes, false, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}
	if err = randomize.Struct(seed, originPrivateEncryptionKeyTwo, originPrivateEncryptionKeyDBTypes, false, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originPrivateEncryptionKeyOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originPrivateEncryptionKeyTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func originPrivateEncryptionKeyBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPrivateEncryptionKey) error {
	*o = OriginPrivateEncryptionKey{}
	return nil
}

func originPrivateEncryptionKeyAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPrivateEncryptionKey) error {
	*o = OriginPrivateEncryptionKey{}
	return nil
}

func originPrivateEncryptionKeyAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OriginPrivateEncryptionKey) error {
	*o = OriginPrivateEncryptionKey{}
	return nil
}

func originPrivateEncryptionKeyBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginPrivateEncryptionKey) error {
	*o = OriginPrivateEncryptionKey{}
	return nil
}

func originPrivateEncryptionKeyAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginPrivateEncryptionKey) error {
	*o = OriginPrivateEncryptionKey{}
	return nil
}

func originPrivateEncryptionKeyBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginPrivateEncryptionKey) error {
	*o = OriginPrivateEncryptionKey{}
	return nil
}

func originPrivateEncryptionKeyAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginPrivateEncryptionKey) error {
	*o = OriginPrivateEncryptionKey{}
	return nil
}

func originPrivateEncryptionKeyBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPrivateEncryptionKey) error {
	*o = OriginPrivateEncryptionKey{}
	return nil
}

func originPrivateEncryptionKeyAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPrivateEncryptionKey) error {
	*o = OriginPrivateEncryptionKey{}
	return nil
}

func testOriginPrivateEncryptionKeysHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OriginPrivateEncryptionKey{}
	o := &OriginPrivateEncryptionKey{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey object: %s", err)
	}

	AddOriginPrivateEncryptionKeyHook(boil.BeforeInsertHook, originPrivateEncryptionKeyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	originPrivateEncryptionKeyBeforeInsertHooks = []OriginPrivateEncryptionKeyHook{}

	AddOriginPrivateEncryptionKeyHook(boil.AfterInsertHook, originPrivateEncryptionKeyAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	originPrivateEncryptionKeyAfterInsertHooks = []OriginPrivateEncryptionKeyHook{}

	AddOriginPrivateEncryptionKeyHook(boil.AfterSelectHook, originPrivateEncryptionKeyAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	originPrivateEncryptionKeyAfterSelectHooks = []OriginPrivateEncryptionKeyHook{}

	AddOriginPrivateEncryptionKeyHook(boil.BeforeUpdateHook, originPrivateEncryptionKeyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	originPrivateEncryptionKeyBeforeUpdateHooks = []OriginPrivateEncryptionKeyHook{}

	AddOriginPrivateEncryptionKeyHook(boil.AfterUpdateHook, originPrivateEncryptionKeyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	originPrivateEncryptionKeyAfterUpdateHooks = []OriginPrivateEncryptionKeyHook{}

	AddOriginPrivateEncryptionKeyHook(boil.BeforeDeleteHook, originPrivateEncryptionKeyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	originPrivateEncryptionKeyBeforeDeleteHooks = []OriginPrivateEncryptionKeyHook{}

	AddOriginPrivateEncryptionKeyHook(boil.AfterDeleteHook, originPrivateEncryptionKeyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	originPrivateEncryptionKeyAfterDeleteHooks = []OriginPrivateEncryptionKeyHook{}

	AddOriginPrivateEncryptionKeyHook(boil.BeforeUpsertHook, originPrivateEncryptionKeyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	originPrivateEncryptionKeyBeforeUpsertHooks = []OriginPrivateEncryptionKeyHook{}

	AddOriginPrivateEncryptionKeyHook(boil.AfterUpsertHook, originPrivateEncryptionKeyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	originPrivateEncryptionKeyAfterUpsertHooks = []OriginPrivateEncryptionKeyHook{}
}

func testOriginPrivateEncryptionKeysInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginPrivateEncryptionKeysInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(originPrivateEncryptionKeyColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginPrivateEncryptionKeyToOneOriginUsingOriginName(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OriginPrivateEncryptionKey
	var foreign Origin

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
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

	slice := OriginPrivateEncryptionKeySlice{&local}
	if err = local.L.LoadOriginName(ctx, tx, false, (*[]*OriginPrivateEncryptionKey)(&slice), nil); err != nil {
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

func testOriginPrivateEncryptionKeyToOneSetOpOriginUsingOriginName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginPrivateEncryptionKey
	var b, c Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originPrivateEncryptionKeyDBTypes, false, strmangle.SetComplement(originPrivateEncryptionKeyPrimaryKeyColumns, originPrivateEncryptionKeyColumnsWithoutDefault)...); err != nil {
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

		if x.R.OriginOPREK[0] != &a {
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

func testOriginPrivateEncryptionKeyToOneRemoveOpOriginUsingOriginName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginPrivateEncryptionKey
	var b Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originPrivateEncryptionKeyDBTypes, false, strmangle.SetComplement(originPrivateEncryptionKeyPrimaryKeyColumns, originPrivateEncryptionKeyColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.OriginOPREK) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOriginPrivateEncryptionKeysReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
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

func testOriginPrivateEncryptionKeysReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginPrivateEncryptionKeySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOriginPrivateEncryptionKeysSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginPrivateEncryptionKeys().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	originPrivateEncryptionKeyDBTypes = map[string]string{`ID`: `bigint`, `OwnerID`: `bigint`, `Name`: `text`, `Revision`: `text`, `FullName`: `text`, `Body`: `bytea`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `Origin`: `text`}
	_                                 = bytes.MinRead
)

func testOriginPrivateEncryptionKeysUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(originPrivateEncryptionKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(originPrivateEncryptionKeyColumns) == len(originPrivateEncryptionKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOriginPrivateEncryptionKeysSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(originPrivateEncryptionKeyColumns) == len(originPrivateEncryptionKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originPrivateEncryptionKeyDBTypes, true, originPrivateEncryptionKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(originPrivateEncryptionKeyColumns, originPrivateEncryptionKeyPrimaryKeyColumns) {
		fields = originPrivateEncryptionKeyColumns
	} else {
		fields = strmangle.SetComplement(
			originPrivateEncryptionKeyColumns,
			originPrivateEncryptionKeyPrimaryKeyColumns,
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

	slice := OriginPrivateEncryptionKeySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOriginPrivateEncryptionKeysUpsert(t *testing.T) {
	t.Parallel()

	if len(originPrivateEncryptionKeyColumns) == len(originPrivateEncryptionKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OriginPrivateEncryptionKey{}
	if err = randomize.Struct(seed, &o, originPrivateEncryptionKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginPrivateEncryptionKey: %s", err)
	}

	count, err := OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, originPrivateEncryptionKeyDBTypes, false, originPrivateEncryptionKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginPrivateEncryptionKey struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginPrivateEncryptionKey: %s", err)
	}

	count, err = OriginPrivateEncryptionKeys().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
