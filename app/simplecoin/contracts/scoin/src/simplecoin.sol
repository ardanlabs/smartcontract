// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract SimpleCoin {
    mapping (address => uint256) public coinBalance;

    event Log(string value);
    event Transfer(address indexed from, address indexed to, uint256 value);

    constructor(uint256 initialSupply) {
        coinBalance[msg.sender] = initialSupply;
    }

    function transfer(address to, uint256 amount) external {
        if (coinBalance[msg.sender] < amount) {
            string memory resp = concatenate("insufficent funds  bal:", uint2str(coinBalance[msg.sender]));
            resp = concatenate(resp, " amount:");
            resp = concatenate(resp, uint2str(amount));
            revert(resp);
        }

        emit Log("starting transfer");

        coinBalance[msg.sender] -= amount;
        coinBalance[to] += amount;

        emit Log("ending transfer");

        emit Transfer(msg.sender, to, amount);
    }

    function concatenate(string memory a, string memory b) public pure returns (string memory) {
        return string(bytes.concat(bytes(a), " ", bytes(b)));
    }

    function uint2str(uint i) internal pure returns (string memory _uintAsString) {
        if (i == 0) {
            return "0";
        }

        uint j = i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }

        bytes memory bstr = new bytes(len);
        uint k = len;
        while (i != 0) {
            k = k-1;
            uint8 temp = (48 + uint8(i - i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            i /= 10;
        }

        return string(bstr);
    }
}