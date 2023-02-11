package formatter

import (
	"github.com/pkg/errors"
	"strings"
	"text/template"
)

// Format keys used to specify certain kinds of output formats
const (
	TableFormatKey  = "table"
	RawFormatKey    = "raw"
	PrettyFormatKey = "pretty"
	JSONFormatKey   = "json"
	YAMLFormatKey   = "yaml"

	DefaultQuietFormat = "{{.ID}}"
	jsonFormat         = "{{json .}}"
	rawFormat          = "{{raw .}}\n"
	yamlFormat         = "{{yaml .}}\n"
)

var (
	formatReplacer = strings.NewReplacer(`\t`, "\t", `\n`, "\n")
)

// Format is the format string rendered using the Context
type Format string

func (f Format) IsBuffered() bool {
	return f.IsTable()
}

// IsTable returns true if the format is a table-type format
func (f Format) IsTable() bool {
	return strings.HasPrefix(string(f), TableFormatKey)
}

// IsJSON returns true if the format is the json format
func (f Format) IsJSON() bool {
	return string(f) == JSONFormatKey
}

func (f Format) IsRaw() bool {
	return string(f) == RawFormatKey
}

func (f Format) IsYAML() bool {
	return string(f) == YAMLFormatKey
}

// Contains returns true if the format contains the substring
func (f Format) Contains(sub string) bool {
	return strings.Contains(string(f), sub)
}

func (f Format) String() string {
	format := string(f)
	switch {
	case f.IsTable():
		format = format[len(TableFormatKey):]
	case f.IsJSON():
		format = jsonFormat
	case f.IsRaw():
		format = rawFormat
	case f.IsYAML():
		format = yamlFormat
	}
	format = strings.Trim(format, " ")
	return formatReplacer.Replace(format)
}

func (f Format) Header() (*template.Template, error) {
	tmpl, err := template.New("").Funcs(HeaderFunctions).Parse(f.String())
	if err != nil {
		return nil, errors.Wrap(err, "template parsing error")
	}
	return tmpl, err
}

func (f Format) Template() (*template.Template, error) {
	tmpl, err := template.New("").Funcs(BasicFunctions).Parse(f.String())
	if err != nil {
		return nil, errors.Wrap(err, "template parsing error")
	}
	return tmpl, err
}
