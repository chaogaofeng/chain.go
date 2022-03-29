package restclient

import (
	"bytes"
	clitx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/glodnet/chain.go/types"
	"github.com/glodnet/chain/app"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/starport/starport/pkg/cosmoscmd"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	HTTPTimeout = 10 // 10s
)

func init() {
	cosmoscmd.SetPrefixes(app.AccountAddressPrefix)
}

type RestClient struct {
	baseUrl       string
	chainID       string
	gasPrice      types.DecCoin
	gasAdjustment types.Dec

	c              *http.Client
	encodingConfig cosmoscmd.EncodingConfig
}

/**
 * 创建SDK客户端
 *
 * @param baseUrl        节点 REST API。 例： http://127.0.0.1:1317
 * @param chainID        网络ChainID。 例： gnchain
 * @param gasPrice       gas手续费。自动计算feeAmount时使用。 例： 0.00002ugnc
 * @param gasAdjustment  gas调整倍数。自动计算gasLimit时使用。例： 1.1
 * @return SDK客户端
 */
func New(baseUrl string, chainID string, gasPrice types.DecCoin, gasAdjustment types.Dec) *RestClient {
	return &RestClient{
		baseUrl:        baseUrl,
		chainID:        chainID,
		gasPrice:       gasPrice,
		gasAdjustment:  gasAdjustment,
		c:              &http.Client{Timeout: time.Duration(HTTPTimeout) * time.Second},
		encodingConfig: cosmoscmd.MakeEncodingConfig(app.ModuleBasics),
	}
}

/**
 * TxSend 发送交易
 *
 * @param priv            私钥。
 * @param msgs            消息事务列表。例：转账列表。
 * @param memo            备注内容。 不可超过最大允许字节数。
 * @param sender          发送者。 accountNumber、sequence自动填充使用。
 * @param accountNumber   账户个数。 0时自动填充， sender's accountNumber。
 * @param sequence        账户发送交易个数。0时自动填充 sender's sequence。
 * @param gasLimit        gas最大可用量（gas用完时，矿工会退出执行，且扣除手续费）。0时自动计算填充，gas * gasAdjustment。
 * @param feeAmount       手续费总额。0时自动计算填充，gasLimit * gasPrice。
 * @param feeGranter	  手续费扣除地址
 * @param timeoutHeight   交易超时高度
 * @return 交易hash
 */
func (client *RestClient) TxSend(priv types.PrivKey, options *types.BuildTxOptions, mode types.BroadcastMode) (string, error) {
	if options.Sender.Empty() {
		options.Sender = types.AccAddress(priv.PubKey().Address())
	}
	txBytes, err := client.TxBuild(options)
	if err != nil {
		return "", err
	}
	txBytes, err = client.TxSign(txBytes, priv, options.AccountNumber, options.Sequence, true)
	if err != nil {
		return "", err
	}
	res, err := client.TxBroadcast(txBytes, mode)
	if err != nil {
		return "", err
	}
	return res.TxResponse.TxHash, nil
}

/**
 * TxBuild 创建交易
 *
 * @param msgs            消息事务列表。例：转账列表。
 * @param memo            备注内容。 不可超过最大允许字节数。
 * @param sender          发送者。 accountNumber、sequence自动填充使用。
 * @param accountNumber   账户个数。 0时自动填充， sender's accountNumber。
 * @param sequence        账户发送交易个数。0时自动填充 sender's sequence。
 * @param gasLimit        gas最大可用量（gas用完时，矿工会退出执行，且扣除手续费）。0时自动计算填充，gas * gasAdjustment。
 * @param feeAmount       手续费总额。0时自动计算填充，gasLimit * gasPrice。
 * @param feeGranter	  手续费扣除地址
 * @param timeoutHeight   交易超时高度
 * @return 交易序列化字节
 */
func (client *RestClient) TxBuild(options *types.BuildTxOptions) ([]byte, error) {
	txBuilder := client.encodingConfig.TxConfig.NewTxBuilder()
	txBuilder.SetMsgs(options.Msgs...)
	txBuilder.SetMemo(options.Memo)
	txBuilder.SetTimeoutHeight(options.TimeoutHeight)
	txBuilder.SetFeeGranter(options.FeeGranter)
	txBuilder.SetGasLimit(options.GasLimit)
	txBuilder.SetFeeAmount(options.FeeAmount)

	if options.AccountNumber == 0 || options.Sequence == 0 {
		resp, err := client.BaseAccountGet(options.Sender.String())
		if err != nil {
			return nil, err
		}
		options.AccountNumber = resp.GetAccountNumber()
		options.Sequence = resp.GetSequence()
	}

	gasLimit := int64(options.GasLimit)
	if options.GasLimit == 0 {
		// Create an empty signature literal as the ante handler will populate with a
		// sentinel pubkey.
		sig := signing.SignatureV2{
			PubKey: &secp256k1.PubKey{},
			Data: &signing.SingleSignatureData{
				SignMode: signing.SignMode_SIGN_MODE_DIRECT,
			},
			Sequence: options.Sequence,
		}
		if err := txBuilder.SetSignatures(sig); err != nil {
			return nil, err
		}

		bz, err := client.encodingConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
		if err != nil {
			return nil, err
		}
		resp, err := client.TxSimulate(bz)
		if err != nil {
			return nil, err
		}
		gasLimit = client.gasAdjustment.MulInt64(int64(resp.GasInfo.GasUsed)).TruncateInt64()
		options.GasLimit = uint64(gasLimit)
	}

	if options.FeeAmount.IsZero() {
		gasFee := types.NewCoin(client.gasPrice.Denom, client.gasPrice.Amount.MulInt64(gasLimit).TruncateInt())
		options.FeeAmount = types.NewCoins(gasFee)
	}
	txBuilder.SetGasLimit(options.GasLimit)
	txBuilder.SetFeeAmount(options.FeeAmount)

	return client.encodingConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
}

/**
 * TxSign 签名交易
 *
 * @param txBytes        交易序列化字节。
 * @param priv           私钥。
 * @param accountNumber  账户编号。来自BuildTx后的返回值。
 * @param sequence       账户已发送交易个数。来自BuildTx后的返回值。
 * @param overwriteSig   是否替换签名。多重签名大部分为false，根据实际情况设置。
 * @return 交易序列化字节
 */
func (client *RestClient) TxSign(txBytes []byte, priv types.PrivKey, accountNumber uint64, sequence uint64, overwriteSig bool) ([]byte, error) {
	tx, err := client.encodingConfig.TxConfig.TxDecoder()(txBytes)
	if err != nil {
		return nil, err
	}
	txBuilder, err := client.encodingConfig.TxConfig.WrapTxBuilder(tx)
	if err != nil {
		return nil, err
	}

	sigData := signing.SingleSignatureData{
		SignMode:  signing.SignMode_SIGN_MODE_DIRECT,
		Signature: nil,
	}
	sig := signing.SignatureV2{
		PubKey:   priv.PubKey(),
		Data:     &sigData,
		Sequence: sequence,
	}

	var prevSignatures []signing.SignatureV2
	if !overwriteSig {
		prevSignatures, err = txBuilder.GetTx().GetSignaturesV2()
		if err != nil {
			return nil, err
		}
	}

	if err := txBuilder.SetSignatures(sig); err != nil {
		return nil, err
	}

	signature, err := clitx.SignWithPrivKey(
		signing.SignMode_SIGN_MODE_DIRECT,
		types.SignerData{
			ChainID:       client.chainID,
			AccountNumber: accountNumber,
			Sequence:      sequence,
		},
		txBuilder,
		priv,
		client.encodingConfig.TxConfig,
		sequence,
	)

	if err != nil {
		return nil, err
	}

	if overwriteSig {
		if err := txBuilder.SetSignatures(signature); err != nil {
			return nil, err
		}
	} else {
		prevSignatures = append(prevSignatures, signature)
		if err := txBuilder.SetSignatures(prevSignatures...); err != nil {
			return nil, err
		}
	}
	return client.encodingConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
}

func (client *RestClient) post(path string, request proto.Message, response proto.Message) error {
	reqBytes, err := client.encodingConfig.Marshaler.MarshalJSON(request)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrMarshaler, "failed to marshal request: %s", err)
	}

	resp, err := client.c.Post(client.baseUrl+path, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return sdkerrors.Wrapf(types.ErrNotAccess, "path %s: %s", path, err)
	}

	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrResponseBody, "failed to read response body: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return sdkerrors.Wrapf(types.ErrResponseStatus, "path %s response code %d: %s", path, resp.StatusCode, string(out))
	}

	err = client.encodingConfig.Marshaler.UnmarshalJSON(out, response)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrMarshaler, "failed to unmarshal response: %s", err)
	}
	return nil
}

func (client *RestClient) get(path string, response proto.Message) error {
	resp, err := client.c.Get(client.baseUrl + path)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrNotAccess, "path %s: %s", path, err)
	}

	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrResponseBody, "failed to read response body: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		return sdkerrors.Wrapf(types.ErrResponseStatus, "path %s response code %d: %s", path, resp.StatusCode, string(out))
	}

	err = client.encodingConfig.Marshaler.UnmarshalJSON(out, response)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrMarshaler, "failed to unmarshal response: %s", err)
	}
	return nil
}

func (client *RestClient) MarshalJSON(message proto.Message) string {
	bts := client.encodingConfig.Marshaler.MustMarshalJSON(message)
	return string(bts)
}
