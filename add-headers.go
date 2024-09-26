package middleware

import (
	"net/http"
	"plugin-dev/util/ctx"
)

// AddFooBarHeader adds custom "Foo: Bar" header to the request
func AddFooBarHeader(rw http.ResponseWriter, r *http.Request) {
	definition := ctx.GetDefinition(r)
	r.Header.Add("Foo", definition.ConfigData["test"].(string))
}
