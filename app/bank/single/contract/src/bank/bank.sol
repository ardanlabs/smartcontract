// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract Bank {

    // Owner represents the address who deployed the contract.
    address public Owner;

    // Version is the current version of Store.
    string public Version;

    // accountBalances represents the amount of money an account has available.
    mapping (address => uint256) private accountBalances;

    // EventLog provides support for external logging.
    event EventLog(string value);

    // =========================================================================

    // constructor is called when the contract is deployed.
    constructor() {
        Owner = msg.sender;
        Version = "0.1.0";
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

        // Add the ante for each player to the pot. The initialization is
        // for the winner's ante.
        uint256 pot = anteWei;
        for (uint i = 0; i < losers.length; i++) {
            if (accountBalances[losers[i]] < anteWei) {
                emit EventLog(string.concat("account balance ", Error.Itoa(accountBalances[losers[i]]), " is less than ante ", Error.Itoa(anteWei)));
                pot += accountBalances[losers[i]];
                accountBalances[losers[i]] = 0;
            } else {
                pot += anteWei;
                accountBalances[losers[i]] -= anteWei;
            }
        }

        emit EventLog(string.concat("ante[", Error.Itoa(anteWei), "] gameFee[", Error.Itoa(gameFeeWei), "] pot[", Error.Itoa(pot), "]"));

        // This should not happen but check to see if the pot is 0 because none
        // of the losers had an account balance.
        if (pot == 0) {
            revert("no pot was created based on account balances");
        }

        // This should not happen but check there is enough in the pot to cover
        // the game fee.
        if (pot < gameFeeWei) {
            accountBalances[Owner] += pot;
            emit EventLog(string.concat("pot less than fee: winner[0] owner[", Error.Itoa(pot), "]"));
            return;
        }

        // Take the game fee from the pot and give the winner the remaining pot
        // and the owner the game fee.
        pot -= gameFeeWei;
        accountBalances[winner] += pot;
        accountBalances[Owner] += gameFeeWei;
        emit EventLog(string.concat("winner[", Error.Itoa(pot), "] owner[", Error.Itoa(gameFeeWei), "]"));
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
