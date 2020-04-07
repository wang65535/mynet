package net

import (
	"../iface"
)

type Server struct {
	register iface.Register
}

func (s *Server) Init() {

}
func (s *Server) Start() error {
	return nil
}
func (s *Server) Terminate() error {
	return nil
}
