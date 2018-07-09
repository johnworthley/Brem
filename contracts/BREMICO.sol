pragma solidity ^0.4.24;

import "./BREMToken.sol";
import "./Userable.sol";

/**
 * @title SafeERC20
 * @dev Wrappers around ERC20 operations that throw on failure.
 * To use this library you can add a `using SafeERC20 for ERC20;` statement to your contract,
 * which allows you to call the safe operations as `token.safeTransfer(...)`, etc.
 */
library SafeERC20 {
    function safeTransfer(ERC20Basic token, address to, uint256 value) internal {
        require(token.transfer(to, value));
    }
    
    function safeTransferFrom(
        ERC20 token,
        address from,
        address to,
        uint256 value
    )
    internal
    {
        require(token.transferFrom(from, to, value));
    }
    
    function safeApprove(ERC20 token, address spender, uint256 value) internal {
        require(token.approve(spender, value));
    }
}

contract BREMICO {
    using SafeMath for uint256;
    using SafeERC20 for BREMToken;

    // The token being sold
    BREMToken public token;

    // Address where funds are collected
    address public wallet;

    // How many token units a buyer gets per wei.
    // The rate is the conversion between wei and the smallest and indivisible token unit.
    // So, if you are using a rate of 1 with a DetailedERC20 token with 3 decimals called TOK
    // 1 wei will give you 1 unit, or 0.001 TOK.
    uint256 public rate;

    // Amount of wei raised
    uint256 public weiRaised;
    
    // ICO description
    string public description;
    
    // ICO documentation hashes
    bytes32[] public docHashes;
    
    // Structure represents ICO stage
    struct Stage {
        uint256 level;
        bool forwarded;
        mapping(address => bool) verified;
        uint256 verificationAmount;
    }
    
    // ICO stages amount
    uint256 totalStages;
    
    // ICO current stage
    uint256 currentStage;
    
    // Auditors contract instance
    Auditable audit;
    
    // All stages
    mapping(uint256 => Stage) stages;

    /**
    * Event for token purchase logging
    * @param purchaser who paid for the tokens
    * @param beneficiary who got the tokens
    * @param value weis paid for purchase
    * @param amount amount of tokens purchased
    */
    event TokenPurchase(
        address indexed purchaser,
        address indexed beneficiary,
        uint256 value,
        uint256 amount
    );

    /**
    * @param _rate Number of token units a buyer gets per wei
    * @param _wallet Address where collected funds will be forwarded to
    * @param _token Address of the token being sold
    */
    constructor(
        uint256 _rate, 
        address _wallet, 
        BREMToken _token,
        uint256 _totalStages, 
        string _description, 
        bytes32[] _docHashes,
        Auditable _auditAddress
    ) 
        public 
    {
        require(_rate > 0);
        require(_wallet != address(0));
        require(_token != address(0));
        require(_auditAddress != address(0));
        
        rate = _rate;
        wallet = _wallet;
        token = _token;
        totalStages = _totalStages;
        description = _description;
        docHashes = _docHashes;
        audit = Auditable(_auditAddress);
        
        // Fill templates for each stage
        for (uint256 i = 0; i < _totalStages; i++) {
            stages[i].level = i;
        }
    }

    // -----------------------------------------
    // Crowdsale external interface
    // -----------------------------------------
    
    /**
    * @dev fallback function ***DO NOT OVERRIDE***
    */
    function () external payable {
        buyTokens(msg.sender);
    }

    /**
    * @dev low level token purchase ***DO NOT OVERRIDE***
    * @param _beneficiary Address performing the token purchase
    */
    function buyTokens(address _beneficiary) public payable {
    
        uint256 weiAmount = msg.value;
        _preValidatePurchase(_beneficiary, weiAmount);
    
        // calculate token amount to be created
        uint256 tokens = _getTokenAmount(weiAmount);
    
        // update state
        weiRaised = weiRaised.add(weiAmount);
    
        _processPurchase(_beneficiary, tokens);
        emit TokenPurchase(
            msg.sender,
            _beneficiary,
            weiAmount,
            tokens
        );
    
        // _updatePurchasingState(_beneficiary, weiAmount);
    
        // _forwardFunds();
        // _postValidatePurchase(_beneficiary, weiAmount);
    }

  // -----------------------------------------
  // Internal interface (extensible)
  // -----------------------------------------

    /**
    * @dev Validation of an incoming purchase. Use require statements to revert state when conditions are not met. Use super to concatenate validations.
    * @param _beneficiary Address performing the token purchase
    * @param _weiAmount Value in wei involved in the purchase
    */
    function _preValidatePurchase(
        address _beneficiary,
        uint256 _weiAmount
    )
        internal pure
    {
        require(_beneficiary != address(0));
        require(_weiAmount != 0);
    }

//   /**
//   * @dev Validation of an executed purchase. Observe state and use revert statements to undo rollback when valid conditions are not met.
//   * @param _beneficiary Address performing the token purchase
//   * @param _weiAmount Value in wei involved in the purchase
//   */
//   function _postValidatePurchase(
//     address _beneficiary,
//     uint256 _weiAmount
//   )
//     internal
//   {
//     // optional override
//   }

    /**
    * @dev Source of tokens. Override this method to modify the way in which the crowdsale ultimately gets and sends its tokens.
    * @param _beneficiary Address performing the token purchase
    * @param _tokenAmount Number of tokens to be emitted
    */
    function _deliverTokens(
        address _beneficiary,
        uint256 _tokenAmount
    )
        internal
    {
        token.safeTransfer(_beneficiary, _tokenAmount);
    }

    /**
    * @dev Executed when a purchase has been validated and is ready to be executed. Not necessarily emits/sends tokens.
    * @param _beneficiary Address receiving the tokens
    * @param _tokenAmount Number of tokens to be purchased
    */
    function _processPurchase(
        address _beneficiary,
        uint256 _tokenAmount
    )
        internal
    {
        _deliverTokens(_beneficiary, _tokenAmount);
    }

//   /**
//   * @dev Override for extensions that require an internal state to check for validity (current user contributions, etc.)
//   * @param _beneficiary Address receiving the tokens
//   * @param _weiAmount Value in wei involved in the purchase
//   */
//   function _updatePurchasingState(
//     address _beneficiary,
//     uint256 _weiAmount
//   )
//     internal
//   {
//     // optional override
//   }

    /**
    * @dev Override to extend the way in which ether is converted to tokens.
    * @param _weiAmount Value in wei to be converted into tokens
    * @return Number of tokens that can be purchased with the specified _weiAmount
    */
    function _getTokenAmount(uint256 _weiAmount)
        internal view returns (uint256)
    {
        return _weiAmount.mul(rate);
    }

//   /**
//   * @dev Determines how ETH is stored/forwarded on purchases.
//   */
//   function _forwardFunds() internal {
//     wallet.transfer(msg.value);
//   }
}