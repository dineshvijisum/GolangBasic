package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	mes := []byte("Hi Dinesh")
	err := ioutil.WriteFile("Dinesh.txt", mes, 0644)
	if err != err {
		fmt.Println(err)
	}
	d,r := ioutil.ReadFile("Dinesh.txt")
     if r != nil {
		fmt.Println(r)
	}
	fmt.Println(string(d))
	}