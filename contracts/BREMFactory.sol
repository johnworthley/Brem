pragma solidity ^0.4.24;

import "./BREMToken.sol";
import "./BREMICO.sol";

contract BREMFactory {
    
    address[] public projects;
    mapping(string => uint256) indexes;
    
    constructor() public {
        projects.push(address(0));
    }
    
    event BREMICOCreated(
        address indexed creator,
        address indexed icoAddress,
        address indexed tokenAddress,
        string name
    );
}