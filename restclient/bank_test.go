package restclient

import (
	"fmt"
	"testing"

	"github.com/glodnet/chain.go/types"
	"github.com/stretchr/testify/assert"
)

func (client *RestClient) TestMsgSend(t *testing.T) {
	addr := types.AccAddress(privKey.PubKey().Address())
	msg, err := types.NewMsgSend(addr.String(), addr.String(), "100000000ugnc") // 100gnc
	assert.NoError(t, err)

	hash, err := rest.TxSend(privKey, &types.BuildTxOptions{
		Sender: addr,
		Msgs: []types.Msg{
			msg,
		},
		Memo: "",
	}, types.BROADCAST_MODE_BLOCK)
	assert.NoError(t, err)
	fmt.Println(hash)
}

func TestBalance(t *testing.T) {
	res, err := rest.Balance("gnc1azlj5whn5rm2xtqeekkdqgwg7036naf0sfqwmu", "ugnc")
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

func TestBalances(t *testing.T) {
	res, err := rest.Balances("gnc1azlj5whn5rm2xtqeekkdqgwg7036naf0sfqwmu")
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

func TestTotalSupply(t *testing.T) {
	var key []byte
	for {
		res, err := rest.TotalSupply(key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}

}

func TestSupplyOf(t *testing.T) {
	res, err := rest.SupplyOf("ugnc")
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

//func TestDenomMetadata(t *testing.T) {
//	res, err := rest.DenomMetadata("ugnc")
//	assert.NoError(t, err)
//	fmt.Println(rest.MarshalJSON(res))
//}

func TestDenomsMetadata(t *testing.T) {
	var key []byte
	for {
		res, err := rest.DenomsMetadata(key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}
