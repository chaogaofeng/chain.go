package types

import (
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type (
	AccountI      = authtypes.AccountI
	BaseAccount   = authtypes.BaseAccount
	ModuleAccount = authtypes.ModuleAccount

	QueryAccountRequest  = authtypes.QueryAccountRequest
	QueryAccountResponse = authtypes.QueryAccountResponse

	QueryAccountsRequest  = authtypes.QueryAccountsRequest
	QueryAccountsResponse = authtypes.QueryAccountsResponse

	QueryAuthParamsRequest  = authtypes.QueryParamsRequest
	QueryAuthParamsResponse = authtypes.QueryParamsResponse

	SignerData = authsigning.SignerData
)
