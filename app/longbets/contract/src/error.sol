// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library Error {
    struct Err {
        bool   isError;
        string msg;
    }

    // None returns an error value indicating no error.
    function None() internal pure returns (Err memory) {
        return Err({isError: false, msg: ""});
    }

    // New constructs an error value with a set of varadic parameters.
    function New(string memory a, string memory b, string memory c, string memory d) internal pure returns (Err memory) {
        return Err({isError: true, msg: string.concat(a, b, c, d)});
    }

    // New constructs an error value with a set of varadic parameters.
    function New(string memory a, string memory b, string memory c) internal pure returns (Err memory) {
        return Err({isError: true, msg: string.concat(a, b, c)});
    }

    // New constructs an error value with a set of varadic parameters.
    function New(string memory a, string memory b) internal pure returns (Err memory) {
        return Err({isError: true, msg: string.concat(a, b)});
    }

    // New constructs an error value with a set of varadic parameters.
    function New(string memory a) internal pure returns (Err memory) {
        return Err({isError: true, msg: a});
    }

    // Itoa converts an unsigned integer to a string.
    function Itoa(uint i) internal pure returns (string memory) {
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

    // Addrtoa converts an address to a string.
    function Addrtoa(address x) internal pure returns (string memory) {
        bytes memory s = new bytes(40);
        
        for (uint i = 0; i < 20; i++) {
            bytes1 b = bytes1(uint8(uint(uint160(x)) / (2**(8*(19 - i)))));
            bytes1 hi = bytes1(uint8(b) / 16);
            bytes1 lo = bytes1(uint8(b) - 16 * uint8(hi));
            s[2*i] = char(hi);
            s[2*i+1] = char(lo);            
        }
        
        return string(s);
    }

    // char is a support function for Addresstoa.
    function char(bytes1 b) internal pure returns (bytes1 c) {
        if (uint8(b) < 10) {
            return bytes1(uint8(b) + 0x30);
        }
        return bytes1(uint8(b) + 0x57);
    }
}
