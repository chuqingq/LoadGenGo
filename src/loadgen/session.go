package loadgen

import (
	"log"
)

type LsSession struct {
	Id     int
	States LsSessionStates
	Vars   LsSessionVars
}

type LsSessionStates map[string]LsSessionState
type LsSessionState interface{}
type LsSessionVars map[string]LsSessionVar
type LsSessionVar interface {
	SessionInit() error
	SessionIterate() error
	SessionTerminate()
}

func NewSession() {
	log.Printf("  NewSession: %d\n", Task.CurId)
	// TODO LsSession后续可以改成静态分配，避免gc
	s := &LsSession{
		Id:     Task.CurId,
		States: make(LsSessionStates),
		Vars:   make(LsSessionVars),
	}
	go func() {
		s.init(Task.Vars)
	}()
	Task.Sessions = append(Task.Sessions, s)
	Task.CurId++
}

func (session *LsSession) init(vars *LsVars) {
	log.Printf("    session.init()\n")

	for _, p := range Plugins {
		err := p.SessionInit(session)
		if err != nil {
			panic("Plugin[" + p.Name() + "] SessionInit error: " + err.Error())
		}
	}
	// vars SessionInit()
	for varname, v := range *vars {
		sessionvar, err := v.Var.NewSessionVar()
		if err != nil {
			panic("failed to NewSessionVar for " + varname + ": " + err.Error())
		}

		session.Vars[varname] = sessionvar
	}

	session.run()
}

func (session *LsSession) run() {
	log.Printf("session.run()\n")

	// TODO 执行Init。如果一个API出错，则退出session
	for _, api := range Script.Init {
		session.runApi(&api)
	}

	for {
		for _, api := range Script.Action {
			session.runApi(&api)
		}

		// TODO 执行Action。每次结束后判断是否要迭代
		continued := true
		if !continued {
			break
		}
	}

	for _, api := range Script.End {
		session.runApi(&api)
	}

	// TODO 执行End。如果API出错，也继续执行
	session.terminate()
}

func (session *LsSession) runApi(api *LsScriptEntry) { // TODO 暂无返回值
	err := api.ApiFunc.Run(api.ParsedArgs, session)
	if err != nil {
		log.Printf("line[%d, %s] error: %s\n", api.Line, api.ApiName, err.Error())
		// TODO 暂时不报错
	}
}

func (session *LsSession) iterate() {
	// 调用所有协议的SessionIterate
	for _, p := range Plugins {
		err := p.SessionIterate(session)
		if err != nil {
			panic("Plugin[" + p.Name() + "] SessionInit error: " + err.Error())
		}
	}
}

func (session *LsSession) terminate() {
	// plugins SessionTerminate()
	for _, p := range Plugins {
		p.SessionTerminate(session)
		delete(session.States, p.Name())
	}

	// TODO vars SessionTerminate()
}
