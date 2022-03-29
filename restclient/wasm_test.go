package restclient

import (
	"fmt"
	"testing"

	"github.com/glodnet/chain.go/types"
	"github.com/stretchr/testify/assert"
)

var (
	testContract string = "gnc1dt3lk455ed360pna38fkhqn0p8y44qndsr77qu73ghyaz2zv4whqpdpj68"
)

func TestMsgStoreCode(t *testing.T) {
	addr := types.AccAddress(privKey.PubKey().Address())
	msg, err := types.NewMsgStoreCode(addr.String(), "../cw_nameservice.wasm", "")
	assert.NoError(t, err)
	hash, err := rest.TxSend(privKey, &types.BuildTxOptions{
		Sender: addr,
		Msgs: []types.Msg{
			msg,
		},
		Memo: "",
	}, types.BROADCAST_MODE_BLOCK)
	assert.NoError(t, err)
	fmt.Println("store", hash)
}

func TestMsgInstantiateContract(t *testing.T) {
	addr := types.AccAddress(privKey.PubKey().Address())
	initArgs := "{\"purchase_price\":{\"amount\":\"100\",\"denom\":\"ugnc\"},\"transfer_price\":{\"amount\":\"999\",\"denom\":\"ugnc\"}}"
	msg, err := types.NewMsgInstantiateContract(addr.String(), 1, initArgs, "test", addr.String(), "")
	assert.NoError(t, err)
	hash, err := rest.TxSend(privKey, &types.BuildTxOptions{
		Sender: addr,
		Msgs: []types.Msg{
			msg,
		},
		Memo: "",
	}, types.BROADCAST_MODE_BLOCK)
	assert.NoError(t, err)
	fmt.Println("instantiate", hash)
}

func TestContractsByCode2(t *testing.T) {
	var key []byte
	for {
		res, err := rest.ContractsByCode(1, key)
		assert.NoError(t, err)
		if len(res.Contracts) > 0 {
			testContract = res.Contracts[len(res.Contracts)-1]
		}
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}

func TestMsgExecuteContract(t *testing.T) {
	addr := types.AccAddress(privKey.PubKey().Address())
	execArgs := "{\"register\":{\"name\":\"fred\"}}"
	msg, err := types.NewMsgExecuteContract(addr.String(), testContract, execArgs, "100ugnc")
	assert.NoError(t, err)
	hash, err := rest.TxSend(privKey, &types.BuildTxOptions{
		Sender: addr,
		Msgs: []types.Msg{
			msg,
		},
		Memo: "",
	}, types.BROADCAST_MODE_BLOCK)
	assert.NoError(t, err)
	fmt.Println("execute", hash)
}

func TestMsgExecuteContract2(t *testing.T) {
	addr := types.AccAddress(privKey.PubKey().Address())
	execArgs := "{\"transfer\":{\"name\":\"fred\",\"to\":\"gnc1azlj5whn5rm2xtqeekkdqgwg7036naf0sfqwmu\"}}"
	msg, err := types.NewMsgExecuteContract(addr.String(), testContract, execArgs, "999ugnc")
	assert.NoError(t, err)
	hash, err := rest.TxSend(privKey, &types.BuildTxOptions{
		Sender: addr,
		Msgs: []types.Msg{
			msg,
		},
		Memo: "",
	}, types.BROADCAST_MODE_BLOCK)
	assert.NoError(t, err)
	fmt.Println("execute", hash)
}

//func TestMsgMigrateContract(t *testing.T) {
//	addr := types.AccAddress(privKey.PubKey().Address())
//	msg, err := types.NewMsgMigrateContract(addr.String(), testContract, 1, "{}", "")
//	assert.NoError(t, err)
//	hash, err := rest.TxSend(privKey, &types.BuildTxOptions{
//		Sender: addr,
//		Msgs: []types.Msg{
//			msg,
//		},
//		Memo: "",
//	}, types.BROADCAST_MODE_BLOCK)
//	assert.NoError(t, err)
//	fmt.Println("instantiate", hash)
//}

func TestCodes(t *testing.T) {
	var key []byte
	for {
		res, err := rest.Codes(key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}

func TestCode(t *testing.T) {
	res, err := rest.Code(1)
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

func TestContractsByCode(t *testing.T) {
	var key []byte
	for {
		res, err := rest.ContractsByCode(1, key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}

func TestContractInfo(t *testing.T) {
	res, err := rest.ContractInfo(testContract)
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

func TestContractHistory(t *testing.T) {
	var key []byte
	for {
		res, err := rest.ContractHistory(testContract, key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}

func TestAllContractState(t *testing.T) {
	var key []byte
	for {
		res, err := rest.AllContractState(testContract, key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}

func TestRawContractState(t *testing.T) {
	res, err := rest.RawContractStateHex(testContract, "000C6E616D657265736F6C76657266726564")
	assert.NoError(t, err)
	// fmt.Println(rest.MarshalJSON(res))
	fmt.Println(string(res.Data))
}

func TestRawContractState2(t *testing.T) {
	res, err := rest.RawContractState(testContract, "nameresolver", "fred")
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
	fmt.Println(string(res.Data))
}

func TestSmartContractState(t *testing.T) {
	res, err := rest.SmartContractState(testContract, "{\"resolve_record\": {\"name\": \"fred\"}}")
	assert.NoError(t, err)
	// fmt.Println(rest.MarshalJSON(res))
	fmt.Println(string(res.Data))
}
