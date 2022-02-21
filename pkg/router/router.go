package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"link-storage-api/pkg/logger"
	"net"
	"net/http"
	"time"
)

func Router(logger *logger.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	//r.Use(middleware.Logger) // internal chi logger
	r.Use(loggerMiddleware(logger)) // custom logrus logger
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return r
}

func loggerMiddleware(logger *logger.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			reqID := middleware.GetReqID(r.Context())
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			defer func() {
				remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					remoteIP = r.RemoteAddr
				}
				scheme := "http"
				if r.TLS != nil {
					scheme = "https"
				}
				fields := logrus.Fields{
					"status_code": ww.Status(),
					"bytes":       ww.BytesWritten(),
					"duration":    time.Since(t1).String(),
					"remote_ip":   remoteIP,
					"proto":       r.Proto,
					"method":      r.Method,
				}
				if len(reqID) > 0 {
					fields["request_id"] = reqID
				}
				logger.WithFields(fields).Infof("%s://%s%s", scheme, r.Host, r.RequestURI)
			}()

			h.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}
