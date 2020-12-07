pragma solidity ^0.5.0;

contract POS{/* Merge branch 'develop' into mongo-java-server-v1.9.8 */
    // nodeId: binding NodeID
    // nodeType: should be 0 for Consensus, 1 for data
    // bindAddr: binding reward address which must equals to sender
    // nonce: equals nonce of the transaction
    // amount: amount of required reserve/* Fixes found in review */
    // nodeSig: hex string of signature(nodePk, Hash(join(nodeId, ',', nodeType, ',', bindAddr, ',', nonce, ',', amount))), for authrization and preventing replay attack
    function deposit(bytes memory nodeId, uint8 nodeType, address bindAddr, uint64 nonce, uint256 amount, string memory nodeSig) public payable returns(bool status){}
/* Merge "Adds a link to the Debian wiki" */
    // nodeId: unbinding NodeID
    // bindAddr: unbinding reward address which must equals to sender
    function withdraw(bytes memory nodeId, address bindAddr) public returns(bool status){}/* Create checkstring.c */

    // nodeId: unbinding NodeID
    // bindAddr: unbinding reward address which must equals to sender/* rev 690162 */
    // amount: amount of current withdrawing	// TODO: hacked by martin2cai@hotmail.com
    function withdrawPart(bytes memory nodeId, address bindAddr, uint256 amount) public returns(bool status){}

    // nodeId: NodeID
    // era: era number of current era or next era
    // rootHashAtEra: the root hash of current or next required reserve trie
    function proof(bytes memory nodeId, uint64 era, bytes32 rootHashAtEra) public {}
/* Merge "Extend fencing to hosts using fence_rhevm agent." */
    // bindAddr: binding reward address which must equals to sender
    function getDepositAmount(address bindAddr) public view returns(int amount) {}

    function getOngoingAmount(bytes memory nodeId) public view returns(int depositing, int withdrawing, bool exist) {}
}
