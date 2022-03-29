package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	ModuleName = "RESTAPI"
)

var (
	ErrNotAccess      = sdkerrors.Register(ModuleName, 2, "failed to access")
	ErrResponseStatus = sdkerrors.Register(ModuleName, 3, "non-200 response status code")
	ErrResponseBody   = sdkerrors.Register(ModuleName, 4, "invalid response body")
	ErrMarshaler      = sdkerrors.Register(ModuleName, 5, "failed to marshal/unmarshal")
)
