package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// this satisfies the http.Handler interface for the hotdog type
func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
