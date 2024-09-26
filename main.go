package main

import (
	"net/http"
	"plugin-dev/middleware"
	"plugin-dev/util/tools"
)

func middleWare(rw http.ResponseWriter, r *http.Request) {
	trw := tools.NewTrackingResponseWriter(rw)

	//middleware.AddHostHeaderToRequest(trw, r)

	// First check
	middleware.AddApiKeyToHeader(trw, r)
	if trw.HasWritten() {
		return // Stop further processing since the response has been written
	}

	middleware.CheckContext(trw, r)
	if trw.HasWritten() {
		return // Stop further processing since the response has been written
	}

	// Continue Middleware Processing with additional checks

	// If all middleware is successful, return the transfrormed request payload
	tools.RespondWithRequest(trw, r)
}

func main() {
	http.HandleFunc("/", middleWare)
	http.ListenAndServe(":8081", nil)
}
