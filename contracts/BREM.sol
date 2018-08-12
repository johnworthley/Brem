pragma solidity ^0.4.24;

import "./BREMFactory.sol";
import "./Userable.sol";
// import "./BRMToken.sol";

contract BREM is BREMFactory {

    // BRMToken public BRM;

    constructor(
        // address _brmAddress,
        // uint256 _icoCreationPrice,
        uint256 _withdrawFeePercent
    ) 
    public 
    {
        require(_withdrawFeePercent >= 0 && _withdrawFeePercent <=100);
        // BRM = BRMToken(_brmAddress);
        // icoCreationPrice = _icoCreationPrice;
        withdrawFeePercent = _withdrawFeePercent;
    }    
    
    function createBREMICO( 
        string _name, 
        string _symbol,
        uint256 _rate,
        uint256 _cap,
        uint256 _closingTime,
        string _docHash
    ) 
    public
    onlyDeveloper
    returns (address tokenAddress, address icoAddress)
    {
        // require(BRM.balanceOf(msg.sender) >= icoCreationPrice);
        require(bytes(_name).length > 0 && bytes(_symbol).length > 0);
        require(indexes[_name] == 0);
        require(_cap >= 100);
        
        // require(BRM.transferFrom(msg.sender, address(this), icoCreationPrice));
        
        
        BREMICO ico = new BREMICO(
            address(this),
            _name,
            _symbol,
            _cap, 
            _rate,
            msg.sender,
            _closingTime,
            _docHash, 
            withdrawFeePercent
        );
        icoAddress = address(ico);
        tokenAddress = ico.token();

        projects[projectsAmount] = icoAddress;
        indexes[_name] = projectsAmount;
        projectsAmount++;
        emit BREMICOCreated(msg.sender, icoAddress, tokenAddress, _name);
    }
    
    // Withdraw collected fies
    function withdrawFees(uint256 _value) 
        public 
        onlySuperuser
     {
        require(address(this).balance >= _value && _value > 0);
        msg.sender.transfer(_value);
    } 

    function () external payable {}    
}