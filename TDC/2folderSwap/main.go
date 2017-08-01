package main

import (
	"fmt"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	os.Chdir("folder")
	wd, _ = os.Getwd()
	fmt.Println(wd)
}
