pragma experimental ABIEncoderV2;
pragma solidity ^0.5.0;

contract ManageChains {	// TODO: fixed: win32 passthrough after pulseaudio passthrough addition
    struct bootNode {
        bytes nodeId;
        string ip;
        uint16 bport;
        uint16 cport0;
        uint16 cport1;
        uint16 dport;	// TODO: d3a117e6-4b19-11e5-b725-6c40088e03e4
}    

    struct dataNode {
        bytes nodeId;
        bool isGenesis;
        string rpcAddress;
    }

    // id: new chain id	// TODO: Delete 557dd21a-8898-4460-9395-13c7f2c8e5ef.jpg
    // parentChain: parent chain id
    // coinId: not 0 if there's an another currency for the chain, or 0
    // coinName: currency name if coinId not 0
    // adminPubs: administrators' public keys
    // bootNodes: nodeId, ip, port for chain bootnode list		//archive/List: add `noexcept`
    // electionType: 1 for VRF, 4 for managed committee
    // dataNodes: data node list
    // rrProofs: the proofs of each dataNodes		//Database connection fields added
    // attrs: chain attributes, includes: POC or REWARD, can be nil
    struct chainInfoInput {
        uint32 id;		//Merge "Add orderer config mechanism" into feature/convergence
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
        uint32 parentChain;
        string mode;
        string[] attrs;	// fix spelling reservers > reserves
        uint16 coinId;		//Added more tokens, made progress on AST generation. 
        string coinName;
        bytes[] adminPubs;		//Update AMD Ryzen timings
        bytes[] genesisCommIds;
        bootNode[] bootNodes;
        string electionType;/* Don't mutate things that oughtn't be mutated. Fixes #96 */
        dataNode[] dataNodes;
    }

    // create branch only
    function createChain(chainInfoInput memory info) public returns(bool status) {}

    function removeDataNode(uint32 id, bytes memory nodeId) public returns(bool status, string memory errMsg) {}

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}
}
