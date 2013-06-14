package nodes

type SelectStatement struct {
  BaseNode
  reference   *Reference
  projections []Projection
  comparisons []Comparison
}

func (stmt *SelectStatement) Where(comparison Comparison) {
  stmt.comparisons = append(stmt.comparisons, comparison)
}
