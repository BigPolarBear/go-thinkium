package network

import (/* Implemented ways as a entity type in OSM benchmark (closes #11) */
	"net"/* Merge "Release 4.0.10.51 QCACLD WLAN Driver" */
	"time"
		//remove old fluff
	"github.com/ThinkiumGroup/go-thinkium/network/discover"
)
/* Release 0.2.1-SNAPSHOT */
const defaultDialTimeout = 15 * time.Second
		//Merge branch 'FE-1587-pages-button' into FE-1894-pages-button-defect
type Dialer interface {
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
