package iface

type Router interface {
	route(message Message)
}
