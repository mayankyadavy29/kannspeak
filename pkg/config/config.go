package config

import (
	"encoding/json"
	"fmt"
	"github.com/mayankyadavy29/kannspeak/pkg/constants"
	"os"
	"path/filepath"
	"runtime"
)

// Config stores the user defined config of the app. By default
type Config struct {
	Pool       constants.PoolEnum
	MaxPoolLen int
}

var configPath string
var config *Config

func DefaultConfig() *Config {
	return &Config{
		Pool:       constants.RANDOM,
		MaxPoolLen: 10,
	}
}

func GetConfig() *Config {
	return config
}

func init() {
	platform := runtime.GOOS
	if platform == "windows" {
		configPath = constants.DefWinCfg
	} else if platform == "linux" {
		configPath = constants.DefLinuxCfg
	}
	LoadConfig()
}

func LoadConfig() {
	cfg := DefaultConfig()
	_, err := os.Open(configPath)
	if err != nil {
		err = os.Mkdir(filepath.Dir(configPath), 0750)
		if err != nil {
			fmt.Println("error while creating KannSpeak directory", err)
			os.Exit(-1)
		}
		cfgFile, err := os.Create(configPath)
		if err != nil {
			fmt.Println(fmt.Errorf("cannot create config file at location:%s", configPath))
			os.Exit(-1)
		}
		cfgBytes, _ := json.Marshal(cfg)
		_, err = cfgFile.Write(cfgBytes)
		if err != nil {
			fmt.Println("error while saving config file")
			os.Exit(-1)
		}
		config = cfg
		return
	}
	cfgBytes, err := os.ReadFile(configPath)
	json.Unmarshal(cfgBytes, cfg)
	config = cfg
}

func SetConfig(cfg *Config) {
	cfgBytes, _ := json.Marshal(cfg)
	err := os.WriteFile(configPath, cfgBytes, 0666)
	if err != nil {
		fmt.Println("error while saving config file")
		os.Exit(-1)
	}
	config = cfg
}
