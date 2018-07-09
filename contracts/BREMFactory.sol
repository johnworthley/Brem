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
    
    function createBREMICO( 
        string _name, 
        string _symbol,
        uint256 _rate,
        uint256 _totalStages,
        string _description,
        bytes32[] _docHashes
    ) 
    public
    returns (address tokenAddress, address icoAddress)
    {
        require(bytes(_name).length > 0 && bytes(_symbol).length > 0);
        require(indexes[_name] == 0);
        
        BREMToken token = new BREMToken(_name, _symbol);
        tokenAddress = address(token);
        
        BREMICO ico = new BREMICO(_rate, msg.sender, token, _totalStages,
            _description, _docHashes);
        icoAddress = address(ico);
        projects.push(icoAddress);
        indexes[_name] = projects.length - 1;
        emit BREMICOCreated(msg.sender, icoAddress, tokenAddress, _name);
    }
}