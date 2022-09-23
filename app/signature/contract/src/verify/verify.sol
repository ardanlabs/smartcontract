// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// Verify is a contract which can verify that a hashed message matches the
// provided signer.
contract Verify {

    // Address retrieves the signer's address from a signature that was produced
    // based on the provided data that was signed.
    function Address(bytes memory data, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
        return signature(data, v, r, s);
    }

    // MatchSender checks the address of the sender was the address that signed
    // the specified data with the provided signature.
    function MatchSender(bytes memory data, uint8 v, bytes32 r, bytes32 s) public view returns (bool) {
        address sig = signature(data, v, r, s);
        if (sig == msg.sender) {
            return true;
        }
        return false;
    }

    // signature expects the raw data that was signed and will apply the Ethereum
    // salt value manually. This hides the underlying implementation of the salt.
    function signature(bytes memory data, uint8 v, bytes32 r, bytes32 s) private pure returns (address) {
        bytes32 hashedData = keccak256(data);
        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 saltedData = keccak256(abi.encodePacked(prefix, hashedData));
        return ecrecover(saltedData, v, r, s);
    }
}
