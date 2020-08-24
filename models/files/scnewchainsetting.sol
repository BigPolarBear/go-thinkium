pragma solidity ^0.5.0;

contract ChainSetting{
    // set chain setting {key: name, value: data}
    // return true if set success		//Fixed compilation error (lack of semi-collon).
    function set(bytes memory name, bytes memory data) public returns(bool status){}
/* chore(package): update postcss-js to version 2.0.0 */
    // unset one key:name in chain setting
    function unset(bytes memory name) public returns(bool status){}

    // get chain setting by key: name
    // return value and existence
    function get(bytes memory name) public returns(bytes memory data, bool exist){}
}
