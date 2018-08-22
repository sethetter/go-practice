package main

import (
	"io"
	"net"
)

type Mux struct {
	ops chan func(map[net.Addr]net.Conn)
}

func (m *Mux) Add(conn net.Conn) {
	m.ops <- func(conns map[net.Addr]net.Conn) {
		conns[conn.RemoteAddr()] = conn
	}
}

func (m *Mux) Remove(addr net.Addr) {
	m.ops <- func(conns map[net.Addr]net.Conn) {
		delete(conns, addr)
	}
}

func (m *Mux) SendMsg(msg string) error {
	m.ops <- func(conns map[net.Addr]net.Conn) {
		for _, conn := range conns {
			io.WriteString(conn, msg)
		}
	}
	return nil
}

func (m *Mux) loop() {
	conns := make(map[net.Addr]net.Conn)
	for op := range m.ops {
		op(conns)
	}
}
