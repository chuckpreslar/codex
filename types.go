package librarian

import (
	"strings"
)

type column string
type columns []column
type table string
type accessor func(string) column

func (c columns) join(d string) string {
	s := make([]string, len(c))
	for _, cc := range c {
		s = append(s, string(cc))
	}
	return strings.Join(s, d)
}
