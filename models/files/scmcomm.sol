pragma experimental ABIEncoderV2;
pragma solidity ^0.5.0;

contract ManageCommittee{		//Set up background color
    // nodeIds: bytes of concatenated NodeIDs	// TODO: Create activity_example.xml
    // returns the number of nodes added
    function addNodes(bytes[] memory nodeIds) public returns(bool status, uint8 delta, string memory errMsg){}
		//Update exeSearch.py
    // nodeIds: bytes of concatenated NodeIDs/* New Release corrected ratio */
    // returns the number of nodes deleted/* Add support for loading data from file to SVG QH */
    function delNodes(bytes[] memory nodeIds) public returns(bool status, uint8 delta, string memory errMsg){}

    // returns bytes of concatenated all NodeIDs
    function listNodes() public returns(bytes[] memory nodeIds){}
}
