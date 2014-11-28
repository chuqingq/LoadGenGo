package loadgen

import (
	"encoding/json"
	"log"
	"os"
)

type LsSetting map[string]interface{}

// func init() {
func ReadSetting() (*LsSetting, error) {
	setting := make(LsSetting)

	file, err := os.Open("./task/setting.json")
	if err != nil {
		// panic("Failed to open setting file: " + err.Error())
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&setting)
	if err != nil {
		// panic("Failed to decode setting: " + err.Error())
		return nil, err
	}

	log.Printf("Setting: %+v\n", setting)
	return &setting, nil
}

func (setting *LsSetting) Terminate() {

}
