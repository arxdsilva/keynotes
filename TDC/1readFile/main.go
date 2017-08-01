package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// dir, _ := ioutil.ReadDir(".")
	// for _, v := range dir {
	// 	fmt.Println(v.Name())
	// }
	v, _ := ioutil.ReadFile("read")
	fmt.Println(string(v))
}
