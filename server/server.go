package server

import (
	"context"
	"github.com/lisuiheng/esm/log"
	"github.com/lisuiheng/fsm"
	flog "github.com/lisuiheng/fsm/log"
	"net"
	"net/http"
	"strconv"
	"time"
)

var (
	startTime time.Time
)

func init() {
	startTime = time.Now().UTC()
}

// Server for the chronograf API
type Server struct {
	Host string `short:"h" long:"host" description:"The IP to listen on" default:"0.0.0.0" env:"HOST"`
	Port int    `short:"p" long:"port" description:"The port to listen on for insecure connections," default:"8000" env:"PORT"`

	Signature string `short:"s" long:"signature" description:"Secret to sign tokens " default:"0dl0mOxyQKVl8Hwo" env:"SIGNATURE"`

	ShowVersion bool   `short:"v" long:"version" description:"Show FSM version info"`
	LogLevel    string `short:"l" long:"log-level" value-name:"choice" choice:"debug" choice:"info" choice:"error" default:"info" description:"Set the logging level" env:"LOG_LEVEL"`

	Version  string
	Listener net.Listener
	handler  http.Handler
}

// Serve starts and runs the fsm server
func (s *Server) Serve(ctx context.Context) error {
	logger := flog.New(flog.ParseLevel(s.LogLevel))
	service := openService(ctx, logger)

	s.handler = NewMux(MuxOpts{
		Logger:    logger,
		Signature: s.Signature,
	}, service)

	// Add chronograf's version header to all requests
	s.handler = Version(s.Version, s.handler)

	listener, err := s.NewListener()
	if err != nil {
		log.Error(err)
		return err
	}
	s.Listener = listener

	scheme := "http"

	logger.
		Info("Serving chronograf at ", scheme, "://", s.Listener.Addr())

	if err := http.Serve(s.Listener, s.handler); err != nil {
		log.Error(err)
		return err
	}

	log.Infow("component", "server", "msg", "Stopped serving fsm at ")

	return nil
}

func openService(ctx context.Context, logger fsm.Logger) Service {
	return Service{
		Logger: logger,
	}
}

// NewListener will an http or https listener depending useTLS()
func (s *Server) NewListener() (net.Listener, error) {
	addr := net.JoinHostPort(s.Host, strconv.Itoa(s.Port))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return listener, nil
}
