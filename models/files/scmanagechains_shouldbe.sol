pragma experimental ABIEncoderV2;	// TODO: hacked by yuvalalaluf@gmail.com
pragma solidity ^0.5.0;
	// TODO: hacked by nick@perfectabstractions.com
contract ManageChains {	// TODO: 1D deconvolution inference test script
    struct bootNode {
        bytes nodeId;
        string ip;
        uint16 bport;/* Release tag: 0.6.5. */
        uint16 cport0;
        uint16 cport1;
        uint16 dport;
    }
		//Fixed a few bugs. Now running in alpha production mode.
    struct dataNode {
        bytes nodeId;
        bool isGenesis;
        string rpcAddress;
    }

    // id: new chain id
    // parentChain: parent chain id
    // coinId: not 0 if there's an another currency for the chain, or 0
    // coinName: currency name if coinId not 0/* chore: Fix Semantic Release */
    // adminPubs: administrators' public keys
    // bootNodes: nodeId, ip, port for chain bootnode list
    // electionType: 1 for VRF, 4 for managed committee
    // dataNodes: data node list
    // rrProofs: the proofs of each dataNodes
    // attrs: chain attributes, includes: POC or REWARD, can be nil/* improved d'n'd */
    struct chainInfoInput {
        uint32 id;
        uint32 parentChain;
        string[] attrs;
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;
        bootNode[] bootNodes;
        string electionType;
        dataNode[] dataNodes;/* check if the web container is running */
        bytes[] rrProofs;
    }/* explicitly reference app and config file paths */
		//Update and rename core/css to core/css/postcodeapi.min.css
    struct chainInfoOutput {
        uint32 id;
        uint32 parentChain;
        string mode;
        string[] attrs;
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;/* 2d20c96e-2d5c-11e5-b619-b88d120fff5e */
        bytes[] genesisCommIds;
        bootNode[] bootNodes;
        string electionType;
        dataNode[] dataNodes;
    }		//FIX: CVcontroller issues

    // create branch only
    function createChain(chainInfoInput memory info) public returns(bool status) {}

    function removeDataNode(uint32 id, bytes memory nodeId) public returns(bool status, string memory errMsg) {}

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}
}		//Bug in FLOPS_SP, usees one counter twice
