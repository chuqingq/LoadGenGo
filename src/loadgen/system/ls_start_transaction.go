package system

import (
	"encoding/json"
	. "loadgen"
	"log"
	"time"
)

type LsStartTransaction struct {
}

func (api *LsStartTransaction) Name() string {
	return "ls_start_transaction"
}

func (api *LsStartTransaction) PluginName() string {
	return pluginName
}

func (api *LsStartTransaction) Init(args json.RawMessage) (interface{}, error) {
	log.Printf("LsStartTransaction.Init()\n")
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

func (api *LsStartTransaction) Run(parsedArgs interface{}, session *LsSession) error {
	log.Printf("LsStartTransaction.Run()")

	trans := session.States[pluginName].(*LsSessionStateSystem).Trans
	trans[parsedArgs.(string)] = time.Now()

	return nil
}

func (api *LsStartTransaction) Terminate(parsedArgs interface{}) {
	log.Printf("LsStartTransaction.Terminate()")
	return
}

func init() {
	PluginApis["ls_start_transaction"] = &LsStartTransaction{}
}
