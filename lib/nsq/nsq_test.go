package nsq

import (
	"classroom/lib/config"
	"log"
	"testing"
)

func init() {
	configpath := "../../config/config.toml"
	err := config.InitConfig(configpath)
	if err != nil {
		panic(err)
	}
	//log.Println(config.Config)
}

func TestInitNsqProducer(t *testing.T) {
	var err error
	err = InitNsqProducer()
	log.Println(NsqProducer)
	err = NsqProducer.Publish("test_topic", []byte("test123"))
	if err != nil {
		panic(err)
	}
	log.Println("success")
}
