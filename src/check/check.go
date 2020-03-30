package main

import "copy"

var (
	c copyfunc.Check
	path = "../copy/copy.json"
)

func main() {
	c = copyfunc.CreateJsonStruct(c, path)
	copyfunc.PrintCopyJson(c)
}