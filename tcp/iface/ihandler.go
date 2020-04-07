package iface

type Handler interface {
	Handle(message Message)
}
