pragma solidity ^0.4.24;

import "./BREMToken.sol";
import "./BREMICO.sol";

contract BREMFactory {
    
    uint256 public projectsAmount;
    
    mapping(uint256 => address) projects;
    
    function getProject(uint256 index) 
        public 
        view 
        returns(address)
    {
        require(index < projectsAmount);
        return projects[index];
    }
    
    mapping(string => uint256) indexes;
    
    function getProjectByName(string projectName)
        public
        view
        returns(address) 
    {
        // TODO: Finish
        return 0;    
    }
    
    event BREMICOCreated(
        address indexed creator,
        address indexed icoAddress,
        address indexed tokenAddress,
        string name
    );
    
    uint256 public icoCreationPrice = 10000;
}