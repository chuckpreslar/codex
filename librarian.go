package librarian

type column string
type columns []column
type table string
type accessor func(string) column

func Search(a accessor) *SelectStatement {
	return &SelectStatement{a: a, projections: []column{}, reference: table(a(""))}
}

func Insert(a accessor) *InsertStatement {
	return &InsertStatement{a: a}
}

func Update(a accessor) *UpdateStatement {
	return &UpdateStatement{a: a}
}

func Delete(a accessor) *DeleteStatement {
	return &DeleteStatement{a: a}
}
