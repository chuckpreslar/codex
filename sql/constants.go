// Package sql provides common SQL constants.
package sql

type Constraint uint8

// Constants for common SQL constraints.
const (
  NOT_NULL Constraint = iota
  UNIQUE
  PRIMARY_KEY
  AUTO_INCREMENT
  FOREIGN_KEY
  CHECK
  DEFAULT
)

type Type uint8

// Constants for common SQL types.
const (
  STRING Type = iota
  TEXT
  INTEGER
  FLOAT
  DECIMAL
  DATE
  DATETIME
  TIME
  TIMESTAMP
)
