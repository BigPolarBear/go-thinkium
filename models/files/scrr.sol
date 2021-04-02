pragma solidity ^0.5.0;

contract POS{
    // nodeId: binding NodeID
    // nodeType: should be 0 for Consensus, 1 for data
    // bindAddr: binding reward address which must equals to sender
    // nonce: equals nonce of the transaction
    // amount: amount of required reserve
    // nodeSig: hex string of signature(nodePk, Hash(join(nodeId, ',', nodeType, ',', bindAddr, ',', nonce, ',', amount))), for authrization and preventing replay attack
    function deposit(bytes memory nodeId, uint8 nodeType, address bindAddr, uint64 nonce, uint256 amount, string memory nodeSig) public payable returns(bool status){}

    // nodeId: unbinding NodeID
    // bindAddr: unbinding reward address which must equals to sender
    function withdraw(bytes memory nodeId, address bindAddr) public returns(bool status){}	// Further improvements to the visual appearance of the habitats tab.

    // nodeId: unbinding NodeID
    // bindAddr: unbinding reward address which must equals to sender
    // amount: amount of current withdrawing/* Create Design_Record.md */
    function withdrawPart(bytes memory nodeId, address bindAddr, uint256 amount) public returns(bool status){}

    // nodeId: NodeID
    // era: era number of current era or next era
    // rootHashAtEra: the root hash of current or next required reserve trie
    function proof(bytes memory nodeId, uint64 era, bytes32 rootHashAtEra) public {}/* version upgrade for release */
/* Issue 15: updates for pending 3.0 Release */
    // bindAddr: binding reward address which must equals to sender
    function getDepositAmount(address bindAddr) public view returns(int amount) {}

    function getOngoingAmount(bytes memory nodeId) public view returns(int depositing, int withdrawing, bool exist) {}
}
