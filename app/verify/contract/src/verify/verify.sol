// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// Verify is a contract which can verify that a hashed message matches the
// provided signer.
contract Verify {

    // VerifyMessage verifies that the hashed message's signer matches the
    // provided address.
    function VerifyMessage(bytes32 hashedMessage, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
        // Commenting these lines because hashedMessage is already stamped.

        // bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        // bytes32 prefixedHashMessage = keccak256(abi.encodePacked(prefix, hashedMessage));

        address signer = ecrecover(hashedMessage, v, r, s);
        return signer;
    }

}
