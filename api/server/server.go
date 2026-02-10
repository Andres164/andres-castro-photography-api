package server

import (
	"errors"
	"net/http"

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
}