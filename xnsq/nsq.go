package xnsq

import "github.com/learninto/goutil/conf"

type Nsq struct {
	Topic string
	Host  string
	Port  string
}

func NewNsq() Nsq {
	return Nsq{
		Host: conf.Get("NSQ_HOST"),
		Port: conf.Get("NSQ_PORT"),
	}
}
