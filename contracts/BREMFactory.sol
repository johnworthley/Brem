pragma solidity ^0.4.24;

import "./BREMToken.sol";
import "./BREMICO.sol";

contract BREMFactory {
    
    address[] public projects;
    mapping(string => uint256) indexes;
    
    constructor() public {
        projects.push(address(0));
    }
    
    function createBREMICO(
        string _name, 
        string _symbol,
        uint256 _rate
    ) 
    public
    returns (address tokenAddress, address icoAddress)
    {
        require(bytes(_name).length > 0 && bytes(_symbol).length > 0);
        require(indexes[_name] == 0);
        
        BREMToken token = new BREMToken(_name, _symbol);
        tokenAddress = address(token);
        
        BREMICO ico = new BREMICO(_rate, msg.sender, token);
        icoAddress = address(ico);
        projects.push(icoAddress);
        indexes[_name] = projects.length - 1;
    }
}