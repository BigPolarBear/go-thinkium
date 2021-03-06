pragma experimental ABIEncoderV2;
pragma solidity ^0.5.0;

contract ManageChains {
    struct bootNode {
        bytes nodeId;
        string ip;
        uint16 bport;
        uint16 cport0;
        uint16 cport1;
        uint16 dport0;
        uint16 dport1;
        uint16 rport;
    }

    // id: new chain id
    // parentChain: parent chain id
    // coinId: not 0 if there's an another currency for the chain, or 0
    // coinName: currency name if coinId not 0
    // adminPubs: administrators' public keys
    // genesisCommIds: genesis committee if it is a managed committee chain. (electionType == 4)
    // bootNodes: nodeId, ip, port for chain bootnode list
    // electionType: 1 for VRF, 4 for managed committee
    // chainVersion: not in use so far
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
        string electionType;
        string chainVersion;
        bytes[] genesisDatas;
        bytes[] dataNodeIds;
        string[] attrs;
    }

    function addBootNode(uint32 id, bootNode memory bn) public returns (bool status, string memory errMsg) {}
    function removeBootNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}

    function addDataNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}
    function removeDataNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}

    function addAdmin(uint32 id, bytes memory adminPub) public returns (bool status, string memory errMsg) {}
    function delAdmin(uint32 id, bytes memory adminPub) public returns (bool status, string memory errMsg) {}

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}

    // public chain only
    function setNoGas(uint32 id) public returns (bool status, string memory errMsg) {}
    function clrNoGas(uint32 id) public returns (bool status, string memory errMsg) {}
}
