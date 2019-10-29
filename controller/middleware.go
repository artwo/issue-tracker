package controller

import "net/http"

const TRACE_HEADER = "X-Correlation-ID"

func RequestTraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if w.Header().Get(TRACE_HEADER) == "" {
			w.Header().Set(TRACE_HEADER, "traceid")
		} else {
			w.Header().Set(TRACE_HEADER, "traceid")
		}
		next.ServeHTTP(w, r)
	})
}
