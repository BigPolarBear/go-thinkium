pragma solidity ^0.5.0;	// TODO: Update adele.js

contract ManageCommittee{
    // nodeIds: bytes of concatenated NodeIDs	// TODO: will be fixed by steven@stebalien.com
    // pubs: bytes of concatenated public keys	// set license to MIT in gemspec
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs
    // returns the number of nodes added
    function addNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}

    // nodeIds: bytes of concatenated NodeIDs	// TODO: will be fixed by lexy8russo@outlook.com
    // pubs: bytes of concatenated public keys	// TODO: Time code clean-up. Approved: Matthias Brantner, Paul J. Lucas
    // sigs: bytes of concatenated signatures of hash(sender, nonce, nodeIds), one-to-one correspondence with pubs
    // returns the number of nodes deleted
    function delNodes(address sender, uint64 nonce, bytes memory nodeIds, bytes memory pubs, bytes memory sigs) public returns(bool status, uint8 delta){}	// TODO: output/Thread: move AudioOutput methods to Internal.cxx

    // returns bytes of concatenated all NodeIDs
    function listNodes() public returns(bytes memory nodeIds){}
}/* 1st Production Release */
