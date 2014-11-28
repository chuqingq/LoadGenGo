package loadgen

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type LsVars map[string]*LsVarConfig

// 从文件中读取变量的配置
type LsVarConfig struct {
	Name  string                 `json:"name"`
	Type  string                 `json:"type"`
	Attrs map[string]interface{} `json:"attrs"`
	Var   LsVar                  `json:"-"`
}

type LsVar interface {
	// *LsVarConfig
	NewSessionVar() (LsSessionVar, error) // 新建一个
	Terminate()                           // 任务停止时释放
}

// 会话中保存的变量，包含当前状态
type LsVarSession interface {
	// *LsVar
	SessionIterate() error
	SessionTerminate()

	GetNextValue() (string, error)
}

// func init() {
func ReadVars() (*LsVars, error) {
	vars := make(LsVars)

	file, err := os.Open("./task/var.json")
	if err != nil {
		// panic("failed to open var file: " + err.Error())
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&vars)
	if err != nil {
		// panic("Failed to decode var: " + err.Error())
		return nil, err
	}

	// 遍历vars，创建*LsVar
	for _, v := range vars {
		switch v.Type {
		case "random":
			err = NewVarRandom(v) // 内部设置Var这个字段
			if err != nil {
				// panic("failed to NewVarRandom: " + err.Error())
				return nil, err
			}
		case "unique":
			err := NewVarUnique(v)
			if err != nil {
				// panic("failed to NewVarUnique: " + err.Error())
				return nil, err
			}
		case "file":
			err := NewVarFile(v)
			if err != nil {
				// panic("failed to NewVarFile: " + err.Error())
				return nil, err
			}
		case "date":
			err := NewVarDate(v)
			if err != nil {
				// panic("failed to NewVarDate: " + err.Error())
				return nil, err
			}
		default:
			return nil, errors.New("var type is invalid: " + v.Type)
		}
	}

	log.Printf("Vars: %+v\n", vars)
	return &vars, nil
}

func (vars *LsVars) Terminate() {
	for name, v := range *vars {
		v.Var.Terminate()
		v.Var = nil

		delete(*vars, name) // TODO ？是否正确？
	}
}
