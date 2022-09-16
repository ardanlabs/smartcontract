// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// Store is a contract which implements a basic key/value store.
contract Store {

    // Version is the current version of Store.
    string public Version;

    // Items is a mapping which will store the key/value data.
    mapping (bytes32 => bytes32) public Items;

    // ItemSet is an event which will output any updates to the key/value
    // data to the transaction receipt's logs.
    event ItemSet(bytes32 key, bytes32 value);

    // The constructor is automatically executed when the contract is deployed.
    constructor() {
        Version = "1.1";
    }

    // SetItem is an external-only function which accepts a key/value pair
    // and updates the contract's internal storage accordingly.
    function SetItem(bytes32 key, bytes32 value) external {
        Items[key] = value;
        emit ItemSet(key, value);
    }
}
