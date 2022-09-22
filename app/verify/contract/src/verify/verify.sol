// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// Verify is a contract which can verify that a hashed message matches the
// provided signer.
contract Verify {

    // AddressFromMessage retrieves the signer's address from a hashed message
    // and signature's r, s, v values.
    function AddressFromMessage(bytes32 hashedMessage, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
        return ecrecover(hashedMessage, v, r, s);
    }

    // VerifySignature verifies that the signature's address matches the address
    // provided.
    function VerifySignature(bytes32 hashedMessage, uint8 v, bytes32 r, bytes32 s, address matches) public pure returns (bool) {
        return ecrecover(hashedMessage, v, r, s) == matches;
    }

}
