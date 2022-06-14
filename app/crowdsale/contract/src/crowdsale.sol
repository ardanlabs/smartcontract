// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./simplecoin.sol";

contract SimpleCrowdsale {

    // Owner represents the address who deployed the contract.
    address public Owner;

    // IsFinalized indicates if the contract has been finalized.
    bool public IsFinalized;

    // IsRefundingAllowed indicates whether refunding is allowed.
    bool public IsRefundingAllowed;

    // StartTime represents the start of the crowdsale funding stage in UNIX epoch.
    uint256 public StartTime;

    // EndTime represents the end of the crowdsale funding stage in UNIX epoch.
    uint256 public EndTime;

    // WeiTokenPrice represents the price of the token being sold.
    uint256 public WeiTokenPrice;

    // WeiInvestmentObjective represents the investment objective, which defines if the crowdsale is successful.
    uint256 public WeiInvestmentObjective;


    // InvestmentAmountOf represents the amount of Ether received from each investor.
    mapping (address => uint256) public InvestmentAmountOf;

    // InvestmentReceived represents the total Ether received from the investors.
    uint256 public InvestmentReceived;

    // InvestmentRefunded represents the total Ether refunded to the investors.
    uint256 public InvestmentRefunded;


    // CrowdsaleToken represents the contract of the token being sold.
    SimpleCoin public CrowdsaleToken;


    // Restricts functions to only be accessed by the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }


    // Invest allows an investor to book crowdsale tokens. (No parameter is necessary
    // to specify the amount of Ether being invested because it’s being sent through
    // the msg.value property.
    function Invest(address beneficiary) public payable {
    }

    // Finalize allows the crowdsale organizer, who is the contract owner, to release
    // tokens to the investors, in case of successful completion, and grant a bonus
    // to the development team, if applicable.
    function Finalize() onlyOwner public {
    }

    // Refund allows an investor to get a refund in case of unsuccessful completion.
    function Refund() public {
    }
}
