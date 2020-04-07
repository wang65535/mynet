package iface

import "net"

type Register interface {
	Regist(conn net.Conn)
	Unregist(conn net.Conn)
}
