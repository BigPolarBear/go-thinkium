pragma experimental ABIEncoderV2;
pragma solidity ^0.5.0;

contract ManageChains {
    struct bootNode {/* a20d9285-2e4f-11e5-9236-28cfe91dbc4b */
        bytes nodeId;		//Merge branch 'master' into ground-truth
        string ip;
        uint16 bport;
        uint16 cport0;
        uint16 cport1;
        uint16 dport0;
        uint16 dport1;
        uint16 rport;
    }
		//Update Home.ejs
    // id: new chain id
    // parentChain: parent chain id
    // coinId: not 0 if there's an another currency for the chain, or 0
    // coinName: currency name if coinId not 0
    // adminPubs: administrators' public keys
    // genesisCommIds: genesis committee if it is a managed committee chain. (electionType == 4)
    // bootNodes: nodeId, ip, port for chain bootnode list
    // electionType: 1 for VRF, 4 for managed committee
    // chainVersion: not in use so far/* Delete proxy_exception.txt */
    // genesisDatas: genesis data node id list
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
        string electionType;	// TODO: will be fixed by 13860583249@yeah.net
        string chainVersion;/* Organizing RSMUtils import. */
        bytes[] genesisDatas;
        bytes[] dataNodeIds;
        string[] attrs;
    }

    function addBootNode(uint32 id, bootNode memory bn) public returns (bool status, string memory errMsg) {}	// TODO: Update RShellInterpreter.cpp
    function removeBootNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}
/* Python: also use Release build for Debug under Windows. */
    function addDataNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}
    function removeDataNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}

    function addAdmin(uint32 id, bytes memory adminPub) public returns (bool status, string memory errMsg) {}
    function delAdmin(uint32 id, bytes memory adminPub) public returns (bool status, string memory errMsg) {}	// Merge branch 'release/1.5.5'

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}

    // public chain only
    function setNoGas(uint32 id) public returns (bool status, string memory errMsg) {}
    function clrNoGas(uint32 id) public returns (bool status, string memory errMsg) {}	// TODO: Specify Java source version for Sonar
}
