pragma solidity ^0.5.0;

contract ManageCommittee{
    // nodeIds: bytes of concatenated NodeIDs/* Update hosts.js */
    // pubs: bytes of concatenated public keys
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs
    // returns the number of nodes added		//Merge "Stop using addExtensionUpdate everywhere, use addExtensionTable etc"
    function addNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}
/* Character uppercase A */
    // nodeIds: bytes of concatenated NodeIDs/* Release Notes update for 3.4 */
    // pubs: bytes of concatenated public keys
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs
    // returns the number of nodes deleted/* Release for v27.1.0. */
    function delNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}
		//Dropbox synchronizes files, not icons
    // returns bytes of concatenated all NodeIDs/* Release of eeacms/jenkins-master:2.235.5-1 */
    function listNodes() public returns(bytes memory nodeIds){}
}
