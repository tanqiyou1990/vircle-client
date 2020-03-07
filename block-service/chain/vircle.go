package chain

import (
	"encoding/json"
	"errors"

	"github.com/micro/go-micro/util/log"
)

const USESSL = false

//根据txid查找交易信息
func getblock(blockHash string) (map[string]interface{}, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return nil, err
	}

	//生成一个新地址
	reqJSON := "{\"method\":\"getblock\",\"params\":[\"" + blockHash + "\"]}"

	returnJSON, err := rpcClient.send(reqJSON)
	if err != nil {
		return nil, err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJSON), &n)
	txMap := (n["result"]).(map[string]interface{})
	return txMap, nil
}

//根据txid查找交易信息
func gettransaction(txid string) (map[string]interface{}, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return nil, err
	}

	//生成一个新地址
	reqJSON := "{\"method\":\"gettransaction\",\"params\":[\"" + txid + "\"]}"

	returnJSON, err := rpcClient.send(reqJSON)
	if err != nil {
		return nil, err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJSON), &n)
	txMap := (n["result"]).(map[string]interface{})
	return txMap, nil
}

//生成新地址
func newAddress() (string, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return "", err
	}

	//生成一个新地址
	reqJSON := "{\"method\":\"getnewaddress\"}"

	returnJSON, err := rpcClient.send(reqJSON)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]string)
	json.Unmarshal([]byte(returnJSON), &n)
	return n["result"], nil
}

//获取UTXO
func listUnspent() (map[string]interface{}, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return nil, err
	}

	reqJSON := "{\"method\":\"listunspent\",\"params\":[1,9999999,[],false,{\"maximumCount\":1,\"minimumAmount\":0.002}]}"

	returnJSON, err := rpcClient.send(reqJSON)
	if err != nil {
		return nil, err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJSON), &n)
	arr := (n["result"]).([]interface{})
	if len(arr) < 1 {
		return nil, errors.New("无可用的UTXO")
	}
	uxo := (arr[0]).(map[string]interface{})
	return uxo, nil
}

//dumpprivkey [UTXO address]
func dumpPrivkey(address string) (string, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return "", err
	}

	reqJSON := "{\"method\":\"dumpprivkey\",\"params\":[\"" + address + "\"]}"

	returnJSON, err := rpcClient.send(reqJSON)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJSON), &n)

	key := (n["result"]).(string)
	if len(key) > 0 {
		return key, nil
	}

	return "", errors.New("dumpprivkey 异常")
}

func createRawTransaction(inputs [1]interface{}, outputs map[string]interface{}) (string, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return "", err
	}

	txStr, err := json.Marshal(inputs)
	if err != nil {
		return "", err
	}
	dataStr, err := json.Marshal(outputs)
	if err != nil {
		return "", err
	}

	reqJSON := "{\"method\":\"createrawtransaction\",\"params\":[" + string(txStr[:]) + "," + string(dataStr[:]) + "]}"
	log.Log(reqJSON)
	returnJSON, err := rpcClient.send(reqJSON)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJSON), &n)
	return (n["result"]).(string), nil
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	// Trim the beginning and trailing " character
	return string(b[1 : len(b)-1])
}

func signRawTransactionWithKey(hex string, dumpPrivkeyList [1]string, utxoInfoList [1]interface{}) (string, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return "", err
	}

	dumpPrivkeyListStr, err := json.Marshal(dumpPrivkeyList)
	if err != nil {
		return "", err
	}
	utxoInfoListStr, err := json.Marshal(utxoInfoList)
	if err != nil {
		return "", err
	}

	reqJSON := "{\"method\":\"signrawtransactionwithkey\",\"params\":[\"" + hex + "\"," + string(dumpPrivkeyListStr[:]) + "," + string(utxoInfoListStr[:]) + "]}"

	returnJSON, err := rpcClient.send(reqJSON)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJSON), &n)
	hexJSON := (n["result"]).(map[string]interface{})
	return (hexJSON["hex"]).(string), nil
}

func sendRawTransaction(sign string) (string, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return "", err
	}

	reqJSON := "{\"method\":\"sendrawtransaction\",\"params\":[\"" + sign + "\"]}"

	returnJSON, err := rpcClient.send(reqJSON)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJSON), &n)
	return (n["result"]).(string), nil
}
