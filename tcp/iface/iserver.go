package iface

type Server interface {
	Init()
	Start() error
	Terminate() error
}
