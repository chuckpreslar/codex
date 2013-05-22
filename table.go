package librarian

import (
	"fmt"
)

func Table(t string) accessor {
	return func(c string) column {
		switch c {
		case "":
			return column(fmt.Sprintf("\"%s\"", t))
		case "*":
			return column(fmt.Sprintf("\"%s\".%s", t, c))
		default:
			return column(fmt.Sprintf("\"%s\".\"%s\"", t, c))
		}
	}
}
