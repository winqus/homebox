// Code generated by ent, DO NOT EDIT.

package maintenanceentry

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the maintenanceentry type in the database.
	Label = "maintenance_entry"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldItemID holds the string denoting the item_id field in the database.
	FieldItemID = "item_id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldScheduledDate holds the string denoting the scheduled_date field in the database.
	FieldScheduledDate = "scheduled_date"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCost holds the string denoting the cost field in the database.
	FieldCost = "cost"
	// EdgeItem holds the string denoting the item edge name in mutations.
	EdgeItem = "item"
	// Table holds the table name of the maintenanceentry in the database.
	Table = "maintenance_entries"
	// ItemTable is the table that holds the item relation/edge.
	ItemTable = "maintenance_entries"
	// ItemInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemInverseTable = "items"
	// ItemColumn is the table column denoting the item relation/edge.
	ItemColumn = "item_id"
)

// Columns holds all SQL columns for maintenanceentry fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldItemID,
	FieldDate,
	FieldScheduledDate,
	FieldName,
	FieldDescription,
	FieldCost,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultCost holds the default value on creation for the "cost" field.
	DefaultCost float64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
