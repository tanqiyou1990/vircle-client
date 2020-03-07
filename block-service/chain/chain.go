package chain

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/micro/go-micro/util/log"
)

// CheckResult 用于返回数据上链信息查询结果
type CheckResult struct {
	Transaction string `json:"transaction"`
	BlockData   string `json:"blockData"`
}

// CheckBlockData 返回区块信息的JSON字符串.
// 参数为交易哈希
func CheckBlockData(txid string) (string, error) {

	//根据txid查询交易是否打包
	txOBJ, err := gettransaction(txid)
	if err != nil {
		return "", errors.New("查询交易信息失败")
	}
	if txOBJ["blockhash"] == nil {
		return "", errors.New("交易尚未打包区块")
	}
	blockHash := (txOBJ["blockhash"]).(string)
	blockOBJ, err := getblock(blockHash)
	if err != nil {
		return "", errors.New("查询区块信息失败")
	}

	blockJSON, err := json.Marshal(blockOBJ)
	if err != nil {
		return "", errors.New("区块数据转换失败")
	}

	txJSON, err := json.Marshal(txOBJ)
	if err != nil {
		return "", errors.New("交易数据转换失败")
	}

	result := &CheckResult{
		Transaction: string(txJSON),
		BlockData:   string(blockJSON),
	}

	resultJSON, err := json.Marshal(result)

	return string(resultJSON), nil
}

// CreateTxWithContent 数据上链
// 参数为待上链的数据字符串
func CreateTxWithContent(data string) (string, error) {

	//生成一个新地址
	address, err := newAddress()
	if err != nil {
		log.Log(err)
		return "", err
	}
	utxo, err := listUnspent()
	if err != nil {
		log.Log(err)
		return "", err
	}
	utxoAddress := (utxo["address"]).(string)
	txid := utxo["txid"]
	vout := utxo["vout"]
	amount := (utxo["amount"]).(float64)

	//dumpprivkey [UTXO address]
	dumpPrivkey, err := dumpPrivkey(utxoAddress)
	if err != nil {
		log.Log("dumpPrivkey:", err)
		return "", err
	}

	txMap := make(map[string]interface{})
	txMap["txid"] = txid
	txMap["vout"] = vout
	inputs := [1]interface{}{txMap}

	hexData := hex.EncodeToString([]byte(data))

	outputs := make(map[string]interface{})
	outputs["data"] = hexData
	outputs[address] = FloatRound(amount-0.001, 8)

	hex, err := createRawTransaction(inputs, outputs)
	if err != nil {
		log.Log("createRawTransaction:", err)
		return "", err
	}

	dumpPrivkeyList := [1]string{dumpPrivkey}
	utxoInfoMap := make(map[string]interface{})
	utxoInfoMap["txid"] = txid
	utxoInfoMap["vout"] = vout
	utxoInfoMap["scriptPubKey"] = utxo["scriptPubKey"]
	utxoInfoMap["amount"] = amount
	utxoInfoList := [1]interface{}{utxoInfoMap}

	txHex, err := signRawTransactionWithKey(hex, dumpPrivkeyList, utxoInfoList)
	if err != nil {
		log.Log("signRawTransactionWithKey:", err)
		return "", err
	}

	transactionID, err := sendRawTransaction(txHex)
	if err != nil {
		log.Log("sendRawTransaction:", err)
		return "", err
	}

	return transactionID, nil

}

// FloatRound 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
