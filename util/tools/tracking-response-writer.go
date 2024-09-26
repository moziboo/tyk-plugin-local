package tools

import (
	"net/http"
)

// trackingResponseWriter wraps around http.ResponseWriter to track whether
// headers or body have been written to the HTTP response.
type TrackingResponseWriter struct {
	http.ResponseWriter
	WroteHeader bool
	WroteBody   bool
}

// WriteHeader tracks the call to write headers and forwards to the original ResponseWriter.
func (trw *TrackingResponseWriter) WriteHeader(statusCode int) {
	trw.WroteHeader = true
	trw.ResponseWriter.WriteHeader(statusCode)
}

// Write tracks the call to write the body and forwards to the original ResponseWriter.
func (trw *TrackingResponseWriter) Write(data []byte) (int, error) {
	if !trw.WroteBody {
		trw.WroteBody = true
	}
	return trw.ResponseWriter.Write(data)
}

func (trw *TrackingResponseWriter) HasWritten() bool {
	return trw.WroteHeader || trw.WroteBody
}

// NewTrackingResponseWriter creates a new instance of TrackingResponseWriter.
func NewTrackingResponseWriter(w http.ResponseWriter) *TrackingResponseWriter {
	return &TrackingResponseWriter{ResponseWriter: w}
}
