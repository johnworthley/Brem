pragma solidity ^0.4.24;

import "./BREMToken.sol";
import "./BREMICO.sol";

contract BREMFactory {
    
    function createBREMICO(
        string _name, 
        string _symbol,
        uint256 _rate
    ) 
    public
    returns (address tokenAddress, address icoAddress)
    {
        // TODO: Checking for name and _symbol
        
        BREMToken token = new BREMToken(_name, _symbol);
        tokenAddress = address(token);
        
        BREMICO ico = new BREMICO(_rate, msg.sender, token);
        icoAddress = address(ico);
    }
}