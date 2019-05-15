package core

import (
	"encoding/json"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

var Config config

type config struct {
	BasePath string
	Logger   Logger
	Database map[string]Database
	Listen   string
	AppEnv   string
	AppName  string
}

type Logger struct {
	Debug        bool
	OutFile      bool
	LogPath      string
	MaxAge       time.Duration //日志最大保存时间单位小时
	RotationTime time.Duration //日志切割时间间隔单位小时
}

type Database struct {
	DriverName     string
	DataSourceName string
}

func (c *config) Init() {
	appenv := os.Getenv("APPENV")
	if appenv == "" {
		Config.AppEnv = "dev"
	}

	file, err := os.Open("config." + Config.AppEnv + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	cwd, _ := os.Getwd()
	if Config.BasePath != "" {
		cwd = Config.BasePath
	} else {
		Config.BasePath = cwd
	}
}
