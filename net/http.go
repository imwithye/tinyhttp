package net

import (
	"fmt"
	"net"
	"net/http"
	"tinyhttp/log"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode("release")
	gin.DisableConsoleColor()
}

type Server struct {
	*http.Server

	tcp net.Listener
	dir string
}

func (s *Server) URL() string {
	return fmt.Sprintf("http://%s:%d", s.Host(), s.Port())
}

func (s *Server) Host() net.IP {
	return s.tcp.Addr().(*net.TCPAddr).IP
}

func (s *Server) Port() int {
	return s.tcp.Addr().(*net.TCPAddr).Port
}

func (s *Server) Dir() string {
	return s.dir
}

func (s *Server) Serve() error {
	log.Info(fmt.Sprintf("Serving %s", s.Dir()))
	log.Info(fmt.Sprintf("Listening on HTTP port %s", s.URL()))
	fmt.Printf("\n\n")
	return s.Server.Serve(s.tcp)
}

func NewServer(host net.IP, port int, dir string) (*Server, error) {
	log.Trace("create tcp listener, looking for available port")
	tcp, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	router := gin.Default()
	router.StaticFS("/", http.Dir(dir))
	server := &Server{&http.Server{Handler: router}, tcp, dir}
	return server, nil
}
