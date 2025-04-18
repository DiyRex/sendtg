package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type AppConfig struct {
	Bots  map[string]string `json:"bots"`
	Chats map[string]int64  `json:"chats"`
}

var config AppConfig
var configPath = filepath.Join(os.Getenv("HOME"), ".sendtg", "config.json")

func InitConfig() {
	config = AppConfig{
		Bots:  map[string]string{},
		Chats: map[string]int64{},
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		_ = os.MkdirAll(filepath.Dir(configPath), 0755)
		SaveConfig()
	} else {
		data, err := os.ReadFile(configPath)
		if err == nil {
			_ = json.Unmarshal(data, &config)
		}
		// Re-init if unmarshaling gave nil maps
		if config.Bots == nil {
			config.Bots = make(map[string]string)
		}
		if config.Chats == nil {
			config.Chats = make(map[string]int64)
		}
	}
}


func SaveConfig() {
	data, _ := json.MarshalIndent(config, "", "  ")
	_ = os.WriteFile(configPath, data, 0644)
}

func GetConfig() AppConfig {
	return config
}

func UpdateChat(name string, id int64) {
	if config.Chats == nil {
		config.Chats = make(map[string]int64)
	}
	config.Chats[name] = id
	SaveConfig()
}

func UpdateBot(name, token string) {
	if config.Bots == nil {
		config.Bots = make(map[string]string)
	}
	config.Bots[name] = token
	SaveConfig()
}

func DeleteBot(name string) {
	delete(config.Bots, name)
	SaveConfig()
}



func DeleteChat(name string) {
	delete(config.Chats, name)
	SaveConfig()
}

