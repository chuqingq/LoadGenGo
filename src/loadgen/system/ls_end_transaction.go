package system

import (
	"encoding/json"
	"../../loadgen"
	"log"
	"time"
)

type LsEndTransaction struct {
	loadgen.LsPluginApi
}

func (api *LsEndTransaction) Name() string {
	return "ls_end_transaction"
}

func (api *LsEndTransaction) PluginName() string {
	return pluginName
}

// func (api *LsEndTransaction) Init(args interface{}) (interface{}, error) {
func (api *LsEndTransaction) Init(args json.RawMessage) (interface{}, error) {
	log.Printf("ls_end_transaction.Init()")
	type startTrans struct {
		Name string `json:"transaction_name"`
	}
	var tran startTrans
	err := json.Unmarshal(args, &tran)
	if err != nil {
		return nil, err
	}
	return tran.Name, nil
}

func (api *LsEndTransaction) Run(parsedArgs interface{}, session *loadgen.LsSession) error {
	log.Printf("ls_end_transaction.Run()")

	startTime := session.States[pluginName].(*LsSessionStateSystem).Trans[parsedArgs.(string)]

	duration := time.Since(startTime)
	log.Printf("duration=%d\n", duration)
	// TODO
	return nil
}

func (api *LsEndTransaction) Terminate(parsedArgs interface{}) {
	log.Printf("ls_end_transaction.Terminate()")
	return
}

func init() {
	loadgen.PluginApis["ls_end_transaction"] = &LsEndTransaction{}
}
