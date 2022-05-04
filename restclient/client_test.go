package restclient

import (
	"fmt"
	"github.com/glodnet/chain/app"
	"testing"

	"github.com/glodnet/chain.go/types"

	"github.com/glodnet/chain.go/key"
	"github.com/stretchr/testify/assert"
)

var (
	rest = New(
		"http://127.0.0.1:1317",
		"gnchain",
		types.NewDecCoinFromDec("ugnc", types.NewDecFromIntWithPrec(types.NewInt(15), 2)), // 0.15ugnc
		types.NewDecFromIntWithPrec(types.NewInt(15), 1),                                  // 1.5
		app.AccountAddressPrefix,
	)

	mnemonic     = "apology false junior asset sphere puppy upset dirt miracle rice horn spell ring vast wrist crisp snake oak give cement pause swallow barely clever"
	privKeyBz, _ = key.DerivePrivKeyBz(mnemonic, types.FullFundraiserPath)
	privKey, _   = key.PrivKeyGen(privKeyBz)
)

func Test_SendTransaction(t *testing.T) {
	mnemonic := "apology false junior asset sphere puppy upset dirt miracle rice horn spell ring vast wrist crisp snake oak give cement pause swallow barely clever"
	privKeyBz, err := key.DerivePrivKeyBz(mnemonic, types.FullFundraiserPath)
	assert.NoError(t, err)
	privKey, err := key.PrivKeyGen(privKeyBz)
	assert.NoError(t, err)

	addr := types.AccAddress(privKey.PubKey().Address())
	assert.Equal(t, addr.String(), "gnc1azlj5whn5rm2xtqeekkdqgwg7036naf0sfqwmu")

	toAddr, err := types.AccAddressFromBech32("gnc1azlj5whn5rm2xtqeekkdqgwg7036naf0sfqwmu")
	assert.NoError(t, err)

	msg, err := types.NewMsgSend(addr.String(), toAddr.String(), "100000000ugnc") // 100gnc
	assert.NoError(t, err)

	txOptions := &types.BuildTxOptions{
		Sender: addr,
		Msgs: []types.Msg{
			msg,
		},
		Memo: "",
	}

	txHash, err := rest.TxSend(privKey, txOptions, types.BROADCAST_MODE_BLOCK)
	assert.NoError(t, err)

	fmt.Println(txHash)
}
