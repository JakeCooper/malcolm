package proxy

import (
	"net"
)

func New(from net.Conn, to net.Conn) *Proxy {
	return &Proxy{
		From: from,
		To:   to,
		End:  make(chan bool),
	}
}

type Proxy struct {
	From net.Conn
	To   net.Conn
	End  chan (bool)
}

func (p *Proxy) monoconn(connFrom net.Conn, connTo net.Conn) {
	buff := make([]byte, 0xffff) // 64kb buffer
	for {
		n, err := connFrom.Read(buff)
		if err != nil {
			// Includes io.EOF
			p.End <- true
		}
		b := buff[:n]
		n, err = connTo.Write(b)
	}
}

func (p *Proxy) Proxy() {
	defer p.From.Close()
	go p.monoconn(p.From, p.To)
	go p.monoconn(p.To, p.From)
	<-p.End
}
