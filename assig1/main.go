package main

import "fmt"

var Data map[string]string

func init() {
	Data = make(map[string]string)
}
func Add(k, v string) {
	Data[k] = v
}

func main() {
	Add("name", "raj")
	Add("lastname", "das")
	fmt.Println(Data)
}
