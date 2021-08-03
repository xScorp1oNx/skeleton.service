package routing

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"skeleton.service/cars/handlers"
	_ "skeleton.service/docs"
	"time"
)

var srv *http.Server

func InitHttpServer(port string) error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Mount("/api/doc", httpSwagger.WrapHandler)

	//router.Options("/api/*", optionsHandler) // here you can set options handlers for CORS

	protectedGroup := router.Group(nil)
	//protectedGroup.Use(jwtMiddleware.Handler) // here you can set JWT middleware handler for verify Bearer JWT token

	protectedGroup.Route("/api/{api_version}/car", func(r chi.Router) {
		r.Get("/", MakeHandlesChain(appendTracingContext, handlers.GetCar))   // you can use cors middleware handler for set headers to Response
		r.Post("/", MakeHandlesChain(appendTracingContext, handlers.PostCar)) // you can use cors middleware handler for set headers to Response
	})

	srv = &http.Server{
		Addr:              "0.0.0.0:" + port,
		Handler:           router,
		ReadTimeout:       300 * time.Second,
		ReadHeaderTimeout: 300 * time.Second,
		WriteTimeout:      300 * time.Second,
		IdleTimeout:       300 * time.Second,
	}

	log.Printf("Running http server on port %s", port)

	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
