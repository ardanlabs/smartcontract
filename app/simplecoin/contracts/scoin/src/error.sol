// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library Error {
    struct Err {
        bool   isError;
        string msg;
    }

    function New(string memory message) internal pure returns (Err memory) {
        return Err({isError: true, msg: message});
    }

    function None() internal pure returns (Err memory) {
        return Err({isError: false, msg: ""});
    }
}
