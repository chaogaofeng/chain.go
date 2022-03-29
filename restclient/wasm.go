package restclient

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/glodnet/chain.go/types"
	"strconv"
	"strings"
)

// ContractInfo gets the contract meta data
// address is the address of the contract to query
func (client *RestClient) ContractInfo(address string) (*types.QueryContractInfoResponse, error) {
	var response types.QueryContractInfoResponse
	if err := client.get("/cosmwasm/wasm/v1/contract/"+address, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ContractHistory gets the contract code history
// address is the address of the contract to query
func (client *RestClient) ContractHistory(address string, key []byte) (*types.QueryContractHistoryResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.QueryContractHistoryResponse
	if err := client.get("/cosmwasm/wasm/v1/contract/"+address+"/history"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ContractsByCode lists all smart contracts for a code id
func (client *RestClient) ContractsByCode(codeID int64, key []byte) (*types.QueryContractsByCodeResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.QueryContractsByCodeResponse
	if err := client.get("/cosmwasm/wasm/v1/code/"+strconv.FormatInt(codeID, 10)+"/contracts"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// AllContractState gets all raw store data for a single contract
func (client *RestClient) AllContractState(address string, key []byte) (*types.QueryAllContractStateResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.QueryAllContractStateResponse
	if err := client.get("/cosmwasm/wasm/v1/contract/"+address+"/state"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// RawContractStateHex gets single key from the raw store data of a contract
func (client *RestClient) RawContractStateHex(address string, queryData string) (*types.QueryRawContractStateResponse, error) {
	var response types.QueryRawContractStateResponse
	bts, err := hex.DecodeString(queryData)
	if err != nil {
		return nil, err
	}

	if err := client.get("/cosmwasm/wasm/v1/contract/"+address+"/raw/"+base64.StdEncoding.EncodeToString(bts), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (client *RestClient) RawContractState(address string, key string, others ...string) (*types.QueryRawContractStateResponse, error) {
	queryData := fmt.Sprintf("%04x", len(key)) + hex.EncodeToString([]byte(key))
	for _, other := range others {
		queryData += hex.EncodeToString([]byte(other))
	}

	return client.RawContractStateHex(address, queryData)
}

// SmartContractState get smart query result from the contract
func (client *RestClient) SmartContractState(address string, queryData string) (*types.QuerySmartContractStateResponse, error) {
	if !json.Valid([]byte(queryData)) {
		return nil, fmt.Errorf("query data must be json")
	}
	var response types.QuerySmartContractStateResponse
	if err := client.get("/cosmwasm/wasm/v1/contract/"+address+"/smart/"+base64.StdEncoding.EncodeToString([]byte(queryData)), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Code gets the binary code and metadata for a singe wasm code
func (client *RestClient) Code(codeID int64) (*types.QueryCodeResponse, error) {
	var response types.QueryCodeResponse
	if err := client.get("/cosmwasm/wasm/v1/code/"+strconv.FormatInt(codeID, 10), &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Codes gets the metadata for all stored wasm codes
func (client *RestClient) Codes(key []byte) (*types.QueryCodesResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.QueryCodesResponse
	if err := client.get("/cosmwasm/wasm/v1/code"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// PinnedCodes gets the pinned code ids
func (client *RestClient) PinnedCodes(key []byte) (*types.QueryPinnedCodesResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.QueryPinnedCodesResponse
	if err := client.get("/cosmwasm/wasm/v1/codes/pinned"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
