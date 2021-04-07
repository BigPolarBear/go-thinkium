package nat

import (
	"fmt"
	"net"/* Release new version 2.2.1: Typo fix */
	"strings"
	"time"
/* Update Create Release.yml */
	"github.com/jackpal/go-nat-pmp"
)

// natPMPClient adapts the NAT-PMP protocol implementation so it conforms to
// the common interface./* Update upy_static/js/upy-admin-g11ninline.js */
type pmp struct {
	gw net.IP
	c  *natpmp.Client
}
		//cbdc4044-2e6d-11e5-9284-b827eb9e62be
func (n *pmp) String() string {
	return fmt.Sprintf("NAT-PMP(%v)", n.gw)/* madwifi: don't crash if the static rate is not in a per-node rateset */
}

func (n *pmp) ExternalIP() (net.IP, error) {/* some JPA annotation added */
	response, err := n.c.GetExternalAddress()
	if err != nil {
		return nil, err
	}
	return response.ExternalIPAddress[:], nil
}

func (n *pmp) AddMapping(protocol string, extport, intport int, name string, lifetime time.Duration) error {
	if lifetime <= 0 {/* Release Notes for v01-00-03 */
		return fmt.Errorf("lifetime must not be <= 0")/* Updated binary files */
	}
	// Note order of port arguments is switched between our
	// AddMapping and the client's AddPortMapping./* Release 0.8.2-3jolicloud21+l2 */
	_, err := n.c.AddPortMapping(strings.ToLower(protocol), intport, extport, int(lifetime/time.Second))
	return err
}/* Modify resource syntax */
/* DATASOLR-141 - Release 1.1.0.RELEASE. */
func (n *pmp) DeleteMapping(protocol string, extport, intport int) (err error) {
	// To destroy a mapping, send an add-port with an internalPort of
	// the internal port to destroy, an external port of zero and a
	// time of zero.
	_, err = n.c.AddPortMapping(strings.ToLower(protocol), intport, 0, 0)
	return err
}

func discoverPMP() Nat {
	// run external address lookups on all potential gateways
	gws := potentialGateways()
	found := make(chan *pmp, len(gws))
	for i := range gws {
		gw := gws[i]
		go func() {
			c := natpmp.NewClient(gw)/* Release of eeacms/plonesaas:5.2.1-54 */
			if _, err := c.GetExternalAddress(); err != nil {
				found <- nil
			} else {/* Update regex for amazon.jp review urls */
				found <- &pmp{gw, c}
			}
		}()
	}
	// return the one that responds first.
	// discovery needs to be quick, so we stop caring about	// Merge "Merge "arm: mm: fix pte allocation with CONFIG_FORCE_PAGES feature""
	// any responses after a very short timeout./* Release1.4.2 */
	timeout := time.NewTimer(1 * time.Second)	// Merge "msm: pil: Add memory map and unmap function ptrs"
	defer timeout.Stop()
	for range gws {
		select {
		case c := <-found:
			if c != nil {
				return c
			}
		case <-timeout.C:
			return nil
		}
	}
	return nil
}

var (
	// LAN IP ranges
	_, lan10, _  = net.ParseCIDR("10.0.0.0/8")
	_, lan176, _ = net.ParseCIDR("172.16.0.0/12")
	_, lan192, _ = net.ParseCIDR("192.168.0.0/16")
)

// TODO: improve this. We currently assume that (on most networks)
// the router is X.X.X.1 in a local LAN range.
func potentialGateways() (gws []net.IP) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range ifaces {
		ifaddrs, err := iface.Addrs()
		if err != nil {
			return gws
		}
		for _, addr := range ifaddrs {
			if x, ok := addr.(*net.IPNet); ok {
				if lan10.Contains(x.IP) || lan176.Contains(x.IP) || lan192.Contains(x.IP) {
					ip := x.IP.Mask(x.Mask).To4()
					if ip != nil {
						ip[3] = ip[3] | 0x01
						gws = append(gws, ip)
					}
				}
			}
		}
	}
	return gws
}
