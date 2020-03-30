package copyfunc

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	
	"blockchain"
)

//被反射的结构体
type Check struct {
	Blocks []*blockchain.Block `json:"Blocks`
}

//打开副本的json文件
func OpenCopy(pathOfCopy string) ([]byte, error) {
	copyFile, err := os.OpenFile(pathOfCopy, os.O_RDWR, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer copyFile.Close()

	return ioutil.ReadAll(copyFile)
}

//json反射到结构体
func CreateJsonStruct(c Check, path string) Check {
	content, err := OpenCopy(path)
	if err != nil {
		fmt.Println("open file error: " + err.Error())
	}

	err = json.Unmarshal(content, &c)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}

	return c
}

//查看区块链数据
func PrintCopyJson(c Check) {
	for _, j := range c.Blocks {
		fmt.Printf("区块编号: %d\n", j.Index)
		fmt.Printf("上一个区块哈希值: %s\n", j.PrevBlockHash)
		fmt.Printf("当前区块哈希值: %s\n", j.Hash)
		fmt.Printf("区块内数据: %s\n", j.Data)
		fmt.Printf("打包时间: %s\n", j.Timestamp)
		fmt.Println("-----------------------------------------------")
	}
}