pragma solidity ^0.4.24;

import "./BREMFactory.sol";
import "./Userable.sol";
import "./BRMToken.sol";

contract BREM is Userable, BREMFactory, BRMToken {
    
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
        require(balanceOf(msg.sender) >= icoCreationPrice);
        require(bytes(_name).length > 0 && bytes(_symbol).length > 0);
        require(indexes[_name] == 0);
        
        require(approve(msg.sender, icoCreationPrice));
        require(transferFrom(msg.sender, address(this), icoCreationPrice));
        
        
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
    
}