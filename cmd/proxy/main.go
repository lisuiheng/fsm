package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func ReverseProxy() gin.HandlerFunc {
	remote, err := url.Parse("http://localhost:8002")
	if err != nil {
		log.Fatalln("remote: ", err)
	}
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = c.Request.URL.Path
			if _, ok := req.Header["User-Agent"]; !ok {
				// explicitly disable User-Agent so it's not set to default value
				req.Header.Set("User-Agent", "")
			}
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.Transport = &transport{http.DefaultTransport}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()

	router.GET("/api/proxy", ReverseProxy())
	err := http.ListenAndServe(":8001", router)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

type transport struct {
	http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(b))
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	b = []byte("1")
	body := ioutil.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	return resp, nil
}
