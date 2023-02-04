# Formatter

[![Build Status](https://travis-ci.org/mantyr/formatter.svg?branch=master)](https://travis-ci.org/mantyr/formatter)
[![GoDoc](https://godoc.org/github.com/mantyr/formatter?status.png)](http://godoc.org/github.com/mantyr/formatter)
[![Go Report Card](https://goreportcard.com/badge/github.com/mantyr/formatter?v=1)][goreport]
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

This is not a stable version.

## Description

Tabular data formatting package.

### Supports formats

- [x] table text/template
- [x] text/template
- [x] raw
- [x] pretty
- [x] json

## Installation

    $ go get github.com/mantyr/formatter

## Example

```go
package main

import (
	"github.com/mantyr/formatter"
)

func main() {
	err = Print(...)
	if err != nil {
		panic(err)
	}
}

func Print(items []interface{}) error {
	format := formatter.Format("{{table .ID\t.Name\t}}")
	header := formatter.Header(map[string]string{
		"ID":   "Identifier",
		"Name": "Name",
	})
	f, err := formatter.New(os.Stdout)
	if err != nil {
		return err
	}
	defer f.Flush()

	err = f.SetFormat(format)
	if err != nil {
		return err
	}

	err = f.SetHeader(header)
	if err != nil {
		return err
	}

	for _, item := range items {
		err = f.Write(item)
		if err != nil {
			return err
		}
	}
	return nil
}
```

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr

[build_status]: https://travis-ci.org/mantyr/formatter
[godoc]:        http://godoc.org/github.com/mantyr/formatter
[goreport]:     https://goreportcard.com/report/github.com/mantyr/formatter
