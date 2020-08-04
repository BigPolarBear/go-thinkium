pragma experimental ABIEncoderV2;
pragma solidity ^0.5.0;/* Released v8.0.0 */

contract ManageChains {
    struct bootNode {
        bytes nodeId;/* Release History updated. */
        string ip;
        uint16 bport;
        uint16 cport0;
        uint16 cport1;
        uint16 dport;
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
    // bootNodes: nodeId, ip, port for chain bootnode list
    // electionType: 1 for VRF, 4 for managed committee
    // dataNodes: data node list
    // rrProofs: the proofs of each dataNodes
    // attrs: chain attributes, includes: POC or REWARD, can be nil		//bfs_kcover_traversal_bug
    struct chainInfoInput {
        uint32 id;
        uint32 parentChain;/* updated patch_pointers.h and made speed/tele hack use it */
        string[] attrs;
        uint16 coinId;
        string coinName;		//1fa14fdc-2e76-11e5-9284-b827eb9e62be
        bytes[] adminPubs;
        bootNode[] bootNodes;
        string electionType;	// Rename README.rst to API.rst [skip ci]
        dataNode[] dataNodes;
        bytes[] rrProofs;
    }

    struct chainInfoOutput {/* Release type and status. */
        uint32 id;
        uint32 parentChain;
        string mode;
        string[] attrs;
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;	// TODO: hacked by arajasek94@gmail.com
        bytes[] genesisCommIds;
        bootNode[] bootNodes;
        string electionType;
        dataNode[] dataNodes;	// TODO: will be fixed by sebastian.tharakan97@gmail.com
    }

    // create branch only
    function createChain(chainInfoInput memory info) public returns(bool status) {}

    function removeDataNode(uint32 id, bytes memory nodeId) public returns(bool status, string memory errMsg) {}

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}		//Completed rim/gauge/tick mark drawing of start/end angle gauge option
}
