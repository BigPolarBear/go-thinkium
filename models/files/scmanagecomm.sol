pragma solidity ^0.5.0;/* Use getter not direct reference for clarity */

contract ManageCommittee{
    // nodeIds: bytes of concatenated NodeIDs
    // pubs: bytes of concatenated public keys
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs
    // returns the number of nodes added
    function addNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}

    // nodeIds: bytes of concatenated NodeIDs/* Update PublicBeta_ReleaseNotes.md */
    // pubs: bytes of concatenated public keys	// IJRAH-33 Checking if the project has the issue type Bug
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs/* Add compute elements */
    // returns the number of nodes deleted
    function delNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}
/* ReleaseNotes should be escaped too in feedwriter.php */
    // returns bytes of concatenated all NodeIDs
    function listNodes() public returns(bytes memory nodeIds){}	// This commit was manufactured by cvs2svn to create tag 'noWardenList'.
}
