package restclient

import (
	"encoding/base64"
	"github.com/glodnet/chain.go/types"
	"strconv"
	"strings"
)

// NodeInfo queries the current node info.
func (client *RestClient) NodeInfo() (*types.GetNodeInfoResponse, error) {
	var response types.GetNodeInfoResponse
	if err := client.get("/cosmos/base/tendermint/v1beta1/node_info", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Syncing queries node syncing.
func (client *RestClient) Syncing() (*types.GetSyncingResponse, error) {
	var response types.GetSyncingResponse
	if err := client.get("/cosmos/base/tendermint/v1beta1/syncing", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BlockLatest returns the latest block.
func (client *RestClient) BlockLatest() (*types.GetLatestBlockResponse, error) {
	var response types.GetLatestBlockResponse
	if err := client.get("/cosmos/base/tendermint/v1beta1/blocks/latest", &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// BlockByHeight queries block for given height.
func (client *RestClient) BlockByHeight(height int64) (*types.GetBlockByHeightResponse, error) {
	var response types.GetBlockByHeightResponse
	if err := client.get("/cosmos/base/tendermint/v1beta1/blocks/"+strconv.FormatInt(height, 10), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ValidatorSetLatest queries latest validator-set.
func (client *RestClient) ValidatorSetLatest(key []byte) (*types.GetLatestValidatorSetResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.GetLatestValidatorSetResponse
	if err := client.get("/cosmos/base/tendermint/v1beta1/validatorsets/latest"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ValidatorSetByHeight queries validator-set at a given height.
func (client *RestClient) ValidatorSetByHeight(height int64, key []byte) (*types.GetValidatorSetByHeightResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.GetValidatorSetByHeightResponse
	if err := client.get("/cosmos/base/tendermint/v1beta1/validatorsets/"+strconv.FormatInt(height, 10)+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
