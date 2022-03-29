package types

import (
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

type (
	SimulateRequest  = tx.SimulateRequest
	SimulateResponse = tx.SimulateResponse

	GetTxRequest  = tx.GetTxRequest
	GetTxResponse = tx.GetTxResponse

	BroadcastMode = tx.BroadcastMode

	BroadcastTxRequest  = tx.BroadcastTxRequest
	BroadcastTxResponse = tx.BroadcastTxResponse

	GetTxsEventRequest  = tx.GetTxsEventRequest
	GetTxsEventResponse = tx.GetTxsEventResponse

	SignatureV2 = signing.SignatureV2
)

var (
	BROADCAST_MODE_BLOCK = tx.BroadcastMode_BROADCAST_MODE_BLOCK
	BROADCAST_MODE_SYNC  = tx.BroadcastMode_BROADCAST_MODE_SYNC
	BROADCAST_MODE_ASYNC = tx.BroadcastMode_BROADCAST_MODE_ASYNC
)
