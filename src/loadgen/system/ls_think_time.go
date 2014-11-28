package system

import (
	"encoding/json"
	. "loadgen"
	"log"
	"time"
)

type LsThinkTime struct {
}

func (api *LsThinkTime) Name() string {
	return "ls_think_time"
}

func (api *LsThinkTime) PluginName() string {
	return pluginName
}

// func (api *LsThinkTime) Init(args interface{}) (interface{}, error) {
func (api *LsThinkTime) Init(args json.RawMessage) (interface{}, error) {
	log.Printf("ls_think_time.Init()")

	type thinkTime struct {
		MilliSeconds time.Duration `json:"milliseconds"`
	}
	var time thinkTime
	err := json.Unmarshal(args, &time)
	if err != nil {
		return nil, err
	}

	log.Printf("thinktime: %d\n", time.MilliSeconds)
	return time.MilliSeconds, nil
}

func (api *LsThinkTime) Run(parsedArgs interface{}, session *LsSession) error {
	log.Printf("session[]: ls_think_time.Run()\n") // TODO 要有session
	time.Sleep(parsedArgs.(time.Duration) * time.Millisecond)
	return nil
}

func (api *LsThinkTime) Terminate(parsedArgs interface{}) {
	log.Printf("ls_think_time.Terminate()")
	return
}

func init() {
	PluginApis["ls_think_time"] = &LsThinkTime{}
}
