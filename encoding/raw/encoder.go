package raw

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

type Encoder struct {
	wr io.Writer
}

func NewEncoder(wr io.Writer) *Encoder {
	return &Encoder{
		wr: wr,
	}
}

func (e *Encoder) Encode(v interface{}) error {
	if v == nil {
		return errors.New("empty v")
	}
	value := reflect.ValueOf(v)
	typeInfo := value.Type()

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			itemName := typeInfo.Field(i).Name
			tag := ParseTag(typeInfo.Field(i).Tag.Get("raw"))
			itemColumn := tag.Get("column")
			if itemColumn != "" {
				itemName = itemColumn
			}
			itemValue := value.Field(i).String()
			itemLine := fmt.Sprintf("%s: %s\n", itemName, itemValue)
			n, err := e.wr.Write([]byte(itemLine))
			if err != nil {
				return err
			}
			if n != len(itemLine) {
				return fmt.Errorf("expected %d bytes but actual %d bytes", len(itemLine), n)
			}
		}
	}
	_, err := e.wr.Write([]byte("\n"))
	return err
}
