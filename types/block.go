package types

import "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"

type (
	GetNodeInfoRequest              = tmservice.GetNodeInfoRequest
	GetNodeInfoResponse             = tmservice.GetNodeInfoResponse
	GetSyncingRequest               = tmservice.GetSyncingRequest
	GetSyncingResponse              = tmservice.GetSyncingResponse
	GetLatestBlockRequest           = tmservice.GetLatestBlockRequest
	GetLatestBlockResponse          = tmservice.GetLatestBlockResponse
	GetBlockByHeightRequest         = tmservice.GetBlockByHeightRequest
	GetBlockByHeightResponse        = tmservice.GetBlockByHeightResponse
	GetLatestValidatorSetRequest    = tmservice.GetLatestValidatorSetRequest
	GetLatestValidatorSetResponse   = tmservice.GetLatestValidatorSetResponse
	GetValidatorSetByHeightRequest  = tmservice.GetValidatorSetByHeightRequest
	GetValidatorSetByHeightResponse = tmservice.GetValidatorSetByHeightResponse
)
