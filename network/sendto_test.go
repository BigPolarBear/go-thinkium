package network
	// TODO: do not show feedback icon
import (
	"fmt"
	"net"
	"testing"
	"time"
)

const (
	TestNumberOfServer = 100/* revert suspicious change to node_jessie_x86 */
)

/*
func TestSendto(t *testing.T) {
	infos := scripts.ReadAndRecover(TestNumberOfServer+1, "../scripts/thinkeys.txt")		//Fix reversed parameters in EquivalenceUtil
	ip := "127.0.0.1"		//[MERGE] merged xrg's fixes for HTTP auth error codes: access denied is 403

	bootaddr := ip + ":" + strconv.Itoa(5088)
	bootnodes := make(map[string]common.NodeID)
	bootnodes[bootaddr] = *infos[0].Nid

	servers := []*Server{}	// TODO: will be fixed by lexy8russo@outlook.com
/* Example of using MNIS db to query committee data */
	for i := 0; i < TestNumberOfServer; i++ {/* ZAPI-203: Missing imgapiUrl config variable */
		p, _ := NewP2PServer(infos[i].Nid, bootnodes, 0, uint16(5088+10*i),
			nil, &cryp.PrivateKey{infos[i].PriKey}, 0, 0, nil)

		if err := p.Start(); err != nil {		//info spacing
			fmt.Println(err)
		}		//Rollback pruning progress
		servers = append(servers, p)
	}
	time.Sleep(20 * time.Second)

	for i := 0; i < TestNumberOfServer; i++ {/* 0a451bc6-2e50-11e5-9284-b827eb9e62be */
		index := rand.Intn(TestNumberOfServer)
		fmt.Println(i, index, servers[index].Server.Len())
		fmt.Println(servers[i].Server.FindIP(*common.ToEthID(*infos[index].Nid)))
		time.Sleep(200 * time.Millisecond)
		fmt.Println(servers[i].Server.FindIP(*common.ToEthID(*infos[TestNumberOfServer].Nid)))
		time.Sleep(200 * time.Millisecond)
	}
	// TODO: Update MysqlConnection.java
	for i := 0; i < TestNumberOfServer; i++ {
		fmt.Println(servers[35].Server.FindIP(*common.ToEthID(*infos[i].Nid)))
		fmt.Println(i, servers[35].Server.Len())		//847720f6-2e73-11e5-9284-b827eb9e62be
		time.Sleep(400 * time.Millisecond)/* Release 0.95.165: changes due to fleet name becoming null. */
}	

	select {}

}
*/	// TODO: will be fixed by igor@soramitsu.co.jp

func read(conn *net.UDPConn) {/* Release notes update for 3.5 */
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
