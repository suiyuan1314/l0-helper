// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package endpointv2view

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

// LzReceiveParam is an auto generated low-level Go binding around an user-defined struct.
type LzReceiveParam struct {
	Origin    Origin
	Receiver  common.Address
	Guid      [32]byte
	Message   []byte
	ExtraData []byte
	Gas       *big.Int
	Value     *big.Int
}

// NativeDropParam is an auto generated low-level Go binding around an user-defined struct.
type NativeDropParam struct {
	Receiver common.Address
	Amount   *big.Int
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// EndpointV2ViewMetaData contains all meta data concerning the EndpointV2View contract.
var EndpointV2ViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"LzExecutor_Executed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzExecutor_ReceiveLibViewNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LzExecutor_Verifying\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer_NativeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Transfer_ToAddressIsZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"NativeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiveLib\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiveLibView\",\"type\":\"address\"}],\"name\":\"ReceiveLibViewSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EMPTY_PAYLOAD_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NIL_PAYLOAD_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiveLib\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLzReceiveParam\",\"name\":\"_lzReceiveParam\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structNativeDropParam[]\",\"name\":\"_nativeDropParams\",\"type\":\"tuple[]\"}],\"name\":\"commitAndExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpointV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"executable\",\"outputs\":[{\"internalType\":\"enumExecutionState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"initializable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiveUln302\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiveUln302View\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localEid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiveLib\",\"type\":\"address\"}],\"name\":\"receiveLibToView\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiveLibView\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveUln302\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiveLib\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiveLibView\",\"type\":\"address\"}],\"name\":\"setReceiveLibView\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiveLib\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"verifiable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EndpointV2ViewABI is the input ABI used to generate the binding from.
// Deprecated: Use EndpointV2ViewMetaData.ABI instead.
var EndpointV2ViewABI = EndpointV2ViewMetaData.ABI

// EndpointV2View is an auto generated Go binding around an Ethereum contract.
type EndpointV2View struct {
	EndpointV2ViewCaller     // Read-only binding to the contract
	EndpointV2ViewTransactor // Write-only binding to the contract
	EndpointV2ViewFilterer   // Log filterer for contract events
}

// EndpointV2ViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type EndpointV2ViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointV2ViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EndpointV2ViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointV2ViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EndpointV2ViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EndpointV2ViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EndpointV2ViewSession struct {
	Contract     *EndpointV2View   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EndpointV2ViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EndpointV2ViewCallerSession struct {
	Contract *EndpointV2ViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EndpointV2ViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EndpointV2ViewTransactorSession struct {
	Contract     *EndpointV2ViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EndpointV2ViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type EndpointV2ViewRaw struct {
	Contract *EndpointV2View // Generic contract binding to access the raw methods on
}

// EndpointV2ViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EndpointV2ViewCallerRaw struct {
	Contract *EndpointV2ViewCaller // Generic read-only contract binding to access the raw methods on
}

// EndpointV2ViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EndpointV2ViewTransactorRaw struct {
	Contract *EndpointV2ViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEndpointV2View creates a new instance of EndpointV2View, bound to a specific deployed contract.
func NewEndpointV2View(address common.Address, backend bind.ContractBackend) (*EndpointV2View, error) {
	contract, err := bindEndpointV2View(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EndpointV2View{EndpointV2ViewCaller: EndpointV2ViewCaller{contract: contract}, EndpointV2ViewTransactor: EndpointV2ViewTransactor{contract: contract}, EndpointV2ViewFilterer: EndpointV2ViewFilterer{contract: contract}}, nil
}

// NewEndpointV2ViewCaller creates a new read-only instance of EndpointV2View, bound to a specific deployed contract.
func NewEndpointV2ViewCaller(address common.Address, caller bind.ContractCaller) (*EndpointV2ViewCaller, error) {
	contract, err := bindEndpointV2View(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EndpointV2ViewCaller{contract: contract}, nil
}

// NewEndpointV2ViewTransactor creates a new write-only instance of EndpointV2View, bound to a specific deployed contract.
func NewEndpointV2ViewTransactor(address common.Address, transactor bind.ContractTransactor) (*EndpointV2ViewTransactor, error) {
	contract, err := bindEndpointV2View(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EndpointV2ViewTransactor{contract: contract}, nil
}

// NewEndpointV2ViewFilterer creates a new log filterer instance of EndpointV2View, bound to a specific deployed contract.
func NewEndpointV2ViewFilterer(address common.Address, filterer bind.ContractFilterer) (*EndpointV2ViewFilterer, error) {
	contract, err := bindEndpointV2View(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EndpointV2ViewFilterer{contract: contract}, nil
}

// bindEndpointV2View binds a generic wrapper to an already deployed contract.
func bindEndpointV2View(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EndpointV2ViewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EndpointV2View *EndpointV2ViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EndpointV2View.Contract.EndpointV2ViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EndpointV2View *EndpointV2ViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EndpointV2View.Contract.EndpointV2ViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EndpointV2View *EndpointV2ViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EndpointV2View.Contract.EndpointV2ViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EndpointV2View *EndpointV2ViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EndpointV2View.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EndpointV2View *EndpointV2ViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EndpointV2View.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EndpointV2View *EndpointV2ViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EndpointV2View.Contract.contract.Transact(opts, method, params...)
}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2View *EndpointV2ViewCaller) EMPTYPAYLOADHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "EMPTY_PAYLOAD_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2View *EndpointV2ViewSession) EMPTYPAYLOADHASH() ([32]byte, error) {
	return _EndpointV2View.Contract.EMPTYPAYLOADHASH(&_EndpointV2View.CallOpts)
}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2View *EndpointV2ViewCallerSession) EMPTYPAYLOADHASH() ([32]byte, error) {
	return _EndpointV2View.Contract.EMPTYPAYLOADHASH(&_EndpointV2View.CallOpts)
}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2View *EndpointV2ViewCaller) NILPAYLOADHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "NIL_PAYLOAD_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2View *EndpointV2ViewSession) NILPAYLOADHASH() ([32]byte, error) {
	return _EndpointV2View.Contract.NILPAYLOADHASH(&_EndpointV2View.CallOpts)
}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_EndpointV2View *EndpointV2ViewCallerSession) NILPAYLOADHASH() ([32]byte, error) {
	return _EndpointV2View.Contract.NILPAYLOADHASH(&_EndpointV2View.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_EndpointV2View *EndpointV2ViewCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_EndpointV2View *EndpointV2ViewSession) Endpoint() (common.Address, error) {
	return _EndpointV2View.Contract.Endpoint(&_EndpointV2View.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_EndpointV2View *EndpointV2ViewCallerSession) Endpoint() (common.Address, error) {
	return _EndpointV2View.Contract.Endpoint(&_EndpointV2View.CallOpts)
}

// Executable is a free data retrieval call binding the contract method 0x4b4b2efb.
//
// Solidity: function executable((uint32,bytes32,uint64) _origin, address _receiver) view returns(uint8)
func (_EndpointV2View *EndpointV2ViewCaller) Executable(opts *bind.CallOpts, _origin Origin, _receiver common.Address) (uint8, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "executable", _origin, _receiver)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Executable is a free data retrieval call binding the contract method 0x4b4b2efb.
//
// Solidity: function executable((uint32,bytes32,uint64) _origin, address _receiver) view returns(uint8)
func (_EndpointV2View *EndpointV2ViewSession) Executable(_origin Origin, _receiver common.Address) (uint8, error) {
	return _EndpointV2View.Contract.Executable(&_EndpointV2View.CallOpts, _origin, _receiver)
}

// Executable is a free data retrieval call binding the contract method 0x4b4b2efb.
//
// Solidity: function executable((uint32,bytes32,uint64) _origin, address _receiver) view returns(uint8)
func (_EndpointV2View *EndpointV2ViewCallerSession) Executable(_origin Origin, _receiver common.Address) (uint8, error) {
	return _EndpointV2View.Contract.Executable(&_EndpointV2View.CallOpts, _origin, _receiver)
}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2View *EndpointV2ViewCaller) Initializable(opts *bind.CallOpts, _origin Origin, _receiver common.Address) (bool, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "initializable", _origin, _receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2View *EndpointV2ViewSession) Initializable(_origin Origin, _receiver common.Address) (bool, error) {
	return _EndpointV2View.Contract.Initializable(&_EndpointV2View.CallOpts, _origin, _receiver)
}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_EndpointV2View *EndpointV2ViewCallerSession) Initializable(_origin Origin, _receiver common.Address) (bool, error) {
	return _EndpointV2View.Contract.Initializable(&_EndpointV2View.CallOpts, _origin, _receiver)
}

// LocalEid is a free data retrieval call binding the contract method 0x72607537.
//
// Solidity: function localEid() view returns(uint32)
func (_EndpointV2View *EndpointV2ViewCaller) LocalEid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "localEid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalEid is a free data retrieval call binding the contract method 0x72607537.
//
// Solidity: function localEid() view returns(uint32)
func (_EndpointV2View *EndpointV2ViewSession) LocalEid() (uint32, error) {
	return _EndpointV2View.Contract.LocalEid(&_EndpointV2View.CallOpts)
}

// LocalEid is a free data retrieval call binding the contract method 0x72607537.
//
// Solidity: function localEid() view returns(uint32)
func (_EndpointV2View *EndpointV2ViewCallerSession) LocalEid() (uint32, error) {
	return _EndpointV2View.Contract.LocalEid(&_EndpointV2View.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EndpointV2View *EndpointV2ViewCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EndpointV2View *EndpointV2ViewSession) Owner() (common.Address, error) {
	return _EndpointV2View.Contract.Owner(&_EndpointV2View.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EndpointV2View *EndpointV2ViewCallerSession) Owner() (common.Address, error) {
	return _EndpointV2View.Contract.Owner(&_EndpointV2View.CallOpts)
}

// ReceiveLibToView is a free data retrieval call binding the contract method 0x6f178835.
//
// Solidity: function receiveLibToView(address receiveLib) view returns(address receiveLibView)
func (_EndpointV2View *EndpointV2ViewCaller) ReceiveLibToView(opts *bind.CallOpts, receiveLib common.Address) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "receiveLibToView", receiveLib)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiveLibToView is a free data retrieval call binding the contract method 0x6f178835.
//
// Solidity: function receiveLibToView(address receiveLib) view returns(address receiveLibView)
func (_EndpointV2View *EndpointV2ViewSession) ReceiveLibToView(receiveLib common.Address) (common.Address, error) {
	return _EndpointV2View.Contract.ReceiveLibToView(&_EndpointV2View.CallOpts, receiveLib)
}

// ReceiveLibToView is a free data retrieval call binding the contract method 0x6f178835.
//
// Solidity: function receiveLibToView(address receiveLib) view returns(address receiveLibView)
func (_EndpointV2View *EndpointV2ViewCallerSession) ReceiveLibToView(receiveLib common.Address) (common.Address, error) {
	return _EndpointV2View.Contract.ReceiveLibToView(&_EndpointV2View.CallOpts, receiveLib)
}

// ReceiveUln302 is a free data retrieval call binding the contract method 0x843c7b0e.
//
// Solidity: function receiveUln302() view returns(address)
func (_EndpointV2View *EndpointV2ViewCaller) ReceiveUln302(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "receiveUln302")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiveUln302 is a free data retrieval call binding the contract method 0x843c7b0e.
//
// Solidity: function receiveUln302() view returns(address)
func (_EndpointV2View *EndpointV2ViewSession) ReceiveUln302() (common.Address, error) {
	return _EndpointV2View.Contract.ReceiveUln302(&_EndpointV2View.CallOpts)
}

// ReceiveUln302 is a free data retrieval call binding the contract method 0x843c7b0e.
//
// Solidity: function receiveUln302() view returns(address)
func (_EndpointV2View *EndpointV2ViewCallerSession) ReceiveUln302() (common.Address, error) {
	return _EndpointV2View.Contract.ReceiveUln302(&_EndpointV2View.CallOpts)
}

// Verifiable is a free data retrieval call binding the contract method 0xe1e3a7df.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver, address _receiveLib, bytes32 _payloadHash) view returns(bool)
func (_EndpointV2View *EndpointV2ViewCaller) Verifiable(opts *bind.CallOpts, _origin Origin, _receiver common.Address, _receiveLib common.Address, _payloadHash [32]byte) (bool, error) {
	var out []interface{}
	err := _EndpointV2View.contract.Call(opts, &out, "verifiable", _origin, _receiver, _receiveLib, _payloadHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verifiable is a free data retrieval call binding the contract method 0xe1e3a7df.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver, address _receiveLib, bytes32 _payloadHash) view returns(bool)
func (_EndpointV2View *EndpointV2ViewSession) Verifiable(_origin Origin, _receiver common.Address, _receiveLib common.Address, _payloadHash [32]byte) (bool, error) {
	return _EndpointV2View.Contract.Verifiable(&_EndpointV2View.CallOpts, _origin, _receiver, _receiveLib, _payloadHash)
}

// Verifiable is a free data retrieval call binding the contract method 0xe1e3a7df.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver, address _receiveLib, bytes32 _payloadHash) view returns(bool)
func (_EndpointV2View *EndpointV2ViewCallerSession) Verifiable(_origin Origin, _receiver common.Address, _receiveLib common.Address, _payloadHash [32]byte) (bool, error) {
	return _EndpointV2View.Contract.Verifiable(&_EndpointV2View.CallOpts, _origin, _receiver, _receiveLib, _payloadHash)
}

// CommitAndExecute is a paid mutator transaction binding the contract method 0xdcfdeb60.
//
// Solidity: function commitAndExecute(address _receiveLib, ((uint32,bytes32,uint64),address,bytes32,bytes,bytes,uint256,uint256) _lzReceiveParam, (address,uint256)[] _nativeDropParams) payable returns()
func (_EndpointV2View *EndpointV2ViewTransactor) CommitAndExecute(opts *bind.TransactOpts, _receiveLib common.Address, _lzReceiveParam LzReceiveParam, _nativeDropParams []NativeDropParam) (*types.Transaction, error) {
	return _EndpointV2View.contract.Transact(opts, "commitAndExecute", _receiveLib, _lzReceiveParam, _nativeDropParams)
}

// CommitAndExecute is a paid mutator transaction binding the contract method 0xdcfdeb60.
//
// Solidity: function commitAndExecute(address _receiveLib, ((uint32,bytes32,uint64),address,bytes32,bytes,bytes,uint256,uint256) _lzReceiveParam, (address,uint256)[] _nativeDropParams) payable returns()
func (_EndpointV2View *EndpointV2ViewSession) CommitAndExecute(_receiveLib common.Address, _lzReceiveParam LzReceiveParam, _nativeDropParams []NativeDropParam) (*types.Transaction, error) {
	return _EndpointV2View.Contract.CommitAndExecute(&_EndpointV2View.TransactOpts, _receiveLib, _lzReceiveParam, _nativeDropParams)
}

// CommitAndExecute is a paid mutator transaction binding the contract method 0xdcfdeb60.
//
// Solidity: function commitAndExecute(address _receiveLib, ((uint32,bytes32,uint64),address,bytes32,bytes,bytes,uint256,uint256) _lzReceiveParam, (address,uint256)[] _nativeDropParams) payable returns()
func (_EndpointV2View *EndpointV2ViewTransactorSession) CommitAndExecute(_receiveLib common.Address, _lzReceiveParam LzReceiveParam, _nativeDropParams []NativeDropParam) (*types.Transaction, error) {
	return _EndpointV2View.Contract.CommitAndExecute(&_EndpointV2View.TransactOpts, _receiveLib, _lzReceiveParam, _nativeDropParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _receiveUln302, address _receiveUln302View, address _endpoint) returns()
func (_EndpointV2View *EndpointV2ViewTransactor) Initialize(opts *bind.TransactOpts, _receiveUln302 common.Address, _receiveUln302View common.Address, _endpoint common.Address) (*types.Transaction, error) {
	return _EndpointV2View.contract.Transact(opts, "initialize", _receiveUln302, _receiveUln302View, _endpoint)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _receiveUln302, address _receiveUln302View, address _endpoint) returns()
func (_EndpointV2View *EndpointV2ViewSession) Initialize(_receiveUln302 common.Address, _receiveUln302View common.Address, _endpoint common.Address) (*types.Transaction, error) {
	return _EndpointV2View.Contract.Initialize(&_EndpointV2View.TransactOpts, _receiveUln302, _receiveUln302View, _endpoint)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _receiveUln302, address _receiveUln302View, address _endpoint) returns()
func (_EndpointV2View *EndpointV2ViewTransactorSession) Initialize(_receiveUln302 common.Address, _receiveUln302View common.Address, _endpoint common.Address) (*types.Transaction, error) {
	return _EndpointV2View.Contract.Initialize(&_EndpointV2View.TransactOpts, _receiveUln302, _receiveUln302View, _endpoint)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EndpointV2View *EndpointV2ViewTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EndpointV2View.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EndpointV2View *EndpointV2ViewSession) RenounceOwnership() (*types.Transaction, error) {
	return _EndpointV2View.Contract.RenounceOwnership(&_EndpointV2View.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EndpointV2View *EndpointV2ViewTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EndpointV2View.Contract.RenounceOwnership(&_EndpointV2View.TransactOpts)
}

// SetReceiveLibView is a paid mutator transaction binding the contract method 0xd73d37c6.
//
// Solidity: function setReceiveLibView(address _receiveLib, address _receiveLibView) returns()
func (_EndpointV2View *EndpointV2ViewTransactor) SetReceiveLibView(opts *bind.TransactOpts, _receiveLib common.Address, _receiveLibView common.Address) (*types.Transaction, error) {
	return _EndpointV2View.contract.Transact(opts, "setReceiveLibView", _receiveLib, _receiveLibView)
}

// SetReceiveLibView is a paid mutator transaction binding the contract method 0xd73d37c6.
//
// Solidity: function setReceiveLibView(address _receiveLib, address _receiveLibView) returns()
func (_EndpointV2View *EndpointV2ViewSession) SetReceiveLibView(_receiveLib common.Address, _receiveLibView common.Address) (*types.Transaction, error) {
	return _EndpointV2View.Contract.SetReceiveLibView(&_EndpointV2View.TransactOpts, _receiveLib, _receiveLibView)
}

// SetReceiveLibView is a paid mutator transaction binding the contract method 0xd73d37c6.
//
// Solidity: function setReceiveLibView(address _receiveLib, address _receiveLibView) returns()
func (_EndpointV2View *EndpointV2ViewTransactorSession) SetReceiveLibView(_receiveLib common.Address, _receiveLibView common.Address) (*types.Transaction, error) {
	return _EndpointV2View.Contract.SetReceiveLibView(&_EndpointV2View.TransactOpts, _receiveLib, _receiveLibView)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EndpointV2View *EndpointV2ViewTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EndpointV2View.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EndpointV2View *EndpointV2ViewSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EndpointV2View.Contract.TransferOwnership(&_EndpointV2View.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EndpointV2View *EndpointV2ViewTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EndpointV2View.Contract.TransferOwnership(&_EndpointV2View.TransactOpts, newOwner)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address _to, uint256 _amount) returns()
func (_EndpointV2View *EndpointV2ViewTransactor) WithdrawNative(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EndpointV2View.contract.Transact(opts, "withdrawNative", _to, _amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address _to, uint256 _amount) returns()
func (_EndpointV2View *EndpointV2ViewSession) WithdrawNative(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EndpointV2View.Contract.WithdrawNative(&_EndpointV2View.TransactOpts, _to, _amount)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x07b18bde.
//
// Solidity: function withdrawNative(address _to, uint256 _amount) returns()
func (_EndpointV2View *EndpointV2ViewTransactorSession) WithdrawNative(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _EndpointV2View.Contract.WithdrawNative(&_EndpointV2View.TransactOpts, _to, _amount)
}

// EndpointV2ViewInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EndpointV2View contract.
type EndpointV2ViewInitializedIterator struct {
	Event *EndpointV2ViewInitialized // Event containing the contract specifics and raw log

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
func (it *EndpointV2ViewInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ViewInitialized)
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
		it.Event = new(EndpointV2ViewInitialized)
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
func (it *EndpointV2ViewInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ViewInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ViewInitialized represents a Initialized event raised by the EndpointV2View contract.
type EndpointV2ViewInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EndpointV2View *EndpointV2ViewFilterer) FilterInitialized(opts *bind.FilterOpts) (*EndpointV2ViewInitializedIterator, error) {

	logs, sub, err := _EndpointV2View.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EndpointV2ViewInitializedIterator{contract: _EndpointV2View.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EndpointV2View *EndpointV2ViewFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EndpointV2ViewInitialized) (event.Subscription, error) {

	logs, sub, err := _EndpointV2View.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ViewInitialized)
				if err := _EndpointV2View.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EndpointV2View *EndpointV2ViewFilterer) ParseInitialized(log types.Log) (*EndpointV2ViewInitialized, error) {
	event := new(EndpointV2ViewInitialized)
	if err := _EndpointV2View.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2ViewNativeWithdrawnIterator is returned from FilterNativeWithdrawn and is used to iterate over the raw logs and unpacked data for NativeWithdrawn events raised by the EndpointV2View contract.
type EndpointV2ViewNativeWithdrawnIterator struct {
	Event *EndpointV2ViewNativeWithdrawn // Event containing the contract specifics and raw log

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
func (it *EndpointV2ViewNativeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ViewNativeWithdrawn)
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
		it.Event = new(EndpointV2ViewNativeWithdrawn)
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
func (it *EndpointV2ViewNativeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ViewNativeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ViewNativeWithdrawn represents a NativeWithdrawn event raised by the EndpointV2View contract.
type EndpointV2ViewNativeWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeWithdrawn is a free log retrieval operation binding the contract event 0xc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed504.
//
// Solidity: event NativeWithdrawn(address _to, uint256 _amount)
func (_EndpointV2View *EndpointV2ViewFilterer) FilterNativeWithdrawn(opts *bind.FilterOpts) (*EndpointV2ViewNativeWithdrawnIterator, error) {

	logs, sub, err := _EndpointV2View.contract.FilterLogs(opts, "NativeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &EndpointV2ViewNativeWithdrawnIterator{contract: _EndpointV2View.contract, event: "NativeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNativeWithdrawn is a free log subscription operation binding the contract event 0xc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed504.
//
// Solidity: event NativeWithdrawn(address _to, uint256 _amount)
func (_EndpointV2View *EndpointV2ViewFilterer) WatchNativeWithdrawn(opts *bind.WatchOpts, sink chan<- *EndpointV2ViewNativeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _EndpointV2View.contract.WatchLogs(opts, "NativeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ViewNativeWithdrawn)
				if err := _EndpointV2View.contract.UnpackLog(event, "NativeWithdrawn", log); err != nil {
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

// ParseNativeWithdrawn is a log parse operation binding the contract event 0xc303ca808382409472acbbf899c316cf439f409f6584aae22df86dfa3c9ed504.
//
// Solidity: event NativeWithdrawn(address _to, uint256 _amount)
func (_EndpointV2View *EndpointV2ViewFilterer) ParseNativeWithdrawn(log types.Log) (*EndpointV2ViewNativeWithdrawn, error) {
	event := new(EndpointV2ViewNativeWithdrawn)
	if err := _EndpointV2View.contract.UnpackLog(event, "NativeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2ViewOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EndpointV2View contract.
type EndpointV2ViewOwnershipTransferredIterator struct {
	Event *EndpointV2ViewOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EndpointV2ViewOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ViewOwnershipTransferred)
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
		it.Event = new(EndpointV2ViewOwnershipTransferred)
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
func (it *EndpointV2ViewOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ViewOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ViewOwnershipTransferred represents a OwnershipTransferred event raised by the EndpointV2View contract.
type EndpointV2ViewOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EndpointV2View *EndpointV2ViewFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EndpointV2ViewOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EndpointV2View.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EndpointV2ViewOwnershipTransferredIterator{contract: _EndpointV2View.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EndpointV2View *EndpointV2ViewFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EndpointV2ViewOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EndpointV2View.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ViewOwnershipTransferred)
				if err := _EndpointV2View.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EndpointV2View *EndpointV2ViewFilterer) ParseOwnershipTransferred(log types.Log) (*EndpointV2ViewOwnershipTransferred, error) {
	event := new(EndpointV2ViewOwnershipTransferred)
	if err := _EndpointV2View.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EndpointV2ViewReceiveLibViewSetIterator is returned from FilterReceiveLibViewSet and is used to iterate over the raw logs and unpacked data for ReceiveLibViewSet events raised by the EndpointV2View contract.
type EndpointV2ViewReceiveLibViewSetIterator struct {
	Event *EndpointV2ViewReceiveLibViewSet // Event containing the contract specifics and raw log

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
func (it *EndpointV2ViewReceiveLibViewSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EndpointV2ViewReceiveLibViewSet)
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
		it.Event = new(EndpointV2ViewReceiveLibViewSet)
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
func (it *EndpointV2ViewReceiveLibViewSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EndpointV2ViewReceiveLibViewSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EndpointV2ViewReceiveLibViewSet represents a ReceiveLibViewSet event raised by the EndpointV2View contract.
type EndpointV2ViewReceiveLibViewSet struct {
	ReceiveLib     common.Address
	ReceiveLibView common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReceiveLibViewSet is a free log retrieval operation binding the contract event 0x142c46535a86ac791981f3f16bdfd58291f3f03fc3fd111646f3f0e4eb326b63.
//
// Solidity: event ReceiveLibViewSet(address _receiveLib, address _receiveLibView)
func (_EndpointV2View *EndpointV2ViewFilterer) FilterReceiveLibViewSet(opts *bind.FilterOpts) (*EndpointV2ViewReceiveLibViewSetIterator, error) {

	logs, sub, err := _EndpointV2View.contract.FilterLogs(opts, "ReceiveLibViewSet")
	if err != nil {
		return nil, err
	}
	return &EndpointV2ViewReceiveLibViewSetIterator{contract: _EndpointV2View.contract, event: "ReceiveLibViewSet", logs: logs, sub: sub}, nil
}

// WatchReceiveLibViewSet is a free log subscription operation binding the contract event 0x142c46535a86ac791981f3f16bdfd58291f3f03fc3fd111646f3f0e4eb326b63.
//
// Solidity: event ReceiveLibViewSet(address _receiveLib, address _receiveLibView)
func (_EndpointV2View *EndpointV2ViewFilterer) WatchReceiveLibViewSet(opts *bind.WatchOpts, sink chan<- *EndpointV2ViewReceiveLibViewSet) (event.Subscription, error) {

	logs, sub, err := _EndpointV2View.contract.WatchLogs(opts, "ReceiveLibViewSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EndpointV2ViewReceiveLibViewSet)
				if err := _EndpointV2View.contract.UnpackLog(event, "ReceiveLibViewSet", log); err != nil {
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

// ParseReceiveLibViewSet is a log parse operation binding the contract event 0x142c46535a86ac791981f3f16bdfd58291f3f03fc3fd111646f3f0e4eb326b63.
//
// Solidity: event ReceiveLibViewSet(address _receiveLib, address _receiveLibView)
func (_EndpointV2View *EndpointV2ViewFilterer) ParseReceiveLibViewSet(log types.Log) (*EndpointV2ViewReceiveLibViewSet, error) {
	event := new(EndpointV2ViewReceiveLibViewSet)
	if err := _EndpointV2View.contract.UnpackLog(event, "ReceiveLibViewSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
