package chain

import (
	"os"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strconv"
)

const (
	VERSION           = 0.1
	RPCCLIENT_TIMEOUT = 30
)

// 钱包连接参数
type rpcClient struct {
	serverAddr string
	user       string
	passwd     string
	httpClient *http.Client
}

//连接配置
func CreateClient(useSSL bool) (c *rpcClient, err error) {
	host := os.Getenv("RPC_HOST")
	port,err := strconv.Atoi(os.Getenv("RPC_PORT"))
	user := os.Getenv("RPC_USER")
	passwd := os.Getenv("RPC_PASSWD")
	var serverAddr string
	var httpClient *http.Client
	if useSSL {
		serverAddr = "https://"
		t := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		httpClient = &http.Client{Transport: t}
	} else {
		serverAddr = "http://"
		httpClient = &http.Client{}
	}
	c = &rpcClient{serverAddr: fmt.Sprintf("%s%s:%d", serverAddr, host, port), user: user, passwd: passwd, httpClient: httpClient}
	return
}

// 超时处理
func (c *rpcClient) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.httpClient.Do(req)
		done <- result{resp, err}
	}()
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("Timeout reading data from server")
	}
}

//通信
func (c *rpcClient) send(reqJson string) (retJSON string, err error) {
	connectTimer := time.NewTimer(RPCCLIENT_TIMEOUT * time.Second)
	reqJsonByte := []byte(reqJson)
	payloadBuffer := bytes.NewReader(reqJsonByte)
	req, err := http.NewRequest("POST", c.serverAddr, payloadBuffer)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Accept", "application/json")
	if len(c.user) > 0 || len(c.passwd) > 0 {
		req.SetBasicAuth(c.user, c.passwd)
	}
	resp, err := c.doTimeoutRequest(connectTimer, req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("HTTP error: " + resp.Status)
		return
	}
	retJSON = string(data)
	return
}