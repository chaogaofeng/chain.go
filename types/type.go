package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	FullFundraiserPath = "m/44'/118'/0'/0/0"
)

type (
	// Coin nolint
	Coin = sdk.Coin
	// Coins nolint
	Coins = sdk.Coins
	// DecCoin nolint
	DecCoin = sdk.DecCoin
	// DecCoins nolint
	DecCoins = sdk.DecCoins

	// Int nolint
	Int = sdk.Int
	// Dec nolint
	Dec = sdk.Dec

	// AccAddress nolint
	AccAddress = sdk.AccAddress
	// ValAddress nolint
	ValAddress = sdk.ValAddress
	// ConsAddress nolint
	ConsAddress = sdk.ConsAddress

	Msg = sdk.Msg

	TxResponse = sdk.TxResponse
)

var (
	NewCoin         = sdk.NewCoin
	NewInt64Coin    = sdk.NewInt64Coin
	NewCoins        = sdk.NewCoins
	NewDecCoin      = sdk.NewDecCoin
	NewInt64DecCoin = sdk.NewInt64DecCoin
	NewDecCoins     = sdk.NewDecCoins

	ParseCoinNormalized  = sdk.ParseCoinNormalized
	ParseCoinsNormalized = sdk.ParseCoinsNormalized
	ParseDecCoin         = sdk.ParseDecCoin
	ParseDecCoins        = sdk.ParseDecCoins

	NewInt                   = sdk.NewInt
	NewIntFromBigInt         = sdk.NewIntFromBigInt
	NewIntFromString         = sdk.NewIntFromString
	NewIntFromUint64         = sdk.NewIntFromUint64
	NewIntWithDecimal        = sdk.NewIntWithDecimal
	NewDec                   = sdk.NewDec
	NewDecCoinFromCoin       = sdk.NewDecCoinFromCoin
	NewDecCoinFromDec        = sdk.NewDecCoinFromDec
	NewDecFromBigInt         = sdk.NewDecFromBigInt
	NewDecFromBigIntWithPrec = sdk.NewDecFromBigIntWithPrec
	NewDecFromInt            = sdk.NewDecFromInt
	NewDecFromIntWithPrec    = sdk.NewDecFromIntWithPrec
	NewDecFromStr            = sdk.NewDecFromStr
	NewDecWithPrec           = sdk.NewDecWithPrec
	AccAddressFromBech32     = sdk.AccAddressFromBech32
	AccAddressFromHex        = sdk.AccAddressFromHex
	ValAddressFromBech32     = sdk.ValAddressFromBech32
	ValAddressFromHex        = sdk.ValAddressFromHex
	ConsAddressFromBech32    = sdk.ConsAddressFromBech32
	ConsAddressFromHex       = sdk.ConsAddressFromHex
)
