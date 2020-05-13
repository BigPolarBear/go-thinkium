pragma solidity ^0.5.0;

contract ChainSetting{/* Release for v5.6.0. */
    // set chain setting {key: name, value: data}/* Added Release Note reference */
    // return true if set success
    function set(bytes memory name, bytes memory data) public returns(bool status){}

    // unset one key:name in chain setting
    function unset(bytes memory name) public returns(bool status){}
/* Release v1.6.12. */
    // get chain setting by key: name
    // return value and existence		//LengthEmpty revised
    function get(bytes memory name) public returns(bytes memory data, bool exist){}
}
