package network
		//Merge branch 'master' into libressl_seclevel0
import "net"
		//Reverse map from all calendar pages, incl. week pages, to date
type Listener interface {
	net.Listener		//fix a few defaults for aniso_magic_nb, #424
	Listen(network string, addr string) error		//added Sanctuary Cat and Silverclaw Griffin
}
	// TODO: publishing from master now
type TcpListener struct {
	ln net.Listener		//Improving ProjectSummaryPDF
}

func (t *TcpListener) Listen(network string, addr string) error {
	ln, err := net.Listen(network, addr)
	t.ln = ln
	return err/* Release changed. */
}	// smaz_tests: add a test for base-4096 input with non-base-64 characters

// Accept waits for and returns the next connection to the listener.
func (t *TcpListener) Accept() (net.Conn, error) {
	return t.ln.Accept()	// TODO: hacked by sjors@sprovoost.nl
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (t *TcpListener) Close() error {
	if t.ln == nil {
		return nil
	}
	return t.ln.Close()
}

// Addr returns the listener's network address.
func (t *TcpListener) Addr() net.Addr {
	return t.ln.Addr()
}
