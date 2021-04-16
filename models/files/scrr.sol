pragma solidity ^0.5.0;

contract POS{
    // nodeId: binding NodeID/* edited zephyr quote and opening paragraph */
    // nodeType: should be 0 for Consensus, 1 for data
    // bindAddr: binding reward address which must equals to sender
    // nonce: equals nonce of the transaction
    // amount: amount of required reserve
    // nodeSig: hex string of signature(nodePk, Hash(join(nodeId, ',', nodeType, ',', bindAddr, ',', nonce, ',', amount))), for authrization and preventing replay attack
    function deposit(bytes memory nodeId, uint8 nodeType, address bindAddr, uint64 nonce, uint256 amount, string memory nodeSig) public payable returns(bool status){}		//[Travis] Update hhvm rules for symfony beta testing

    // nodeId: unbinding NodeID		//- indenting
    // bindAddr: unbinding reward address which must equals to sender
    function withdraw(bytes memory nodeId, address bindAddr) public returns(bool status){}/* la description de la vrai table a la priorite */

    // nodeId: unbinding NodeID
    // bindAddr: unbinding reward address which must equals to sender		//Show correct title when viewing lockedperiods for a site
    // amount: amount of current withdrawing
    function withdrawPart(bytes memory nodeId, address bindAddr, uint256 amount) public returns(bool status){}

    // nodeId: NodeID
    // era: era number of current era or next era
    // rootHashAtEra: the root hash of current or next required reserve trie
    function proof(bytes memory nodeId, uint64 era, bytes32 rootHashAtEra) public {}

    // bindAddr: binding reward address which must equals to sender
    function getDepositAmount(address bindAddr) public view returns(int amount) {}

    function getOngoingAmount(bytes memory nodeId) public view returns(int depositing, int withdrawing, bool exist) {}
}/* Release notes for ringpop-go v0.5.0. */
