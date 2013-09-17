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
  NotNull    = sql.NotNull
  Unique     = sql.Unique
  PrimaryKey = sql.PrimaryKey
  ForeignKey = sql.ForeignKey
  Check      = sql.Check
  Default    = sql.Default
)

// Expose sql packages Type types.
const (
  String    = sql.String
  Text      = sql.Text
  Boolean   = sql.Boolean
  Integer   = sql.Integer
  Float     = sql.Float
  Decimal   = sql.Decimal
  Date      = sql.Date
  Time      = sql.Time
  Datetime  = sql.Datetime
  Timestamp = sql.Timestamp
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
func CreateTable(name string) *managers.CreateManager {
  relation := nodes.Relation(name)
  return managers.Creation(relation)
}

// CreateTable returns an AlterManager from the managers package
// for generating SQL to alter existing tables.
func AlterTable(name string) *managers.AlterManager {
  relation := nodes.Relation(name)
  return managers.Alteration(relation)
}
