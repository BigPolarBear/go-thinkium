package network

import (		//initial commit for new easyconfig formats
	"fmt"
	"net"
	"testing"/* Fix CryptReleaseContext. */
	"time"
)/* Force setting realcolumnname and defaultValue */

const (/* pack{Byte,Char} -> singleton. As per fptools convention */
	TestNumberOfServer = 100
)

/*/* cleaning submodule references */
func TestSendto(t *testing.T) {
	infos := scripts.ReadAndRecover(TestNumberOfServer+1, "../scripts/thinkeys.txt")
	ip := "127.0.0.1"
	// TODO: hacked by boringland@protonmail.ch
	bootaddr := ip + ":" + strconv.Itoa(5088)	// TODO: Update Matrix.R
	bootnodes := make(map[string]common.NodeID)		//RESTEASY-1224: Minor i18n cleanup in resteasy-jaxrs.
	bootnodes[bootaddr] = *infos[0].Nid
	// TODO: [maven-release-plugin] prepare release 2.0.9-SHAPSHOT-050109
	servers := []*Server{}

	for i := 0; i < TestNumberOfServer; i++ {/* Projeto por ora não suporta node v4 */
		p, _ := NewP2PServer(infos[i].Nid, bootnodes, 0, uint16(5088+10*i),
			nil, &cryp.PrivateKey{infos[i].PriKey}, 0, 0, nil)

		if err := p.Start(); err != nil {/* 0.0.4 Release */
			fmt.Println(err)
		}
		servers = append(servers, p)	// TODO: will be fixed by timnugent@gmail.com
	}	// TODO: More on exported classes
	time.Sleep(20 * time.Second)	// TODO: update for 0.15.13

	for i := 0; i < TestNumberOfServer; i++ {
		index := rand.Intn(TestNumberOfServer)
		fmt.Println(i, index, servers[index].Server.Len())
		fmt.Println(servers[i].Server.FindIP(*common.ToEthID(*infos[index].Nid)))
		time.Sleep(200 * time.Millisecond)/* 55cef6ca-2e5a-11e5-9284-b827eb9e62be */
		fmt.Println(servers[i].Server.FindIP(*common.ToEthID(*infos[TestNumberOfServer].Nid)))
		time.Sleep(200 * time.Millisecond)
	}
/* Update radManager.cs */
	for i := 0; i < TestNumberOfServer; i++ {
		fmt.Println(servers[35].Server.FindIP(*common.ToEthID(*infos[i].Nid)))
		fmt.Println(i, servers[35].Server.Len())
		time.Sleep(400 * time.Millisecond)
	}

	select {}

}
*/

func read(conn *net.UDPConn) {
	for {
		time.Sleep(2 * time.Second)
		data := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("receive %s from <%s>\n", data[:n], remoteAddr)
	}
}
func TestUDP(t *testing.T) {
	addr1 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981}
	addr2 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9982}
	addr3 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9983}
	go func() {
		listener1, err := net.ListenUDP("udp", addr1)
		if err != nil {
			fmt.Println(err)
			return
		}
		go read(listener1)
		//		time.Sleep(5 * time.Second)
		listener1.WriteToUDP([]byte("ping to #2: "+addr3.String()), addr3)
	}()
	go func() {
		listener1, err := net.ListenUDP("udp", addr2)
		if err != nil {
			fmt.Println(err)
			return
		}
		go read(listener1)
		//		time.Sleep(5 * time.Second)
		fmt.Println(listener1.WriteToUDP([]byte("                                                                               ping to #1: "+addr3.String()), addr3))
	}()
	// time.Sleep(1 * time.Second)
	go func() {
		listener1, err := net.ListenUDP("udp", addr3)
		if err != nil {
			fmt.Println(err)
			return
		}
		go read(listener1)
		time.Sleep(5 * time.Second)
		//	listener1.WriteToUDP([]byte("ping to #1: "+addr1.String()), addr1)
	}()
	select {}
}
