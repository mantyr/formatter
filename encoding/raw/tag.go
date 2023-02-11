package raw

import (
	"strings"
)

type Tag struct {
	data map[string]string
}

func (t *Tag) Set(key, value string) {
	t.data[key] = value
}

func (t *Tag) Get(key string) string {
	return t.data[key]
}

func (t *Tag) Is(key string) bool {
	_, ok := t.data[key]
	return ok
}

func NewTag() *Tag {
	return &Tag{
		data: make(map[string]string),
	}
}

func ParseTag(s string) *Tag {
	t := NewTag()
	items := strings.Split(s, ",")
	for _, item := range items {
		delim := strings.Index(item, ":")
		switch {
		case delim == 0:
			continue
		case delim > 0:
			t.Set(item[:delim], item[delim+1:])
		}
	}
	return t
}
