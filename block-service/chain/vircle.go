package chain

import (
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/util/log"
)

const (
	USESSL	= false
)


type Utxo struct {
	txid 			string	`json:txid`
	vout			int		`json:vout`
	address			string	`json:address`
	scriptPubKey	string	`json:scriptPubKey`
	amount			float64	`json:amount`
	confirmations	int		`json:confirmations`
	spendable		bool	`json:spendable`
	solvable		bool	`json:solvable`
	desc			string	`json:desc`
	safe			bool	`json:safe`
	stakeable		bool	`json:stakeable`
}


//生成新地址
func newAddress() (string, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return "", err
	}

	//生成一个新地址
	reqJson := "{\"method\":\"getnewaddress\"}"

	returnJson, err := rpcClient.send(reqJson)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]string)
	json.Unmarshal([]byte(returnJson), &n)
	return n["result"], nil
}

//获取UTXO
func listUnspent() (map[string]interface{}, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return nil, err
	}

	reqJson :=  "{\"method\":\"listunspent\",\"params\":[1,9999999,[],false,{\"maximumCount\":1,\"minimumAmount\":0.002}]}"

	returnJson, err := rpcClient.send(reqJson)
	if err != nil {
		return nil, err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJson), &n)
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

	reqJson :=  "{\"method\":\"dumpprivkey\",\"params\":[\"" + address + "\"]}"

	returnJson, err := rpcClient.send(reqJson)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJson), &n)

	key := (n["result"]).(string)
	if len(key) > 0 {
		return key,nil
	}

	return "",errors.New("dumpprivkey 异常")
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

	reqJson := "{\"method\":\"createrawtransaction\",\"params\":[" + string(txStr[:]) + "," + string(dataStr[:]) + "]}"
	log.Log(reqJson)
	returnJson, err := rpcClient.send(reqJson)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJson), &n)
	return (n["result"]).(string), nil
}

func jsonEscape(i string) string {
    b, err := json.Marshal(i)
    if err != nil {
        panic(err)
    }
    // Trim the beginning and trailing " character
    return string(b[1:len(b)-1])
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

	reqJson := "{\"method\":\"signrawtransactionwithkey\",\"params\":[\"" + hex + "\"," + string(dumpPrivkeyListStr[:]) + "," + string(utxoInfoListStr[:]) + "]}"

	returnJson, err := rpcClient.send(reqJson)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJson), &n)
	hexJson := (n["result"]).(map[string]interface{})
	return (hexJson["hex"]).(string), nil
}

func sendRawTransaction(sign string) (string, error) {
	rpcClient, err := CreateClient(USESSL)
	if err != nil {
		return "", err
	}

	reqJson :=  "{\"method\":\"sendrawtransaction\",\"params\":[\"" + sign + "\"]}"

	returnJson, err := rpcClient.send(reqJson)
	if err != nil {
		return "", err
	}

	//json 取值
	n := make(map[string]interface{})
	json.Unmarshal([]byte(returnJson), &n)
	return (n["result"]).(string),nil
}