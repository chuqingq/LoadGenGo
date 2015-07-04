package main

import (
	"../../loadgen"
	_ "../http"
	_ "../system"
	"log"
	"runtime"
)

func main() {
	log.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(runtime.NumCPU()))

	// script
	script, err := loadgen.ReadScript()
	if err != nil {
		panic("NewScript error: " + err.Error())
	}
	err = script.Load(loadgen.Plugins, loadgen.PluginApis)
	if err != nil {
		panic(err.Error())
	}
	defer script.Unload()

	// callmodel
	callModel, err := loadgen.ReadCallModel()
	if err != nil {
		panic("ReadCallModel error: " + err.Error())
	}
	defer callModel.Terminate()

	// setting
	setting, err := loadgen.ReadSetting()
	if err != nil {
		panic("ReadSetting error: " + err.Error())
	}
	defer setting.Terminate()

	// vars
	vars, err := loadgen.ReadVars()
	if err != nil {
		panic("ReadVars error: " + err.Error())
	}

	// task
	err = loadgen.Task.Init(
		// &loadgen.Setting,
		setting,
		// &loadgen.Vars,
		vars,
		// &loadgen.CallModel,
		callModel,
		// &loadgen.Script,
		script,
	)
	if err != nil {
		panic(err.Error())
	}
	defer loadgen.Task.Terminate()

	err = loadgen.Task.Run()
	if err != nil {
		panic(err.Error())
	}
}
