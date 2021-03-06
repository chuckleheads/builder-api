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

func testOriginPackages(t *testing.T) {
	t.Parallel()

	query := OriginPackages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOriginPackagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
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

	count, err := OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginPackagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OriginPackages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginPackagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginPackageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOriginPackagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OriginPackageExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OriginPackage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OriginPackageExists to return true, but got false.")
	}
}

func testOriginPackagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	originPackageFound, err := FindOriginPackage(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if originPackageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOriginPackagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OriginPackages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOriginPackagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OriginPackages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOriginPackagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	originPackageOne := &OriginPackage{}
	originPackageTwo := &OriginPackage{}
	if err = randomize.Struct(seed, originPackageOne, originPackageDBTypes, false, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}
	if err = randomize.Struct(seed, originPackageTwo, originPackageDBTypes, false, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originPackageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originPackageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginPackages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOriginPackagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	originPackageOne := &OriginPackage{}
	originPackageTwo := &OriginPackage{}
	if err = randomize.Struct(seed, originPackageOne, originPackageDBTypes, false, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}
	if err = randomize.Struct(seed, originPackageTwo, originPackageDBTypes, false, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = originPackageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = originPackageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func originPackageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPackage) error {
	*o = OriginPackage{}
	return nil
}

func originPackageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPackage) error {
	*o = OriginPackage{}
	return nil
}

func originPackageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OriginPackage) error {
	*o = OriginPackage{}
	return nil
}

func originPackageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginPackage) error {
	*o = OriginPackage{}
	return nil
}

func originPackageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OriginPackage) error {
	*o = OriginPackage{}
	return nil
}

func originPackageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginPackage) error {
	*o = OriginPackage{}
	return nil
}

func originPackageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OriginPackage) error {
	*o = OriginPackage{}
	return nil
}

func originPackageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPackage) error {
	*o = OriginPackage{}
	return nil
}

func originPackageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OriginPackage) error {
	*o = OriginPackage{}
	return nil
}

func testOriginPackagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OriginPackage{}
	o := &OriginPackage{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, originPackageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OriginPackage object: %s", err)
	}

	AddOriginPackageHook(boil.BeforeInsertHook, originPackageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	originPackageBeforeInsertHooks = []OriginPackageHook{}

	AddOriginPackageHook(boil.AfterInsertHook, originPackageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	originPackageAfterInsertHooks = []OriginPackageHook{}

	AddOriginPackageHook(boil.AfterSelectHook, originPackageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	originPackageAfterSelectHooks = []OriginPackageHook{}

	AddOriginPackageHook(boil.BeforeUpdateHook, originPackageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	originPackageBeforeUpdateHooks = []OriginPackageHook{}

	AddOriginPackageHook(boil.AfterUpdateHook, originPackageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	originPackageAfterUpdateHooks = []OriginPackageHook{}

	AddOriginPackageHook(boil.BeforeDeleteHook, originPackageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	originPackageBeforeDeleteHooks = []OriginPackageHook{}

	AddOriginPackageHook(boil.AfterDeleteHook, originPackageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	originPackageAfterDeleteHooks = []OriginPackageHook{}

	AddOriginPackageHook(boil.BeforeUpsertHook, originPackageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	originPackageBeforeUpsertHooks = []OriginPackageHook{}

	AddOriginPackageHook(boil.AfterUpsertHook, originPackageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	originPackageAfterUpsertHooks = []OriginPackageHook{}
}

func testOriginPackagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginPackagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(originPackageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOriginPackageToManyPackageOriginChannelPackages(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginPackage
	var b, c OriginChannelPackage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, originChannelPackageDBTypes, false, originChannelPackageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, originChannelPackageDBTypes, false, originChannelPackageColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.PackageID = a.ID
	c.PackageID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.PackageOriginChannelPackages().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.PackageID == b.PackageID {
			bFound = true
		}
		if v.PackageID == c.PackageID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OriginPackageSlice{&a}
	if err = a.L.LoadPackageOriginChannelPackages(ctx, tx, false, (*[]*OriginPackage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PackageOriginChannelPackages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.PackageOriginChannelPackages = nil
	if err = a.L.LoadPackageOriginChannelPackages(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.PackageOriginChannelPackages); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOriginPackageToManyAddOpPackageOriginChannelPackages(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginPackage
	var b, c, d, e OriginChannelPackage

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originPackageDBTypes, false, strmangle.SetComplement(originPackagePrimaryKeyColumns, originPackageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OriginChannelPackage{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, originChannelPackageDBTypes, false, strmangle.SetComplement(originChannelPackagePrimaryKeyColumns, originChannelPackageColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*OriginChannelPackage{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddPackageOriginChannelPackages(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.PackageID {
			t.Error("foreign key was wrong value", a.ID, first.PackageID)
		}
		if a.ID != second.PackageID {
			t.Error("foreign key was wrong value", a.ID, second.PackageID)
		}

		if first.R.Package != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Package != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.PackageOriginChannelPackages[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.PackageOriginChannelPackages[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.PackageOriginChannelPackages().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testOriginPackageToOneOriginUsingOriginName(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OriginPackage
	var foreign Origin

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
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

	slice := OriginPackageSlice{&local}
	if err = local.L.LoadOriginName(ctx, tx, false, (*[]*OriginPackage)(&slice), nil); err != nil {
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

func testOriginPackageToOneSetOpOriginUsingOriginName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginPackage
	var b, c Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originPackageDBTypes, false, strmangle.SetComplement(originPackagePrimaryKeyColumns, originPackageColumnsWithoutDefault)...); err != nil {
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

		if x.R.OriginOPA[0] != &a {
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

func testOriginPackageToOneRemoveOpOriginUsingOriginName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OriginPackage
	var b Origin

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, originPackageDBTypes, false, strmangle.SetComplement(originPackagePrimaryKeyColumns, originPackageColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.OriginOPA) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOriginPackagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
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

func testOriginPackagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OriginPackageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOriginPackagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OriginPackages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	originPackageDBTypes = map[string]string{`ID`: `bigint`, `OwnerID`: `bigint`, `Name`: `text`, `Ident`: `text`, `IdentArray`: `ARRAYtext`, `Checksum`: `text`, `Manifest`: `text`, `Config`: `text`, `Target`: `text`, `Deps`: `ARRAYtext`, `Tdeps`: `ARRAYtext`, `Exposes`: `ARRAYinteger`, `SchedulerSync`: `boolean`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`, `Visibility`: `enum.origin_package_visibility('public','private','hidden')`, `IdentVector`: `tsvector`, `Origin`: `text`}
	_                    = bytes.MinRead
)

func testOriginPackagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(originPackagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(originPackageColumns) == len(originPackagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOriginPackagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(originPackageColumns) == len(originPackagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OriginPackage{}
	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, originPackageDBTypes, true, originPackagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(originPackageColumns, originPackagePrimaryKeyColumns) {
		fields = originPackageColumns
	} else {
		fields = strmangle.SetComplement(
			originPackageColumns,
			originPackagePrimaryKeyColumns,
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

	slice := OriginPackageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOriginPackagesUpsert(t *testing.T) {
	t.Parallel()

	if len(originPackageColumns) == len(originPackagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OriginPackage{}
	if err = randomize.Struct(seed, &o, originPackageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginPackage: %s", err)
	}

	count, err := OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, originPackageDBTypes, false, originPackagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OriginPackage struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OriginPackage: %s", err)
	}

	count, err = OriginPackages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
