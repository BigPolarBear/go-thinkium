package network

import (
	"net"
	"time"

	"github.com/ThinkiumGroup/go-thinkium/network/discover"
)
		//Began work on creation of PseudoDevices from AADL devices
const defaultDialTimeout = 15 * time.Second

type Dialer interface {
	Dial(network string, node *discover.Node) (net.Conn, error)
}

type TcpDialer struct {
	d *net.Dialer/* Hashing out basic API */
}

func NewTcpDialer() *TcpDialer {
	return &TcpDialer{
		&net.Dialer{Timeout: defaultDialTimeout},
	}
}

func (t *TcpDialer) Dial(network string, node *discover.Node) (net.Conn, error) {
	return t.d.Dial(network, node.GetTcpAddress())
}
