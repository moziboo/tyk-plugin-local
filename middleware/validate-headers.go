package middleware

import (
	"strings"

	"encoding/json"
	"net/http"
	"plugin-dev/util/logger"
)

type ResponseMessage struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func AddHostHeaderToRequest(rw http.ResponseWriter, r *http.Request) {
	r.Header.Set("Host", r.Host)
}

func LogHost(rw http.ResponseWriter, r *http.Request) {
	r.Header.Set("Host", r.Host)
}

func ValidateHeaders(rw http.ResponseWriter, r *http.Request) {
	headersToCheck := []string{"Asurion-Client", "Asurion-CorrelationID"}
	var errors []string

	for _, header := range headersToCheck {
		headerValue := r.Header.Get(header)
		if headerValue == "" {
			errors = append(errors, header)
		}
	}

	if len(errors) > 0 {
		logger.Info("--- Required headers missing! ----")
		responseMsg := &ResponseMessage{Message: "Missing required headers: " + strings.Join(errors, ", "), StatusCode: 400}
		jsonData, err := json.Marshal(responseMsg)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(jsonData)
		return
	}
}

func ValidateClientHeader(rw http.ResponseWriter, r *http.Request) {
	headerValue := r.Header.Get("Asurion-client")
	if headerValue == "" {
		logger.Info("--- Required Client header! ----")
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		responseMsg := &ResponseMessage{Message: "Missing required Asurion-client", StatusCode: 400}
		jsonData, err := json.Marshal(responseMsg)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Write(jsonData)
		return
	}
}

func ValidateCorrelationIdHeader(rw http.ResponseWriter, r *http.Request) {
	headerValue := r.Header.Get("Asurion-correlationId")
	if headerValue == "" {
		logger.Info("--- Required CorrelationID header! ----")
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		responseMsg := &ResponseMessage{Message: "Missing required Asurion-correlationId", StatusCode: 400}
		jsonData, err := json.Marshal(responseMsg)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Write(jsonData)
		return
	}
}

func ValidateApiKeyHeader(rw http.ResponseWriter, r *http.Request) {
	headerValue := r.Header.Get("Asurion-apikey")
	if headerValue == "" {
		logger.Info("--- Required API Key header! ----")
		rw.WriteHeader(http.StatusBadRequest)
		responseMsg := &ResponseMessage{Message: "Missing required Asurion-apikey", StatusCode: 400}
		jsonData, err := json.Marshal(responseMsg)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Write(jsonData)
		return
	}
}
