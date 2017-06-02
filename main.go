package main

import "bk-is-better/src"
import "fmt"

func main() {
	err := BkIsBetter.Initialize()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Complete")
	}
}
