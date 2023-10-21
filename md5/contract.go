// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package md5

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/precompile/contract"
	"github.com/ava-labs/subnet-evm/vmerrs"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// Gas costs for each function. These are set to 1 by default.
	// You should set a gas cost for each function in your contract.
	// Generally, you should not set gas costs very low as this may cause your network to be vulnerable to DoS attacks.
	// There are some predefined gas costs in contract/utils.go that you can use.
	HashWithMD5GasCost uint64 = 1 /* SET A GAS COST HERE */
)

// CUSTOM CODE STARTS HERE
// Reference imports to suppress errors from unused imports. This code and any unnecessary imports can be removed.
var (
	_ = abi.JSON
	_ = errors.New
	_ = big.NewInt
	_ = vmerrs.ErrOutOfGas
	_ = common.Big0
)

// Singleton StatefulPrecompiledContract and signatures.
var (

	// Md5RawABI contains the raw ABI of Md5 contract.
	//go:embed contract.abi
	Md5RawABI string

	Md5ABI = contract.ParseABI(Md5RawABI)

	Md5Precompile = createMd5Precompile()
)

// UnpackHashWithMD5Input attempts to unpack [input] into the string type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackHashWithMD5Input(input []byte) (string, error) {
	res, err := Md5ABI.UnpackInput("hashWithMD5", input)
	if err != nil {
		return "", err
	}
	unpacked := *abi.ConvertType(res[0], new(string)).(*string)
	return unpacked, nil
}

// PackHashWithMD5 packs [value] of type string into the appropriate arguments for hashWithMD5.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackHashWithMD5(value string) ([]byte, error) {
	return Md5ABI.Pack("hashWithMD5", value)
}

// PackHashWithMD5Output attempts to pack given hash of type [16]byte
// to conform the ABI outputs.
func PackHashWithMD5Output(hash [16]byte) ([]byte, error) {
	return Md5ABI.PackOutput("hashWithMD5", hash)
}

// UnpackHashWithMD5Output attempts to unpack given [output] into the [16]byte type output
// assumes that [output] does not include selector (omits first 4 func signature bytes)
func UnpackHashWithMD5Output(output []byte) ([16]byte, error) {
	res, err := Md5ABI.Unpack("hashWithMD5", output)
	if err != nil {
		return [16]byte{}, err
	}
	unpacked := *abi.ConvertType(res[0], new([16]byte)).(*[16]byte)
	return unpacked, nil
}

func hashWithMD5(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, HashWithMD5GasCost); err != nil {
		return nil, 0, err
	}
	// attempts to unpack [input] into the arguments to the HashWithMD5Input.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackHashWithMD5Input(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	_ = inputStruct // CUSTOM CODE OPERATES ON INPUT

	var output [16]byte // CUSTOM CODE FOR AN OUTPUT
	output = md5.Sum([]byte(inputStruct))
	packedOutput, err := PackHashWithMD5Output(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// createMd5Precompile returns a StatefulPrecompiledContract with getters and setters for the precompile.

func createMd5Precompile() contract.StatefulPrecompiledContract {
	var functions []*contract.StatefulPrecompileFunction

	abiFunctionMap := map[string]contract.RunStatefulPrecompileFunc{
		"hashWithMD5": hashWithMD5,
	}

	for name, function := range abiFunctionMap {
		method, ok := Md5ABI.Methods[name]
		if !ok {
			panic(fmt.Errorf("given method (%s) does not exist in the ABI", name))
		}
		functions = append(functions, contract.NewStatefulPrecompileFunction(method.ID, function))
	}
	// Construct the contract with no fallback function.
	statefulContract, err := contract.NewStatefulPrecompileContract(nil, functions)
	if err != nil {
		panic(err)
	}
	return statefulContract
}
