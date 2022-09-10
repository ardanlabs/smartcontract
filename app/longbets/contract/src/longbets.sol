// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./longbet.sol";
import "./error.sol";

// LongBets contract will track, deploy, and mediate all bets.
// Mediators may approve the addition of new mediators with a majority vote.
// Mediators may approve the removal of other mediators with a majority vote.
contract LongBets {

    // ===========================================================================
    // Contract Level Variables
  
    // A mapping of all known LongBet contracts managed by this contract.
    mapping (uint => LongBet) public deployedBets;
    uint public numBets;

    // A mapping of all approved mediators.
    mapping (address => bool) public mediators;
    uint public numMediators;

    // A mapping of newly proposed mediators and their vote counts.
    mapping (address => uint) public proposedMediators;

    // A mapping of votes in favor of approving a new proposed mediator.
    // [proposedMediator][mediator][vote]
    mapping (address => mapping(address => bool)) public proposedMediatorVotes;
    mapping (address => uint) public totalProposedMediatorVotes;

    // A mapping of proposed mediator removals and their vote counts.
    mapping (address => mapping(address => bool)) public proposedMediatorRemovals;
    mapping (address => uint) public totalProposedMediatorRemovalVotes;

    // ===========================================================================
    // Modifiers

    modifier mediatorOnly {
        if (!mediators[msg.sender]) {
            revert("must be a mediator");
        }
        _;
    }

    // ===========================================================================
    // Constructor

    constructor(address _creator) {
        // Add the creator of this contract as the initial mediator.
        mediators[_creator] = true;
    }

    // ===========================================================================
    // API

    // Creates a new LongBet contract.
    function bet(string calldata _prediction, string[] calldata _rules) public returns (address, uint) {
        deployedBets[numBets] = new LongBet(address(this), payable(msg.sender), _prediction, _rules, block.timestamp + (2 * 365 days));
        numBets++;
        
        return (address(deployedBets[numBets-1]), numBets-1);
    }

    function proposeMediator(address _mediator) mediatorOnly public {
        if (proposedMediatorVotes[_mediator][msg.sender] == true) {
            revert("you have already voted in favor of this mediator");
        }

        proposedMediatorVotes[_mediator][msg.sender] = true;
        totalProposedMediatorVotes[_mediator]++;

        // If more than half of existing mediators have approved of the new mediator,
        // add them to the mediators mapping.
        if (totalProposedMediatorVotes[_mediator] > numMediators/2) {
            mediators[_mediator] = true;
            numMediators++;
            
            // TODO: Resolve error with below line
            // delete proposedMediatorVotes[_mediator];
            totalProposedMediatorVotes[_mediator] = 0;
        }
    }

    function proposeMediatorRemoval(address _mediator) mediatorOnly public {
        if (proposedMediatorRemovals[_mediator][msg.sender] == true) {
            revert("you have already voted to remove this mediator");
        }

        proposedMediatorRemovals[_mediator][msg.sender] = true;
        totalProposedMediatorRemovalVotes[_mediator]++;

        // If more than half of existing mediators have approved of the removal of a
        // mediator, remove them from the mediators mapping.
        if (totalProposedMediatorRemovalVotes[_mediator] > numMediators/2) {
            delete mediators[_mediator];
            numMediators--;
            
            // TODO: Resolve error with below line.
            // delete proposedMediatorRemovals[_mediator];
            totalProposedMediatorVotes[_mediator] = 0;
        }
    }

    function approveLongBet(uint betId) mediatorOnly public {
        deployedBets[betId].mediatorApproves(payable(msg.sender));
    }

    function ruleLongBetTrue(uint betId) mediatorOnly public {
        deployedBets[betId].mediatorRulesPredictionTrue(payable(msg.sender));
    }

    function ruleLongBetFalse(uint betId) mediatorOnly public {
        deployedBets[betId].mediatorRulesPredictionFalse(payable(msg.sender));
    }
}
