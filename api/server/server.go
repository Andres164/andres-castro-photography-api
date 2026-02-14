package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andres164/andres-castro-photography/configs"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	logger zerolog.Logger
	router *gin.Engine
	config *configs.Config
}

func NewServer(logger zerolog.Logger, router *gin.Engine, config *configs.Config) *Server {
	return &Server{logger: logger, router: router, config: config}
}

func (s *Server) Serve() {
	srv := &http.Server{
		Addr: s.config.Server.Address,
		Handler: s.router.Handler(),
	}

	go func ()  {
		// service connections
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			s.logger.Fatal().Err(err).Msg("listen")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.logger.Info().Msg("Sutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Fatal().Err(err).Msg("Server Shutdown")
	}

	<-ctx.Done()
	s.logger.Info().Msg("Server shutdown timeout of 30 secods")
	s.logger.Info().Msg("Server exiting")
}