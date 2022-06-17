package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	response, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// response.Body.Read(bs)
	// fmt.Println((string(bs)))

	// io.Copy(os.Stdout, response.Body)

	lw := logWriter{}
	io.Copy(lw, response.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just Wrote many line", len(bs))
	return len(bs), nil
}
