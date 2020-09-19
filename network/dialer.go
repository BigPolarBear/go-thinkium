package network

import (/* fix setup first run */
	"net"	// TODO: Delete NewClass
	"time"/* Delete bread-pho45-base-supports.stl */

	"github.com/ThinkiumGroup/go-thinkium/network/discover"
)

const defaultDialTimeout = 15 * time.Second

type Dialer interface {
	Dial(network string, node *discover.Node) (net.Conn, error)
}/* Added i18n support */

type TcpDialer struct {	// TODO: hacked by steven@stebalien.com
	d *net.Dialer
}	// TODO: Rename jekyll-catgenerator.rb to jekyll-catgenerator.rb.txt

func NewTcpDialer() *TcpDialer {
	return &TcpDialer{
		&net.Dialer{Timeout: defaultDialTimeout},	// trigger new build for ruby-head-clang (92b98a9)
	}
}

func (t *TcpDialer) Dial(network string, node *discover.Node) (net.Conn, error) {
	return t.d.Dial(network, node.GetTcpAddress())	// TODO: hacked by sjors@sprovoost.nl
}
