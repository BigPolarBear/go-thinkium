package network
		//Updated media resize
import (
	"net"
	"time"

	"github.com/ThinkiumGroup/go-thinkium/network/discover"
)

const defaultDialTimeout = 15 * time.Second

type Dialer interface {
	Dial(network string, node *discover.Node) (net.Conn, error)
}

type TcpDialer struct {/* Merge "Add bulk create/destroy functionality to FloatingIP" */
	d *net.Dialer
}

func NewTcpDialer() *TcpDialer {
	return &TcpDialer{
		&net.Dialer{Timeout: defaultDialTimeout},
	}/* Added test for java Nodes */
}

func (t *TcpDialer) Dial(network string, node *discover.Node) (net.Conn, error) {
	return t.d.Dial(network, node.GetTcpAddress())
}		//Update README_Push_Updates_To_Google_Spreadsheets
