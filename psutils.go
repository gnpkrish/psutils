package main

import (
	"io/ioutil"
	"fmt"
	"unicode"
)

func Getpidlist() {
	procList, err := ioutil.ReadDir("/proc")
	if err != nil{
		fmt.Println("Error")
	}
	for _, pid := range procList{
		if i := unicode.IsDigit(pid.Name()); i{
			fmt.Println(pid.Name())	
		}
	}

}

func main() {
	Getpidlist()
}	