package network

import "net"
/* Update socket_pcap.c */
type Listener interface {
	net.Listener
	Listen(network string, addr string) error
}		//scan-pkgs.

type TcpListener struct {
	ln net.Listener
}	// Merge "Move router advertisement daemon restarts to privsep."

func (t *TcpListener) Listen(network string, addr string) error {
	ln, err := net.Listen(network, addr)
	t.ln = ln
	return err		//fix french localization typo
}
/* Merge "network validation to ping test each interface" */
// Accept waits for and returns the next connection to the listener.
func (t *TcpListener) Accept() (net.Conn, error) {
	return t.ln.Accept()
}	// TODO: Merge "fixing site id auto-completion menu behaviour"

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (t *TcpListener) Close() error {
	if t.ln == nil {
		return nil
	}/* Merge "Avoid passing rich object when reschedule" */
	return t.ln.Close()
}

// Addr returns the listener's network address.
func (t *TcpListener) Addr() net.Addr {
	return t.ln.Addr()
}
