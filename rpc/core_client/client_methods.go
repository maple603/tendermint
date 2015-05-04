// File generated by github.com/ebuchman/rpc-gen

package core_client

import (
	"fmt"
	"github.com/tendermint/tendermint/account"
	"github.com/tendermint/tendermint/binary"
	"github.com/tendermint/tendermint/rpc"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"io/ioutil"
	"net/http"
)

type Client interface {
	BlockchainInfo(minHeight uint, maxHeight uint) (*ctypes.ResponseBlockchainInfo, error)
	BroadcastTx(tx types.Tx) (*ctypes.ResponseBroadcastTx, error)
	Call(address []byte, data []byte) (*ctypes.ResponseCall, error)
	CallCode(code []byte, data []byte) (*ctypes.ResponseCall, error)
	DumpConsensusState() (*ctypes.ResponseDumpConsensusState, error)
	DumpStorage(address []byte) (*ctypes.ResponseDumpStorage, error)
	GenPrivAccount() (*ctypes.ResponseGenPrivAccount, error)
	GetAccount(address []byte) (*ctypes.ResponseGetAccount, error)
	GetBlock(height uint) (*ctypes.ResponseGetBlock, error)
	GetStorage(address []byte, key []byte) (*ctypes.ResponseGetStorage, error)
	ListAccounts() (*ctypes.ResponseListAccounts, error)
	ListUnconfirmedTxs() (*ctypes.ResponseListUnconfirmedTxs, error)
	ListValidators() (*ctypes.ResponseListValidators, error)
	NetInfo() (*ctypes.ResponseNetInfo, error)
	SignTx(tx types.Tx, privAccounts []*account.PrivAccount) (*ctypes.ResponseSignTx, error)
	Status() (*ctypes.ResponseStatus, error)
}

func (c *ClientHTTP) BlockchainInfo(minHeight uint, maxHeight uint) (*ctypes.ResponseBlockchainInfo, error) {
	values, err := argsToURLValues([]string{"minHeight", "maxHeight"}, minHeight, maxHeight)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["BlockchainInfo"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseBlockchainInfo `json:"result"`
		Error   string                         `json:"error"`
		Id      string                         `json:"id"`
		JSONRPC string                         `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) BroadcastTx(tx types.Tx) (*ctypes.ResponseBroadcastTx, error) {
	values, err := argsToURLValues([]string{"tx"}, tx)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["BroadcastTx"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseBroadcastTx `json:"result"`
		Error   string                      `json:"error"`
		Id      string                      `json:"id"`
		JSONRPC string                      `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) Call(address []byte, data []byte) (*ctypes.ResponseCall, error) {
	values, err := argsToURLValues([]string{"address", "data"}, address, data)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["Call"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseCall `json:"result"`
		Error   string               `json:"error"`
		Id      string               `json:"id"`
		JSONRPC string               `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) CallCode(code []byte, data []byte) (*ctypes.ResponseCall, error) {
	values, err := argsToURLValues([]string{"code", "data"}, code, data)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["CallCode"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseCall `json:"result"`
		Error   string               `json:"error"`
		Id      string               `json:"id"`
		JSONRPC string               `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) DumpConsensusState() (*ctypes.ResponseDumpConsensusState, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["DumpConsensusState"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseDumpConsensusState `json:"result"`
		Error   string                             `json:"error"`
		Id      string                             `json:"id"`
		JSONRPC string                             `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) DumpStorage(address []byte) (*ctypes.ResponseDumpStorage, error) {
	values, err := argsToURLValues([]string{"address"}, address)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["DumpStorage"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseDumpStorage `json:"result"`
		Error   string                      `json:"error"`
		Id      string                      `json:"id"`
		JSONRPC string                      `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) GenPrivAccount() (*ctypes.ResponseGenPrivAccount, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["GenPrivAccount"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseGenPrivAccount `json:"result"`
		Error   string                         `json:"error"`
		Id      string                         `json:"id"`
		JSONRPC string                         `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) GetAccount(address []byte) (*ctypes.ResponseGetAccount, error) {
	values, err := argsToURLValues([]string{"address"}, address)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["GetAccount"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseGetAccount `json:"result"`
		Error   string                     `json:"error"`
		Id      string                     `json:"id"`
		JSONRPC string                     `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) GetBlock(height uint) (*ctypes.ResponseGetBlock, error) {
	values, err := argsToURLValues([]string{"height"}, height)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["GetBlock"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseGetBlock `json:"result"`
		Error   string                   `json:"error"`
		Id      string                   `json:"id"`
		JSONRPC string                   `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) GetStorage(address []byte, key []byte) (*ctypes.ResponseGetStorage, error) {
	values, err := argsToURLValues([]string{"address", "key"}, address, key)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["GetStorage"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseGetStorage `json:"result"`
		Error   string                     `json:"error"`
		Id      string                     `json:"id"`
		JSONRPC string                     `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) ListAccounts() (*ctypes.ResponseListAccounts, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["ListAccounts"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseListAccounts `json:"result"`
		Error   string                       `json:"error"`
		Id      string                       `json:"id"`
		JSONRPC string                       `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) ListUnconfirmedTxs() (*ctypes.ResponseListUnconfirmedTxs, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["ListUnconfirmedTxs"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseListUnconfirmedTxs `json:"result"`
		Error   string                             `json:"error"`
		Id      string                             `json:"id"`
		JSONRPC string                             `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) ListValidators() (*ctypes.ResponseListValidators, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["ListValidators"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseListValidators `json:"result"`
		Error   string                         `json:"error"`
		Id      string                         `json:"id"`
		JSONRPC string                         `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) NetInfo() (*ctypes.ResponseNetInfo, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["NetInfo"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseNetInfo `json:"result"`
		Error   string                  `json:"error"`
		Id      string                  `json:"id"`
		JSONRPC string                  `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) SignTx(tx types.Tx, privAccounts []*account.PrivAccount) (*ctypes.ResponseSignTx, error) {
	values, err := argsToURLValues([]string{"tx", "privAccounts"}, tx, privAccounts)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["SignTx"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseSignTx `json:"result"`
		Error   string                 `json:"error"`
		Id      string                 `json:"id"`
		JSONRPC string                 `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientHTTP) Status() (*ctypes.ResponseStatus, error) {
	values, err := argsToURLValues(nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.addr+reverseFuncMap["Status"], values)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseStatus `json:"result"`
		Error   string                 `json:"error"`
		Id      string                 `json:"id"`
		JSONRPC string                 `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) BlockchainInfo(minHeight uint, maxHeight uint) (*ctypes.ResponseBlockchainInfo, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["BlockchainInfo"],
		Params:  []interface{}{minHeight, maxHeight},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseBlockchainInfo `json:"result"`
		Error   string                         `json:"error"`
		Id      string                         `json:"id"`
		JSONRPC string                         `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) BroadcastTx(tx types.Tx) (*ctypes.ResponseBroadcastTx, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["BroadcastTx"],
		Params:  []interface{}{tx},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseBroadcastTx `json:"result"`
		Error   string                      `json:"error"`
		Id      string                      `json:"id"`
		JSONRPC string                      `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) Call(address []byte, data []byte) (*ctypes.ResponseCall, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["Call"],
		Params:  []interface{}{address, data},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseCall `json:"result"`
		Error   string               `json:"error"`
		Id      string               `json:"id"`
		JSONRPC string               `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) CallCode(code []byte, data []byte) (*ctypes.ResponseCall, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["CallCode"],
		Params:  []interface{}{code, data},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseCall `json:"result"`
		Error   string               `json:"error"`
		Id      string               `json:"id"`
		JSONRPC string               `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) DumpConsensusState() (*ctypes.ResponseDumpConsensusState, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["DumpConsensusState"],
		Params:  []interface{}{},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseDumpConsensusState `json:"result"`
		Error   string                             `json:"error"`
		Id      string                             `json:"id"`
		JSONRPC string                             `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) DumpStorage(address []byte) (*ctypes.ResponseDumpStorage, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["DumpStorage"],
		Params:  []interface{}{address},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseDumpStorage `json:"result"`
		Error   string                      `json:"error"`
		Id      string                      `json:"id"`
		JSONRPC string                      `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) GenPrivAccount() (*ctypes.ResponseGenPrivAccount, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["GenPrivAccount"],
		Params:  []interface{}{},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseGenPrivAccount `json:"result"`
		Error   string                         `json:"error"`
		Id      string                         `json:"id"`
		JSONRPC string                         `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) GetAccount(address []byte) (*ctypes.ResponseGetAccount, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["GetAccount"],
		Params:  []interface{}{address},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseGetAccount `json:"result"`
		Error   string                     `json:"error"`
		Id      string                     `json:"id"`
		JSONRPC string                     `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) GetBlock(height uint) (*ctypes.ResponseGetBlock, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["GetBlock"],
		Params:  []interface{}{height},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseGetBlock `json:"result"`
		Error   string                   `json:"error"`
		Id      string                   `json:"id"`
		JSONRPC string                   `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) GetStorage(address []byte, key []byte) (*ctypes.ResponseGetStorage, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["GetStorage"],
		Params:  []interface{}{address, key},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseGetStorage `json:"result"`
		Error   string                     `json:"error"`
		Id      string                     `json:"id"`
		JSONRPC string                     `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) ListAccounts() (*ctypes.ResponseListAccounts, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["ListAccounts"],
		Params:  []interface{}{},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseListAccounts `json:"result"`
		Error   string                       `json:"error"`
		Id      string                       `json:"id"`
		JSONRPC string                       `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) ListUnconfirmedTxs() (*ctypes.ResponseListUnconfirmedTxs, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["ListUnconfirmedTxs"],
		Params:  []interface{}{},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseListUnconfirmedTxs `json:"result"`
		Error   string                             `json:"error"`
		Id      string                             `json:"id"`
		JSONRPC string                             `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) ListValidators() (*ctypes.ResponseListValidators, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["ListValidators"],
		Params:  []interface{}{},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseListValidators `json:"result"`
		Error   string                         `json:"error"`
		Id      string                         `json:"id"`
		JSONRPC string                         `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) NetInfo() (*ctypes.ResponseNetInfo, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["NetInfo"],
		Params:  []interface{}{},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseNetInfo `json:"result"`
		Error   string                  `json:"error"`
		Id      string                  `json:"id"`
		JSONRPC string                  `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) SignTx(tx types.Tx, privAccounts []*account.PrivAccount) (*ctypes.ResponseSignTx, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["SignTx"],
		Params:  []interface{}{tx, privAccounts},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseSignTx `json:"result"`
		Error   string                 `json:"error"`
		Id      string                 `json:"id"`
		JSONRPC string                 `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}

func (c *ClientJSON) Status() (*ctypes.ResponseStatus, error) {
	request := rpc.RPCRequest{
		JSONRPC: "2.0",
		Method:  reverseFuncMap["Status"],
		Params:  []interface{}{},
		Id:      0,
	}
	body, err := c.RequestResponse(request)
	if err != nil {
		return nil, err
	}
	var response struct {
		Result  *ctypes.ResponseStatus `json:"result"`
		Error   string                 `json:"error"`
		Id      string                 `json:"id"`
		JSONRPC string                 `json:"jsonrpc"`
	}
	binary.ReadJSON(&response, body, &err)
	if err != nil {
		return nil, err
	}
	if response.Error != "" {
		return nil, fmt.Errorf(response.Error)
	}
	return response.Result, nil
}