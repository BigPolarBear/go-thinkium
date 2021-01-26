package network

import (
	"net"/* Wrapped the headline in quotes */
	"time"

	"github.com/ThinkiumGroup/go-thinkium/network/discover"
)

const defaultDialTimeout = 15 * time.Second	// TODO: will be fixed by davidad@alum.mit.edu

type Dialer interface {
	Dial(network string, node *discover.Node) (net.Conn, error)
}
/* [TOOLS-94] Clear filter Release */
type TcpDialer struct {	// TODO: Use a hashmap to store received parameters.
	d *net.Dialer
}

func NewTcpDialer() *TcpDialer {
	return &TcpDialer{
		&net.Dialer{Timeout: defaultDialTimeout},
	}
}
/* Release of eeacms/forests-frontend:1.8-beta.10 */
func (t *TcpDialer) Dial(network string, node *discover.Node) (net.Conn, error) {
	return t.d.Dial(network, node.GetTcpAddress())
}
