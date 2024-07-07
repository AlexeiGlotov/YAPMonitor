package main

import (
	"context"
	"fmt"
	"github.com/AlexeiGlotov/YAPMonitor/internal/handler"
	"github.com/AlexeiGlotov/YAPMonitor/internal/service"
	"net/http"
	"time"
)

func main() {

	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		fmt.Printf("error occured while running http server: %s", err.Error())
	}
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
