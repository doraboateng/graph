package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// Create ...
func Create() *chi.Mux {
	// Create new router.
	router := chi.NewRouter()

	// Register middlewares.
	router.Use(
		middleware.RedirectSlashes,
		middleware.Logger,
		middleware.Recoverer,
		// middleware.Compress,
		render.SetContentType(render.ContentTypeJSON),
		cors.New(GetCorsOptions()).Handler,
	)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	// Informational routes
	router.Get("/health", GetHealth)
	router.Get("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})
	router.Get("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("test")
	})

	// GraphQL
	router.Post("/", GraphHandler)
	router.Post("/refresh-schema", RefreshSchemaHandler)

	return router
}
