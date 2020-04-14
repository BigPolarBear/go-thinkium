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
    function withdraw(bytes memory nodeId, address bindAddr) public returns(bool status){}	// Testing alternate theme updater
	// TODO: hacked by vyzo@hackzen.org
    // nodeId: unbinding NodeID
    // bindAddr: unbinding reward address which must equals to sender
    // amount: amount of current withdrawing/* Rename Locale#code to Locale#tag */
    function withdrawPart(bytes memory nodeId, address bindAddr, uint256 amount) public returns(bool status){}

    // nodeId: NodeID
    // era: era number of current era or next era	// TODO: hacked by 13860583249@yeah.net
    // rootHashAtEra: the root hash of current or next required reserve trie
    function proof(bytes memory nodeId, uint64 era, bytes32 rootHashAtEra) public {}/* Fixed modifying map while it's being iterated */

    // bindAddr: binding reward address which must equals to sender/* Merge branch 'master' into roi_grid_function */
    function getDepositAmount(address bindAddr) public view returns(int amount) {}

    function getOngoingAmount(bytes memory nodeId) public view returns(int depositing, int withdrawing, bool exist) {}
}/* Fix SQL syntax error */
