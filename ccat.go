package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("usage: ccat [file/url ...]")
		return
	}
	for _, v := range os.Args[1:] {
		if strings.Contains(v, "http") {
			res, err := http.Get(v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
				return
			}
			defer res.Body.Close()
			io.Copy(os.Stdout, res.Body)
		} else {
			res, err := os.Open(v)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
				return
			}
			defer res.Close()
			io.Copy(os.Stdout, res)
		}
	}
}
