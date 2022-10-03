package Json

import (
	"encoding/json"
	"os"
)

func JsonLineToStruct(dataJson string, structure any) error {
	b := []byte(dataJson)
	err := json.Unmarshal(b, &structure)
	if err != nil {
		return err
	}
	return nil
}

func JsonToStruct(namefile string, structure any) error {
	dataJson, err := os.ReadFile(namefile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataJson, &structure)
	if err != nil {
		return err
	}
	return nil
}

func StructToJson(structure any) (string, error) {
	u, err := json.Marshal(structure)
	if err != nil {
		return "", nil
	}
	return string(u), nil // {"Name":"Bob","Age":10,"Active":true}
}
