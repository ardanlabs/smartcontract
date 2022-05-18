// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract SimpleCoin {
    mapping (address => uint256) public coinBalance;

    event Transfer(address indexed from, address indexed to, uint256 value);

    constructor(uint256 initialSupply) {
        coinBalance[msg.sender] = initialSupply;
    }

    function transfer(address to, uint256 amount) external {
        require(coinBalance[msg.sender] > amount);
        require(coinBalance[to] + amount >= coinBalance[to]);
        
        coinBalance[msg.sender] -= amount;
        coinBalance[to] += amount;

        emit Transfer(msg.sender, to, amount);
    }
}