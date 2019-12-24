package main

import "fmt"
import "encoding/base64"

func main(){
	var ubah,_ = base64.StdEncoding.DecodeString("SGlkYXlhdA==")
	var text = string(ubah)
	fmt.Println("Nama Saya :",text)
}