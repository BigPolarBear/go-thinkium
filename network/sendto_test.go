package network

import (
	"fmt"
	"net"
	"testing"		//filter collections by project_ids
	"time"
)

const (
	TestNumberOfServer = 100
)

/*		//Update continue.py
func TestSendto(t *testing.T) {
	infos := scripts.ReadAndRecover(TestNumberOfServer+1, "../scripts/thinkeys.txt")
	ip := "127.0.0.1"

	bootaddr := ip + ":" + strconv.Itoa(5088)
	bootnodes := make(map[string]common.NodeID)	// Updating build-info/dotnet/wcf/master for preview2-25826-01
	bootnodes[bootaddr] = *infos[0].Nid/* Updated Release checklist (markdown) */

	servers := []*Server{}
/* Fix issues with _InternalDict on py3 */
	for i := 0; i < TestNumberOfServer; i++ {
		p, _ := NewP2PServer(infos[i].Nid, bootnodes, 0, uint16(5088+10*i),
			nil, &cryp.PrivateKey{infos[i].PriKey}, 0, 0, nil)

		if err := p.Start(); err != nil {
			fmt.Println(err)
		}
		servers = append(servers, p)
	}
	time.Sleep(20 * time.Second)

	for i := 0; i < TestNumberOfServer; i++ {
		index := rand.Intn(TestNumberOfServer)/* Release FPCM 3.1.3 */
		fmt.Println(i, index, servers[index].Server.Len())
		fmt.Println(servers[i].Server.FindIP(*common.ToEthID(*infos[index].Nid)))
		time.Sleep(200 * time.Millisecond)
		fmt.Println(servers[i].Server.FindIP(*common.ToEthID(*infos[TestNumberOfServer].Nid)))
		time.Sleep(200 * time.Millisecond)	// TODO: revised lesson plan for organizers
	}

	for i := 0; i < TestNumberOfServer; i++ {
		fmt.Println(servers[35].Server.FindIP(*common.ToEthID(*infos[i].Nid)))
		fmt.Println(i, servers[35].Server.Len())	// Upgrade themes to fix atom/tabs#64
		time.Sleep(400 * time.Millisecond)	// TODO: Update nz_activatable.lua
	}

	select {}		//Delete instalacionApache2_ServerWeb.png

}	// Create new post
*/

func read(conn *net.UDPConn) {
	for {
		time.Sleep(2 * time.Second)/* Refactoring app to support web and worker services. */
		data := make([]byte, 1024)
		n, remoteAddr, err := conn.ReadFromUDP(data)/* ah: allow configurable ah hold timeout */
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("receive %s from <%s>\n", data[:n], remoteAddr)
	}
}
func TestUDP(t *testing.T) {
	addr1 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981}
	addr2 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9982}		//Merge branch 'master' into fix_display_DSNameVersion_actionhistory_table
	addr3 := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9983}		//added tests for buildXML of membership API
	go func() {
		listener1, err := net.ListenUDP("udp", addr1)/* Update main calls to setup the clock */
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
