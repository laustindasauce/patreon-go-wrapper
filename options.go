package patreon

import (
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

type options struct {
	fields  map[string]string
	include string
	size    int
	cursor  string
}

type requestOption func(*options)

// WithFields specifies the resource attributes you want to be returned by API.
func WithFields(resource string, fields ...string) requestOption {
	return func(o *options) {
		if o.fields == nil {
			o.fields = make(map[string]string)
		}
		o.fields[resource] = strings.Join(fields, ",")
	}
}

// WithIncludes specifies the related resources you want to be returned by API.
func WithIncludes(include ...string) requestOption {
	return func(o *options) {
		o.include = strings.Join(include, ",")
	}
}

// WithPageSize specifies the number of items to return.
func WithPageSize(size int) requestOption {
	return func(o *options) {
		o.size = size
	}
}

// WithCursor controls cursor-based pagination. Cursor will also be extracted from navigation links for convenience.
func WithCursor(cursor string) requestOption {
	return func(o *options) {
		u, err := url.ParseRequestURI(cursor)
		if err == nil {
			cursor = u.Query().Get("page[cursor]")
		}

		o.cursor = cursor
	}
}

func getOptions(opts ...requestOption) options {
	cfg := options{}
	for _, fn := range opts {
		fn(&cfg)
	}

	return cfg
}

// getObjectFields will get all fields for an object
func getObjectFields(i interface{}) []string {
	v := reflect.ValueOf(i)
	typeOfS := v.Type()

	var fields []string

	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, toSnakeCase(typeOfS.Field(i).Name))
	}

	return fields
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
