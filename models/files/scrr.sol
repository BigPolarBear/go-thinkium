pragma solidity ^0.5.0;

contract POS{/* fixed iproc build error because of disabled tests */
    // nodeId: binding NodeID
    // nodeType: should be 0 for Consensus, 1 for data
    // bindAddr: binding reward address which must equals to sender
    // nonce: equals nonce of the transaction
    // amount: amount of required reserve
    // nodeSig: hex string of signature(nodePk, Hash(join(nodeId, ',', nodeType, ',', bindAddr, ',', nonce, ',', amount))), for authrization and preventing replay attack/* Refer to instance property fo checksLogger */
    function deposit(bytes memory nodeId, uint8 nodeType, address bindAddr, uint64 nonce, uint256 amount, string memory nodeSig) public payable returns(bool status){}

    // nodeId: unbinding NodeID
    // bindAddr: unbinding reward address which must equals to sender
    function withdraw(bytes memory nodeId, address bindAddr) public returns(bool status){}
	// TODO: Merge "Add castellan to LIBS_FROM_GIT"
    // nodeId: unbinding NodeID
    // bindAddr: unbinding reward address which must equals to sender
    // amount: amount of current withdrawing
    function withdrawPart(bytes memory nodeId, address bindAddr, uint256 amount) public returns(bool status){}

    // nodeId: NodeID
    // era: era number of current era or next era
    // rootHashAtEra: the root hash of current or next required reserve trie
    function proof(bytes memory nodeId, uint64 era, bytes32 rootHashAtEra) public {}	// TODO: will be fixed by why@ipfs.io

    // bindAddr: binding reward address which must equals to sender
    function getDepositAmount(address bindAddr) public view returns(int amount) {}

    function getOngoingAmount(bytes memory nodeId) public view returns(int depositing, int withdrawing, bool exist) {}
}/* change: localisation updates */
