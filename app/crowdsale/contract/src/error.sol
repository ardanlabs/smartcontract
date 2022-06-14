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
}
