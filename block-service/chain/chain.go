package chain

import (
	"encoding/hex"
	"strconv"
	"fmt"
	"github.com/micro/go-micro/util/log"
)


func CreateTxWithContent(data string) (string, error) {

	//生成一个新地址
	address,err := newAddress()
	if err != nil {
		log.Log(err)
		return "", err
	}
	utxo,err := listUnspent()
	if err != nil {
		log.Log(err)
		return "", err
	}
	utxoAddress := (utxo["address"]).(string)
	txid := utxo["txid"]
	vout := utxo["vout"]
	amount := (utxo["amount"]).(float64)

	//dumpprivkey [UTXO address]
	dumpPrivkey,err := dumpPrivkey(utxoAddress)
	if err != nil {
		log.Log("dumpPrivkey:",err)
		return "", err
	}

	txMap := make(map[string]interface{})
	txMap["txid"] = txid
	txMap["vout"] = vout
	inputs := [1]interface{}{txMap}

	hexData := hex.EncodeToString([]byte(data))

	outputs := make(map[string]interface{})
	outputs["data"] = hexData
	outputs[address] = FloatRound(amount - 0.001, 8) 

	hex,err := createRawTransaction(inputs, outputs)
	if err != nil {
		log.Log("createRawTransaction:",err)
		return "", err
	}

	dumpPrivkeyList := [1]string{dumpPrivkey}
	utxoInfoMap := make(map[string]interface{})
	utxoInfoMap["txid"] = txid
	utxoInfoMap["vout"] = vout
	utxoInfoMap["scriptPubKey"] = utxo["scriptPubKey"]
	utxoInfoMap["amount"] = amount
	utxoInfoList := [1]interface{}{utxoInfoMap}

	txHex,err := signRawTransactionWithKey(hex, dumpPrivkeyList, utxoInfoList)
	if err != nil {
		log.Log("signRawTransactionWithKey:",err)
		return "", err
	}

	transactionId,err := sendRawTransaction(txHex)
	if err != nil {
		log.Log("sendRawTransaction:",err)
		return "", err
	}

	return transactionId, nil

}

// 截取小数位数
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}




