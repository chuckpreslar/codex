// Package sql provides common SQL constants for the codex package.
package sql

type Constraint uint8

// SQL Constraint constants.
const(
  NOT_NULL Constraint = iota
  UNIQUE
  PRIMARY_KEY
  FOREIGN_KEY
  CHECK
  DEFAULT
)

type Type uint8

// SQL Type constants.
const (
  STRING Type = iota
  TEXT
  BOOLEAN
  INTEGER
  FLOAT
  DECIMAL
  DATE
  TIME
  DATETIME
  TIMESTAMP
)
