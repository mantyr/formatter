// nolint:errcheck
package formatter

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"text/tabwriter"
	"text/template"
)

type formatter struct {
	header Header
	format Format

	templates struct {
		header *template.Template
		line   *template.Template
	}

	// buffer for table
	buffer *bytes.Buffer

	output io.Writer
}

func New(writer io.Writer) (Formatter, error) {
	if writer == nil {
		return nil, errors.New("empty writer")
	}
	return &formatter{
		output: writer,
	}, nil
}

func (f *formatter) SetFormat(format Format) (err error) {
	if format == "" {
		return fmt.Errorf("empty format")
	}
	f.format = format
	f.templates.header, err = format.Header()
	if err != nil {
		return err
	}
	f.templates.line, err = format.Template()
	if err != nil {
		return err
	}
	if format.IsBuffered() {
		f.buffer = bytes.NewBufferString("")
	}
	return nil
}

func (f *formatter) SetHeader(header Header) (err error) {
	if header == nil {
		return fmt.Errorf("empty header")
	}
	f.header = header
	return nil
}

func (f *formatter) Write(item interface{}) (err error) {
	if item == nil {
		return errors.New("empty item")
	}
	if f.buffer != nil {
		err = f.templates.line.Execute(f.buffer, item)
		if err != nil {
			return err
		}
		_, err = f.buffer.WriteString("\n")
		return err
	}
	err = f.templates.line.Execute(f.output, item)
	if err != nil {
		return err
	}
	_, err = f.output.Write([]byte("\n"))
	return err
}

func (f *formatter) Flush() error {
	switch {
	case f.format.IsTable():
		t := tabwriter.NewWriter(f.output, 10, 1, 3, ' ', 0)
		if f.header != nil {
			err := f.templates.header.Execute(t, f.header)
			if err != nil {
				return err
			}
			t.Write([]byte("\n"))
		}
		f.buffer.WriteTo(t)
		t.Flush()
		return nil
	case f.buffer != nil:
		f.buffer.WriteTo(f.output)
		return nil
	}
	return nil
}
