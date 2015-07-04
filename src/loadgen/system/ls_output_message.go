package system

import (
	"encoding/json"
	"../../loadgen"
	"log"
)

type LsOutputMessage struct {
	loadgen.LsPluginApi
}

func (api *LsOutputMessage) Name() string {
	return "ls_output_message"
}

func (api *LsOutputMessage) PluginName() string {
	return pluginName
}

// func (api *LsOutputMessage) Init(args interface{}) (interface{}, error) {
func (api *LsOutputMessage) Init(args json.RawMessage) (interface{}, error) {
	log.Printf("ls_output_message.Init()")

	type outputMessage struct {
		Message string `json:"message"`
	}
	var msg outputMessage
	err := json.Unmarshal(args, &msg)
	if err != nil {
		return nil, err
	}

	return msg.Message, nil
}

func (api *LsOutputMessage) Run(parsedArgs interface{}, session *loadgen.LsSession) error {
	log.Printf("ls_output_message.Run()")

	log.Printf("====%s====\n", parsedArgs.(string))
	return nil
}

func (api *LsOutputMessage) Terminate(parsedArgs interface{}) {
	log.Printf("ls_output_message.Terminate()")
	return
}

func init() {
	loadgen.PluginApis["ls_output_message"] = &LsOutputMessage{}
}
