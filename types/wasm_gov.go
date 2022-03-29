package types

import (
	"fmt"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

type (
	StoreCodeProposal           = wasmtypes.StoreCodeProposal
	InstantiateContractProposal = wasmtypes.InstantiateContractProposal
	MigrateContractProposal     = wasmtypes.MigrateContractProposal
	ExecuteContractProposal     = wasmtypes.ExecuteContractProposal
	UpdateAdminProposal         = wasmtypes.UpdateAdminProposal
	ClearAdminProposal          = wasmtypes.ClearAdminProposal
	SudoContractProposal        = wasmtypes.SudoContractProposal
	PinCodesProposal            = wasmtypes.PinCodesProposal
	UnpinCodesProposal          = wasmtypes.UnpinCodesProposal
)

func NewStoreCodeProposal(sender string, wasmFile string, instantiateByAddress string) (*StoreCodeProposal, error) {
	return nil, fmt.Errorf("unimpl")
}

func NewInstantiateContractProposal(sender string, wasmFile string, instantiateByAddress string) (*InstantiateContractProposal, error) {
	return nil, fmt.Errorf("unimpl")
}

func NewMigrateContractProposal(sender string, wasmFile string, instantiateByAddress string) (*MigrateContractProposal, error) {
	return nil, fmt.Errorf("unimpl")
}

func NewExecuteContractProposal(sender string, wasmFile string, instantiateByAddress string) (*ExecuteContractProposal, error) {
	return nil, fmt.Errorf("unimpl")
}

func NewUpdateAdminProposal(sender string, wasmFile string, instantiateByAddress string) (*UpdateAdminProposal, error) {
	return nil, fmt.Errorf("unimpl")
}

func NewClearAdminProposal(sender string, wasmFile string, instantiateByAddress string) (*ClearAdminProposal, error) {
	return nil, fmt.Errorf("unimpl")
}

func NewSudoContractProposal(sender string, wasmFile string, instantiateByAddress string) (*SudoContractProposal, error) {
	return nil, fmt.Errorf("unimpl")
}

func NewPinCodesProposal(sender string, wasmFile string, instantiateByAddress string) (*PinCodesProposal, error) {
	return nil, fmt.Errorf("unimpl")
}

func NewUnpinCodesProposal(sender string, wasmFile string, instantiateByAddress string) (*UnpinCodesProposal, error) {
	return nil, fmt.Errorf("unimpl")
}
