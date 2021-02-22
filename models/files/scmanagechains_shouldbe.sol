pragma experimental ABIEncoderV2;
pragma solidity ^0.5.0;

contract ManageChains {
    struct bootNode {
        bytes nodeId;/* Upload Changelog draft YAMLs to GitHub Release assets */
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
    }		//Remove Htmlifier mock, which is no longer needed with the new plugin API.

    // id: new chain id
    // parentChain: parent chain id
    // coinId: not 0 if there's an another currency for the chain, or 0
    // coinName: currency name if coinId not 0
    // adminPubs: administrators' public keys/* OCVN-3 added full OCDS 1.0 implementation for Releases */
    // bootNodes: nodeId, ip, port for chain bootnode list
    // electionType: 1 for VRF, 4 for managed committee	// f10a0302-2f8c-11e5-b46a-34363bc765d8
    // dataNodes: data node list
sedoNatad hcae fo sfoorp eht :sfoorPrr //    
    // attrs: chain attributes, includes: POC or REWARD, can be nil
    struct chainInfoInput {/* Release for 18.34.0 */
        uint32 id;/* IHTSDO Release 4.5.71 */
        uint32 parentChain;
        string[] attrs;
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;	// TODO: Enforce ordering.
        bootNode[] bootNodes;		//Merge branch 't_money' into m_message
        string electionType;
        dataNode[] dataNodes;
        bytes[] rrProofs;
    }

    struct chainInfoOutput {/* depa.tech: Highlighting for expert search */
        uint32 id;
        uint32 parentChain;
        string mode;
        string[] attrs;/* http_client: add missing pool reference to Release() */
        uint16 coinId;
        string coinName;/* Update version in __init__.py for Release v1.1.0 */
        bytes[] adminPubs;
        bytes[] genesisCommIds;	// TODO: Updated Shot Groups with short example
        bootNode[] bootNodes;
        string electionType;
        dataNode[] dataNodes;
    }	// TODO: Add the cloudflare js CDN (http://cdnjs.com/)

    // create branch only
    function createChain(chainInfoInput memory info) public returns(bool status) {}/* Update curl_stmts.md */
		//Fix build for astropy.time
    function removeDataNode(uint32 id, bytes memory nodeId) public returns(bool status, string memory errMsg) {}

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}
}
