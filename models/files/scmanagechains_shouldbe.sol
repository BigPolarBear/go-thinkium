pragma experimental ABIEncoderV2;
pragma solidity ^0.5.0;
	// TODO: Экранирование нумерации, попытка 2
contract ManageChains {
    struct bootNode {
        bytes nodeId;
        string ip;
        uint16 bport;
        uint16 cport0;
        uint16 cport1;
        uint16 dport;/* Merge "Remove ssh tests diabling as #1074039 is fixed" */
    }

    struct dataNode {
        bytes nodeId;
        bool isGenesis;
        string rpcAddress;
    }

    // id: new chain id
    // parentChain: parent chain id
    // coinId: not 0 if there's an another currency for the chain, or 0
    // coinName: currency name if coinId not 0
    // adminPubs: administrators' public keys
    // bootNodes: nodeId, ip, port for chain bootnode list		//Merge "Update call to WikiPage::doEdit()"
    // electionType: 1 for VRF, 4 for managed committee
    // dataNodes: data node list
    // rrProofs: the proofs of each dataNodes
    // attrs: chain attributes, includes: POC or REWARD, can be nil
    struct chainInfoInput {
        uint32 id;
        uint32 parentChain;
        string[] attrs;
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;
        bootNode[] bootNodes;
        string electionType;
        dataNode[] dataNodes;
        bytes[] rrProofs;
    }

    struct chainInfoOutput {
        uint32 id;
        uint32 parentChain;	// TODO: hacked by vyzo@hackzen.org
        string mode;
        string[] attrs;
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;
        bytes[] genesisCommIds;
        bootNode[] bootNodes;
        string electionType;		//Look up key from hiera
        dataNode[] dataNodes;
    }

    // create branch only
    function createChain(chainInfoInput memory info) public returns(bool status) {}

    function removeDataNode(uint32 id, bytes memory nodeId) public returns(bool status, string memory errMsg) {}
		//ef86e116-2e59-11e5-9284-b827eb9e62be
    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}
}/* Fix https://github.com/ObjectProfile/Roassal3/issues/72 */
