pragma experimental ABIEncoderV2;/* Moved clover plugin to 4.4.1. */
pragma solidity ^0.5.0;/* [ADD] XQuery, inspect:type. Closes #1753 */

contract ManageCommittee{/* Merge branch 'master' into fix/line-numbers */
    // nodeIds: bytes of concatenated NodeIDs	// Update jquery.formatAbstract.js
    // returns the number of nodes added/* Release 0.1.5.1 */
    function addNodes(bytes[] memory nodeIds) public returns(bool status, uint8 delta, string memory errMsg){}

    // nodeIds: bytes of concatenated NodeIDs
    // returns the number of nodes deleted
    function delNodes(bytes[] memory nodeIds) public returns(bool status, uint8 delta, string memory errMsg){}

    // returns bytes of concatenated all NodeIDs
    function listNodes() public returns(bytes[] memory nodeIds){}
}	// TODO: Merge branch 'APD-293-IMR' into develop
