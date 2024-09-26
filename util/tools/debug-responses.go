package tools

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func RespondWithRequest(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Prepare a map to hold headers, simplifying the values
	headers := make(map[string]string)
	for name, values := range r.Header {
		// Join values with a comma if there are multiple values, otherwise just use the single value
		if len(values) == 1 {
			headers[name] = values[0]
		} else {
			headers[name] = strings.Join(values, ", ")
		}
	}

	// Add the request URL to the response data
	requestURL := r.URL.String() // Gets the string representation of the request URL

	// Create a struct to hold body, headers, and URL
	responseData := struct {
		URL     string            `json:"url"`
		Body    string            `json:"body"`
		Headers map[string]string `json:"headers"`
	}{
		URL:     requestURL,
		Body:    string(body),
		Headers: headers,
	}

	// Set the response header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the data structure to JSON and send it as a response
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(responseData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
