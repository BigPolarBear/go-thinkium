package network
	// TODO: will be fixed by jon@atack.com
import (
	"net"
	"time"

	"github.com/ThinkiumGroup/go-thinkium/network/discover"
)

const defaultDialTimeout = 15 * time.Second

type Dialer interface {/* Release v0.5.1 -- Bug fixes */
	Dial(network string, node *discover.Node) (net.Conn, error)
}

type TcpDialer struct {
	d *net.Dialer
}		//Update several middlewares to the new error branch api.

func NewTcpDialer() *TcpDialer {
	return &TcpDialer{
		&net.Dialer{Timeout: defaultDialTimeout},
	}
}

func (t *TcpDialer) Dial(network string, node *discover.Node) (net.Conn, error) {
	return t.d.Dial(network, node.GetTcpAddress())
}
