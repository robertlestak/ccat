package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type stringArr []string

var (
	body      string
	headers   stringArr
	cookies   stringArr
	basicAuth string
	reqMethod string
)

func (h *stringArr) String() string {
	var hs string
	hs = strings.Join(headers, ", ")
	return hs
}

func (h *stringArr) Set(v string) error {
	*h = append(*h, v)
	return nil
}

func init() {
	flag.StringVar(&body, "b", "", "Body to send when requesting remote resources")
	flag.Var(&cookies, "c", "Cookies to send when requesting remote resources")
	flag.Var(&headers, "H", "Headers to send when requesting remote resources")
	flag.StringVar(&reqMethod, "m", "GET", "Method when requesting remote resources")
	flag.StringVar(&basicAuth, "u", "", "Basic auth to send when requesting remote resources")
	flag.Parse()
	reqMethod = strings.ToUpper(reqMethod)
}

func setHeaders(req *http.Request, h []string) {
	for _, v := range h {
		hs := strings.Split(v, "=")
		if len(hs) <= 1 {
			continue
		}
		hk := hs[0]
		hv := hs[1]
		req.Header.Add(hk, hv)
	}
}

func setBasicAuth(req *http.Request, b string) {
	bs := strings.Split(b, ":")
	if len(bs) <= 1 {
		return
	}
	bu := bs[0]
	bp := bs[1]
	req.SetBasicAuth(bu, bp)
}

func setCookies(req *http.Request, c []string) {
	for _, v := range c {
		cs := strings.Split(v, "=")
		if len(cs) <= 1 {
			continue
		}
		cn := cs[0]
		cv := cs[1]
		req.AddCookie(&http.Cookie{Name: cn, Value: cv})
	}
}

func setRequest(req *http.Request) {
	setHeaders(req, headers)
	setBasicAuth(req, basicAuth)
	setCookies(req, cookies)
}

func main() {
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("usage: ccat [options] [file/url ...]")
		flag.PrintDefaults()
		return
	}
	for _, v := range args {
		if strings.Contains(v, "http") {
			c := &http.Client{}
			req, err := http.NewRequest(reqMethod, v, bytes.NewBufferString(body))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
				return
			}
			setRequest(req)
			res, err := c.Do(req)
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
