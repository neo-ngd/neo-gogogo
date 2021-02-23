package rpc

import "github.com/stretchr/testify/mock"

type RpcClientMock struct {
	mock.Mock
}

func (r *RpcClientMock) ClaimGas(s string) ClaimGasResponse {
	args := r.Called(s)
	return args.Get(0).(ClaimGasResponse)
}

func (r *RpcClientMock) GetAccountState(s string) GetAccountStateResponse {
	args := r.Called(s)
	return args.Get(0).(GetAccountStateResponse)
}
func (r *RpcClientMock) GetApplicationLog(s string) GetApplicationLogResponse {
	args := r.Called(s)
	return args.Get(0).(GetApplicationLogResponse)
}
func (r *RpcClientMock) GetAssetState(s string) GetAssetStateResponse {
	args := r.Called(s)
	return args.Get(0).(GetAssetStateResponse)
}
func (r *RpcClientMock) GetBalance(s string) GetBalanceResponse {
	args := r.Called(s)
	return args.Get(0).(GetBalanceResponse)
}
func (r *RpcClientMock) GetBestBlockHash() GetBestBlockHashResponse {
	args := r.Called()
	return args.Get(0).(GetBestBlockHashResponse)
}
func (r *RpcClientMock) GetBlockByHash(s string) GetBlockResponse {
	args := r.Called(s)
	return args.Get(0).(GetBlockResponse)
}
func (r *RpcClientMock) GetBlockByIndex(n uint32) GetBlockResponse {
	args := r.Called(n)
	return args.Get(0).(GetBlockResponse)
}
func (r *RpcClientMock) GetBlockCount() GetBlockCountResponse {
	args := r.Called()
	return args.Get(0).(GetBlockCountResponse)
}
func (r *RpcClientMock) GetBlockHeaderByHash(s string) GetBlockHeaderResponse {
	args := r.Called(s)
	return args.Get(0).(GetBlockHeaderResponse)
}
func (r *RpcClientMock) GetBlockHash(n uint32) GetBlockHashResponse {
	args := r.Called(n)
	return args.Get(0).(GetBlockHashResponse)
}

func (r *RpcClientMock) GetClaimable(s string) GetClaimableResponse {
	args := r.Called(s)
	return args.Get(0).(GetClaimableResponse)
}

func (r *RpcClientMock) GetConnectionCount() GetConnectionCountResponse {
	args := r.Called()
	return args.Get(0).(GetConnectionCountResponse)
}
func (r *RpcClientMock) GetContractState(s string) GetContractStateResponse {
	args := r.Called(s)
	return args.Get(0).(GetContractStateResponse)
}
func (r *RpcClientMock) GetNep5Balances(s string) GetNep5BalancesResponse {
	args := r.Called(s)
	return args.Get(0).(GetNep5BalancesResponse)
}
func (r *RpcClientMock) GetNep5Transfers(s string) GetNep5TransfersResponse {
	args := r.Called(s)
	return args.Get(0).(GetNep5TransfersResponse)
}
func (r *RpcClientMock) GetNewAddress() GetNewAddressResponse {
	args := r.Called()
	return args.Get(0).(GetNewAddressResponse)
}
func (r *RpcClientMock) GetPeers() GetPeersResponse {
	args := r.Called()
	return args.Get(0).(GetPeersResponse)
}
func (r *RpcClientMock) GetRawMemPool() GetRawMemPoolResponse {
	args := r.Called()
	return args.Get(0).(GetRawMemPoolResponse)
}
func (r *RpcClientMock) GetRawTransaction(s string) GetRawTransactionResponse {
	args := r.Called(s)
	return args.Get(0).(GetRawTransactionResponse)
}
func (r *RpcClientMock) GetStorage(s1 string, s2 string) GetStorageResponse {
	args := r.Called(s1, s2)
	return args.Get(0).(GetStorageResponse)
}
func (r *RpcClientMock) GetTransactionHeight(s string) GetTransactionHeightResponse {
	args := r.Called(s)
	return args.Get(0).(GetTransactionHeightResponse)
}
func (r *RpcClientMock) GetTxOut(s string, n int) GetTxOutResponse {
	args := r.Called(s)
	return args.Get(0).(GetTxOutResponse)
}
func (r *RpcClientMock) GetUnclaimed(s string) GetUnclaimedResponse {
	args := r.Called(s)
	return args.Get(0).(GetUnclaimedResponse)
}
func (r *RpcClientMock) GetUnclaimedGas() GetUnclaimedGasResponse {
	args := r.Called()
	return args.Get(0).(GetUnclaimedGasResponse)
}

func (r *RpcClientMock) GetUnspents(s string) GetUnspentsResponse {
	args := r.Called(s)
	return args.Get(0).(GetUnspentsResponse)
}

func (r *RpcClientMock) GetValidators() GetValidatorsResponse {
	args := r.Called()
	return args.Get(0).(GetValidatorsResponse)
}
func (r *RpcClientMock) GetVersion() GetVersionResponse {
	args := r.Called()
	return args.Get(0).(GetVersionResponse)
}
func (r *RpcClientMock) GetWalletHeight() GetWalletHeightResponse {
	args := r.Called()
	return args.Get(0).(GetWalletHeightResponse)
}
func (r *RpcClientMock) ImportPrivKey(s string) ImportPrivKeyResponse {
	args := r.Called(s)
	return args.Get(0).(ImportPrivKeyResponse)
}
func (r *RpcClientMock) InvokeFunction(s1 string, s2 string, s3 string, a ...interface{}) InvokeFunctionResponse {
	args := r.Called(s1, s2, a, s3)
	return args.Get(0).(InvokeFunctionResponse)
}
func (r *RpcClientMock) InvokeScript(s1 string, s2 string) InvokeScriptResponse {
	args := r.Called(s1, s2)
	return args.Get(0).(InvokeScriptResponse)
}
func (r *RpcClientMock) ListPlugins() ListPluginsResponse {
	args := r.Called()
	return args.Get(0).(ListPluginsResponse)
}
func (r *RpcClientMock) ListAddress() ListAddressResponse {
	args := r.Called()
	return args.Get(0).(ListAddressResponse)
}
func (r *RpcClientMock) SendFrom(assetId string, from string, to string, amount uint32, fee float32, changeAddress string) SendFromResponse {
	args := r.Called(assetId, from, to, amount, fee, changeAddress)
	return args.Get(0).(SendFromResponse)
}
func (r *RpcClientMock) SendRawTransaction(s string) SendRawTransactionResponse {
	args := r.Called(s)
	return args.Get(0).(SendRawTransactionResponse)
}
func (r *RpcClientMock) SendToAddress(assetId string, to string, amount uint32, fee float32, changeAddress string) SendToAddressResponse {
	args := r.Called(assetId, to, amount, fee, changeAddress)
	return args.Get(0).(SendToAddressResponse)
}
func (r *RpcClientMock) SubmitBlock(s string) SubmitBlockResponse {
	args := r.Called(s)
	return args.Get(0).(SubmitBlockResponse)
}
func (r *RpcClientMock) ValidateAddress(s string) ValidateAddressResponse {
	args := r.Called(s)
	return args.Get(0).(ValidateAddressResponse)
}
