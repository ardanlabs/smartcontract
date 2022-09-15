// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

/*
██████╗  █████╗ ███╗   ██╗██╗  ██╗██████╗ ██████╗  ██████╗ ██╗  ██╗██╗   ██╗
██╔══██╗██╔══██╗████╗  ██║██║ ██╔╝██╔══██╗██╔══██╗██╔═══██╗╚██╗██╔╝╚██╗ ██╔╝
██████╔╝███████║██╔██╗ ██║█████╔╝ ██████╔╝██████╔╝██║   ██║ ╚███╔╝  ╚████╔╝
██╔══██╗██╔══██║██║╚██╗██║██╔═██╗ ██╔═══╝ ██╔══██╗██║   ██║ ██╔██╗   ╚██╔╝
██████╔╝██║  ██║██║ ╚████║██║  ██╗██║     ██║  ██║╚██████╔╝██╔╝ ██╗   ██║
╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝

The "proxy" contract model uses two contracts: a dedicated Proxy and an
external API. The Proxy calls out to the API for business logic, creating
a system that can be upgraded without losing the Proxy contract's stored
data by updating the address of the external API within the Proxy.

                                                       ┌────────────┐
                                                       │            │
                                                 ┌────►│   API_v1   │
                                                 │     │            │
                                                 │     └────────────┘
                                                 │
 ┌────────────┐       ┌───────────┐              │     ┌────────────┐
 │            │       │           │              │     │            │
 │  Customer  ├──TX──►│   Proxy   ├─delegatecall─┼────►│   API_v2   │
 │            │       │           │              │     │            │
 └────────────┘       └───────────┘              │     └────────────┘
                                                 │
                                                 │     ┌────────────┐
                                                 │     │            │
                                                 └────►│   API_v3   │
                                                       │            │
                                                       └────────────┘

A caveat is that the storage models for both the Proxy and the API must match
almost exactly. This is because the methods in the API expect the variables
it interacts with to be at certain points in the contract's storage.

Smart contract storage is a largely "slots" based system. Each slot can
potentially contain multiple variables, if they are small enough to fit, but
they will always be allocated according to the order in which they are declared
and not via any automated optimization strategy. This makes the allocations
predictable.

Because they are predictable, as long as the same variables are declared in the
same order for both the Proxy and the API, they will match up for the purposes
of the delegatecall.

+--------------+------------+------------+
| Storage Slot | Proxy      | API        |
+--------------+------------+------------+
| Slot 1       | Variable 1 | Variable 1 |
+--------------+------------+------------+
| Slot 2       | Variable 2 | Variable 2 |
+--------------+------------+------------+
| Slot 3       | Variable 3 | Variable 3 |
+--------------+------------+------------+

A mismatch will cause the delegatecall to reference the wrong location in storage,
possibly corrupting the data.

+--------------+------------+------------+
| Storage Slot | Proxy      | API        |
+--------------+------------+------------+
| Slot 1       | Variable 1 | Variable 1 |
+--------------+------------+------------+
| Slot 2       | Variable 2 | Variable 3 | <- Mismatch, delegatecall trying to reach
+--------------+------------+------------+    Variable 3 will instead be working with
| Slot 3       | Variable 3 |            |    Variable 2.
+--------------+------------+------------+

The API contract should be deployed first, so that it is available to the Proxy when
it is deployed. Other API versions can be switched to via the Proxy's UpgradeAPI
function.

*/

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