package types

import (
	"fmt"
	wasmUtils "github.com/CosmWasm/wasmd/x/wasm/client/utils"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"io/ioutil"
)

type (
	MsgStoreCode           = wasmtypes.MsgStoreCode
	MsgInstantiateContract = wasmtypes.MsgInstantiateContract
	MsgExecuteContract     = wasmtypes.MsgExecuteContract
	MsgMigrateContract     = wasmtypes.MsgMigrateContract
	MsgUpdateAdmin         = wasmtypes.MsgUpdateAdmin
	MsgClearAdmin          = wasmtypes.MsgClearAdmin

	QueryContractInfoRequest  = wasmtypes.QueryContractInfoRequest
	QueryContractInfoResponse = wasmtypes.QueryContractInfoResponse

	QueryContractHistoryRequest  = wasmtypes.QueryContractHistoryRequest
	QueryContractHistoryResponse = wasmtypes.QueryContractHistoryResponse

	QueryContractsByCodeRequest  = wasmtypes.QueryContractsByCodeRequest
	QueryContractsByCodeResponse = wasmtypes.QueryContractsByCodeResponse

	QueryAllContractStateRequest  = wasmtypes.QueryAllContractStateRequest
	QueryAllContractStateResponse = wasmtypes.QueryAllContractStateResponse

	QueryRawContractStateRequest  = wasmtypes.QueryRawContractStateRequest
	QueryRawContractStateResponse = wasmtypes.QueryRawContractStateResponse

	QuerySmartContractStateRequest  = wasmtypes.QuerySmartContractStateRequest
	QuerySmartContractStateResponse = wasmtypes.QuerySmartContractStateResponse

	QueryCodeRequest  = wasmtypes.QueryCodeRequest
	QueryCodeResponse = wasmtypes.QueryCodeResponse

	QueryCodesRequest  = wasmtypes.QueryCodesRequest
	QueryCodesResponse = wasmtypes.QueryCodesResponse

	QueryPinnedCodesRequest  = wasmtypes.QueryPinnedCodesRequest
	QueryPinnedCodesResponse = wasmtypes.QueryPinnedCodesResponse
)

// NewMsgStoreCode upload code to be reused.
func NewMsgStoreCode(sender string, wasmFile string, instantiateByAddress string) (*MsgStoreCode, error) {
	wasm, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		return nil, err
	}
	if wasmUtils.IsWasm(wasm) {
		wasm, err = wasmUtils.GzipIt(wasm)

		if err != nil {
			return nil, err
		}
	} else if !wasmUtils.IsGzip(wasm) {
		return nil, fmt.Errorf("invalid input file. Use wasm binary or gzip")
	}

	var perm *wasmtypes.AccessConfig
	if len(instantiateByAddress) == 0 {
		perm = &wasmtypes.AllowEverybody
	} else {
		allowedAddr, err := AccAddressFromBech32(instantiateByAddress)
		if err != nil {
			return nil, fmt.Errorf("address %s: %s", instantiateByAddress, err)
		}
		x := wasmtypes.AccessTypeOnlyAddress.With(allowedAddr)
		perm = &x
	}

	msg := &wasmtypes.MsgStoreCode{
		Sender:                sender,
		WASMByteCode:          wasm,
		InstantiatePermission: perm,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

// NewMsgInstantiateContract instantiate a contract from previously uploaded code.
func NewMsgInstantiateContract(sender string, codeID uint64, initArgs string, label string, admin string, amountStr string) (*MsgInstantiateContract, error) {
	amount, err := ParseCoinsNormalized(amountStr)
	if err != nil {
		return nil, fmt.Errorf("amount %s: %s", amountStr, err)
	}
	msg := &MsgInstantiateContract{
		Sender: sender,
		CodeID: codeID,
		Label:  label,
		Funds:  amount,
		Msg:    []byte(initArgs),
		Admin:  admin,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

// NewMsgExecuteContract execute a command on a wasm contract
func NewMsgExecuteContract(sender string, contractAddr string, execArgs string, amountStr string) (*MsgExecuteContract, error) {
	amount, err := ParseCoinsNormalized(amountStr)
	if err != nil {
		return nil, fmt.Errorf("amount %s: %s", amountStr, err)
	}

	msg := &MsgExecuteContract{
		Sender:   sender,
		Contract: contractAddr,
		Funds:    amount,
		Msg:      []byte(execArgs),
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

//NewMsgMigrateContract migrate a wasm contract to a new code version
func NewMsgMigrateContract(sender string, contractAddr string, codeID uint64, migrateArgs string) (*MsgMigrateContract, error) {
	msg := &MsgMigrateContract{
		Sender:   sender,
		Contract: contractAddr,
		CodeID:   codeID,
		Msg:      []byte(migrateArgs),
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

// NewMsgUpdateAdmin set new admin for a contract
func NewMsgUpdateAdmin(sender string, contractAddr string, admin string) (*MsgUpdateAdmin, error) {
	msg := &MsgUpdateAdmin{
		Sender:   sender,
		Contract: contractAddr,
		NewAdmin: admin,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}

// NewMsgClearAdmin clears admin for a contract to prevent further migrations
func NewMsgClearAdmin(sender string, contractAddr string) (*MsgClearAdmin, error) {
	msg := &MsgClearAdmin{
		Sender:   sender,
		Contract: contractAddr,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}
