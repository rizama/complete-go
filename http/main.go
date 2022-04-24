package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
