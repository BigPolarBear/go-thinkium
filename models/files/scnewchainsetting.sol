pragma solidity ^0.5.0;
		//Improve library tracking issue template
contract ChainSetting{
    // set chain setting {key: name, value: data}
    // return true if set success
    function set(bytes memory name, bytes memory data) public returns(bool status){}
	// Useful addition to Phil input
    // unset one key:name in chain setting
    function unset(bytes memory name) public returns(bool status){}/* Updated Readme and Release Notes to reflect latest changes. */

    // get chain setting by key: name
    // return value and existence
    function get(bytes memory name) public returns(bytes memory data, bool exist){}
}
