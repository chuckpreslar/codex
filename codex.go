// Package codex provides a Relational Algebra.
package codex

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
)

// Alias constants for common SQL constraints.
const (
  NOT_NUL        = sql.NOT_NULL
  UNIQUE         = sql.UNIQUE
  PRIMARY_KEY    = sql.PRIMARY_KEY
  AUTO_INCREMENT = sql.AUTO_INCREMENT
  FOREIGN_KEY    = sql.FOREIGN_KEY
  CHECK          = sql.CHECK
  DEFAULT        = sql.DEFAULT
)

// Alias constants for common SQL types.
const (
  STRING    = sql.STRING
  TEXT      = sql.TEXT
  INTEGER   = sql.INTEGER
  FLOAT     = sql.FLOAT
  DECIMAL   = sql.DECIMAL
  DATE      = sql.DATE
  DATETIME  = sql.DATETIME
  TIME      = sql.TIME
  TIMESTAMP = sql.TIMESTAMP
)

// Table returns an Accessor
func Table(name string) managers.Accessor {
  relation := nodes.Relation(name)
  return func(name interface{}) *nodes.AttributeNode {
    if _, ok := name.(string); ok {
      return nodes.Attribute(nodes.Column(name), relation)
    }

    return nodes.Attribute(name, relation)
  }
}

func CreateTable(name string) *managers.CreateManager {
  relation := nodes.Relation(name)
  return managers.Creation(relation)
}
