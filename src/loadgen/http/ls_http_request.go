package http

import (
	"encoding/json"
	. "loadgen"
	"log"
)

// 注册的API，全局唯一
type LsHttpRequest struct {
}

// args：每行API使用一个
type HttpReq struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Version string            `json:"version"`
	Headers map[string]string `json:"headers"`
	Body    []string          `json:"body"`
}

type MainHttpReq struct {
	HttpReq
	ReferReqs []HttpReq `json:"refers"`
}

func (api *LsHttpRequest) Name() string {
	return "ls_http_request"
}

func (api *LsHttpRequest) PluginName() string {
	return "ls_http"
}

// func (api *LsHttpRequest) Init(args interface{}) (interface{}, error) {
func (api *LsHttpRequest) Init(args json.RawMessage) (interface{}, error) {
	log.Printf("LsHttpRequest.Init()\n")

	var mainHttpReq MainHttpReq
	err := json.Unmarshal(args, &mainHttpReq)
	if err != nil {
		return nil, err
	}
	log.Printf("http_request.init() %#v\n", mainHttpReq)
	return mainHttpReq, nil
}

func (api *LsHttpRequest) Run(parsedArgs interface{}, session *LsSession) error {
	log.Printf("LsHttpRequest.Run()\n")
	return nil
}

func (api *LsHttpRequest) Terminate(parsedArgs interface{}) {
	log.Printf("LsHttpRequest.Terminate()\n")
	return
}

func init() {
	PluginApis["ls_http_request"] = &LsHttpRequest{}
}
