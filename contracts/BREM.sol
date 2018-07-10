pragma solidity ^0.4.24;

import "./BREMFactory.sol";
import "./Userable.sol";

contract BREM is Userable, BREMFactory {
    
    function createBREMICO( 
        string _name, 
        string _symbol,
        uint256 _rate,
        uint256 _cap,
        uint256 _totalStages,
        string _description,
        bytes32[] _docHashes
    ) 
    public
    onlyDeveloper
    returns (address tokenAddress, address icoAddress)
    {
        require(bytes(_name).length > 0 && bytes(_symbol).length > 0);
        require(indexes[_name] == 0);
        
        BREMToken token = new BREMToken(_name, _symbol);
        tokenAddress = address(token);
        
        BREMICO ico = new BREMICO(_cap, _rate, msg.sender, token, _totalStages,
            _description, _docHashes, this);
        icoAddress = address(ico);
        projects.push(icoAddress);
        indexes[_name] = projects.length - 1;
        emit BREMICOCreated(msg.sender, icoAddress, tokenAddress, _name);
    }
    
}