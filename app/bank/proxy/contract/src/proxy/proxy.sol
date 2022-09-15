// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../error.sol";

contract BankProxy {

    // Owner represents the address who deployed the contract.
    address public Owner;

    // API represents the address of the API contract.
    address public API;

    // accountBalances represents the amount of money an account has available.
    mapping (address => uint256) private accountBalances;

    // EventLog provides support for external logging.
    event EventLog(string value);

    // =========================================================================

    // constructor is called when the contract is deployed.
    constructor(address api) {
        Owner = msg.sender;
        API = api;
    }

    // =========================================================================
    // Owner Only Calls

    // onlyOwner can be used to restrict access to a function for only the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // Reconcile settles the accounting for a game that was played.
    function Reconcile(address winner, address[] calldata losers, uint256 anteWei, uint256 gameFeeWei) onlyOwner public {
        // Calls the API contract Reconcile function.
        (bool success,) = API.delegatecall(
            abi.encodeWithSignature("Reconcile(address, address[] calldata, uint256, uint256)", winner, losers, anteWei, gameFeeWei)
        );
        emit EventLog(string.concat("reconcile[", Error.Booltoa(success), "]"));
    }

    // UpgradeAPI upgrades the API to a new version, or rollback to a previous one.
    function UpgradeAPI(address api) onlyOwner public {
        API = api;
        emit EventLog(string.concat("upgrade[",Error.Addrtoa(api),"]"));
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
        accountBalances[msg.sender] += msg.value;
        emit EventLog(string.concat("deposit[", Error.Addrtoa(msg.sender), "] balance[", Error.Itoa(accountBalances[msg.sender]), "]"));
    }

    // Withdraw the given amount from the account balance.
    function Withdraw() payable public {
        address payable account = payable(msg.sender);

        if (accountBalances[msg.sender] == 0) {
            revert("not enough balance");
        }

        uint256 amount = accountBalances[msg.sender];
        account.transfer(amount);        
        accountBalances[msg.sender] = 0;

        emit EventLog(string.concat("withdraw[", Error.Addrtoa(msg.sender), "] amount[", Error.Itoa(amount), "]"));
    }
}
