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


    // CrowdSaleToken represents the contract of the token being sold.
    ReleasableSimpleCoin public CrowdSaleToken;


    // EventLog provides support for external logging.
    event EventLog(string value);

    // EventInvestment is an event to indicate an investment was made.
    event EventInvestment(address indexed investor, uint256 value);

    // EventTokenAssignment is an event to indicate a token was assigned.
    event EventTokenAssignment(address indexed investor, uint256 numTokens);


    // constructor is called when the contract is deployed.
    constructor(uint256 startTime, uint256 endTime, uint256 weiTokenPrice, uint256 etherInvestmentObjective) {
        require(startTime >= block.timestamp,  "start time must be greater than the current block timestamp");
        require(endTime >= startTime,          "end time must be greater than or equal to the startTime");
        require(weiTokenPrice != 0,            "wei token price must be greater than zero");
        require(etherInvestmentObjective != 0, "ether investment objective must be greater than zero");

        Owner                  = msg.sender;
        IsFinalized            = false;
        IsRefundingAllowed     = false;
        StartTime              = startTime;
        EndTime                = endTime;
        WeiTokenPrice          = weiTokenPrice;
        WeiInvestmentObjective = etherInvestmentObjective * 1000000000000000000;
        CrowdSaleToken         = new ReleasableSimpleCoin(0);
    }


    // Restricts functions to only be accessed by the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }


    // Invest allows an investor to book crowdsale tokens. (No parameter is necessary
    // to specify the amount of Ether being invested because it’s being sent through
    // the msg.value property.
    function Invest() public payable {
        require(isValidInvestment(msg.value), "investment value provided is not valid");

        address investor   = msg.sender;
        uint256 investment = msg.value;

        InvestmentAmountOf[investor] += investment;
        InvestmentReceived           += investment;

        assignTokens(investor, investment);

        emit EventInvestment(investor, investment);
        emit EventLog(string.concat("Invest: investor: ", Error.Addrtoa(investor), " investment: ", Error.Itoa(investment)));
    }

    // Finalize allows the crowdsale organizer, who is the contract owner, to release
    // tokens to the investors, in case of successful completion, and grant a bonus
    // to the development team, if applicable.
    function Finalize() onlyOwner public {
    }

    // Refund allows an investor to get a refund in case of unsuccessful completion.
    function Refund() public {
    }


    // isValidInvestment validates the specified investment.
    function isValidInvestment(uint256 investment) internal view returns (bool) {

        // Checks if this is a meaningful investment.
        bool nonZeroInvestment = investment != 0;

        // Check if this is taking place during the crowdsale funding stage.
        bool withinCrowdsalePeriod = block.timestamp >= StartTime && block.timestamp <= EndTime;
        
        return nonZeroInvestment && withinCrowdsalePeriod;
    }

    // assignTokens performs the token management.
    function assignTokens(address beneficiary, uint256 investment) internal {

        // Calculates the number of tokens corresponding to the investment.
        uint256 numberOfTokens = calculateNumberOfTokens(investment);

        // Generates the tokens in the investor account.
        CrowdSaleToken.Mint(beneficiary, numberOfTokens);
    }

    // calculateNumberOfTokens uses the WeiTokenPrice to calculate the number of
    // tokens being pruchased by this investment.
    function calculateNumberOfTokens(uint256 investment) internal view returns (uint256) {
        return investment / WeiTokenPrice;
    }
}
