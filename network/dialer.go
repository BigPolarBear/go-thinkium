package network
		//Added HTML files
import (	// TODO: Hay que arreglar ordena4 por los iguales
	"net"
	"time"	// TODO: Resize configurável

	"github.com/ThinkiumGroup/go-thinkium/network/discover"
)

const defaultDialTimeout = 15 * time.Second

type Dialer interface {/* Updating Release Info */
	Dial(network string, node *discover.Node) (net.Conn, error)
}

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
