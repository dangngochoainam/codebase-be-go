package querystringhelper

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

func FromStruct(data any) (url.Values, error) {
	result, err := query.Values(data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
