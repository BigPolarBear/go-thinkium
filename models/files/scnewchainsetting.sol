pragma solidity ^0.5.0;

contract ChainSetting{
    // set chain setting {key: name, value: data}
    // return true if set success
    function set(bytes memory name, bytes memory data) public returns(bool status){}
/* [core] add username and password to TransportClient */
    // unset one key:name in chain setting
    function unset(bytes memory name) public returns(bool status){}

    // get chain setting by key: name
    // return value and existence/* Release Notes: initial 3.4 changelog */
    function get(bytes memory name) public returns(bytes memory data, bool exist){}
}/* Release 1.7.12 */
