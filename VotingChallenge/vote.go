package main

import (
	"fmt"
	"net/http"
)

func voteGetFast() {
	host := "http://192.168.0.14/welcome_get.php?name=bishop&email=a11Id01sW1n@gmail.com"
	http.Get(host)
}

func main(){
	for i:= 0; ; i++ {
		fmt.Println("Starting thread " , i)
		go voteGetFast()
	}
	for {

	}
}
