package loadgen

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

// 总呼叫模型LsCallModel

type LsCallModel struct {
	Type      string             `json:"type"` // CALLMODEL_VUSER | CALLMODEL_CAPS | CALLMODELS_TIPS
	InitValue int                `json:"init"`
	Phases    []LsCallModelPhase `json:"phases"`

	CurValue int // 当前会话数
	CurPhase int // 保存当前阶段
}

// 呼叫模型阶段LsCallModelPhase

type LsCallModelPhase struct {
	Accelerate int           `json:"accelerate"`
	Unit       time.Duration `json:"unit"`     // 每Unit秒，增加Accelerate
	Dest       int           `json:"dest"`     // 直到Dest
	Duration   time.Duration `json:"duration"` // 之后再呼叫Duration
}

func ReadCallModel() (*LsCallModel, error) {
	file, err := os.Open("./task/callmodel.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var callModel LsCallModel
	err = decoder.Decode(&callModel)
	if err != nil {
		return nil, err
	}

	log.Printf("CallModel: %+v\n", callModel)
	return &callModel, nil
}

func (callmodel *LsCallModel) Run() {
	// 先启动init
	log.Printf("init")
	for i := 0; i < callmodel.InitValue; i++ {
		NewSession()
	}
	callmodel.CurValue = callmodel.InitValue
	log.Printf("after init")

	// 运行每个phase
	log.Printf("phases")
	for ; callmodel.CurPhase < len(callmodel.Phases); callmodel.CurPhase++ {
		log.Printf("  phase[%d], cur:%d\n", callmodel.CurPhase, callmodel.CurValue)
		phase := &callmodel.Phases[callmodel.CurPhase]
		// 每unit秒，增加accelerate个session
		for callmodel.CurValue < phase.Dest { // TODO 当前简单写小于
			time.Sleep(phase.Unit * time.Second)
			log.Printf("    sleep: %d\n", phase.Unit)

			for i := 0; i < phase.Accelerate; i++ {
				NewSession()
				callmodel.CurValue++
			}
			log.Printf("    accelerate: %d\n", phase.Accelerate)
		}

		time.Sleep(phase.Duration * time.Second)
		log.Printf("  phase sleep; %d\n", phase.Duration)
	}
	log.Printf("after phases")
}

func (setting *LsCallModel) Terminate() {
	// do nothing
}
