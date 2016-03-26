package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func postForm() {
	for {
		host := "http://localhost/Voting_Challenge.html"
		resp, err := http.PostForm(host,
			url.Values{ "name": {"Bishop"}, "email": {"a11Id01sW1n@gmail.com"}} )
	}
}

func main(){
	for i:= 0; i < 100; i++ {
		fmt.Println("Starting thread " , i)
		go postForm()
	}
	for {

	}
}
