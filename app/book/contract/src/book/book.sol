// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

/*
    - Don't accept a place bet with the date already expired
    - Validate the fee is not greater than the bet amount
*/

contract Book {

    // These constants define the different states a bet can exist in.
    uint8 private constant STATE_NOTEXISTS  = 0;
    uint8 private constant STATE_LIVE       = 1;
    uint8 private constant STATE_RECONCILED = 2;
    uint8 private constant STATE_CANCELLED  = 3;

    // =========================================================================

    // BetInfo represents the details about a bet.
    struct BetInfo {
        uint8     State;
        address[] Participants;
        address   Moderator;
        uint256   AmountBetWei;
        uint256   Expiration;
    }

    // Bet is used to manage bet logic.
    struct Bet {
        BetInfo                   Info;
        mapping (address => bool) IsParticipant;
    }

    // Account represents account information for an account.
    struct Account {
        bool    Exists;
        uint256 Balance;
        uint    Nonce;
    }

    // =========================================================================

    // Owner represents the address who deployed the contract.
    address public Owner;

    // accounts represents the acount information for all participants,
    // moderators, and the Owner.
    mapping (address => Account) private accounts;

    // bets represents current bets, organized by Bet ID.
    mapping (string => Bet) private bets;

    // EventLog provides support for external logging.
    event EventLog(string value);

    // =========================================================================

    // constructor is called when the contract is deployed.
    constructor() {
        Owner = msg.sender;
    }

    // =========================================================================
    // Owner Called API's

    // onlyOwner can be used to restrict access to a function for only the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // Drain will drain the full value of the contract to the contract owner.
    function Drain() onlyOwner payable public {
        address payable account = payable(msg.sender);
        uint256 bal = address(this).balance;

        account.transfer(bal);
        emit EventLog(string.concat("drain[", Error.Addrtoa(account), "] amount[", Error.Itoa(bal), "]"));
    }

    // AccountBalance returns the specified account's balance and amount bet.
    function AccountBalance(address account) onlyOwner view public returns (uint) {
        return accounts[account].Balance;
    }

    // Nonce will retrieve the current nonce for a given account.
    function Nonce(address account) onlyOwner view public returns (uint) {
        return accounts[account].Nonce;
    }

    // BetDetails returns the details about the specified bet.
    function BetDetails(string memory betID) onlyOwner view public returns (BetInfo memory) {
        if (bets[betID].Info.State == STATE_NOTEXISTS) {
            revert("bet id does not exist");
        }

        return bets[betID].Info;
    }

    // PlaceBet will add a bet to the system that is considered a live bet.
    function PlaceBet(
        string    memory   betID,        // Unique Bet identifier
        uint256            amountBetWei, // Amount each participant is betting
        uint256            amountFeeWei, // Amount each participant is paying in fees
        uint256            expiration,   // Time the bet expires
        address            moderator,    // Address of the moderator
        address[] memory   participants, // List of participant addresses
        uint[]    memory   nonces,       // Nonce used per participant for signing
        bytes[]   calldata signatures    // List of participant signatures
    ) onlyOwner public {

        // Check to see if this bet already exists.
        if (bets[betID].Info.State != STATE_NOTEXISTS) {
            revert("bet id already exists");
        }

        // Calculate the total cost to each participant.
        uint256 totalCostWei = amountBetWei + amountFeeWei;

        // Validate the signatures, balances, nounces.
        for (uint i = 0; i < participants.length; i++) {
            address        participant = participants[i];
            uint           nonce       = nonces[i];
            bytes calldata signature   = signatures[i];

            // Ensure the participant has a sufficient balance for the bet.
            if (accounts[participant].Balance < totalCostWei) {
                revert(string.concat(Error.Addrtoa(participant), " has an insufficient balance"));
            }

            // Ensure the expected nonce for this participant is provided.
            if (accounts[participant].Nonce != nonce) {
                revert(string.concat(Error.Addrtoa(participant), " has an invalid nonce"));
            }

            // Reconstruct the data should have been signed by this participant.
            bytes32 hashData = keccak256(abi.encode(betID, participant, nonce));

            // Retrieve the participant's public address from the signature.
            (address addr, Error.Err memory err) = extractAddress(hashData, signature);
            if (err.isError) {
                revert(err.msg);
            }

            // Ensure the address retrieved from the signature matches the participant.
            if (addr != participant) {
                revert(string.concat(Error.Addrtoa(participant), " address doesn't match signature"));
            }
        }

        // Construct a bet from the provided details.
        bets[betID].Info = BetInfo(
            {
                State:        STATE_LIVE,
                Participants: participants,
                Moderator:    moderator,
                Expiration:   expiration,
                AmountBetWei: amountBetWei
            }
        );

        // Move the funds from the participant's balance into the betting pool.
        for (uint i = 0; i < participants.length; i++) {
            address participant = participants[i];

            accounts[participant].Balance -= totalCostWei;
            accounts[participant].Nonce++;
            accounts[Owner].Balance += amountFeeWei;

            // Mark this participant as part of this bet.
            bets[betID].IsParticipant[participant] = true;
        }        

        // Check if we need to add an account for the moderator.
        if (!accounts[moderator].Exists) {
            accounts[moderator] = Account(true, 0, 0);
        }

        emit EventLog(string.concat(betID, " has been added to the system"));
    }

    // ReconcileBet allows a moderator to reconcile a bet.
    function ReconcileBet(
        string    memory   betID,     // Unique Bet identifier
        uint               nonce,     // Nonce used by moderator for signing
        bytes     calldata signature, // Moderator signature
        address[] memory   winners    // List of winner addresses
    ) onlyOwner public {

        // Capture the bet information.
        Bet storage bet = bets[betID];

        // Ensure the bet is live.
        if (bet.Info.State != STATE_LIVE) {
            revert("bet is not live");
        }

        // Ensure the bet has passed its expiration.
        if (block.timestamp < bet.Info.Expiration) {
            revert(string.concat("bet has not yet expired : block.timestamp[", Error.Itoa(block.timestamp), "] expiration[", Error.Itoa(bet.Info.Expiration), "]"));
        }

        // Ensure the nonce used by the moderator is the expected nonce.
        if (accounts[bet.Info.Moderator].Nonce != nonce) {
            revert("invalid moderator nonce");
        }

        // Reconstruct the data that was signed by the moderator.
        bytes32 hashData = keccak256(abi.encode(betID, bet.Info.Moderator, nonce));

        // Retrieve the moderator's public address from the signature.
        (address mod, Error.Err memory err) = extractAddress(hashData, signature);
        if (err.isError) {
            revert(err.msg);
        }

        // Ensure the moderator on file for the bet is the one that signed to
        // reconcile the bet.
        if (mod != bet.Info.Moderator) {
            revert("invalid moderator signature");
        }

        // Ensure each of the winners exist in the participation list.
        for (uint i = 0; i < winners.length; i++) {
            if (!bet.IsParticipant[winners[i]]) {
                revert("winner address is not a participant");
            }
        }

        // Calculate the total winnings for each winner.
        uint256 totalWinnings   = bet.Info.AmountBetWei * 2;
        uint256 amountPerWinner = totalWinnings / winners.length;

        // Give each of the winners the amount listed in the bet.
        for (uint i = 0; i < winners.length; i++) {
            accounts[winners[i]].Balance += amountPerWinner;
        }

        // Increment the moderator's nonce.
        accounts[bet.Info.Moderator].Nonce++;

        // Change the state of the bet to reconciled and set the amount to zero.
        bet.Info.State        = STATE_RECONCILED;
        bet.Info.AmountBetWei = 0;

        emit EventLog(string.concat(betID, " has been reconciled")); 
    }

    // CancelBetModerator allows the moderator to cancel a bet at any time.
    function CancelBetModerator(
        string  memory   betID,
        uint256          amountFeeWei,
        uint             nonce,
        bytes   calldata signature
    ) onlyOwner public {

        // Capture the bet information.
        Bet storage bet = bets[betID];

        // Ensure the bet is live.
        if (bet.Info.State != STATE_LIVE) {
            revert("bet is not live");
        }

        // Ensure the nonce used by the moderator is the expected nonce.
        if (accounts[bet.Info.Moderator].Nonce != nonce) {
            revert("invalid moderator nonce");
        }

        // Reconstruct the data that was signed by the moderator.
        bytes32 hashData = keccak256(abi.encode(betID, bet.Info.Moderator, nonce));

        // Retrieve the moderator's public address from the signature.
        (address mod, Error.Err memory err) = extractAddress(hashData, signature);
        if (err.isError) {
            revert(err.msg);
        }

        // Ensure the moderator on file for the bet is the one that signed to
        // reconcile the bet.
        if (mod != bet.Info.Moderator) {
            revert("invalid moderator signature");
        }

        // Return the money back to the participants minus the fee.
        uint256 totalAmount = bet.Info.AmountBetWei - amountFeeWei;
        for (uint i = 0; i < bet.Info.Participants.length; i++) {
            accounts[bet.Info.Participants[i]].Balance += totalAmount;
            accounts[Owner].Balance += amountFeeWei;
        }

        // Increment the moderator's nonce.
        accounts[bet.Info.Moderator].Nonce++;

        // Change the state of the bet to cancelled and set the amount to zero.
        bet.Info.State        = STATE_CANCELLED;
        bet.Info.AmountBetWei = 0;

        emit EventLog(string.concat(betID, " has been cancelled by moderator")); 
    }

    // CancelBetParticipants allows all the participants to cancel a bet.
    function CancelBetParticipants(
        string  memory   betID,
        uint256          amountFeeWei,
        uint[]  memory   nonces,
        bytes[] calldata signatures
    ) onlyOwner public {

        // Capture the bet information.
        Bet storage bet = bets[betID];

        // Ensure the bet is live.
        if (bet.Info.State != STATE_LIVE) {
            revert("bet is not live");
        }

        // Ensure we have a proper number of signatures and nonces.
        if ((bet.Info.Participants.length != signatures.length) || (bet.Info.Participants.length != nonces.length)) {
            revert("invalid number of signatures or nonces");
        }

        // Ensure the we have proper signatures from all the participants.
        for (uint i = 0; i < bet.Info.Participants.length; i++) {
            address        participant = bet.Info.Participants[i];
            uint           nonce       = nonces[i];
            bytes calldata signature   = signatures[i];

            // Ensure the nonce used by the participant is the expected nonce.
            if (accounts[participant].Nonce != nonce) {
                revert(string.concat(Error.Addrtoa(participant), "] has an invalid nonce"));
            }

            // Reconstruct the data that was signed by the participant.
            bytes32 hashData = keccak256(abi.encode(betID, participant, nonce));

            // Retrieve the participant's public address from the signature.
            (address addr, Error.Err memory err) = extractAddress(hashData, signature);
            if (err.isError) {
                revert(err.msg);
            }

            // Ensure the participant's signature matches the address of file.
            if (addr != participant) {
                revert(string.concat(Error.Addrtoa(participant), " address doesn't match signature"));
            }

            // Increment the nonce value for this participant.
            accounts[participant].Nonce++;
        }

        // Return the money back to the participants minus the fee.
        uint256 totalAmount = bet.Info.AmountBetWei - amountFeeWei;
        for (uint i = 0; i < bet.Info.Participants.length; i++) {
            accounts[bet.Info.Participants[i]].Balance += totalAmount;
            accounts[Owner].Balance += amountFeeWei;
        }

        // Change the state of the bet to cancelled and set the amount to zero.
        bet.Info.State        = STATE_CANCELLED;
        bet.Info.AmountBetWei = 0;

        emit EventLog(string.concat(betID, " has been cancelled by all participants")); 
    }

    // CancelBetOwner allows the owner to cancel a bet at any time.
    function CancelBetOwner(
        string  memory betID,
        uint256        amountFeeWei
    ) onlyOwner public {

        // Capture the bet information.
        Bet storage bet = bets[betID];

        // Ensure the bet is live.
        if (bet.Info.State != STATE_LIVE) {
            revert("bet is not live");
        }

        // Return the money back to the participants minus the fee.
        uint256 totalAmount = bet.Info.AmountBetWei - amountFeeWei;
        for (uint i = 0; i < bet.Info.Participants.length; i++) {
            accounts[bet.Info.Participants[i]].Balance += totalAmount;
            accounts[Owner].Balance += amountFeeWei;
        }

        // Change the state of the bet to cancelled and set the amount to zero.
        bet.Info.State        = STATE_CANCELLED;
        bet.Info.AmountBetWei = 0;

        emit EventLog(string.concat(betID, " has been cancelled by owner")); 
    }

    // =========================================================================
    // Account Called API's

    // CancelBetExpired allows an individual participant to cancel a bet 30 days
    // after cancelation.
    function CancelBetExpired(
        string memory betID
    ) public {

        // Capture the bet information.
        Bet storage bet = bets[betID];

        // Ensure the bet is live.
        if (bet.Info.State != STATE_LIVE) {
            revert("bet is not live");
        }

        // Ensure the bet has expired.
        if (block.timestamp < bet.Info.Expiration + 30 days) {
            revert("bets may only be canceled 30+ days after expiration");
        }

        // Ensure the canceler is one of the participants.
        if (!bet.IsParticipant[msg.sender]) {
            revert("canceler is not authorized to cancel this bet");
        }

        // Return the money back to the participants minus the fee.
        for (uint i = 0; i < bet.Info.Participants.length; i++) {
            accounts[bet.Info.Participants[i]].Balance += bet.Info.AmountBetWei;
        }

        // Change the state of the bet to cancelled and set the amount to zero.
        bet.Info.State        = STATE_CANCELLED;
        bet.Info.AmountBetWei = 0;

        emit EventLog(string.concat(betID, " has been cancelled by ", Error.Addrtoa(msg.sender), " since expired")); 
    }

    // Balance returns the balance of the caller.
    function Balance() view public returns (uint) {
        return accounts[msg.sender].Balance;
    }

    // Deposit the given amount to the account balance.
    function Deposit() payable public {
        accounts[msg.sender].Balance += msg.value;

        emit EventLog(string.concat("deposit[", Error.Addrtoa(msg.sender), "] balance[", Error.Itoa(accounts[msg.sender].Balance), "]"));
    }

    // Withdraw all of the available balance from the account.
    function Withdraw() payable public {
        address payable account = payable(msg.sender);

        uint bal = accounts[msg.sender].Balance;
        if (bal == 0) {
            revert("not enough balance");
        }

        account.transfer(bal);
        accounts[msg.sender].Balance -= bal;

        emit EventLog(string.concat("withdraw[", Error.Addrtoa(msg.sender), "] amount[", Error.Itoa(bal), "]"));
    }

    // =========================================================================
    // Private Functions

    // extractAddress expects the raw data that was signed and will apply the Ethereum
    // salt value manually. This hides the underlying implementation of the salt.
    function extractAddress(bytes32 hashData, bytes calldata sig) private pure returns (address, Error.Err memory) {
        if (sig.length != 65) {
            return (address(0), Error.New("invalid signature length"));
        }

        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 saltedData = keccak256(abi.encodePacked(prefix, hashData));

        bytes32 r = bytes32(sig[:32]);
        bytes32 s = bytes32(sig[32:64]);
        uint8 v = uint8(sig[64]);

        return (ecrecover(saltedData, v, r, s), Error.None());
    }
}
