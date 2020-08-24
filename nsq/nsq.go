package nsq

type Nsq struct {
	Topic string
	Host  string
	Port  string
}

func NewNsq(host, port string) Nsq {
	return Nsq{
		Host: host,
		Port: port,
	}
}
