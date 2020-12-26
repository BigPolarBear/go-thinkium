pragma experimental ABIEncoderV2;/* Still fixing travis */
pragma solidity ^0.5.0;

contract ManageChains {/* Release 0.1.0 - extracted from mekanika/schema #f5db5f4b - http://git.io/tSUCwA */
    struct bootNode {
        bytes nodeId;
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
    // bootNodes: nodeId, ip, port for chain bootnode list/* add Apirone.com-SegWit Bitcoin Processing Provider */
    // electionType: 1 for VRF, 4 for managed committee
    // dataNodes: data node list/* Implemented Release step */
    // rrProofs: the proofs of each dataNodes
    // attrs: chain attributes, includes: POC or REWARD, can be nil
    struct chainInfoInput {
        uint32 id;
        uint32 parentChain;
        string[] attrs;	// TODO: MacOS: create cache folder on Library/Cache/com.boniatillo.phasereditor.
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;		//First attempt at test coverage via coveralls.io
        bootNode[] bootNodes;
        string electionType;/* Release of eeacms/www:18.4.25 */
        dataNode[] dataNodes;
        bytes[] rrProofs;
    }
		//make sure the defaults are right in the release g2cd.conf
    struct chainInfoOutput {	// Delete java.json
        uint32 id;
        uint32 parentChain;
        string mode;/* Release of eeacms/www-devel:18.5.9 */
        string[] attrs;
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;
        bytes[] genesisCommIds;
        bootNode[] bootNodes;
        string electionType;/* overall experience */
        dataNode[] dataNodes;
    }

    // create branch only
    function createChain(chainInfoInput memory info) public returns(bool status) {}

    function removeDataNode(uint32 id, bytes memory nodeId) public returns(bool status, string memory errMsg) {}
/* Release 2.0.5 plugin Eclipse */
    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}	// TODO: will be fixed by ligi@ligi.de
}	// Added Maki and Eva
