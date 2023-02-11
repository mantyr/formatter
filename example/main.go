package main

import (
	"os"

	"github.com/mantyr/formatter"
)

func main() {
	items := []Item{
		{ID: "1231", Name: "name1"},
		{ID: "1232", Name: "name2sdasdasd"},
		{ID: "1233sdasasdasd", Name: "name3"},
		{ID: "1234", Name: "name4"},
	}
	//	format := formatter.Format("{{.ID}}\t{{.Name}}\t")
	//	format := formatter.Format("table {{.}}")
	format := formatter.Format("json")
	header := formatter.Header(map[string]string{
		"ID":   "Identifier",
		"Name": "Name",
	})
	f, err := formatter.New(os.Stdout)
	if err != nil {
		panic(err)
	}
	defer f.Flush()

	err = f.SetFormat(format)
	if err != nil {
		panic(err)
	}

	err = f.SetHeader(header)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		err = f.Write(item)
		if err != nil {
			panic(err)
		}
	}
}

type Item struct {
	ID   string
	Name string
}
