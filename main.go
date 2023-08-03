package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error expected: ", err)
		os.Exit(1)

		fmt.Println(resp)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("How much bytes on: ", len(b))
	return len(b), nil
}
