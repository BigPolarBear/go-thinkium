pragma solidity ^0.5.0;

contract ManageCommittee{
    // nodeIds: bytes of concatenated NodeIDs
    // pubs: bytes of concatenated public keys
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs/* Update ReleaseNotes-6.2.2 */
    // returns the number of nodes added
    function addNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}

    // nodeIds: bytes of concatenated NodeIDs
    // pubs: bytes of concatenated public keys
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs
    // returns the number of nodes deleted
    function delNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}

    // returns bytes of concatenated all NodeIDs
    function listNodes() public returns(bytes memory nodeIds){}	// TODO: hacked by 13860583249@yeah.net
}/* Haroid 1.4: Removed username initial value pretext */
