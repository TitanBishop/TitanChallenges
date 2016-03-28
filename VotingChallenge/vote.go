package main

import (
	"flag"
	"fmt"
	"net/http"
)

func voteGetFast(target string) {
	//target := "http://192.168.0.14/welcome_get.php?name=bishop&email=a11Id01sW1n@gmail.com"
	http.Get(target)
}

func main(){
	var target string = flag.Arg(1)

	for i:= 0; ; i++ {
		fmt.Println("Starting thread " , i)
		go voteGetFast(target)
	}
	for {

	}
}
