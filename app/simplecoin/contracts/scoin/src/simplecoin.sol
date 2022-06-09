// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract SimpleCoin {

    // =========================================================================
    
    mapping (address => uint256) public coinBalance;
    mapping (address => mapping (address => uint256)) public allowance;

    // =========================================================================

    event Log(string value);
    event Transfer(address indexed from, address indexed to, uint256 value);

    // =========================================================================

    constructor(uint256 initialSupply) {
        coinBalance[msg.sender] = initialSupply;
    }

    // =========================================================================

    function transfer(address to, uint256 amount) public {
        Error.Err memory err = validateTransfer(msg.sender, to, amount);
        if (err.isError) {
            revert(err.msg);
        }

        emit Log("starting transfer");
        {
            coinBalance[msg.sender] -= amount;
            coinBalance[to] += amount;
        }
        emit Log("ending transfer");

        emit Transfer(msg.sender, to, amount);
    }

    function transferFrom(address from, address to, uint256 amount) public {
        Error.Err memory err = validateTransfer(from, to, amount);
        if (err.isError) {
            revert(err.msg);
        }

        emit Log("starting transfer");
        {
            coinBalance[from] -= amount;
            coinBalance[to] += amount;
            allowance[from][msg.sender] -= amount;
        }
        emit Log("ending transfer");

        emit Transfer(from, to, amount);
    }

    function authorize(address authorizedAccount, uint256 allowanceAmount) public returns (bool) {
        allowance[msg.sender][authorizedAccount] = allowanceAmount;
        return true;
    }

    // =========================================================================

    function validateTransfer(address from, address to, uint256 amount) internal view returns (Error.Err memory) {
        
        // Check the caller isn't sending money to address zero.
        if (to == address(0)) {
            return Error.New("can't send money to address 0x0");
        }

        // Check that the sender has enough coins.
        if (coinBalance[from] < amount) {
            return Error.New(string(abi.encodePacked("insufficent funds  bal:", uint2str(coinBalance[msg.sender]), "  amount: ", uint2str(amount))));
        }

        // Check the amount is not zero or the amount value caused the unsigned
        // int to restart back to zero.
        if (coinBalance[to]+amount <= coinBalance[to]) {
            return Error.New(string(abi.encodePacked("invalid amount: ", uint2str(amount))));
        }

        return Error.None();
    }
 
    function uint2str(uint i) internal pure returns (string memory) {
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