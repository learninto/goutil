package nsq

import (
	"strings"

	nsq "github.com/nsqio/go-nsq"
)

func (m Nsq) Publish(topic string, body []byte) error {
	config := nsq.NewConfig()

	addr := strings.Builder{}
	addr.WriteString(m.Host)
	addr.WriteString(":")
	addr.WriteString(m.Port)

	w, _ := nsq.NewProducer(addr.String(), config)
	if err := w.Publish(topic, body); err != nil {
		return err
	}

	w.Stop()
	return nil
}
