package xnsq

import (
	"strings"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

//初始化消费者
func (m Nsq) InitConsumer(topic, channel string, handler nsq.Handler) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second          //设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		// TODO 写日志
		panic(err)
	}
	c.AddHandler(handler) // 添加消费者接口

	//建立NSQLookupd连接
	addr := strings.Builder{}
	addr.WriteString(m.Host)
	addr.WriteString(":")
	addr.WriteString(m.Port)
	if err := c.ConnectToNSQD(addr.String()); err != nil {
		// TODO 写日志
		panic(err)
	}

	<-c.StopChan
}
