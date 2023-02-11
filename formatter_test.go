package formatter

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey" //nolint:golint,stylecheck
)

type Item struct {
	ID   string `raw:"column:id"`
	Name string `raw:"column:name"`
}

func TestFormatter(t *testing.T) {
	Convey("Checking formatter", t, func() {
		items := []Item{
			{ID: "1231", Name: "name1"},
			{ID: "1232", Name: "name2sdasdasd"},
			{ID: "1233sdasasdasd", Name: "name3"},
			{ID: "1234", Name: "name4"},
		}
		header := Header(map[string]string{
			"ID":   "Identifier",
			"Name": "Name",
		})
		buf := bytes.NewBufferString("")

		f, err := New(buf)
		So(err, ShouldBeNil)
		So(f, ShouldNotBeNil)

		err = f.SetHeader(header)
		So(err, ShouldBeNil)

		Convey("raw", func() {
			format := Format("raw")
			expected := strings.Join(
				[]string{
					`id: 1231`,
					`name: name1`,
					``,
					`id: 1232`,
					`name: name2sdasdasd`,
					``,
					`id: 1233sdasasdasd`,
					`name: name3`,
					``,
					`id: 1234`,
					`name: name4`,
					``,
					``,
				},
				"\n",
			)
			check(buf, f, format, items, expected)
		})
		Convey("yaml", func() {
			format := Format("yaml")
			expected := strings.Join(
				[]string{
					`id: "1231"`,
					`name: name1`,
					``,
					`id: "1232"`,
					`name: name2sdasdasd`,
					``,
					`id: 1233sdasasdasd`,
					`name: name3`,
					``,
					`id: "1234"`,
					`name: name4`,
					``,
					``,
				},
				"\n",
			)
			check(buf, f, format, items, expected)
		})
		Convey("{{.ID}}\\t{{.Name}}", func() {
			format := Format("{{.ID}}\t{{.Name}}")
			expected := strings.Join(
				[]string{
					"1231\tname1",
					"1232\tname2sdasdasd",
					"1233sdasasdasd\tname3",
					"1234\tname4",
					``,
				},
				"\n",
			)
			check(buf, f, format, items, expected)
		})
		Convey("table {{.ID}}\\t{{.Name}}", func() {
			format := Format("table {{.ID}}\t{{.Name}}")
			expected := strings.Join(
				[]string{
					"Identifier       Name",
					"1231             name1",
					"1232             name2sdasdasd",
					"1233sdasasdasd   name3",
					"1234             name4",
					``,
				},
				"\n",
			)
			check(buf, f, format, items, expected)
		})
		Convey("json", func() {
			format := Format("json")
			expected := strings.Join(
				[]string{
					`{"ID":"1231","Name":"name1"}`,
					`{"ID":"1232","Name":"name2sdasdasd"}`,
					`{"ID":"1233sdasasdasd","Name":"name3"}`,
					`{"ID":"1234","Name":"name4"}`,
					``,
				},
				"\n",
			)
			check(buf, f, format, items, expected)
		})
		Convey("Bad format", func() {
			Convey("No field in struct - no table", func() {
				format := Format("{{.DDD}}")
				expected := "bad format"
				checkError(buf, f, format, items, expected)
			})
			Convey("No field in struct - table", func() {
				format := Format("table {{.DDD}}")
				expected := "bad format"
				checkError(buf, f, format, items, expected)
			})
			Convey("text/template format error", func() {
				format := Format("table {{.OS}}\t{{.Architecture]}\t{{.CreatedAt}}")
				err := f.SetFormat(format)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "template parsing error: template: :1: bad character U+005D ']'")
			})
		})
	})
}

func check(buf *bytes.Buffer, f Formatter, format Format, items []Item, expected string) {
	So(f.SetFormat(format), ShouldBeNil)
	for _, item := range items {
		So(f.Write(item), ShouldBeNil)
	}
	So(f.Flush(), ShouldBeNil)
	So(buf.String(), ShouldEqual, expected)
}

func checkError(buf *bytes.Buffer, f Formatter, format Format, items []Item, expectedError string) {
	for _, item := range items {
		err := f.Write(item)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldEqual, expectedError)
	}
}
