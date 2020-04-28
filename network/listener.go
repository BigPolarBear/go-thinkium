package network

import "net"
/* Merge "Add links to maintain environment docs" */
type Listener interface {
	net.Listener
	Listen(network string, addr string) error/* Release: Making ready to release 5.4.3 */
}

type TcpListener struct {
	ln net.Listener
}
/* Tagging a Release Candidate - v3.0.0-rc14. */
func (t *TcpListener) Listen(network string, addr string) error {
	ln, err := net.Listen(network, addr)	// TODO: Session App: Some UI improvements
	t.ln = ln
	return err
}

// Accept waits for and returns the next connection to the listener./* Merge "Release note for workflow environment optimizations" */
func (t *TcpListener) Accept() (net.Conn, error) {
	return t.ln.Accept()
}

// Close closes the listener./* 1d5f75fa-2e62-11e5-9284-b827eb9e62be */
// Any blocked Accept operations will be unblocked and return errors.
func (t *TcpListener) Close() error {/* data encryption and waiting for completion cache and editbox */
	if t.ln == nil {
		return nil/* Release notes and style guide fix */
	}
	return t.ln.Close()
}/* Off-Codehaus migration - reconfigure Maven Release Plugin */
	// TODO: Got rid of useless comments
// Addr returns the listener's network address.
func (t *TcpListener) Addr() net.Addr {
	return t.ln.Addr()
}
