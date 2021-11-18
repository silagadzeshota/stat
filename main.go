package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "math/big"
  "net/http"
  "strconv"
)

type EthereumBLockHeader struct {
  Result EthereumInnerBlockHeader
}

type EthereumInnerBlockHeader struct {
  Hash string
  Number string
  Miner string
  Difficulty string
  TotalDifficulty string
  Nonce string
}

func main() {

  stats := make(map[string]int)
  for k := 301; k<= 350; k++ {
    //run tron node json rpc request
  	i := new(big.Int)
  	i.SetString(strconv.Itoa(k), 10)

  	//shlog.Log(shlog.Info, "Getting block by number - {\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"0x"+ fmt.Sprintf("%x", i)+ "\", false],\"id\":1}")
  	response, err := jsonRPC([]byte("{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"0x"+fmt.Sprintf("%x", i)+"\", false],\"id\":1}"), "http://127.0.0.1:59796", "POST")
  	if err != nil {
  		fmt.Println("Node cannot get the block header by last block number: "+err.Error())
  		return
  	}

  	//parse returned json
  	var rpcBlockHeader EthereumBLockHeader
  	err = json.Unmarshal([]byte(response), &rpcBlockHeader)
  	if err != nil {
  		fmt.Println("Cannot parsse rpc output for block number")
  		return
  	}
    stats[rpcBlockHeader.Result.Miner]+=1
  }

  for i, v := range stats {
    fmt.Println(i," ",v)
  }

}


func jsonRPC(jsonStr []byte, url string, requestType string) (string, error) {
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println( err.Error())
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
