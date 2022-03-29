package restclient

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountGet(t *testing.T) {
	res, err := rest.AccountGet("gnc1azlj5whn5rm2xtqeekkdqgwg7036naf0sfqwmu")
	assert.NoError(t, err)

	fmt.Println(rest.MarshalJSON(res))
}

func TestBaseAccountGet(t *testing.T) {
	res, err := rest.BaseAccountGet("gnc1azlj5whn5rm2xtqeekkdqgwg7036naf0sfqwmu")
	assert.NoError(t, err)

	fmt.Println(rest.MarshalJSON(res))
}

func TestAccountsGet(t *testing.T) {
	var key []byte
	for {
		res, err := rest.AccountsGet(key)
		assert.NoError(t, err)
		fmt.Println(rest.MarshalJSON(res))
		key = res.Pagination.NextKey
		if len(key) == 0 {
			break
		}
	}
}
