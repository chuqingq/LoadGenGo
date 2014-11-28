package main

import (
	"encoding/json"
	"io/ioutil"
	"loadgen"
	loadgenhttp "loadgen/http"
	"log"
	"net/http"
	"os"
)

type hPair struct {
	Request httpReq
	// Response httpRes
}

type httpReq struct {
	Method  string
	URL     string
	Version string
	Headers http.Header
	Body    string
}

// type httpRes struct {
// 	Version string
// 	Status  string
// 	Headers http.Header
// 	Body    string
// }

func main() {
	// 从当前目录的http_recordstream文件中读取
	stream, err := readStreamFile()
	if err != nil {
		panic("readStreamFile error: " + err.Error())
	}

	events := make([]*hPair, 0, 1)
	err = json.Unmarshal(stream, &events)
	if err != nil {
		panic("unmarshal stream error: " + err.Error())
	}

	// 读取脚本
	script, err := loadgen.ReadScript()
	if err != nil {
		panic("ReadScript error: " + err.Error())
	}
	// log.Printf("oldscript=%#v\n", script)
	action := script.Action[:]

	// 为events分组
	group := make([]*httpReq, 0, 1)
	referer := ""
	for _, e := range events {
		ref := e.Request.Headers.Get("Referer")
		if ref == "" { // 当前是新的一个页面
			if len(group) != 0 {
				// action = append(action, group)
				action = addGroup(action, group)
				// log.Printf("newaction=%#v\n", action)

				script.Action = action
				// log.Printf("newscript=%#v\n", script)

				group = make([]*httpReq, 0, 1)
				referer = e.Request.URL
			} else {
				group = append(group, &e.Request)
				referer = ref
			}
		} else if ref == referer { // 当前不是新的页面
			group = append(group, &e.Request)
		} else {
			log.Printf("====unexpected", e.Request.URL)
		}
	}

	// TODO 把分组后的结果放在原脚本的action最后
	script.Write()
}

func addGroup(action []loadgen.LsScriptEntry, group []*httpReq) []loadgen.LsScriptEntry {
	// TODO 1 group转成MainHttpReq
	mainreq := group[0]
	// referreqs := group[1:]

	mainHttpReq := &loadgenhttp.MainHttpReq{
		HttpReq: loadgenhttp.HttpReq{
			Method:  mainreq.Method,
			URL:     mainreq.URL,
			Version: mainreq.Version,
			// TODO 1.1 需要处理 mainreq.Headers,
			Body: append(make([]string, 0, 0), mainreq.Body),
		},
		// referreqs, TODO 1.2 需要处理
	}
	// 2 把MainHttpReq转成json []byte
	bytes, err := json.Marshal(mainHttpReq)
	if err != nil {
		panic(err.Error())
	}
	// 3 把[]byte保存成RawMessage
	msg := &json.RawMessage{}
	err = msg.UnmarshalJSON(bytes)
	if err != nil {
		panic(err.Error())
	}

	// 4 设置RawMessage并增加到action中（后续会marshal）
	entry := loadgen.LsScriptEntry{
		ApiName: "ls_http_request",
		Args:    *msg,
		// map[string]interface{}{
		// 	"Method":        mainreq.Method,
		// 	"URL":           mainreq.URL,
		// 	"Version":       mainreq.Version,
		// 	"Headers":       mainreq.Headers,
		// 	"Body":          mainreq.Body,
		// 	"ReferRequests": referreqs,
		// },
	}
	action = append(action, entry)
	return action
}

// CANCEL 分组。取消，本模块只保存码流，不处理成脚本，分组应该放在httpparser中
func readStreamFile() ([]byte, error) {
	file, err := os.Open("http_recordstream")
	if err != nil {
		return nil, err
	}

	stream, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return stream, nil
}

// TODO callmodel/setting/vars都要改成显式init
