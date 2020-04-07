package iface

import "net"

type Message interface {
	GetConn() net.Conn
	GetContext() MessageContext
	GetData() []byte
}
