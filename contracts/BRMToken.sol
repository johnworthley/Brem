pragma solidity ^0.4.24;

import "./BREMToken.sol";

contract BRMToken is MintableToken {
    
    string public constant name = "BRM";
    
    string public constant symbol = "BRM";
    
    uint8 public constant decimals = 18;
    
}