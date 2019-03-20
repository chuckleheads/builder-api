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

func testOriginInvitations(t *testing.T) {
	t.Parallel()

	query := OriginInvitations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOriginInvitationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
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

	count, err := OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginInvitationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OriginInvitations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginInvitationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginInvitationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginInvitationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OriginInvitationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OriginInvitation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OriginInvitationExists to return true, but got false.")
	}
}

func testOriginInvitationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	originInvitationFound, err := FindOriginInvitation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if originInvitationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOriginInvitationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OriginInvitations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOriginInvitationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OriginInvitations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOriginInvitationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	originInvitationOne := &OriginInvitation{}
	originInvitationTwo := &OriginInvitation{}
	if err = randomize.Struct(seed, originInvitationOne, originInvitationDBTypes, false, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}
	if err = randomize.Struct(seed, originInvitationTwo, originInvitationDBTypes, false, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originInvitationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originInvitationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginInvitations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOriginInvitationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	originInvitationOne := &OriginInvitation{}
	originInvitationTwo := &OriginInvitation{}
	if err = randomize.Struct(seed, originInvitationOne, originInvitationDBTypes, false, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}
	if err = randomize.Struct(seed, originInvitationTwo, originInvitationDBTypes, false, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originInvitationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originInvitationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func originInvitationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginInvitation) error {
	*o = OriginInvitation{}
	return nil
}

func originInvitationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginInvitation) error {
	*o = OriginInvitation{}
	return nil
}

func originInvitationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OriginInvitation) error {
	*o = OriginInvitation{}
	return nil
}

func originInvitationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginInvitation) error {
	*o = OriginInvitation{}
	return nil
}

func originInvitationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginInvitation) error {
	*o = OriginInvitation{}
	return nil
}

func originInvitationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginInvitation) error {
	*o = OriginInvitation{}
	return nil
}

func originInvitationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginInvitation) error {
	*o = OriginInvitation{}
	return nil
}

func originInvitationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginInvitation) error {
	*o = OriginInvitation{}
	return nil
}

func originInvitationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginInvitation) error {
	*o = OriginInvitation{}
	return nil
}

func testOriginInvitationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OriginInvitation{}
	o := &OriginInvitation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, originInvitationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OriginInvitation object: %s", err)
	}

	AddOriginInvitationHook(boil.BeforeInsertHook, originInvitationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	originInvitationBeforeInsertHooks = []OriginInvitationHook{}

	AddOriginInvitationHook(boil.AfterInsertHook, originInvitationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	originInvitationAfterInsertHooks = []OriginInvitationHook{}

	AddOriginInvitationHook(boil.AfterSelectHook, originInvitationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	originInvitationAfterSelectHooks = []OriginInvitationHook{}

	AddOriginInvitationHook(boil.BeforeUpdateHook, originInvitationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	originInvitationBeforeUpdateHooks = []OriginInvitationHook{}

	AddOriginInvitationHook(boil.AfterUpdateHook, originInvitationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	originInvitationAfterUpdateHooks = []OriginInvitationHook{}

	AddOriginInvitationHook(boil.BeforeDeleteHook, originInvitationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	originInvitationBeforeDeleteHooks = []OriginInvitationHook{}

	AddOriginInvitationHook(boil.AfterDeleteHook, originInvitationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	originInvitationAfterDeleteHooks = []OriginInvitationHook{}

	AddOriginInvitationHook(boil.BeforeUpsertHook, originInvitationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	originInvitationBeforeUpsertHooks = []OriginInvitationHook{}

	AddOriginInvitationHook(boil.AfterUpsertHook, originInvitationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	originInvitationAfterUpsertHooks = []OriginInvitationHook{}
}

func testOriginInvitationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginInvitationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(originInvitationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginInvitationToOneOriginUsingOrigin(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OriginInvitation
	var foreign Origin

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
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

	slice := OriginInvitationSlice{&local}
	if err = local.L.LoadOrigin(ctx, tx, false, (*[]*OriginInvitation)(&slice), nil); err != nil {
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

func testOriginInvitationToOneSetOpOriginUsingOrigin(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginInvitation
	var b, c Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originInvitationDBTypes, false, strmangle.SetComplement(originInvitationPrimaryKeyColumns, originInvitationColumnsWithoutDefault)...); err != nil {
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

		if x.R.OriginOI[0] != &a {
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

func testOriginInvitationToOneRemoveOpOriginUsingOrigin(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginInvitation
	var b Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originInvitationDBTypes, false, strmangle.SetComplement(originInvitationPrimaryKeyColumns, originInvitationColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.OriginOI) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOriginInvitationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
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

func testOriginInvitationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginInvitationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOriginInvitationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginInvitations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	originInvitationDBTypes = map[string]string{`ID`: `bigint`, `OriginName`: `text`, `AccountID`: `bigint`, `AccountName`: `text`, `OwnerID`: `bigint`, `Ignored`: `boolean`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                       = bytes.MinRead
)

func testOriginInvitationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(originInvitationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(originInvitationColumns) == len(originInvitationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOriginInvitationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(originInvitationColumns) == len(originInvitationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginInvitation{}
	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originInvitationDBTypes, true, originInvitationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(originInvitationColumns, originInvitationPrimaryKeyColumns) {
		fields = originInvitationColumns
	} else {
		fields = strmangle.SetComplement(
			originInvitationColumns,
			originInvitationPrimaryKeyColumns,
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

	slice := OriginInvitationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOriginInvitationsUpsert(t *testing.T) {
	t.Parallel()

	if len(originInvitationColumns) == len(originInvitationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OriginInvitation{}
	if err = randomize.Struct(seed, &o, originInvitationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginInvitation: %s", err)
	}

	count, err := OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, originInvitationDBTypes, false, originInvitationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginInvitation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginInvitation: %s", err)
	}

	count, err = OriginInvitations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
