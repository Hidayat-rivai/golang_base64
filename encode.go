package main

import "fmt"
import "encoding/base64"

func main(){
	var text = "Hidayat"
	var ubah = base64.StdEncoding.EncodeToString([]byte (text))
	fmt.Println("Nama Saya :",ubah)
}