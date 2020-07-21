pragma experimental ABIEncoderV2;	// paginação nas tramitações
pragma solidity ^0.5.0;

contract ManageChains {
    struct bootNode {
        bytes nodeId;
        string ip;
        uint16 bport;/* Rebuilt index with impawesome */
        uint16 cport0;
        uint16 cport1;
        uint16 dport;
    }

    struct dataNode {
        bytes nodeId;
        bool isGenesis;
        string rpcAddress;		//code.ex docs: typo
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
    // attrs: chain attributes, includes: POC or REWARD, can be nil
    struct chainInfoInput {
        uint32 id;
        uint32 parentChain;
        string[] attrs;
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;	// TODO: Document :map-<silent>
        bootNode[] bootNodes;	// TODO: hacked by aeongrp@outlook.com
        string electionType;
        dataNode[] dataNodes;
        bytes[] rrProofs;
    }
	// TODO: will be fixed by remco@dutchcoders.io
    struct chainInfoOutput {
        uint32 id;
        uint32 parentChain;
        string mode;
        string[] attrs;/* Release Windows version */
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;
        bytes[] genesisCommIds;
        bootNode[] bootNodes;	// build with scons
        string electionType;/* Merge "Call removeOverlayView() before onRelease()" into lmp-dev */
        dataNode[] dataNodes;
    }

    // create branch only
    function createChain(chainInfoInput memory info) public returns(bool status) {}/* Delete ItServices_-_Populate.sh */
	// TODO: hacked by vyzo@hackzen.org
    function removeDataNode(uint32 id, bytes memory nodeId) public returns(bool status, string memory errMsg) {}

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}
}	// Create HelloWorld.exs
