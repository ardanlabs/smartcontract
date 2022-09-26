// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

// Verify is a contract which can verify that a hashed message matches the
// provided signer.
contract Verify {

    // Address retrieves the signer's address from a signature that was produced
    // based on the provided data that was signed.
    function Address(bytes32 data, bytes calldata sig) public pure returns (address) {
        (address addr,Error.Err memory err) = extractAddress(data, sig);
        if (err.isError) {
            revert(err.msg);
        }

        return addr;
    }

    // MatchSender checks the address of the sender was the address that signed
    // the specified data with the provided signature.
    function MatchSender(bytes32 data, bytes calldata sig) public view returns (bool) {
        (address addr,Error.Err memory err) = extractAddress(data, sig);
        if (err.isError) {
            revert(err.msg);
        }

        if (addr == msg.sender) {
            return true;
        }

        return false;
    }

    // extractAddress expects the raw data that was signed and will apply the Ethereum
    // salt value manually. This hides the underlying implementation of the salt.
    function extractAddress(bytes32 data, bytes calldata sig) private pure returns (address, Error.Err memory) {
        if (sig.length != 65) {
            return (address(0), Error.New("invalid signature length"));
        }

        bytes32 r = bytes32(sig[:32]);
        bytes32 s = bytes32(sig[32:64]);
        uint8 v = uint8(sig[64]);

        return (ecrecover(data, v, r, s), Error.None());
    }
}
