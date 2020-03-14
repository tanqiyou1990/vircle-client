package ipfs

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

// UploadIPFS 上传至IPFS
func UploadIPFS(str string) (string, error) {
	ipfsURL := os.Getenv("IPFS_URL")
	sh = shell.NewShell(ipfsURL)
	cid, err := sh.Add(strings.NewReader(str))
	if err != nil {
		return "", err
	}
	return cid, nil
}

// UploadURLFile 上传文件至IPFS
func UploadURLFile(url string) (string, error) {

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	ipfsURL := os.Getenv("IPFS_URL")
	sh = shell.NewShell(ipfsURL)
	cid, err := sh.Add(bufio.NewReader(res.Body))
	if err != nil {
		return "", err
	}
	return cid, nil
}

// CatIPFS 查询IPFS数据
func CatIPFS(cid string) (string, error) {
	ipfsURL := os.Getenv("IPFS_URL")
	sh = shell.NewShell(ipfsURL)

	read, err := sh.Cat(cid)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(read)

	return string(body), nil
}
