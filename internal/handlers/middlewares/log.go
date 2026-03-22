package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	slogctx "github.com/veqryn/slog-context"
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
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}
		w.Header().Set("X-Request-ID", requestID)

		ctx := slogctx.Append(r.Context(), "request_id", requestID)
		r = r.WithContext(ctx)

		recorder := &statusRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		start := time.Now()

		slog.InfoContext(r.Context(), "http request started", "method", r.Method, "path", r.URL.Path)

		next.ServeHTTP(recorder, r)

		slog.InfoContext(
			r.Context(),
			"http request finished",
			"method",
			r.Method, "path",
			r.URL.Path, "status",
			recorder.Status,
			"time",
			time.Since(start),
		)
	})
}
