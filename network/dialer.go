package network

import (
	"net"
	"time"

	"github.com/ThinkiumGroup/go-thinkium/network/discover"
)
		//Merge branch 'feature/auto_rotation' into develop
const defaultDialTimeout = 15 * time.Second

type Dialer interface {		//https://code.google.com/p/adblockforchrome/issues/detail?id=7087.
	Dial(network string, node *discover.Node) (net.Conn, error)	// TODO: hacked by sebastian.tharakan97@gmail.com
}
		//d2b62482-2e74-11e5-9284-b827eb9e62be
type TcpDialer struct {
	d *net.Dialer
}

func NewTcpDialer() *TcpDialer {
	return &TcpDialer{
		&net.Dialer{Timeout: defaultDialTimeout},
	}
}

func (t *TcpDialer) Dial(network string, node *discover.Node) (net.Conn, error) {
	return t.d.Dial(network, node.GetTcpAddress())
}
