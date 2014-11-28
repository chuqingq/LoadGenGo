package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// httphandler通知outputstream保存
var httpchan chan *hPair = make(chan *hPair)

var outputfile = flag.String("output", "./http_recordstream", "http record stream output file")

func main() {
	flag.Parse()

	// outputStream：接收http码流，并保存
	go outputStream()

	// httpproxy：把http码流发给outputStream
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:10080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = ""
	// 向真实服务器发请求
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Printf("%s %s : %s\n", r.Method, r.URL, resp.Status)

	// 响应头域
	for k, v := range resp.Header {
		if k == "Proxy-Connection" {
			continue
		}
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	// 响应状态码。需要放在头域后面
	w.WriteHeader(resp.StatusCode)

	// 响应消息体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		panic(err)
	}
	w.Write(body)

	// 把请求和响应按照指定的格式发给outputStream
	// req
	r.Header.Del("Proxy-Connection")
	req := httpReq{
		Method:  r.Method,
		URL:     r.URL.String(),
		Version: r.Proto,
		Headers: r.Header,
		Body:    "body", // TODO ioutil.ReadAll(r.req.Body)
	}
	// res
	res := httpRes{
		Version: resp.Proto,
		Status:  resp.Status,
		Headers: resp.Header,
		Body:    "", // TODO
	}
	httpchan <- &hPair{req, res}
}

func outputStream() {
	// 保存接收到的httpPair，相同referer的放在一个[]*hPair
	var events []*hPair = make([]*hPair, 0, 64)

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt, os.Kill)

	for {
		select {
		case pair := <-httpchan:
			events = append(events, pair)
		case <-sigchan:
			log.Printf("recorded %d events\n", len(events))
			saveStream(events)
			os.Exit(0)
		}
	}
}

type hPair struct {
	Request  httpReq
	Response httpRes
}

type httpReq struct {
	Method  string
	URL     string
	Version string
	Headers http.Header
	Body    string
}

type httpRes struct {
	Version string
	Status  string
	Headers http.Header
	Body    string
}

func saveStream(events []*hPair) {
	buf, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		panic("json.Marshal error: " + err.Error())
	}

	// 保存到文件中
	file, err := os.Create(*outputfile)
	if err != nil {
		panic("create outputstream file err: " + err.Error())
	}
	defer file.Close()

	_, err = file.Write(buf)
	if err != nil {
		panic("write outputstream err: " + err.Error())
	}
}
