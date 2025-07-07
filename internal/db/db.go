package db

import (
	"encoding/json"
	"os"
)

const UserFile = "internal/db/data/users.json"
const TaskFile = "internal/db/data/tasks.json"

func LoadData(filename string, data interface{}) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)
}

func SaveData(filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}
