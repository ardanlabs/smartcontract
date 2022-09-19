// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract Bank {

    // API represents the address of the API contract.
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
    constructor(address api) {
        API = api;
        Owner = msg.sender;

        (bool success, bytes memory data) = API.call(abi.encodeWithSignature("Version()"));
        emit EventLog(string.concat("upgrade[",Error.Addrtoa(api),"] success[", Error.Booltoa(success), "] version[", string(data), "]"));
    }

    // =========================================================================
    // Owner Only Calls

    // onlyOwner can be used to restrict access to a function for only the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // UpgradeAPI upgrades the API to a new version, or rollback to a previous one.
    function UpgradeAPI(address api) onlyOwner public {
        API = api;

        (bool success, bytes memory data) = API.call(abi.encodeWithSignature("Version()"));
        emit EventLog(string.concat("upgrade[",Error.Addrtoa(api),"] success[", Error.Booltoa(success), "] version[", string(data), "]"));
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

        emit EventLog(string.concat("reconcile[", Error.Booltoa(success), "]"));
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
