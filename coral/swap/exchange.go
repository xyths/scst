// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

import (
	"math/big"
	"strings"

	sero "github.com/sero-cash/go-sero"
	"github.com/sero-cash/go-sero/accounts/abi"
	"github.com/sero-cash/go-sero/accounts/abi/bind"
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/core/types"
	"github.com/sero-cash/go-sero/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = sero.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Liquidity is an auto generated low-level Go binding around an user-defined struct.
type Liquidity struct {
	Value     *big.Int
	NextValue *big.Int
	NextIndex *big.Int
	PrevIndex *big.Int
	Flag      bool
}

// Pair is an auto generated low-level Go binding around an user-defined struct.
type Pair struct {
	TokenA       [32]byte
	TokenB       [32]byte
	ReserveA     *big.Int
	ReserveB     *big.Int
	TotalShares  *big.Int
	MyShare      *big.Int
	ShareRreward *big.Int
	Mining       bool
}

// Volume is an auto generated low-level Go binding around an user-defined struct.
type Volume struct {
	Index *big.Int
	Value *big.Int
}

// SwapExchangeABI is the input ABI used to generate the binding from.
const SwapExchangeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenPool\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"cancelInvest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_sharesBurned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minTokenA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minTokenB\",\"type\":\"uint256\"}],\"name\":\"divestLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"tokenIn\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"estimateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"tokenOut\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"estimateSwapBuy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"feeRateMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"tokens\",\"type\":\"bytes32[]\"}],\"name\":\"getGroupTokens\",\"outputs\":[{\"internalType\":\"bytes32[][]\",\"name\":\"rets\",\"type\":\"bytes32[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"}],\"name\":\"getTokens\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"rets\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"hasPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tokenA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"tokenB\",\"type\":\"bytes32\"}],\"name\":\"hashKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"investAmount\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minShares\",\"type\":\"uint256\"}],\"name\":\"investLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lastIndexsMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenB\",\"type\":\"string\"}],\"name\":\"liquidityOfPair\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"internalType\":\"structLiquidity[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevIndex\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"internalType\":\"structLiquidity[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintDayIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"output\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"pairInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"tokenA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"tokenB\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"myShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shareRreward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"mining\",\"type\":\"bool\"}],\"internalType\":\"structPair\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pairKeys\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"pairList\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"tokenA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"tokenB\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"myShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shareRreward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"mining\",\"type\":\"bool\"}],\"internalType\":\"structPair[]\",\"name\":\"rets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"pairListByToken\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"tokenA\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"tokenB\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"myShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shareRreward\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"mining\",\"type\":\"bool\"}],\"internalType\":\"structPair[]\",\"name\":\"rets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"rateMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenB\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenB\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_baseToken\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setTokenBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"shareReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_minTokensReceived\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenB\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"volumeDayOfPair\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenA\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenB\",\"type\":\"string\"}],\"name\":\"volumesOfPair\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structVolume[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structVolume[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"withdrawShareReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// SwapExchange is an auto generated Go binding around an Ethereum contract.
type SwapExchange struct {
	SwapExchangeCaller     // Read-only binding to the contract
	SwapExchangeTransactor // Write-only binding to the contract
	SwapExchangeFilterer   // Log filterer for contract events
}

// SwapExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapExchangeSession struct {
	Contract     *SwapExchange     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapExchangeCallerSession struct {
	Contract *SwapExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SwapExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapExchangeTransactorSession struct {
	Contract     *SwapExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SwapExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapExchangeRaw struct {
	Contract *SwapExchange // Generic contract binding to access the raw methods on
}

// SwapExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapExchangeCallerRaw struct {
	Contract *SwapExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// SwapExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapExchangeTransactorRaw struct {
	Contract *SwapExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapExchange creates a new instance of SwapExchange, bound to a specific deployed contract.
func NewSwapExchange(address common.Address, backend bind.ContractBackend) (*SwapExchange, error) {
	contract, err := bindSwapExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapExchange{SwapExchangeCaller: SwapExchangeCaller{contract: contract}, SwapExchangeTransactor: SwapExchangeTransactor{contract: contract}, SwapExchangeFilterer: SwapExchangeFilterer{contract: contract}}, nil
}

// NewSwapExchangeCaller creates a new read-only instance of SwapExchange, bound to a specific deployed contract.
func NewSwapExchangeCaller(address common.Address, caller bind.ContractCaller) (*SwapExchangeCaller, error) {
	contract, err := bindSwapExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapExchangeCaller{contract: contract}, nil
}

// NewSwapExchangeTransactor creates a new write-only instance of SwapExchange, bound to a specific deployed contract.
func NewSwapExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapExchangeTransactor, error) {
	contract, err := bindSwapExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapExchangeTransactor{contract: contract}, nil
}

// NewSwapExchangeFilterer creates a new log filterer instance of SwapExchange, bound to a specific deployed contract.
func NewSwapExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapExchangeFilterer, error) {
	contract, err := bindSwapExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapExchangeFilterer{contract: contract}, nil
}

// bindSwapExchange binds a generic wrapper to an already deployed contract.
func bindSwapExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapExchange *SwapExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwapExchange.Contract.SwapExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapExchange *SwapExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExchange.Contract.SwapExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapExchange *SwapExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapExchange.Contract.SwapExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapExchange *SwapExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SwapExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapExchange *SwapExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapExchange *SwapExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapExchange.Contract.contract.Transact(opts, method, params...)
}

// EstimateSwap is a free data retrieval call binding the contract method 0x737f4194.
//
// Solidity: function estimateSwap(bytes32 key, bytes32 tokenIn, uint256 amountIn) view returns(uint256 value)
func (_SwapExchange *SwapExchangeCaller) EstimateSwap(opts *bind.CallOpts, key [32]byte, tokenIn [32]byte, amountIn *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "estimateSwap", key, tokenIn, amountIn)
	return *ret0, err
}

// EstimateSwap is a free data retrieval call binding the contract method 0x737f4194.
//
// Solidity: function estimateSwap(bytes32 key, bytes32 tokenIn, uint256 amountIn) view returns(uint256 value)
func (_SwapExchange *SwapExchangeSession) EstimateSwap(key [32]byte, tokenIn [32]byte, amountIn *big.Int) (*big.Int, error) {
	return _SwapExchange.Contract.EstimateSwap(&_SwapExchange.CallOpts, key, tokenIn, amountIn)
}

// EstimateSwap is a free data retrieval call binding the contract method 0x737f4194.
//
// Solidity: function estimateSwap(bytes32 key, bytes32 tokenIn, uint256 amountIn) view returns(uint256 value)
func (_SwapExchange *SwapExchangeCallerSession) EstimateSwap(key [32]byte, tokenIn [32]byte, amountIn *big.Int) (*big.Int, error) {
	return _SwapExchange.Contract.EstimateSwap(&_SwapExchange.CallOpts, key, tokenIn, amountIn)
}

// EstimateSwapBuy is a free data retrieval call binding the contract method 0xf5e1f98c.
//
// Solidity: function estimateSwapBuy(bytes32 key, bytes32 tokenOut, uint256 amountOut) view returns(uint256 amountIn)
func (_SwapExchange *SwapExchangeCaller) EstimateSwapBuy(opts *bind.CallOpts, key [32]byte, tokenOut [32]byte, amountOut *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "estimateSwapBuy", key, tokenOut, amountOut)
	return *ret0, err
}

// EstimateSwapBuy is a free data retrieval call binding the contract method 0xf5e1f98c.
//
// Solidity: function estimateSwapBuy(bytes32 key, bytes32 tokenOut, uint256 amountOut) view returns(uint256 amountIn)
func (_SwapExchange *SwapExchangeSession) EstimateSwapBuy(key [32]byte, tokenOut [32]byte, amountOut *big.Int) (*big.Int, error) {
	return _SwapExchange.Contract.EstimateSwapBuy(&_SwapExchange.CallOpts, key, tokenOut, amountOut)
}

// EstimateSwapBuy is a free data retrieval call binding the contract method 0xf5e1f98c.
//
// Solidity: function estimateSwapBuy(bytes32 key, bytes32 tokenOut, uint256 amountOut) view returns(uint256 amountIn)
func (_SwapExchange *SwapExchangeCallerSession) EstimateSwapBuy(key [32]byte, tokenOut [32]byte, amountOut *big.Int) (*big.Int, error) {
	return _SwapExchange.Contract.EstimateSwapBuy(&_SwapExchange.CallOpts, key, tokenOut, amountOut)
}

// FeeRateMap is a free data retrieval call binding the contract method 0xdfab082d.
//
// Solidity: function feeRateMap(bytes32 ) view returns(uint256)
func (_SwapExchange *SwapExchangeCaller) FeeRateMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "feeRateMap", arg0)
	return *ret0, err
}

// FeeRateMap is a free data retrieval call binding the contract method 0xdfab082d.
//
// Solidity: function feeRateMap(bytes32 ) view returns(uint256)
func (_SwapExchange *SwapExchangeSession) FeeRateMap(arg0 [32]byte) (*big.Int, error) {
	return _SwapExchange.Contract.FeeRateMap(&_SwapExchange.CallOpts, arg0)
}

// FeeRateMap is a free data retrieval call binding the contract method 0xdfab082d.
//
// Solidity: function feeRateMap(bytes32 ) view returns(uint256)
func (_SwapExchange *SwapExchangeCallerSession) FeeRateMap(arg0 [32]byte) (*big.Int, error) {
	return _SwapExchange.Contract.FeeRateMap(&_SwapExchange.CallOpts, arg0)
}

// GetGroupTokens is a free data retrieval call binding the contract method 0x20a3aa7e.
//
// Solidity: function getGroupTokens(bytes32[] tokens) view returns(bytes32[][] rets)
func (_SwapExchange *SwapExchangeCaller) GetGroupTokens(opts *bind.CallOpts, tokens [][32]byte) ([][][32]byte, error) {
	var (
		ret0 = new([][][32]byte)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "getGroupTokens", tokens)
	return *ret0, err
}

// GetGroupTokens is a free data retrieval call binding the contract method 0x20a3aa7e.
//
// Solidity: function getGroupTokens(bytes32[] tokens) view returns(bytes32[][] rets)
func (_SwapExchange *SwapExchangeSession) GetGroupTokens(tokens [][32]byte) ([][][32]byte, error) {
	return _SwapExchange.Contract.GetGroupTokens(&_SwapExchange.CallOpts, tokens)
}

// GetGroupTokens is a free data retrieval call binding the contract method 0x20a3aa7e.
//
// Solidity: function getGroupTokens(bytes32[] tokens) view returns(bytes32[][] rets)
func (_SwapExchange *SwapExchangeCallerSession) GetGroupTokens(tokens [][32]byte) ([][][32]byte, error) {
	return _SwapExchange.Contract.GetGroupTokens(&_SwapExchange.CallOpts, tokens)
}

// GetTokens is a free data retrieval call binding the contract method 0xecbbc033.
//
// Solidity: function getTokens(bytes32 token) view returns(bytes32[] rets)
func (_SwapExchange *SwapExchangeCaller) GetTokens(opts *bind.CallOpts, token [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "getTokens", token)
	return *ret0, err
}

// GetTokens is a free data retrieval call binding the contract method 0xecbbc033.
//
// Solidity: function getTokens(bytes32 token) view returns(bytes32[] rets)
func (_SwapExchange *SwapExchangeSession) GetTokens(token [32]byte) ([][32]byte, error) {
	return _SwapExchange.Contract.GetTokens(&_SwapExchange.CallOpts, token)
}

// GetTokens is a free data retrieval call binding the contract method 0xecbbc033.
//
// Solidity: function getTokens(bytes32 token) view returns(bytes32[] rets)
func (_SwapExchange *SwapExchangeCallerSession) GetTokens(token [32]byte) ([][32]byte, error) {
	return _SwapExchange.Contract.GetTokens(&_SwapExchange.CallOpts, token)
}

// HasPair is a free data retrieval call binding the contract method 0x610a5f3a.
//
// Solidity: function hasPair(bytes32 _key) view returns(bool)
func (_SwapExchange *SwapExchangeCaller) HasPair(opts *bind.CallOpts, _key [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "hasPair", _key)
	return *ret0, err
}

// HasPair is a free data retrieval call binding the contract method 0x610a5f3a.
//
// Solidity: function hasPair(bytes32 _key) view returns(bool)
func (_SwapExchange *SwapExchangeSession) HasPair(_key [32]byte) (bool, error) {
	return _SwapExchange.Contract.HasPair(&_SwapExchange.CallOpts, _key)
}

// HasPair is a free data retrieval call binding the contract method 0x610a5f3a.
//
// Solidity: function hasPair(bytes32 _key) view returns(bool)
func (_SwapExchange *SwapExchangeCallerSession) HasPair(_key [32]byte) (bool, error) {
	return _SwapExchange.Contract.HasPair(&_SwapExchange.CallOpts, _key)
}

// HashKey is a free data retrieval call binding the contract method 0x72091e9b.
//
// Solidity: function hashKey(bytes32 tokenA, bytes32 tokenB) pure returns(bytes32)
func (_SwapExchange *SwapExchangeCaller) HashKey(opts *bind.CallOpts, tokenA [32]byte, tokenB [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "hashKey", tokenA, tokenB)
	return *ret0, err
}

// HashKey is a free data retrieval call binding the contract method 0x72091e9b.
//
// Solidity: function hashKey(bytes32 tokenA, bytes32 tokenB) pure returns(bytes32)
func (_SwapExchange *SwapExchangeSession) HashKey(tokenA [32]byte, tokenB [32]byte) ([32]byte, error) {
	return _SwapExchange.Contract.HashKey(&_SwapExchange.CallOpts, tokenA, tokenB)
}

// HashKey is a free data retrieval call binding the contract method 0x72091e9b.
//
// Solidity: function hashKey(bytes32 tokenA, bytes32 tokenB) pure returns(bytes32)
func (_SwapExchange *SwapExchangeCallerSession) HashKey(tokenA [32]byte, tokenB [32]byte) ([32]byte, error) {
	return _SwapExchange.Contract.HashKey(&_SwapExchange.CallOpts, tokenA, tokenB)
}

// InvestAmount is a free data retrieval call binding the contract method 0x0311904c.
//
// Solidity: function investAmount() view returns(bytes32 token, uint256 value)
func (_SwapExchange *SwapExchangeCaller) InvestAmount(opts *bind.CallOpts) (struct {
	Token [32]byte
	Value *big.Int
}, error) {
	ret := new(struct {
		Token [32]byte
		Value *big.Int
	})
	out := ret
	err := _SwapExchange.contract.Call(opts, out, "investAmount")
	return *ret, err
}

// InvestAmount is a free data retrieval call binding the contract method 0x0311904c.
//
// Solidity: function investAmount() view returns(bytes32 token, uint256 value)
func (_SwapExchange *SwapExchangeSession) InvestAmount() (struct {
	Token [32]byte
	Value *big.Int
}, error) {
	return _SwapExchange.Contract.InvestAmount(&_SwapExchange.CallOpts)
}

// InvestAmount is a free data retrieval call binding the contract method 0x0311904c.
//
// Solidity: function investAmount() view returns(bytes32 token, uint256 value)
func (_SwapExchange *SwapExchangeCallerSession) InvestAmount() (struct {
	Token [32]byte
	Value *big.Int
}, error) {
	return _SwapExchange.Contract.InvestAmount(&_SwapExchange.CallOpts)
}

// LastIndexsMap is a free data retrieval call binding the contract method 0x7294460a.
//
// Solidity: function lastIndexsMap(address , bytes32 ) view returns(uint256)
func (_SwapExchange *SwapExchangeCaller) LastIndexsMap(opts *bind.CallOpts, arg0 common.ContractAddress, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "lastIndexsMap", arg0, arg1)
	return *ret0, err
}

// LastIndexsMap is a free data retrieval call binding the contract method 0x7294460a.
//
// Solidity: function lastIndexsMap(address , bytes32 ) view returns(uint256)
func (_SwapExchange *SwapExchangeSession) LastIndexsMap(arg0 common.ContractAddress, arg1 [32]byte) (*big.Int, error) {
	return _SwapExchange.Contract.LastIndexsMap(&_SwapExchange.CallOpts, arg0, arg1)
}

// LastIndexsMap is a free data retrieval call binding the contract method 0x7294460a.
//
// Solidity: function lastIndexsMap(address , bytes32 ) view returns(uint256)
func (_SwapExchange *SwapExchangeCallerSession) LastIndexsMap(arg0 common.ContractAddress, arg1 [32]byte) (*big.Int, error) {
	return _SwapExchange.Contract.LastIndexsMap(&_SwapExchange.CallOpts, arg0, arg1)
}

// LiquidityOfPair is a free data retrieval call binding the contract method 0xecdce0de.
//
// Solidity: function liquidityOfPair(string tokenA, string tokenB) view returns((uint256,uint256,uint256,uint256,bool)[], (uint256,uint256,uint256,uint256,bool)[])
func (_SwapExchange *SwapExchangeCaller) LiquidityOfPair(opts *bind.CallOpts, tokenA string, tokenB string) ([]Liquidity, []Liquidity, error) {
	var (
		ret0 = new([]Liquidity)
		ret1 = new([]Liquidity)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SwapExchange.contract.Call(opts, out, "liquidityOfPair", tokenA, tokenB)
	return *ret0, *ret1, err
}

// LiquidityOfPair is a free data retrieval call binding the contract method 0xecdce0de.
//
// Solidity: function liquidityOfPair(string tokenA, string tokenB) view returns((uint256,uint256,uint256,uint256,bool)[], (uint256,uint256,uint256,uint256,bool)[])
func (_SwapExchange *SwapExchangeSession) LiquidityOfPair(tokenA string, tokenB string) ([]Liquidity, []Liquidity, error) {
	return _SwapExchange.Contract.LiquidityOfPair(&_SwapExchange.CallOpts, tokenA, tokenB)
}

// LiquidityOfPair is a free data retrieval call binding the contract method 0xecdce0de.
//
// Solidity: function liquidityOfPair(string tokenA, string tokenB) view returns((uint256,uint256,uint256,uint256,bool)[], (uint256,uint256,uint256,uint256,bool)[])
func (_SwapExchange *SwapExchangeCallerSession) LiquidityOfPair(tokenA string, tokenB string) ([]Liquidity, []Liquidity, error) {
	return _SwapExchange.Contract.LiquidityOfPair(&_SwapExchange.CallOpts, tokenA, tokenB)
}

// MintDayIndex is a free data retrieval call binding the contract method 0x18e35fcf.
//
// Solidity: function mintDayIndex() view returns(uint256)
func (_SwapExchange *SwapExchangeCaller) MintDayIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "mintDayIndex")
	return *ret0, err
}

// MintDayIndex is a free data retrieval call binding the contract method 0x18e35fcf.
//
// Solidity: function mintDayIndex() view returns(uint256)
func (_SwapExchange *SwapExchangeSession) MintDayIndex() (*big.Int, error) {
	return _SwapExchange.Contract.MintDayIndex(&_SwapExchange.CallOpts)
}

// MintDayIndex is a free data retrieval call binding the contract method 0x18e35fcf.
//
// Solidity: function mintDayIndex() view returns(uint256)
func (_SwapExchange *SwapExchangeCallerSession) MintDayIndex() (*big.Int, error) {
	return _SwapExchange.Contract.MintDayIndex(&_SwapExchange.CallOpts)
}

// Output is a free data retrieval call binding the contract method 0x64f9016b.
//
// Solidity: function output(uint256 index) view returns(uint256)
func (_SwapExchange *SwapExchangeCaller) Output(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "output", index)
	return *ret0, err
}

// Output is a free data retrieval call binding the contract method 0x64f9016b.
//
// Solidity: function output(uint256 index) view returns(uint256)
func (_SwapExchange *SwapExchangeSession) Output(index *big.Int) (*big.Int, error) {
	return _SwapExchange.Contract.Output(&_SwapExchange.CallOpts, index)
}

// Output is a free data retrieval call binding the contract method 0x64f9016b.
//
// Solidity: function output(uint256 index) view returns(uint256)
func (_SwapExchange *SwapExchangeCallerSession) Output(index *big.Int) (*big.Int, error) {
	return _SwapExchange.Contract.Output(&_SwapExchange.CallOpts, index)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapExchange *SwapExchangeCaller) Owner(opts *bind.CallOpts) (common.ContractAddress, error) {
	var (
		ret0 = new(common.ContractAddress)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapExchange *SwapExchangeSession) Owner() (common.ContractAddress, error) {
	return _SwapExchange.Contract.Owner(&_SwapExchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapExchange *SwapExchangeCallerSession) Owner() (common.ContractAddress, error) {
	return _SwapExchange.Contract.Owner(&_SwapExchange.CallOpts)
}

// PairInfo is a free data retrieval call binding the contract method 0xabc48614.
//
// Solidity: function pairInfo(bytes32 key) view returns((bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bool))
func (_SwapExchange *SwapExchangeCaller) PairInfo(opts *bind.CallOpts, key [32]byte) (Pair, error) {
	var (
		ret0 = new(Pair)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "pairInfo", key)
	return *ret0, err
}

// PairInfo is a free data retrieval call binding the contract method 0xabc48614.
//
// Solidity: function pairInfo(bytes32 key) view returns((bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bool))
func (_SwapExchange *SwapExchangeSession) PairInfo(key [32]byte) (Pair, error) {
	return _SwapExchange.Contract.PairInfo(&_SwapExchange.CallOpts, key)
}

// PairInfo is a free data retrieval call binding the contract method 0xabc48614.
//
// Solidity: function pairInfo(bytes32 key) view returns((bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bool))
func (_SwapExchange *SwapExchangeCallerSession) PairInfo(key [32]byte) (Pair, error) {
	return _SwapExchange.Contract.PairInfo(&_SwapExchange.CallOpts, key)
}

// PairKeys is a free data retrieval call binding the contract method 0xad8c5668.
//
// Solidity: function pairKeys(uint256 ) view returns(bytes32)
func (_SwapExchange *SwapExchangeCaller) PairKeys(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "pairKeys", arg0)
	return *ret0, err
}

// PairKeys is a free data retrieval call binding the contract method 0xad8c5668.
//
// Solidity: function pairKeys(uint256 ) view returns(bytes32)
func (_SwapExchange *SwapExchangeSession) PairKeys(arg0 *big.Int) ([32]byte, error) {
	return _SwapExchange.Contract.PairKeys(&_SwapExchange.CallOpts, arg0)
}

// PairKeys is a free data retrieval call binding the contract method 0xad8c5668.
//
// Solidity: function pairKeys(uint256 ) view returns(bytes32)
func (_SwapExchange *SwapExchangeCallerSession) PairKeys(arg0 *big.Int) ([32]byte, error) {
	return _SwapExchange.Contract.PairKeys(&_SwapExchange.CallOpts, arg0)
}

// PairList is a free data retrieval call binding the contract method 0x435a508a.
//
// Solidity: function pairList(uint256 _start, uint256 _end) view returns((bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bool)[] rets)
func (_SwapExchange *SwapExchangeCaller) PairList(opts *bind.CallOpts, _start *big.Int, _end *big.Int) ([]Pair, error) {
	var (
		ret0 = new([]Pair)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "pairList", _start, _end)
	return *ret0, err
}

// PairList is a free data retrieval call binding the contract method 0x435a508a.
//
// Solidity: function pairList(uint256 _start, uint256 _end) view returns((bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bool)[] rets)
func (_SwapExchange *SwapExchangeSession) PairList(_start *big.Int, _end *big.Int) ([]Pair, error) {
	return _SwapExchange.Contract.PairList(&_SwapExchange.CallOpts, _start, _end)
}

// PairList is a free data retrieval call binding the contract method 0x435a508a.
//
// Solidity: function pairList(uint256 _start, uint256 _end) view returns((bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bool)[] rets)
func (_SwapExchange *SwapExchangeCallerSession) PairList(_start *big.Int, _end *big.Int) ([]Pair, error) {
	return _SwapExchange.Contract.PairList(&_SwapExchange.CallOpts, _start, _end)
}

// PairListByToken is a free data retrieval call binding the contract method 0x8336762f.
//
// Solidity: function pairListByToken(bytes32 token, uint256 _start, uint256 _end) view returns((bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bool)[] rets)
func (_SwapExchange *SwapExchangeCaller) PairListByToken(opts *bind.CallOpts, token [32]byte, _start *big.Int, _end *big.Int) ([]Pair, error) {
	var (
		ret0 = new([]Pair)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "pairListByToken", token, _start, _end)
	return *ret0, err
}

// PairListByToken is a free data retrieval call binding the contract method 0x8336762f.
//
// Solidity: function pairListByToken(bytes32 token, uint256 _start, uint256 _end) view returns((bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bool)[] rets)
func (_SwapExchange *SwapExchangeSession) PairListByToken(token [32]byte, _start *big.Int, _end *big.Int) ([]Pair, error) {
	return _SwapExchange.Contract.PairListByToken(&_SwapExchange.CallOpts, token, _start, _end)
}

// PairListByToken is a free data retrieval call binding the contract method 0x8336762f.
//
// Solidity: function pairListByToken(bytes32 token, uint256 _start, uint256 _end) view returns((bytes32,bytes32,uint256,uint256,uint256,uint256,uint256,bool)[] rets)
func (_SwapExchange *SwapExchangeCallerSession) PairListByToken(token [32]byte, _start *big.Int, _end *big.Int) ([]Pair, error) {
	return _SwapExchange.Contract.PairListByToken(&_SwapExchange.CallOpts, token, _start, _end)
}

// RateMap is a free data retrieval call binding the contract method 0x90b23000.
//
// Solidity: function rateMap(bytes32 ) view returns(uint256)
func (_SwapExchange *SwapExchangeCaller) RateMap(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "rateMap", arg0)
	return *ret0, err
}

// RateMap is a free data retrieval call binding the contract method 0x90b23000.
//
// Solidity: function rateMap(bytes32 ) view returns(uint256)
func (_SwapExchange *SwapExchangeSession) RateMap(arg0 [32]byte) (*big.Int, error) {
	return _SwapExchange.Contract.RateMap(&_SwapExchange.CallOpts, arg0)
}

// RateMap is a free data retrieval call binding the contract method 0x90b23000.
//
// Solidity: function rateMap(bytes32 ) view returns(uint256)
func (_SwapExchange *SwapExchangeCallerSession) RateMap(arg0 [32]byte) (*big.Int, error) {
	return _SwapExchange.Contract.RateMap(&_SwapExchange.CallOpts, arg0)
}

// ShareReward is a free data retrieval call binding the contract method 0xfbb453ca.
//
// Solidity: function shareReward(bytes32 key) view returns(uint256 reward)
func (_SwapExchange *SwapExchangeCaller) ShareReward(opts *bind.CallOpts, key [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "shareReward", key)
	return *ret0, err
}

// ShareReward is a free data retrieval call binding the contract method 0xfbb453ca.
//
// Solidity: function shareReward(bytes32 key) view returns(uint256 reward)
func (_SwapExchange *SwapExchangeSession) ShareReward(key [32]byte) (*big.Int, error) {
	return _SwapExchange.Contract.ShareReward(&_SwapExchange.CallOpts, key)
}

// ShareReward is a free data retrieval call binding the contract method 0xfbb453ca.
//
// Solidity: function shareReward(bytes32 key) view returns(uint256 reward)
func (_SwapExchange *SwapExchangeCallerSession) ShareReward(key [32]byte) (*big.Int, error) {
	return _SwapExchange.Contract.ShareReward(&_SwapExchange.CallOpts, key)
}

// StartDay is a free data retrieval call binding the contract method 0x0539272a.
//
// Solidity: function startDay() view returns(uint256)
func (_SwapExchange *SwapExchangeCaller) StartDay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SwapExchange.contract.Call(opts, out, "startDay")
	return *ret0, err
}

// StartDay is a free data retrieval call binding the contract method 0x0539272a.
//
// Solidity: function startDay() view returns(uint256)
func (_SwapExchange *SwapExchangeSession) StartDay() (*big.Int, error) {
	return _SwapExchange.Contract.StartDay(&_SwapExchange.CallOpts)
}

// StartDay is a free data retrieval call binding the contract method 0x0539272a.
//
// Solidity: function startDay() view returns(uint256)
func (_SwapExchange *SwapExchangeCallerSession) StartDay() (*big.Int, error) {
	return _SwapExchange.Contract.StartDay(&_SwapExchange.CallOpts)
}

// VolumeDayOfPair is a free data retrieval call binding the contract method 0xe94f8f2a.
//
// Solidity: function volumeDayOfPair(string tokenA, string tokenB, uint256 time) view returns(uint256, uint256)
func (_SwapExchange *SwapExchangeCaller) VolumeDayOfPair(opts *bind.CallOpts, tokenA string, tokenB string, time *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SwapExchange.contract.Call(opts, out, "volumeDayOfPair", tokenA, tokenB, time)
	return *ret0, *ret1, err
}

// VolumeDayOfPair is a free data retrieval call binding the contract method 0xe94f8f2a.
//
// Solidity: function volumeDayOfPair(string tokenA, string tokenB, uint256 time) view returns(uint256, uint256)
func (_SwapExchange *SwapExchangeSession) VolumeDayOfPair(tokenA string, tokenB string, time *big.Int) (*big.Int, *big.Int, error) {
	return _SwapExchange.Contract.VolumeDayOfPair(&_SwapExchange.CallOpts, tokenA, tokenB, time)
}

// VolumeDayOfPair is a free data retrieval call binding the contract method 0xe94f8f2a.
//
// Solidity: function volumeDayOfPair(string tokenA, string tokenB, uint256 time) view returns(uint256, uint256)
func (_SwapExchange *SwapExchangeCallerSession) VolumeDayOfPair(tokenA string, tokenB string, time *big.Int) (*big.Int, *big.Int, error) {
	return _SwapExchange.Contract.VolumeDayOfPair(&_SwapExchange.CallOpts, tokenA, tokenB, time)
}

// VolumesOfPair is a free data retrieval call binding the contract method 0x87399950.
//
// Solidity: function volumesOfPair(string tokenA, string tokenB) view returns((uint256,uint256)[], (uint256,uint256)[])
func (_SwapExchange *SwapExchangeCaller) VolumesOfPair(opts *bind.CallOpts, tokenA string, tokenB string) ([]Volume, []Volume, error) {
	var (
		ret0 = new([]Volume)
		ret1 = new([]Volume)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _SwapExchange.contract.Call(opts, out, "volumesOfPair", tokenA, tokenB)
	return *ret0, *ret1, err
}

// VolumesOfPair is a free data retrieval call binding the contract method 0x87399950.
//
// Solidity: function volumesOfPair(string tokenA, string tokenB) view returns((uint256,uint256)[], (uint256,uint256)[])
func (_SwapExchange *SwapExchangeSession) VolumesOfPair(tokenA string, tokenB string) ([]Volume, []Volume, error) {
	return _SwapExchange.Contract.VolumesOfPair(&_SwapExchange.CallOpts, tokenA, tokenB)
}

// VolumesOfPair is a free data retrieval call binding the contract method 0x87399950.
//
// Solidity: function volumesOfPair(string tokenA, string tokenB) view returns((uint256,uint256)[], (uint256,uint256)[])
func (_SwapExchange *SwapExchangeCallerSession) VolumesOfPair(tokenA string, tokenB string) ([]Volume, []Volume, error) {
	return _SwapExchange.Contract.VolumesOfPair(&_SwapExchange.CallOpts, tokenA, tokenB)
}

// CancelInvest is a paid mutator transaction binding the contract method 0xb4f81dcc.
//
// Solidity: function cancelInvest() returns(bool)
func (_SwapExchange *SwapExchangeTransactor) CancelInvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExchange.contract.Transact(opts, "cancelInvest")
}

// CancelInvest is a paid mutator transaction binding the contract method 0xb4f81dcc.
//
// Solidity: function cancelInvest() returns(bool)
func (_SwapExchange *SwapExchangeSession) CancelInvest() (*types.Transaction, error) {
	return _SwapExchange.Contract.CancelInvest(&_SwapExchange.TransactOpts)
}

// CancelInvest is a paid mutator transaction binding the contract method 0xb4f81dcc.
//
// Solidity: function cancelInvest() returns(bool)
func (_SwapExchange *SwapExchangeTransactorSession) CancelInvest() (*types.Transaction, error) {
	return _SwapExchange.Contract.CancelInvest(&_SwapExchange.TransactOpts)
}

// DivestLiquidity is a paid mutator transaction binding the contract method 0x756925ff.
//
// Solidity: function divestLiquidity(bytes32 key, uint256 _sharesBurned, uint256 _minTokenA, uint256 _minTokenB) returns()
func (_SwapExchange *SwapExchangeTransactor) DivestLiquidity(opts *bind.TransactOpts, key [32]byte, _sharesBurned *big.Int, _minTokenA *big.Int, _minTokenB *big.Int) (*types.Transaction, error) {
	return _SwapExchange.contract.Transact(opts, "divestLiquidity", key, _sharesBurned, _minTokenA, _minTokenB)
}

// DivestLiquidity is a paid mutator transaction binding the contract method 0x756925ff.
//
// Solidity: function divestLiquidity(bytes32 key, uint256 _sharesBurned, uint256 _minTokenA, uint256 _minTokenB) returns()
func (_SwapExchange *SwapExchangeSession) DivestLiquidity(key [32]byte, _sharesBurned *big.Int, _minTokenA *big.Int, _minTokenB *big.Int) (*types.Transaction, error) {
	return _SwapExchange.Contract.DivestLiquidity(&_SwapExchange.TransactOpts, key, _sharesBurned, _minTokenA, _minTokenB)
}

// DivestLiquidity is a paid mutator transaction binding the contract method 0x756925ff.
//
// Solidity: function divestLiquidity(bytes32 key, uint256 _sharesBurned, uint256 _minTokenA, uint256 _minTokenB) returns()
func (_SwapExchange *SwapExchangeTransactorSession) DivestLiquidity(key [32]byte, _sharesBurned *big.Int, _minTokenA *big.Int, _minTokenB *big.Int) (*types.Transaction, error) {
	return _SwapExchange.Contract.DivestLiquidity(&_SwapExchange.TransactOpts, key, _sharesBurned, _minTokenA, _minTokenB)
}

// InvestLiquidity is a paid mutator transaction binding the contract method 0x12675e01.
//
// Solidity: function investLiquidity(uint256 _minShares) payable returns()
func (_SwapExchange *SwapExchangeTransactor) InvestLiquidity(opts *bind.TransactOpts, _minShares *big.Int) (*types.Transaction, error) {
	return _SwapExchange.contract.Transact(opts, "investLiquidity", _minShares)
}

// InvestLiquidity is a paid mutator transaction binding the contract method 0x12675e01.
//
// Solidity: function investLiquidity(uint256 _minShares) payable returns()
func (_SwapExchange *SwapExchangeSession) InvestLiquidity(_minShares *big.Int) (*types.Transaction, error) {
	return _SwapExchange.Contract.InvestLiquidity(&_SwapExchange.TransactOpts, _minShares)
}

// InvestLiquidity is a paid mutator transaction binding the contract method 0x12675e01.
//
// Solidity: function investLiquidity(uint256 _minShares) payable returns()
func (_SwapExchange *SwapExchangeTransactorSession) InvestLiquidity(_minShares *big.Int) (*types.Transaction, error) {
	return _SwapExchange.Contract.InvestLiquidity(&_SwapExchange.TransactOpts, _minShares)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x9029feed.
//
// Solidity: function setFeeRate(string tokenA, string tokenB, uint256 _feeRate) returns()
func (_SwapExchange *SwapExchangeTransactor) SetFeeRate(opts *bind.TransactOpts, tokenA string, tokenB string, _feeRate *big.Int) (*types.Transaction, error) {
	return _SwapExchange.contract.Transact(opts, "setFeeRate", tokenA, tokenB, _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x9029feed.
//
// Solidity: function setFeeRate(string tokenA, string tokenB, uint256 _feeRate) returns()
func (_SwapExchange *SwapExchangeSession) SetFeeRate(tokenA string, tokenB string, _feeRate *big.Int) (*types.Transaction, error) {
	return _SwapExchange.Contract.SetFeeRate(&_SwapExchange.TransactOpts, tokenA, tokenB, _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x9029feed.
//
// Solidity: function setFeeRate(string tokenA, string tokenB, uint256 _feeRate) returns()
func (_SwapExchange *SwapExchangeTransactorSession) SetFeeRate(tokenA string, tokenB string, _feeRate *big.Int) (*types.Transaction, error) {
	return _SwapExchange.Contract.SetFeeRate(&_SwapExchange.TransactOpts, tokenA, tokenB, _feeRate)
}

// SetTokenBase is a paid mutator transaction binding the contract method 0x3ad7bacc.
//
// Solidity: function setTokenBase(string tokenA, string tokenB, string _baseToken, uint256 _rate) returns()
func (_SwapExchange *SwapExchangeTransactor) SetTokenBase(opts *bind.TransactOpts, tokenA string, tokenB string, _baseToken string, _rate *big.Int) (*types.Transaction, error) {
	return _SwapExchange.contract.Transact(opts, "setTokenBase", tokenA, tokenB, _baseToken, _rate)
}

// SetTokenBase is a paid mutator transaction binding the contract method 0x3ad7bacc.
//
// Solidity: function setTokenBase(string tokenA, string tokenB, string _baseToken, uint256 _rate) returns()
func (_SwapExchange *SwapExchangeSession) SetTokenBase(tokenA string, tokenB string, _baseToken string, _rate *big.Int) (*types.Transaction, error) {
	return _SwapExchange.Contract.SetTokenBase(&_SwapExchange.TransactOpts, tokenA, tokenB, _baseToken, _rate)
}

// SetTokenBase is a paid mutator transaction binding the contract method 0x3ad7bacc.
//
// Solidity: function setTokenBase(string tokenA, string tokenB, string _baseToken, uint256 _rate) returns()
func (_SwapExchange *SwapExchangeTransactorSession) SetTokenBase(tokenA string, tokenB string, _baseToken string, _rate *big.Int) (*types.Transaction, error) {
	return _SwapExchange.Contract.SetTokenBase(&_SwapExchange.TransactOpts, tokenA, tokenB, _baseToken, _rate)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_SwapExchange *SwapExchangeTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapExchange.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_SwapExchange *SwapExchangeSession) Start() (*types.Transaction, error) {
	return _SwapExchange.Contract.Start(&_SwapExchange.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_SwapExchange *SwapExchangeTransactorSession) Start() (*types.Transaction, error) {
	return _SwapExchange.Contract.Start(&_SwapExchange.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x82daa6f6.
//
// Solidity: function swap(bytes32 key, uint256 _minTokensReceived, uint256 _timeout, address _recipient) payable returns(uint256)
func (_SwapExchange *SwapExchangeTransactor) Swap(opts *bind.TransactOpts, key [32]byte, _minTokensReceived *big.Int, _timeout *big.Int, _recipient common.ContractAddress) (*types.Transaction, error) {
	return _SwapExchange.contract.Transact(opts, "swap", key, _minTokensReceived, _timeout, _recipient)
}

// Swap is a paid mutator transaction binding the contract method 0x82daa6f6.
//
// Solidity: function swap(bytes32 key, uint256 _minTokensReceived, uint256 _timeout, address _recipient) payable returns(uint256)
func (_SwapExchange *SwapExchangeSession) Swap(key [32]byte, _minTokensReceived *big.Int, _timeout *big.Int, _recipient common.ContractAddress) (*types.Transaction, error) {
	return _SwapExchange.Contract.Swap(&_SwapExchange.TransactOpts, key, _minTokensReceived, _timeout, _recipient)
}

// Swap is a paid mutator transaction binding the contract method 0x82daa6f6.
//
// Solidity: function swap(bytes32 key, uint256 _minTokensReceived, uint256 _timeout, address _recipient) payable returns(uint256)
func (_SwapExchange *SwapExchangeTransactorSession) Swap(key [32]byte, _minTokensReceived *big.Int, _timeout *big.Int, _recipient common.ContractAddress) (*types.Transaction, error) {
	return _SwapExchange.Contract.Swap(&_SwapExchange.TransactOpts, key, _minTokensReceived, _timeout, _recipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapExchange *SwapExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.ContractAddress) (*types.Transaction, error) {
	return _SwapExchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapExchange *SwapExchangeSession) TransferOwnership(newOwner common.ContractAddress) (*types.Transaction, error) {
	return _SwapExchange.Contract.TransferOwnership(&_SwapExchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapExchange *SwapExchangeTransactorSession) TransferOwnership(newOwner common.ContractAddress) (*types.Transaction, error) {
	return _SwapExchange.Contract.TransferOwnership(&_SwapExchange.TransactOpts, newOwner)
}

// WithdrawShareReward is a paid mutator transaction binding the contract method 0xc127a9b0.
//
// Solidity: function withdrawShareReward(bytes32 key) returns()
func (_SwapExchange *SwapExchangeTransactor) WithdrawShareReward(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _SwapExchange.contract.Transact(opts, "withdrawShareReward", key)
}

// WithdrawShareReward is a paid mutator transaction binding the contract method 0xc127a9b0.
//
// Solidity: function withdrawShareReward(bytes32 key) returns()
func (_SwapExchange *SwapExchangeSession) WithdrawShareReward(key [32]byte) (*types.Transaction, error) {
	return _SwapExchange.Contract.WithdrawShareReward(&_SwapExchange.TransactOpts, key)
}

// WithdrawShareReward is a paid mutator transaction binding the contract method 0xc127a9b0.
//
// Solidity: function withdrawShareReward(bytes32 key) returns()
func (_SwapExchange *SwapExchangeTransactorSession) WithdrawShareReward(key [32]byte) (*types.Transaction, error) {
	return _SwapExchange.Contract.WithdrawShareReward(&_SwapExchange.TransactOpts, key)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapExchange *SwapExchangeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	//return _SwapExchange.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
	return nil, nil
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapExchange *SwapExchangeSession) Receive() (*types.Transaction, error) {
	return _SwapExchange.Contract.Receive(&_SwapExchange.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapExchange *SwapExchangeTransactorSession) Receive() (*types.Transaction, error) {
	return _SwapExchange.Contract.Receive(&_SwapExchange.TransactOpts)
}

// SwapExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SwapExchange contract.
type SwapExchangeOwnershipTransferredIterator struct {
	Event *SwapExchangeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log    // Log channel receiving the found contract events
	sub  sero.Subscription // Subscription for errors, completion and termination
	done bool              // Whether the subscription completed delivering logs
	fail error             // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapExchangeOwnershipTransferred)
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
		it.Event = new(SwapExchangeOwnershipTransferred)
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
func (it *SwapExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the SwapExchange contract.
type SwapExchangeOwnershipTransferred struct {
	PreviousOwner common.ContractAddress
	NewOwner      common.ContractAddress
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapExchange *SwapExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.ContractAddress, newOwner []common.ContractAddress) (*SwapExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapExchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapExchangeOwnershipTransferredIterator{contract: _SwapExchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapExchange *SwapExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapExchangeOwnershipTransferred, previousOwner []common.ContractAddress, newOwner []common.ContractAddress) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapExchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapExchangeOwnershipTransferred)
				if err := _SwapExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SwapExchange *SwapExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*SwapExchangeOwnershipTransferred, error) {
	event := new(SwapExchangeOwnershipTransferred)
	if err := _SwapExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
