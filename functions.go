package formatter

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"
)

// BasicFunctions are the set of initial
// functions provided to every template.
var BasicFunctions = template.FuncMap{
	"json":     MarshalJSON,
	"split":    strings.Split,
	"join":     strings.Join,
	"title":    strings.Title, //nolint:staticcheck // strings.Title is deprecated, but we only use it for ASCII, so replacing with golang.org/x/text is out of scope
	"lower":    strings.ToLower,
	"upper":    strings.ToUpper,
	"pad":      PadWithSpace,
	"truncate": TruncateWithLength,
}

// HeaderFunctions are used to created headers of a table.
// This is a replacement of basicFunctions for header generation
// because we want the header to remain intact.
// Some functions like `pad` are not overridden (to preserve alignment
// with the columns).
var HeaderFunctions = template.FuncMap{
	"json": func(v string) string {
		return v
	},
	"split": func(v string, _ string) string {
		// we want the table header to show the name of the column, and not
		// split the table header itself. Using a different signature
		// here, and return a string instead of []string
		return v
	},
	"join": func(v string, _ string) string {
		// table headers are always a string, so use a different signature
		// for the "join" function (string instead of []string)
		return v
	},
	"title": func(v string) string {
		return v
	},
	"lower": func(v string) string {
		return v
	},
	"upper": func(v string) string {
		return v
	},
	"truncate": func(v string, _ int) string {
		return v
	},
}

func MarshalJSON(v interface{}) string {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.Encode(v)
	return strings.TrimSpace(buf.String())
}

// PadWithSpace adds whitespace to the input if the input is non-empty
func PadWithSpace(source string, prefix, suffix int) string {
	if source == "" {
		return source
	}
	return strings.Repeat(" ", prefix) + source + strings.Repeat(" ", suffix)
}

// TruncateWithLength truncates the source string up to the length provided by the input
func TruncateWithLength(source string, length int) string {
	if len(source) < length {
		return source
	}
	return source[:length]
}
