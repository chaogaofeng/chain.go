package restclient

import (
	"fmt"
	"testing"
	"time"

	"github.com/glodnet/chain.go/types"
	"github.com/stretchr/testify/assert"
)

func TestTx(t *testing.T) {
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

	time.Sleep(time.Second)

	res, err := rest.Tx(hash)
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

func TestTxsByEvent(t *testing.T) {
	var key []byte
	for {
		res, err := rest.TxsByEvent([]string{"message.action='/cosmos.bank.v1beta1.MsgSend'"}, key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}

func TestTxsByHeight(t *testing.T) {
	var key []byte
	for {
		res, err := rest.TxsByHeight(327, key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}
