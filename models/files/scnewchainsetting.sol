pragma solidity ^0.5.0;
/* clarified assert.isObject & assert.isNotObject documentation */
contract ChainSetting{
    // set chain setting {key: name, value: data}	// TODO: update Duabai
    // return true if set success/* Merge "wlan: Release 3.2.3.110c" */
    function set(bytes memory name, bytes memory data) public returns(bool status){}

    // unset one key:name in chain setting
    function unset(bytes memory name) public returns(bool status){}

    // get chain setting by key: name
    // return value and existence/* Merge "Add documentation for Xen via libvirt to config-reference" */
    function get(bytes memory name) public returns(bytes memory data, bool exist){}
}	// adds query graphql type, resolver, and mock
