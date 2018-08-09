// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AuditableABI is the input ABI used to generate the binding from.
const AuditableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_AUDITOR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAuditor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSuperuser\",\"type\":\"address\"}],\"name\":\"transferSuperuser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isSuperuser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAuditor\",\"type\":\"address\"}],\"name\":\"addAuditor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_SUPERUSER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// AuditableBin is the compiled bytecode used for deploying new contracts.
const AuditableBin = `0x60c0604052600960809081527f737570657275736572000000000000000000000000000000000000000000000060a052610043903390640100000000610048810204565b61018b565b6100bf826000836040518082805190602001908083835b6020831061007e5780518252601f19909201916020918201910161005f565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505064010000000061016681026107a51704565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610128578181015183820152602001610110565b50505050905090810190601f1680156101555780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6108168061019a6000396000f30060806040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c8114610092578063217fe6c6146100fb5780633f029aaf1461017657806349b905571461020057806357c393fa14610221578063bceee05e14610242578063e429cef114610263578063ebb4f48414610284575b600080fd5b34801561009e57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100f9958335600160a060020a03169536956044949193909101919081908401838280828437509497506102999650505050505050565b005b34801561010757600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610162958335600160a060020a03169536956044949193909101919081908401838280828437509497506103079650505050505050565b604080519115158252519081900360200190f35b34801561018257600080fd5b5061018b61037a565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101c55781810151838201526020016101ad565b50505050905090810190601f1680156101f25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020c57600080fd5b50610162600160a060020a03600435166103b1565b34801561022d57600080fd5b506100f9600160a060020a03600435166103f8565b34801561024e57600080fd5b50610162600160a060020a0360043516610497565b34801561026f57600080fd5b506100f9600160a060020a03600435166104c6565b34801561029057600080fd5b5061018b610547565b610303826000836040518082805190602001908083835b602083106102cf5780518252601f1990920191602091820191016102b0565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505061056c565b5050565b6000610373836000846040518082805190602001908083835b6020831061033f5780518252601f199092019160209182019101610320565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610581565b9392505050565b60408051808201909152600781527f61756469746f7200000000000000000000000000000000000000000000000000602082015281565b60006103f2826040805190810160405280600781526020017f61756469746f7200000000000000000000000000000000000000000000000000815250610307565b92915050565b610425336040805190810160405280600981526020016000805160206107cb833981519152815250610299565b600160a060020a038116151561043a57600080fd5b610467336040805190810160405280600981526020016000805160206107cb8339815191528152506105a0565b610494816040805190810160405280600981526020016000805160206107cb8339815191528152506106b1565b50565b60006103f2826040805190810160405280600981526020016000805160206107cb833981519152815250610307565b6104f3336040805190810160405280600981526020016000805160206107cb833981519152815250610299565b600160a060020a038116151561050857600080fd5b610494816040805190810160405280600781526020017f61756469746f72000000000000000000000000000000000000000000000000008152506106b1565b60408051808201909152600981526000805160206107cb833981519152602082015281565b6105768282610581565b151561030357600080fd5b600160a060020a03166000908152602091909152604090205460ff1690565b61060a826000836040518082805190602001908083835b602083106105d65780518252601f1990920191602091820191016105b7565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610783565b81600160a060020a03167fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a826040518080602001828103825283818151815260200191508051906020019080838360005b8381101561067357818101518382015260200161065b565b50505050905090810190601f1680156106a05780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b61071b826000836040518082805190602001908083835b602083106106e75780518252601f1990920191602091820191016106c8565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929150506107a5565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360008381101561067357818101518382015260200161065b565b600160a060020a0316600090815260209190915260409020805460ff19169055565b600160a060020a0316600090815260209190915260409020805460ff1916600117905556007375706572757365720000000000000000000000000000000000000000000000a165627a7a723058203852b5785612141d1e5038c929a59e18525b4571a3d0903aa3a1afbc2ab31ed70029`

// DeployAuditable deploys a new Ethereum contract, binding an instance of Auditable to it.
func DeployAuditable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Auditable, error) {
	parsed, err := abi.JSON(strings.NewReader(AuditableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuditableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auditable{AuditableCaller: AuditableCaller{contract: contract}, AuditableTransactor: AuditableTransactor{contract: contract}, AuditableFilterer: AuditableFilterer{contract: contract}}, nil
}

// Auditable is an auto generated Go binding around an Ethereum contract.
type Auditable struct {
	AuditableCaller     // Read-only binding to the contract
	AuditableTransactor // Write-only binding to the contract
	AuditableFilterer   // Log filterer for contract events
}

// AuditableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuditableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuditableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuditableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuditableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuditableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuditableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuditableSession struct {
	Contract     *Auditable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuditableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuditableCallerSession struct {
	Contract *AuditableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AuditableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuditableTransactorSession struct {
	Contract     *AuditableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AuditableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuditableRaw struct {
	Contract *Auditable // Generic contract binding to access the raw methods on
}

// AuditableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuditableCallerRaw struct {
	Contract *AuditableCaller // Generic read-only contract binding to access the raw methods on
}

// AuditableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuditableTransactorRaw struct {
	Contract *AuditableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuditable creates a new instance of Auditable, bound to a specific deployed contract.
func NewAuditable(address common.Address, backend bind.ContractBackend) (*Auditable, error) {
	contract, err := bindAuditable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auditable{AuditableCaller: AuditableCaller{contract: contract}, AuditableTransactor: AuditableTransactor{contract: contract}, AuditableFilterer: AuditableFilterer{contract: contract}}, nil
}

// NewAuditableCaller creates a new read-only instance of Auditable, bound to a specific deployed contract.
func NewAuditableCaller(address common.Address, caller bind.ContractCaller) (*AuditableCaller, error) {
	contract, err := bindAuditable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuditableCaller{contract: contract}, nil
}

// NewAuditableTransactor creates a new write-only instance of Auditable, bound to a specific deployed contract.
func NewAuditableTransactor(address common.Address, transactor bind.ContractTransactor) (*AuditableTransactor, error) {
	contract, err := bindAuditable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuditableTransactor{contract: contract}, nil
}

// NewAuditableFilterer creates a new log filterer instance of Auditable, bound to a specific deployed contract.
func NewAuditableFilterer(address common.Address, filterer bind.ContractFilterer) (*AuditableFilterer, error) {
	contract, err := bindAuditable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuditableFilterer{contract: contract}, nil
}

// bindAuditable binds a generic wrapper to an already deployed contract.
func bindAuditable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuditableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auditable *AuditableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Auditable.Contract.AuditableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auditable *AuditableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auditable.Contract.AuditableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auditable *AuditableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auditable.Contract.AuditableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auditable *AuditableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Auditable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auditable *AuditableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auditable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auditable *AuditableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auditable.Contract.contract.Transact(opts, method, params...)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_Auditable *AuditableCaller) ROLEAUDITOR(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Auditable.contract.Call(opts, out, "ROLE_AUDITOR")
	return *ret0, err
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_Auditable *AuditableSession) ROLEAUDITOR() (string, error) {
	return _Auditable.Contract.ROLEAUDITOR(&_Auditable.CallOpts)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_Auditable *AuditableCallerSession) ROLEAUDITOR() (string, error) {
	return _Auditable.Contract.ROLEAUDITOR(&_Auditable.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Auditable *AuditableCaller) ROLESUPERUSER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Auditable.contract.Call(opts, out, "ROLE_SUPERUSER")
	return *ret0, err
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Auditable *AuditableSession) ROLESUPERUSER() (string, error) {
	return _Auditable.Contract.ROLESUPERUSER(&_Auditable.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Auditable *AuditableCallerSession) ROLESUPERUSER() (string, error) {
	return _Auditable.Contract.ROLESUPERUSER(&_Auditable.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Auditable *AuditableCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _Auditable.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Auditable *AuditableSession) CheckRole(_operator common.Address, _role string) error {
	return _Auditable.Contract.CheckRole(&_Auditable.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Auditable *AuditableCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _Auditable.Contract.CheckRole(&_Auditable.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Auditable *AuditableCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Auditable.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Auditable *AuditableSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Auditable.Contract.HasRole(&_Auditable.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Auditable *AuditableCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Auditable.Contract.HasRole(&_Auditable.CallOpts, _operator, _role)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_Auditable *AuditableCaller) IsAuditor(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Auditable.contract.Call(opts, out, "isAuditor", _addr)
	return *ret0, err
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_Auditable *AuditableSession) IsAuditor(_addr common.Address) (bool, error) {
	return _Auditable.Contract.IsAuditor(&_Auditable.CallOpts, _addr)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_Auditable *AuditableCallerSession) IsAuditor(_addr common.Address) (bool, error) {
	return _Auditable.Contract.IsAuditor(&_Auditable.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Auditable *AuditableCaller) IsSuperuser(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Auditable.contract.Call(opts, out, "isSuperuser", _addr)
	return *ret0, err
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Auditable *AuditableSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _Auditable.Contract.IsSuperuser(&_Auditable.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Auditable *AuditableCallerSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _Auditable.Contract.IsSuperuser(&_Auditable.CallOpts, _addr)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_Auditable *AuditableTransactor) AddAuditor(opts *bind.TransactOpts, _newAuditor common.Address) (*types.Transaction, error) {
	return _Auditable.contract.Transact(opts, "addAuditor", _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_Auditable *AuditableSession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _Auditable.Contract.AddAuditor(&_Auditable.TransactOpts, _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_Auditable *AuditableTransactorSession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _Auditable.Contract.AddAuditor(&_Auditable.TransactOpts, _newAuditor)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Auditable *AuditableTransactor) TransferSuperuser(opts *bind.TransactOpts, _newSuperuser common.Address) (*types.Transaction, error) {
	return _Auditable.contract.Transact(opts, "transferSuperuser", _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Auditable *AuditableSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _Auditable.Contract.TransferSuperuser(&_Auditable.TransactOpts, _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Auditable *AuditableTransactorSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _Auditable.Contract.TransferSuperuser(&_Auditable.TransactOpts, _newSuperuser)
}

// AuditableRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the Auditable contract.
type AuditableRoleAddedIterator struct {
	Event *AuditableRoleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuditableRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuditableRoleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuditableRoleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuditableRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuditableRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuditableRoleAdded represents a RoleAdded event raised by the Auditable contract.
type AuditableRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Auditable *AuditableFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*AuditableRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Auditable.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AuditableRoleAddedIterator{contract: _Auditable.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Auditable *AuditableFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *AuditableRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Auditable.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuditableRoleAdded)
				if err := _Auditable.contract.UnpackLog(event, "RoleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// AuditableRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the Auditable contract.
type AuditableRoleRemovedIterator struct {
	Event *AuditableRoleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AuditableRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuditableRoleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AuditableRoleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AuditableRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuditableRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuditableRoleRemoved represents a RoleRemoved event raised by the Auditable contract.
type AuditableRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Auditable *AuditableFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*AuditableRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Auditable.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AuditableRoleRemovedIterator{contract: _Auditable.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Auditable *AuditableFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *AuditableRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Auditable.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuditableRoleRemoved)
				if err := _Auditable.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMABI is the input ABI used to generate the binding from.
const BREMABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addDeveloper\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_AUDITOR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_withdrawFeePercent\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFeePercent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawFeePercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAuditor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_projectName\",\"type\":\"string\"}],\"name\":\"getProjectByName\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSuperuser\",\"type\":\"address\"}],\"name\":\"transferSuperuser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isDeveloper\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"icoCreationPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"projectsAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"signUp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"login\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isSuperuser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BRM\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_icoCreationPrice\",\"type\":\"uint256\"}],\"name\":\"setIcoCreationPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEVELOPER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAuditor\",\"type\":\"address\"}],\"name\":\"addAuditor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_SUPERUSER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_cap\",\"type\":\"uint256\"},{\"name\":\"_closingTime\",\"type\":\"uint256\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_docHash\",\"type\":\"string\"}],\"name\":\"createBREMICO\",\"outputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"icoAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getProject\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_brmAddress\",\"type\":\"address\"},{\"name\":\"_icoCreationPrice\",\"type\":\"uint256\"},{\"name\":\"_withdrawFeePercent\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"icoAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BREMICOCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// BREMBin is the compiled bytecode used for deploying new contracts.
const BREMBin = `0x60806040523480156200001157600080fd5b5060405160608062003da6833981016040818152825160208085015194830151838501909352600984527f737570657275736572000000000000000000000000000000000000000000000090840152929162000078903390640100000000620000c4810204565b600081101580156200008b575060648111155b15156200009757600080fd5b60078054600160a060020a031916600160a060020a0394909416939093179092556005556006556200020f565b62000140826000836040518082805190602001908083835b60208310620000fd5780518252601f199092019160209182019101620000dc565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050640100000000620001ea8102620018001704565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b83811015620001ab57818101518382015260200162000191565b50505050905090810190601f168015620001d95780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b613b87806200021f6000396000f3006080604052600436106200014a5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c81146200014c578063217fe6c614620001b65780632b7d748914620002345780633f029aaf146200024c578063447cb89614620002dc578063495ef70514620002f757806349b9055714620003215780634b77f269146200034557806357c393fa14620003bd5780635e318e0714620003e15780635eca4a7014620003fc57806380d2370f14620004205780639cdd99131462000438578063adcfb4fc1462000450578063b34e97e814620004ac578063bceee05e14620004c4578063bfb7703014620004e8578063c4f918101462000500578063e31cd22e146200051b578063e429cef11462000533578063ebb4f4841462000557578063ecdf8ac0146200056f578063f0f3f2c814620006bd575b005b3480156200015957600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526200014a958335600160a060020a0316953695604494919390910191908190840183828082843750949750620006d89650505050505050565b348015620001c357600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855262000220958335600160a060020a03169536956044949193909101919081908401838280828437509497506200074a9650505050505050565b604080519115158252519081900360200190f35b3480156200024157600080fd5b506200014a620007c1565b3480156200025957600080fd5b506200026462000800565b6040805160208082528351818301528351919283929083019185019080838360005b83811015620002a057818101518382015260200162000286565b50505050905090810190601f168015620002ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015620002e957600080fd5b506200014a60043562000837565b3480156200030457600080fd5b506200030f6200089b565b60408051918252519081900360200190f35b3480156200032e57600080fd5b5062000220600160a060020a0360043516620008a1565b3480156200035257600080fd5b506040805160206004803580820135601f8101849004840285018401909552848452620003a1943694929360249392840191908190840183828082843750949750620008ea9650505050505050565b60408051600160a060020a039092168252519081900360200190f35b348015620003ca57600080fd5b506200014a600160a060020a0360043516620009a6565b348015620003ee57600080fd5b506200014a60043562000a4f565b3480156200040957600080fd5b5062000220600160a060020a036004351662000acb565b3480156200042d57600080fd5b506200030f62000afd565b3480156200044557600080fd5b506200030f62000b03565b3480156200045d57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526200026494369492936024939284019190819084018382808284375094975062000b099650505050505050565b348015620004b957600080fd5b506200026462000c40565b348015620004d157600080fd5b5062000220600160a060020a036004351662000d17565b348015620004f557600080fd5b50620003a162000d49565b3480156200050d57600080fd5b506200014a60043562000d58565b3480156200052857600080fd5b506200026462000d9d565b3480156200054057600080fd5b506200014a600160a060020a036004351662000dc3565b3480156200056457600080fd5b506200026462000e4a565b3480156200057c57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526200069794369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f60608a01358b0180359182018390048302840183018552818452989b8a359b838c01359b958601359a91995097506080909401955091935091820191819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975062000e709650505050505050565b60408051600160a060020a03938416815291909216602082015281519081900390910190f35b348015620006ca57600080fd5b50620003a1600435620015ab565b62000746826000836040518082805190602001908083835b60208310620007115780518252601f199092019160209182019101620006f0565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050620015d9565b5050565b6000620007ba836000846040518082805190602001908083835b60208310620007855780518252601f19909201916020918201910162000764565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050620015f1565b9392505050565b331515620007ce57600080fd5b620007fe3360408051908101604052806009815260200160008051602062003b3c83398151915281525062001610565b565b60408051808201909152600781527f61756469746f7200000000000000000000000000000000000000000000000000602082015281565b620008673360408051908101604052806009815260200160008051602062003b1c833981519152815250620006d8565b6006548114156200087757600080fd5b600081101580156200088a575060648111155b15156200089657600080fd5b600655565b60065481565b6000620008e4826040805190810160405280600781526020017f61756469746f72000000000000000000000000000000000000000000000000008152506200074a565b92915050565b60008060008351111515620008fe57600080fd5b6004836040518082805190602001908083835b60208310620009325780518252601f19909201916020918201910162000911565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094205460008181526003909252939020549293505050600160a060020a031615156200098857600080fd5b600090815260036020526040902054600160a060020a031692915050565b620009d63360408051908101604052806009815260200160008051602062003b1c833981519152815250620006d8565b600160a060020a0381161515620009ec57600080fd5b62000a1c3360408051908101604052806009815260200160008051602062003b1c83398151915281525062001728565b62000a4c8160408051908101604052806009815260200160008051602062003b1c83398151915281525062001610565b50565b62000a7f3360408051908101604052806009815260200160008051602062003b1c833981519152815250620006d8565b3031811180159062000a915750600081115b151562000a9d57600080fd5b604051339082156108fc029083906000818181858888f1935050505015801562000746573d6000803e3d6000fd5b6000620008e48260408051908101604052806009815260200160008051602062003b3c8339815191528152506200074a565b60055481565b60025481565b6060816000815111151562000b1d57600080fd5b3360009081526001602081905260409091205460029181161561010002600019011604151562000b6b57336000908152600160209081526040909120845162000b699286019062001847565b505b62000b763362000d17565b15801562000b8c575062000b8a33620008a1565b155b1562000b9c5762000b9c620007c1565b3360009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801562000c335780601f1062000c075761010080835404028352916020019162000c33565b820191906000526020600020905b81548152906001019060200180831162000c1557829003601f168201915b5050505050915050919050565b3360009081526001602081905260408220546060926002928216156101000260001901909116919091041162000c7557600080fd5b3360009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801562000d0c5780601f1062000ce05761010080835404028352916020019162000d0c565b820191906000526020600020905b81548152906001019060200180831162000cee57829003601f168201915b505050505090505b90565b6000620008e48260408051908101604052806009815260200160008051602062003b1c8339815191528152506200074a565b600754600160a060020a031681565b62000d883360408051908101604052806009815260200160008051602062003b1c833981519152815250620006d8565b60055481141562000d9857600080fd5b600555565b604080518082019091526009815260008051602062003b3c833981519152602082015281565b62000df33360408051908101604052806009815260200160008051602062003b1c833981519152815250620006d8565b600160a060020a038116151562000e0957600080fd5b62000a4c816040805190810160405280600781526020017f61756469746f720000000000000000000000000000000000000000000000000081525062001610565b604080518082019091526009815260008051602062003b1c833981519152602082015281565b60008060008062000ea63360408051908101604052806009815260200160008051602062003b3c833981519152815250620006d8565b600554600754604080517f70a082310000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a03909216916370a08231916024808201926020929091908290030181600087803b15801562000f1057600080fd5b505af115801562000f25573d6000803e3d6000fd5b505050506040513d602081101562000f3c57600080fd5b5051101562000f4a57600080fd5b60008b5111801562000f5d575060008a51115b151562000f6957600080fd5b60048b6040518082805190602001908083835b6020831062000f9d5780518252601f19909201916020918201910162000f7c565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205415915062000fd8905057600080fd5b606488101562000fe757600080fd5b600754600554604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481019290925251600160a060020a03909216916323b872dd916064808201926020929091908290030181600087803b1580156200105e57600080fd5b505af115801562001073573d6000803e3d6000fd5b505050506040513d60208110156200108a57600080fd5b505115156200109857600080fd5b8a8a620010a4620018cc565b604080825283519082015282518190602080830191606084019187019080838360005b83811015620010e1578181015183820152602001620010c7565b50505050905090810190601f1680156200110f5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015620011445781810151838201526020016200112a565b50505050905090810190601f168015620011725780820380516001836020036101000a031916815260200191505b50945050505050604051809103906000f08015801562001196573d6000803e3d6000fd5b509150819350308b898b33868c8c8c30600654620011b3620018dd565b808c600160a060020a0316600160a060020a03168152602001806020018b81526020018a815260200189600160a060020a0316600160a060020a0316815260200188600160a060020a0316600160a060020a03168152602001878152602001806020018060200186600160a060020a0316600160a060020a0316815260200185815260200184810384528e818151815260200191508051906020019080838360005b838110156200126f57818101518382015260200162001255565b50505050905090810190601f1680156200129d5780820380516001836020036101000a031916815260200191505b5084810383528851815288516020918201918a019080838360005b83811015620012d2578181015183820152602001620012b8565b50505050905090810190601f168015620013005780820380516001836020036101000a031916815260200191505b50848103825287518152875160209182019189019080838360005b83811015620013355781810151838201526020016200131b565b50505050905090810190601f168015620013635780820380516001836020036101000a031916815260200191505b509e505050505050505050505050505050604051809103906000f08015801562001391573d6000803e3d6000fd5b50905080925081600160a060020a031663f2fde38b846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082600160a060020a0316600160a060020a03168152602001915050600060405180830381600087803b1580156200140c57600080fd5b505af115801562001421573d6000803e3d6000fd5b505060028054600090815260036020908152604091829020805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038a16179055915490518f51919450600493508f92909182918401908083835b602083106200149a5780518252601f19909201916020918201910162001479565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208190555060026000815480929190600101919050555083600160a060020a031683600160a060020a031633600160a060020a03167fcbef3bd4943085dfb59f0609e5f754aec11e22a5838e77e0d0c0bff98a0dcdae8e6040518080602001828103825283818151815260200191508051906020019080838360005b838110156200156257818101518382015260200162001548565b50505050905090810190601f168015620015905780820380516001836020036101000a031916815260200191505b509250505060405180910390a4505097509795505050505050565b6002546000908210620015bd57600080fd5b50600090815260036020526040902054600160a060020a031690565b620015e58282620015f1565b15156200074657600080fd5b600160a060020a03166000908152602091909152604090205460ff1690565b6200167e826000836040518082805190602001908083835b60208310620016495780518252601f19909201916020918201910162001628565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505062001800565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b83811015620016e9578181015183820152602001620016cf565b50505050905090810190601f168015620017175780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b62001796826000836040518082805190602001908083835b60208310620017615780518252601f19909201916020918201910162001740565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505062001825565b81600160a060020a03167fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a8260405180806020018281038252838181518152602001915080519060200190808383600083811015620016e9578181015183820152602001620016cf565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b600160a060020a0316600090815260209190915260409020805460ff19169055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200188a57805160ff1916838001178555620018ba565b82800160010185558215620018ba579182015b82811115620018ba5782518255916020019190600101906200189d565b50620018c8929150620018ee565b5090565b604051610e43806200190c83390190565b6040516113cd806200274f83390190565b62000d1491905b80821115620018c85760008155600101620018f5560060806040526003805460a060020a60ff021916905534801561002057600080fd5b50604051610e43380380610e4383398101604052805160208083015160038054600160a060020a0319163317905591830180519093929092019161006a9160049190850190610086565b50805161007e906005906020840190610086565b505050610121565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c757805160ff19168380011785556100f4565b828001600101855582156100f4579182015b828111156100f45782518255916020019190600101906100d9565b50610100929150610104565b5090565b61011e91905b80821115610100576000815560010161010a565b90565b610d13806101306000396000f3006080604052600436106100f05763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100f557806306fdde031461011e578063095ea7b3146101a857806318160ddd146101cc57806323b872dd146101f3578063313ce5671461021d57806340c10f1914610248578063661884631461026c57806370a08231146102905780637d64bcb4146102b15780638da5cb5b146102c657806395d89b41146102f7578063a9059cbb1461030c578063d0482d0414610330578063d73dd62314610356578063dd62ed3e1461037a578063f2fde38b146103a1575b600080fd5b34801561010157600080fd5b5061010a6103c2565b604080519115158252519081900360200190f35b34801561012a57600080fd5b506101336103e3565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016d578181015183820152602001610155565b50505050905090810190601f16801561019a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101b457600080fd5b5061010a600160a060020a0360043516602435610471565b3480156101d857600080fd5b506101e16104d7565b60408051918252519081900360200190f35b3480156101ff57600080fd5b5061010a600160a060020a03600435811690602435166044356104dd565b34801561022957600080fd5b50610232610644565b6040805160ff9092168252519081900360200190f35b34801561025457600080fd5b5061010a600160a060020a0360043516602435610649565b34801561027857600080fd5b5061010a600160a060020a0360043516602435610753565b34801561029c57600080fd5b506101e1600160a060020a0360043516610843565b3480156102bd57600080fd5b5061010a61085e565b3480156102d257600080fd5b506102db610904565b60408051600160a060020a039092168252519081900360200190f35b34801561030357600080fd5b50610133610913565b34801561031857600080fd5b5061010a600160a060020a036004351660243561096e565b34801561033c57600080fd5b50610354600160a060020a0360043516602435610a3f565b005b34801561036257600080fd5b5061010a600160a060020a0360043516602435610b46565b34801561038657600080fd5b506101e1600160a060020a0360043581169060243516610bdf565b3480156103ad57600080fd5b50610354600160a060020a0360043516610c0a565b60035474010000000000000000000000000000000000000000900460ff1681565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104695780601f1061043e57610100808354040283529160200191610469565b820191906000526020600020905b81548152906001019060200180831161044c57829003601f168201915b505050505081565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b6000600160a060020a03831615156104f457600080fd5b600160a060020a03841660009081526001602052604090205482111561051957600080fd5b600160a060020a038416600090815260026020908152604080832033845290915290205482111561054957600080fd5b600160a060020a038416600090815260016020526040902054610572908363ffffffff610c9f16565b600160a060020a0380861660009081526001602052604080822093909355908516815220546105a7908363ffffffff610cb116565b600160a060020a0380851660009081526001602090815260408083209490945591871681526002825282812033825290915220546105eb908363ffffffff610c9f16565b600160a060020a0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610cc8833981519152929181900390910190a35060019392505050565b601281565b600354600090600160a060020a0316331461066357600080fd5b60035474010000000000000000000000000000000000000000900460ff161561068b57600080fd5b60005461069e908363ffffffff610cb116565b6000908155600160a060020a0384168152600160205260409020546106c9908363ffffffff610cb116565b600160a060020a038416600081815260016020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a2604080518381529051600160a060020a03851691600091600080516020610cc88339815191529181900360200190a350600192915050565b336000908152600260209081526040808320600160a060020a0386168452909152812054808311156107a857336000908152600260209081526040808320600160a060020a03881684529091528120556107dd565b6107b8818463ffffffff610c9f16565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526001602052604090205490565b600354600090600160a060020a0316331461087857600080fd5b60035474010000000000000000000000000000000000000000900460ff16156108a057600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b600354600160a060020a031681565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104695780601f1061043e57610100808354040283529160200191610469565b6000600160a060020a038316151561098557600080fd5b336000908152600160205260409020548211156109a157600080fd5b336000908152600160205260409020546109c1908363ffffffff610c9f16565b3360009081526001602052604080822092909255600160a060020a038516815220546109f3908363ffffffff610cb116565b600160a060020a038416600081815260016020908152604091829020939093558051858152905191923392600080516020610cc88339815191529281900390910190a350600192915050565b600354600160a060020a03163314610a5657600080fd5b600160a060020a038216600090815260016020526040902054811115610a7b57600080fd5b600160a060020a038216600090815260016020526040902054610aa4908263ffffffff610c9f16565b600160a060020a03831660009081526001602052604081209190915554610ad1908263ffffffff610c9f16565b600055604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a03851691600080516020610cc88339815191529181900360200190a35050565b336000908152600260209081526040808320600160a060020a0386168452909152812054610b7a908363ffffffff610cb116565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a03163314610c2157600080fd5b600160a060020a0381161515610c3657600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610cab57fe5b50900390565b600082820183811015610cc057fe5b93925050505600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a72305820f411444b49ce555e9034faf955522fd13f9ccde69d0c577c64eaaafe9c322b3a002960806040523480156200001157600080fd5b50604051620013cd380380620013cd83398101604090815281516020830151918301516060840151608085015160a086015160c087015160e08801516101008901516101208a01516101408b0151989a998a019997989697959694959394928401939190910191600160a060020a038b1615156200008e57600080fd5b89516000106200009d57600080fd5b60008911620000ab57600080fd5b60008811620000b957600080fd5b600160a060020a0387161515620000cf57600080fd5b600160a060020a0386161515620000e557600080fd5b600160a060020a0382161515620000fb57600080fd5b428510156200010957600080fd5b600081101580156200011c575060648111155b15156200012857600080fd5b89516200013d9060019060208d0190620001ea565b5060148054600160a060020a03808e16600160a060020a03199283161790925560058b905560038a9055600280548a84169083161790556000805492891692909116919091179055601085905583516200019f906006906020870190620001ea565b508251620001b5906007906020860190620001ea565b5060088054600160a060020a031916600160a060020a039390931692909217909155601355506200028f975050505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022d57805160ff19168380011785556200025d565b828001600101855582156200025d579182015b828111156200025d57825182559160200191906001019062000240565b506200026b9291506200026f565b5090565b6200028c91905b808211156200026b576000815560010162000276565b90565b61112e806200029f6000396000f3006080604052600436106101745763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461017f5780631515bc2b1461020957806318e4ac351461023257806327e235e3146102535780632c4e722e146102865780632e1a7d4d1461029b578063338cdca1146102a6578063355274ea146102d45780633d6df0d5146102e95780634042b66f146102fe57806344aa57001461031357806345e4f34a14610328578063495ef7051461034957806349b905571461035e5780634b6753bc1461037f5780634f93594514610394578063521eb273146103a9578063590e1ae3146103da57806368124f9a146103ef5780637284e4161461040457806393f80ff41461041957806397781f0214610431578063c8e7f59514610446578063e429cef11461045b578063e44e76a01461047c578063ea78d1ab14610491578063ec8ac4d8146104a6578063fad5e0c4146104ba578063fc0c546a146104cf575b61017d336104e4565b005b34801561018b57600080fd5b50610194610598565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ce5781810151838201526020016101b6565b50505050905090810190601f1680156101fb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561021557600080fd5b5061021e610625565b604080519115158252519081900360200190f35b34801561023e57600080fd5b5061021e600160a060020a036004351661062d565b34801561025f57600080fd5b50610274600160a060020a036004351661064b565b60408051918252519081900360200190f35b34801561029257600080fd5b5061027461065d565b61017d600435610663565b3480156102b257600080fd5b506102bb61072b565b6040805192835260208301919091528051918290030190f35b3480156102e057600080fd5b50610274610734565b3480156102f557600080fd5b5061019461073a565b34801561030a57600080fd5b50610274610795565b34801561031f57600080fd5b5061021e61079b565b34801561033457600080fd5b50610274600160a060020a03600435166107c6565b34801561035557600080fd5b506102746107d8565b34801561036a57600080fd5b5061021e600160a060020a03600435166107de565b34801561038b57600080fd5b506102746107fc565b3480156103a057600080fd5b5061021e610802565b3480156103b557600080fd5b506103be61080d565b60408051600160a060020a039092168252519081900360200190f35b3480156103e657600080fd5b5061017d61081c565b3480156103fb57600080fd5b5061017d61094c565b34801561041057600080fd5b50610194610bba565b34801561042557600080fd5b506103be600435610c15565b34801561043d57600080fd5b5061021e610c42565b34801561045257600080fd5b50610274610c7e565b34801561046757600080fd5b5061017d600160a060020a0360043516610c84565b34801561048857600080fd5b5061021e610e60565b34801561049d57600080fd5b5061017d610e69565b61017d600160a060020a03600435166104e4565b3480156104c657600080fd5b506103be610f43565b3480156104db57600080fd5b506103be610f52565b600c54600090819060ff1680156104fd57506010544211155b151561050857600080fd5b3491506105158383610f61565b61051e82610f86565b600454909150610534908363ffffffff610fa316565b6004556105418382610fbd565b60408051838152602081018390528151600160a060020a0386169233927f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18929081900390910190a36105938383610fc7565b505050565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561061d5780601f106105f25761010080835404028352916020019161061d565b820191906000526020600020905b81548152906001019060200180831161060057829003601f168201915b505050505081565b601054421190565b600160a060020a03166000908152600f602052604090205460ff1690565b60116020526000908152604090205481565b60035481565b600254600160a060020a0316331461067a57600080fd5b600c5460ff16151561068b57600080fd5b610693610802565b80156106a257506106a2610625565b15156106ad57600080fd5b606481101580156106bf575030318111155b15156106ca57600080fd5b600d54156106d757600080fd5b6064303182900310156107085760408051808201909152303180825260006020909201829052600d55600e55610728565b6040805180820190915281815260006020909101819052600d829055600e555b50565b600d54600e5482565b60055481565b6007805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561061d5780601f106105f25761010080835404028352916020019161061d565b60045481565b60006107a5610625565b80156107b457506107b4610802565b80156107c1575060643031105b905090565b60126020526000908152604090205481565b60135481565b600160a060020a031660009081526009602052604090205460ff1690565b60105481565b600554600454101590565b600254600160a060020a031681565b600080610827610625565b80156108385750610836610802565b155b151561084357600080fd5b336000908152601160205260408120541161085d57600080fd5b505033600081815260116020908152604080832080546012909352818420805491859055849055835482517fd0482d04000000000000000000000000000000000000000000000000000000008152600481019690965260248601829052915192949093600160a060020a039092169263d0482d049260448084019391929182900301818387803b1580156108f057600080fd5b505af1158015610904573d6000803e3d6000fd5b505060045461091c925090508363ffffffff61101216565b600455604051339083156108fc029084906000818181858888f19350505050158015610593573d6000803e3d6000fd5b600854604080517f49b905570000000000000000000000000000000000000000000000000000000081523360048201529051600092839283928392600160a060020a0316916349b9055791602480830192602092919082900301818787803b1580156109b757600080fd5b505af11580156109cb573d6000803e3d6000fd5b505050506040513d60208110156109e157600080fd5b505115156109ee57600080fd5b600d546000106109fd57600080fd5b336000908152600f602052604090205460ff1615610a1a57600080fd5b600c5460ff161515610a2b57600080fd5b610a33610802565b8015610a425750610a42610625565b1515610a4d57600080fd5b3360009081526009602052604090205460ff161515610a6b57600080fd5b336000908152600f60205260409020805460ff19166001908117909155600e805490910190819055600b541415610bb457600d8054604080518082019091526000808252602090910181905291829055600e829055945092505b600b54831015610b06576000838152600a6020908152604080832054600160a060020a03168352600f9091529020805460ff19169055600190920191610ac5565b601354610b2a90610b1e86606463ffffffff61102416565b9063ffffffff61103b16565b9150610b3c848363ffffffff61101216565b601454604051919250600160a060020a03169083156108fc029084906000818181858888f19350505050158015610b77573d6000803e3d6000fd5b50600254604051600160a060020a039091169082156108fc029083906000818181858888f19350505050158015610bb2573d6000803e3d6000fd5b505b50505050565b6006805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561061d5780601f106105f25761010080835404028352916020019161061d565b600b546000908210610c2657600080fd5b506000908152600a6020526040902054600160a060020a031690565b6000610c4c610802565b8015610c5b5750610c5b610625565b8015610c695750600d546000105b80156107c15750610c7861079b565b15905090565b600b5481565b601054421115610c9357600080fd5b600c5460ff1615610ca357600080fd5b600854604080517fbceee05e0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163bceee05e916024808201926020929091908290030181600087803b158015610d0957600080fd5b505af1158015610d1d573d6000803e3d6000fd5b505050506040513d6020811015610d3357600080fd5b50511515610d4057600080fd5b600854604080517f49b90557000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915191909216916349b905579160248083019260209291908290030181600087803b158015610da857600080fd5b505af1158015610dbc573d6000803e3d6000fd5b505050506040513d6020811015610dd257600080fd5b50511515610ddf57600080fd5b600160a060020a03811660009081526009602052604090205460ff1615610e0557600080fd5b600160a060020a03166000818152600960209081526040808320805460ff19166001908117909155600b80548552600a9093529220805473ffffffffffffffffffffffffffffffffffffffff19169093179092558154019055565b600c5460ff1681565b601054421115610e7857600080fd5b600c5460ff1615610e8857600080fd5b600854604080517fbceee05e0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163bceee05e916024808201926020929091908290030181600087803b158015610eee57600080fd5b505af1158015610f02573d6000803e3d6000fd5b505050506040513d6020811015610f1857600080fd5b50511515610f2557600080fd5b600b54600010610f3457600080fd5b600c805460ff19166001179055565b601454600160a060020a031681565b600054600160a060020a031681565b600160a060020a0382161515610f7657600080fd5b801515610f8257600080fd5b5050565b6000610f9d6003548361103b90919063ffffffff16565b92915050565b600082820183811015610fb257fe5b8091505b5092915050565b610f828282611066565b600160a060020a0382166000908152601160205260409020805482019055610fee81610f86565b600160a060020a039092166000908152601260205260409020805490920190915550565b60008282111561101e57fe5b50900390565b600080828481151561103257fe5b04949350505050565b60008083151561104e5760009150610fb6565b5082820282848281151561105e57fe5b0414610fb257fe5b60008054604080517f40c10f19000000000000000000000000000000000000000000000000000000008152600160a060020a03868116600483015260248201869052915191909216926340c10f1992604480820193602093909283900390910190829087803b1580156110d857600080fd5b505af11580156110ec573d6000803e3d6000fd5b505050506040513d6020811015610bb457600080fd00a165627a7a72305820670adf9b43ff56ea291ac7240a801a880392912a6ad5023622a0ce670320d27900297375706572757365720000000000000000000000000000000000000000000000646576656c6f7065720000000000000000000000000000000000000000000000a165627a7a7230582077912e6abda7ce87a5139b93a8a62601a2252714c45c94f0cede620ca10284c60029`

// DeployBREM deploys a new Ethereum contract, binding an instance of BREM to it.
func DeployBREM(auth *bind.TransactOpts, backend bind.ContractBackend, _brmAddress common.Address, _icoCreationPrice *big.Int, _withdrawFeePercent *big.Int) (common.Address, *types.Transaction, *BREM, error) {
	parsed, err := abi.JSON(strings.NewReader(BREMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BREMBin), backend, _brmAddress, _icoCreationPrice, _withdrawFeePercent)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BREM{BREMCaller: BREMCaller{contract: contract}, BREMTransactor: BREMTransactor{contract: contract}, BREMFilterer: BREMFilterer{contract: contract}}, nil
}

// BREM is an auto generated Go binding around an Ethereum contract.
type BREM struct {
	BREMCaller     // Read-only binding to the contract
	BREMTransactor // Write-only binding to the contract
	BREMFilterer   // Log filterer for contract events
}

// BREMCaller is an auto generated read-only Go binding around an Ethereum contract.
type BREMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BREMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BREMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BREMSession struct {
	Contract     *BREM             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BREMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BREMCallerSession struct {
	Contract *BREMCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BREMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BREMTransactorSession struct {
	Contract     *BREMTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BREMRaw is an auto generated low-level Go binding around an Ethereum contract.
type BREMRaw struct {
	Contract *BREM // Generic contract binding to access the raw methods on
}

// BREMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BREMCallerRaw struct {
	Contract *BREMCaller // Generic read-only contract binding to access the raw methods on
}

// BREMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BREMTransactorRaw struct {
	Contract *BREMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBREM creates a new instance of BREM, bound to a specific deployed contract.
func NewBREM(address common.Address, backend bind.ContractBackend) (*BREM, error) {
	contract, err := bindBREM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BREM{BREMCaller: BREMCaller{contract: contract}, BREMTransactor: BREMTransactor{contract: contract}, BREMFilterer: BREMFilterer{contract: contract}}, nil
}

// NewBREMCaller creates a new read-only instance of BREM, bound to a specific deployed contract.
func NewBREMCaller(address common.Address, caller bind.ContractCaller) (*BREMCaller, error) {
	contract, err := bindBREM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BREMCaller{contract: contract}, nil
}

// NewBREMTransactor creates a new write-only instance of BREM, bound to a specific deployed contract.
func NewBREMTransactor(address common.Address, transactor bind.ContractTransactor) (*BREMTransactor, error) {
	contract, err := bindBREM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BREMTransactor{contract: contract}, nil
}

// NewBREMFilterer creates a new log filterer instance of BREM, bound to a specific deployed contract.
func NewBREMFilterer(address common.Address, filterer bind.ContractFilterer) (*BREMFilterer, error) {
	contract, err := bindBREM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BREMFilterer{contract: contract}, nil
}

// bindBREM binds a generic wrapper to an already deployed contract.
func bindBREM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BREMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BREM *BREMRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BREM.Contract.BREMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BREM *BREMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREM.Contract.BREMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BREM *BREMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BREM.Contract.BREMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BREM *BREMCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BREM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BREM *BREMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BREM *BREMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BREM.Contract.contract.Transact(opts, method, params...)
}

// BRM is a free data retrieval call binding the contract method 0xbfb77030.
//
// Solidity: function BRM() constant returns(address)
func (_BREM *BREMCaller) BRM(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "BRM")
	return *ret0, err
}

// BRM is a free data retrieval call binding the contract method 0xbfb77030.
//
// Solidity: function BRM() constant returns(address)
func (_BREM *BREMSession) BRM() (common.Address, error) {
	return _BREM.Contract.BRM(&_BREM.CallOpts)
}

// BRM is a free data retrieval call binding the contract method 0xbfb77030.
//
// Solidity: function BRM() constant returns(address)
func (_BREM *BREMCallerSession) BRM() (common.Address, error) {
	return _BREM.Contract.BRM(&_BREM.CallOpts)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_BREM *BREMCaller) ROLEAUDITOR(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "ROLE_AUDITOR")
	return *ret0, err
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_BREM *BREMSession) ROLEAUDITOR() (string, error) {
	return _BREM.Contract.ROLEAUDITOR(&_BREM.CallOpts)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_BREM *BREMCallerSession) ROLEAUDITOR() (string, error) {
	return _BREM.Contract.ROLEAUDITOR(&_BREM.CallOpts)
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_BREM *BREMCaller) ROLEDEVELOPER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "ROLE_DEVELOPER")
	return *ret0, err
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_BREM *BREMSession) ROLEDEVELOPER() (string, error) {
	return _BREM.Contract.ROLEDEVELOPER(&_BREM.CallOpts)
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_BREM *BREMCallerSession) ROLEDEVELOPER() (string, error) {
	return _BREM.Contract.ROLEDEVELOPER(&_BREM.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_BREM *BREMCaller) ROLESUPERUSER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "ROLE_SUPERUSER")
	return *ret0, err
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_BREM *BREMSession) ROLESUPERUSER() (string, error) {
	return _BREM.Contract.ROLESUPERUSER(&_BREM.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_BREM *BREMCallerSession) ROLESUPERUSER() (string, error) {
	return _BREM.Contract.ROLESUPERUSER(&_BREM.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_BREM *BREMCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _BREM.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_BREM *BREMSession) CheckRole(_operator common.Address, _role string) error {
	return _BREM.Contract.CheckRole(&_BREM.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_BREM *BREMCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _BREM.Contract.CheckRole(&_BREM.CallOpts, _operator, _role)
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(_index uint256) constant returns(address)
func (_BREM *BREMCaller) GetProject(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "getProject", _index)
	return *ret0, err
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(_index uint256) constant returns(address)
func (_BREM *BREMSession) GetProject(_index *big.Int) (common.Address, error) {
	return _BREM.Contract.GetProject(&_BREM.CallOpts, _index)
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(_index uint256) constant returns(address)
func (_BREM *BREMCallerSession) GetProject(_index *big.Int) (common.Address, error) {
	return _BREM.Contract.GetProject(&_BREM.CallOpts, _index)
}

// GetProjectByName is a free data retrieval call binding the contract method 0x4b77f269.
//
// Solidity: function getProjectByName(_projectName string) constant returns(address)
func (_BREM *BREMCaller) GetProjectByName(opts *bind.CallOpts, _projectName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "getProjectByName", _projectName)
	return *ret0, err
}

// GetProjectByName is a free data retrieval call binding the contract method 0x4b77f269.
//
// Solidity: function getProjectByName(_projectName string) constant returns(address)
func (_BREM *BREMSession) GetProjectByName(_projectName string) (common.Address, error) {
	return _BREM.Contract.GetProjectByName(&_BREM.CallOpts, _projectName)
}

// GetProjectByName is a free data retrieval call binding the contract method 0x4b77f269.
//
// Solidity: function getProjectByName(_projectName string) constant returns(address)
func (_BREM *BREMCallerSession) GetProjectByName(_projectName string) (common.Address, error) {
	return _BREM.Contract.GetProjectByName(&_BREM.CallOpts, _projectName)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_BREM *BREMCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_BREM *BREMSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _BREM.Contract.HasRole(&_BREM.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_BREM *BREMCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _BREM.Contract.HasRole(&_BREM.CallOpts, _operator, _role)
}

// IcoCreationPrice is a free data retrieval call binding the contract method 0x80d2370f.
//
// Solidity: function icoCreationPrice() constant returns(uint256)
func (_BREM *BREMCaller) IcoCreationPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "icoCreationPrice")
	return *ret0, err
}

// IcoCreationPrice is a free data retrieval call binding the contract method 0x80d2370f.
//
// Solidity: function icoCreationPrice() constant returns(uint256)
func (_BREM *BREMSession) IcoCreationPrice() (*big.Int, error) {
	return _BREM.Contract.IcoCreationPrice(&_BREM.CallOpts)
}

// IcoCreationPrice is a free data retrieval call binding the contract method 0x80d2370f.
//
// Solidity: function icoCreationPrice() constant returns(uint256)
func (_BREM *BREMCallerSession) IcoCreationPrice() (*big.Int, error) {
	return _BREM.Contract.IcoCreationPrice(&_BREM.CallOpts)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_BREM *BREMCaller) IsAuditor(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "isAuditor", _addr)
	return *ret0, err
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_BREM *BREMSession) IsAuditor(_addr common.Address) (bool, error) {
	return _BREM.Contract.IsAuditor(&_BREM.CallOpts, _addr)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_BREM *BREMCallerSession) IsAuditor(_addr common.Address) (bool, error) {
	return _BREM.Contract.IsAuditor(&_BREM.CallOpts, _addr)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_BREM *BREMCaller) IsDeveloper(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "isDeveloper", _addr)
	return *ret0, err
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_BREM *BREMSession) IsDeveloper(_addr common.Address) (bool, error) {
	return _BREM.Contract.IsDeveloper(&_BREM.CallOpts, _addr)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_BREM *BREMCallerSession) IsDeveloper(_addr common.Address) (bool, error) {
	return _BREM.Contract.IsDeveloper(&_BREM.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_BREM *BREMCaller) IsSuperuser(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "isSuperuser", _addr)
	return *ret0, err
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_BREM *BREMSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _BREM.Contract.IsSuperuser(&_BREM.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_BREM *BREMCallerSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _BREM.Contract.IsSuperuser(&_BREM.CallOpts, _addr)
}

// Login is a free data retrieval call binding the contract method 0xb34e97e8.
//
// Solidity: function login() constant returns(string)
func (_BREM *BREMCaller) Login(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "login")
	return *ret0, err
}

// Login is a free data retrieval call binding the contract method 0xb34e97e8.
//
// Solidity: function login() constant returns(string)
func (_BREM *BREMSession) Login() (string, error) {
	return _BREM.Contract.Login(&_BREM.CallOpts)
}

// Login is a free data retrieval call binding the contract method 0xb34e97e8.
//
// Solidity: function login() constant returns(string)
func (_BREM *BREMCallerSession) Login() (string, error) {
	return _BREM.Contract.Login(&_BREM.CallOpts)
}

// ProjectsAmount is a free data retrieval call binding the contract method 0x9cdd9913.
//
// Solidity: function projectsAmount() constant returns(uint256)
func (_BREM *BREMCaller) ProjectsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "projectsAmount")
	return *ret0, err
}

// ProjectsAmount is a free data retrieval call binding the contract method 0x9cdd9913.
//
// Solidity: function projectsAmount() constant returns(uint256)
func (_BREM *BREMSession) ProjectsAmount() (*big.Int, error) {
	return _BREM.Contract.ProjectsAmount(&_BREM.CallOpts)
}

// ProjectsAmount is a free data retrieval call binding the contract method 0x9cdd9913.
//
// Solidity: function projectsAmount() constant returns(uint256)
func (_BREM *BREMCallerSession) ProjectsAmount() (*big.Int, error) {
	return _BREM.Contract.ProjectsAmount(&_BREM.CallOpts)
}

// WithdrawFeePercent is a free data retrieval call binding the contract method 0x495ef705.
//
// Solidity: function withdrawFeePercent() constant returns(uint256)
func (_BREM *BREMCaller) WithdrawFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREM.contract.Call(opts, out, "withdrawFeePercent")
	return *ret0, err
}

// WithdrawFeePercent is a free data retrieval call binding the contract method 0x495ef705.
//
// Solidity: function withdrawFeePercent() constant returns(uint256)
func (_BREM *BREMSession) WithdrawFeePercent() (*big.Int, error) {
	return _BREM.Contract.WithdrawFeePercent(&_BREM.CallOpts)
}

// WithdrawFeePercent is a free data retrieval call binding the contract method 0x495ef705.
//
// Solidity: function withdrawFeePercent() constant returns(uint256)
func (_BREM *BREMCallerSession) WithdrawFeePercent() (*big.Int, error) {
	return _BREM.Contract.WithdrawFeePercent(&_BREM.CallOpts)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_BREM *BREMTransactor) AddAuditor(opts *bind.TransactOpts, _newAuditor common.Address) (*types.Transaction, error) {
	return _BREM.contract.Transact(opts, "addAuditor", _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_BREM *BREMSession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _BREM.Contract.AddAuditor(&_BREM.TransactOpts, _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_BREM *BREMTransactorSession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _BREM.Contract.AddAuditor(&_BREM.TransactOpts, _newAuditor)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_BREM *BREMTransactor) AddDeveloper(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREM.contract.Transact(opts, "addDeveloper")
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_BREM *BREMSession) AddDeveloper() (*types.Transaction, error) {
	return _BREM.Contract.AddDeveloper(&_BREM.TransactOpts)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_BREM *BREMTransactorSession) AddDeveloper() (*types.Transaction, error) {
	return _BREM.Contract.AddDeveloper(&_BREM.TransactOpts)
}

// CreateBREMICO is a paid mutator transaction binding the contract method 0xecdf8ac0.
//
// Solidity: function createBREMICO(_name string, _symbol string, _rate uint256, _cap uint256, _closingTime uint256, _description string, _docHash string) returns(tokenAddress address, icoAddress address)
func (_BREM *BREMTransactor) CreateBREMICO(opts *bind.TransactOpts, _name string, _symbol string, _rate *big.Int, _cap *big.Int, _closingTime *big.Int, _description string, _docHash string) (*types.Transaction, error) {
	return _BREM.contract.Transact(opts, "createBREMICO", _name, _symbol, _rate, _cap, _closingTime, _description, _docHash)
}

// CreateBREMICO is a paid mutator transaction binding the contract method 0xecdf8ac0.
//
// Solidity: function createBREMICO(_name string, _symbol string, _rate uint256, _cap uint256, _closingTime uint256, _description string, _docHash string) returns(tokenAddress address, icoAddress address)
func (_BREM *BREMSession) CreateBREMICO(_name string, _symbol string, _rate *big.Int, _cap *big.Int, _closingTime *big.Int, _description string, _docHash string) (*types.Transaction, error) {
	return _BREM.Contract.CreateBREMICO(&_BREM.TransactOpts, _name, _symbol, _rate, _cap, _closingTime, _description, _docHash)
}

// CreateBREMICO is a paid mutator transaction binding the contract method 0xecdf8ac0.
//
// Solidity: function createBREMICO(_name string, _symbol string, _rate uint256, _cap uint256, _closingTime uint256, _description string, _docHash string) returns(tokenAddress address, icoAddress address)
func (_BREM *BREMTransactorSession) CreateBREMICO(_name string, _symbol string, _rate *big.Int, _cap *big.Int, _closingTime *big.Int, _description string, _docHash string) (*types.Transaction, error) {
	return _BREM.Contract.CreateBREMICO(&_BREM.TransactOpts, _name, _symbol, _rate, _cap, _closingTime, _description, _docHash)
}

// SetIcoCreationPrice is a paid mutator transaction binding the contract method 0xc4f91810.
//
// Solidity: function setIcoCreationPrice(_icoCreationPrice uint256) returns()
func (_BREM *BREMTransactor) SetIcoCreationPrice(opts *bind.TransactOpts, _icoCreationPrice *big.Int) (*types.Transaction, error) {
	return _BREM.contract.Transact(opts, "setIcoCreationPrice", _icoCreationPrice)
}

// SetIcoCreationPrice is a paid mutator transaction binding the contract method 0xc4f91810.
//
// Solidity: function setIcoCreationPrice(_icoCreationPrice uint256) returns()
func (_BREM *BREMSession) SetIcoCreationPrice(_icoCreationPrice *big.Int) (*types.Transaction, error) {
	return _BREM.Contract.SetIcoCreationPrice(&_BREM.TransactOpts, _icoCreationPrice)
}

// SetIcoCreationPrice is a paid mutator transaction binding the contract method 0xc4f91810.
//
// Solidity: function setIcoCreationPrice(_icoCreationPrice uint256) returns()
func (_BREM *BREMTransactorSession) SetIcoCreationPrice(_icoCreationPrice *big.Int) (*types.Transaction, error) {
	return _BREM.Contract.SetIcoCreationPrice(&_BREM.TransactOpts, _icoCreationPrice)
}

// SetWithdrawFeePercent is a paid mutator transaction binding the contract method 0x447cb896.
//
// Solidity: function setWithdrawFeePercent(_withdrawFeePercent uint256) returns()
func (_BREM *BREMTransactor) SetWithdrawFeePercent(opts *bind.TransactOpts, _withdrawFeePercent *big.Int) (*types.Transaction, error) {
	return _BREM.contract.Transact(opts, "setWithdrawFeePercent", _withdrawFeePercent)
}

// SetWithdrawFeePercent is a paid mutator transaction binding the contract method 0x447cb896.
//
// Solidity: function setWithdrawFeePercent(_withdrawFeePercent uint256) returns()
func (_BREM *BREMSession) SetWithdrawFeePercent(_withdrawFeePercent *big.Int) (*types.Transaction, error) {
	return _BREM.Contract.SetWithdrawFeePercent(&_BREM.TransactOpts, _withdrawFeePercent)
}

// SetWithdrawFeePercent is a paid mutator transaction binding the contract method 0x447cb896.
//
// Solidity: function setWithdrawFeePercent(_withdrawFeePercent uint256) returns()
func (_BREM *BREMTransactorSession) SetWithdrawFeePercent(_withdrawFeePercent *big.Int) (*types.Transaction, error) {
	return _BREM.Contract.SetWithdrawFeePercent(&_BREM.TransactOpts, _withdrawFeePercent)
}

// SignUp is a paid mutator transaction binding the contract method 0xadcfb4fc.
//
// Solidity: function signUp(_name string) returns(string)
func (_BREM *BREMTransactor) SignUp(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _BREM.contract.Transact(opts, "signUp", _name)
}

// SignUp is a paid mutator transaction binding the contract method 0xadcfb4fc.
//
// Solidity: function signUp(_name string) returns(string)
func (_BREM *BREMSession) SignUp(_name string) (*types.Transaction, error) {
	return _BREM.Contract.SignUp(&_BREM.TransactOpts, _name)
}

// SignUp is a paid mutator transaction binding the contract method 0xadcfb4fc.
//
// Solidity: function signUp(_name string) returns(string)
func (_BREM *BREMTransactorSession) SignUp(_name string) (*types.Transaction, error) {
	return _BREM.Contract.SignUp(&_BREM.TransactOpts, _name)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_BREM *BREMTransactor) TransferSuperuser(opts *bind.TransactOpts, _newSuperuser common.Address) (*types.Transaction, error) {
	return _BREM.contract.Transact(opts, "transferSuperuser", _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_BREM *BREMSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _BREM.Contract.TransferSuperuser(&_BREM.TransactOpts, _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_BREM *BREMTransactorSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _BREM.Contract.TransferSuperuser(&_BREM.TransactOpts, _newSuperuser)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x5e318e07.
//
// Solidity: function withdrawFees(_value uint256) returns()
func (_BREM *BREMTransactor) WithdrawFees(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BREM.contract.Transact(opts, "withdrawFees", _value)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x5e318e07.
//
// Solidity: function withdrawFees(_value uint256) returns()
func (_BREM *BREMSession) WithdrawFees(_value *big.Int) (*types.Transaction, error) {
	return _BREM.Contract.WithdrawFees(&_BREM.TransactOpts, _value)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x5e318e07.
//
// Solidity: function withdrawFees(_value uint256) returns()
func (_BREM *BREMTransactorSession) WithdrawFees(_value *big.Int) (*types.Transaction, error) {
	return _BREM.Contract.WithdrawFees(&_BREM.TransactOpts, _value)
}

// BREMBREMICOCreatedIterator is returned from FilterBREMICOCreated and is used to iterate over the raw logs and unpacked data for BREMICOCreated events raised by the BREM contract.
type BREMBREMICOCreatedIterator struct {
	Event *BREMBREMICOCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMBREMICOCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMBREMICOCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMBREMICOCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMBREMICOCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMBREMICOCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMBREMICOCreated represents a BREMICOCreated event raised by the BREM contract.
type BREMBREMICOCreated struct {
	Creator      common.Address
	IcoAddress   common.Address
	TokenAddress common.Address
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBREMICOCreated is a free log retrieval operation binding the contract event 0xcbef3bd4943085dfb59f0609e5f754aec11e22a5838e77e0d0c0bff98a0dcdae.
//
// Solidity: e BREMICOCreated(creator indexed address, icoAddress indexed address, tokenAddress indexed address, name string)
func (_BREM *BREMFilterer) FilterBREMICOCreated(opts *bind.FilterOpts, creator []common.Address, icoAddress []common.Address, tokenAddress []common.Address) (*BREMBREMICOCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var icoAddressRule []interface{}
	for _, icoAddressItem := range icoAddress {
		icoAddressRule = append(icoAddressRule, icoAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BREM.contract.FilterLogs(opts, "BREMICOCreated", creatorRule, icoAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BREMBREMICOCreatedIterator{contract: _BREM.contract, event: "BREMICOCreated", logs: logs, sub: sub}, nil
}

// WatchBREMICOCreated is a free log subscription operation binding the contract event 0xcbef3bd4943085dfb59f0609e5f754aec11e22a5838e77e0d0c0bff98a0dcdae.
//
// Solidity: e BREMICOCreated(creator indexed address, icoAddress indexed address, tokenAddress indexed address, name string)
func (_BREM *BREMFilterer) WatchBREMICOCreated(opts *bind.WatchOpts, sink chan<- *BREMBREMICOCreated, creator []common.Address, icoAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var icoAddressRule []interface{}
	for _, icoAddressItem := range icoAddress {
		icoAddressRule = append(icoAddressRule, icoAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BREM.contract.WatchLogs(opts, "BREMICOCreated", creatorRule, icoAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMBREMICOCreated)
				if err := _BREM.contract.UnpackLog(event, "BREMICOCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the BREM contract.
type BREMRoleAddedIterator struct {
	Event *BREMRoleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMRoleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMRoleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMRoleAdded represents a RoleAdded event raised by the BREM contract.
type BREMRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_BREM *BREMFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*BREMRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BREM.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BREMRoleAddedIterator{contract: _BREM.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_BREM *BREMFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *BREMRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BREM.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMRoleAdded)
				if err := _BREM.contract.UnpackLog(event, "RoleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the BREM contract.
type BREMRoleRemovedIterator struct {
	Event *BREMRoleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMRoleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMRoleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMRoleRemoved represents a RoleRemoved event raised by the BREM contract.
type BREMRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_BREM *BREMFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*BREMRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BREM.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BREMRoleRemovedIterator{contract: _BREM.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_BREM *BREMFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *BREMRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BREM.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMRoleRemoved)
				if err := _BREM.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMFactoryABI is the input ABI used to generate the binding from.
const BREMFactoryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addDeveloper\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_AUDITOR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_withdrawFeePercent\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFeePercent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawFeePercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAuditor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_projectName\",\"type\":\"string\"}],\"name\":\"getProjectByName\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSuperuser\",\"type\":\"address\"}],\"name\":\"transferSuperuser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isDeveloper\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"icoCreationPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"projectsAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"signUp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"login\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isSuperuser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_icoCreationPrice\",\"type\":\"uint256\"}],\"name\":\"setIcoCreationPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEVELOPER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAuditor\",\"type\":\"address\"}],\"name\":\"addAuditor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_SUPERUSER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getProject\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"icoAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"BREMICOCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// BREMFactoryBin is the compiled bytecode used for deploying new contracts.
const BREMFactoryBin = `0x60c0604052600960809081527f737570657275736572000000000000000000000000000000000000000000000060a052610043903390640100000000610048810204565b61018b565b6100bf826000836040518082805190602001908083835b6020831061007e5780518252601f19909201916020918201910161005f565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929150506401000000006101668102610e221704565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610128578181015183820152602001610110565b50505050905090810190601f1680156101555780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b610f4d8061019a6000396000f3006080604052600436106101115763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c8114610116578063217fe6c61461017f5780632b7d7489146101fa5780633f029aaf1461020f578063447cb89614610299578063495ef705146102b157806349b90557146102d85780634b77f269146102f957806357c393fa1461036e5780635eca4a701461038f57806380d2370f146103b05780639cdd9913146103c5578063adcfb4fc146103da578063b34e97e814610433578063bceee05e14610448578063c4f9181014610469578063e31cd22e14610481578063e429cef114610496578063ebb4f484146104b7578063f0f3f2c8146104cc575b600080fd5b34801561012257600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261017d958335600160a060020a03169536956044949193909101919081908401838280828437509497506104e49650505050505050565b005b34801561018b57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526101e6958335600160a060020a03169536956044949193909101919081908401838280828437509497506105529650505050505050565b604080519115158252519081900360200190f35b34801561020657600080fd5b5061017d6105c5565b34801561021b57600080fd5b50610224610612565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561025e578181015183820152602001610246565b50505050905090810190601f16801561028b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102a557600080fd5b5061017d600435610649565b3480156102bd57600080fd5b506102c66106a7565b60408051918252519081900360200190f35b3480156102e457600080fd5b506101e6600160a060020a03600435166106ad565b34801561030557600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526103529436949293602493928401919081908401838280828437509497506106f49650505050505050565b60408051600160a060020a039092168252519081900360200190f35b34801561037a57600080fd5b5061017d600160a060020a03600435166107ac565b34801561039b57600080fd5b506101e6600160a060020a036004351661084b565b3480156103bc57600080fd5b506102c661088c565b3480156103d157600080fd5b506102c6610892565b3480156103e657600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526102249436949293602493928401919081908401838280828437509497506108989650505050505050565b34801561043f57600080fd5b506102246109bf565b34801561045457600080fd5b506101e6600160a060020a0360043516610a91565b34801561047557600080fd5b5061017d600435610ac0565b34801561048d57600080fd5b50610224610b01565b3480156104a257600080fd5b5061017d600160a060020a0360043516610b38565b3480156104c357600080fd5b50610224610bb9565b3480156104d857600080fd5b50610352600435610bde565b61054e826000836040518082805190602001908083835b6020831061051a5780518252601f1990920191602091820191016104fb565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610c0b565b5050565b60006105be836000846040518082805190602001908083835b6020831061058a5780518252601f19909201916020918201910161056b565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610c20565b9392505050565b3315156105d157600080fd5b610610336040805190810160405280600981526020017f646576656c6f7065720000000000000000000000000000000000000000000000815250610c3f565b565b60408051808201909152600781527f61756469746f7200000000000000000000000000000000000000000000000000602082015281565b61067633604080519081016040528060098152602001600080516020610f028339815191528152506104e4565b60065481141561068557600080fd5b60008110158015610697575060648111155b15156106a257600080fd5b600655565b60065481565b60006106ee826040805190810160405280600781526020017f61756469746f7200000000000000000000000000000000000000000000000000815250610552565b92915050565b6000806000835111151561070757600080fd5b6004836040518082805190602001908083835b602083106107395780518252601f19909201916020918201910161071a565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094205460008181526003909252939020549293505050600160a060020a0316151561078e57600080fd5b600090815260036020526040902054600160a060020a031692915050565b6107d933604080519081016040528060098152602001600080516020610f028339815191528152506104e4565b600160a060020a03811615156107ee57600080fd5b61081b33604080519081016040528060098152602001600080516020610f02833981519152815250610d50565b61084881604080519081016040528060098152602001600080516020610f02833981519152815250610c3f565b50565b60006106ee826040805190810160405280600981526020017f646576656c6f7065720000000000000000000000000000000000000000000000815250610552565b60055481565b60025481565b606081600081511115156108ab57600080fd5b336000908152600160208190526040909120546002918116156101000260001901160415156108f65733600090815260016020908152604090912084516108f492860190610e69565b505b6108ff33610a91565b1580156109125750610910336106ad565b155b1561091f5761091f6105c5565b3360009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156109b25780601f10610987576101008083540402835291602001916109b2565b820191906000526020600020905b81548152906001019060200180831161099557829003601f168201915b5050505050915050919050565b336000908152600160208190526040822054606092600292821615610100026000190190911691909104116109f357600080fd5b3360009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f8101839004830284018301909452838352919290830182828015610a865780601f10610a5b57610100808354040283529160200191610a86565b820191906000526020600020905b815481529060010190602001808311610a6957829003601f168201915b505050505090505b90565b60006106ee82604080519081016040528060098152602001600080516020610f02833981519152815250610552565b610aed33604080519081016040528060098152602001600080516020610f028339815191528152506104e4565b600554811415610afc57600080fd5b600555565b60408051808201909152600981527f646576656c6f7065720000000000000000000000000000000000000000000000602082015281565b610b6533604080519081016040528060098152602001600080516020610f028339815191528152506104e4565b600160a060020a0381161515610b7a57600080fd5b610848816040805190810160405280600781526020017f61756469746f7200000000000000000000000000000000000000000000000000815250610c3f565b6040805180820190915260098152600080516020610f02833981519152602082015281565b6002546000908210610bef57600080fd5b50600090815260036020526040902054600160a060020a031690565b610c158282610c20565b151561054e57600080fd5b600160a060020a03166000908152602091909152604090205460ff1690565b610ca9826000836040518082805190602001908083835b60208310610c755780518252601f199092019160209182019101610c56565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610e22565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610d12578181015183820152602001610cfa565b50505050905090810190601f168015610d3f5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b610dba826000836040518082805190602001908083835b60208310610d865780518252601f199092019160209182019101610d67565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610e47565b81600160a060020a03167fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a8260405180806020018281038252838181518152602001915080519060200190808383600083811015610d12578181015183820152602001610cfa565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b600160a060020a0316600090815260209190915260409020805460ff19169055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610eaa57805160ff1916838001178555610ed7565b82800160010185558215610ed7579182015b82811115610ed7578251825591602001919060010190610ebc565b50610ee3929150610ee7565b5090565b610a8e91905b80821115610ee35760008155600101610eed56007375706572757365720000000000000000000000000000000000000000000000a165627a7a723058205e39ef63654049512619beaaef7bc4ca0a70bac642c6dfbc06e286e0d4e0a6ba0029`

// DeployBREMFactory deploys a new Ethereum contract, binding an instance of BREMFactory to it.
func DeployBREMFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BREMFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(BREMFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BREMFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BREMFactory{BREMFactoryCaller: BREMFactoryCaller{contract: contract}, BREMFactoryTransactor: BREMFactoryTransactor{contract: contract}, BREMFactoryFilterer: BREMFactoryFilterer{contract: contract}}, nil
}

// BREMFactory is an auto generated Go binding around an Ethereum contract.
type BREMFactory struct {
	BREMFactoryCaller     // Read-only binding to the contract
	BREMFactoryTransactor // Write-only binding to the contract
	BREMFactoryFilterer   // Log filterer for contract events
}

// BREMFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BREMFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BREMFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BREMFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BREMFactorySession struct {
	Contract     *BREMFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BREMFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BREMFactoryCallerSession struct {
	Contract *BREMFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BREMFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BREMFactoryTransactorSession struct {
	Contract     *BREMFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BREMFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BREMFactoryRaw struct {
	Contract *BREMFactory // Generic contract binding to access the raw methods on
}

// BREMFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BREMFactoryCallerRaw struct {
	Contract *BREMFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BREMFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BREMFactoryTransactorRaw struct {
	Contract *BREMFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBREMFactory creates a new instance of BREMFactory, bound to a specific deployed contract.
func NewBREMFactory(address common.Address, backend bind.ContractBackend) (*BREMFactory, error) {
	contract, err := bindBREMFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BREMFactory{BREMFactoryCaller: BREMFactoryCaller{contract: contract}, BREMFactoryTransactor: BREMFactoryTransactor{contract: contract}, BREMFactoryFilterer: BREMFactoryFilterer{contract: contract}}, nil
}

// NewBREMFactoryCaller creates a new read-only instance of BREMFactory, bound to a specific deployed contract.
func NewBREMFactoryCaller(address common.Address, caller bind.ContractCaller) (*BREMFactoryCaller, error) {
	contract, err := bindBREMFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BREMFactoryCaller{contract: contract}, nil
}

// NewBREMFactoryTransactor creates a new write-only instance of BREMFactory, bound to a specific deployed contract.
func NewBREMFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BREMFactoryTransactor, error) {
	contract, err := bindBREMFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BREMFactoryTransactor{contract: contract}, nil
}

// NewBREMFactoryFilterer creates a new log filterer instance of BREMFactory, bound to a specific deployed contract.
func NewBREMFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BREMFactoryFilterer, error) {
	contract, err := bindBREMFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BREMFactoryFilterer{contract: contract}, nil
}

// bindBREMFactory binds a generic wrapper to an already deployed contract.
func bindBREMFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BREMFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BREMFactory *BREMFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BREMFactory.Contract.BREMFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BREMFactory *BREMFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMFactory.Contract.BREMFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BREMFactory *BREMFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BREMFactory.Contract.BREMFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BREMFactory *BREMFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BREMFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BREMFactory *BREMFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BREMFactory *BREMFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BREMFactory.Contract.contract.Transact(opts, method, params...)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_BREMFactory *BREMFactoryCaller) ROLEAUDITOR(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "ROLE_AUDITOR")
	return *ret0, err
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_BREMFactory *BREMFactorySession) ROLEAUDITOR() (string, error) {
	return _BREMFactory.Contract.ROLEAUDITOR(&_BREMFactory.CallOpts)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_BREMFactory *BREMFactoryCallerSession) ROLEAUDITOR() (string, error) {
	return _BREMFactory.Contract.ROLEAUDITOR(&_BREMFactory.CallOpts)
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_BREMFactory *BREMFactoryCaller) ROLEDEVELOPER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "ROLE_DEVELOPER")
	return *ret0, err
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_BREMFactory *BREMFactorySession) ROLEDEVELOPER() (string, error) {
	return _BREMFactory.Contract.ROLEDEVELOPER(&_BREMFactory.CallOpts)
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_BREMFactory *BREMFactoryCallerSession) ROLEDEVELOPER() (string, error) {
	return _BREMFactory.Contract.ROLEDEVELOPER(&_BREMFactory.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_BREMFactory *BREMFactoryCaller) ROLESUPERUSER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "ROLE_SUPERUSER")
	return *ret0, err
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_BREMFactory *BREMFactorySession) ROLESUPERUSER() (string, error) {
	return _BREMFactory.Contract.ROLESUPERUSER(&_BREMFactory.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_BREMFactory *BREMFactoryCallerSession) ROLESUPERUSER() (string, error) {
	return _BREMFactory.Contract.ROLESUPERUSER(&_BREMFactory.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_BREMFactory *BREMFactoryCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _BREMFactory.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_BREMFactory *BREMFactorySession) CheckRole(_operator common.Address, _role string) error {
	return _BREMFactory.Contract.CheckRole(&_BREMFactory.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_BREMFactory *BREMFactoryCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _BREMFactory.Contract.CheckRole(&_BREMFactory.CallOpts, _operator, _role)
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(_index uint256) constant returns(address)
func (_BREMFactory *BREMFactoryCaller) GetProject(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "getProject", _index)
	return *ret0, err
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(_index uint256) constant returns(address)
func (_BREMFactory *BREMFactorySession) GetProject(_index *big.Int) (common.Address, error) {
	return _BREMFactory.Contract.GetProject(&_BREMFactory.CallOpts, _index)
}

// GetProject is a free data retrieval call binding the contract method 0xf0f3f2c8.
//
// Solidity: function getProject(_index uint256) constant returns(address)
func (_BREMFactory *BREMFactoryCallerSession) GetProject(_index *big.Int) (common.Address, error) {
	return _BREMFactory.Contract.GetProject(&_BREMFactory.CallOpts, _index)
}

// GetProjectByName is a free data retrieval call binding the contract method 0x4b77f269.
//
// Solidity: function getProjectByName(_projectName string) constant returns(address)
func (_BREMFactory *BREMFactoryCaller) GetProjectByName(opts *bind.CallOpts, _projectName string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "getProjectByName", _projectName)
	return *ret0, err
}

// GetProjectByName is a free data retrieval call binding the contract method 0x4b77f269.
//
// Solidity: function getProjectByName(_projectName string) constant returns(address)
func (_BREMFactory *BREMFactorySession) GetProjectByName(_projectName string) (common.Address, error) {
	return _BREMFactory.Contract.GetProjectByName(&_BREMFactory.CallOpts, _projectName)
}

// GetProjectByName is a free data retrieval call binding the contract method 0x4b77f269.
//
// Solidity: function getProjectByName(_projectName string) constant returns(address)
func (_BREMFactory *BREMFactoryCallerSession) GetProjectByName(_projectName string) (common.Address, error) {
	return _BREMFactory.Contract.GetProjectByName(&_BREMFactory.CallOpts, _projectName)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_BREMFactory *BREMFactoryCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_BREMFactory *BREMFactorySession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _BREMFactory.Contract.HasRole(&_BREMFactory.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_BREMFactory *BREMFactoryCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _BREMFactory.Contract.HasRole(&_BREMFactory.CallOpts, _operator, _role)
}

// IcoCreationPrice is a free data retrieval call binding the contract method 0x80d2370f.
//
// Solidity: function icoCreationPrice() constant returns(uint256)
func (_BREMFactory *BREMFactoryCaller) IcoCreationPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "icoCreationPrice")
	return *ret0, err
}

// IcoCreationPrice is a free data retrieval call binding the contract method 0x80d2370f.
//
// Solidity: function icoCreationPrice() constant returns(uint256)
func (_BREMFactory *BREMFactorySession) IcoCreationPrice() (*big.Int, error) {
	return _BREMFactory.Contract.IcoCreationPrice(&_BREMFactory.CallOpts)
}

// IcoCreationPrice is a free data retrieval call binding the contract method 0x80d2370f.
//
// Solidity: function icoCreationPrice() constant returns(uint256)
func (_BREMFactory *BREMFactoryCallerSession) IcoCreationPrice() (*big.Int, error) {
	return _BREMFactory.Contract.IcoCreationPrice(&_BREMFactory.CallOpts)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_BREMFactory *BREMFactoryCaller) IsAuditor(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "isAuditor", _addr)
	return *ret0, err
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_BREMFactory *BREMFactorySession) IsAuditor(_addr common.Address) (bool, error) {
	return _BREMFactory.Contract.IsAuditor(&_BREMFactory.CallOpts, _addr)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_BREMFactory *BREMFactoryCallerSession) IsAuditor(_addr common.Address) (bool, error) {
	return _BREMFactory.Contract.IsAuditor(&_BREMFactory.CallOpts, _addr)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_BREMFactory *BREMFactoryCaller) IsDeveloper(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "isDeveloper", _addr)
	return *ret0, err
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_BREMFactory *BREMFactorySession) IsDeveloper(_addr common.Address) (bool, error) {
	return _BREMFactory.Contract.IsDeveloper(&_BREMFactory.CallOpts, _addr)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_BREMFactory *BREMFactoryCallerSession) IsDeveloper(_addr common.Address) (bool, error) {
	return _BREMFactory.Contract.IsDeveloper(&_BREMFactory.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_BREMFactory *BREMFactoryCaller) IsSuperuser(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "isSuperuser", _addr)
	return *ret0, err
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_BREMFactory *BREMFactorySession) IsSuperuser(_addr common.Address) (bool, error) {
	return _BREMFactory.Contract.IsSuperuser(&_BREMFactory.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_BREMFactory *BREMFactoryCallerSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _BREMFactory.Contract.IsSuperuser(&_BREMFactory.CallOpts, _addr)
}

// Login is a free data retrieval call binding the contract method 0xb34e97e8.
//
// Solidity: function login() constant returns(string)
func (_BREMFactory *BREMFactoryCaller) Login(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "login")
	return *ret0, err
}

// Login is a free data retrieval call binding the contract method 0xb34e97e8.
//
// Solidity: function login() constant returns(string)
func (_BREMFactory *BREMFactorySession) Login() (string, error) {
	return _BREMFactory.Contract.Login(&_BREMFactory.CallOpts)
}

// Login is a free data retrieval call binding the contract method 0xb34e97e8.
//
// Solidity: function login() constant returns(string)
func (_BREMFactory *BREMFactoryCallerSession) Login() (string, error) {
	return _BREMFactory.Contract.Login(&_BREMFactory.CallOpts)
}

// ProjectsAmount is a free data retrieval call binding the contract method 0x9cdd9913.
//
// Solidity: function projectsAmount() constant returns(uint256)
func (_BREMFactory *BREMFactoryCaller) ProjectsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "projectsAmount")
	return *ret0, err
}

// ProjectsAmount is a free data retrieval call binding the contract method 0x9cdd9913.
//
// Solidity: function projectsAmount() constant returns(uint256)
func (_BREMFactory *BREMFactorySession) ProjectsAmount() (*big.Int, error) {
	return _BREMFactory.Contract.ProjectsAmount(&_BREMFactory.CallOpts)
}

// ProjectsAmount is a free data retrieval call binding the contract method 0x9cdd9913.
//
// Solidity: function projectsAmount() constant returns(uint256)
func (_BREMFactory *BREMFactoryCallerSession) ProjectsAmount() (*big.Int, error) {
	return _BREMFactory.Contract.ProjectsAmount(&_BREMFactory.CallOpts)
}

// WithdrawFeePercent is a free data retrieval call binding the contract method 0x495ef705.
//
// Solidity: function withdrawFeePercent() constant returns(uint256)
func (_BREMFactory *BREMFactoryCaller) WithdrawFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMFactory.contract.Call(opts, out, "withdrawFeePercent")
	return *ret0, err
}

// WithdrawFeePercent is a free data retrieval call binding the contract method 0x495ef705.
//
// Solidity: function withdrawFeePercent() constant returns(uint256)
func (_BREMFactory *BREMFactorySession) WithdrawFeePercent() (*big.Int, error) {
	return _BREMFactory.Contract.WithdrawFeePercent(&_BREMFactory.CallOpts)
}

// WithdrawFeePercent is a free data retrieval call binding the contract method 0x495ef705.
//
// Solidity: function withdrawFeePercent() constant returns(uint256)
func (_BREMFactory *BREMFactoryCallerSession) WithdrawFeePercent() (*big.Int, error) {
	return _BREMFactory.Contract.WithdrawFeePercent(&_BREMFactory.CallOpts)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_BREMFactory *BREMFactoryTransactor) AddAuditor(opts *bind.TransactOpts, _newAuditor common.Address) (*types.Transaction, error) {
	return _BREMFactory.contract.Transact(opts, "addAuditor", _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_BREMFactory *BREMFactorySession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _BREMFactory.Contract.AddAuditor(&_BREMFactory.TransactOpts, _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_BREMFactory *BREMFactoryTransactorSession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _BREMFactory.Contract.AddAuditor(&_BREMFactory.TransactOpts, _newAuditor)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_BREMFactory *BREMFactoryTransactor) AddDeveloper(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMFactory.contract.Transact(opts, "addDeveloper")
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_BREMFactory *BREMFactorySession) AddDeveloper() (*types.Transaction, error) {
	return _BREMFactory.Contract.AddDeveloper(&_BREMFactory.TransactOpts)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_BREMFactory *BREMFactoryTransactorSession) AddDeveloper() (*types.Transaction, error) {
	return _BREMFactory.Contract.AddDeveloper(&_BREMFactory.TransactOpts)
}

// SetIcoCreationPrice is a paid mutator transaction binding the contract method 0xc4f91810.
//
// Solidity: function setIcoCreationPrice(_icoCreationPrice uint256) returns()
func (_BREMFactory *BREMFactoryTransactor) SetIcoCreationPrice(opts *bind.TransactOpts, _icoCreationPrice *big.Int) (*types.Transaction, error) {
	return _BREMFactory.contract.Transact(opts, "setIcoCreationPrice", _icoCreationPrice)
}

// SetIcoCreationPrice is a paid mutator transaction binding the contract method 0xc4f91810.
//
// Solidity: function setIcoCreationPrice(_icoCreationPrice uint256) returns()
func (_BREMFactory *BREMFactorySession) SetIcoCreationPrice(_icoCreationPrice *big.Int) (*types.Transaction, error) {
	return _BREMFactory.Contract.SetIcoCreationPrice(&_BREMFactory.TransactOpts, _icoCreationPrice)
}

// SetIcoCreationPrice is a paid mutator transaction binding the contract method 0xc4f91810.
//
// Solidity: function setIcoCreationPrice(_icoCreationPrice uint256) returns()
func (_BREMFactory *BREMFactoryTransactorSession) SetIcoCreationPrice(_icoCreationPrice *big.Int) (*types.Transaction, error) {
	return _BREMFactory.Contract.SetIcoCreationPrice(&_BREMFactory.TransactOpts, _icoCreationPrice)
}

// SetWithdrawFeePercent is a paid mutator transaction binding the contract method 0x447cb896.
//
// Solidity: function setWithdrawFeePercent(_withdrawFeePercent uint256) returns()
func (_BREMFactory *BREMFactoryTransactor) SetWithdrawFeePercent(opts *bind.TransactOpts, _withdrawFeePercent *big.Int) (*types.Transaction, error) {
	return _BREMFactory.contract.Transact(opts, "setWithdrawFeePercent", _withdrawFeePercent)
}

// SetWithdrawFeePercent is a paid mutator transaction binding the contract method 0x447cb896.
//
// Solidity: function setWithdrawFeePercent(_withdrawFeePercent uint256) returns()
func (_BREMFactory *BREMFactorySession) SetWithdrawFeePercent(_withdrawFeePercent *big.Int) (*types.Transaction, error) {
	return _BREMFactory.Contract.SetWithdrawFeePercent(&_BREMFactory.TransactOpts, _withdrawFeePercent)
}

// SetWithdrawFeePercent is a paid mutator transaction binding the contract method 0x447cb896.
//
// Solidity: function setWithdrawFeePercent(_withdrawFeePercent uint256) returns()
func (_BREMFactory *BREMFactoryTransactorSession) SetWithdrawFeePercent(_withdrawFeePercent *big.Int) (*types.Transaction, error) {
	return _BREMFactory.Contract.SetWithdrawFeePercent(&_BREMFactory.TransactOpts, _withdrawFeePercent)
}

// SignUp is a paid mutator transaction binding the contract method 0xadcfb4fc.
//
// Solidity: function signUp(_name string) returns(string)
func (_BREMFactory *BREMFactoryTransactor) SignUp(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _BREMFactory.contract.Transact(opts, "signUp", _name)
}

// SignUp is a paid mutator transaction binding the contract method 0xadcfb4fc.
//
// Solidity: function signUp(_name string) returns(string)
func (_BREMFactory *BREMFactorySession) SignUp(_name string) (*types.Transaction, error) {
	return _BREMFactory.Contract.SignUp(&_BREMFactory.TransactOpts, _name)
}

// SignUp is a paid mutator transaction binding the contract method 0xadcfb4fc.
//
// Solidity: function signUp(_name string) returns(string)
func (_BREMFactory *BREMFactoryTransactorSession) SignUp(_name string) (*types.Transaction, error) {
	return _BREMFactory.Contract.SignUp(&_BREMFactory.TransactOpts, _name)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_BREMFactory *BREMFactoryTransactor) TransferSuperuser(opts *bind.TransactOpts, _newSuperuser common.Address) (*types.Transaction, error) {
	return _BREMFactory.contract.Transact(opts, "transferSuperuser", _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_BREMFactory *BREMFactorySession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _BREMFactory.Contract.TransferSuperuser(&_BREMFactory.TransactOpts, _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_BREMFactory *BREMFactoryTransactorSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _BREMFactory.Contract.TransferSuperuser(&_BREMFactory.TransactOpts, _newSuperuser)
}

// BREMFactoryBREMICOCreatedIterator is returned from FilterBREMICOCreated and is used to iterate over the raw logs and unpacked data for BREMICOCreated events raised by the BREMFactory contract.
type BREMFactoryBREMICOCreatedIterator struct {
	Event *BREMFactoryBREMICOCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMFactoryBREMICOCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMFactoryBREMICOCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMFactoryBREMICOCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMFactoryBREMICOCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMFactoryBREMICOCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMFactoryBREMICOCreated represents a BREMICOCreated event raised by the BREMFactory contract.
type BREMFactoryBREMICOCreated struct {
	Creator      common.Address
	IcoAddress   common.Address
	TokenAddress common.Address
	Name         string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBREMICOCreated is a free log retrieval operation binding the contract event 0xcbef3bd4943085dfb59f0609e5f754aec11e22a5838e77e0d0c0bff98a0dcdae.
//
// Solidity: e BREMICOCreated(creator indexed address, icoAddress indexed address, tokenAddress indexed address, name string)
func (_BREMFactory *BREMFactoryFilterer) FilterBREMICOCreated(opts *bind.FilterOpts, creator []common.Address, icoAddress []common.Address, tokenAddress []common.Address) (*BREMFactoryBREMICOCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var icoAddressRule []interface{}
	for _, icoAddressItem := range icoAddress {
		icoAddressRule = append(icoAddressRule, icoAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BREMFactory.contract.FilterLogs(opts, "BREMICOCreated", creatorRule, icoAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BREMFactoryBREMICOCreatedIterator{contract: _BREMFactory.contract, event: "BREMICOCreated", logs: logs, sub: sub}, nil
}

// WatchBREMICOCreated is a free log subscription operation binding the contract event 0xcbef3bd4943085dfb59f0609e5f754aec11e22a5838e77e0d0c0bff98a0dcdae.
//
// Solidity: e BREMICOCreated(creator indexed address, icoAddress indexed address, tokenAddress indexed address, name string)
func (_BREMFactory *BREMFactoryFilterer) WatchBREMICOCreated(opts *bind.WatchOpts, sink chan<- *BREMFactoryBREMICOCreated, creator []common.Address, icoAddress []common.Address, tokenAddress []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var icoAddressRule []interface{}
	for _, icoAddressItem := range icoAddress {
		icoAddressRule = append(icoAddressRule, icoAddressItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _BREMFactory.contract.WatchLogs(opts, "BREMICOCreated", creatorRule, icoAddressRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMFactoryBREMICOCreated)
				if err := _BREMFactory.contract.UnpackLog(event, "BREMICOCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMFactoryRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the BREMFactory contract.
type BREMFactoryRoleAddedIterator struct {
	Event *BREMFactoryRoleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMFactoryRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMFactoryRoleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMFactoryRoleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMFactoryRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMFactoryRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMFactoryRoleAdded represents a RoleAdded event raised by the BREMFactory contract.
type BREMFactoryRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_BREMFactory *BREMFactoryFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*BREMFactoryRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BREMFactory.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BREMFactoryRoleAddedIterator{contract: _BREMFactory.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_BREMFactory *BREMFactoryFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *BREMFactoryRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BREMFactory.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMFactoryRoleAdded)
				if err := _BREMFactory.contract.UnpackLog(event, "RoleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMFactoryRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the BREMFactory contract.
type BREMFactoryRoleRemovedIterator struct {
	Event *BREMFactoryRoleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMFactoryRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMFactoryRoleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMFactoryRoleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMFactoryRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMFactoryRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMFactoryRoleRemoved represents a RoleRemoved event raised by the BREMFactory contract.
type BREMFactoryRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_BREMFactory *BREMFactoryFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*BREMFactoryRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BREMFactory.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &BREMFactoryRoleRemovedIterator{contract: _BREMFactory.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_BREMFactory *BREMFactoryFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *BREMFactoryRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BREMFactory.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMFactoryRoleRemoved)
				if err := _BREMFactory.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMICOABI is the input ABI used to generate the binding from.
const BREMICOABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_auditor\",\"type\":\"address\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"request\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"confirmAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"docHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWithdrawn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balancesInToken\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"withdrawFeePercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_auditor\",\"type\":\"address\"}],\"name\":\"isAuditor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"confirmWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getAuditor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRequested\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"auditorsAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_auditor\",\"type\":\"address\"}],\"name\":\"addAuditor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"auditSelected\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishAuditorSelection\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brem\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_brem\",\"type\":\"address\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_cap\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_closingTime\",\"type\":\"uint256\"},{\"name\":\"_description\",\"type\":\"string\"},{\"name\":\"_docHash\",\"type\":\"string\"},{\"name\":\"_auditAddress\",\"type\":\"address\"},{\"name\":\"_withdrawFeePercent\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// BREMICOBin is the compiled bytecode used for deploying new contracts.
const BREMICOBin = `0x60806040523480156200001157600080fd5b50604051620013cd380380620013cd83398101604090815281516020830151918301516060840151608085015160a086015160c087015160e08801516101008901516101208a01516101408b0151989a998a019997989697959694959394928401939190910191600160a060020a038b1615156200008e57600080fd5b89516000106200009d57600080fd5b60008911620000ab57600080fd5b60008811620000b957600080fd5b600160a060020a0387161515620000cf57600080fd5b600160a060020a0386161515620000e557600080fd5b600160a060020a0382161515620000fb57600080fd5b428510156200010957600080fd5b600081101580156200011c575060648111155b15156200012857600080fd5b89516200013d9060019060208d0190620001ea565b5060148054600160a060020a03808e16600160a060020a03199283161790925560058b905560038a9055600280548a84169083161790556000805492891692909116919091179055601085905583516200019f906006906020870190620001ea565b508251620001b5906007906020860190620001ea565b5060088054600160a060020a031916600160a060020a039390931692909217909155601355506200028f975050505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022d57805160ff19168380011785556200025d565b828001600101855582156200025d579182015b828111156200025d57825182559160200191906001019062000240565b506200026b9291506200026f565b5090565b6200028c91905b808211156200026b576000815560010162000276565b90565b61112e806200029f6000396000f3006080604052600436106101745763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461017f5780631515bc2b1461020957806318e4ac351461023257806327e235e3146102535780632c4e722e146102865780632e1a7d4d1461029b578063338cdca1146102a6578063355274ea146102d45780633d6df0d5146102e95780634042b66f146102fe57806344aa57001461031357806345e4f34a14610328578063495ef7051461034957806349b905571461035e5780634b6753bc1461037f5780634f93594514610394578063521eb273146103a9578063590e1ae3146103da57806368124f9a146103ef5780637284e4161461040457806393f80ff41461041957806397781f0214610431578063c8e7f59514610446578063e429cef11461045b578063e44e76a01461047c578063ea78d1ab14610491578063ec8ac4d8146104a6578063fad5e0c4146104ba578063fc0c546a146104cf575b61017d336104e4565b005b34801561018b57600080fd5b50610194610598565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101ce5781810151838201526020016101b6565b50505050905090810190601f1680156101fb5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561021557600080fd5b5061021e610625565b604080519115158252519081900360200190f35b34801561023e57600080fd5b5061021e600160a060020a036004351661062d565b34801561025f57600080fd5b50610274600160a060020a036004351661064b565b60408051918252519081900360200190f35b34801561029257600080fd5b5061027461065d565b61017d600435610663565b3480156102b257600080fd5b506102bb61072b565b6040805192835260208301919091528051918290030190f35b3480156102e057600080fd5b50610274610734565b3480156102f557600080fd5b5061019461073a565b34801561030a57600080fd5b50610274610795565b34801561031f57600080fd5b5061021e61079b565b34801561033457600080fd5b50610274600160a060020a03600435166107c6565b34801561035557600080fd5b506102746107d8565b34801561036a57600080fd5b5061021e600160a060020a03600435166107de565b34801561038b57600080fd5b506102746107fc565b3480156103a057600080fd5b5061021e610802565b3480156103b557600080fd5b506103be61080d565b60408051600160a060020a039092168252519081900360200190f35b3480156103e657600080fd5b5061017d61081c565b3480156103fb57600080fd5b5061017d61094c565b34801561041057600080fd5b50610194610bba565b34801561042557600080fd5b506103be600435610c15565b34801561043d57600080fd5b5061021e610c42565b34801561045257600080fd5b50610274610c7e565b34801561046757600080fd5b5061017d600160a060020a0360043516610c84565b34801561048857600080fd5b5061021e610e60565b34801561049d57600080fd5b5061017d610e69565b61017d600160a060020a03600435166104e4565b3480156104c657600080fd5b506103be610f43565b3480156104db57600080fd5b506103be610f52565b600c54600090819060ff1680156104fd57506010544211155b151561050857600080fd5b3491506105158383610f61565b61051e82610f86565b600454909150610534908363ffffffff610fa316565b6004556105418382610fbd565b60408051838152602081018390528151600160a060020a0386169233927f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18929081900390910190a36105938383610fc7565b505050565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561061d5780601f106105f25761010080835404028352916020019161061d565b820191906000526020600020905b81548152906001019060200180831161060057829003601f168201915b505050505081565b601054421190565b600160a060020a03166000908152600f602052604090205460ff1690565b60116020526000908152604090205481565b60035481565b600254600160a060020a0316331461067a57600080fd5b600c5460ff16151561068b57600080fd5b610693610802565b80156106a257506106a2610625565b15156106ad57600080fd5b606481101580156106bf575030318111155b15156106ca57600080fd5b600d54156106d757600080fd5b6064303182900310156107085760408051808201909152303180825260006020909201829052600d55600e55610728565b6040805180820190915281815260006020909101819052600d829055600e555b50565b600d54600e5482565b60055481565b6007805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561061d5780601f106105f25761010080835404028352916020019161061d565b60045481565b60006107a5610625565b80156107b457506107b4610802565b80156107c1575060643031105b905090565b60126020526000908152604090205481565b60135481565b600160a060020a031660009081526009602052604090205460ff1690565b60105481565b600554600454101590565b600254600160a060020a031681565b600080610827610625565b80156108385750610836610802565b155b151561084357600080fd5b336000908152601160205260408120541161085d57600080fd5b505033600081815260116020908152604080832080546012909352818420805491859055849055835482517fd0482d04000000000000000000000000000000000000000000000000000000008152600481019690965260248601829052915192949093600160a060020a039092169263d0482d049260448084019391929182900301818387803b1580156108f057600080fd5b505af1158015610904573d6000803e3d6000fd5b505060045461091c925090508363ffffffff61101216565b600455604051339083156108fc029084906000818181858888f19350505050158015610593573d6000803e3d6000fd5b600854604080517f49b905570000000000000000000000000000000000000000000000000000000081523360048201529051600092839283928392600160a060020a0316916349b9055791602480830192602092919082900301818787803b1580156109b757600080fd5b505af11580156109cb573d6000803e3d6000fd5b505050506040513d60208110156109e157600080fd5b505115156109ee57600080fd5b600d546000106109fd57600080fd5b336000908152600f602052604090205460ff1615610a1a57600080fd5b600c5460ff161515610a2b57600080fd5b610a33610802565b8015610a425750610a42610625565b1515610a4d57600080fd5b3360009081526009602052604090205460ff161515610a6b57600080fd5b336000908152600f60205260409020805460ff19166001908117909155600e805490910190819055600b541415610bb457600d8054604080518082019091526000808252602090910181905291829055600e829055945092505b600b54831015610b06576000838152600a6020908152604080832054600160a060020a03168352600f9091529020805460ff19169055600190920191610ac5565b601354610b2a90610b1e86606463ffffffff61102416565b9063ffffffff61103b16565b9150610b3c848363ffffffff61101216565b601454604051919250600160a060020a03169083156108fc029084906000818181858888f19350505050158015610b77573d6000803e3d6000fd5b50600254604051600160a060020a039091169082156108fc029083906000818181858888f19350505050158015610bb2573d6000803e3d6000fd5b505b50505050565b6006805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561061d5780601f106105f25761010080835404028352916020019161061d565b600b546000908210610c2657600080fd5b506000908152600a6020526040902054600160a060020a031690565b6000610c4c610802565b8015610c5b5750610c5b610625565b8015610c695750600d546000105b80156107c15750610c7861079b565b15905090565b600b5481565b601054421115610c9357600080fd5b600c5460ff1615610ca357600080fd5b600854604080517fbceee05e0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163bceee05e916024808201926020929091908290030181600087803b158015610d0957600080fd5b505af1158015610d1d573d6000803e3d6000fd5b505050506040513d6020811015610d3357600080fd5b50511515610d4057600080fd5b600854604080517f49b90557000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915191909216916349b905579160248083019260209291908290030181600087803b158015610da857600080fd5b505af1158015610dbc573d6000803e3d6000fd5b505050506040513d6020811015610dd257600080fd5b50511515610ddf57600080fd5b600160a060020a03811660009081526009602052604090205460ff1615610e0557600080fd5b600160a060020a03166000818152600960209081526040808320805460ff19166001908117909155600b80548552600a9093529220805473ffffffffffffffffffffffffffffffffffffffff19169093179092558154019055565b600c5460ff1681565b601054421115610e7857600080fd5b600c5460ff1615610e8857600080fd5b600854604080517fbceee05e0000000000000000000000000000000000000000000000000000000081523360048201529051600160a060020a039092169163bceee05e916024808201926020929091908290030181600087803b158015610eee57600080fd5b505af1158015610f02573d6000803e3d6000fd5b505050506040513d6020811015610f1857600080fd5b50511515610f2557600080fd5b600b54600010610f3457600080fd5b600c805460ff19166001179055565b601454600160a060020a031681565b600054600160a060020a031681565b600160a060020a0382161515610f7657600080fd5b801515610f8257600080fd5b5050565b6000610f9d6003548361103b90919063ffffffff16565b92915050565b600082820183811015610fb257fe5b8091505b5092915050565b610f828282611066565b600160a060020a0382166000908152601160205260409020805482019055610fee81610f86565b600160a060020a039092166000908152601260205260409020805490920190915550565b60008282111561101e57fe5b50900390565b600080828481151561103257fe5b04949350505050565b60008083151561104e5760009150610fb6565b5082820282848281151561105e57fe5b0414610fb257fe5b60008054604080517f40c10f19000000000000000000000000000000000000000000000000000000008152600160a060020a03868116600483015260248201869052915191909216926340c10f1992604480820193602093909283900390910190829087803b1580156110d857600080fd5b505af11580156110ec573d6000803e3d6000fd5b505050506040513d6020811015610bb457600080fd00a165627a7a72305820670adf9b43ff56ea291ac7240a801a880392912a6ad5023622a0ce670320d2790029`

// DeployBREMICO deploys a new Ethereum contract, binding an instance of BREMICO to it.
func DeployBREMICO(auth *bind.TransactOpts, backend bind.ContractBackend, _brem common.Address, _name string, _cap *big.Int, _rate *big.Int, _wallet common.Address, _token common.Address, _closingTime *big.Int, _description string, _docHash string, _auditAddress common.Address, _withdrawFeePercent *big.Int) (common.Address, *types.Transaction, *BREMICO, error) {
	parsed, err := abi.JSON(strings.NewReader(BREMICOABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BREMICOBin), backend, _brem, _name, _cap, _rate, _wallet, _token, _closingTime, _description, _docHash, _auditAddress, _withdrawFeePercent)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BREMICO{BREMICOCaller: BREMICOCaller{contract: contract}, BREMICOTransactor: BREMICOTransactor{contract: contract}, BREMICOFilterer: BREMICOFilterer{contract: contract}}, nil
}

// BREMICO is an auto generated Go binding around an Ethereum contract.
type BREMICO struct {
	BREMICOCaller     // Read-only binding to the contract
	BREMICOTransactor // Write-only binding to the contract
	BREMICOFilterer   // Log filterer for contract events
}

// BREMICOCaller is an auto generated read-only Go binding around an Ethereum contract.
type BREMICOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMICOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BREMICOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMICOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BREMICOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMICOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BREMICOSession struct {
	Contract     *BREMICO          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BREMICOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BREMICOCallerSession struct {
	Contract *BREMICOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BREMICOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BREMICOTransactorSession struct {
	Contract     *BREMICOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BREMICORaw is an auto generated low-level Go binding around an Ethereum contract.
type BREMICORaw struct {
	Contract *BREMICO // Generic contract binding to access the raw methods on
}

// BREMICOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BREMICOCallerRaw struct {
	Contract *BREMICOCaller // Generic read-only contract binding to access the raw methods on
}

// BREMICOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BREMICOTransactorRaw struct {
	Contract *BREMICOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBREMICO creates a new instance of BREMICO, bound to a specific deployed contract.
func NewBREMICO(address common.Address, backend bind.ContractBackend) (*BREMICO, error) {
	contract, err := bindBREMICO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BREMICO{BREMICOCaller: BREMICOCaller{contract: contract}, BREMICOTransactor: BREMICOTransactor{contract: contract}, BREMICOFilterer: BREMICOFilterer{contract: contract}}, nil
}

// NewBREMICOCaller creates a new read-only instance of BREMICO, bound to a specific deployed contract.
func NewBREMICOCaller(address common.Address, caller bind.ContractCaller) (*BREMICOCaller, error) {
	contract, err := bindBREMICO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BREMICOCaller{contract: contract}, nil
}

// NewBREMICOTransactor creates a new write-only instance of BREMICO, bound to a specific deployed contract.
func NewBREMICOTransactor(address common.Address, transactor bind.ContractTransactor) (*BREMICOTransactor, error) {
	contract, err := bindBREMICO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BREMICOTransactor{contract: contract}, nil
}

// NewBREMICOFilterer creates a new log filterer instance of BREMICO, bound to a specific deployed contract.
func NewBREMICOFilterer(address common.Address, filterer bind.ContractFilterer) (*BREMICOFilterer, error) {
	contract, err := bindBREMICO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BREMICOFilterer{contract: contract}, nil
}

// bindBREMICO binds a generic wrapper to an already deployed contract.
func bindBREMICO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BREMICOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BREMICO *BREMICORaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BREMICO.Contract.BREMICOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BREMICO *BREMICORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMICO.Contract.BREMICOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BREMICO *BREMICORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BREMICO.Contract.BREMICOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BREMICO *BREMICOCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BREMICO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BREMICO *BREMICOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMICO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BREMICO *BREMICOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BREMICO.Contract.contract.Transact(opts, method, params...)
}

// AuditSelected is a free data retrieval call binding the contract method 0xe44e76a0.
//
// Solidity: function auditSelected() constant returns(bool)
func (_BREMICO *BREMICOCaller) AuditSelected(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "auditSelected")
	return *ret0, err
}

// AuditSelected is a free data retrieval call binding the contract method 0xe44e76a0.
//
// Solidity: function auditSelected() constant returns(bool)
func (_BREMICO *BREMICOSession) AuditSelected() (bool, error) {
	return _BREMICO.Contract.AuditSelected(&_BREMICO.CallOpts)
}

// AuditSelected is a free data retrieval call binding the contract method 0xe44e76a0.
//
// Solidity: function auditSelected() constant returns(bool)
func (_BREMICO *BREMICOCallerSession) AuditSelected() (bool, error) {
	return _BREMICO.Contract.AuditSelected(&_BREMICO.CallOpts)
}

// AuditorsAmount is a free data retrieval call binding the contract method 0xc8e7f595.
//
// Solidity: function auditorsAmount() constant returns(uint256)
func (_BREMICO *BREMICOCaller) AuditorsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "auditorsAmount")
	return *ret0, err
}

// AuditorsAmount is a free data retrieval call binding the contract method 0xc8e7f595.
//
// Solidity: function auditorsAmount() constant returns(uint256)
func (_BREMICO *BREMICOSession) AuditorsAmount() (*big.Int, error) {
	return _BREMICO.Contract.AuditorsAmount(&_BREMICO.CallOpts)
}

// AuditorsAmount is a free data retrieval call binding the contract method 0xc8e7f595.
//
// Solidity: function auditorsAmount() constant returns(uint256)
func (_BREMICO *BREMICOCallerSession) AuditorsAmount() (*big.Int, error) {
	return _BREMICO.Contract.AuditorsAmount(&_BREMICO.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_BREMICO *BREMICOCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_BREMICO *BREMICOSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _BREMICO.Contract.Balances(&_BREMICO.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_BREMICO *BREMICOCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _BREMICO.Contract.Balances(&_BREMICO.CallOpts, arg0)
}

// BalancesInToken is a free data retrieval call binding the contract method 0x45e4f34a.
//
// Solidity: function balancesInToken( address) constant returns(uint256)
func (_BREMICO *BREMICOCaller) BalancesInToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "balancesInToken", arg0)
	return *ret0, err
}

// BalancesInToken is a free data retrieval call binding the contract method 0x45e4f34a.
//
// Solidity: function balancesInToken( address) constant returns(uint256)
func (_BREMICO *BREMICOSession) BalancesInToken(arg0 common.Address) (*big.Int, error) {
	return _BREMICO.Contract.BalancesInToken(&_BREMICO.CallOpts, arg0)
}

// BalancesInToken is a free data retrieval call binding the contract method 0x45e4f34a.
//
// Solidity: function balancesInToken( address) constant returns(uint256)
func (_BREMICO *BREMICOCallerSession) BalancesInToken(arg0 common.Address) (*big.Int, error) {
	return _BREMICO.Contract.BalancesInToken(&_BREMICO.CallOpts, arg0)
}

// Brem is a free data retrieval call binding the contract method 0xfad5e0c4.
//
// Solidity: function brem() constant returns(address)
func (_BREMICO *BREMICOCaller) Brem(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "brem")
	return *ret0, err
}

// Brem is a free data retrieval call binding the contract method 0xfad5e0c4.
//
// Solidity: function brem() constant returns(address)
func (_BREMICO *BREMICOSession) Brem() (common.Address, error) {
	return _BREMICO.Contract.Brem(&_BREMICO.CallOpts)
}

// Brem is a free data retrieval call binding the contract method 0xfad5e0c4.
//
// Solidity: function brem() constant returns(address)
func (_BREMICO *BREMICOCallerSession) Brem() (common.Address, error) {
	return _BREMICO.Contract.Brem(&_BREMICO.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_BREMICO *BREMICOCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_BREMICO *BREMICOSession) Cap() (*big.Int, error) {
	return _BREMICO.Contract.Cap(&_BREMICO.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_BREMICO *BREMICOCallerSession) Cap() (*big.Int, error) {
	return _BREMICO.Contract.Cap(&_BREMICO.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_BREMICO *BREMICOCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "capReached")
	return *ret0, err
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_BREMICO *BREMICOSession) CapReached() (bool, error) {
	return _BREMICO.Contract.CapReached(&_BREMICO.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() constant returns(bool)
func (_BREMICO *BREMICOCallerSession) CapReached() (bool, error) {
	return _BREMICO.Contract.CapReached(&_BREMICO.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_BREMICO *BREMICOCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "closingTime")
	return *ret0, err
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_BREMICO *BREMICOSession) ClosingTime() (*big.Int, error) {
	return _BREMICO.Contract.ClosingTime(&_BREMICO.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() constant returns(uint256)
func (_BREMICO *BREMICOCallerSession) ClosingTime() (*big.Int, error) {
	return _BREMICO.Contract.ClosingTime(&_BREMICO.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_BREMICO *BREMICOCaller) Description(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "description")
	return *ret0, err
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_BREMICO *BREMICOSession) Description() (string, error) {
	return _BREMICO.Contract.Description(&_BREMICO.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() constant returns(string)
func (_BREMICO *BREMICOCallerSession) Description() (string, error) {
	return _BREMICO.Contract.Description(&_BREMICO.CallOpts)
}

// DocHash is a free data retrieval call binding the contract method 0x3d6df0d5.
//
// Solidity: function docHash() constant returns(string)
func (_BREMICO *BREMICOCaller) DocHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "docHash")
	return *ret0, err
}

// DocHash is a free data retrieval call binding the contract method 0x3d6df0d5.
//
// Solidity: function docHash() constant returns(string)
func (_BREMICO *BREMICOSession) DocHash() (string, error) {
	return _BREMICO.Contract.DocHash(&_BREMICO.CallOpts)
}

// DocHash is a free data retrieval call binding the contract method 0x3d6df0d5.
//
// Solidity: function docHash() constant returns(string)
func (_BREMICO *BREMICOCallerSession) DocHash() (string, error) {
	return _BREMICO.Contract.DocHash(&_BREMICO.CallOpts)
}

// GetAuditor is a free data retrieval call binding the contract method 0x93f80ff4.
//
// Solidity: function getAuditor(_index uint256) constant returns(address)
func (_BREMICO *BREMICOCaller) GetAuditor(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "getAuditor", _index)
	return *ret0, err
}

// GetAuditor is a free data retrieval call binding the contract method 0x93f80ff4.
//
// Solidity: function getAuditor(_index uint256) constant returns(address)
func (_BREMICO *BREMICOSession) GetAuditor(_index *big.Int) (common.Address, error) {
	return _BREMICO.Contract.GetAuditor(&_BREMICO.CallOpts, _index)
}

// GetAuditor is a free data retrieval call binding the contract method 0x93f80ff4.
//
// Solidity: function getAuditor(_index uint256) constant returns(address)
func (_BREMICO *BREMICOCallerSession) GetAuditor(_index *big.Int) (common.Address, error) {
	return _BREMICO.Contract.GetAuditor(&_BREMICO.CallOpts, _index)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_BREMICO *BREMICOCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "hasClosed")
	return *ret0, err
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_BREMICO *BREMICOSession) HasClosed() (bool, error) {
	return _BREMICO.Contract.HasClosed(&_BREMICO.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() constant returns(bool)
func (_BREMICO *BREMICOCallerSession) HasClosed() (bool, error) {
	return _BREMICO.Contract.HasClosed(&_BREMICO.CallOpts)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_auditor address) constant returns(bool)
func (_BREMICO *BREMICOCaller) IsAuditor(opts *bind.CallOpts, _auditor common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "isAuditor", _auditor)
	return *ret0, err
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_auditor address) constant returns(bool)
func (_BREMICO *BREMICOSession) IsAuditor(_auditor common.Address) (bool, error) {
	return _BREMICO.Contract.IsAuditor(&_BREMICO.CallOpts, _auditor)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_auditor address) constant returns(bool)
func (_BREMICO *BREMICOCallerSession) IsAuditor(_auditor common.Address) (bool, error) {
	return _BREMICO.Contract.IsAuditor(&_BREMICO.CallOpts, _auditor)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x18e4ac35.
//
// Solidity: function isConfirmed(_auditor address) constant returns(bool)
func (_BREMICO *BREMICOCaller) IsConfirmed(opts *bind.CallOpts, _auditor common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "isConfirmed", _auditor)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x18e4ac35.
//
// Solidity: function isConfirmed(_auditor address) constant returns(bool)
func (_BREMICO *BREMICOSession) IsConfirmed(_auditor common.Address) (bool, error) {
	return _BREMICO.Contract.IsConfirmed(&_BREMICO.CallOpts, _auditor)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x18e4ac35.
//
// Solidity: function isConfirmed(_auditor address) constant returns(bool)
func (_BREMICO *BREMICOCallerSession) IsConfirmed(_auditor common.Address) (bool, error) {
	return _BREMICO.Contract.IsConfirmed(&_BREMICO.CallOpts, _auditor)
}

// IsRequested is a free data retrieval call binding the contract method 0x97781f02.
//
// Solidity: function isRequested() constant returns(bool)
func (_BREMICO *BREMICOCaller) IsRequested(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "isRequested")
	return *ret0, err
}

// IsRequested is a free data retrieval call binding the contract method 0x97781f02.
//
// Solidity: function isRequested() constant returns(bool)
func (_BREMICO *BREMICOSession) IsRequested() (bool, error) {
	return _BREMICO.Contract.IsRequested(&_BREMICO.CallOpts)
}

// IsRequested is a free data retrieval call binding the contract method 0x97781f02.
//
// Solidity: function isRequested() constant returns(bool)
func (_BREMICO *BREMICOCallerSession) IsRequested() (bool, error) {
	return _BREMICO.Contract.IsRequested(&_BREMICO.CallOpts)
}

// IsWithdrawn is a free data retrieval call binding the contract method 0x44aa5700.
//
// Solidity: function isWithdrawn() constant returns(bool)
func (_BREMICO *BREMICOCaller) IsWithdrawn(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "isWithdrawn")
	return *ret0, err
}

// IsWithdrawn is a free data retrieval call binding the contract method 0x44aa5700.
//
// Solidity: function isWithdrawn() constant returns(bool)
func (_BREMICO *BREMICOSession) IsWithdrawn() (bool, error) {
	return _BREMICO.Contract.IsWithdrawn(&_BREMICO.CallOpts)
}

// IsWithdrawn is a free data retrieval call binding the contract method 0x44aa5700.
//
// Solidity: function isWithdrawn() constant returns(bool)
func (_BREMICO *BREMICOCallerSession) IsWithdrawn() (bool, error) {
	return _BREMICO.Contract.IsWithdrawn(&_BREMICO.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BREMICO *BREMICOCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BREMICO *BREMICOSession) Name() (string, error) {
	return _BREMICO.Contract.Name(&_BREMICO.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BREMICO *BREMICOCallerSession) Name() (string, error) {
	return _BREMICO.Contract.Name(&_BREMICO.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_BREMICO *BREMICOCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_BREMICO *BREMICOSession) Rate() (*big.Int, error) {
	return _BREMICO.Contract.Rate(&_BREMICO.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_BREMICO *BREMICOCallerSession) Rate() (*big.Int, error) {
	return _BREMICO.Contract.Rate(&_BREMICO.CallOpts)
}

// Request is a free data retrieval call binding the contract method 0x338cdca1.
//
// Solidity: function request() constant returns(value uint256, confirmAmount uint256)
func (_BREMICO *BREMICOCaller) Request(opts *bind.CallOpts) (struct {
	Value         *big.Int
	ConfirmAmount *big.Int
}, error) {
	ret := new(struct {
		Value         *big.Int
		ConfirmAmount *big.Int
	})
	out := ret
	err := _BREMICO.contract.Call(opts, out, "request")
	return *ret, err
}

// Request is a free data retrieval call binding the contract method 0x338cdca1.
//
// Solidity: function request() constant returns(value uint256, confirmAmount uint256)
func (_BREMICO *BREMICOSession) Request() (struct {
	Value         *big.Int
	ConfirmAmount *big.Int
}, error) {
	return _BREMICO.Contract.Request(&_BREMICO.CallOpts)
}

// Request is a free data retrieval call binding the contract method 0x338cdca1.
//
// Solidity: function request() constant returns(value uint256, confirmAmount uint256)
func (_BREMICO *BREMICOCallerSession) Request() (struct {
	Value         *big.Int
	ConfirmAmount *big.Int
}, error) {
	return _BREMICO.Contract.Request(&_BREMICO.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BREMICO *BREMICOCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BREMICO *BREMICOSession) Token() (common.Address, error) {
	return _BREMICO.Contract.Token(&_BREMICO.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BREMICO *BREMICOCallerSession) Token() (common.Address, error) {
	return _BREMICO.Contract.Token(&_BREMICO.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_BREMICO *BREMICOCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_BREMICO *BREMICOSession) Wallet() (common.Address, error) {
	return _BREMICO.Contract.Wallet(&_BREMICO.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_BREMICO *BREMICOCallerSession) Wallet() (common.Address, error) {
	return _BREMICO.Contract.Wallet(&_BREMICO.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_BREMICO *BREMICOCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_BREMICO *BREMICOSession) WeiRaised() (*big.Int, error) {
	return _BREMICO.Contract.WeiRaised(&_BREMICO.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_BREMICO *BREMICOCallerSession) WeiRaised() (*big.Int, error) {
	return _BREMICO.Contract.WeiRaised(&_BREMICO.CallOpts)
}

// WithdrawFeePercent is a free data retrieval call binding the contract method 0x495ef705.
//
// Solidity: function withdrawFeePercent() constant returns(uint256)
func (_BREMICO *BREMICOCaller) WithdrawFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMICO.contract.Call(opts, out, "withdrawFeePercent")
	return *ret0, err
}

// WithdrawFeePercent is a free data retrieval call binding the contract method 0x495ef705.
//
// Solidity: function withdrawFeePercent() constant returns(uint256)
func (_BREMICO *BREMICOSession) WithdrawFeePercent() (*big.Int, error) {
	return _BREMICO.Contract.WithdrawFeePercent(&_BREMICO.CallOpts)
}

// WithdrawFeePercent is a free data retrieval call binding the contract method 0x495ef705.
//
// Solidity: function withdrawFeePercent() constant returns(uint256)
func (_BREMICO *BREMICOCallerSession) WithdrawFeePercent() (*big.Int, error) {
	return _BREMICO.Contract.WithdrawFeePercent(&_BREMICO.CallOpts)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_auditor address) returns()
func (_BREMICO *BREMICOTransactor) AddAuditor(opts *bind.TransactOpts, _auditor common.Address) (*types.Transaction, error) {
	return _BREMICO.contract.Transact(opts, "addAuditor", _auditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_auditor address) returns()
func (_BREMICO *BREMICOSession) AddAuditor(_auditor common.Address) (*types.Transaction, error) {
	return _BREMICO.Contract.AddAuditor(&_BREMICO.TransactOpts, _auditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_auditor address) returns()
func (_BREMICO *BREMICOTransactorSession) AddAuditor(_auditor common.Address) (*types.Transaction, error) {
	return _BREMICO.Contract.AddAuditor(&_BREMICO.TransactOpts, _auditor)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_BREMICO *BREMICOTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _BREMICO.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_BREMICO *BREMICOSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _BREMICO.Contract.BuyTokens(&_BREMICO.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_BREMICO *BREMICOTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _BREMICO.Contract.BuyTokens(&_BREMICO.TransactOpts, _beneficiary)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0x68124f9a.
//
// Solidity: function confirmWithdraw() returns()
func (_BREMICO *BREMICOTransactor) ConfirmWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMICO.contract.Transact(opts, "confirmWithdraw")
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0x68124f9a.
//
// Solidity: function confirmWithdraw() returns()
func (_BREMICO *BREMICOSession) ConfirmWithdraw() (*types.Transaction, error) {
	return _BREMICO.Contract.ConfirmWithdraw(&_BREMICO.TransactOpts)
}

// ConfirmWithdraw is a paid mutator transaction binding the contract method 0x68124f9a.
//
// Solidity: function confirmWithdraw() returns()
func (_BREMICO *BREMICOTransactorSession) ConfirmWithdraw() (*types.Transaction, error) {
	return _BREMICO.Contract.ConfirmWithdraw(&_BREMICO.TransactOpts)
}

// FinishAuditorSelection is a paid mutator transaction binding the contract method 0xea78d1ab.
//
// Solidity: function finishAuditorSelection() returns()
func (_BREMICO *BREMICOTransactor) FinishAuditorSelection(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMICO.contract.Transact(opts, "finishAuditorSelection")
}

// FinishAuditorSelection is a paid mutator transaction binding the contract method 0xea78d1ab.
//
// Solidity: function finishAuditorSelection() returns()
func (_BREMICO *BREMICOSession) FinishAuditorSelection() (*types.Transaction, error) {
	return _BREMICO.Contract.FinishAuditorSelection(&_BREMICO.TransactOpts)
}

// FinishAuditorSelection is a paid mutator transaction binding the contract method 0xea78d1ab.
//
// Solidity: function finishAuditorSelection() returns()
func (_BREMICO *BREMICOTransactorSession) FinishAuditorSelection() (*types.Transaction, error) {
	return _BREMICO.Contract.FinishAuditorSelection(&_BREMICO.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_BREMICO *BREMICOTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMICO.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_BREMICO *BREMICOSession) Refund() (*types.Transaction, error) {
	return _BREMICO.Contract.Refund(&_BREMICO.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_BREMICO *BREMICOTransactorSession) Refund() (*types.Transaction, error) {
	return _BREMICO.Contract.Refund(&_BREMICO.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_value uint256) returns()
func (_BREMICO *BREMICOTransactor) Withdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BREMICO.contract.Transact(opts, "withdraw", _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_value uint256) returns()
func (_BREMICO *BREMICOSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _BREMICO.Contract.Withdraw(&_BREMICO.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_value uint256) returns()
func (_BREMICO *BREMICOTransactorSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _BREMICO.Contract.Withdraw(&_BREMICO.TransactOpts, _value)
}

// BREMICOTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the BREMICO contract.
type BREMICOTokenPurchaseIterator struct {
	Event *BREMICOTokenPurchase // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMICOTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMICOTokenPurchase)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMICOTokenPurchase)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMICOTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMICOTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMICOTokenPurchase represents a TokenPurchase event raised by the BREMICO contract.
type BREMICOTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: e TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_BREMICO *BREMICOFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*BREMICOTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BREMICO.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &BREMICOTokenPurchaseIterator{contract: _BREMICO.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: e TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_BREMICO *BREMICOFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *BREMICOTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BREMICO.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMICOTokenPurchase)
				if err := _BREMICO.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMTokenABI is the input ABI used to generate the binding from.
const BREMTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_burner\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnForRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burned\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BREMTokenBin is the compiled bytecode used for deploying new contracts.
const BREMTokenBin = `0x60806040526003805460a060020a60ff021916905534801561002057600080fd5b50604051610e43380380610e4383398101604052805160208083015160038054600160a060020a0319163317905591830180519093929092019161006a9160049190850190610086565b50805161007e906005906020840190610086565b505050610121565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c757805160ff19168380011785556100f4565b828001600101855582156100f4579182015b828111156100f45782518255916020019190600101906100d9565b50610100929150610104565b5090565b61011e91905b80821115610100576000815560010161010a565b90565b610d13806101306000396000f3006080604052600436106100f05763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100f557806306fdde031461011e578063095ea7b3146101a857806318160ddd146101cc57806323b872dd146101f3578063313ce5671461021d57806340c10f1914610248578063661884631461026c57806370a08231146102905780637d64bcb4146102b15780638da5cb5b146102c657806395d89b41146102f7578063a9059cbb1461030c578063d0482d0414610330578063d73dd62314610356578063dd62ed3e1461037a578063f2fde38b146103a1575b600080fd5b34801561010157600080fd5b5061010a6103c2565b604080519115158252519081900360200190f35b34801561012a57600080fd5b506101336103e3565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016d578181015183820152602001610155565b50505050905090810190601f16801561019a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101b457600080fd5b5061010a600160a060020a0360043516602435610471565b3480156101d857600080fd5b506101e16104d7565b60408051918252519081900360200190f35b3480156101ff57600080fd5b5061010a600160a060020a03600435811690602435166044356104dd565b34801561022957600080fd5b50610232610644565b6040805160ff9092168252519081900360200190f35b34801561025457600080fd5b5061010a600160a060020a0360043516602435610649565b34801561027857600080fd5b5061010a600160a060020a0360043516602435610753565b34801561029c57600080fd5b506101e1600160a060020a0360043516610843565b3480156102bd57600080fd5b5061010a61085e565b3480156102d257600080fd5b506102db610904565b60408051600160a060020a039092168252519081900360200190f35b34801561030357600080fd5b50610133610913565b34801561031857600080fd5b5061010a600160a060020a036004351660243561096e565b34801561033c57600080fd5b50610354600160a060020a0360043516602435610a3f565b005b34801561036257600080fd5b5061010a600160a060020a0360043516602435610b46565b34801561038657600080fd5b506101e1600160a060020a0360043581169060243516610bdf565b3480156103ad57600080fd5b50610354600160a060020a0360043516610c0a565b60035474010000000000000000000000000000000000000000900460ff1681565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104695780601f1061043e57610100808354040283529160200191610469565b820191906000526020600020905b81548152906001019060200180831161044c57829003601f168201915b505050505081565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b6000600160a060020a03831615156104f457600080fd5b600160a060020a03841660009081526001602052604090205482111561051957600080fd5b600160a060020a038416600090815260026020908152604080832033845290915290205482111561054957600080fd5b600160a060020a038416600090815260016020526040902054610572908363ffffffff610c9f16565b600160a060020a0380861660009081526001602052604080822093909355908516815220546105a7908363ffffffff610cb116565b600160a060020a0380851660009081526001602090815260408083209490945591871681526002825282812033825290915220546105eb908363ffffffff610c9f16565b600160a060020a0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610cc8833981519152929181900390910190a35060019392505050565b601281565b600354600090600160a060020a0316331461066357600080fd5b60035474010000000000000000000000000000000000000000900460ff161561068b57600080fd5b60005461069e908363ffffffff610cb116565b6000908155600160a060020a0384168152600160205260409020546106c9908363ffffffff610cb116565b600160a060020a038416600081815260016020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a2604080518381529051600160a060020a03851691600091600080516020610cc88339815191529181900360200190a350600192915050565b336000908152600260209081526040808320600160a060020a0386168452909152812054808311156107a857336000908152600260209081526040808320600160a060020a03881684529091528120556107dd565b6107b8818463ffffffff610c9f16565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526001602052604090205490565b600354600090600160a060020a0316331461087857600080fd5b60035474010000000000000000000000000000000000000000900460ff16156108a057600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b600354600160a060020a031681565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104695780601f1061043e57610100808354040283529160200191610469565b6000600160a060020a038316151561098557600080fd5b336000908152600160205260409020548211156109a157600080fd5b336000908152600160205260409020546109c1908363ffffffff610c9f16565b3360009081526001602052604080822092909255600160a060020a038516815220546109f3908363ffffffff610cb116565b600160a060020a038416600081815260016020908152604091829020939093558051858152905191923392600080516020610cc88339815191529281900390910190a350600192915050565b600354600160a060020a03163314610a5657600080fd5b600160a060020a038216600090815260016020526040902054811115610a7b57600080fd5b600160a060020a038216600090815260016020526040902054610aa4908263ffffffff610c9f16565b600160a060020a03831660009081526001602052604081209190915554610ad1908263ffffffff610c9f16565b600055604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a03851691600080516020610cc88339815191529181900360200190a35050565b336000908152600260209081526040808320600160a060020a0386168452909152812054610b7a908363ffffffff610cb116565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a03163314610c2157600080fd5b600160a060020a0381161515610c3657600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610cab57fe5b50900390565b600082820183811015610cc057fe5b93925050505600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a72305820f411444b49ce555e9034faf955522fd13f9ccde69d0c577c64eaaafe9c322b3a0029`

// DeployBREMToken deploys a new Ethereum contract, binding an instance of BREMToken to it.
func DeployBREMToken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *BREMToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BREMTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BREMTokenBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BREMToken{BREMTokenCaller: BREMTokenCaller{contract: contract}, BREMTokenTransactor: BREMTokenTransactor{contract: contract}, BREMTokenFilterer: BREMTokenFilterer{contract: contract}}, nil
}

// BREMToken is an auto generated Go binding around an Ethereum contract.
type BREMToken struct {
	BREMTokenCaller     // Read-only binding to the contract
	BREMTokenTransactor // Write-only binding to the contract
	BREMTokenFilterer   // Log filterer for contract events
}

// BREMTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BREMTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BREMTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BREMTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BREMTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BREMTokenSession struct {
	Contract     *BREMToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BREMTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BREMTokenCallerSession struct {
	Contract *BREMTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BREMTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BREMTokenTransactorSession struct {
	Contract     *BREMTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BREMTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BREMTokenRaw struct {
	Contract *BREMToken // Generic contract binding to access the raw methods on
}

// BREMTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BREMTokenCallerRaw struct {
	Contract *BREMTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BREMTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BREMTokenTransactorRaw struct {
	Contract *BREMTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBREMToken creates a new instance of BREMToken, bound to a specific deployed contract.
func NewBREMToken(address common.Address, backend bind.ContractBackend) (*BREMToken, error) {
	contract, err := bindBREMToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BREMToken{BREMTokenCaller: BREMTokenCaller{contract: contract}, BREMTokenTransactor: BREMTokenTransactor{contract: contract}, BREMTokenFilterer: BREMTokenFilterer{contract: contract}}, nil
}

// NewBREMTokenCaller creates a new read-only instance of BREMToken, bound to a specific deployed contract.
func NewBREMTokenCaller(address common.Address, caller bind.ContractCaller) (*BREMTokenCaller, error) {
	contract, err := bindBREMToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BREMTokenCaller{contract: contract}, nil
}

// NewBREMTokenTransactor creates a new write-only instance of BREMToken, bound to a specific deployed contract.
func NewBREMTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BREMTokenTransactor, error) {
	contract, err := bindBREMToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BREMTokenTransactor{contract: contract}, nil
}

// NewBREMTokenFilterer creates a new log filterer instance of BREMToken, bound to a specific deployed contract.
func NewBREMTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BREMTokenFilterer, error) {
	contract, err := bindBREMToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BREMTokenFilterer{contract: contract}, nil
}

// bindBREMToken binds a generic wrapper to an already deployed contract.
func bindBREMToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BREMTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BREMToken *BREMTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BREMToken.Contract.BREMTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BREMToken *BREMTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMToken.Contract.BREMTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BREMToken *BREMTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BREMToken.Contract.BREMTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BREMToken *BREMTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BREMToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BREMToken *BREMTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BREMToken *BREMTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BREMToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BREMToken *BREMTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BREMToken *BREMTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BREMToken.Contract.Allowance(&_BREMToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BREMToken *BREMTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BREMToken.Contract.Allowance(&_BREMToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BREMToken *BREMTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BREMToken *BREMTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BREMToken.Contract.BalanceOf(&_BREMToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BREMToken *BREMTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BREMToken.Contract.BalanceOf(&_BREMToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BREMToken *BREMTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BREMToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BREMToken *BREMTokenSession) Decimals() (uint8, error) {
	return _BREMToken.Contract.Decimals(&_BREMToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BREMToken *BREMTokenCallerSession) Decimals() (uint8, error) {
	return _BREMToken.Contract.Decimals(&_BREMToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_BREMToken *BREMTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BREMToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_BREMToken *BREMTokenSession) MintingFinished() (bool, error) {
	return _BREMToken.Contract.MintingFinished(&_BREMToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_BREMToken *BREMTokenCallerSession) MintingFinished() (bool, error) {
	return _BREMToken.Contract.MintingFinished(&_BREMToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BREMToken *BREMTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREMToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BREMToken *BREMTokenSession) Name() (string, error) {
	return _BREMToken.Contract.Name(&_BREMToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BREMToken *BREMTokenCallerSession) Name() (string, error) {
	return _BREMToken.Contract.Name(&_BREMToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BREMToken *BREMTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BREMToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BREMToken *BREMTokenSession) Owner() (common.Address, error) {
	return _BREMToken.Contract.Owner(&_BREMToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BREMToken *BREMTokenCallerSession) Owner() (common.Address, error) {
	return _BREMToken.Contract.Owner(&_BREMToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BREMToken *BREMTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BREMToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BREMToken *BREMTokenSession) Symbol() (string, error) {
	return _BREMToken.Contract.Symbol(&_BREMToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BREMToken *BREMTokenCallerSession) Symbol() (string, error) {
	return _BREMToken.Contract.Symbol(&_BREMToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BREMToken *BREMTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BREMToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BREMToken *BREMTokenSession) TotalSupply() (*big.Int, error) {
	return _BREMToken.Contract.TotalSupply(&_BREMToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BREMToken *BREMTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BREMToken.Contract.TotalSupply(&_BREMToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BREMToken *BREMTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BREMToken *BREMTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.Approve(&_BREMToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BREMToken *BREMTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.Approve(&_BREMToken.TransactOpts, _spender, _value)
}

// BurnForRefund is a paid mutator transaction binding the contract method 0xd0482d04.
//
// Solidity: function burnForRefund(_burner address, _value uint256) returns()
func (_BREMToken *BREMTokenTransactor) BurnForRefund(opts *bind.TransactOpts, _burner common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.contract.Transact(opts, "burnForRefund", _burner, _value)
}

// BurnForRefund is a paid mutator transaction binding the contract method 0xd0482d04.
//
// Solidity: function burnForRefund(_burner address, _value uint256) returns()
func (_BREMToken *BREMTokenSession) BurnForRefund(_burner common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.BurnForRefund(&_BREMToken.TransactOpts, _burner, _value)
}

// BurnForRefund is a paid mutator transaction binding the contract method 0xd0482d04.
//
// Solidity: function burnForRefund(_burner address, _value uint256) returns()
func (_BREMToken *BREMTokenTransactorSession) BurnForRefund(_burner common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.BurnForRefund(&_BREMToken.TransactOpts, _burner, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BREMToken *BREMTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BREMToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BREMToken *BREMTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.DecreaseApproval(&_BREMToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BREMToken *BREMTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.DecreaseApproval(&_BREMToken.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_BREMToken *BREMTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BREMToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_BREMToken *BREMTokenSession) FinishMinting() (*types.Transaction, error) {
	return _BREMToken.Contract.FinishMinting(&_BREMToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_BREMToken *BREMTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _BREMToken.Contract.FinishMinting(&_BREMToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BREMToken *BREMTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BREMToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BREMToken *BREMTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.IncreaseApproval(&_BREMToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BREMToken *BREMTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.IncreaseApproval(&_BREMToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_BREMToken *BREMTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BREMToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_BREMToken *BREMTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.Mint(&_BREMToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_BREMToken *BREMTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.Mint(&_BREMToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BREMToken *BREMTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BREMToken *BREMTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.Transfer(&_BREMToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BREMToken *BREMTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.Transfer(&_BREMToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BREMToken *BREMTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BREMToken *BREMTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.TransferFrom(&_BREMToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BREMToken *BREMTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BREMToken.Contract.TransferFrom(&_BREMToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BREMToken *BREMTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BREMToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BREMToken *BREMTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BREMToken.Contract.TransferOwnership(&_BREMToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BREMToken *BREMTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BREMToken.Contract.TransferOwnership(&_BREMToken.TransactOpts, newOwner)
}

// BREMTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BREMToken contract.
type BREMTokenApprovalIterator struct {
	Event *BREMTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMTokenApproval represents a Approval event raised by the BREMToken contract.
type BREMTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BREMToken *BREMTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BREMTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BREMToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BREMTokenApprovalIterator{contract: _BREMToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BREMToken *BREMTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BREMTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BREMToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMTokenApproval)
				if err := _BREMToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BREMToken contract.
type BREMTokenBurnIterator struct {
	Event *BREMTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMTokenBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMTokenBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMTokenBurn represents a Burn event raised by the BREMToken contract.
type BREMTokenBurn struct {
	Burned common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burned indexed address, value uint256)
func (_BREMToken *BREMTokenFilterer) FilterBurn(opts *bind.FilterOpts, burned []common.Address) (*BREMTokenBurnIterator, error) {

	var burnedRule []interface{}
	for _, burnedItem := range burned {
		burnedRule = append(burnedRule, burnedItem)
	}

	logs, sub, err := _BREMToken.contract.FilterLogs(opts, "Burn", burnedRule)
	if err != nil {
		return nil, err
	}
	return &BREMTokenBurnIterator{contract: _BREMToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burned indexed address, value uint256)
func (_BREMToken *BREMTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BREMTokenBurn, burned []common.Address) (event.Subscription, error) {

	var burnedRule []interface{}
	for _, burnedItem := range burned {
		burnedRule = append(burnedRule, burnedItem)
	}

	logs, sub, err := _BREMToken.contract.WatchLogs(opts, "Burn", burnedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMTokenBurn)
				if err := _BREMToken.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BREMToken contract.
type BREMTokenMintIterator struct {
	Event *BREMTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMTokenMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMTokenMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMTokenMint represents a Mint event raised by the BREMToken contract.
type BREMTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_BREMToken *BREMTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*BREMTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BREMToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &BREMTokenMintIterator{contract: _BREMToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_BREMToken *BREMTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BREMTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BREMToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMTokenMint)
				if err := _BREMToken.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the BREMToken contract.
type BREMTokenMintFinishedIterator struct {
	Event *BREMTokenMintFinished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMTokenMintFinished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMTokenMintFinished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMTokenMintFinished represents a MintFinished event raised by the BREMToken contract.
type BREMTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_BREMToken *BREMTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*BREMTokenMintFinishedIterator, error) {

	logs, sub, err := _BREMToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &BREMTokenMintFinishedIterator{contract: _BREMToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_BREMToken *BREMTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *BREMTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _BREMToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMTokenMintFinished)
				if err := _BREMToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BREMToken contract.
type BREMTokenOwnershipTransferredIterator struct {
	Event *BREMTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BREMToken contract.
type BREMTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BREMToken *BREMTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BREMTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BREMToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BREMTokenOwnershipTransferredIterator{contract: _BREMToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BREMToken *BREMTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BREMTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BREMToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMTokenOwnershipTransferred)
				if err := _BREMToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BREMTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BREMToken contract.
type BREMTokenTransferIterator struct {
	Event *BREMTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BREMTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BREMTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BREMTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BREMTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BREMTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BREMTokenTransfer represents a Transfer event raised by the BREMToken contract.
type BREMTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BREMToken *BREMTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BREMTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BREMToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BREMTokenTransferIterator{contract: _BREMToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BREMToken *BREMTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BREMTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BREMToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BREMTokenTransfer)
				if err := _BREMToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BRMTokenABI is the input ABI used to generate the binding from.
const BRMTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BRMTokenBin is the compiled bytecode used for deploying new contracts.
const BRMTokenBin = `0x608060405260038054600160a860020a03191633179055610b2c806100256000396000f3006080604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100ea57806306fdde0314610113578063095ea7b31461019d57806318160ddd146101c157806323b872dd146101e8578063313ce5671461021257806340c10f191461023d578063661884631461026157806370a08231146102855780637d64bcb4146102a65780638da5cb5b146102bb57806395d89b4114610113578063a9059cbb146102ec578063d73dd62314610310578063dd62ed3e14610334578063f2fde38b1461035b575b600080fd5b3480156100f657600080fd5b506100ff61037e565b604080519115158252519081900360200190f35b34801561011f57600080fd5b5061012861039f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561016257818101518382015260200161014a565b50505050905090810190601f16801561018f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101a957600080fd5b506100ff600160a060020a03600435166024356103d6565b3480156101cd57600080fd5b506101d661043c565b60408051918252519081900360200190f35b3480156101f457600080fd5b506100ff600160a060020a0360043581169060243516604435610442565b34801561021e57600080fd5b506102276105bb565b6040805160ff9092168252519081900360200190f35b34801561024957600080fd5b506100ff600160a060020a03600435166024356105c0565b34801561026d57600080fd5b506100ff600160a060020a03600435166024356106dc565b34801561029157600080fd5b506101d6600160a060020a03600435166107cc565b3480156102b257600080fd5b506100ff6107e7565b3480156102c757600080fd5b506102d061088d565b60408051600160a060020a039092168252519081900360200190f35b3480156102f857600080fd5b506100ff600160a060020a036004351660243561089c565b34801561031c57600080fd5b506100ff600160a060020a036004351660243561097f565b34801561034057600080fd5b506101d6600160a060020a0360043581169060243516610a18565b34801561036757600080fd5b5061037c600160a060020a0360043516610a43565b005b60035474010000000000000000000000000000000000000000900460ff1681565b60408051808201909152600381527f42524d0000000000000000000000000000000000000000000000000000000000602082015281565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b6000600160a060020a038316151561045957600080fd5b600160a060020a03841660009081526001602052604090205482111561047e57600080fd5b600160a060020a03841660009081526002602090815260408083203384529091529020548211156104ae57600080fd5b600160a060020a0384166000908152600160205260409020546104d7908363ffffffff610ad816565b600160a060020a03808616600090815260016020526040808220939093559085168152205461050c908363ffffffff610aea16565b600160a060020a038085166000908152600160209081526040808320949094559187168152600282528281203382529091522054610550908363ffffffff610ad816565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b601281565b600354600090600160a060020a031633146105da57600080fd5b60035474010000000000000000000000000000000000000000900460ff161561060257600080fd5b600054610615908363ffffffff610aea16565b6000908155600160a060020a038416815260016020526040902054610640908363ffffffff610aea16565b600160a060020a038416600081815260016020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a2604080518381529051600160a060020a038516916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120548083111561073157336000908152600260209081526040808320600160a060020a0388168452909152812055610766565b610741818463ffffffff610ad816565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526001602052604090205490565b600354600090600160a060020a0316331461080157600080fd5b60035474010000000000000000000000000000000000000000900460ff161561082957600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b600354600160a060020a031681565b6000600160a060020a03831615156108b357600080fd5b336000908152600160205260409020548211156108cf57600080fd5b336000908152600160205260409020546108ef908363ffffffff610ad816565b3360009081526001602052604080822092909255600160a060020a03851681522054610921908363ffffffff610aea16565b600160a060020a0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120546109b3908363ffffffff610aea16565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a03163314610a5a57600080fd5b600160a060020a0381161515610a6f57600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610ae457fe5b50900390565b600082820183811015610af957fe5b93925050505600a165627a7a72305820e95520f0167098c11fc6a267a7f77b9e17eee0731694b9116735e261a0adc7540029`

// DeployBRMToken deploys a new Ethereum contract, binding an instance of BRMToken to it.
func DeployBRMToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BRMToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BRMTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BRMTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BRMToken{BRMTokenCaller: BRMTokenCaller{contract: contract}, BRMTokenTransactor: BRMTokenTransactor{contract: contract}, BRMTokenFilterer: BRMTokenFilterer{contract: contract}}, nil
}

// BRMToken is an auto generated Go binding around an Ethereum contract.
type BRMToken struct {
	BRMTokenCaller     // Read-only binding to the contract
	BRMTokenTransactor // Write-only binding to the contract
	BRMTokenFilterer   // Log filterer for contract events
}

// BRMTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BRMTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BRMTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BRMTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BRMTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BRMTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BRMTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BRMTokenSession struct {
	Contract     *BRMToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BRMTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BRMTokenCallerSession struct {
	Contract *BRMTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BRMTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BRMTokenTransactorSession struct {
	Contract     *BRMTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BRMTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BRMTokenRaw struct {
	Contract *BRMToken // Generic contract binding to access the raw methods on
}

// BRMTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BRMTokenCallerRaw struct {
	Contract *BRMTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BRMTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BRMTokenTransactorRaw struct {
	Contract *BRMTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBRMToken creates a new instance of BRMToken, bound to a specific deployed contract.
func NewBRMToken(address common.Address, backend bind.ContractBackend) (*BRMToken, error) {
	contract, err := bindBRMToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BRMToken{BRMTokenCaller: BRMTokenCaller{contract: contract}, BRMTokenTransactor: BRMTokenTransactor{contract: contract}, BRMTokenFilterer: BRMTokenFilterer{contract: contract}}, nil
}

// NewBRMTokenCaller creates a new read-only instance of BRMToken, bound to a specific deployed contract.
func NewBRMTokenCaller(address common.Address, caller bind.ContractCaller) (*BRMTokenCaller, error) {
	contract, err := bindBRMToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BRMTokenCaller{contract: contract}, nil
}

// NewBRMTokenTransactor creates a new write-only instance of BRMToken, bound to a specific deployed contract.
func NewBRMTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BRMTokenTransactor, error) {
	contract, err := bindBRMToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BRMTokenTransactor{contract: contract}, nil
}

// NewBRMTokenFilterer creates a new log filterer instance of BRMToken, bound to a specific deployed contract.
func NewBRMTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BRMTokenFilterer, error) {
	contract, err := bindBRMToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BRMTokenFilterer{contract: contract}, nil
}

// bindBRMToken binds a generic wrapper to an already deployed contract.
func bindBRMToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BRMTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BRMToken *BRMTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BRMToken.Contract.BRMTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BRMToken *BRMTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BRMToken.Contract.BRMTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BRMToken *BRMTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BRMToken.Contract.BRMTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BRMToken *BRMTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BRMToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BRMToken *BRMTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BRMToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BRMToken *BRMTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BRMToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BRMToken *BRMTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BRMToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BRMToken *BRMTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BRMToken.Contract.Allowance(&_BRMToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BRMToken *BRMTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BRMToken.Contract.Allowance(&_BRMToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BRMToken *BRMTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BRMToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BRMToken *BRMTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BRMToken.Contract.BalanceOf(&_BRMToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BRMToken *BRMTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BRMToken.Contract.BalanceOf(&_BRMToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BRMToken *BRMTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BRMToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BRMToken *BRMTokenSession) Decimals() (uint8, error) {
	return _BRMToken.Contract.Decimals(&_BRMToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_BRMToken *BRMTokenCallerSession) Decimals() (uint8, error) {
	return _BRMToken.Contract.Decimals(&_BRMToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_BRMToken *BRMTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BRMToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_BRMToken *BRMTokenSession) MintingFinished() (bool, error) {
	return _BRMToken.Contract.MintingFinished(&_BRMToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_BRMToken *BRMTokenCallerSession) MintingFinished() (bool, error) {
	return _BRMToken.Contract.MintingFinished(&_BRMToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BRMToken *BRMTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BRMToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BRMToken *BRMTokenSession) Name() (string, error) {
	return _BRMToken.Contract.Name(&_BRMToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BRMToken *BRMTokenCallerSession) Name() (string, error) {
	return _BRMToken.Contract.Name(&_BRMToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BRMToken *BRMTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BRMToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BRMToken *BRMTokenSession) Owner() (common.Address, error) {
	return _BRMToken.Contract.Owner(&_BRMToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BRMToken *BRMTokenCallerSession) Owner() (common.Address, error) {
	return _BRMToken.Contract.Owner(&_BRMToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BRMToken *BRMTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BRMToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BRMToken *BRMTokenSession) Symbol() (string, error) {
	return _BRMToken.Contract.Symbol(&_BRMToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_BRMToken *BRMTokenCallerSession) Symbol() (string, error) {
	return _BRMToken.Contract.Symbol(&_BRMToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BRMToken *BRMTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BRMToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BRMToken *BRMTokenSession) TotalSupply() (*big.Int, error) {
	return _BRMToken.Contract.TotalSupply(&_BRMToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BRMToken *BRMTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BRMToken.Contract.TotalSupply(&_BRMToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BRMToken *BRMTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BRMToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BRMToken *BRMTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.Approve(&_BRMToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BRMToken *BRMTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.Approve(&_BRMToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BRMToken *BRMTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BRMToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BRMToken *BRMTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.DecreaseApproval(&_BRMToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BRMToken *BRMTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.DecreaseApproval(&_BRMToken.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_BRMToken *BRMTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BRMToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_BRMToken *BRMTokenSession) FinishMinting() (*types.Transaction, error) {
	return _BRMToken.Contract.FinishMinting(&_BRMToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_BRMToken *BRMTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _BRMToken.Contract.FinishMinting(&_BRMToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BRMToken *BRMTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BRMToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BRMToken *BRMTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.IncreaseApproval(&_BRMToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BRMToken *BRMTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.IncreaseApproval(&_BRMToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_BRMToken *BRMTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BRMToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_BRMToken *BRMTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.Mint(&_BRMToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_BRMToken *BRMTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.Mint(&_BRMToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BRMToken *BRMTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BRMToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BRMToken *BRMTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.Transfer(&_BRMToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BRMToken *BRMTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.Transfer(&_BRMToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BRMToken *BRMTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BRMToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BRMToken *BRMTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.TransferFrom(&_BRMToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BRMToken *BRMTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BRMToken.Contract.TransferFrom(&_BRMToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BRMToken *BRMTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BRMToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BRMToken *BRMTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BRMToken.Contract.TransferOwnership(&_BRMToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BRMToken *BRMTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BRMToken.Contract.TransferOwnership(&_BRMToken.TransactOpts, newOwner)
}

// BRMTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BRMToken contract.
type BRMTokenApprovalIterator struct {
	Event *BRMTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BRMTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BRMTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BRMTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BRMTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BRMTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BRMTokenApproval represents a Approval event raised by the BRMToken contract.
type BRMTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BRMToken *BRMTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BRMTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BRMToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BRMTokenApprovalIterator{contract: _BRMToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BRMToken *BRMTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BRMTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BRMToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BRMTokenApproval)
				if err := _BRMToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BRMTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BRMToken contract.
type BRMTokenMintIterator struct {
	Event *BRMTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BRMTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BRMTokenMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BRMTokenMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BRMTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BRMTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BRMTokenMint represents a Mint event raised by the BRMToken contract.
type BRMTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_BRMToken *BRMTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*BRMTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BRMToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &BRMTokenMintIterator{contract: _BRMToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_BRMToken *BRMTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BRMTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BRMToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BRMTokenMint)
				if err := _BRMToken.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BRMTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the BRMToken contract.
type BRMTokenMintFinishedIterator struct {
	Event *BRMTokenMintFinished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BRMTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BRMTokenMintFinished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BRMTokenMintFinished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BRMTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BRMTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BRMTokenMintFinished represents a MintFinished event raised by the BRMToken contract.
type BRMTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_BRMToken *BRMTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*BRMTokenMintFinishedIterator, error) {

	logs, sub, err := _BRMToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &BRMTokenMintFinishedIterator{contract: _BRMToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_BRMToken *BRMTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *BRMTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _BRMToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BRMTokenMintFinished)
				if err := _BRMToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BRMTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BRMToken contract.
type BRMTokenOwnershipTransferredIterator struct {
	Event *BRMTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BRMTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BRMTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BRMTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BRMTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BRMTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BRMTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BRMToken contract.
type BRMTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BRMToken *BRMTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BRMTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BRMToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BRMTokenOwnershipTransferredIterator{contract: _BRMToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BRMToken *BRMTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BRMTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BRMToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BRMTokenOwnershipTransferred)
				if err := _BRMToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BRMTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BRMToken contract.
type BRMTokenTransferIterator struct {
	Event *BRMTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BRMTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BRMTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BRMTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BRMTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BRMTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BRMTokenTransfer represents a Transfer event raised by the BRMToken contract.
type BRMTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BRMToken *BRMTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BRMTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BRMToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BRMTokenTransferIterator{contract: _BRMToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BRMToken *BRMTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BRMTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BRMToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BRMTokenTransfer)
				if err := _BRMToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x608060405234801561001057600080fd5b50610281806100206000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610082578063a9059cbb146100b0575b600080fd5b34801561006757600080fd5b506100706100f5565b60408051918252519081900360200190f35b34801561008e57600080fd5b5061007073ffffffffffffffffffffffffffffffffffffffff600435166100fb565b3480156100bc57600080fd5b506100e173ffffffffffffffffffffffffffffffffffffffff60043516602435610123565b604080519115158252519081900360200190f35b60005481565b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205490565b600073ffffffffffffffffffffffffffffffffffffffff8316151561014757600080fd5b3360009081526001602052604090205482111561016357600080fd5b33600090815260016020526040902054610183908363ffffffff61022d16565b336000908152600160205260408082209290925573ffffffffffffffffffffffffffffffffffffffff8516815220546101c2908363ffffffff61023f16565b73ffffffffffffffffffffffffffffffffffffffff84166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b60008282111561023957fe5b50900390565b60008282018381101561024e57fe5b93925050505600a165627a7a723058200e0df7df457c677f39b20af040153bf6a4f8babfeb1ef67b1dc85a781435fb560029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BasicTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BurnableTokenABI is the input ABI used to generate the binding from.
const BurnableTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_burner\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burnForRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burned\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BurnableTokenBin is the compiled bytecode used for deploying new contracts.
const BurnableTokenBin = `0x608060405260038054600160a860020a03191633179055610b3a806100256000396000f3006080604052600436106100cf5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100d4578063095ea7b3146100fd57806318160ddd1461012157806323b872dd1461014857806340c10f1914610172578063661884631461019657806370a08231146101ba5780637d64bcb4146101db5780638da5cb5b146101f0578063a9059cbb14610221578063d0482d0414610245578063d73dd6231461026b578063dd62ed3e1461028f578063f2fde38b146102b6575b600080fd5b3480156100e057600080fd5b506100e96102d7565b604080519115158252519081900360200190f35b34801561010957600080fd5b506100e9600160a060020a03600435166024356102f8565b34801561012d57600080fd5b5061013661035e565b60408051918252519081900360200190f35b34801561015457600080fd5b506100e9600160a060020a0360043581169060243516604435610364565b34801561017e57600080fd5b506100e9600160a060020a03600435166024356104cb565b3480156101a257600080fd5b506100e9600160a060020a03600435166024356105d5565b3480156101c657600080fd5b50610136600160a060020a03600435166106c5565b3480156101e757600080fd5b506100e96106e0565b3480156101fc57600080fd5b50610205610786565b60408051600160a060020a039092168252519081900360200190f35b34801561022d57600080fd5b506100e9600160a060020a0360043516602435610795565b34801561025157600080fd5b50610269600160a060020a0360043516602435610866565b005b34801561027757600080fd5b506100e9600160a060020a036004351660243561096d565b34801561029b57600080fd5b50610136600160a060020a0360043581169060243516610a06565b3480156102c257600080fd5b50610269600160a060020a0360043516610a31565b60035474010000000000000000000000000000000000000000900460ff1681565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b6000600160a060020a038316151561037b57600080fd5b600160a060020a0384166000908152600160205260409020548211156103a057600080fd5b600160a060020a03841660009081526002602090815260408083203384529091529020548211156103d057600080fd5b600160a060020a0384166000908152600160205260409020546103f9908363ffffffff610ac616565b600160a060020a03808616600090815260016020526040808220939093559085168152205461042e908363ffffffff610ad816565b600160a060020a038085166000908152600160209081526040808320949094559187168152600282528281203382529091522054610472908363ffffffff610ac616565b600160a060020a0380861660008181526002602090815260408083203384528252918290209490945580518681529051928716939192600080516020610aef833981519152929181900390910190a35060019392505050565b600354600090600160a060020a031633146104e557600080fd5b60035474010000000000000000000000000000000000000000900460ff161561050d57600080fd5b600054610520908363ffffffff610ad816565b6000908155600160a060020a03841681526001602052604090205461054b908363ffffffff610ad816565b600160a060020a038416600081815260016020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a2604080518381529051600160a060020a03851691600091600080516020610aef8339815191529181900360200190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120548083111561062a57336000908152600260209081526040808320600160a060020a038816845290915281205561065f565b61063a818463ffffffff610ac616565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526001602052604090205490565b600354600090600160a060020a031633146106fa57600080fd5b60035474010000000000000000000000000000000000000000900460ff161561072257600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b600354600160a060020a031681565b6000600160a060020a03831615156107ac57600080fd5b336000908152600160205260409020548211156107c857600080fd5b336000908152600160205260409020546107e8908363ffffffff610ac616565b3360009081526001602052604080822092909255600160a060020a0385168152205461081a908363ffffffff610ad816565b600160a060020a038416600081815260016020908152604091829020939093558051858152905191923392600080516020610aef8339815191529281900390910190a350600192915050565b600354600160a060020a0316331461087d57600080fd5b600160a060020a0382166000908152600160205260409020548111156108a257600080fd5b600160a060020a0382166000908152600160205260409020546108cb908263ffffffff610ac616565b600160a060020a038316600090815260016020526040812091909155546108f8908263ffffffff610ac616565b600055604080518281529051600160a060020a038416917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518281529051600091600160a060020a03851691600080516020610aef8339815191529181900360200190a35050565b336000908152600260209081526040808320600160a060020a03861684529091528120546109a1908363ffffffff610ad816565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a03163314610a4857600080fd5b600160a060020a0381161515610a5d57600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600082821115610ad257fe5b50900390565b600082820183811015610ae757fe5b93925050505600ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa165627a7a72305820d11612cd1e7065deb52989337645268363558d8ec5e3911ad77619fc1110327d0029`

// DeployBurnableToken deploys a new Ethereum contract, binding an instance of BurnableToken to it.
func DeployBurnableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BurnableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// BurnableToken is an auto generated Go binding around an Ethereum contract.
type BurnableToken struct {
	BurnableTokenCaller     // Read-only binding to the contract
	BurnableTokenTransactor // Write-only binding to the contract
	BurnableTokenFilterer   // Log filterer for contract events
}

// BurnableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnableTokenSession struct {
	Contract     *BurnableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnableTokenCallerSession struct {
	Contract *BurnableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BurnableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnableTokenTransactorSession struct {
	Contract     *BurnableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BurnableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnableTokenRaw struct {
	Contract *BurnableToken // Generic contract binding to access the raw methods on
}

// BurnableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnableTokenCallerRaw struct {
	Contract *BurnableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BurnableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnableTokenTransactorRaw struct {
	Contract *BurnableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnableToken creates a new instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableToken(address common.Address, backend bind.ContractBackend) (*BurnableToken, error) {
	contract, err := bindBurnableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// NewBurnableTokenCaller creates a new read-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenCaller(address common.Address, caller bind.ContractCaller) (*BurnableTokenCaller, error) {
	contract, err := bindBurnableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenCaller{contract: contract}, nil
}

// NewBurnableTokenTransactor creates a new write-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnableTokenTransactor, error) {
	contract, err := bindBurnableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransactor{contract: contract}, nil
}

// NewBurnableTokenFilterer creates a new log filterer instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnableTokenFilterer, error) {
	contract, err := bindBurnableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenFilterer{contract: contract}, nil
}

// bindBurnableToken binds a generic wrapper to an already deployed contract.
func bindBurnableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.BurnableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BurnableToken *BurnableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BurnableToken *BurnableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.Allowance(&_BurnableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_BurnableToken *BurnableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.Allowance(&_BurnableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableToken *BurnableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableToken *BurnableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableToken *BurnableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_BurnableToken *BurnableTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_BurnableToken *BurnableTokenSession) MintingFinished() (bool, error) {
	return _BurnableToken.Contract.MintingFinished(&_BurnableToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_BurnableToken *BurnableTokenCallerSession) MintingFinished() (bool, error) {
	return _BurnableToken.Contract.MintingFinished(&_BurnableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnableToken *BurnableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnableToken *BurnableTokenSession) Owner() (common.Address, error) {
	return _BurnableToken.Contract.Owner(&_BurnableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnableToken *BurnableTokenCallerSession) Owner() (common.Address, error) {
	return _BurnableToken.Contract.Owner(&_BurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Approve(&_BurnableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Approve(&_BurnableToken.TransactOpts, _spender, _value)
}

// BurnForRefund is a paid mutator transaction binding the contract method 0xd0482d04.
//
// Solidity: function burnForRefund(_burner address, _value uint256) returns()
func (_BurnableToken *BurnableTokenTransactor) BurnForRefund(opts *bind.TransactOpts, _burner common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "burnForRefund", _burner, _value)
}

// BurnForRefund is a paid mutator transaction binding the contract method 0xd0482d04.
//
// Solidity: function burnForRefund(_burner address, _value uint256) returns()
func (_BurnableToken *BurnableTokenSession) BurnForRefund(_burner common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnForRefund(&_BurnableToken.TransactOpts, _burner, _value)
}

// BurnForRefund is a paid mutator transaction binding the contract method 0xd0482d04.
//
// Solidity: function burnForRefund(_burner address, _value uint256) returns()
func (_BurnableToken *BurnableTokenTransactorSession) BurnForRefund(_burner common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnForRefund(&_BurnableToken.TransactOpts, _burner, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BurnableToken *BurnableTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.DecreaseApproval(&_BurnableToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.DecreaseApproval(&_BurnableToken.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_BurnableToken *BurnableTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_BurnableToken *BurnableTokenSession) FinishMinting() (*types.Transaction, error) {
	return _BurnableToken.Contract.FinishMinting(&_BurnableToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _BurnableToken.Contract.FinishMinting(&_BurnableToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BurnableToken *BurnableTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.IncreaseApproval(&_BurnableToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.IncreaseApproval(&_BurnableToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_BurnableToken *BurnableTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Mint(&_BurnableToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Mint(&_BurnableToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.TransferFrom(&_BurnableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.TransferFrom(&_BurnableToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BurnableToken *BurnableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BurnableToken *BurnableTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BurnableToken.Contract.TransferOwnership(&_BurnableToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BurnableToken *BurnableTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BurnableToken.Contract.TransferOwnership(&_BurnableToken.TransactOpts, newOwner)
}

// BurnableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BurnableToken contract.
type BurnableTokenApprovalIterator struct {
	Event *BurnableTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BurnableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BurnableTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BurnableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenApproval represents a Approval event raised by the BurnableToken contract.
type BurnableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BurnableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenApprovalIterator{contract: _BurnableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BurnableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenApproval)
				if err := _BurnableToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BurnableTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BurnableToken contract.
type BurnableTokenBurnIterator struct {
	Event *BurnableTokenBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BurnableTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BurnableTokenBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BurnableTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenBurn represents a Burn event raised by the BurnableToken contract.
type BurnableTokenBurn struct {
	Burned common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burned indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterBurn(opts *bind.FilterOpts, burned []common.Address) (*BurnableTokenBurnIterator, error) {

	var burnedRule []interface{}
	for _, burnedItem := range burned {
		burnedRule = append(burnedRule, burnedItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Burn", burnedRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenBurnIterator{contract: _BurnableToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burned indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BurnableTokenBurn, burned []common.Address) (event.Subscription, error) {

	var burnedRule []interface{}
	for _, burnedItem := range burned {
		burnedRule = append(burnedRule, burnedItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Burn", burnedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenBurn)
				if err := _BurnableToken.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BurnableTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BurnableToken contract.
type BurnableTokenMintIterator struct {
	Event *BurnableTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BurnableTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BurnableTokenMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BurnableTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenMint represents a Mint event raised by the BurnableToken contract.
type BurnableTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*BurnableTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMintIterator{contract: _BurnableToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BurnableTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenMint)
				if err := _BurnableToken.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BurnableTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the BurnableToken contract.
type BurnableTokenMintFinishedIterator struct {
	Event *BurnableTokenMintFinished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BurnableTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenMintFinished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BurnableTokenMintFinished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BurnableTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenMintFinished represents a MintFinished event raised by the BurnableToken contract.
type BurnableTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_BurnableToken *BurnableTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*BurnableTokenMintFinishedIterator, error) {

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &BurnableTokenMintFinishedIterator{contract: _BurnableToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_BurnableToken *BurnableTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *BurnableTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenMintFinished)
				if err := _BurnableToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BurnableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BurnableToken contract.
type BurnableTokenOwnershipTransferredIterator struct {
	Event *BurnableTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BurnableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BurnableTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BurnableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BurnableToken contract.
type BurnableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BurnableToken *BurnableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BurnableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenOwnershipTransferredIterator{contract: _BurnableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BurnableToken *BurnableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenOwnershipTransferred)
				if err := _BurnableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// BurnableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BurnableToken contract.
type BurnableTokenTransferIterator struct {
	Event *BurnableTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BurnableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BurnableTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BurnableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenTransfer represents a Transfer event raised by the BurnableToken contract.
type BurnableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransferIterator{contract: _BurnableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenTransfer)
				if err := _BurnableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DeveloperableABI is the input ABI used to generate the binding from.
const DeveloperableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addDeveloper\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_AUDITOR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAuditor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSuperuser\",\"type\":\"address\"}],\"name\":\"transferSuperuser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isDeveloper\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isSuperuser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEVELOPER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAuditor\",\"type\":\"address\"}],\"name\":\"addAuditor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_SUPERUSER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// DeveloperableBin is the compiled bytecode used for deploying new contracts.
const DeveloperableBin = `0x60c0604052600960809081527f737570657275736572000000000000000000000000000000000000000000000060a052610043903390640100000000610048810204565b61018b565b6100bf826000836040518082805190602001908083835b6020831061007e5780518252601f19909201916020918201910161005f565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505064010000000061016681026108b41704565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610128578181015183820152602001610110565b50505050905090810190601f1680156101555780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6109478061019a6000396000f3006080604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c81146100b3578063217fe6c61461011c5780632b7d7489146101975780633f029aaf146101ac57806349b905571461023657806357c393fa146102575780635eca4a7014610278578063bceee05e14610299578063e31cd22e146102ba578063e429cef1146102cf578063ebb4f484146102f0575b600080fd5b3480156100bf57600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261011a958335600160a060020a03169536956044949193909101919081908401838280828437509497506103059650505050505050565b005b34801561012857600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610183958335600160a060020a03169536956044949193909101919081908401838280828437509497506103739650505050505050565b604080519115158252519081900360200190f35b3480156101a357600080fd5b5061011a6103e6565b3480156101b857600080fd5b506101c1610433565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101fb5781810151838201526020016101e3565b50505050905090810190601f1680156102285780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561024257600080fd5b50610183600160a060020a036004351661046a565b34801561026357600080fd5b5061011a600160a060020a03600435166104b1565b34801561028457600080fd5b50610183600160a060020a0360043516610550565b3480156102a557600080fd5b50610183600160a060020a0360043516610591565b3480156102c657600080fd5b506101c16105c0565b3480156102db57600080fd5b5061011a600160a060020a03600435166105f7565b3480156102fc57600080fd5b506101c1610678565b61036f826000836040518082805190602001908083835b6020831061033b5780518252601f19909201916020918201910161031c565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505061069d565b5050565b60006103df836000846040518082805190602001908083835b602083106103ab5780518252601f19909201916020918201910161038c565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929150506106b2565b9392505050565b3315156103f257600080fd5b610431336040805190810160405280600981526020017f646576656c6f70657200000000000000000000000000000000000000000000008152506106d1565b565b60408051808201909152600781527f61756469746f7200000000000000000000000000000000000000000000000000602082015281565b60006104ab826040805190810160405280600781526020017f61756469746f7200000000000000000000000000000000000000000000000000815250610373565b92915050565b6104de336040805190810160405280600981526020016000805160206108fc833981519152815250610305565b600160a060020a03811615156104f357600080fd5b610520336040805190810160405280600981526020016000805160206108fc8339815191528152506107e2565b61054d816040805190810160405280600981526020016000805160206108fc8339815191528152506106d1565b50565b60006104ab826040805190810160405280600981526020017f646576656c6f7065720000000000000000000000000000000000000000000000815250610373565b60006104ab826040805190810160405280600981526020016000805160206108fc833981519152815250610373565b60408051808201909152600981527f646576656c6f7065720000000000000000000000000000000000000000000000602082015281565b610624336040805190810160405280600981526020016000805160206108fc833981519152815250610305565b600160a060020a038116151561063957600080fd5b61054d816040805190810160405280600781526020017f61756469746f72000000000000000000000000000000000000000000000000008152506106d1565b60408051808201909152600981526000805160206108fc833981519152602082015281565b6106a782826106b2565b151561036f57600080fd5b600160a060020a03166000908152602091909152604090205460ff1690565b61073b826000836040518082805190602001908083835b602083106107075780518252601f1990920191602091820191016106e8565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929150506108b4565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b838110156107a457818101518382015260200161078c565b50505050905090810190601f1680156107d15780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b61084c826000836040518082805190602001908083835b602083106108185780518252601f1990920191602091820191016107f9565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929150506108d9565b81600160a060020a03167fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a82604051808060200182810382528381815181526020019150805190602001908083836000838110156107a457818101518382015260200161078c565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b600160a060020a0316600090815260209190915260409020805460ff1916905556007375706572757365720000000000000000000000000000000000000000000000a165627a7a72305820e84bedfa18533c06858323958e07be01a891a50600e8f3d29fba7867c3f0a1cb0029`

// DeployDeveloperable deploys a new Ethereum contract, binding an instance of Developerable to it.
func DeployDeveloperable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Developerable, error) {
	parsed, err := abi.JSON(strings.NewReader(DeveloperableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DeveloperableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Developerable{DeveloperableCaller: DeveloperableCaller{contract: contract}, DeveloperableTransactor: DeveloperableTransactor{contract: contract}, DeveloperableFilterer: DeveloperableFilterer{contract: contract}}, nil
}

// Developerable is an auto generated Go binding around an Ethereum contract.
type Developerable struct {
	DeveloperableCaller     // Read-only binding to the contract
	DeveloperableTransactor // Write-only binding to the contract
	DeveloperableFilterer   // Log filterer for contract events
}

// DeveloperableCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeveloperableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeveloperableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeveloperableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeveloperableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeveloperableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeveloperableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeveloperableSession struct {
	Contract     *Developerable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeveloperableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeveloperableCallerSession struct {
	Contract *DeveloperableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DeveloperableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeveloperableTransactorSession struct {
	Contract     *DeveloperableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DeveloperableRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeveloperableRaw struct {
	Contract *Developerable // Generic contract binding to access the raw methods on
}

// DeveloperableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeveloperableCallerRaw struct {
	Contract *DeveloperableCaller // Generic read-only contract binding to access the raw methods on
}

// DeveloperableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeveloperableTransactorRaw struct {
	Contract *DeveloperableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeveloperable creates a new instance of Developerable, bound to a specific deployed contract.
func NewDeveloperable(address common.Address, backend bind.ContractBackend) (*Developerable, error) {
	contract, err := bindDeveloperable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Developerable{DeveloperableCaller: DeveloperableCaller{contract: contract}, DeveloperableTransactor: DeveloperableTransactor{contract: contract}, DeveloperableFilterer: DeveloperableFilterer{contract: contract}}, nil
}

// NewDeveloperableCaller creates a new read-only instance of Developerable, bound to a specific deployed contract.
func NewDeveloperableCaller(address common.Address, caller bind.ContractCaller) (*DeveloperableCaller, error) {
	contract, err := bindDeveloperable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeveloperableCaller{contract: contract}, nil
}

// NewDeveloperableTransactor creates a new write-only instance of Developerable, bound to a specific deployed contract.
func NewDeveloperableTransactor(address common.Address, transactor bind.ContractTransactor) (*DeveloperableTransactor, error) {
	contract, err := bindDeveloperable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeveloperableTransactor{contract: contract}, nil
}

// NewDeveloperableFilterer creates a new log filterer instance of Developerable, bound to a specific deployed contract.
func NewDeveloperableFilterer(address common.Address, filterer bind.ContractFilterer) (*DeveloperableFilterer, error) {
	contract, err := bindDeveloperable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeveloperableFilterer{contract: contract}, nil
}

// bindDeveloperable binds a generic wrapper to an already deployed contract.
func bindDeveloperable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeveloperableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Developerable *DeveloperableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Developerable.Contract.DeveloperableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Developerable *DeveloperableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Developerable.Contract.DeveloperableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Developerable *DeveloperableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Developerable.Contract.DeveloperableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Developerable *DeveloperableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Developerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Developerable *DeveloperableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Developerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Developerable *DeveloperableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Developerable.Contract.contract.Transact(opts, method, params...)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_Developerable *DeveloperableCaller) ROLEAUDITOR(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Developerable.contract.Call(opts, out, "ROLE_AUDITOR")
	return *ret0, err
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_Developerable *DeveloperableSession) ROLEAUDITOR() (string, error) {
	return _Developerable.Contract.ROLEAUDITOR(&_Developerable.CallOpts)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_Developerable *DeveloperableCallerSession) ROLEAUDITOR() (string, error) {
	return _Developerable.Contract.ROLEAUDITOR(&_Developerable.CallOpts)
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_Developerable *DeveloperableCaller) ROLEDEVELOPER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Developerable.contract.Call(opts, out, "ROLE_DEVELOPER")
	return *ret0, err
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_Developerable *DeveloperableSession) ROLEDEVELOPER() (string, error) {
	return _Developerable.Contract.ROLEDEVELOPER(&_Developerable.CallOpts)
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_Developerable *DeveloperableCallerSession) ROLEDEVELOPER() (string, error) {
	return _Developerable.Contract.ROLEDEVELOPER(&_Developerable.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Developerable *DeveloperableCaller) ROLESUPERUSER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Developerable.contract.Call(opts, out, "ROLE_SUPERUSER")
	return *ret0, err
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Developerable *DeveloperableSession) ROLESUPERUSER() (string, error) {
	return _Developerable.Contract.ROLESUPERUSER(&_Developerable.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Developerable *DeveloperableCallerSession) ROLESUPERUSER() (string, error) {
	return _Developerable.Contract.ROLESUPERUSER(&_Developerable.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Developerable *DeveloperableCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _Developerable.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Developerable *DeveloperableSession) CheckRole(_operator common.Address, _role string) error {
	return _Developerable.Contract.CheckRole(&_Developerable.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Developerable *DeveloperableCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _Developerable.Contract.CheckRole(&_Developerable.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Developerable *DeveloperableCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Developerable.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Developerable *DeveloperableSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Developerable.Contract.HasRole(&_Developerable.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Developerable *DeveloperableCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Developerable.Contract.HasRole(&_Developerable.CallOpts, _operator, _role)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_Developerable *DeveloperableCaller) IsAuditor(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Developerable.contract.Call(opts, out, "isAuditor", _addr)
	return *ret0, err
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_Developerable *DeveloperableSession) IsAuditor(_addr common.Address) (bool, error) {
	return _Developerable.Contract.IsAuditor(&_Developerable.CallOpts, _addr)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_Developerable *DeveloperableCallerSession) IsAuditor(_addr common.Address) (bool, error) {
	return _Developerable.Contract.IsAuditor(&_Developerable.CallOpts, _addr)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_Developerable *DeveloperableCaller) IsDeveloper(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Developerable.contract.Call(opts, out, "isDeveloper", _addr)
	return *ret0, err
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_Developerable *DeveloperableSession) IsDeveloper(_addr common.Address) (bool, error) {
	return _Developerable.Contract.IsDeveloper(&_Developerable.CallOpts, _addr)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_Developerable *DeveloperableCallerSession) IsDeveloper(_addr common.Address) (bool, error) {
	return _Developerable.Contract.IsDeveloper(&_Developerable.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Developerable *DeveloperableCaller) IsSuperuser(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Developerable.contract.Call(opts, out, "isSuperuser", _addr)
	return *ret0, err
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Developerable *DeveloperableSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _Developerable.Contract.IsSuperuser(&_Developerable.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Developerable *DeveloperableCallerSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _Developerable.Contract.IsSuperuser(&_Developerable.CallOpts, _addr)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_Developerable *DeveloperableTransactor) AddAuditor(opts *bind.TransactOpts, _newAuditor common.Address) (*types.Transaction, error) {
	return _Developerable.contract.Transact(opts, "addAuditor", _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_Developerable *DeveloperableSession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _Developerable.Contract.AddAuditor(&_Developerable.TransactOpts, _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_Developerable *DeveloperableTransactorSession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _Developerable.Contract.AddAuditor(&_Developerable.TransactOpts, _newAuditor)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_Developerable *DeveloperableTransactor) AddDeveloper(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Developerable.contract.Transact(opts, "addDeveloper")
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_Developerable *DeveloperableSession) AddDeveloper() (*types.Transaction, error) {
	return _Developerable.Contract.AddDeveloper(&_Developerable.TransactOpts)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_Developerable *DeveloperableTransactorSession) AddDeveloper() (*types.Transaction, error) {
	return _Developerable.Contract.AddDeveloper(&_Developerable.TransactOpts)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Developerable *DeveloperableTransactor) TransferSuperuser(opts *bind.TransactOpts, _newSuperuser common.Address) (*types.Transaction, error) {
	return _Developerable.contract.Transact(opts, "transferSuperuser", _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Developerable *DeveloperableSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _Developerable.Contract.TransferSuperuser(&_Developerable.TransactOpts, _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Developerable *DeveloperableTransactorSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _Developerable.Contract.TransferSuperuser(&_Developerable.TransactOpts, _newSuperuser)
}

// DeveloperableRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the Developerable contract.
type DeveloperableRoleAddedIterator struct {
	Event *DeveloperableRoleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeveloperableRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeveloperableRoleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeveloperableRoleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeveloperableRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeveloperableRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeveloperableRoleAdded represents a RoleAdded event raised by the Developerable contract.
type DeveloperableRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Developerable *DeveloperableFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*DeveloperableRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Developerable.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DeveloperableRoleAddedIterator{contract: _Developerable.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Developerable *DeveloperableFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *DeveloperableRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Developerable.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeveloperableRoleAdded)
				if err := _Developerable.contract.UnpackLog(event, "RoleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DeveloperableRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the Developerable contract.
type DeveloperableRoleRemovedIterator struct {
	Event *DeveloperableRoleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DeveloperableRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeveloperableRoleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DeveloperableRoleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DeveloperableRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeveloperableRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeveloperableRoleRemoved represents a RoleRemoved event raised by the Developerable contract.
type DeveloperableRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Developerable *DeveloperableFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*DeveloperableRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Developerable.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DeveloperableRoleRemovedIterator{contract: _Developerable.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Developerable *DeveloperableFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *DeveloperableRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Developerable.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeveloperableRoleRemoved)
				if err := _Developerable.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BasicTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MintableTokenABI is the input ABI used to generate the binding from.
const MintableTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// MintableTokenBin is the compiled bytecode used for deploying new contracts.
const MintableTokenBin = `0x608060405260038054600160a860020a03191633179055610a1a806100256000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166305d2035b81146100c9578063095ea7b3146100f257806318160ddd1461011657806323b872dd1461013d57806340c10f1914610167578063661884631461018b57806370a08231146101af5780637d64bcb4146101d05780638da5cb5b146101e5578063a9059cbb14610216578063d73dd6231461023a578063dd62ed3e1461025e578063f2fde38b14610285575b600080fd5b3480156100d557600080fd5b506100de6102a8565b604080519115158252519081900360200190f35b3480156100fe57600080fd5b506100de600160a060020a03600435166024356102c9565b34801561012257600080fd5b5061012b61032f565b60408051918252519081900360200190f35b34801561014957600080fd5b506100de600160a060020a0360043581169060243516604435610335565b34801561017357600080fd5b506100de600160a060020a03600435166024356104ae565b34801561019757600080fd5b506100de600160a060020a03600435166024356105ca565b3480156101bb57600080fd5b5061012b600160a060020a03600435166106ba565b3480156101dc57600080fd5b506100de6106d5565b3480156101f157600080fd5b506101fa61077b565b60408051600160a060020a039092168252519081900360200190f35b34801561022257600080fd5b506100de600160a060020a036004351660243561078a565b34801561024657600080fd5b506100de600160a060020a036004351660243561086d565b34801561026a57600080fd5b5061012b600160a060020a0360043581169060243516610906565b34801561029157600080fd5b506102a6600160a060020a0360043516610931565b005b60035474010000000000000000000000000000000000000000900460ff1681565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b6000600160a060020a038316151561034c57600080fd5b600160a060020a03841660009081526001602052604090205482111561037157600080fd5b600160a060020a03841660009081526002602090815260408083203384529091529020548211156103a157600080fd5b600160a060020a0384166000908152600160205260409020546103ca908363ffffffff6109c616565b600160a060020a0380861660009081526001602052604080822093909355908516815220546103ff908363ffffffff6109d816565b600160a060020a038085166000908152600160209081526040808320949094559187168152600282528281203382529091522054610443908363ffffffff6109c616565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b600354600090600160a060020a031633146104c857600080fd5b60035474010000000000000000000000000000000000000000900460ff16156104f057600080fd5b600054610503908363ffffffff6109d816565b6000908155600160a060020a03841681526001602052604090205461052e908363ffffffff6109d816565b600160a060020a038416600081815260016020908152604091829020939093558051858152905191927f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592918290030190a2604080518381529051600160a060020a038516916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120548083111561061f57336000908152600260209081526040808320600160a060020a0388168452909152812055610654565b61062f818463ffffffff6109c616565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526001602052604090205490565b600354600090600160a060020a031633146106ef57600080fd5b60035474010000000000000000000000000000000000000000900460ff161561071757600080fd5b6003805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790556040517fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0890600090a150600190565b600354600160a060020a031681565b6000600160a060020a03831615156107a157600080fd5b336000908152600160205260409020548211156107bd57600080fd5b336000908152600160205260409020546107dd908363ffffffff6109c616565b3360009081526001602052604080822092909255600160a060020a0385168152205461080f908363ffffffff6109d816565b600160a060020a0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120546108a1908363ffffffff6109d816565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b600354600160a060020a0316331461094857600080fd5b600160a060020a038116151561095d57600080fd5b600354604051600160a060020a038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a36003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000828211156109d257fe5b50900390565b6000828201838110156109e757fe5b93925050505600a165627a7a72305820c60b465e0d2f2cfd078954778664dfb01c9bc507509d70fb7a010079751217190029`

// DeployMintableToken deploys a new Ethereum contract, binding an instance of MintableToken to it.
func DeployMintableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MintableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MintableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MintableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MintableToken{MintableTokenCaller: MintableTokenCaller{contract: contract}, MintableTokenTransactor: MintableTokenTransactor{contract: contract}, MintableTokenFilterer: MintableTokenFilterer{contract: contract}}, nil
}

// MintableToken is an auto generated Go binding around an Ethereum contract.
type MintableToken struct {
	MintableTokenCaller     // Read-only binding to the contract
	MintableTokenTransactor // Write-only binding to the contract
	MintableTokenFilterer   // Log filterer for contract events
}

// MintableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintableTokenSession struct {
	Contract     *MintableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintableTokenCallerSession struct {
	Contract *MintableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MintableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintableTokenTransactorSession struct {
	Contract     *MintableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MintableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintableTokenRaw struct {
	Contract *MintableToken // Generic contract binding to access the raw methods on
}

// MintableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintableTokenCallerRaw struct {
	Contract *MintableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MintableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintableTokenTransactorRaw struct {
	Contract *MintableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintableToken creates a new instance of MintableToken, bound to a specific deployed contract.
func NewMintableToken(address common.Address, backend bind.ContractBackend) (*MintableToken, error) {
	contract, err := bindMintableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintableToken{MintableTokenCaller: MintableTokenCaller{contract: contract}, MintableTokenTransactor: MintableTokenTransactor{contract: contract}, MintableTokenFilterer: MintableTokenFilterer{contract: contract}}, nil
}

// NewMintableTokenCaller creates a new read-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenCaller(address common.Address, caller bind.ContractCaller) (*MintableTokenCaller, error) {
	contract, err := bindMintableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintableTokenCaller{contract: contract}, nil
}

// NewMintableTokenTransactor creates a new write-only instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MintableTokenTransactor, error) {
	contract, err := bindMintableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintableTokenTransactor{contract: contract}, nil
}

// NewMintableTokenFilterer creates a new log filterer instance of MintableToken, bound to a specific deployed contract.
func NewMintableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MintableTokenFilterer, error) {
	contract, err := bindMintableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintableTokenFilterer{contract: contract}, nil
}

// bindMintableToken binds a generic wrapper to an already deployed contract.
func bindMintableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.MintableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.MintableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableToken *MintableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MintableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableToken *MintableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableToken *MintableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_MintableToken *MintableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_MintableToken *MintableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableToken.Contract.Allowance(&_MintableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_MintableToken *MintableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableToken.Contract.Allowance(&_MintableToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MintableToken *MintableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MintableToken *MintableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MintableToken.Contract.BalanceOf(&_MintableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_MintableToken *MintableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _MintableToken.Contract.BalanceOf(&_MintableToken.CallOpts, _owner)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MintableToken *MintableTokenCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MintableToken *MintableTokenSession) MintingFinished() (bool, error) {
	return _MintableToken.Contract.MintingFinished(&_MintableToken.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_MintableToken *MintableTokenCallerSession) MintingFinished() (bool, error) {
	return _MintableToken.Contract.MintingFinished(&_MintableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MintableToken *MintableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MintableToken *MintableTokenSession) Owner() (common.Address, error) {
	return _MintableToken.Contract.Owner(&_MintableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MintableToken *MintableTokenCallerSession) Owner() (common.Address, error) {
	return _MintableToken.Contract.Owner(&_MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MintableToken *MintableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MintableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MintableToken *MintableTokenSession) TotalSupply() (*big.Int, error) {
	return _MintableToken.Contract.TotalSupply(&_MintableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MintableToken *MintableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MintableToken.Contract.TotalSupply(&_MintableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Approve(&_MintableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Approve(&_MintableToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_MintableToken *MintableTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.DecreaseApproval(&_MintableToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.DecreaseApproval(&_MintableToken.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenSession) FinishMinting() (*types.Transaction, error) {
	return _MintableToken.Contract.FinishMinting(&_MintableToken.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_MintableToken *MintableTokenTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _MintableToken.Contract.FinishMinting(&_MintableToken.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_MintableToken *MintableTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.IncreaseApproval(&_MintableToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.IncreaseApproval(&_MintableToken.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MintableToken *MintableTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Mint(&_MintableToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Transfer(&_MintableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.Transfer(&_MintableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferFrom(&_MintableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_MintableToken *MintableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferFrom(&_MintableToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MintableToken *MintableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MintableToken *MintableTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferOwnership(&_MintableToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MintableToken *MintableTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintableToken.Contract.TransferOwnership(&_MintableToken.TransactOpts, newOwner)
}

// MintableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MintableToken contract.
type MintableTokenApprovalIterator struct {
	Event *MintableTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MintableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MintableTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MintableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenApproval represents a Approval event raised by the MintableToken contract.
type MintableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_MintableToken *MintableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MintableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenApprovalIterator{contract: _MintableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_MintableToken *MintableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MintableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenApproval)
				if err := _MintableToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MintableTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the MintableToken contract.
type MintableTokenMintIterator struct {
	Event *MintableTokenMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MintableTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MintableTokenMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MintableTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenMint represents a Mint event raised by the MintableToken contract.
type MintableTokenMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_MintableToken *MintableTokenFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*MintableTokenMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenMintIterator{contract: _MintableToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_MintableToken *MintableTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *MintableTokenMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenMint)
				if err := _MintableToken.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MintableTokenMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the MintableToken contract.
type MintableTokenMintFinishedIterator struct {
	Event *MintableTokenMintFinished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MintableTokenMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenMintFinished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MintableTokenMintFinished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MintableTokenMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenMintFinished represents a MintFinished event raised by the MintableToken contract.
type MintableTokenMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_MintableToken *MintableTokenFilterer) FilterMintFinished(opts *bind.FilterOpts) (*MintableTokenMintFinishedIterator, error) {

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &MintableTokenMintFinishedIterator{contract: _MintableToken.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_MintableToken *MintableTokenFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *MintableTokenMintFinished) (event.Subscription, error) {

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenMintFinished)
				if err := _MintableToken.contract.UnpackLog(event, "MintFinished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MintableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MintableToken contract.
type MintableTokenOwnershipTransferredIterator struct {
	Event *MintableTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MintableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MintableTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MintableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the MintableToken contract.
type MintableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MintableToken *MintableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenOwnershipTransferredIterator{contract: _MintableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MintableToken *MintableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MintableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenOwnershipTransferred)
				if err := _MintableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MintableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MintableToken contract.
type MintableTokenTransferIterator struct {
	Event *MintableTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MintableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MintableTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MintableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableTokenTransfer represents a Transfer event raised by the MintableToken contract.
type MintableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_MintableToken *MintableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MintableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MintableTokenTransferIterator{contract: _MintableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_MintableToken *MintableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MintableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableTokenTransfer)
				if err := _MintableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610173806100326000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b14610081575b600080fd5b34801561005c57600080fd5b506100656100a4565b60408051600160a060020a039092168252519081900360200190f35b34801561008d57600080fd5b506100a2600160a060020a03600435166100b3565b005b600054600160a060020a031681565b600054600160a060020a031633146100ca57600080fd5b600160a060020a03811615156100df57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058200a48a1402449ffcfcc72a5c8f45742ed8f168983070c6fb0aa6b1926c2df4ad00029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RBACABI is the input ABI used to generate the binding from.
const RBACABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// RBACBin is the compiled bytecode used for deploying new contracts.
const RBACBin = `0x608060405234801561001057600080fd5b5061029c806100206000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c8114610050578063217fe6c6146100c6575b600080fd5b34801561005c57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100c495833573ffffffffffffffffffffffffffffffffffffffff1695369560449491939091019190819084018382808284375094975061014e9650505050505050565b005b3480156100d257600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261013a95833573ffffffffffffffffffffffffffffffffffffffff169536956044949193909101919081908401838280828437509497506101bc9650505050505050565b604080519115158252519081900360200190f35b6101b8826000836040518082805190602001908083835b602083106101845780518252601f199092019160209182019101610165565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505061022f565b5050565b6000610228836000846040518082805190602001908083835b602083106101f45780518252601f1990920191602091820191016101d5565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610244565b9392505050565b6102398282610244565b15156101b857600080fd5b73ffffffffffffffffffffffffffffffffffffffff166000908152602091909152604090205460ff16905600a165627a7a7230582051139924331553082811337294c8cdc873b0b880db07b853963f13c50281df4f0029`

// DeployRBAC deploys a new Ethereum contract, binding an instance of RBAC to it.
func DeployRBAC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RBAC, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RBACBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RBAC{RBACCaller: RBACCaller{contract: contract}, RBACTransactor: RBACTransactor{contract: contract}, RBACFilterer: RBACFilterer{contract: contract}}, nil
}

// RBAC is an auto generated Go binding around an Ethereum contract.
type RBAC struct {
	RBACCaller     // Read-only binding to the contract
	RBACTransactor // Write-only binding to the contract
	RBACFilterer   // Log filterer for contract events
}

// RBACCaller is an auto generated read-only Go binding around an Ethereum contract.
type RBACCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RBACTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RBACFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RBACSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RBACSession struct {
	Contract     *RBAC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RBACCallerSession struct {
	Contract *RBACCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RBACTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RBACTransactorSession struct {
	Contract     *RBACTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RBACRaw is an auto generated low-level Go binding around an Ethereum contract.
type RBACRaw struct {
	Contract *RBAC // Generic contract binding to access the raw methods on
}

// RBACCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RBACCallerRaw struct {
	Contract *RBACCaller // Generic read-only contract binding to access the raw methods on
}

// RBACTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RBACTransactorRaw struct {
	Contract *RBACTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRBAC creates a new instance of RBAC, bound to a specific deployed contract.
func NewRBAC(address common.Address, backend bind.ContractBackend) (*RBAC, error) {
	contract, err := bindRBAC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RBAC{RBACCaller: RBACCaller{contract: contract}, RBACTransactor: RBACTransactor{contract: contract}, RBACFilterer: RBACFilterer{contract: contract}}, nil
}

// NewRBACCaller creates a new read-only instance of RBAC, bound to a specific deployed contract.
func NewRBACCaller(address common.Address, caller bind.ContractCaller) (*RBACCaller, error) {
	contract, err := bindRBAC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RBACCaller{contract: contract}, nil
}

// NewRBACTransactor creates a new write-only instance of RBAC, bound to a specific deployed contract.
func NewRBACTransactor(address common.Address, transactor bind.ContractTransactor) (*RBACTransactor, error) {
	contract, err := bindRBAC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RBACTransactor{contract: contract}, nil
}

// NewRBACFilterer creates a new log filterer instance of RBAC, bound to a specific deployed contract.
func NewRBACFilterer(address common.Address, filterer bind.ContractFilterer) (*RBACFilterer, error) {
	contract, err := bindRBAC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RBACFilterer{contract: contract}, nil
}

// bindRBAC binds a generic wrapper to an already deployed contract.
func bindRBAC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RBACABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBAC *RBACRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBAC.Contract.RBACCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBAC *RBACRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBAC.Contract.RBACTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBAC *RBACRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBAC.Contract.RBACTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RBAC *RBACCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RBAC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RBAC *RBACTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RBAC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RBAC *RBACTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RBAC.Contract.contract.Transact(opts, method, params...)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_RBAC *RBACCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _RBAC.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_RBAC *RBACSession) CheckRole(_operator common.Address, _role string) error {
	return _RBAC.Contract.CheckRole(&_RBAC.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_RBAC *RBACCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _RBAC.Contract.CheckRole(&_RBAC.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_RBAC *RBACCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RBAC.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_RBAC *RBACSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _RBAC.Contract.HasRole(&_RBAC.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_RBAC *RBACCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _RBAC.Contract.HasRole(&_RBAC.CallOpts, _operator, _role)
}

// RBACRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the RBAC contract.
type RBACRoleAddedIterator struct {
	Event *RBACRoleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RBACRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RBACRoleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RBACRoleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RBACRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RBACRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RBACRoleAdded represents a RoleAdded event raised by the RBAC contract.
type RBACRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_RBAC *RBACFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*RBACRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RBAC.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &RBACRoleAddedIterator{contract: _RBAC.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_RBAC *RBACFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *RBACRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RBAC.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RBACRoleAdded)
				if err := _RBAC.contract.UnpackLog(event, "RoleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RBACRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the RBAC contract.
type RBACRoleRemovedIterator struct {
	Event *RBACRoleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RBACRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RBACRoleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RBACRoleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RBACRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RBACRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RBACRoleRemoved represents a RoleRemoved event raised by the RBAC contract.
type RBACRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_RBAC *RBACFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*RBACRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RBAC.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &RBACRoleRemovedIterator{contract: _RBAC.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_RBAC *RBACFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *RBACRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RBAC.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RBACRoleRemoved)
				if err := _RBAC.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBin is the compiled bytecode used for deploying new contracts.
const RolesBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058201f3f63b29fa126f492efdc4230c520464db75465018d399f6a3d630faa5085140029`

// DeployRoles deploys a new Ethereum contract, binding an instance of Roles to it.
func DeployRoles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Roles, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// Roles is an auto generated Go binding around an Ethereum contract.
type Roles struct {
	RolesCaller     // Read-only binding to the contract
	RolesTransactor // Write-only binding to the contract
	RolesFilterer   // Log filterer for contract events
}

// RolesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolesSession struct {
	Contract     *Roles            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolesCallerSession struct {
	Contract *RolesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RolesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolesTransactorSession struct {
	Contract     *RolesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolesRaw struct {
	Contract *Roles // Generic contract binding to access the raw methods on
}

// RolesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolesCallerRaw struct {
	Contract *RolesCaller // Generic read-only contract binding to access the raw methods on
}

// RolesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolesTransactorRaw struct {
	Contract *RolesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoles creates a new instance of Roles, bound to a specific deployed contract.
func NewRoles(address common.Address, backend bind.ContractBackend) (*Roles, error) {
	contract, err := bindRoles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// NewRolesCaller creates a new read-only instance of Roles, bound to a specific deployed contract.
func NewRolesCaller(address common.Address, caller bind.ContractCaller) (*RolesCaller, error) {
	contract, err := bindRoles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolesCaller{contract: contract}, nil
}

// NewRolesTransactor creates a new write-only instance of Roles, bound to a specific deployed contract.
func NewRolesTransactor(address common.Address, transactor bind.ContractTransactor) (*RolesTransactor, error) {
	contract, err := bindRoles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolesTransactor{contract: contract}, nil
}

// NewRolesFilterer creates a new log filterer instance of Roles, bound to a specific deployed contract.
func NewRolesFilterer(address common.Address, filterer bind.ContractFilterer) (*RolesFilterer, error) {
	contract, err := bindRoles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolesFilterer{contract: contract}, nil
}

// bindRoles binds a generic wrapper to an already deployed contract.
func bindRoles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.RolesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
const SafeERC20Bin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820a841a449eddaa841d711a7c004c387bab8e2c2743db3ffef70771f4d36c30efe0029`

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820d5f900b7bdc5efb3e14346567b9883a5dc6ca366a0be3bab7f09b019731d122b0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x608060405234801561001057600080fd5b506106ba806100206000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461009257806318160ddd146100ca57806323b872dd146100f1578063661884631461011b57806370a082311461013f578063a9059cbb14610160578063d73dd62314610184578063dd62ed3e146101a8575b600080fd5b34801561009e57600080fd5b506100b6600160a060020a03600435166024356101cf565b604080519115158252519081900360200190f35b3480156100d657600080fd5b506100df610235565b60408051918252519081900360200190f35b3480156100fd57600080fd5b506100b6600160a060020a036004358116906024351660443561023b565b34801561012757600080fd5b506100b6600160a060020a03600435166024356103b4565b34801561014b57600080fd5b506100df600160a060020a03600435166104a4565b34801561016c57600080fd5b506100b6600160a060020a03600435166024356104bf565b34801561019057600080fd5b506100b6600160a060020a03600435166024356105a2565b3480156101b457600080fd5b506100df600160a060020a036004358116906024351661063b565b336000818152600260209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60005481565b6000600160a060020a038316151561025257600080fd5b600160a060020a03841660009081526001602052604090205482111561027757600080fd5b600160a060020a03841660009081526002602090815260408083203384529091529020548211156102a757600080fd5b600160a060020a0384166000908152600160205260409020546102d0908363ffffffff61066616565b600160a060020a038086166000908152600160205260408082209390935590851681522054610305908363ffffffff61067816565b600160a060020a038085166000908152600160209081526040808320949094559187168152600282528281203382529091522054610349908363ffffffff61066616565b600160a060020a03808616600081815260026020908152604080832033845282529182902094909455805186815290519287169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060019392505050565b336000908152600260209081526040808320600160a060020a03861684529091528120548083111561040957336000908152600260209081526040808320600160a060020a038816845290915281205561043e565b610419818463ffffffff61066616565b336000908152600260209081526040808320600160a060020a03891684529091529020555b336000818152600260209081526040808320600160a060020a0389168085529083529281902054815190815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b600160a060020a031660009081526001602052604090205490565b6000600160a060020a03831615156104d657600080fd5b336000908152600160205260409020548211156104f257600080fd5b33600090815260016020526040902054610512908363ffffffff61066616565b3360009081526001602052604080822092909255600160a060020a03851681522054610544908363ffffffff61067816565b600160a060020a0384166000818152600160209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b336000908152600260209081526040808320600160a060020a03861684529091528120546105d6908363ffffffff61067816565b336000818152600260209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b600160a060020a03918216600090815260026020908152604080832093909416825291909152205490565b60008282111561067257fe5b50900390565b60008282018381101561068757fe5b93925050505600a165627a7a723058203a961bb8b20122ca6fe4ac49eab49b18fb1243b25d71352d52db9610180578380029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StandardTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SuperuserableABI is the input ABI used to generate the binding from.
const SuperuserableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSuperuser\",\"type\":\"address\"}],\"name\":\"transferSuperuser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isSuperuser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_SUPERUSER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// SuperuserableBin is the compiled bytecode used for deploying new contracts.
const SuperuserableBin = `0x608060405234801561001057600080fd5b5061005f336040805190810160405280600981526020017f7375706572757365720000000000000000000000000000000000000000000000815250610064640100000000026401000000009004565b6101a7565b6100db826000836040518082805190602001908083835b6020831061009a5780518252601f19909201916020918201910161007b565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505064010000000061018281026106341704565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b8381101561014457818101518382015260200161012c565b50505050905090810190601f1680156101715780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6106a5806101b66000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c8114610071578063217fe6c6146100da57806357c393fa14610155578063bceee05e14610176578063ebb4f48414610197575b600080fd5b34801561007d57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100d8958335600160a060020a03169536956044949193909101919081908401838280828437509497506102219650505050505050565b005b3480156100e657600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610141958335600160a060020a031695369560449491939091019190819084018382808284375094975061028f9650505050505050565b604080519115158252519081900360200190f35b34801561016157600080fd5b506100d8600160a060020a0360043516610302565b34801561018257600080fd5b50610141600160a060020a03600435166103a1565b3480156101a357600080fd5b506101ac6103d6565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101e65781810151838201526020016101ce565b50505050905090810190601f1680156102135780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61028b826000836040518082805190602001908083835b602083106102575780518252601f199092019160209182019101610238565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929150506103fb565b5050565b60006102fb836000846040518082805190602001908083835b602083106102c75780518252601f1990920191602091820191016102a8565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610410565b9392505050565b61032f3360408051908101604052806009815260200160008051602061065a833981519152815250610221565b600160a060020a038116151561034457600080fd5b6103713360408051908101604052806009815260200160008051602061065a83398151915281525061042f565b61039e8160408051908101604052806009815260200160008051602061065a833981519152815250610540565b50565b60006103d08260408051908101604052806009815260200160008051602061065a83398151915281525061028f565b92915050565b604080518082019091526009815260008051602061065a833981519152602082015281565b6104058282610410565b151561028b57600080fd5b600160a060020a03166000908152602091909152604090205460ff1690565b610499826000836040518082805190602001908083835b602083106104655780518252601f199092019160209182019101610446565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610612565b81600160a060020a03167fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a826040518080602001828103825283818151815260200191508051906020019080838360005b838110156105025781810151838201526020016104ea565b50505050905090810190601f16801561052f5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b6105aa826000836040518082805190602001908083835b602083106105765780518252601f199092019160209182019101610557565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610634565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b70048982604051808060200182810382528381815181526020019150805190602001908083836000838110156105025781810151838201526020016104ea565b600160a060020a0316600090815260209190915260409020805460ff19169055565b600160a060020a0316600090815260209190915260409020805460ff1916600117905556007375706572757365720000000000000000000000000000000000000000000000a165627a7a72305820507da756e0759d3c4d6a5dd2203b6acd861fe97a23ce7356474cd23f532675de0029`

// DeploySuperuserable deploys a new Ethereum contract, binding an instance of Superuserable to it.
func DeploySuperuserable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Superuserable, error) {
	parsed, err := abi.JSON(strings.NewReader(SuperuserableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SuperuserableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Superuserable{SuperuserableCaller: SuperuserableCaller{contract: contract}, SuperuserableTransactor: SuperuserableTransactor{contract: contract}, SuperuserableFilterer: SuperuserableFilterer{contract: contract}}, nil
}

// Superuserable is an auto generated Go binding around an Ethereum contract.
type Superuserable struct {
	SuperuserableCaller     // Read-only binding to the contract
	SuperuserableTransactor // Write-only binding to the contract
	SuperuserableFilterer   // Log filterer for contract events
}

// SuperuserableCaller is an auto generated read-only Go binding around an Ethereum contract.
type SuperuserableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperuserableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SuperuserableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperuserableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SuperuserableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperuserableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SuperuserableSession struct {
	Contract     *Superuserable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SuperuserableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SuperuserableCallerSession struct {
	Contract *SuperuserableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SuperuserableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SuperuserableTransactorSession struct {
	Contract     *SuperuserableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SuperuserableRaw is an auto generated low-level Go binding around an Ethereum contract.
type SuperuserableRaw struct {
	Contract *Superuserable // Generic contract binding to access the raw methods on
}

// SuperuserableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SuperuserableCallerRaw struct {
	Contract *SuperuserableCaller // Generic read-only contract binding to access the raw methods on
}

// SuperuserableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SuperuserableTransactorRaw struct {
	Contract *SuperuserableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuperuserable creates a new instance of Superuserable, bound to a specific deployed contract.
func NewSuperuserable(address common.Address, backend bind.ContractBackend) (*Superuserable, error) {
	contract, err := bindSuperuserable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Superuserable{SuperuserableCaller: SuperuserableCaller{contract: contract}, SuperuserableTransactor: SuperuserableTransactor{contract: contract}, SuperuserableFilterer: SuperuserableFilterer{contract: contract}}, nil
}

// NewSuperuserableCaller creates a new read-only instance of Superuserable, bound to a specific deployed contract.
func NewSuperuserableCaller(address common.Address, caller bind.ContractCaller) (*SuperuserableCaller, error) {
	contract, err := bindSuperuserable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SuperuserableCaller{contract: contract}, nil
}

// NewSuperuserableTransactor creates a new write-only instance of Superuserable, bound to a specific deployed contract.
func NewSuperuserableTransactor(address common.Address, transactor bind.ContractTransactor) (*SuperuserableTransactor, error) {
	contract, err := bindSuperuserable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SuperuserableTransactor{contract: contract}, nil
}

// NewSuperuserableFilterer creates a new log filterer instance of Superuserable, bound to a specific deployed contract.
func NewSuperuserableFilterer(address common.Address, filterer bind.ContractFilterer) (*SuperuserableFilterer, error) {
	contract, err := bindSuperuserable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SuperuserableFilterer{contract: contract}, nil
}

// bindSuperuserable binds a generic wrapper to an already deployed contract.
func bindSuperuserable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SuperuserableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Superuserable *SuperuserableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Superuserable.Contract.SuperuserableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Superuserable *SuperuserableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Superuserable.Contract.SuperuserableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Superuserable *SuperuserableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Superuserable.Contract.SuperuserableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Superuserable *SuperuserableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Superuserable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Superuserable *SuperuserableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Superuserable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Superuserable *SuperuserableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Superuserable.Contract.contract.Transact(opts, method, params...)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Superuserable *SuperuserableCaller) ROLESUPERUSER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Superuserable.contract.Call(opts, out, "ROLE_SUPERUSER")
	return *ret0, err
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Superuserable *SuperuserableSession) ROLESUPERUSER() (string, error) {
	return _Superuserable.Contract.ROLESUPERUSER(&_Superuserable.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Superuserable *SuperuserableCallerSession) ROLESUPERUSER() (string, error) {
	return _Superuserable.Contract.ROLESUPERUSER(&_Superuserable.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Superuserable *SuperuserableCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _Superuserable.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Superuserable *SuperuserableSession) CheckRole(_operator common.Address, _role string) error {
	return _Superuserable.Contract.CheckRole(&_Superuserable.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Superuserable *SuperuserableCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _Superuserable.Contract.CheckRole(&_Superuserable.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Superuserable *SuperuserableCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Superuserable.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Superuserable *SuperuserableSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Superuserable.Contract.HasRole(&_Superuserable.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Superuserable *SuperuserableCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Superuserable.Contract.HasRole(&_Superuserable.CallOpts, _operator, _role)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Superuserable *SuperuserableCaller) IsSuperuser(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Superuserable.contract.Call(opts, out, "isSuperuser", _addr)
	return *ret0, err
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Superuserable *SuperuserableSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _Superuserable.Contract.IsSuperuser(&_Superuserable.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Superuserable *SuperuserableCallerSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _Superuserable.Contract.IsSuperuser(&_Superuserable.CallOpts, _addr)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Superuserable *SuperuserableTransactor) TransferSuperuser(opts *bind.TransactOpts, _newSuperuser common.Address) (*types.Transaction, error) {
	return _Superuserable.contract.Transact(opts, "transferSuperuser", _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Superuserable *SuperuserableSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _Superuserable.Contract.TransferSuperuser(&_Superuserable.TransactOpts, _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Superuserable *SuperuserableTransactorSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _Superuserable.Contract.TransferSuperuser(&_Superuserable.TransactOpts, _newSuperuser)
}

// SuperuserableRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the Superuserable contract.
type SuperuserableRoleAddedIterator struct {
	Event *SuperuserableRoleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperuserableRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperuserableRoleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperuserableRoleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperuserableRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperuserableRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperuserableRoleAdded represents a RoleAdded event raised by the Superuserable contract.
type SuperuserableRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Superuserable *SuperuserableFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*SuperuserableRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Superuserable.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SuperuserableRoleAddedIterator{contract: _Superuserable.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Superuserable *SuperuserableFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *SuperuserableRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Superuserable.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperuserableRoleAdded)
				if err := _Superuserable.contract.UnpackLog(event, "RoleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SuperuserableRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the Superuserable contract.
type SuperuserableRoleRemovedIterator struct {
	Event *SuperuserableRoleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SuperuserableRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperuserableRoleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SuperuserableRoleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SuperuserableRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperuserableRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperuserableRoleRemoved represents a RoleRemoved event raised by the Superuserable contract.
type SuperuserableRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Superuserable *SuperuserableFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*SuperuserableRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Superuserable.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SuperuserableRoleRemovedIterator{contract: _Superuserable.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Superuserable *SuperuserableFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *SuperuserableRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Superuserable.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperuserableRoleRemoved)
				if err := _Superuserable.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UserableABI is the input ABI used to generate the binding from.
const UserableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addDeveloper\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_AUDITOR\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isAuditor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newSuperuser\",\"type\":\"address\"}],\"name\":\"transferSuperuser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isDeveloper\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"signUp\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"login\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isSuperuser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_DEVELOPER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newAuditor\",\"type\":\"address\"}],\"name\":\"addAuditor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_SUPERUSER\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"role\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"}]"

// UserableBin is the compiled bytecode used for deploying new contracts.
const UserableBin = `0x60c0604052600960809081527f737570657275736572000000000000000000000000000000000000000000000060a052610043903390640100000000610048810204565b61018b565b6100bf826000836040518082805190602001908083835b6020831061007e5780518252601f19909201916020918201910161005f565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929150506401000000006101668102610b311704565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610128578181015183820152602001610110565b50505050905090810190601f1680156101555780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b610c5c8061019a6000396000f3006080604052600436106100c45763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630988ca8c81146100c9578063217fe6c6146101325780632b7d7489146101ad5780633f029aaf146101c257806349b905571461024c57806357c393fa1461026d5780635eca4a701461028e578063adcfb4fc146102af578063b34e97e814610308578063bceee05e1461031d578063e31cd22e1461033e578063e429cef114610353578063ebb4f48414610374575b600080fd5b3480156100d557600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610130958335600160a060020a03169536956044949193909101919081908401838280828437509497506103899650505050505050565b005b34801561013e57600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610199958335600160a060020a03169536956044949193909101919081908401838280828437509497506103f79650505050505050565b604080519115158252519081900360200190f35b3480156101b957600080fd5b5061013061046a565b3480156101ce57600080fd5b506101d76104b7565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102115781810151838201526020016101f9565b50505050905090810190601f16801561023e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561025857600080fd5b50610199600160a060020a03600435166104ee565b34801561027957600080fd5b50610130600160a060020a0360043516610535565b34801561029a57600080fd5b50610199600160a060020a03600435166105d4565b3480156102bb57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101d79436949293602493928401919081908401838280828437509497506106159650505050505050565b34801561031457600080fd5b506101d761073c565b34801561032957600080fd5b50610199600160a060020a036004351661080e565b34801561034a57600080fd5b506101d761083d565b34801561035f57600080fd5b50610130600160a060020a0360043516610874565b34801561038057600080fd5b506101d76108f5565b6103f3826000836040518082805190602001908083835b602083106103bf5780518252601f1990920191602091820191016103a0565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505061091a565b5050565b6000610463836000846040518082805190602001908083835b6020831061042f5780518252601f199092019160209182019101610410565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092209291505061092f565b9392505050565b33151561047657600080fd5b6104b5336040805190810160405280600981526020017f646576656c6f706572000000000000000000000000000000000000000000000081525061094e565b565b60408051808201909152600781527f61756469746f7200000000000000000000000000000000000000000000000000602082015281565b600061052f826040805190810160405280600781526020017f61756469746f72000000000000000000000000000000000000000000000000008152506103f7565b92915050565b61056233604080519081016040528060098152602001600080516020610c11833981519152815250610389565b600160a060020a038116151561057757600080fd5b6105a433604080519081016040528060098152602001600080516020610c11833981519152815250610a5f565b6105d181604080519081016040528060098152602001600080516020610c1183398151915281525061094e565b50565b600061052f826040805190810160405280600981526020017f646576656c6f70657200000000000000000000000000000000000000000000008152506103f7565b6060816000815111151561062857600080fd5b3360009081526001602081905260409091205460029181161561010002600019011604151561067357336000908152600160209081526040909120845161067192860190610b78565b505b61067c3361080e565b15801561068f575061068d336104ee565b155b1561069c5761069c61046a565b3360009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801561072f5780601f106107045761010080835404028352916020019161072f565b820191906000526020600020905b81548152906001019060200180831161071257829003601f168201915b5050505050915050919050565b3360009081526001602081905260408220546060926002928216156101000260001901909116919091041161077057600080fd5b3360009081526001602081815260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156108035780601f106107d857610100808354040283529160200191610803565b820191906000526020600020905b8154815290600101906020018083116107e657829003601f168201915b505050505090505b90565b600061052f82604080519081016040528060098152602001600080516020610c118339815191528152506103f7565b60408051808201909152600981527f646576656c6f7065720000000000000000000000000000000000000000000000602082015281565b6108a133604080519081016040528060098152602001600080516020610c11833981519152815250610389565b600160a060020a03811615156108b657600080fd5b6105d1816040805190810160405280600781526020017f61756469746f720000000000000000000000000000000000000000000000000081525061094e565b6040805180820190915260098152600080516020610c11833981519152602082015281565b610924828261092f565b15156103f357600080fd5b600160a060020a03166000908152602091909152604090205460ff1690565b6109b8826000836040518082805190602001908083835b602083106109845780518252601f199092019160209182019101610965565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610b31565b81600160a060020a03167fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489826040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a21578181015183820152602001610a09565b50505050905090810190601f168015610a4e5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b610ac9826000836040518082805190602001908083835b60208310610a955780518252601f199092019160209182019101610a76565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050610b56565b81600160a060020a03167fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a8260405180806020018281038252838181518152602001915080519060200190808383600083811015610a21578181015183820152602001610a09565b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b600160a060020a0316600090815260209190915260409020805460ff19169055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610bb957805160ff1916838001178555610be6565b82800160010185558215610be6579182015b82811115610be6578251825591602001919060010190610bcb565b50610bf2929150610bf6565b5090565b61080b91905b80821115610bf25760008155600101610bfc56007375706572757365720000000000000000000000000000000000000000000000a165627a7a72305820c68d29fc5679dd52d916509060ff2536a01618bfca604e755153ab8fdb47db770029`

// DeployUserable deploys a new Ethereum contract, binding an instance of Userable to it.
func DeployUserable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Userable, error) {
	parsed, err := abi.JSON(strings.NewReader(UserableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Userable{UserableCaller: UserableCaller{contract: contract}, UserableTransactor: UserableTransactor{contract: contract}, UserableFilterer: UserableFilterer{contract: contract}}, nil
}

// Userable is an auto generated Go binding around an Ethereum contract.
type Userable struct {
	UserableCaller     // Read-only binding to the contract
	UserableTransactor // Write-only binding to the contract
	UserableFilterer   // Log filterer for contract events
}

// UserableCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserableSession struct {
	Contract     *Userable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserableCallerSession struct {
	Contract *UserableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UserableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserableTransactorSession struct {
	Contract     *UserableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UserableRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserableRaw struct {
	Contract *Userable // Generic contract binding to access the raw methods on
}

// UserableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserableCallerRaw struct {
	Contract *UserableCaller // Generic read-only contract binding to access the raw methods on
}

// UserableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserableTransactorRaw struct {
	Contract *UserableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserable creates a new instance of Userable, bound to a specific deployed contract.
func NewUserable(address common.Address, backend bind.ContractBackend) (*Userable, error) {
	contract, err := bindUserable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Userable{UserableCaller: UserableCaller{contract: contract}, UserableTransactor: UserableTransactor{contract: contract}, UserableFilterer: UserableFilterer{contract: contract}}, nil
}

// NewUserableCaller creates a new read-only instance of Userable, bound to a specific deployed contract.
func NewUserableCaller(address common.Address, caller bind.ContractCaller) (*UserableCaller, error) {
	contract, err := bindUserable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserableCaller{contract: contract}, nil
}

// NewUserableTransactor creates a new write-only instance of Userable, bound to a specific deployed contract.
func NewUserableTransactor(address common.Address, transactor bind.ContractTransactor) (*UserableTransactor, error) {
	contract, err := bindUserable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserableTransactor{contract: contract}, nil
}

// NewUserableFilterer creates a new log filterer instance of Userable, bound to a specific deployed contract.
func NewUserableFilterer(address common.Address, filterer bind.ContractFilterer) (*UserableFilterer, error) {
	contract, err := bindUserable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserableFilterer{contract: contract}, nil
}

// bindUserable binds a generic wrapper to an already deployed contract.
func bindUserable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Userable *UserableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Userable.Contract.UserableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Userable *UserableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Userable.Contract.UserableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Userable *UserableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Userable.Contract.UserableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Userable *UserableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Userable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Userable *UserableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Userable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Userable *UserableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Userable.Contract.contract.Transact(opts, method, params...)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_Userable *UserableCaller) ROLEAUDITOR(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Userable.contract.Call(opts, out, "ROLE_AUDITOR")
	return *ret0, err
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_Userable *UserableSession) ROLEAUDITOR() (string, error) {
	return _Userable.Contract.ROLEAUDITOR(&_Userable.CallOpts)
}

// ROLEAUDITOR is a free data retrieval call binding the contract method 0x3f029aaf.
//
// Solidity: function ROLE_AUDITOR() constant returns(string)
func (_Userable *UserableCallerSession) ROLEAUDITOR() (string, error) {
	return _Userable.Contract.ROLEAUDITOR(&_Userable.CallOpts)
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_Userable *UserableCaller) ROLEDEVELOPER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Userable.contract.Call(opts, out, "ROLE_DEVELOPER")
	return *ret0, err
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_Userable *UserableSession) ROLEDEVELOPER() (string, error) {
	return _Userable.Contract.ROLEDEVELOPER(&_Userable.CallOpts)
}

// ROLEDEVELOPER is a free data retrieval call binding the contract method 0xe31cd22e.
//
// Solidity: function ROLE_DEVELOPER() constant returns(string)
func (_Userable *UserableCallerSession) ROLEDEVELOPER() (string, error) {
	return _Userable.Contract.ROLEDEVELOPER(&_Userable.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Userable *UserableCaller) ROLESUPERUSER(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Userable.contract.Call(opts, out, "ROLE_SUPERUSER")
	return *ret0, err
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Userable *UserableSession) ROLESUPERUSER() (string, error) {
	return _Userable.Contract.ROLESUPERUSER(&_Userable.CallOpts)
}

// ROLESUPERUSER is a free data retrieval call binding the contract method 0xebb4f484.
//
// Solidity: function ROLE_SUPERUSER() constant returns(string)
func (_Userable *UserableCallerSession) ROLESUPERUSER() (string, error) {
	return _Userable.Contract.ROLESUPERUSER(&_Userable.CallOpts)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Userable *UserableCaller) CheckRole(opts *bind.CallOpts, _operator common.Address, _role string) error {
	var ()
	out := &[]interface{}{}
	err := _Userable.contract.Call(opts, out, "checkRole", _operator, _role)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Userable *UserableSession) CheckRole(_operator common.Address, _role string) error {
	return _Userable.Contract.CheckRole(&_Userable.CallOpts, _operator, _role)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(_operator address, _role string) constant returns()
func (_Userable *UserableCallerSession) CheckRole(_operator common.Address, _role string) error {
	return _Userable.Contract.CheckRole(&_Userable.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Userable *UserableCaller) HasRole(opts *bind.CallOpts, _operator common.Address, _role string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Userable.contract.Call(opts, out, "hasRole", _operator, _role)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Userable *UserableSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Userable.Contract.HasRole(&_Userable.CallOpts, _operator, _role)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(_operator address, _role string) constant returns(bool)
func (_Userable *UserableCallerSession) HasRole(_operator common.Address, _role string) (bool, error) {
	return _Userable.Contract.HasRole(&_Userable.CallOpts, _operator, _role)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_Userable *UserableCaller) IsAuditor(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Userable.contract.Call(opts, out, "isAuditor", _addr)
	return *ret0, err
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_Userable *UserableSession) IsAuditor(_addr common.Address) (bool, error) {
	return _Userable.Contract.IsAuditor(&_Userable.CallOpts, _addr)
}

// IsAuditor is a free data retrieval call binding the contract method 0x49b90557.
//
// Solidity: function isAuditor(_addr address) constant returns(bool)
func (_Userable *UserableCallerSession) IsAuditor(_addr common.Address) (bool, error) {
	return _Userable.Contract.IsAuditor(&_Userable.CallOpts, _addr)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_Userable *UserableCaller) IsDeveloper(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Userable.contract.Call(opts, out, "isDeveloper", _addr)
	return *ret0, err
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_Userable *UserableSession) IsDeveloper(_addr common.Address) (bool, error) {
	return _Userable.Contract.IsDeveloper(&_Userable.CallOpts, _addr)
}

// IsDeveloper is a free data retrieval call binding the contract method 0x5eca4a70.
//
// Solidity: function isDeveloper(_addr address) constant returns(bool)
func (_Userable *UserableCallerSession) IsDeveloper(_addr common.Address) (bool, error) {
	return _Userable.Contract.IsDeveloper(&_Userable.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Userable *UserableCaller) IsSuperuser(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Userable.contract.Call(opts, out, "isSuperuser", _addr)
	return *ret0, err
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Userable *UserableSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _Userable.Contract.IsSuperuser(&_Userable.CallOpts, _addr)
}

// IsSuperuser is a free data retrieval call binding the contract method 0xbceee05e.
//
// Solidity: function isSuperuser(_addr address) constant returns(bool)
func (_Userable *UserableCallerSession) IsSuperuser(_addr common.Address) (bool, error) {
	return _Userable.Contract.IsSuperuser(&_Userable.CallOpts, _addr)
}

// Login is a free data retrieval call binding the contract method 0xb34e97e8.
//
// Solidity: function login() constant returns(string)
func (_Userable *UserableCaller) Login(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Userable.contract.Call(opts, out, "login")
	return *ret0, err
}

// Login is a free data retrieval call binding the contract method 0xb34e97e8.
//
// Solidity: function login() constant returns(string)
func (_Userable *UserableSession) Login() (string, error) {
	return _Userable.Contract.Login(&_Userable.CallOpts)
}

// Login is a free data retrieval call binding the contract method 0xb34e97e8.
//
// Solidity: function login() constant returns(string)
func (_Userable *UserableCallerSession) Login() (string, error) {
	return _Userable.Contract.Login(&_Userable.CallOpts)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_Userable *UserableTransactor) AddAuditor(opts *bind.TransactOpts, _newAuditor common.Address) (*types.Transaction, error) {
	return _Userable.contract.Transact(opts, "addAuditor", _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_Userable *UserableSession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _Userable.Contract.AddAuditor(&_Userable.TransactOpts, _newAuditor)
}

// AddAuditor is a paid mutator transaction binding the contract method 0xe429cef1.
//
// Solidity: function addAuditor(_newAuditor address) returns()
func (_Userable *UserableTransactorSession) AddAuditor(_newAuditor common.Address) (*types.Transaction, error) {
	return _Userable.Contract.AddAuditor(&_Userable.TransactOpts, _newAuditor)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_Userable *UserableTransactor) AddDeveloper(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Userable.contract.Transact(opts, "addDeveloper")
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_Userable *UserableSession) AddDeveloper() (*types.Transaction, error) {
	return _Userable.Contract.AddDeveloper(&_Userable.TransactOpts)
}

// AddDeveloper is a paid mutator transaction binding the contract method 0x2b7d7489.
//
// Solidity: function addDeveloper() returns()
func (_Userable *UserableTransactorSession) AddDeveloper() (*types.Transaction, error) {
	return _Userable.Contract.AddDeveloper(&_Userable.TransactOpts)
}

// SignUp is a paid mutator transaction binding the contract method 0xadcfb4fc.
//
// Solidity: function signUp(_name string) returns(string)
func (_Userable *UserableTransactor) SignUp(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Userable.contract.Transact(opts, "signUp", _name)
}

// SignUp is a paid mutator transaction binding the contract method 0xadcfb4fc.
//
// Solidity: function signUp(_name string) returns(string)
func (_Userable *UserableSession) SignUp(_name string) (*types.Transaction, error) {
	return _Userable.Contract.SignUp(&_Userable.TransactOpts, _name)
}

// SignUp is a paid mutator transaction binding the contract method 0xadcfb4fc.
//
// Solidity: function signUp(_name string) returns(string)
func (_Userable *UserableTransactorSession) SignUp(_name string) (*types.Transaction, error) {
	return _Userable.Contract.SignUp(&_Userable.TransactOpts, _name)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Userable *UserableTransactor) TransferSuperuser(opts *bind.TransactOpts, _newSuperuser common.Address) (*types.Transaction, error) {
	return _Userable.contract.Transact(opts, "transferSuperuser", _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Userable *UserableSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _Userable.Contract.TransferSuperuser(&_Userable.TransactOpts, _newSuperuser)
}

// TransferSuperuser is a paid mutator transaction binding the contract method 0x57c393fa.
//
// Solidity: function transferSuperuser(_newSuperuser address) returns()
func (_Userable *UserableTransactorSession) TransferSuperuser(_newSuperuser common.Address) (*types.Transaction, error) {
	return _Userable.Contract.TransferSuperuser(&_Userable.TransactOpts, _newSuperuser)
}

// UserableRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the Userable contract.
type UserableRoleAddedIterator struct {
	Event *UserableRoleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UserableRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserableRoleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UserableRoleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UserableRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserableRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserableRoleAdded represents a RoleAdded event raised by the Userable contract.
type UserableRoleAdded struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Userable *UserableFilterer) FilterRoleAdded(opts *bind.FilterOpts, operator []common.Address) (*UserableRoleAddedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Userable.contract.FilterLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return &UserableRoleAddedIterator{contract: _Userable.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(operator indexed address, role string)
func (_Userable *UserableFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *UserableRoleAdded, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Userable.contract.WatchLogs(opts, "RoleAdded", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserableRoleAdded)
				if err := _Userable.contract.UnpackLog(event, "RoleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UserableRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the Userable contract.
type UserableRoleRemovedIterator struct {
	Event *UserableRoleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UserableRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserableRoleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UserableRoleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UserableRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserableRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserableRoleRemoved represents a RoleRemoved event raised by the Userable contract.
type UserableRoleRemoved struct {
	Operator common.Address
	Role     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Userable *UserableFilterer) FilterRoleRemoved(opts *bind.FilterOpts, operator []common.Address) (*UserableRoleRemovedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Userable.contract.FilterLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return &UserableRoleRemovedIterator{contract: _Userable.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(operator indexed address, role string)
func (_Userable *UserableFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *UserableRoleRemoved, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Userable.contract.WatchLogs(opts, "RoleRemoved", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserableRoleRemoved)
				if err := _Userable.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
