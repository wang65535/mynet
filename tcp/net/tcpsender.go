package net

import (
	"../iface"
)

type SenderConfig struct {
	MaxLength int
}
type Sender struct {
	jobs chan iface.Message
	stop chan bool
}

func (s *Sender) Init(c SenderConfig) {
	s.jobs = make(chan iface.Message, c.MaxLength)
}

func (s *Sender) Start() {
	go func() {
		for {
			select {
			case b := <-s.stop:
				if !b {
					return
				}
				break
			case m := <-s.jobs:
				go s.send(m)
				break
			default:
			}

		}
	}()
}

func (s *Sender) Close() {
	s.stop <- false
}

func (s *Sender) send(message iface.Message) {
	conn := message.GetConn()
	_, err := conn.Write(message.GetData())
	if err != nil {

	}
}

func (s *Sender) Send(message iface.Message) {
	s.jobs <- message
}
