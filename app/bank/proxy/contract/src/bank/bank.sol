// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../error.sol";

contract Bank {

    // API represents the address of the contract to use for delagate calls.
    address public API;

    // Version is the current version of Store.
    string public Version;

    // Owner represents the address who deployed the contract.
    address public Owner;

    // accountBalances represents the amount of money an account has available.
    mapping (address => uint256) private accountBalances;

    // EventLog provides support for external logging.
    event EventLog(string value);

    // =========================================================================

    // constructor is called when the contract is deployed.
    constructor() {
        Owner = msg.sender;
    }

    // =========================================================================
    // Owner Only Calls

    // onlyOwner can be used to restrict access to a function for only the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // SetContract points the bank to the contract to use for logic.
    function SetContract(address contractAddr) onlyOwner public {
        API = contractAddr;

        (bool success, bytes memory data) = API.call(abi.encodeWithSignature("Version()"));
        if (success) {
            Version = string(abi.decode(data, (string)));
        } else {
            Version = "unknown";
        }

        emit EventLog(string.concat("contract[",Error.Addrtoa(API),"] success[", Error.Booltoa(success), "] version[", Version, "]"));
    }

    // Reconcile settles the accounting for a game that was played.
    function Reconcile(address winner, address[] memory losers, uint256 anteWei, uint256 gameFeeWei) onlyOwner public {

        // Calls the API contract's Reconcile function.
        // Notice the string represents the function call with the declared
        // parameters followed by the parameters to pass. You can't have spaces
        // between the commas in the function declaration.
        (bool success,) = API.delegatecall(
            abi.encodeWithSignature("Reconcile(address,address[],uint256,uint256)", winner, losers, anteWei, gameFeeWei)
        );

        emit EventLog(string.concat("success[", Error.Booltoa(success), "]"));
    }

    // AccountBalance returns the current account's balance.
    function AccountBalance(address account) onlyOwner view public returns (uint) {
        return accountBalances[account];
    }

    // =========================================================================
    // Account Only Calls

    // Balance returns the balance of the caller.
    function Balance() view public returns (uint) {
        return accountBalances[msg.sender];
    }

    // Deposit the given amount to the account balance.
    function Deposit() payable public {
        (bool success,) = API.delegatecall(
            abi.encodeWithSignature("Deposit()")
        );

        emit EventLog(string.concat("success[", Error.Booltoa(success), "]"));
    }

    // Withdraw the given amount from the account balance.
    function Withdraw() payable public {
        (bool success,) = API.delegatecall(
            abi.encodeWithSignature("Withdraw()")
        );

        emit EventLog(string.concat("success[", Error.Booltoa(success), "]"));
    }
}
