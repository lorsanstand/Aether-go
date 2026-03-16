package middlewares

import (
	"log"
	"net/http"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("start: %v %v", r.Method, r.URL)

		recorder := &statusRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}
		start := time.Now()

		next.ServeHTTP(recorder, r)

		duration := time.Since(start)
		log.Printf("end: %v, %v, status: %v time: %d.%d ms",
			r.Method,
			r.URL,
			recorder.Status,
			duration.Milliseconds(),
			duration.Microseconds())
	})
}
