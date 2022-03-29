package types

type BuildTxOptions struct {
	Msgs []Msg
	Memo string

	// Optional parameters
	Sender        AccAddress
	AccountNumber uint64
	Sequence      uint64
	GasLimit      uint64
	FeeAmount     Coins

	FeeGranter    AccAddress
	TimeoutHeight uint64
}
