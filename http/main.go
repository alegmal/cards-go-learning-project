package main

import (
	"io"
	"os"
	"fmt"
	"net/http"
)

type logWriter struct {}

func main()  {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// io.Copy(os.Stdout, resp.Body)
}

func (logWriter) Write(bs []byte) (n int, err error){
	fmt.Println(string(bs))
	fmt.Printf("just wrote %v bytes",len(bs))
	
	return len(bs), nil
}