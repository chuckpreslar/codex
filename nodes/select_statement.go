package nodes

type SelectStatement struct {
  BaseNode
  reference   *Reference
  projections []Projection
  filters     []Comparison
}

func (stmt *SelectStatement) Select(projections ...Projection) {
  stmt.projections = append(stmt.projections, projections...)
}

func (stmt *SelectStatement) Where(comparison Comparison) {
  stmt.filters = append(stmt.filters, comparison)
}

func NewSelectStatement(reference *Reference) *SelectStatement {
  stmt := new(SelectStatement)
  stmt.reference = reference
  stmt.projections = make([]Projection, 0)
  stmt.filters = make([]Comparison, 0)
  return stmt
}
