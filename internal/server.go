package internal

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	Cfg    *Config
	Logger *zap.Logger
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Logger.Sugar().Infof("Received: %s request.", req.Method)
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, "Hello, world!\n")
	if err != nil {
		s.Logger.Sugar().Error(err)
	}
}

func NewServer(path string) (*Server, error) {
	if path == "" {
		return nil, fmt.Errorf("path to a config file was not specified")
	}
	cfg, err := InitConfig(path)
	if err != nil {
		return nil, err
	}
	l, err := InitLogger(cfg.LogLevel)
	if err != nil {
		return nil, err
	}
	return &Server{Cfg: cfg, Logger: l}, nil
}

func (s *Server) Run() {
	s.Logger.Sugar().Info("Server started")
	s.Logger.Sugar().Errorf("%v", http.ListenAndServe(s.Cfg.Host+":"+s.Cfg.Port, s))
}
