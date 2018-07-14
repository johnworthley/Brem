pragma solidity ^0.4.24;

import "./BREMFactory.sol";
import "./Userable.sol";
import "./BRMToken.sol";

contract BREM is Userable, BREMFactory {

    BRMToken public BRM;

    constructor(BRMToken _brm) public {
        BRM = _brm;
    }    
    
    function createBREMICO( 
        string _name, 
        string _symbol,
        uint256 _rate,
        uint256 _cap,
        uint256 _openingTime,
        uint256 _closingTime,
        string _description,
        bytes32[] _docHashes
    ) 
    public
    onlyDeveloper
    returns (address tokenAddress, address icoAddress)
    {
        require(BRM.balanceOf(msg.sender) >= icoCreationPrice);
        require(bytes(_name).length > 0 && bytes(_symbol).length > 0);
        require(indexes[_name] == 0);
        
        require(BRM.approve(msg.sender, icoCreationPrice));
        require(BRM.transferFrom(msg.sender, address(this), icoCreationPrice));
        
        
        BREMToken token = new BREMToken(_name, _symbol);
        tokenAddress = address(token);
        
        BREMICO ico = new BREMICO(
            _cap, 
            _rate,
            msg.sender,
            token,
            _openingTime,
            _closingTime,
            _description, 
            _docHashes, 
            this
        );
        icoAddress = address(ico);
        
        projects.push(icoAddress);
        indexes[_name] = projects.length - 1;
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
    
}