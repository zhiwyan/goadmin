package config

import (
	"log"
	"testing"
)

func TestInitConfig(t *testing.T) {
	configpath := "../../config/config.toml"
	err := InitConfig(configpath)
	if err != nil {
		panic(err)
	}

	log.Println(Config)
}

/**
 * 不配置section 获取单个key
 * @param {[type]} t *testing.T [description]
 */
func TestGetKey(t *testing.T) {
	configpath := "../../config/config.toml"
	InitConfig(configpath)
}

func TestGetPort(t *testing.T) {
	configpath := "../../config/config.toml"
	InitConfig(configpath)
	log.Println(GetPort())
}
