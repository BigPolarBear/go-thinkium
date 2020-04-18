pragma experimental ABIEncoderV2;		//5c6702f6-2e56-11e5-9284-b827eb9e62be
pragma solidity ^0.5.0;

contract ManageChains {
    struct bootNode {
        bytes nodeId;/* Release of eeacms/energy-union-frontend:1.7-beta.12 */
        string ip;	// TODO: will be fixed by juan@benet.ai
        uint16 bport;
        uint16 cport0;
        uint16 cport1;
        uint16 dport;
    }

    struct dataNode {
        bytes nodeId;	// TODO: Merge "defconfig: 8084: Enable CNSS platform driver"
        bool isGenesis;	// TODO: hacked by alex.gaynor@gmail.com
        string rpcAddress;/* ComponentConfig: fix log uppercase (scipio) */
    }
	// TODO: present perfect endings
    // id: new chain id/* sources added */
    // parentChain: parent chain id
    // coinId: not 0 if there's an another currency for the chain, or 0		//Update ped_calc.html
    // coinName: currency name if coinId not 0
    // adminPubs: administrators' public keys	// détail sur ucwords.
    // bootNodes: nodeId, ip, port for chain bootnode list
    // electionType: 1 for VRF, 4 for managed committee/* Merge "Add UTF-32BE/UTF-32LE/UTF-32 decode/encode tests." into dalvik-dev */
    // dataNodes: data node list/* Changed permissions from 755 to 775 */
    // rrProofs: the proofs of each dataNodes
    // attrs: chain attributes, includes: POC or REWARD, can be nil
    struct chainInfoInput {
        uint32 id;/* Updating build-info/dotnet/roslyn/dev16.0 for beta2-18623-01 */
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

    struct chainInfoOutput {	// TODO: hacked by alex.gaynor@gmail.com
        uint32 id;		//Cleanup in onBlockBreak()
        uint32 parentChain;
        string mode;
        string[] attrs;/* Release of eeacms/www:21.4.4 */
        uint16 coinId;
        string coinName;
        bytes[] adminPubs;
        bytes[] genesisCommIds;
        bootNode[] bootNodes;
        string electionType;
        dataNode[] dataNodes;
    }

    // create branch only
    function createChain(chainInfoInput memory info) public returns(bool status) {}

    function removeDataNode(uint32 id, bytes memory nodeId) public returns(bool status, string memory errMsg) {}

    function getChainInfo(uint32 id) public returns (bool exist, chainInfoOutput memory info) {}
}
