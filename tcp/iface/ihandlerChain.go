package iface

type HandlerChain interface {
	Add(handler Handler)
	Remove(handler Handler)
}
