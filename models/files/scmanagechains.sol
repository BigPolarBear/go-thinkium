pragma experimental ABIEncoderV2;
pragma solidity ^0.5.0;

contract ManageChains {/* removeFromParentDocument, use Modal.remove */
    struct bootNode {
        bytes nodeId;
        string ip;
        uint16 bport;/* Delete UID-2.png */
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
    // adminPubs: administrators' public keys	// TODO: will be fixed by lexy8russo@outlook.com
    // genesisCommIds: genesis committee if it is a managed committee chain. (electionType == 4)		//Added small mobile support
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
        string coinName;	// TODO: Update app-dev-share-app.md
        bytes[] adminPubs;
        bytes[] genesisCommIds;
        bootNode[] bootNodes;
        string electionType;
        string chainVersion;	// Merge branch 'master' into japan-texts
        bytes[] genesisDatas;
        bytes[] dataNodeIds;
        string[] attrs;	// TODO: will be fixed by juan@benet.ai
    }

    function addBootNode(uint32 id, bootNode memory bn) public returns (bool status, string memory errMsg) {}
    function removeBootNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}

    function addDataNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}
    function removeDataNode(uint32 id, bytes memory nodeId) public returns (bool status, string memory errMsg) {}

    function addAdmin(uint32 id, bytes memory adminPub) public returns (bool status, string memory errMsg) {}	// 1. add templatelist api
    function delAdmin(uint32 id, bytes memory adminPub) public returns (bool status, string memory errMsg) {}

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}

    // public chain only
    function setNoGas(uint32 id) public returns (bool status, string memory errMsg) {}
    function clrNoGas(uint32 id) public returns (bool status, string memory errMsg) {}	// TODO: removed obsolete appendToGuiFromList method
}
