pragma solidity ^0.4.24;

/**
 * @title Roles
 * @author Francisco Giordano (@frangio)
 * @dev Library for managing addresses assigned to a Role.
 * See RBAC.sol for example usage.
 */
library Roles {
    struct Role {
        mapping (address => bool) bearer;
    }

    /**
    * @dev give an address access to this role
    */
    function add(Role storage role, address addr)
        internal
    {
        role.bearer[addr] = true;
    }

    /**
    * @dev remove an address' access to this role
    */
    function remove(Role storage role, address addr)
        internal
    {
        role.bearer[addr] = false;
    }

    /**
    * @dev check if an address has this role
    * // reverts
    */
    function check(Role storage role, address addr)
        view
        internal
    {
        require(has(role, addr));
    }

    /**
    * @dev check if an address has this role
    * @return bool
    */
    function has(Role storage role, address addr)
        view
        internal
    returns (bool)
    {
        return role.bearer[addr];
    }
}

/**
 * @title RBAC (Role-Based Access Control)
 * @author Matt Condon (@Shrugs)
 * @dev Stores and provides setters and getters for roles and addresses.
 * Supports unlimited numbers of roles and addresses.
 * See //contracts/mocks/RBACMock.sol for an example of usage.
 * This RBAC method uses strings to key roles. It may be beneficial
 * for you to write your own implementation of this interface using Enums or similar.
 * It's also recommended that you define constants in the contract, like ROLE_ADMIN below,
 * to avoid typos.
 */
contract RBAC {
    using Roles for Roles.Role;

    mapping (string => Roles.Role) private roles;

    event RoleAdded(address indexed operator, string role);
    event RoleRemoved(address indexed operator, string role);

    /**
    * @dev reverts if addr does not have role
    * @param _operator address
    * @param _role the name of the role
    * // reverts
    */
    function checkRole(address _operator, string _role)
        view
        public
    {
        roles[_role].check(_operator);
    }

    /**
    * @dev determine if addr has role
    * @param _operator address
    * @param _role the name of the role
    * @return bool
    */
    function hasRole(address _operator, string _role)
        view
        public
    returns (bool)
    {
        return roles[_role].has(_operator);
    }

    /**
    * @dev add a role to an address
    * @param _operator address
    * @param _role the name of the role
    */
    function addRole(address _operator, string _role)
        internal
    {
        roles[_role].add(_operator);
        emit RoleAdded(_operator, _role);
    }

    /**
    * @dev remove a role from an address
    * @param _operator address
    * @param _role the name of the role
    */
    function removeRole(address _operator, string _role)
        internal
    {
        roles[_role].remove(_operator);
        emit RoleRemoved(_operator, _role);
    }

    /**
    * @dev modifier to scope access to a single role (uses msg.sender as addr)
    * @param _role the name of the role
    * // reverts
    */
    modifier onlyRole(string _role)
    {
        checkRole(msg.sender, _role);
        _;
    }
}

/**
 * @title Superuser
 * @dev The Superuser contract defines a single superuser who can transfer the ownership 
 * of a contract to a new address, even if he is not the owner. 
 * A superuser can transfer his role to a new address. 
 */
contract Superuserable is RBAC {
    
    string public constant ROLE_SUPERUSER = "superuser";

    constructor () public {
        addRole(msg.sender, ROLE_SUPERUSER);
    }

    /**
    * @dev Throws if called by any account that's not a superuser.
    */
    modifier onlySuperuser() {
        checkRole(msg.sender, ROLE_SUPERUSER);
        _;
    }

    /**
    * @dev getter to determine if address has superuser role
    */
    function isSuperuser(address _addr)
        public
        view
    returns (bool)
    {
        return hasRole(_addr, ROLE_SUPERUSER);
    }

    /**
    * @dev Allows the current superuser to transfer his role to a newSuperuser.
    * @param _newSuperuser The address to transfer ownership to.
    */
    function transferSuperuser(address _newSuperuser) public onlySuperuser {
        require(_newSuperuser != address(0));
        removeRole(msg.sender, ROLE_SUPERUSER);
        addRole(_newSuperuser, ROLE_SUPERUSER);
    }
}

contract Auditable is Superuserable {
    
    string public constant ROLE_AUDITOR = "auditor";
    
    function addAuditor(address _newAuditor) public onlySuperuser {
        require(_newAuditor != address(0));
        addRole(_newAuditor, ROLE_AUDITOR);
    }
    
    function removeAuditor(address _auditor) public onlySuperuser {
        removeRole(_auditor, ROLE_AUDITOR);
    }
    
    modifier onlyAuditor() {
        checkRole(msg.sender, ROLE_AUDITOR);
        _;
    }
    
    function isAuditor(address _addr)
        public
        view
    returns (bool) 
    {
        return hasRole(_addr, ROLE_AUDITOR);
    }
}

contract Developerable is Auditable {
    
    string public constant ROLE_DEVELOPER = "developer";
    
    function addDeveloper() public {
        require(msg.sender != address(0));
        addRole(msg.sender, ROLE_DEVELOPER);
    }
    
    modifier onlyDeveloper() {
        checkRole(msg.sender, ROLE_DEVELOPER);
        _;
    }
    
    function isDeveloper(address _addr) 
        public
        view
    returns (bool)
    {
        return hasRole(_addr, ROLE_DEVELOPER);
    }
}

contract Userable is Developerable {
    
    struct User {
        string name;
    }
    
    mapping(address => User) users;
    
    modifier onlyValidName(string _name) {
        require(bytes(_name).length > 0);
        _;
    }
    
    modifier onlyExistingUser() {
        require(bytes(users[msg.sender].name).length > 0);
        _;
    }
    
    function signUp(string _name) 
        public 
        onlyValidName(_name)
    returns (string)
    {
        if (bytes(users[msg.sender].name).length == 0) {
            users[msg.sender].name = _name;
        }
        
        if (!isSuperuser(msg.sender) && !isAuditor(msg.sender)) {
            addDeveloper();
        }
        
        return users[msg.sender].name;
    }
    
    function login()
        view
        public
        onlyExistingUser
    returns (string)
    {
        return users[msg.sender].name;
    }
}