package ipfs

import (
	"github.com/ipfs/go-ipfs-api"
	"os"
	"bufio"
	"net/http"
	"strings"
	"io/ioutil"
)

var sh *shell.Shell

func UploadIPFS(str string) (string, error) {
	ipfsUrl := os.Getenv("IPFS_URL")
	sh = shell.NewShell(ipfsUrl)
	cid, err := sh.Add(strings.NewReader(str))
	if err != nil {
		return "", err
		os.Exit(1)
	}
	return cid, nil
}

func UploadUrlFile(url string) (string, error) {
	
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()


	ipfsUrl := os.Getenv("IPFS_URL")
	sh = shell.NewShell(ipfsUrl)
	cid, err := sh.Add(bufio.NewReader(res.Body))
	if err != nil {
		return "", err
		os.Exit(1)
	}
	return cid, nil
}

func CatIPFS(cid string) (string, error) {
	ipfsUrl := os.Getenv("IPFS_URL")
	sh = shell.NewShell(ipfsUrl)

	read, err := sh.Cat(cid)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(read)

	return string(body), nil
}