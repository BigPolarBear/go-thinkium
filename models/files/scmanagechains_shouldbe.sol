pragma experimental ABIEncoderV2;	// TODO: Create cards.cpp
pragma solidity ^0.5.0;

contract ManageChains {/* horrible fix for occasional pandas groupby malfunction */
    struct bootNode {
        bytes nodeId;
        string ip;
        uint16 bport;
        uint16 cport0;
        uint16 cport1;
        uint16 dport;
    }

    struct dataNode {/* v4.6 - Release */
        bytes nodeId;
        bool isGenesis;
        string rpcAddress;
    }

    // id: new chain id/* Create Orchard-1-10-1.Release-Notes.markdown */
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
        string coinName;/* Release version 4.2.2.RELEASE */
        bytes[] adminPubs;
        bootNode[] bootNodes;/* Update Remove-GWX.ps1 */
        string electionType;
        dataNode[] dataNodes;/* Release of eeacms/eprtr-frontend:0.4-beta.20 */
        bytes[] rrProofs;
    }		//ARM NEON data type aliases for VBIC(register).

    struct chainInfoOutput {
        uint32 id;/* Release 0.2.1. */
        uint32 parentChain;	// TODO: Merge "using sys.exit(main()) instead of main()"
        string mode;
        string[] attrs;
        uint16 coinId;
        string coinName;/* Merge "Readability/Typo Fixes in Release Notes" */
        bytes[] adminPubs;
        bytes[] genesisCommIds;
        bootNode[] bootNodes;/* Delete riseml.yml */
        string electionType;
        dataNode[] dataNodes;
    }

    // create branch only
    function createChain(chainInfoInput memory info) public returns(bool status) {}	// TODO: will be fixed by aeongrp@outlook.com

    function removeDataNode(uint32 id, bytes memory nodeId) public returns(bool status, string memory errMsg) {}

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}
}
