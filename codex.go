// Package codex provides a Relational Algebra.
package codex

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
  "github.com/chuckpreslar/codex/visitors"
)

// Expose sql packages Constraint types.
const (
  NOT_NULL    = sql.NOT_NULL
  UNIQUE      = sql.UNIQUE
  PRIMARY_KEY = sql.PRIMARY_KEY
  FOREIGN_KEY = sql.FOREIGN_KEY
  CHECK       = sql.CHECK
  DEFAULT     = sql.DEFAULT
)

// Expose sql packages Type types.
const (
  STRING    = sql.STRING
  TEXT      = sql.TEXT
  BOOLEAN   = sql.BOOLEAN
  INTEGER   = sql.INTEGER
  FLOAT     = sql.FLOAT
  DECIMAL   = sql.DECIMAL
  DATE      = sql.DATE
  TIME      = sql.TIME
  DATETIME  = sql.DATETIME
  TIMESTAMP = sql.TIMESTAMP
)

// ToggleDebugMode toggles debugger variable for managers package.
func ToggleDebugMode() {
  visitors.DEBUG = !visitors.DEBUG
}

// Table returns an Accessor from the managers package for
// generating SQL to interact with existing tables.
func Table(name string) managers.Accessor {
  relation := nodes.Relation(name)
  return func(name interface{}) *nodes.AttributeNode {
    if _, ok := name.(string); ok {
      return nodes.Attribute(nodes.Column(name), relation)
    }

    return nodes.Attribute(name, relation)
  }
}

// CreateTable returns an AlterManager from the managers package
// for generating SQL to create new tables.
func CreateTable(name string) *managers.AlterManager {
  relation := nodes.Relation(name)
  return managers.Alteration(relation, true)
}

// CreateTable returns an AlterManager from the managers package
// for generating SQL to alter existing tables.
func AlterTable(name string) *managers.AlterManager {
  relation := nodes.Relation(name)
  return managers.Alteration(relation, false)
}
