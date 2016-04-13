package table

import "github.com/chuckpreslar/codex/lib/nodes"

// Protected attributes ...
const (
	_relation = "_relation"
)

// _protectedEntries ...
var _protectedEntries = []string{_relation}

// isProtectedEntry ...
func isProtectedEntry(name string) bool {
	var index int

	for index = 0; index < len(_protectedEntries); index++ {
		if name == _protectedEntries[index] {
			return true
		}
	}

	return false
}

// Table ...
type Table map[string]*nodes.Attribute

// DefineColumns ...
func (t Table) DefineColumns(names ...string) Table {
	var index int

	for index = 0; index < len(names); index++ {
		if isProtectedEntry(names[index]) {
			continue
		}

		t[names[index]] = nodes.InitializeAttribute(names[index], t.Relation())
	}

	return t
}

// DefineColumn ...
func (t Table) DefineColumn(name string) Table {
	return t.DefineColumns(name)
}

// Columns ...
func (t Table) Columns() []string {
	var (
		columns []string
		column  string
	)

	for column, _ = range t {
		if isProtectedEntry(column) {
			continue
		}

		columns = append(columns, column)
	}

	return columns
}

// Relation ...
func (t Table) Relation() *nodes.Relation {
	return t[_relation].Relation
}

// Name ...
func (t Table) Name() string {
	return t.Relation().Name
}

// New ...
func New(name string) Table {
	var table = make(Table)

	table[_relation] = nodes.InitializeAttribute("", nodes.InitializeRelation(name))

	return table
}
