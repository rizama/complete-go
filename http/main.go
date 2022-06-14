package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// response.Body.Read(bs)
	// fmt.Println((string(bs)))

	io.Copy(os.Stdout, response.Body)
}
