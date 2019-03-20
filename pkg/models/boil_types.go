// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

// Enum values for origin_package_operation
const (
	OriginPackageOperationPromote = "promote"
	OriginPackageOperationDemote  = "demote"
)

// Enum values for package_channel_trigger
const (
	PackageChannelTriggerUnknown   = "unknown"
	PackageChannelTriggerBuilderUI = "builder_ui"
	PackageChannelTriggerHabClient = "hab_client"
)

// Enum values for origin_package_visibility
const (
	OriginPackageVisibilityPublic  = "public"
	OriginPackageVisibilityPrivate = "private"
	OriginPackageVisibilityHidden  = "hidden"
)
