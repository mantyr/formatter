package formatter

import (
	"strings"
)

var (
	labelReplacer = strings.NewReplacer("-", " ", "_", " ", ".", "_")
)

// Header is a map destined to formatter header (table format)
type Header map[string]string

// Label returns the header label for the specified string
func (h Header) Label(name string) string {
	i := strings.IndexRune(name, '.')
	if i >= 0 {
		name = name[i:]
	}
	return labelReplacer.Replace(h[name])
}
