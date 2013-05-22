package librarian

type (
	Statement interface {
		ToSQL() string
	}
)

type (
	SelectStatement struct {
		a           accessor
		projections []column
		reference   table
	}

	InsertStatement struct {
		a accessor
	}
	UpdateStatement struct {
		a accessor
	}
	DeleteStatement struct {
		a accessor
	}
)

func (s *SelectStatement) Select(c ...string) *SelectStatement {
	for _, cc := range c {
		s.projections = append(s.projections, s.a(cc))
	}
	return s
}
