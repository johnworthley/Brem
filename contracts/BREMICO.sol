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
    
    // Value in wei that need owner
    uint256  public cap;
    
    // ICO documentation hashes
    string public docHash;
    
    // Auditors contract instance
    Auditable audit;
    
    // Current ICO auditors
    mapping(address => bool) auditors;
    
    mapping(uint256 => address) auditorsList;
    
    // Current ICO auditors amount
    uint256 public auditorsAmount;
    
    // Is superuser selected all current ICO auditors
    bool public auditSelected;
    
    // Developer's withdraw request
    struct WithdrawRequst {
        uint256 value;
        uint256 confirmAmount;
        mapping(address => bool) confirmed;
    }
    
    WithdrawRequst public request;
    
    // ICO closing time
    uint256 public closingTime;
    
    // Token sale payment registry
    mapping(address => uint256) public balances;
    
    // Withdraw fee percent
    uint256 public withdrawFeePercent;

    // Brem contract address
    address public brem;
    
    modifier onlyWhileOpen {
        // solium-disable-next-line security/no-block-members
        require(auditSelected && block.timestamp <= closingTime);
        _;
    }

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
    */
    constructor(
        address _brem,
        string _name,
        string _symbol,
        uint256 _cap,
        uint256 _rate, 
        address _wallet, 
        uint256 _closingTime,
        string _docHash,
        Auditable _auditAddress,
        uint256 _withdrawFeePercent
    ) 
        public 
    {
        require(_brem != address(0));
        require(bytes(_name).length > 0);
        require(_cap > 0);
        require(_rate > 0);
        require(_wallet != address(0));
        require(_auditAddress != address(0));
        require(_closingTime >= block.timestamp);
        require(_withdrawFeePercent >= 0 && _withdrawFeePercent <= 100);
        
        brem = _brem;
        cap = _cap;
        rate = _rate;
        wallet = _wallet;
        closingTime = _closingTime;
        docHash = _docHash;
        audit = Auditable(_auditAddress);
        withdrawFeePercent = _withdrawFeePercent;

        BREMToken _token = new BREMToken(_name, _symbol, this);
        require(address(_token) != address(0));
        token = _token;
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
    function buyTokens(address _beneficiary) 
        public 
        payable 
        onlyWhileOpen
    {
    
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
    
        _updatePurchasingState(_beneficiary, weiAmount);
    
        // _postValidatePurchase();
    }
    
    function addAuditor(address _auditor) public {
        require(block.timestamp <= closingTime);
        require(!auditSelected);
        require(audit.isSuperuser(msg.sender));
        require(audit.isAuditor(_auditor));
        require(!auditors[_auditor]);
        
        auditors[_auditor] = true;
        auditorsList[auditorsAmount] = _auditor;
        auditorsAmount++;
    }
    
    function finishAuditorSelection() public {
        require(block.timestamp <= closingTime);
        require(!auditSelected);
        require(audit.isSuperuser(msg.sender));
        require(auditorsAmount > 0);
        
        auditSelected = true;
    }
    
    function withdraw(uint256 _value) public payable {
        require(msg.sender == wallet);
        require(auditSelected);
        require(capReached() && hasClosed());
        require(_value >= 100 && _value <= address(this).balance);
        require(request.value == 0);
        
        if (address(this).balance - _value < 100) {
            request = WithdrawRequst(address(this).balance, 0);
        } else {
            request = WithdrawRequst(_value, 0);
        }
    }
    
    // Auditors confirm withdraw
    function confirmWithdraw() public {
        require(audit.isAuditor(msg.sender));
        require(request.value > 0);
        require(!request.confirmed[msg.sender]);
        require(auditSelected);
        require(capReached() && hasClosed());
        require(auditors[msg.sender]);
        
        request.confirmed[msg.sender] = true;
        request.confirmAmount++;
        
        if (request.confirmAmount == auditorsAmount) {
            uint256 _value = request.value;
            request = WithdrawRequst(0, 0);
            for (uint256 i = 0; i < auditorsAmount; i++) {
                request.confirmed[auditorsList[i]] = false;
            }
            uint256 _feeValue = _value.div(100).mul(withdrawFeePercent);
            uint256 _trasferValue = _value.sub(_feeValue);
            brem.transfer(_feeValue);
            wallet.transfer(_trasferValue);
        }
    }
    
    function refund() public {
        require(hasClosed() && !capReached());
        require(balances[msg.sender] > 0);
        
        uint256 _value = balances[msg.sender];
        balances[msg.sender] = 0;
        uint256 _tokenAmount = _getTokenAmount(_value);
        token.burnForRefund(msg.sender, _tokenAmount);
        weiRaised = weiRaised.sub(_value);
        msg.sender.transfer(_value);
    }
    
    /**
    * @dev Checks whether the period in which the crowdsale is open has already elapsed.
    * @return Whether crowdsale period has elapsed
    */
    function hasClosed() public view returns (bool) {
        // solium-disable-next-line security/no-block-members
        return block.timestamp > closingTime;
    }
    
    /**
    * @dev Checks whether the cap has been reached.
    * @return Whether the cap was reached
    */
    function capReached() public view returns (bool) {
        return weiRaised >= cap;
    }
    
    function getAuditor(uint256 _index) public view returns (address) {
        require(_index < auditorsAmount);
        return auditorsList[_index];
    }

    function isAuditor(address _auditor) public view returns (bool) {
        return auditors[_auditor];
    }
    
    function isRequested() public view returns (bool) {
        return capReached() && hasClosed() && 
            request.value > 0 && !isWithdrawn();
    }
    
    function isConfirmed(address _auditor) public view returns (bool) {
        return request.confirmed[_auditor];
    }
    
    function isWithdrawn() public view returns (bool) {
        return hasClosed() && capReached() && address(this).balance < 100;
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
        token.mint(_beneficiary, _tokenAmount);
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

    /**
    * @dev Override for extensions that require an internal state to check for validity (current user contributions, etc.)
    * @param _beneficiary Address receiving the tokens
    * @param _weiAmount Value in wei involved in the purchase
    */
    function _updatePurchasingState(
        address _beneficiary,
        uint256 _weiAmount
    )
        internal
    {
        balances[_beneficiary] += _weiAmount;
    }

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
}