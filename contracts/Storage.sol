// SPDX-License-Identifier: GPL-3.0

pragma solidity >0.7.0 < 0.9.0;
/**
* @title Storage
* @dev store or retrieve a variable value
*/

contract Storage {

    uint256 value;

    function increaseValue(uint256 number) public{
        value = value + number;
    }

    function store(uint256 number) public{
        value = number;
    }

    function retrieve() public view returns (uint256){
        return value;
    }
}