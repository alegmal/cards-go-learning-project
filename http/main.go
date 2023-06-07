package main

import (
	"os"
	"fmt"
	"net/http"
)

func main()  {
	resp, err := http.Get("http://www.google.com")

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}