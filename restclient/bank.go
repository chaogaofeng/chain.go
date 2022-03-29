package restclient

import (
	"encoding/base64"
	"github.com/glodnet/chain.go/types"
	"strings"
)

// Balance queries the balance of a single coin for a single account.
func (client *RestClient) Balance(address string, denom string) (*types.QueryBalanceResponse, error) {
	var response types.QueryBalanceResponse
	if err := client.get("/cosmos/bank/v1beta1/balances/"+address+"/by_denom?denom="+denom, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// AllBalances queries the balance of all coins for a single account.
func (client *RestClient) Balances(address string) (*types.QueryAllBalancesResponse, error) {
	var response types.QueryAllBalancesResponse
	if err := client.get("/cosmos/bank/v1beta1/balances/"+address, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// TotalSupply queries the total supply of all coins.
func (client *RestClient) TotalSupply(key []byte) (*types.QueryTotalSupplyResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.QueryTotalSupplyResponse
	if err := client.get("/cosmos/bank/v1beta1/supply"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// SupplyOf queries the supply of a single coin.
func (client *RestClient) SupplyOf(denom string) (*types.QuerySupplyOfResponse, error) {
	var response types.QuerySupplyOfResponse
	if err := client.get("/cosmos/bank/v1beta1/supply/"+denom, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DenomMetadata queries the client metadata of a given coin denomination.
func (client *RestClient) DenomMetadata(denom string) (*types.QueryDenomMetadataResponse, error) {
	var response types.QueryDenomMetadataResponse
	if err := client.get("/cosmos/bank/v1beta1/denoms_metadata/"+denom, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DenomsMetadata queries the client metadata for all registered coin denominations.
func (client *RestClient) DenomsMetadata(key []byte) (*types.QueryDenomsMetadataResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}

	var response types.QueryDenomsMetadataResponse
	if err := client.get("/cosmos/bank/v1beta1/denoms_metadata"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
