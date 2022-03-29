package restclient

import (
	"encoding/base64"
	"fmt"
	"github.com/glodnet/chain.go/types"
	"strconv"
	"strings"
)

// TxSimulate simulates executing a transaction for estimating gas usage.
func (client *RestClient) TxSimulate(txBytes []byte) (*types.SimulateResponse, error) {
	var response types.SimulateResponse
	if err := client.post("/cosmos/tx/v1beta1/simulate", &types.SimulateRequest{
		TxBytes: txBytes,
	}, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Tx fetches a tx by hash.
func (client *RestClient) Tx(hash string) (*types.GetTxResponse, error) {
	var response types.GetTxResponse
	if err := client.get("/cosmos/tx/v1beta1/txs/"+hash, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// TxBroadcast broadcast transaction.
func (client *RestClient) TxBroadcast(txBytes []byte, mode types.BroadcastMode) (*types.BroadcastTxResponse, error) {
	var response types.BroadcastTxResponse
	if err := client.post("/cosmos/tx/v1beta1/txs", &types.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    mode,
	}, &response); err != nil {
		return nil, err
	}

	txResponse := response.TxResponse
	if txResponse.Code != 0 {
		return nil, fmt.Errorf("tx failed with code %d: %s", txResponse.Code, txResponse.RawLog)
	}
	return &response, nil
}

// TxsGet fetches txs by event.
func (client *RestClient) TxsByEvent(events []string, key []byte) (*types.GetTxsEventResponse, error) {
	var params []string
	for _, event := range events {
		params = append(params, "events="+event)
	}
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.GetTxsEventResponse
	if err := client.get("/cosmos/tx/v1beta1/txs"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// TxsByHeight fetches txs by height.
func (client *RestClient) TxsByHeight(height int64, key []byte) (*types.GetTxsEventResponse, error) {
	return client.TxsByEvent([]string{"tx.height=" + strconv.FormatInt(height, 10)}, key)
}
