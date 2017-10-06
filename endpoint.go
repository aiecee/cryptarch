package cryptarch

import (
	"fmt"
	"strings"
)

const basePath = "http://bungie.net/platform"

type endpoint struct {
	method          string
	path            string
	pathParameters  map[string]string
	queryParameters map[string]string
}

func (e endpoint) build() string {
	path := basePath + e.path
	for pathParam, pathValue := range e.pathParameters {
		path = strings.Replace(path, "{"+pathParam+"}", pathValue, -1)
	}
	if len(e.queryParameters) > 0 {
		path += "?"
		i := 0
		for queryParam, queryValue := range e.queryParameters {
			path += queryParam + "=" + queryValue
			i++
			if i < len(e.queryParameters) {
				path += "&"
			}
		}
	}
	fmt.Println(path)
	return path
}
