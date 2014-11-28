package loadgen

import (
	"encoding/json"
)

var Plugins LsPlugins // 加载时把插件和API保存进来。暂时用两个全局变量
var PluginApis LsPluginApis

func init() {
	Plugins = make(LsPlugins)
	PluginApis = make(LsPluginApis)
}

type LsPlugin interface {
	Name() string

	TaskInit(task *LsTask) error // 解析setting等
	TaskTerminate(task *LsTask)

	SessionInit(*LsSession) error    // 初始化对应的sessionstate和sessionvar
	SessionIterate(*LsSession) error // 更新sessionState和sessionvar
	SessionTerminate(*LsSession)
}

type LsPluginApi interface {
	Name() string
	PluginName() string

	Init(args json.RawMessage) (interface{}, error)
	Run(parsedArgs interface{}, session *LsSession) error
	Terminate(args interface{}) // Init返回的args。例如打开的文件，需要关闭
}

type LsPlugins map[string]LsPlugin
type LsPluginApis map[string]LsPluginApi
