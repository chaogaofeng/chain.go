package restclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeInfo(t *testing.T) {
	res, err := rest.NodeInfo()
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

func TestSyncing(t *testing.T) {
	res, err := rest.Syncing()
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

func TestBlockLatest(t *testing.T) {
	res, err := rest.BlockLatest()
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

func TestBlockByHeight(t *testing.T) {
	res, err := rest.BlockByHeight(1)
	assert.NoError(t, err)
	fmt.Println(rest.MarshalJSON(res))
}

func TestValidatorSetByHeight(t *testing.T) {
	var key []byte
	for {
		res, err := rest.ValidatorSetByHeight(1, key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}

func TestValidatorSetLatest(t *testing.T) {
	var key []byte
	for {
		res, err := rest.ValidatorSetLatest(key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}
