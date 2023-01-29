package nsq

import (
	"goadmin/lib/config"

	"github.com/nsqio/go-nsq"
)

var NsqProducer *nsq.Producer

func InitNsqProducer() error {
	var err error
	NsqProducer, err = nsq.NewProducer(config.Config.NsqAddr, nsq.NewConfig())
	if err != nil {
		return err
	}

	if err = NsqProducer.Ping(); err != nil {
		return err
	}

	return nil
}
