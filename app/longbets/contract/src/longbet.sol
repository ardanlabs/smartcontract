// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

/*

██╗      ██████╗ ███╗   ██╗ ██████╗     ██████╗ ███████╗████████╗███████╗
██║     ██╔═══██╗████╗  ██║██╔════╝     ██╔══██╗██╔════╝╚══██╔══╝██╔════╝
██║     ██║   ██║██╔██╗ ██║██║  ███╗    ██████╔╝█████╗     ██║   ███████╗
██║     ██║   ██║██║╚██╗██║██║   ██║    ██╔══██╗██╔══╝     ██║   ╚════██║
███████╗╚██████╔╝██║ ╚████║╚██████╔╝    ██████╔╝███████╗   ██║   ███████║
╚══════╝ ╚═════╝ ╚═╝  ╚═══╝ ╚═════╝     ╚═════╝ ╚══════╝   ╚═╝   ╚══════╝

When a predictor believes a noteworthy event will happen, they may assert
their prediction with a stake of eth. This will create a new LongBet smart
contract, which can then be used to manage the prediction.

The Long Bet process works as follows:

1. The predictor calls LongBets.Bet with their prediction, a list of proposed
   mediators, and a proposed ruleset. This creates a new LongBet contract.
2. The predictor loads their stakes into the new LongBet contract.
3. A challenger joins the bet by loading a matching stake into the LongBet
   contract.
4. The predictor and challenger negotiate the terms of the bet by proposing:
   a. Alternative rulesets.
   b. Alternative mediator tip percentage.
5. Once the predictor and challenger have matching proposals, they both seek
   approval from a mediator.
6. A mediator approves the bet, solidifying the terms.
7. Once the terms have been fulfilled, a mediator rules in favor of the
   winning participant, allocating the pooled stakes and taking their tip.

Any single mediator may perform a mediator action.

To encourage mediator participation, each mediator action will award the acting
mediator with a tip. Tip default is 1% of total stakes pool unless a different
amount is otherwise agreed-upon by participants. Because mediators will need to
act a minimum of two times, the tip percentage will be subtracted from total
stakes at least twice (applied as percentage of current balance each time).

Potential improvements:
  - Allow for multiple challengers/predictors
  - Possibly allow the predictor to reject a challenger
  - Needs thought put into potential attack vectors, some possibilities may
    exist of locking up a predictor's funds until the expiry date.
  - Possibly allow a mediator to remove a challenger if they are being too
    adversarial with their rules proposals?
*/

// LongBet contract will manage individual bets.
contract LongBet {

    // ===========================================================================
    // Structs

        // Participant defines a person who participates in a longbet.
        struct Participant {

        // The address of the participant.
        address payable addr;

        // The stakes put forth by the participant.
        uint256 stakes;

        // The rules of the bet, proposed by the participant.
        string[] proposedRules;

        // The proposed mediator tip, as a percentage of total stakes.
        uint proposedTip;

        // When the participant is satisfied with the stakes and terms of the bet,
        // they will set this value to true.
        bool seeksApproval;
    }

    // ===========================================================================
    // Contract Level Variables

    // The address of the contract which owns and manages this LongBet.
    address public owner;

    // The primary basis of the bet, e.g.
    // "Man will walk on Mars by the year 2030."
    string public prediction;

    // A series of mutually agreed-upon criteria which must be met in order to
    // satisfy the conditions of the prediction, e.g.
    // "A living person must walk on the surface of the planet Mars."
    // "Must occur before 2030-01-01T00:00:00Z."
    // "The person need not successfully return to Earth."
    string[] public rules;

    // The predictor is the one to propose the bet with the initial prediction,
    // rules, stakes, and mediator.
    Participant public predictor;

    // The challenger is the one to challenge the predictor's prediction, with
    // a matching stake.
    Participant public challenger;

    // When the mediator has approved the stakes and rules, the bet is set and
    // cannot be altered any further by the predictor or challenger.
    bool public mediatorApproved;

    // The date by which the bet will expire. If the terms of the bet cannot be
    // adequately determined or the mediator does not take action by this date
    // then the participants may trigger a refund of their stakes.
    uint public expires;

    // The tip percentage to be disbursed to the mediator in exchange for their
    // services.
    uint public mediatorTip;

    // ===========================================================================
    // Modifiers

    // onlyPredictor can be used to restrict access to a function to the predictor.
    modifier onlyPredictor {
        if (msg.sender != predictor.addr) {
            revert("only the predictor may call this function");
        }
        _;
    }

    // notPredictor prevents the predictor from accessing some functions.
    modifier notPredictor {
        if (msg.sender == predictor.addr) {
            revert("the predictor may not call this function");
        }
        _;
    }

    // onlyChallenger can be used to restrict access to a function to the challenger.
    modifier onlyChallenger {
        if (msg.sender != challenger.addr) {
            revert("only the challenger may call this function");
        }
        _;
    }

    // onlyOwner can be used to restrict access to a function to the owning contract,
    // which is itself managed by a team of mediators.
    modifier onlyOwner {
        if (msg.sender != owner) {
            revert("only the mediator may call this function");
        }
        _;
    }

    // onlyBeforeMediatorApproved can be used to restrict functions to the "setup"
    // stages of the betting process, before the mediator has approved the final
    // terms.
    modifier onlyBeforeMediatorApproval {
        if (mediatorApproved) {
            revert("this function may only be called prior to mediator approval");
        }
        _;
    }

    // onlyAfterMediatorApproval prevents the mediator from performing certain
    // actions before the bet is appropriately agreed-upon by predictor and challenger.
    modifier onlyAfterMediatorApproval {
        if (!mediatorApproved) {
            revert("this function may only be called after mediator approval");
        }
        _;
    }

    modifier matchingStakes {
        if (predictor.stakes != challenger.stakes) {
            revert("this function may only be called when predictor and challenger have matching stakes");
        }
        _;
    }

    modifier matchingRules {

        // Comparing hashes is more gas-efficient than walking the arrays and
        // comparing individual entries.
        if (keccak256(abi.encode(predictor.proposedRules)) != keccak256(abi.encode(challenger.proposedRules))) {
            revert("this function may only be called when predictor and challenger have matching rules");
        }
        _;
    }

    modifier matchingTips {
        if (predictor.proposedTip != challenger.proposedTip) {
            revert("the participants do not have matching proposed tip percentages");
        }
        _;
    }

    modifier matchingApprovals {
        if (!predictor.seeksApproval || !challenger.seeksApproval) {
            revert("this function may only be called when predictor and challenger are both seeking mediator approval");
        }
        _;
    }

    // onlyAfterExpired can be used to restrict a function to be callable only
    // after the bet has expired.
    modifier onlyAfterExpired {
        if (block.timestamp < expires) {
            revert("the bet has not yet expired");
        }
        _;
    }

    // ===========================================================================
    // Constructor

    constructor(address _owner, address payable _predictor, string memory _prediction, string[] memory _rules, uint _expires) {
        owner = _owner;
        predictor = Participant(_predictor, 0, _rules, 1, false);
        prediction = _prediction;
        expires = _expires;
    }

    // ===========================================================================
    // Manage stakes

    // predictorAddStakes allows the predictor to add stakes. Newly added stakes
    // are automatically stored in the contract's balance. May be called multiple
    // times, as raising stakes may be necessary to attract a challenger.
    function predictorAddStakes() onlyPredictor onlyBeforeMediatorApproval payable public {
        if (challenger.stakes > 0) {
            revert("predictor may not raise stakes after a challenger has appeared");
        }
        predictor.stakes += msg.value;
    }

    // challengePrediction allows a challenger to participate by adding their
    // stakes. Newly added stakes are automatically stored in the contract's balance.
    function challengePrediction() notPredictor onlyBeforeMediatorApproval payable public {
        if (msg.value != predictor.stakes) {
            revert("challenger must match predictor stakes");
        }

        if (challenger.addr != address(0)) {
            revert("this prediction already has a challenger");
        }
        
        string[] memory _rules;
        challenger = Participant(payable(msg.sender), msg.value, _rules, 1, false);
    }

    // getPool allows a view if the total pool.
    function getPool() public view returns (uint256) {
        return address(this).balance;
    }

    // ===========================================================================
    // Manage rules

    function predictorProposeRules(string[] calldata _rules) onlyPredictor onlyBeforeMediatorApproval public {
        predictor.proposedRules = _rules;
    }

    function challengerProposeRules(string[] calldata _rules) onlyChallenger onlyBeforeMediatorApproval public {
        challenger.proposedRules = _rules;
    }

    // ===========================================================================
    // Manage tips

    function predictorProposeTip(uint _tip) onlyPredictor onlyBeforeMediatorApproval public {
        predictor.proposedTip = _tip;
    }

    function challengerProposeTip(uint _tip) onlyChallenger onlyBeforeMediatorApproval public {
        challenger.proposedTip = _tip;
    }

    // ===========================================================================
    // Seek mediator approval

    function predictorSeekApproval() onlyPredictor onlyBeforeMediatorApproval matchingStakes matchingRules matchingTips public {
        predictor.seeksApproval = true;
    }

    function challengerSeekApproval() onlyChallenger onlyBeforeMediatorApproval matchingStakes matchingRules public {
        challenger.seeksApproval = true;
    }

    // ===========================================================================
    // Mediator actions, may only be called externally by owner contract

    function mediatorApproves(address payable _mediator) onlyOwner onlyBeforeMediatorApproval matchingStakes matchingRules matchingTips matchingApprovals external {
        mediatorApproved = true;
        rules = predictor.proposedRules;
        mediatorTip = predictor.proposedTip;

        // Mediator gets their tip.
        _mediator.transfer(address(this).balance * (100 / mediatorTip));
    }

    function mediatorRulesPredictionTrue(address payable _mediator) onlyOwner onlyAfterMediatorApproval external {

        // Mediator gets their tip first.
        _mediator.transfer(address(this).balance * (100 / mediatorTip));
        predictor.addr.transfer(address(this).balance);
    }

    function mediatorRulesPredictionFalse(address payable _mediator) onlyOwner onlyAfterMediatorApproval external {

        // Mediator gets their tip first.
        _mediator.transfer(address(this).balance * (100 / mediatorTip));
        challenger.addr.transfer(address(this).balance);
    }

    // ==========================================================================
    // Post-expiry refund requests

    function predictorRequestRefund() onlyPredictor onlyAfterExpired public {
        if (predictor.stakes == 0) {
            revert("predictor has already refunded stakes");
        }

        if (challenger.stakes == 0) {
            
            // Challenger has already refunded, predictor gets remaining balance.
            predictor.addr.transfer(address(this).balance);
        } else {
            
            // Challenger has not refunded, predictor gets half of balance.
            predictor.addr.transfer(address(this).balance / 2);
        }

        predictor.stakes = 0;
    }

    function challengerRequestRefund() onlyChallenger onlyAfterExpired public {
        if (challenger.stakes == 0) {
            revert("challenger has already refunded stakes");
        }

        if (predictor.stakes == 0) {
            
            // Predictor has already refunded, challenger gets remaining balance.
            challenger.addr.transfer(address(this).balance);
        } else {
            
            // Predictor has not refunded, challenger gets half of balance.
            challenger.addr.transfer(address(this).balance / 2);
        }

        challenger.stakes = 0;
    }
}

