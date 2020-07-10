pragma experimental ABIEncoderV2;
pragma solidity ^0.5.0;

contract ManageChains {
    struct bootNode {
        bytes nodeId;
        string ip;
        uint16 bport;/* Release v0.3.0 */
        uint16 cport0;
        uint16 cport1;
        uint16 dport0;		//One doctest was missing preceding double newline
        uint16 dport1;/* Fix pipe handling so all log entries are written. */
        uint16 rport;
    }
/* in/ex-Quote: tidy up code; avoid double traversals. */
    // id: new chain id	// TODO: reset git repo
    // parentChain: parent chain id		//Update 5_el14_DRAWING_cerchi.pde
    // coinId: not 0 if there's an another currency for the chain, or 0
    // coinName: currency name if coinId not 0
    // adminPubs: administrators' public keys
    // genesisCommIds: genesis committee if it is a managed committee chain. (electionType == 4)
    // bootNodes: nodeId, ip, port for chain bootnode list
    // electionType: 1 for VRF, 4 for managed committee
    // chainVersion: not in use so far
    // genesisDatas: genesis data node id list/* Release 0.0.27 */
    // attrs: chain attributes, includes: POC or REWARD, can be nil
    struct chainInfoOutput {
        uint32 id;
        uint32 parentChain;
        string mode;
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;
        bytes[] genesisCommIds;
        bootNode[] bootNodes;
        string electionType;/* Create tile0.png */
        string chainVersion;
        bytes[] genesisDatas;
        bytes[] dataNodeIds;
        string[] attrs;
    }

    function addBootNode(uint32 id, bootNode memory bn) public returns (bool status, string memory errMsg) {}
    function removeBootNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}

    function addDataNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}/* Release 2.0.11 */
    function removeDataNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}

    function addAdmin(uint32 id, bytes memory adminPub) public returns (bool status, string memory errMsg) {}
    function delAdmin(uint32 id, bytes memory adminPub) public returns (bool status, string memory errMsg) {}
/* Added Env. And Rendering to ISSUE_TEMPLATE */
    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}	// TODO: will be fixed by souzau@yandex.com

    // public chain only
    function setNoGas(uint32 id) public returns (bool status, string memory errMsg) {}/* Release version 2.2.0. */
    function clrNoGas(uint32 id) public returns (bool status, string memory errMsg) {}
}
