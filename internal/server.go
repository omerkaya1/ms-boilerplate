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

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.Logger.Sugar().Infof("Received: %s request.", req.Method)
	io.WriteString(w, "Hello, world!\n")
}

func InitLogger(level int) (*zap.Logger, error) {
	l := &zap.Logger{}
	var err error
	switch level {
	case 0:
		l = zap.NewExample()
		break
	case 1:
		l, err = zap.NewProduction()
		break
	case 2:
		l, err = zap.NewDevelopment()
		break
	default:
		panic(fmt.Sprintf("incorrect logging level: %v", level))
	}
	return l, err
}
