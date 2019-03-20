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

func testOriginChannelPackages(t *testing.T) {
	t.Parallel()

	query := OriginChannelPackages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOriginChannelPackagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
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

	count, err := OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginChannelPackagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OriginChannelPackages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginChannelPackagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginChannelPackageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginChannelPackagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OriginChannelPackageExists(ctx, tx, o.ChannelID, o.PackageID)
	if err != nil {
		t.Errorf("Unable to check if OriginChannelPackage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OriginChannelPackageExists to return true, but got false.")
	}
}

func testOriginChannelPackagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	originChannelPackageFound, err := FindOriginChannelPackage(ctx, tx, o.ChannelID, o.PackageID)
	if err != nil {
		t.Error(err)
	}

	if originChannelPackageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOriginChannelPackagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OriginChannelPackages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOriginChannelPackagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OriginChannelPackages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOriginChannelPackagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	originChannelPackageOne := &OriginChannelPackage{}
	originChannelPackageTwo := &OriginChannelPackage{}
	if err = randomize.Struct(seed, originChannelPackageOne, originChannelPackageDBTypes, false, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}
	if err = randomize.Struct(seed, originChannelPackageTwo, originChannelPackageDBTypes, false, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originChannelPackageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originChannelPackageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginChannelPackages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOriginChannelPackagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	originChannelPackageOne := &OriginChannelPackage{}
	originChannelPackageTwo := &OriginChannelPackage{}
	if err = randomize.Struct(seed, originChannelPackageOne, originChannelPackageDBTypes, false, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}
	if err = randomize.Struct(seed, originChannelPackageTwo, originChannelPackageDBTypes, false, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originChannelPackageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originChannelPackageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func originChannelPackageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginChannelPackage) error {
	*o = OriginChannelPackage{}
	return nil
}

func originChannelPackageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginChannelPackage) error {
	*o = OriginChannelPackage{}
	return nil
}

func originChannelPackageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OriginChannelPackage) error {
	*o = OriginChannelPackage{}
	return nil
}

func originChannelPackageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginChannelPackage) error {
	*o = OriginChannelPackage{}
	return nil
}

func originChannelPackageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginChannelPackage) error {
	*o = OriginChannelPackage{}
	return nil
}

func originChannelPackageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginChannelPackage) error {
	*o = OriginChannelPackage{}
	return nil
}

func originChannelPackageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginChannelPackage) error {
	*o = OriginChannelPackage{}
	return nil
}

func originChannelPackageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginChannelPackage) error {
	*o = OriginChannelPackage{}
	return nil
}

func originChannelPackageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginChannelPackage) error {
	*o = OriginChannelPackage{}
	return nil
}

func testOriginChannelPackagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OriginChannelPackage{}
	o := &OriginChannelPackage{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage object: %s", err)
	}

	AddOriginChannelPackageHook(boil.BeforeInsertHook, originChannelPackageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	originChannelPackageBeforeInsertHooks = []OriginChannelPackageHook{}

	AddOriginChannelPackageHook(boil.AfterInsertHook, originChannelPackageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	originChannelPackageAfterInsertHooks = []OriginChannelPackageHook{}

	AddOriginChannelPackageHook(boil.AfterSelectHook, originChannelPackageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	originChannelPackageAfterSelectHooks = []OriginChannelPackageHook{}

	AddOriginChannelPackageHook(boil.BeforeUpdateHook, originChannelPackageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	originChannelPackageBeforeUpdateHooks = []OriginChannelPackageHook{}

	AddOriginChannelPackageHook(boil.AfterUpdateHook, originChannelPackageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	originChannelPackageAfterUpdateHooks = []OriginChannelPackageHook{}

	AddOriginChannelPackageHook(boil.BeforeDeleteHook, originChannelPackageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	originChannelPackageBeforeDeleteHooks = []OriginChannelPackageHook{}

	AddOriginChannelPackageHook(boil.AfterDeleteHook, originChannelPackageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	originChannelPackageAfterDeleteHooks = []OriginChannelPackageHook{}

	AddOriginChannelPackageHook(boil.BeforeUpsertHook, originChannelPackageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	originChannelPackageBeforeUpsertHooks = []OriginChannelPackageHook{}

	AddOriginChannelPackageHook(boil.AfterUpsertHook, originChannelPackageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	originChannelPackageAfterUpsertHooks = []OriginChannelPackageHook{}
}

func testOriginChannelPackagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginChannelPackagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(originChannelPackageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginChannelPackageToOneOriginChannelUsingChannel(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OriginChannelPackage
	var foreign OriginChannel

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, originChannelPackageDBTypes, false, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, originChannelDBTypes, false, originChannelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannel struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ChannelID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Channel().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OriginChannelPackageSlice{&local}
	if err = local.L.LoadChannel(ctx, tx, false, (*[]*OriginChannelPackage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Channel == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Channel = nil
	if err = local.L.LoadChannel(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Channel == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOriginChannelPackageToOneOriginPackageUsingPackage(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OriginChannelPackage
	var foreign OriginPackage

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, originChannelPackageDBTypes, false, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, originPackageDBTypes, false, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PackageID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Package().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OriginChannelPackageSlice{&local}
	if err = local.L.LoadPackage(ctx, tx, false, (*[]*OriginChannelPackage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Package == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Package = nil
	if err = local.L.LoadPackage(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Package == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOriginChannelPackageToOneSetOpOriginChannelUsingChannel(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginChannelPackage
	var b, c OriginChannel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originChannelPackageDBTypes, false, strmangle.SetComplement(originChannelPackagePrimaryKeyColumns, originChannelPackageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, originChannelDBTypes, false, strmangle.SetComplement(originChannelPrimaryKeyColumns, originChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, originChannelDBTypes, false, strmangle.SetComplement(originChannelPrimaryKeyColumns, originChannelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OriginChannel{&b, &c} {
		err = a.SetChannel(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Channel != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ChannelOriginChannelPackages[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ChannelID != x.ID {
			t.Error("foreign key was wrong value", a.ChannelID)
		}

		if exists, err := OriginChannelPackageExists(ctx, tx, a.ChannelID, a.PackageID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testOriginChannelPackageToOneSetOpOriginPackageUsingPackage(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginChannelPackage
	var b, c OriginPackage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originChannelPackageDBTypes, false, strmangle.SetComplement(originChannelPackagePrimaryKeyColumns, originChannelPackageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, originPackageDBTypes, false, strmangle.SetComplement(originPackagePrimaryKeyColumns, originPackageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, originPackageDBTypes, false, strmangle.SetComplement(originPackagePrimaryKeyColumns, originPackageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*OriginPackage{&b, &c} {
		err = a.SetPackage(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Package != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PackageOriginChannelPackages[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PackageID != x.ID {
			t.Error("foreign key was wrong value", a.PackageID)
		}

		if exists, err := OriginChannelPackageExists(ctx, tx, a.ChannelID, a.PackageID); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testOriginChannelPackagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
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

func testOriginChannelPackagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginChannelPackageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOriginChannelPackagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginChannelPackages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	originChannelPackageDBTypes = map[string]string{`ChannelID`: `bigint`, `PackageID`: `bigint`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                           = bytes.MinRead
)

func testOriginChannelPackagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(originChannelPackagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(originChannelPackageColumns) == len(originChannelPackagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOriginChannelPackagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(originChannelPackageColumns) == len(originChannelPackagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginChannelPackage{}
	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originChannelPackageDBTypes, true, originChannelPackagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(originChannelPackageColumns, originChannelPackagePrimaryKeyColumns) {
		fields = originChannelPackageColumns
	} else {
		fields = strmangle.SetComplement(
			originChannelPackageColumns,
			originChannelPackagePrimaryKeyColumns,
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

	slice := OriginChannelPackageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOriginChannelPackagesUpsert(t *testing.T) {
	t.Parallel()

	if len(originChannelPackageColumns) == len(originChannelPackagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OriginChannelPackage{}
	if err = randomize.Struct(seed, &o, originChannelPackageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginChannelPackage: %s", err)
	}

	count, err := OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, originChannelPackageDBTypes, false, originChannelPackagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginChannelPackage struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginChannelPackage: %s", err)
	}

	count, err = OriginChannelPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
