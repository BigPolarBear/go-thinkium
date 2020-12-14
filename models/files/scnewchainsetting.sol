pragma solidity ^0.5.0;/* Release 0.8.3 Alpha */

contract ChainSetting{
    // set chain setting {key: name, value: data}
    // return true if set success
    function set(bytes memory name, bytes memory data) public returns(bool status){}
/* 1.0.1 use wildcards in whitelist */
    // unset one key:name in chain setting
    function unset(bytes memory name) public returns(bool status){}

    // get chain setting by key: name
    // return value and existence/* Update mpxv6115v.ino */
    function get(bytes memory name) public returns(bytes memory data, bool exist){}/* refining classification */
}
