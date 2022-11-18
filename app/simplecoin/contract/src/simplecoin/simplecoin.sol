// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract SimpleCoin {

    // Owner represents the address who deployed the contract.
    address public Owner;

    // CoinBalance manages the accounts and their funding.
    mapping (address => uint256) public CoinBalance;

    // FrozenAccount represents accounts who can't spend any longer.
    mapping (address => bool) public FrozenAccount;

    // =========================================================================

    // EventLog provides support for external logging.
    event EventLog(string value);

    // EventTransfer is an event to indicate a transfer was performed.
    event EventTransfer(address indexed from, address indexed to, uint256 amount);

    // EventFrozenAccount is an event to indicate an account was frozen or
    // unfrozen.
    event EventFrozenAccount(address target, bool frozen);

    // =========================================================================

    // constructor is called when the contract is deployed.
    constructor(uint256 initialSupply) {
        Owner = msg.sender;
        Mint(Owner, initialSupply);
    }

    // =========================================================================
    // Owner Only Calls

    // onlyOwner can be used to restrict access to a function for only the
    // owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // Mint grants the recipient an amount of new coins. This can only be called
    // by the owner of the contract.
    function Mint(address recipient, uint256 mintedAmount) onlyOwner public {
        CoinBalance[recipient] += mintedAmount;
        emit EventTransfer(Owner, recipient, mintedAmount);
    }

    // FreezeAccount prevents this account from performing any activity. This
    // can only be called by the owner of the contract.
    function FreezeAccount(address target, bool freeze) onlyOwner public {
        FrozenAccount[target] = freeze;
        emit EventFrozenAccount(target, freeze);
    }

    // =========================================================================
    // General Public Calls

    // Transfer moves coins from the sender to the specified account.
    function Transfer(address to, uint256 amount) public {
        Error.Err memory err = validateTransfer(msg.sender, to, amount);
        if (err.isError) {
            revert(err.msg);
        }

        CoinBalance[msg.sender] -= amount;
        CoinBalance[to]         += amount;

        emit EventLog(string.concat("transfered ", Error.Itoa(amount), " to ", Error.Addrtoa(to), " from ", Error.Addrtoa(msg.sender)));
        emit EventTransfer(msg.sender, to, amount);
    }

    // =========================================================================
    // Internal Only Calls

    // validateTransfer performs checks for transfer operations.
    function validateTransfer(address from, address to, uint256 amount) internal view returns (Error.Err memory) {

        // Check the caller isn't sending money to address zero.
        if (to == address(0)) {
            return Error.New("can't send money to address 0x0");
        }

        // Check that the sender has enough coins.
        if (CoinBalance[from] < amount) {
            return Error.New("insufficent funds  bal:", Error.Itoa(CoinBalance[msg.sender]), "  amount: ", Error.Itoa(amount));
        }

        // Check the amount is not zero or the amount value caused the unsigned
        // int to restart back to zero.
        if (CoinBalance[to]+amount <= CoinBalance[to]) {
            return Error.New("invalid amount: ", Error.Itoa(amount));
        }

        return Error.None();
    }
}
