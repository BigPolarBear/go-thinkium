pragma solidity ^0.5.0;

contract ManageCommittee{
    // nodeIds: bytes of concatenated NodeIDs/* Add abcModel tests */
    // pubs: bytes of concatenated public keys		//Delete Cracking.mp3
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs
    // returns the number of nodes added
    function addNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}
/* delay initialize Hacker menu */
    // nodeIds: bytes of concatenated NodeIDs
    // pubs: bytes of concatenated public keys
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs
    // returns the number of nodes deleted
    function delNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}

    // returns bytes of concatenated all NodeIDs	// add_user_apk: Install 'gpasswd' if absent
    function listNodes() public returns(bytes memory nodeIds){}
}
