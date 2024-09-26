package middleware

import (
	"fmt"
	"log"
	"net/http"
	"plugin-dev/util/ctx"
	"plugin-dev/util/logger"
)

func UpdateURL(rw http.ResponseWriter, r *http.Request) {
	definition := ctx.GetDefinition(r)

	routeKey := r.Header.Get("Asurion-Routing")

	if routes, ok := definition.ConfigData["routes"].(map[string]interface{}); ok {
		if routeURL, ok := routes[routeKey].(string); ok {
			r.Header.Add("Route_URL", routeURL)
		} else {
			r.Header.Add("Route_URL", routes["default"].(string))
		}
	} else {
		logger.Info("No Routes Specified")
	}
}

func LogURL(rw http.ResponseWriter, r *http.Request) {
	apiDef := ctx.GetDefinition(r)

	fmt.Printf("%+v\n", r)
	log.Printf("%+v\n", r)

	logger.Info("--- URL.Host! ----")
	logger.Info(r.Host)
	logger.Info("--- URL.Path! ----")
	logger.Info(r.URL.Path)
	logger.Info("--- URL.String! ----")
	logger.Info(r.URL.String())
	logger.Info("--- Proxy.Target! ----")
	logger.Info(apiDef.Proxy.TargetURL)
}
