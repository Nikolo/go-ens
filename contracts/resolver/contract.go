// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package resolver

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunction\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"OffchainLookup\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"StorageHandledByL2\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"}],\"internalType\":\"structIWriteDeferral.domainData\",\"name\":\"sender\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes4\",\"name\":\"functionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"internalType\":\"structIWriteDeferral.parameter[]\",\"name\":\"parameters\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIWriteDeferral.messageData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"StorageHandledByOffChainDatabase\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimeoutDurationTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimeoutDurationTooShort\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"}],\"name\":\"ABIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddrChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddress\",\"type\":\"bytes\"}],\"name\":\"AddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"ContenthashChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGatewayManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGatewayManager\",\"type\":\"address\"}],\"name\":\"GatewayManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"previousUrl\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"newUrl\",\"type\":\"string\"}],\"name\":\"GatewayUrlSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousContractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newContractAddress\",\"type\":\"address\"}],\"name\":\"L2HandlerContractAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"previousChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newChainId\",\"type\":\"uint256\"}],\"name\":\"L2HandlerDefaultChainIdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"previousUrl\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"newUrl\",\"type\":\"string\"}],\"name\":\"OffChainDatabaseHandlerURLChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousDuration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"OffChainDatabaseTimeoutDurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"PubkeyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addedSigner\",\"type\":\"address\"}],\"name\":\"SignerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousSignerManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSignerManager\",\"type\":\"address\"}],\"name\":\"SignerManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedSigner\",\"type\":\"address\"}],\"name\":\"SignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"indexedKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"TextChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentTypes\",\"type\":\"uint256\"}],\"name\":\"ABI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signersToAdd\",\"type\":\"address[]\"}],\"name\":\"addSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGatewayManager\",\"type\":\"address\"}],\"name\":\"changeGatewayManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSignerManager\",\"type\":\"address\"}],\"name\":\"changeSignerManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"contenthash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newSignerManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newGatewayManager\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"newGatewayUrl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newOffChainDatabaseUrl\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"newOffChainDatabaseTimeoutDuration\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"newSigners\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"makeSignatureHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offChainDatabaseUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"pubkey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"signersToRemove\",\"type\":\"address[]\"}],\"name\":\"removeSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"name\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"resolveWithProof\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"contentType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"setABI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"coinType\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"setAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"setContenthash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUrl\",\"type\":\"string\"}],\"name\":\"setGatewayUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"setOffChainDatabaseTimoutDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newUrl\",\"type\":\"string\"}],\"name\":\"setOffChainDatabaseUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"setPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"setText\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"text\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_Contract *ContractCaller) ABI(opts *bind.CallOpts, node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "ABI", node, contentTypes)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_Contract *ContractSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _Contract.Contract.ABI(&_Contract.CallOpts, node, contentTypes)
}

// ABI is a free data retrieval call binding the contract method 0x2203ab56.
//
// Solidity: function ABI(bytes32 node, uint256 contentTypes) view returns(uint256, bytes)
func (_Contract *ContractCallerSession) ABI(node [32]byte, contentTypes *big.Int) (*big.Int, []byte, error) {
	return _Contract.Contract.ABI(&_Contract.CallOpts, node, contentTypes)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Contract *ContractCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Contract *ContractSession) Addr(node [32]byte) (common.Address, error) {
	return _Contract.Contract.Addr(&_Contract.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Contract *ContractCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _Contract.Contract.Addr(&_Contract.CallOpts, node)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_Contract *ContractCaller) Addr0(opts *bind.CallOpts, node [32]byte, coinType *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "addr0", node, coinType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_Contract *ContractSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _Contract.Contract.Addr0(&_Contract.CallOpts, node, coinType)
}

// Addr0 is a free data retrieval call binding the contract method 0xf1cb7e06.
//
// Solidity: function addr(bytes32 node, uint256 coinType) view returns(bytes)
func (_Contract *ContractCallerSession) Addr0(node [32]byte, coinType *big.Int) ([]byte, error) {
	return _Contract.Contract.Addr0(&_Contract.CallOpts, node, coinType)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_Contract *ContractCaller) Contenthash(opts *bind.CallOpts, node [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "contenthash", node)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_Contract *ContractSession) Contenthash(node [32]byte) ([]byte, error) {
	return _Contract.Contract.Contenthash(&_Contract.CallOpts, node)
}

// Contenthash is a free data retrieval call binding the contract method 0xbc1c58d1.
//
// Solidity: function contenthash(bytes32 node) view returns(bytes)
func (_Contract *ContractCallerSession) Contenthash(node [32]byte) ([]byte, error) {
	return _Contract.Contract.Contenthash(&_Contract.CallOpts, node)
}

// GatewayManager is a free data retrieval call binding the contract method 0xd0414c9d.
//
// Solidity: function gatewayManager() view returns(address)
func (_Contract *ContractCaller) GatewayManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "gatewayManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayManager is a free data retrieval call binding the contract method 0xd0414c9d.
//
// Solidity: function gatewayManager() view returns(address)
func (_Contract *ContractSession) GatewayManager() (common.Address, error) {
	return _Contract.Contract.GatewayManager(&_Contract.CallOpts)
}

// GatewayManager is a free data retrieval call binding the contract method 0xd0414c9d.
//
// Solidity: function gatewayManager() view returns(address)
func (_Contract *ContractCallerSession) GatewayManager() (common.Address, error) {
	return _Contract.Contract.GatewayManager(&_Contract.CallOpts)
}

// GatewayUrl is a free data retrieval call binding the contract method 0x8bf165d9.
//
// Solidity: function gatewayUrl() view returns(string)
func (_Contract *ContractCaller) GatewayUrl(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "gatewayUrl")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GatewayUrl is a free data retrieval call binding the contract method 0x8bf165d9.
//
// Solidity: function gatewayUrl() view returns(string)
func (_Contract *ContractSession) GatewayUrl() (string, error) {
	return _Contract.Contract.GatewayUrl(&_Contract.CallOpts)
}

// GatewayUrl is a free data retrieval call binding the contract method 0x8bf165d9.
//
// Solidity: function gatewayUrl() view returns(string)
func (_Contract *ContractCallerSession) GatewayUrl() (string, error) {
	return _Contract.Contract.GatewayUrl(&_Contract.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address account) view returns(bool)
func (_Contract *ContractCaller) IsSigner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isSigner", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address account) view returns(bool)
func (_Contract *ContractSession) IsSigner(account common.Address) (bool, error) {
	return _Contract.Contract.IsSigner(&_Contract.CallOpts, account)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address account) view returns(bool)
func (_Contract *ContractCallerSession) IsSigner(account common.Address) (bool, error) {
	return _Contract.Contract.IsSigner(&_Contract.CallOpts, account)
}

// MakeSignatureHash is a free data retrieval call binding the contract method 0x0e829d0e.
//
// Solidity: function makeSignatureHash(uint64 expires, bytes request, bytes result) view returns(bytes32)
func (_Contract *ContractCaller) MakeSignatureHash(opts *bind.CallOpts, expires uint64, request []byte, result []byte) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "makeSignatureHash", expires, request, result)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeSignatureHash is a free data retrieval call binding the contract method 0x0e829d0e.
//
// Solidity: function makeSignatureHash(uint64 expires, bytes request, bytes result) view returns(bytes32)
func (_Contract *ContractSession) MakeSignatureHash(expires uint64, request []byte, result []byte) ([32]byte, error) {
	return _Contract.Contract.MakeSignatureHash(&_Contract.CallOpts, expires, request, result)
}

// MakeSignatureHash is a free data retrieval call binding the contract method 0x0e829d0e.
//
// Solidity: function makeSignatureHash(uint64 expires, bytes request, bytes result) view returns(bytes32)
func (_Contract *ContractCallerSession) MakeSignatureHash(expires uint64, request []byte, result []byte) ([32]byte, error) {
	return _Contract.Contract.MakeSignatureHash(&_Contract.CallOpts, expires, request, result)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Contract *ContractSession) Name(node [32]byte) (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Contract *ContractCallerSession) Name(node [32]byte) (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts, node)
}

// OffChainDatabaseUrl is a free data retrieval call binding the contract method 0xb7441ef9.
//
// Solidity: function offChainDatabaseUrl() view returns(string)
func (_Contract *ContractCaller) OffChainDatabaseUrl(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "offChainDatabaseUrl")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OffChainDatabaseUrl is a free data retrieval call binding the contract method 0xb7441ef9.
//
// Solidity: function offChainDatabaseUrl() view returns(string)
func (_Contract *ContractSession) OffChainDatabaseUrl() (string, error) {
	return _Contract.Contract.OffChainDatabaseUrl(&_Contract.CallOpts)
}

// OffChainDatabaseUrl is a free data retrieval call binding the contract method 0xb7441ef9.
//
// Solidity: function offChainDatabaseUrl() view returns(string)
func (_Contract *ContractCallerSession) OffChainDatabaseUrl() (string, error) {
	return _Contract.Contract.OffChainDatabaseUrl(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_Contract *ContractCaller) Pubkey(opts *bind.CallOpts, node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "pubkey", node)

	outstruct := new(struct {
		X [32]byte
		Y [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.X = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Y = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_Contract *ContractSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _Contract.Contract.Pubkey(&_Contract.CallOpts, node)
}

// Pubkey is a free data retrieval call binding the contract method 0xc8690233.
//
// Solidity: function pubkey(bytes32 node) view returns(bytes32 x, bytes32 y)
func (_Contract *ContractCallerSession) Pubkey(node [32]byte) (struct {
	X [32]byte
	Y [32]byte
}, error) {
	return _Contract.Contract.Pubkey(&_Contract.CallOpts, node)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (_Contract *ContractCaller) Resolve(opts *bind.CallOpts, name []byte, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolve", name, data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (_Contract *ContractSession) Resolve(name []byte, data []byte) ([]byte, error) {
	return _Contract.Contract.Resolve(&_Contract.CallOpts, name, data)
}

// Resolve is a free data retrieval call binding the contract method 0x9061b923.
//
// Solidity: function resolve(bytes name, bytes data) view returns(bytes)
func (_Contract *ContractCallerSession) Resolve(name []byte, data []byte) ([]byte, error) {
	return _Contract.Contract.Resolve(&_Contract.CallOpts, name, data)
}

// ResolveWithProof is a free data retrieval call binding the contract method 0xf4d4d2f8.
//
// Solidity: function resolveWithProof(bytes response, bytes extraData) view returns(bytes)
func (_Contract *ContractCaller) ResolveWithProof(opts *bind.CallOpts, response []byte, extraData []byte) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "resolveWithProof", response, extraData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ResolveWithProof is a free data retrieval call binding the contract method 0xf4d4d2f8.
//
// Solidity: function resolveWithProof(bytes response, bytes extraData) view returns(bytes)
func (_Contract *ContractSession) ResolveWithProof(response []byte, extraData []byte) ([]byte, error) {
	return _Contract.Contract.ResolveWithProof(&_Contract.CallOpts, response, extraData)
}

// ResolveWithProof is a free data retrieval call binding the contract method 0xf4d4d2f8.
//
// Solidity: function resolveWithProof(bytes response, bytes extraData) view returns(bytes)
func (_Contract *ContractCallerSession) ResolveWithProof(response []byte, extraData []byte) ([]byte, error) {
	return _Contract.Contract.ResolveWithProof(&_Contract.CallOpts, response, extraData)
}

// SignerManager is a free data retrieval call binding the contract method 0xcba89d58.
//
// Solidity: function signerManager() view returns(address)
func (_Contract *ContractCaller) SignerManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "signerManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerManager is a free data retrieval call binding the contract method 0xcba89d58.
//
// Solidity: function signerManager() view returns(address)
func (_Contract *ContractSession) SignerManager() (common.Address, error) {
	return _Contract.Contract.SignerManager(&_Contract.CallOpts)
}

// SignerManager is a free data retrieval call binding the contract method 0xcba89d58.
//
// Solidity: function signerManager() view returns(address)
func (_Contract *ContractCallerSession) SignerManager() (common.Address, error) {
	return _Contract.Contract.SignerManager(&_Contract.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x46f0975a.
//
// Solidity: function signers() view returns(address[])
func (_Contract *ContractCaller) Signers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "signers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x46f0975a.
//
// Solidity: function signers() view returns(address[])
func (_Contract *ContractSession) Signers() ([]common.Address, error) {
	return _Contract.Contract.Signers(&_Contract.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x46f0975a.
//
// Solidity: function signers() view returns(address[])
func (_Contract *ContractCallerSession) Signers() ([]common.Address, error) {
	return _Contract.Contract.Signers(&_Contract.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Contract *ContractSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, interfaceID)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Contract *ContractCaller) Text(opts *bind.CallOpts, node [32]byte, key string) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "text", node, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Contract *ContractSession) Text(node [32]byte, key string) (string, error) {
	return _Contract.Contract.Text(&_Contract.CallOpts, node, key)
}

// Text is a free data retrieval call binding the contract method 0x59d1d43c.
//
// Solidity: function text(bytes32 node, string key) view returns(string)
func (_Contract *ContractCallerSession) Text(node [32]byte, key string) (string, error) {
	return _Contract.Contract.Text(&_Contract.CallOpts, node, key)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signersToAdd) returns()
func (_Contract *ContractTransactor) AddSigners(opts *bind.TransactOpts, signersToAdd []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addSigners", signersToAdd)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signersToAdd) returns()
func (_Contract *ContractSession) AddSigners(signersToAdd []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddSigners(&_Contract.TransactOpts, signersToAdd)
}

// AddSigners is a paid mutator transaction binding the contract method 0xe8906a2d.
//
// Solidity: function addSigners(address[] signersToAdd) returns()
func (_Contract *ContractTransactorSession) AddSigners(signersToAdd []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddSigners(&_Contract.TransactOpts, signersToAdd)
}

// ChangeGatewayManager is a paid mutator transaction binding the contract method 0x1baacbf7.
//
// Solidity: function changeGatewayManager(address newGatewayManager) returns()
func (_Contract *ContractTransactor) ChangeGatewayManager(opts *bind.TransactOpts, newGatewayManager common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "changeGatewayManager", newGatewayManager)
}

// ChangeGatewayManager is a paid mutator transaction binding the contract method 0x1baacbf7.
//
// Solidity: function changeGatewayManager(address newGatewayManager) returns()
func (_Contract *ContractSession) ChangeGatewayManager(newGatewayManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeGatewayManager(&_Contract.TransactOpts, newGatewayManager)
}

// ChangeGatewayManager is a paid mutator transaction binding the contract method 0x1baacbf7.
//
// Solidity: function changeGatewayManager(address newGatewayManager) returns()
func (_Contract *ContractTransactorSession) ChangeGatewayManager(newGatewayManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeGatewayManager(&_Contract.TransactOpts, newGatewayManager)
}

// ChangeSignerManager is a paid mutator transaction binding the contract method 0xfb50b898.
//
// Solidity: function changeSignerManager(address newSignerManager) returns()
func (_Contract *ContractTransactor) ChangeSignerManager(opts *bind.TransactOpts, newSignerManager common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "changeSignerManager", newSignerManager)
}

// ChangeSignerManager is a paid mutator transaction binding the contract method 0xfb50b898.
//
// Solidity: function changeSignerManager(address newSignerManager) returns()
func (_Contract *ContractSession) ChangeSignerManager(newSignerManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeSignerManager(&_Contract.TransactOpts, newSignerManager)
}

// ChangeSignerManager is a paid mutator transaction binding the contract method 0xfb50b898.
//
// Solidity: function changeSignerManager(address newSignerManager) returns()
func (_Contract *ContractTransactorSession) ChangeSignerManager(newSignerManager common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ChangeSignerManager(&_Contract.TransactOpts, newSignerManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xa330a42b.
//
// Solidity: function initialize(address newOwner, address newSignerManager, address newGatewayManager, string newGatewayUrl, string newOffChainDatabaseUrl, uint256 newOffChainDatabaseTimeoutDuration, address[] newSigners) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, newOwner common.Address, newSignerManager common.Address, newGatewayManager common.Address, newGatewayUrl string, newOffChainDatabaseUrl string, newOffChainDatabaseTimeoutDuration *big.Int, newSigners []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", newOwner, newSignerManager, newGatewayManager, newGatewayUrl, newOffChainDatabaseUrl, newOffChainDatabaseTimeoutDuration, newSigners)
}

// Initialize is a paid mutator transaction binding the contract method 0xa330a42b.
//
// Solidity: function initialize(address newOwner, address newSignerManager, address newGatewayManager, string newGatewayUrl, string newOffChainDatabaseUrl, uint256 newOffChainDatabaseTimeoutDuration, address[] newSigners) returns()
func (_Contract *ContractSession) Initialize(newOwner common.Address, newSignerManager common.Address, newGatewayManager common.Address, newGatewayUrl string, newOffChainDatabaseUrl string, newOffChainDatabaseTimeoutDuration *big.Int, newSigners []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, newOwner, newSignerManager, newGatewayManager, newGatewayUrl, newOffChainDatabaseUrl, newOffChainDatabaseTimeoutDuration, newSigners)
}

// Initialize is a paid mutator transaction binding the contract method 0xa330a42b.
//
// Solidity: function initialize(address newOwner, address newSignerManager, address newGatewayManager, string newGatewayUrl, string newOffChainDatabaseUrl, uint256 newOffChainDatabaseTimeoutDuration, address[] newSigners) returns()
func (_Contract *ContractTransactorSession) Initialize(newOwner common.Address, newSignerManager common.Address, newGatewayManager common.Address, newGatewayUrl string, newOffChainDatabaseUrl string, newOffChainDatabaseTimeoutDuration *big.Int, newSigners []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, newOwner, newSignerManager, newGatewayManager, newGatewayUrl, newOffChainDatabaseUrl, newOffChainDatabaseTimeoutDuration, newSigners)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signersToRemove) returns()
func (_Contract *ContractTransactor) RemoveSigners(opts *bind.TransactOpts, signersToRemove []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeSigners", signersToRemove)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signersToRemove) returns()
func (_Contract *ContractSession) RemoveSigners(signersToRemove []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveSigners(&_Contract.TransactOpts, signersToRemove)
}

// RemoveSigners is a paid mutator transaction binding the contract method 0x8d361e43.
//
// Solidity: function removeSigners(address[] signersToRemove) returns()
func (_Contract *ContractTransactorSession) RemoveSigners(signersToRemove []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveSigners(&_Contract.TransactOpts, signersToRemove)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_Contract *ContractTransactor) SetABI(opts *bind.TransactOpts, node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setABI", node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_Contract *ContractSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetABI(&_Contract.TransactOpts, node, contentType, data)
}

// SetABI is a paid mutator transaction binding the contract method 0x623195b0.
//
// Solidity: function setABI(bytes32 node, uint256 contentType, bytes data) returns()
func (_Contract *ContractTransactorSession) SetABI(node [32]byte, contentType *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetABI(&_Contract.TransactOpts, node, contentType, data)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_Contract *ContractTransactor) SetAddr(opts *bind.TransactOpts, node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAddr", node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_Contract *ContractSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetAddr(&_Contract.TransactOpts, node, coinType, a)
}

// SetAddr is a paid mutator transaction binding the contract method 0x8b95dd71.
//
// Solidity: function setAddr(bytes32 node, uint256 coinType, bytes a) returns()
func (_Contract *ContractTransactorSession) SetAddr(node [32]byte, coinType *big.Int, a []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetAddr(&_Contract.TransactOpts, node, coinType, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_Contract *ContractTransactor) SetAddr0(opts *bind.TransactOpts, node [32]byte, a common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setAddr0", node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_Contract *ContractSession) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAddr0(&_Contract.TransactOpts, node, a)
}

// SetAddr0 is a paid mutator transaction binding the contract method 0xd5fa2b00.
//
// Solidity: function setAddr(bytes32 node, address a) returns()
func (_Contract *ContractTransactorSession) SetAddr0(node [32]byte, a common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetAddr0(&_Contract.TransactOpts, node, a)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_Contract *ContractTransactor) SetContenthash(opts *bind.TransactOpts, node [32]byte, hash []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setContenthash", node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_Contract *ContractSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetContenthash(&_Contract.TransactOpts, node, hash)
}

// SetContenthash is a paid mutator transaction binding the contract method 0x304e6ade.
//
// Solidity: function setContenthash(bytes32 node, bytes hash) returns()
func (_Contract *ContractTransactorSession) SetContenthash(node [32]byte, hash []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetContenthash(&_Contract.TransactOpts, node, hash)
}

// SetGatewayUrl is a paid mutator transaction binding the contract method 0x1787080e.
//
// Solidity: function setGatewayUrl(string newUrl) returns()
func (_Contract *ContractTransactor) SetGatewayUrl(opts *bind.TransactOpts, newUrl string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGatewayUrl", newUrl)
}

// SetGatewayUrl is a paid mutator transaction binding the contract method 0x1787080e.
//
// Solidity: function setGatewayUrl(string newUrl) returns()
func (_Contract *ContractSession) SetGatewayUrl(newUrl string) (*types.Transaction, error) {
	return _Contract.Contract.SetGatewayUrl(&_Contract.TransactOpts, newUrl)
}

// SetGatewayUrl is a paid mutator transaction binding the contract method 0x1787080e.
//
// Solidity: function setGatewayUrl(string newUrl) returns()
func (_Contract *ContractTransactorSession) SetGatewayUrl(newUrl string) (*types.Transaction, error) {
	return _Contract.Contract.SetGatewayUrl(&_Contract.TransactOpts, newUrl)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_Contract *ContractTransactor) SetName(opts *bind.TransactOpts, node [32]byte, name string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setName", node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_Contract *ContractSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _Contract.Contract.SetName(&_Contract.TransactOpts, node, name)
}

// SetName is a paid mutator transaction binding the contract method 0x77372213.
//
// Solidity: function setName(bytes32 node, string name) returns()
func (_Contract *ContractTransactorSession) SetName(node [32]byte, name string) (*types.Transaction, error) {
	return _Contract.Contract.SetName(&_Contract.TransactOpts, node, name)
}

// SetOffChainDatabaseTimoutDuration is a paid mutator transaction binding the contract method 0x27a5f3af.
//
// Solidity: function setOffChainDatabaseTimoutDuration(uint256 newDuration) returns()
func (_Contract *ContractTransactor) SetOffChainDatabaseTimoutDuration(opts *bind.TransactOpts, newDuration *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOffChainDatabaseTimoutDuration", newDuration)
}

// SetOffChainDatabaseTimoutDuration is a paid mutator transaction binding the contract method 0x27a5f3af.
//
// Solidity: function setOffChainDatabaseTimoutDuration(uint256 newDuration) returns()
func (_Contract *ContractSession) SetOffChainDatabaseTimoutDuration(newDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetOffChainDatabaseTimoutDuration(&_Contract.TransactOpts, newDuration)
}

// SetOffChainDatabaseTimoutDuration is a paid mutator transaction binding the contract method 0x27a5f3af.
//
// Solidity: function setOffChainDatabaseTimoutDuration(uint256 newDuration) returns()
func (_Contract *ContractTransactorSession) SetOffChainDatabaseTimoutDuration(newDuration *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetOffChainDatabaseTimoutDuration(&_Contract.TransactOpts, newDuration)
}

// SetOffChainDatabaseUrl is a paid mutator transaction binding the contract method 0x8e56d4a1.
//
// Solidity: function setOffChainDatabaseUrl(string newUrl) returns()
func (_Contract *ContractTransactor) SetOffChainDatabaseUrl(opts *bind.TransactOpts, newUrl string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOffChainDatabaseUrl", newUrl)
}

// SetOffChainDatabaseUrl is a paid mutator transaction binding the contract method 0x8e56d4a1.
//
// Solidity: function setOffChainDatabaseUrl(string newUrl) returns()
func (_Contract *ContractSession) SetOffChainDatabaseUrl(newUrl string) (*types.Transaction, error) {
	return _Contract.Contract.SetOffChainDatabaseUrl(&_Contract.TransactOpts, newUrl)
}

// SetOffChainDatabaseUrl is a paid mutator transaction binding the contract method 0x8e56d4a1.
//
// Solidity: function setOffChainDatabaseUrl(string newUrl) returns()
func (_Contract *ContractTransactorSession) SetOffChainDatabaseUrl(newUrl string) (*types.Transaction, error) {
	return _Contract.Contract.SetOffChainDatabaseUrl(&_Contract.TransactOpts, newUrl)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_Contract *ContractTransactor) SetPubkey(opts *bind.TransactOpts, node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPubkey", node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_Contract *ContractSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetPubkey(&_Contract.TransactOpts, node, x, y)
}

// SetPubkey is a paid mutator transaction binding the contract method 0x29cd62ea.
//
// Solidity: function setPubkey(bytes32 node, bytes32 x, bytes32 y) returns()
func (_Contract *ContractTransactorSession) SetPubkey(node [32]byte, x [32]byte, y [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SetPubkey(&_Contract.TransactOpts, node, x, y)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_Contract *ContractTransactor) SetText(opts *bind.TransactOpts, node [32]byte, key string, value string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setText", node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_Contract *ContractSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _Contract.Contract.SetText(&_Contract.TransactOpts, node, key, value)
}

// SetText is a paid mutator transaction binding the contract method 0x10f13a8c.
//
// Solidity: function setText(bytes32 node, string key, string value) returns()
func (_Contract *ContractTransactorSession) SetText(node [32]byte, key string, value string) (*types.Transaction, error) {
	return _Contract.Contract.SetText(&_Contract.TransactOpts, node, key, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// ContractABIChangedIterator is returned from FilterABIChanged and is used to iterate over the raw logs and unpacked data for ABIChanged events raised by the Contract contract.
type ContractABIChangedIterator struct {
	Event *ContractABIChanged // Event containing the contract specifics and raw log

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
func (it *ContractABIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractABIChanged)
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
		it.Event = new(ContractABIChanged)
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
func (it *ContractABIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractABIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractABIChanged represents a ABIChanged event raised by the Contract contract.
type ContractABIChanged struct {
	Node        [32]byte
	ContentType *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterABIChanged is a free log retrieval operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_Contract *ContractFilterer) FilterABIChanged(opts *bind.FilterOpts, node [][32]byte, contentType []*big.Int) (*ContractABIChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return &ContractABIChangedIterator{contract: _Contract.contract, event: "ABIChanged", logs: logs, sub: sub}, nil
}

// WatchABIChanged is a free log subscription operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_Contract *ContractFilterer) WatchABIChanged(opts *bind.WatchOpts, sink chan<- *ContractABIChanged, node [][32]byte, contentType []*big.Int) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var contentTypeRule []interface{}
	for _, contentTypeItem := range contentType {
		contentTypeRule = append(contentTypeRule, contentTypeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ABIChanged", nodeRule, contentTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractABIChanged)
				if err := _Contract.contract.UnpackLog(event, "ABIChanged", log); err != nil {
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

// ParseABIChanged is a log parse operation binding the contract event 0xaa121bbeef5f32f5961a2a28966e769023910fc9479059ee3495d4c1a696efe3.
//
// Solidity: event ABIChanged(bytes32 indexed node, uint256 indexed contentType)
func (_Contract *ContractFilterer) ParseABIChanged(log types.Log) (*ContractABIChanged, error) {
	event := new(ContractABIChanged)
	if err := _Contract.contract.UnpackLog(event, "ABIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAddrChangedIterator is returned from FilterAddrChanged and is used to iterate over the raw logs and unpacked data for AddrChanged events raised by the Contract contract.
type ContractAddrChangedIterator struct {
	Event *ContractAddrChanged // Event containing the contract specifics and raw log

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
func (it *ContractAddrChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddrChanged)
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
		it.Event = new(ContractAddrChanged)
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
func (it *ContractAddrChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddrChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddrChanged represents a AddrChanged event raised by the Contract contract.
type ContractAddrChanged struct {
	Node [32]byte
	A    common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddrChanged is a free log retrieval operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Contract *ContractFilterer) FilterAddrChanged(opts *bind.FilterOpts, node [][32]byte) (*ContractAddrChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractAddrChangedIterator{contract: _Contract.contract, event: "AddrChanged", logs: logs, sub: sub}, nil
}

// WatchAddrChanged is a free log subscription operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Contract *ContractFilterer) WatchAddrChanged(opts *bind.WatchOpts, sink chan<- *ContractAddrChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddrChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddrChanged)
				if err := _Contract.contract.UnpackLog(event, "AddrChanged", log); err != nil {
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

// ParseAddrChanged is a log parse operation binding the contract event 0x52d7d861f09ab3d26239d492e8968629f95e9e318cf0b73bfddc441522a15fd2.
//
// Solidity: event AddrChanged(bytes32 indexed node, address a)
func (_Contract *ContractFilterer) ParseAddrChanged(log types.Log) (*ContractAddrChanged, error) {
	event := new(ContractAddrChanged)
	if err := _Contract.contract.UnpackLog(event, "AddrChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAddressChangedIterator is returned from FilterAddressChanged and is used to iterate over the raw logs and unpacked data for AddressChanged events raised by the Contract contract.
type ContractAddressChangedIterator struct {
	Event *ContractAddressChanged // Event containing the contract specifics and raw log

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
func (it *ContractAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAddressChanged)
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
		it.Event = new(ContractAddressChanged)
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
func (it *ContractAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAddressChanged represents a AddressChanged event raised by the Contract contract.
type ContractAddressChanged struct {
	Node       [32]byte
	CoinType   *big.Int
	NewAddress []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressChanged is a free log retrieval operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_Contract *ContractFilterer) FilterAddressChanged(opts *bind.FilterOpts, node [][32]byte) (*ContractAddressChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractAddressChangedIterator{contract: _Contract.contract, event: "AddressChanged", logs: logs, sub: sub}, nil
}

// WatchAddressChanged is a free log subscription operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_Contract *ContractFilterer) WatchAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractAddressChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AddressChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "AddressChanged", log); err != nil {
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

// ParseAddressChanged is a log parse operation binding the contract event 0x65412581168e88a1e60c6459d7f44ae83ad0832e670826c05a4e2476b57af752.
//
// Solidity: event AddressChanged(bytes32 indexed node, uint256 coinType, bytes newAddress)
func (_Contract *ContractFilterer) ParseAddressChanged(log types.Log) (*ContractAddressChanged, error) {
	event := new(ContractAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "AddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractContenthashChangedIterator is returned from FilterContenthashChanged and is used to iterate over the raw logs and unpacked data for ContenthashChanged events raised by the Contract contract.
type ContractContenthashChangedIterator struct {
	Event *ContractContenthashChanged // Event containing the contract specifics and raw log

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
func (it *ContractContenthashChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractContenthashChanged)
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
		it.Event = new(ContractContenthashChanged)
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
func (it *ContractContenthashChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractContenthashChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractContenthashChanged represents a ContenthashChanged event raised by the Contract contract.
type ContractContenthashChanged struct {
	Node [32]byte
	Hash []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterContenthashChanged is a free log retrieval operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_Contract *ContractFilterer) FilterContenthashChanged(opts *bind.FilterOpts, node [][32]byte) (*ContractContenthashChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractContenthashChangedIterator{contract: _Contract.contract, event: "ContenthashChanged", logs: logs, sub: sub}, nil
}

// WatchContenthashChanged is a free log subscription operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_Contract *ContractFilterer) WatchContenthashChanged(opts *bind.WatchOpts, sink chan<- *ContractContenthashChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ContenthashChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractContenthashChanged)
				if err := _Contract.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
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

// ParseContenthashChanged is a log parse operation binding the contract event 0xe379c1624ed7e714cc0937528a32359d69d5281337765313dba4e081b72d7578.
//
// Solidity: event ContenthashChanged(bytes32 indexed node, bytes hash)
func (_Contract *ContractFilterer) ParseContenthashChanged(log types.Log) (*ContractContenthashChanged, error) {
	event := new(ContractContenthashChanged)
	if err := _Contract.contract.UnpackLog(event, "ContenthashChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractGatewayManagerChangedIterator is returned from FilterGatewayManagerChanged and is used to iterate over the raw logs and unpacked data for GatewayManagerChanged events raised by the Contract contract.
type ContractGatewayManagerChangedIterator struct {
	Event *ContractGatewayManagerChanged // Event containing the contract specifics and raw log

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
func (it *ContractGatewayManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGatewayManagerChanged)
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
		it.Event = new(ContractGatewayManagerChanged)
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
func (it *ContractGatewayManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractGatewayManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractGatewayManagerChanged represents a GatewayManagerChanged event raised by the Contract contract.
type ContractGatewayManagerChanged struct {
	PreviousGatewayManager common.Address
	NewGatewayManager      common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterGatewayManagerChanged is a free log retrieval operation binding the contract event 0x053cf517826bba468cbbc3577a5113ac1d1f3a2a9577e7a95f196c52dfb9e1c9.
//
// Solidity: event GatewayManagerChanged(address indexed previousGatewayManager, address indexed newGatewayManager)
func (_Contract *ContractFilterer) FilterGatewayManagerChanged(opts *bind.FilterOpts, previousGatewayManager []common.Address, newGatewayManager []common.Address) (*ContractGatewayManagerChangedIterator, error) {

	var previousGatewayManagerRule []interface{}
	for _, previousGatewayManagerItem := range previousGatewayManager {
		previousGatewayManagerRule = append(previousGatewayManagerRule, previousGatewayManagerItem)
	}
	var newGatewayManagerRule []interface{}
	for _, newGatewayManagerItem := range newGatewayManager {
		newGatewayManagerRule = append(newGatewayManagerRule, newGatewayManagerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "GatewayManagerChanged", previousGatewayManagerRule, newGatewayManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractGatewayManagerChangedIterator{contract: _Contract.contract, event: "GatewayManagerChanged", logs: logs, sub: sub}, nil
}

// WatchGatewayManagerChanged is a free log subscription operation binding the contract event 0x053cf517826bba468cbbc3577a5113ac1d1f3a2a9577e7a95f196c52dfb9e1c9.
//
// Solidity: event GatewayManagerChanged(address indexed previousGatewayManager, address indexed newGatewayManager)
func (_Contract *ContractFilterer) WatchGatewayManagerChanged(opts *bind.WatchOpts, sink chan<- *ContractGatewayManagerChanged, previousGatewayManager []common.Address, newGatewayManager []common.Address) (event.Subscription, error) {

	var previousGatewayManagerRule []interface{}
	for _, previousGatewayManagerItem := range previousGatewayManager {
		previousGatewayManagerRule = append(previousGatewayManagerRule, previousGatewayManagerItem)
	}
	var newGatewayManagerRule []interface{}
	for _, newGatewayManagerItem := range newGatewayManager {
		newGatewayManagerRule = append(newGatewayManagerRule, newGatewayManagerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "GatewayManagerChanged", previousGatewayManagerRule, newGatewayManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractGatewayManagerChanged)
				if err := _Contract.contract.UnpackLog(event, "GatewayManagerChanged", log); err != nil {
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

// ParseGatewayManagerChanged is a log parse operation binding the contract event 0x053cf517826bba468cbbc3577a5113ac1d1f3a2a9577e7a95f196c52dfb9e1c9.
//
// Solidity: event GatewayManagerChanged(address indexed previousGatewayManager, address indexed newGatewayManager)
func (_Contract *ContractFilterer) ParseGatewayManagerChanged(log types.Log) (*ContractGatewayManagerChanged, error) {
	event := new(ContractGatewayManagerChanged)
	if err := _Contract.contract.UnpackLog(event, "GatewayManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractGatewayUrlSetIterator is returned from FilterGatewayUrlSet and is used to iterate over the raw logs and unpacked data for GatewayUrlSet events raised by the Contract contract.
type ContractGatewayUrlSetIterator struct {
	Event *ContractGatewayUrlSet // Event containing the contract specifics and raw log

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
func (it *ContractGatewayUrlSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractGatewayUrlSet)
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
		it.Event = new(ContractGatewayUrlSet)
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
func (it *ContractGatewayUrlSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractGatewayUrlSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractGatewayUrlSet represents a GatewayUrlSet event raised by the Contract contract.
type ContractGatewayUrlSet struct {
	PreviousUrl common.Hash
	NewUrl      common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGatewayUrlSet is a free log retrieval operation binding the contract event 0x89e3f76277bf08e4371bcc677d04c7d3ab8fcef918f05d5e9668c532f3bd1207.
//
// Solidity: event GatewayUrlSet(string indexed previousUrl, string indexed newUrl)
func (_Contract *ContractFilterer) FilterGatewayUrlSet(opts *bind.FilterOpts, previousUrl []string, newUrl []string) (*ContractGatewayUrlSetIterator, error) {

	var previousUrlRule []interface{}
	for _, previousUrlItem := range previousUrl {
		previousUrlRule = append(previousUrlRule, previousUrlItem)
	}
	var newUrlRule []interface{}
	for _, newUrlItem := range newUrl {
		newUrlRule = append(newUrlRule, newUrlItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "GatewayUrlSet", previousUrlRule, newUrlRule)
	if err != nil {
		return nil, err
	}
	return &ContractGatewayUrlSetIterator{contract: _Contract.contract, event: "GatewayUrlSet", logs: logs, sub: sub}, nil
}

// WatchGatewayUrlSet is a free log subscription operation binding the contract event 0x89e3f76277bf08e4371bcc677d04c7d3ab8fcef918f05d5e9668c532f3bd1207.
//
// Solidity: event GatewayUrlSet(string indexed previousUrl, string indexed newUrl)
func (_Contract *ContractFilterer) WatchGatewayUrlSet(opts *bind.WatchOpts, sink chan<- *ContractGatewayUrlSet, previousUrl []string, newUrl []string) (event.Subscription, error) {

	var previousUrlRule []interface{}
	for _, previousUrlItem := range previousUrl {
		previousUrlRule = append(previousUrlRule, previousUrlItem)
	}
	var newUrlRule []interface{}
	for _, newUrlItem := range newUrl {
		newUrlRule = append(newUrlRule, newUrlItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "GatewayUrlSet", previousUrlRule, newUrlRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractGatewayUrlSet)
				if err := _Contract.contract.UnpackLog(event, "GatewayUrlSet", log); err != nil {
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

// ParseGatewayUrlSet is a log parse operation binding the contract event 0x89e3f76277bf08e4371bcc677d04c7d3ab8fcef918f05d5e9668c532f3bd1207.
//
// Solidity: event GatewayUrlSet(string indexed previousUrl, string indexed newUrl)
func (_Contract *ContractFilterer) ParseGatewayUrlSet(log types.Log) (*ContractGatewayUrlSet, error) {
	event := new(ContractGatewayUrlSet)
	if err := _Contract.contract.UnpackLog(event, "GatewayUrlSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

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
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
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
		it.Event = new(ContractInitialized)
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
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractL2HandlerContractAddressChangedIterator is returned from FilterL2HandlerContractAddressChanged and is used to iterate over the raw logs and unpacked data for L2HandlerContractAddressChanged events raised by the Contract contract.
type ContractL2HandlerContractAddressChangedIterator struct {
	Event *ContractL2HandlerContractAddressChanged // Event containing the contract specifics and raw log

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
func (it *ContractL2HandlerContractAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractL2HandlerContractAddressChanged)
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
		it.Event = new(ContractL2HandlerContractAddressChanged)
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
func (it *ContractL2HandlerContractAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractL2HandlerContractAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractL2HandlerContractAddressChanged represents a L2HandlerContractAddressChanged event raised by the Contract contract.
type ContractL2HandlerContractAddressChanged struct {
	ChainId                 *big.Int
	PreviousContractAddress common.Address
	NewContractAddress      common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterL2HandlerContractAddressChanged is a free log retrieval operation binding the contract event 0xb7a54307f58b341a0bb3d5dbbd490d3dc474b99c9335f80891b183929f02c4cd.
//
// Solidity: event L2HandlerContractAddressChanged(uint256 indexed chainId, address indexed previousContractAddress, address indexed newContractAddress)
func (_Contract *ContractFilterer) FilterL2HandlerContractAddressChanged(opts *bind.FilterOpts, chainId []*big.Int, previousContractAddress []common.Address, newContractAddress []common.Address) (*ContractL2HandlerContractAddressChangedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var previousContractAddressRule []interface{}
	for _, previousContractAddressItem := range previousContractAddress {
		previousContractAddressRule = append(previousContractAddressRule, previousContractAddressItem)
	}
	var newContractAddressRule []interface{}
	for _, newContractAddressItem := range newContractAddress {
		newContractAddressRule = append(newContractAddressRule, newContractAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "L2HandlerContractAddressChanged", chainIdRule, previousContractAddressRule, newContractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractL2HandlerContractAddressChangedIterator{contract: _Contract.contract, event: "L2HandlerContractAddressChanged", logs: logs, sub: sub}, nil
}

// WatchL2HandlerContractAddressChanged is a free log subscription operation binding the contract event 0xb7a54307f58b341a0bb3d5dbbd490d3dc474b99c9335f80891b183929f02c4cd.
//
// Solidity: event L2HandlerContractAddressChanged(uint256 indexed chainId, address indexed previousContractAddress, address indexed newContractAddress)
func (_Contract *ContractFilterer) WatchL2HandlerContractAddressChanged(opts *bind.WatchOpts, sink chan<- *ContractL2HandlerContractAddressChanged, chainId []*big.Int, previousContractAddress []common.Address, newContractAddress []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var previousContractAddressRule []interface{}
	for _, previousContractAddressItem := range previousContractAddress {
		previousContractAddressRule = append(previousContractAddressRule, previousContractAddressItem)
	}
	var newContractAddressRule []interface{}
	for _, newContractAddressItem := range newContractAddress {
		newContractAddressRule = append(newContractAddressRule, newContractAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "L2HandlerContractAddressChanged", chainIdRule, previousContractAddressRule, newContractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractL2HandlerContractAddressChanged)
				if err := _Contract.contract.UnpackLog(event, "L2HandlerContractAddressChanged", log); err != nil {
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

// ParseL2HandlerContractAddressChanged is a log parse operation binding the contract event 0xb7a54307f58b341a0bb3d5dbbd490d3dc474b99c9335f80891b183929f02c4cd.
//
// Solidity: event L2HandlerContractAddressChanged(uint256 indexed chainId, address indexed previousContractAddress, address indexed newContractAddress)
func (_Contract *ContractFilterer) ParseL2HandlerContractAddressChanged(log types.Log) (*ContractL2HandlerContractAddressChanged, error) {
	event := new(ContractL2HandlerContractAddressChanged)
	if err := _Contract.contract.UnpackLog(event, "L2HandlerContractAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractL2HandlerDefaultChainIdChangedIterator is returned from FilterL2HandlerDefaultChainIdChanged and is used to iterate over the raw logs and unpacked data for L2HandlerDefaultChainIdChanged events raised by the Contract contract.
type ContractL2HandlerDefaultChainIdChangedIterator struct {
	Event *ContractL2HandlerDefaultChainIdChanged // Event containing the contract specifics and raw log

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
func (it *ContractL2HandlerDefaultChainIdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractL2HandlerDefaultChainIdChanged)
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
		it.Event = new(ContractL2HandlerDefaultChainIdChanged)
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
func (it *ContractL2HandlerDefaultChainIdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractL2HandlerDefaultChainIdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractL2HandlerDefaultChainIdChanged represents a L2HandlerDefaultChainIdChanged event raised by the Contract contract.
type ContractL2HandlerDefaultChainIdChanged struct {
	PreviousChainId *big.Int
	NewChainId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterL2HandlerDefaultChainIdChanged is a free log retrieval operation binding the contract event 0x8f398717609692d5f4243541f192192f783256965570651dd3b78b7422cd0133.
//
// Solidity: event L2HandlerDefaultChainIdChanged(uint256 indexed previousChainId, uint256 indexed newChainId)
func (_Contract *ContractFilterer) FilterL2HandlerDefaultChainIdChanged(opts *bind.FilterOpts, previousChainId []*big.Int, newChainId []*big.Int) (*ContractL2HandlerDefaultChainIdChangedIterator, error) {

	var previousChainIdRule []interface{}
	for _, previousChainIdItem := range previousChainId {
		previousChainIdRule = append(previousChainIdRule, previousChainIdItem)
	}
	var newChainIdRule []interface{}
	for _, newChainIdItem := range newChainId {
		newChainIdRule = append(newChainIdRule, newChainIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "L2HandlerDefaultChainIdChanged", previousChainIdRule, newChainIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractL2HandlerDefaultChainIdChangedIterator{contract: _Contract.contract, event: "L2HandlerDefaultChainIdChanged", logs: logs, sub: sub}, nil
}

// WatchL2HandlerDefaultChainIdChanged is a free log subscription operation binding the contract event 0x8f398717609692d5f4243541f192192f783256965570651dd3b78b7422cd0133.
//
// Solidity: event L2HandlerDefaultChainIdChanged(uint256 indexed previousChainId, uint256 indexed newChainId)
func (_Contract *ContractFilterer) WatchL2HandlerDefaultChainIdChanged(opts *bind.WatchOpts, sink chan<- *ContractL2HandlerDefaultChainIdChanged, previousChainId []*big.Int, newChainId []*big.Int) (event.Subscription, error) {

	var previousChainIdRule []interface{}
	for _, previousChainIdItem := range previousChainId {
		previousChainIdRule = append(previousChainIdRule, previousChainIdItem)
	}
	var newChainIdRule []interface{}
	for _, newChainIdItem := range newChainId {
		newChainIdRule = append(newChainIdRule, newChainIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "L2HandlerDefaultChainIdChanged", previousChainIdRule, newChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractL2HandlerDefaultChainIdChanged)
				if err := _Contract.contract.UnpackLog(event, "L2HandlerDefaultChainIdChanged", log); err != nil {
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

// ParseL2HandlerDefaultChainIdChanged is a log parse operation binding the contract event 0x8f398717609692d5f4243541f192192f783256965570651dd3b78b7422cd0133.
//
// Solidity: event L2HandlerDefaultChainIdChanged(uint256 indexed previousChainId, uint256 indexed newChainId)
func (_Contract *ContractFilterer) ParseL2HandlerDefaultChainIdChanged(log types.Log) (*ContractL2HandlerDefaultChainIdChanged, error) {
	event := new(ContractL2HandlerDefaultChainIdChanged)
	if err := _Contract.contract.UnpackLog(event, "L2HandlerDefaultChainIdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the Contract contract.
type ContractNameChangedIterator struct {
	Event *ContractNameChanged // Event containing the contract specifics and raw log

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
func (it *ContractNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNameChanged)
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
		it.Event = new(ContractNameChanged)
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
func (it *ContractNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNameChanged represents a NameChanged event raised by the Contract contract.
type ContractNameChanged struct {
	Node [32]byte
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Contract *ContractFilterer) FilterNameChanged(opts *bind.FilterOpts, node [][32]byte) (*ContractNameChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractNameChangedIterator{contract: _Contract.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Contract *ContractFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *ContractNameChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NameChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNameChanged)
				if err := _Contract.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0xb7d29e911041e8d9b843369e890bcb72c9388692ba48b65ac54e7214c4c348f7.
//
// Solidity: event NameChanged(bytes32 indexed node, string name)
func (_Contract *ContractFilterer) ParseNameChanged(log types.Log) (*ContractNameChanged, error) {
	event := new(ContractNameChanged)
	if err := _Contract.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOffChainDatabaseHandlerURLChangedIterator is returned from FilterOffChainDatabaseHandlerURLChanged and is used to iterate over the raw logs and unpacked data for OffChainDatabaseHandlerURLChanged events raised by the Contract contract.
type ContractOffChainDatabaseHandlerURLChangedIterator struct {
	Event *ContractOffChainDatabaseHandlerURLChanged // Event containing the contract specifics and raw log

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
func (it *ContractOffChainDatabaseHandlerURLChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOffChainDatabaseHandlerURLChanged)
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
		it.Event = new(ContractOffChainDatabaseHandlerURLChanged)
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
func (it *ContractOffChainDatabaseHandlerURLChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOffChainDatabaseHandlerURLChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOffChainDatabaseHandlerURLChanged represents a OffChainDatabaseHandlerURLChanged event raised by the Contract contract.
type ContractOffChainDatabaseHandlerURLChanged struct {
	PreviousUrl common.Hash
	NewUrl      common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOffChainDatabaseHandlerURLChanged is a free log retrieval operation binding the contract event 0x80cd27d55de1340891dfb68b682bfcabcf01cd9b15fe94e75c5364db15848999.
//
// Solidity: event OffChainDatabaseHandlerURLChanged(string indexed previousUrl, string indexed newUrl)
func (_Contract *ContractFilterer) FilterOffChainDatabaseHandlerURLChanged(opts *bind.FilterOpts, previousUrl []string, newUrl []string) (*ContractOffChainDatabaseHandlerURLChangedIterator, error) {

	var previousUrlRule []interface{}
	for _, previousUrlItem := range previousUrl {
		previousUrlRule = append(previousUrlRule, previousUrlItem)
	}
	var newUrlRule []interface{}
	for _, newUrlItem := range newUrl {
		newUrlRule = append(newUrlRule, newUrlItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OffChainDatabaseHandlerURLChanged", previousUrlRule, newUrlRule)
	if err != nil {
		return nil, err
	}
	return &ContractOffChainDatabaseHandlerURLChangedIterator{contract: _Contract.contract, event: "OffChainDatabaseHandlerURLChanged", logs: logs, sub: sub}, nil
}

// WatchOffChainDatabaseHandlerURLChanged is a free log subscription operation binding the contract event 0x80cd27d55de1340891dfb68b682bfcabcf01cd9b15fe94e75c5364db15848999.
//
// Solidity: event OffChainDatabaseHandlerURLChanged(string indexed previousUrl, string indexed newUrl)
func (_Contract *ContractFilterer) WatchOffChainDatabaseHandlerURLChanged(opts *bind.WatchOpts, sink chan<- *ContractOffChainDatabaseHandlerURLChanged, previousUrl []string, newUrl []string) (event.Subscription, error) {

	var previousUrlRule []interface{}
	for _, previousUrlItem := range previousUrl {
		previousUrlRule = append(previousUrlRule, previousUrlItem)
	}
	var newUrlRule []interface{}
	for _, newUrlItem := range newUrl {
		newUrlRule = append(newUrlRule, newUrlItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OffChainDatabaseHandlerURLChanged", previousUrlRule, newUrlRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOffChainDatabaseHandlerURLChanged)
				if err := _Contract.contract.UnpackLog(event, "OffChainDatabaseHandlerURLChanged", log); err != nil {
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

// ParseOffChainDatabaseHandlerURLChanged is a log parse operation binding the contract event 0x80cd27d55de1340891dfb68b682bfcabcf01cd9b15fe94e75c5364db15848999.
//
// Solidity: event OffChainDatabaseHandlerURLChanged(string indexed previousUrl, string indexed newUrl)
func (_Contract *ContractFilterer) ParseOffChainDatabaseHandlerURLChanged(log types.Log) (*ContractOffChainDatabaseHandlerURLChanged, error) {
	event := new(ContractOffChainDatabaseHandlerURLChanged)
	if err := _Contract.contract.UnpackLog(event, "OffChainDatabaseHandlerURLChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOffChainDatabaseTimeoutDurationSetIterator is returned from FilterOffChainDatabaseTimeoutDurationSet and is used to iterate over the raw logs and unpacked data for OffChainDatabaseTimeoutDurationSet events raised by the Contract contract.
type ContractOffChainDatabaseTimeoutDurationSetIterator struct {
	Event *ContractOffChainDatabaseTimeoutDurationSet // Event containing the contract specifics and raw log

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
func (it *ContractOffChainDatabaseTimeoutDurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOffChainDatabaseTimeoutDurationSet)
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
		it.Event = new(ContractOffChainDatabaseTimeoutDurationSet)
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
func (it *ContractOffChainDatabaseTimeoutDurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOffChainDatabaseTimeoutDurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOffChainDatabaseTimeoutDurationSet represents a OffChainDatabaseTimeoutDurationSet event raised by the Contract contract.
type ContractOffChainDatabaseTimeoutDurationSet struct {
	PreviousDuration *big.Int
	NewDuration      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOffChainDatabaseTimeoutDurationSet is a free log retrieval operation binding the contract event 0x96c6caf9f42ebbbb9ba734bc86cf4885bd740fa956cdc9e642cafdb29d1b0e9b.
//
// Solidity: event OffChainDatabaseTimeoutDurationSet(uint256 previousDuration, uint256 newDuration)
func (_Contract *ContractFilterer) FilterOffChainDatabaseTimeoutDurationSet(opts *bind.FilterOpts) (*ContractOffChainDatabaseTimeoutDurationSetIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OffChainDatabaseTimeoutDurationSet")
	if err != nil {
		return nil, err
	}
	return &ContractOffChainDatabaseTimeoutDurationSetIterator{contract: _Contract.contract, event: "OffChainDatabaseTimeoutDurationSet", logs: logs, sub: sub}, nil
}

// WatchOffChainDatabaseTimeoutDurationSet is a free log subscription operation binding the contract event 0x96c6caf9f42ebbbb9ba734bc86cf4885bd740fa956cdc9e642cafdb29d1b0e9b.
//
// Solidity: event OffChainDatabaseTimeoutDurationSet(uint256 previousDuration, uint256 newDuration)
func (_Contract *ContractFilterer) WatchOffChainDatabaseTimeoutDurationSet(opts *bind.WatchOpts, sink chan<- *ContractOffChainDatabaseTimeoutDurationSet) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OffChainDatabaseTimeoutDurationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOffChainDatabaseTimeoutDurationSet)
				if err := _Contract.contract.UnpackLog(event, "OffChainDatabaseTimeoutDurationSet", log); err != nil {
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

// ParseOffChainDatabaseTimeoutDurationSet is a log parse operation binding the contract event 0x96c6caf9f42ebbbb9ba734bc86cf4885bd740fa956cdc9e642cafdb29d1b0e9b.
//
// Solidity: event OffChainDatabaseTimeoutDurationSet(uint256 previousDuration, uint256 newDuration)
func (_Contract *ContractFilterer) ParseOffChainDatabaseTimeoutDurationSet(log types.Log) (*ContractOffChainDatabaseTimeoutDurationSet, error) {
	event := new(ContractOffChainDatabaseTimeoutDurationSet)
	if err := _Contract.contract.UnpackLog(event, "OffChainDatabaseTimeoutDurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractPubkeyChangedIterator is returned from FilterPubkeyChanged and is used to iterate over the raw logs and unpacked data for PubkeyChanged events raised by the Contract contract.
type ContractPubkeyChangedIterator struct {
	Event *ContractPubkeyChanged // Event containing the contract specifics and raw log

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
func (it *ContractPubkeyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPubkeyChanged)
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
		it.Event = new(ContractPubkeyChanged)
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
func (it *ContractPubkeyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPubkeyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPubkeyChanged represents a PubkeyChanged event raised by the Contract contract.
type ContractPubkeyChanged struct {
	Node [32]byte
	X    [32]byte
	Y    [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPubkeyChanged is a free log retrieval operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_Contract *ContractFilterer) FilterPubkeyChanged(opts *bind.FilterOpts, node [][32]byte) (*ContractPubkeyChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return &ContractPubkeyChangedIterator{contract: _Contract.contract, event: "PubkeyChanged", logs: logs, sub: sub}, nil
}

// WatchPubkeyChanged is a free log subscription operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_Contract *ContractFilterer) WatchPubkeyChanged(opts *bind.WatchOpts, sink chan<- *ContractPubkeyChanged, node [][32]byte) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PubkeyChanged", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPubkeyChanged)
				if err := _Contract.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
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

// ParsePubkeyChanged is a log parse operation binding the contract event 0x1d6f5e03d3f63eb58751986629a5439baee5079ff04f345becb66e23eb154e46.
//
// Solidity: event PubkeyChanged(bytes32 indexed node, bytes32 x, bytes32 y)
func (_Contract *ContractFilterer) ParsePubkeyChanged(log types.Log) (*ContractPubkeyChanged, error) {
	event := new(ContractPubkeyChanged)
	if err := _Contract.contract.UnpackLog(event, "PubkeyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSignerAddedIterator is returned from FilterSignerAdded and is used to iterate over the raw logs and unpacked data for SignerAdded events raised by the Contract contract.
type ContractSignerAddedIterator struct {
	Event *ContractSignerAdded // Event containing the contract specifics and raw log

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
func (it *ContractSignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSignerAdded)
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
		it.Event = new(ContractSignerAdded)
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
func (it *ContractSignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSignerAdded represents a SignerAdded event raised by the Contract contract.
type ContractSignerAdded struct {
	AddedSigner common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSignerAdded is a free log retrieval operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed addedSigner)
func (_Contract *ContractFilterer) FilterSignerAdded(opts *bind.FilterOpts, addedSigner []common.Address) (*ContractSignerAddedIterator, error) {

	var addedSignerRule []interface{}
	for _, addedSignerItem := range addedSigner {
		addedSignerRule = append(addedSignerRule, addedSignerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SignerAdded", addedSignerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSignerAddedIterator{contract: _Contract.contract, event: "SignerAdded", logs: logs, sub: sub}, nil
}

// WatchSignerAdded is a free log subscription operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed addedSigner)
func (_Contract *ContractFilterer) WatchSignerAdded(opts *bind.WatchOpts, sink chan<- *ContractSignerAdded, addedSigner []common.Address) (event.Subscription, error) {

	var addedSignerRule []interface{}
	for _, addedSignerItem := range addedSigner {
		addedSignerRule = append(addedSignerRule, addedSignerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SignerAdded", addedSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSignerAdded)
				if err := _Contract.contract.UnpackLog(event, "SignerAdded", log); err != nil {
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

// ParseSignerAdded is a log parse operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed addedSigner)
func (_Contract *ContractFilterer) ParseSignerAdded(log types.Log) (*ContractSignerAdded, error) {
	event := new(ContractSignerAdded)
	if err := _Contract.contract.UnpackLog(event, "SignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSignerManagerChangedIterator is returned from FilterSignerManagerChanged and is used to iterate over the raw logs and unpacked data for SignerManagerChanged events raised by the Contract contract.
type ContractSignerManagerChangedIterator struct {
	Event *ContractSignerManagerChanged // Event containing the contract specifics and raw log

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
func (it *ContractSignerManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSignerManagerChanged)
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
		it.Event = new(ContractSignerManagerChanged)
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
func (it *ContractSignerManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSignerManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSignerManagerChanged represents a SignerManagerChanged event raised by the Contract contract.
type ContractSignerManagerChanged struct {
	PreviousSignerManager common.Address
	NewSignerManager      common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSignerManagerChanged is a free log retrieval operation binding the contract event 0xec6fde59c41700c846fcc431ef18452f74214c199e384fcdb37ebb16baf77a50.
//
// Solidity: event SignerManagerChanged(address indexed previousSignerManager, address indexed newSignerManager)
func (_Contract *ContractFilterer) FilterSignerManagerChanged(opts *bind.FilterOpts, previousSignerManager []common.Address, newSignerManager []common.Address) (*ContractSignerManagerChangedIterator, error) {

	var previousSignerManagerRule []interface{}
	for _, previousSignerManagerItem := range previousSignerManager {
		previousSignerManagerRule = append(previousSignerManagerRule, previousSignerManagerItem)
	}
	var newSignerManagerRule []interface{}
	for _, newSignerManagerItem := range newSignerManager {
		newSignerManagerRule = append(newSignerManagerRule, newSignerManagerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SignerManagerChanged", previousSignerManagerRule, newSignerManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSignerManagerChangedIterator{contract: _Contract.contract, event: "SignerManagerChanged", logs: logs, sub: sub}, nil
}

// WatchSignerManagerChanged is a free log subscription operation binding the contract event 0xec6fde59c41700c846fcc431ef18452f74214c199e384fcdb37ebb16baf77a50.
//
// Solidity: event SignerManagerChanged(address indexed previousSignerManager, address indexed newSignerManager)
func (_Contract *ContractFilterer) WatchSignerManagerChanged(opts *bind.WatchOpts, sink chan<- *ContractSignerManagerChanged, previousSignerManager []common.Address, newSignerManager []common.Address) (event.Subscription, error) {

	var previousSignerManagerRule []interface{}
	for _, previousSignerManagerItem := range previousSignerManager {
		previousSignerManagerRule = append(previousSignerManagerRule, previousSignerManagerItem)
	}
	var newSignerManagerRule []interface{}
	for _, newSignerManagerItem := range newSignerManager {
		newSignerManagerRule = append(newSignerManagerRule, newSignerManagerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SignerManagerChanged", previousSignerManagerRule, newSignerManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSignerManagerChanged)
				if err := _Contract.contract.UnpackLog(event, "SignerManagerChanged", log); err != nil {
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

// ParseSignerManagerChanged is a log parse operation binding the contract event 0xec6fde59c41700c846fcc431ef18452f74214c199e384fcdb37ebb16baf77a50.
//
// Solidity: event SignerManagerChanged(address indexed previousSignerManager, address indexed newSignerManager)
func (_Contract *ContractFilterer) ParseSignerManagerChanged(log types.Log) (*ContractSignerManagerChanged, error) {
	event := new(ContractSignerManagerChanged)
	if err := _Contract.contract.UnpackLog(event, "SignerManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSignerRemovedIterator is returned from FilterSignerRemoved and is used to iterate over the raw logs and unpacked data for SignerRemoved events raised by the Contract contract.
type ContractSignerRemovedIterator struct {
	Event *ContractSignerRemoved // Event containing the contract specifics and raw log

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
func (it *ContractSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSignerRemoved)
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
		it.Event = new(ContractSignerRemoved)
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
func (it *ContractSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSignerRemoved represents a SignerRemoved event raised by the Contract contract.
type ContractSignerRemoved struct {
	RemovedSigner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignerRemoved is a free log retrieval operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed removedSigner)
func (_Contract *ContractFilterer) FilterSignerRemoved(opts *bind.FilterOpts, removedSigner []common.Address) (*ContractSignerRemovedIterator, error) {

	var removedSignerRule []interface{}
	for _, removedSignerItem := range removedSigner {
		removedSignerRule = append(removedSignerRule, removedSignerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SignerRemoved", removedSignerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSignerRemovedIterator{contract: _Contract.contract, event: "SignerRemoved", logs: logs, sub: sub}, nil
}

// WatchSignerRemoved is a free log subscription operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed removedSigner)
func (_Contract *ContractFilterer) WatchSignerRemoved(opts *bind.WatchOpts, sink chan<- *ContractSignerRemoved, removedSigner []common.Address) (event.Subscription, error) {

	var removedSignerRule []interface{}
	for _, removedSignerItem := range removedSigner {
		removedSignerRule = append(removedSignerRule, removedSignerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SignerRemoved", removedSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSignerRemoved)
				if err := _Contract.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
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

// ParseSignerRemoved is a log parse operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed removedSigner)
func (_Contract *ContractFilterer) ParseSignerRemoved(log types.Log) (*ContractSignerRemoved, error) {
	event := new(ContractSignerRemoved)
	if err := _Contract.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTextChangedIterator is returned from FilterTextChanged and is used to iterate over the raw logs and unpacked data for TextChanged events raised by the Contract contract.
type ContractTextChangedIterator struct {
	Event *ContractTextChanged // Event containing the contract specifics and raw log

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
func (it *ContractTextChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTextChanged)
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
		it.Event = new(ContractTextChanged)
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
func (it *ContractTextChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTextChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTextChanged represents a TextChanged event raised by the Contract contract.
type ContractTextChanged struct {
	Node       [32]byte
	IndexedKey common.Hash
	Key        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTextChanged is a free log retrieval operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Contract *ContractFilterer) FilterTextChanged(opts *bind.FilterOpts, node [][32]byte, indexedKey []string) (*ContractTextChangedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return &ContractTextChangedIterator{contract: _Contract.contract, event: "TextChanged", logs: logs, sub: sub}, nil
}

// WatchTextChanged is a free log subscription operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Contract *ContractFilterer) WatchTextChanged(opts *bind.WatchOpts, sink chan<- *ContractTextChanged, node [][32]byte, indexedKey []string) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var indexedKeyRule []interface{}
	for _, indexedKeyItem := range indexedKey {
		indexedKeyRule = append(indexedKeyRule, indexedKeyItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TextChanged", nodeRule, indexedKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTextChanged)
				if err := _Contract.contract.UnpackLog(event, "TextChanged", log); err != nil {
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

// ParseTextChanged is a log parse operation binding the contract event 0xd8c9334b1a9c2f9da342a0a2b32629c1a229b6445dad78947f674b44444a7550.
//
// Solidity: event TextChanged(bytes32 indexed node, string indexed indexedKey, string key)
func (_Contract *ContractFilterer) ParseTextChanged(log types.Log) (*ContractTextChanged, error) {
	event := new(ContractTextChanged)
	if err := _Contract.contract.UnpackLog(event, "TextChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
