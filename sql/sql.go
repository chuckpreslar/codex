// Package sql provides common SQL constants for the codex package.
package sql

type Constraint uint8

// SQL Constraint constants.
const (
	NotNull Constraint = iota
	Unique
	PrimaryKey
	ForeignKey
	Check
	Default
)

type Type uint8

// SQL Type constants.
const (
	String Type = iota
	Text
	Boolean
	Integer
	Float
	Decimal
	Date
	Time
	Datetime
	Timestamp
)
