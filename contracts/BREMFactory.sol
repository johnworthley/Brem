pragma solidity ^0.4.24;

import "./BREMToken.sol";
import "./BREMICO.sol";

contract BREMFactory is Userable{
    
    uint256 public projectsAmount;
    
    mapping(uint256 => address) projects;
    
    function getProject(uint256 _index) 
        public 
        view 
        returns(address)
    {
        require(_index < projectsAmount);
        return projects[_index];
    }
    
    mapping(string => uint256) indexes;
    
    function getProjectByName(string _projectName)
        public
        view
        returns(address) 
    {
        require(bytes(_projectName).length > 0);
        uint256 index = indexes[_projectName];
        require(projects[index] != address(0));
        
        return projects[index];    
    }
    
    event BREMICOCreated(
        address indexed creator,
        address indexed icoAddress,
        address indexed tokenAddress,
        string name
    );
    
    uint256 public icoCreationPrice;

    function setIcoCreationPrice(uint256 _icoCreationPrice)
        public 
        onlySuperuser
    {
        require(_icoCreationPrice != icoCreationPrice);
        icoCreationPrice = _icoCreationPrice;
    }

    uint256 public withdrawFeePercent;

    function setWithdrawFeePercent(uint256 _withdrawFeePercent)
        public
        onlySuperuser
    {
        require(_withdrawFeePercent != withdrawFeePercent);
        require(_withdrawFeePercent >= 0 && _withdrawFeePercent <=100);

        withdrawFeePercent = _withdrawFeePercent;
    }
}