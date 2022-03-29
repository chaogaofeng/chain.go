package types

import (
	"fmt"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type (
	MsgSend      = banktypes.MsgSend
	MsgMultiSend = banktypes.MsgMultiSend

	QueryBalanceRequest  = banktypes.QueryBalanceRequest
	QueryBalanceResponse = banktypes.QueryBalanceResponse

	QueryAllBalancesRequest  = banktypes.QueryAllBalancesRequest
	QueryAllBalancesResponse = banktypes.QueryAllBalancesResponse

	QueryTotalSupplyRequest  = banktypes.QueryTotalSupplyRequest
	QueryTotalSupplyResponse = banktypes.QueryTotalSupplyResponse

	QuerySupplyOfRequest  = banktypes.QuerySupplyOfRequest
	QuerySupplyOfResponse = banktypes.QuerySupplyOfResponse

	QueryDenomsMetadataRequest  = banktypes.QueryDenomsMetadataRequest
	QueryDenomsMetadataResponse = banktypes.QueryDenomsMetadataResponse

	QueryDenomMetadataRequest  = banktypes.QueryDenomMetadataRequest
	QueryDenomMetadataResponse = banktypes.QueryDenomMetadataResponse

	QueryBankParamsRequest  = banktypes.QueryParamsRequest
	QueryBankParamsResponse = banktypes.QueryParamsResponse
)

var (
	NewMsgMultiSend = banktypes.NewMsgMultiSend
)

func NewMsgSend(from string, to string, amountStr string) (*MsgSend, error) {
	amount, err := ParseCoinsNormalized(amountStr)
	if err != nil {
		return nil, fmt.Errorf("amount %s: %s", amountStr, err)
	}
	msg := &MsgSend{
		FromAddress: from,
		ToAddress:   to,
		Amount:      amount,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	return msg, nil
}
