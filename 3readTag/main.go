package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// bad, 'array'/slice de bytes com cada letra
	v, _ := ioutil.ReadFile("readOrigin")
	fmt.Println(v)
	// f, _ := os.Open("readOrigin")
	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
}
