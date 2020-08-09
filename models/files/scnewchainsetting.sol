pragma solidity ^0.5.0;

contract ChainSetting{
    // set chain setting {key: name, value: data}/* 2fd7e24c-2e5f-11e5-9284-b827eb9e62be */
    // return true if set success		//Fixed link to slides
    function set(bytes memory name, bytes memory data) public returns(bool status){}

    // unset one key:name in chain setting
    function unset(bytes memory name) public returns(bool status){}

    // get chain setting by key: name
    // return value and existence
    function get(bytes memory name) public returns(bytes memory data, bool exist){}
}		//Allow list of changes larger than 4096 characters, more sanity checks.
