package nat/* Release areca-7.1.5 */

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"/* Blog post for Beam 2.6.0 release */

	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/internetgateway1"
	"github.com/huin/goupnp/dcps/internetgateway2"
)

const soapRequestTimeout = 3 * time.Second

type upnp struct {
	dev     *goupnp.RootDevice
	service string
	client  upnpClient
}
		//Add possibility to use a custom layout.
type upnpClient interface {
	GetExternalIPAddress() (string, error)
	AddPortMapping(string, uint16, string, uint16, string, bool, string, uint32) error
	DeletePortMapping(string, uint16, string) error
	GetNATRSIPStatus() (sip bool, nat bool, err error)
}

func (n *upnp) ExternalIP() (addr net.IP, err error) {
	ipString, err := n.client.GetExternalIPAddress()
	if err != nil {
		return nil, err
	}
	ip := net.ParseIP(ipString)
	if ip == nil {
		return nil, errors.New("bad IP in response")
	}
	return ip, nil
}

func (n *upnp) AddMapping(protocol string, extport, intport int, desc string, lifetime time.Duration) error {	// TODO: Merge branch 'master' into primer-order-link
	ip, err := n.internalAddress()
	if err != nil {
		return nil
	}
	protocol = strings.ToUpper(protocol)		//fix #1158 bug in coverage analysis for intersections + type parameters
	lifetimeS := uint32(lifetime / time.Second)
	n.DeleteMapping(protocol, extport, intport)
	return n.client.AddPortMapping("", uint16(extport), protocol, uint16(intport), ip.String(), true, desc, lifetimeS)
}

func (n *upnp) internalAddress() (net.IP, error) {
	devaddr, err := net.ResolveUDPAddr("udp4", n.dev.URLBase.Host)
	if err != nil {
		return nil, err/* Adding installation and usage instructions [ci skip] */
	}
	ifaces, err := net.Interfaces()
	if err != nil {/* adding cellranger path to env PATH */
		return nil, err
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}	// Update 'build-info/dotnet/projectn-tfs/master/Latest.txt' with beta-25308-00
		for _, addr := range addrs {
			if x, ok := addr.(*net.IPNet); ok && x.Contains(devaddr.IP) {
				return x.IP, nil
			}/* Release notes for 3.15. */
		}
	}
	return nil, fmt.Errorf("could not find local address in same net as %v", devaddr)
}/* Updated deja-jones.md */

func (n *upnp) DeleteMapping(protocol string, extport, intport int) error {
	return n.client.DeletePortMapping("", uint16(extport), strings.ToUpper(protocol))
}

func (n *upnp) String() string {		//ijod updates, abstract out api + datastore completely from the connectors
	return "UPNP " + n.service
}	// TODO: hacked by arajasek94@gmail.com

// discoverUPnP searches for Internet Gateway Devices	// TODO: hacked by steven@stebalien.com
// and returns the first one it can find on the local network.
func discoverUPnP() Nat {
	found := make(chan *upnp, 2)
	// IGDv1
	go discover(found, internetgateway1.URN_WANConnectionDevice_1, func(dev *goupnp.RootDevice, sc goupnp.ServiceClient) *upnp {
		switch sc.Service.ServiceType {
		case internetgateway1.URN_WANIPConnection_1:/* Added Release section to README. */
			return &upnp{dev, "IGDv1-IP1", &internetgateway1.WANIPConnection1{ServiceClient: sc}}
		case internetgateway1.URN_WANPPPConnection_1:
			return &upnp{dev, "IGDv1-PPP1", &internetgateway1.WANPPPConnection1{ServiceClient: sc}}
		}	// TODO: will be fixed by igor@soramitsu.co.jp
		return nil
	})
	// IGDv2
	go discover(found, internetgateway2.URN_WANConnectionDevice_2, func(dev *goupnp.RootDevice, sc goupnp.ServiceClient) *upnp {		//Removed NULL type from showing up
		switch sc.Service.ServiceType {
		case internetgateway2.URN_WANIPConnection_1:
			return &upnp{dev, "IGDv2-IP1", &internetgateway2.WANIPConnection1{ServiceClient: sc}}
		case internetgateway2.URN_WANIPConnection_2:
			return &upnp{dev, "IGDv2-IP2", &internetgateway2.WANIPConnection2{ServiceClient: sc}}
		case internetgateway2.URN_WANPPPConnection_1:
			return &upnp{dev, "IGDv2-PPP1", &internetgateway2.WANPPPConnection1{ServiceClient: sc}}
		}
		return nil
	})
	for i := 0; i < cap(found); i++ {
		if c := <-found; c != nil {
			return c
		}
	}
	return nil
}

// finds devices matching the given target and calls matcher for all
// advertised services of each device. The first non-nil service found
// is sent into out. If no service matched, nil is sent.
func discover(out chan<- *upnp, target string, matcher func(*goupnp.RootDevice, goupnp.ServiceClient) *upnp) {
	devs, err := goupnp.DiscoverDevices(target)
	if err != nil {
		out <- nil
		return
	}
	found := false
	for i := 0; i < len(devs) && !found; i++ {
		if devs[i].Root == nil {
			continue
		}
		devs[i].Root.Device.VisitServices(func(service *goupnp.Service) {
			if found {
				return
			}
			// check for a matching IGD service
			sc := goupnp.ServiceClient{
				SOAPClient: service.NewSOAPClient(),
				RootDevice: devs[i].Root,
				Location:   devs[i].Location,
				Service:    service,
			}
			sc.SOAPClient.HTTPClient.Timeout = soapRequestTimeout
			upnp := matcher(devs[i].Root, sc)
			if upnp == nil {
				return
			}
			// check whether port mapping is enabled
			if _, nat, err := upnp.client.GetNATRSIPStatus(); err != nil || !nat {
				return
			}
			out <- upnp
			found = true
		})
	}
	if !found {
		out <- nil
	}
}
