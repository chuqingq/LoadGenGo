package loadgen

import (
	"log"
)

var Task LsTask

type LsTask struct {
	Setting   *LsSetting
	Vars      *LsVars
	CallModel *LsCallModel
	Script    *LsScript

	CurId    int // 当前分配到的会话Id
	Sessions []*LsSession

	Stats LsStats
}

type LsStats map[string]map[string]*int64

func init() {
	Task.Stats = make(LsStats)
	Task.Sessions = make([]*LsSession, 0)
}

func (task *LsTask) Init(setting *LsSetting, vars *LsVars,
	callmodel *LsCallModel, script *LsScript) error {
	task.Setting = setting
	task.Vars = vars
	task.CallModel = callmodel
	task.Script = script

	// task.Sessions = make([]*LsSession, 64 /* 默认长度 */)

	// 调用Plugin的TaskInit
	for _, p := range Plugins {
		if err := p.TaskInit(task); err != nil {
			panic(err.Error())
		}
	}
	return nil
}

// TODO 阻塞执行，直到任务结束
func (task *LsTask) Run() error {
	log.Printf("====task.Run()")
	task.CallModel.Run()
	// TODO task逻辑：统计会话启动和停止
	//  停止有两种场景：
	//    1、如果callmodel停止，并且已启动会话数等于已停止会话数
	//    2、平滑停止：停止呼叫模型，等待正在运行的会话自动结束（第一种场景）
	//    2、立刻停止：停止呼叫模型，停止正在运行的会话
	// TODO 启动callmodel协程
	return nil
}

func (task *LsTask) Terminate() {
	log.Printf("LsTask.Terminate()\n")

	for _, p := range Plugins {
		p.TaskTerminate(task)
	}
}
