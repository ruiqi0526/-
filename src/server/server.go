package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"html/template"

	"copy"
	"blockchain"
)

var blockChain *blockchain.BlockChain
var pathOfCopy string = "../copy/copy.json"
var c copyfunc.Check

//创建服务
func run() {
	http.HandleFunc("/blockchain/get", blockChainGetHandler)
	http.HandleFunc("/blockchain/write", blockChainWriteHandler)
	bs := http.FileServer(http.Dir("../bs"))
	http.Handle("/bs/", http.StripPrefix("/bs/", bs))
	http.ListenAndServe("localhost:8080", nil)
}

//打包区块数据
func blockChainGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(blockChain, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonResult := string(bytes)
	
	//往副本里写数据
	writeIntoCopy(jsonResult)

	c = copyfunc.CreateJsonStruct(c, pathOfCopy)

	//调用HTML文件
	t, err := template.ParseFiles("1.html")
	if err != nil {
		log.Fatal(err)
	}

	//显示在浏览器
	err = t.Execute(w, c.Blocks)
	if err != nil {
		log.Fatal(err)
	}
}

//往区块写入数据
func blockChainWriteHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data")
	blockChain.SendData(blockData)
	blockChainGetHandler(w, r)
}

//数据写入副本
func writeIntoCopy(json string) {
	err := ioutil.WriteFile(pathOfCopy, []byte(json), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

//程序入口
func main() {
	blockChain = blockchain.NewBlockChain()
	run()
}