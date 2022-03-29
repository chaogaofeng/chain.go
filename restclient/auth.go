package restclient

import (
	"encoding/base64"
	"github.com/glodnet/chain.go/types"
	"strings"
)

// BaseAccountGet returns account details based on address.
func (client *RestClient) BaseAccountGet(address string) (*types.BaseAccount, error) {
	var response types.QueryAccountResponse
	if err := client.get("/cosmos/auth/v1beta1/accounts/"+address, &response); err != nil {
		return nil, err
	}
	return response.Account.GetCachedValue().(*types.BaseAccount), nil
}

// AccountGet returns account details based on address.
func (client *RestClient) AccountGet(address string) (*types.QueryAccountResponse, error) {
	var response types.QueryAccountResponse
	if err := client.get("/cosmos/auth/v1beta1/accounts/"+address, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// AccountsGet returns all the existing accounts
func (client *RestClient) AccountsGet(key []byte) (*types.QueryAccountsResponse, error) {
	var params []string
	if len(key) > 0 {
		params = append(params, "pagination.key="+base64.StdEncoding.EncodeToString(key))
	}
	query := ""
	if len(params) > 0 {
		query = "?" + strings.Join(params, "&")
	}
	var response types.QueryAccountsResponse
	if err := client.get("/cosmos/auth/v1beta1/accounts"+query, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
