package iface

type Sender interface {
	Init()
	Start()
	Terminate()
	Send()
}
