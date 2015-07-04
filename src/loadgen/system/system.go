package system

import (
	. "../../loadgen"
	"time"
)

func init() {
	// plugin
	plugin := &LsPluginSystem{PluginName: "ls_system", PluginState: &LsSystemState{}}
	Plugins[plugin.PluginName] = plugin
	// api的注册放在各自api中
}

// plugin

type LsPluginSystem struct {
	PluginName  string
	PluginState *LsSystemState
	// LsPluginInterface // Name(), State(), PluginInit()
}

type LsSessionStateSystem struct {
	Trans map[string]time.Time // 记录事务的开始时间
}

func (s *LsPluginSystem) Name() string {
	return s.PluginName
}

func (s *LsPluginSystem) State() *LsSystemState {
	return s.PluginState
}

// func (s *LsPluginSystem) PluginInit(setting interface{}) error {
// 	return nil
// }

// func (s *LsPluginSystem) PluginTerminate() {
// }

func (s *LsPluginSystem) TaskInit(task *LsTask) error {
	// TODO 解析setting中的ignore_think_time
	return nil
}

func (s *LsPluginSystem) TaskTerminate(task *LsTask) {
}

func (s *LsPluginSystem) SessionInit(session *LsSession) error {
	session.States[pluginName] = &LsSessionStateSystem{
		Trans: make(map[string]time.Time),
	}
	return nil
}

func (s *LsPluginSystem) SessionIterate(session *LsSession) error {
	return nil
}

func (s *LsPluginSystem) SessionTerminate(session *LsSession) {
}

// state

type LsSystemState struct { // TODO 是否要大写？
	IgnoreThinkTime bool
}

const (
	pluginName = "ls_system"
)
