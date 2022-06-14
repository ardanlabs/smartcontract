// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library Error {
    struct Err {
        bool   isError;
        string msg;
    }

    function None() internal pure returns (Err memory) {
        return Err({isError: false, msg: ""});
    }

    function New(string memory a, string memory b, string memory c, string memory d) internal pure returns (Err memory) {
        return Err({isError: true, msg: string.concat(a, b, c, d)});
    }

    function New(string memory a, string memory b, string memory c) internal pure returns (Err memory) {
        return Err({isError: true, msg: string.concat(a, b, c)});
    }

    function New(string memory a, string memory b) internal pure returns (Err memory) {
        return Err({isError: true, msg: string.concat(a, b)});
    }

    function New(string memory a) internal pure returns (Err memory) {
        return Err({isError: true, msg: a});
    }
}
