package loadgen

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var Script LsScript

type LsScript struct {
	Init   []LsScriptEntry `json:"init"`
	Action []LsScriptEntry `json:"action"`
	End    []LsScriptEntry `json:"end"`
}

type LsScriptEntry struct {
	// json中的内容
	// LineNo  int                    `json:"line"`
	ApiName string          `json:"api"`
	Args    json.RawMessage `json:"args"`

	// 解析后的内容
	Line       int         `json:"-"`
	ApiFunc    LsPluginApi `json:"-"`
	ParsedArgs interface{} `json:"-"`
}

const scriptFile = "./task/script.json"

func ReadScript() (*LsScript, error) {
	file, err := os.Open(scriptFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	// script := new(LsScript)
	script := &Script
	err = decoder.Decode(script)
	if err != nil {
		return nil, err
	}

	return script, nil
}

func (script *LsScript) Load(plugins LsPlugins, apis LsPluginApis) error {
	log.Printf("script.Load()\n")

	// 根据plugin和api的加载情况，设置LsScriptEntry中的Line和ApiFunc，执行init
	for index, entry := range script.Init {
		entry.Line = index
		api, ok := apis[entry.ApiName]
		if !ok {
			return errors.New("api " + entry.ApiName + " not found")
		}
		entry.ApiFunc = api

		// 执行api的init接口
		var err error
		entry.ParsedArgs, err = entry.ApiFunc.Init(entry.Args)
		if err != nil {
			return err
		}

		script.Init[index] = entry
	}

	for index, entry := range script.Action {
		entry.Line = index
		api, ok := apis[entry.ApiName]
		if !ok {
			return errors.New("api " + entry.ApiName + " not found")
		}
		entry.ApiFunc = api

		// 执行api的init接口
		var err error
		entry.ParsedArgs, err = entry.ApiFunc.Init(entry.Args)
		if err != nil {
			return err
		}

		script.Action[index] = entry
	}

	for index, entry := range script.End {
		entry.Line = index
		api, ok := apis[entry.ApiName]
		if !ok {
			return errors.New("api " + entry.ApiName + " not found")
		}
		entry.ApiFunc = api

		// 执行api的init接口
		var err error
		entry.ParsedArgs, err = entry.ApiFunc.Init(entry.Args)
		if err != nil {
			return err
		}

		script.End[index] = entry
	}

	return nil
}

func (script *LsScript) Unload() {
	// do nothing
}

func (script *LsScript) Write() error {
	buf, err := json.MarshalIndent(script, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(scriptFile)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(buf)
	return err
}
