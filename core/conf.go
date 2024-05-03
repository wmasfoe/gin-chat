package core

import (
	"github.com/sirupsen/logrus"
	"go-admin/config"
	"go-admin/global"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// 读取 settings.yaml 的数据
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		logrus.Errorf("get config error %v", err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init unmarshal: %v", err)
	}
	logrus.Info("config yaml load success!")

	// 添加到全局变量
	global.CONFIG = c
}
