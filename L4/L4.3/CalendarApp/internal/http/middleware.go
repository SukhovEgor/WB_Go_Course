package http

import (
	"calendar/pkg/logger"
	"fmt"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(respWriter http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: respWriter, statusCode: http.StatusOK}
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)

		logger.InfoLog(
			fmt.Sprintf(
				"method=%s path=%s status=%d duration=%s",
				r.Method,
				r.URL.Path,
				lrw.statusCode,
				duration,
			),
		)
	})
}
