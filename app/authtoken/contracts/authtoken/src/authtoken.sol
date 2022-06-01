// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract AuthorizedToken {

    // Enumerations
    enum UserType {TokenHolder, Admin, Owner}

    // User defined types
    struct AccountInfo {
        address account;
        string firstName;
        string lastName;
        UserType userType;
        bool exists;
    }

    // State variable definitions
    mapping (address => uint256) public tokenBalances;
    mapping (address => AccountInfo) public registeredAccounts;
    mapping (address => bool) public frozenAccounts;
    address public owner;

    // Constants
    uint256 public constant maxTransferLimit = 15000;

    // Events
    event Transfer(address indexed from, address indexed to, uint256 value);
    event FrozenAccount(address target, bool frozen);

    // onlyOwner adds a prerequisite that only the owner can execute the
    // function this is attached to.
    modifier onlyOwner {
        if (msg.sender != owner) {
            revert("only the owner can perform this operation");
        }
        _;
    }

    // Constructor that is called when the contract is deployed.
    constructor(uint256 initialSupply) {
        owner = msg.sender;
        registerAccount(msg.sender, "owner", "owner", true);
        mintToken(owner, initialSupply);
    }

    // =========================================================================

    // registerAccount adds the specified account to the registeredAccounts
    // map. Only registered account can be used in the system.
    function registerAccount(address account, string memory firstName, string memory lastName, bool isAdmin) public onlyOwner {
        if (registeredAccounts[account].exists) {
            revert("account already registered");
        }
        
        UserType userType = UserType.TokenHolder;
        if (isAdmin) {
            userType = UserType.Admin;
        }
        
        AccountInfo memory accountInfo = AccountInfo(
            {
                account: account,
                firstName: firstName,
                lastName: lastName,
                userType: userType,
                exists: true
            }
        );
        
        registeredAccounts[account] = accountInfo;
    }

    // mintToken gives the recipient the specified amount of coins. This function
    // can only be executed by the owner of this contract.
    function mintToken(address recipient, uint256 mintedAmount) public onlyOwner {
        if (!registeredAccounts[recipient].exists) {
            revert("account not registered");
        }

        tokenBalances[recipient] += mintedAmount;
       
        emit Transfer(owner, recipient, mintedAmount);
    }

    // transfer allows the sender to transfer coins from their account to the
    // specified account.
    function transfer(address to, uint256 amount) external {
        if (amount > maxTransferLimit) {
            revert("amount specified is greater than max transfer limit of 15000");
        }
        if (!registeredAccounts[msg.sender].exists) {
            revert("sending account not registered");
        }
        if (!registeredAccounts[to].exists) {
            revert("to account not registered");
        }

        tokenBalances[msg.sender] -= amount;
        tokenBalances[to] += amount;
        
        emit Transfer(msg.sender, to, amount);
    }

    // freezeAccount sets the target account into a freeze state. This function
    // can only be executed by the owner of this contract.
    function freezeAccount(address target, bool freeze) external onlyOwner {
        if (!registeredAccounts[target].exists) {
            revert("target account not registered");
        }

        frozenAccounts[target] = freeze;
       
        emit FrozenAccount(target, freeze);
    }
}