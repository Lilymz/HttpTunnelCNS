package main

import "time"

type RedisConfig struct {
	Ip       string `json:"Ip"`
	Port     int    `json:"Port"`
	Password string `json:"Password"`
}
type CnsConfig struct {
	UdpTimeOut int           `json:"UdpTimeOut"`
	TcpTimeOut int           `json:"TcpTimeOut"`
	ListenAddr []string      `json:"ListenAddr"`
	KeepAlive  time.Duration `json:"KeepAlive"`
	Account    string        `json:"Account"`
	*RedisConfig
	EnableTCPToUDP bool `json:"EnableTCPToUDP"`
	EnableHttpDNS  bool `json:"EnableHttpDNS"`
	EnableTFO      bool `json:"EnableTFO"`
}

var config = &CnsConfig{
	UdpTimeOut: 30,
	TcpTimeOut: 300,
	ListenAddr: []string{"7998"},
	KeepAlive:  time.Minute * 5,
	Account:    "cnsProxy",
	RedisConfig: &RedisConfig{
		Ip:       "",
		Port:     6389,
		Password: "cnsRedis",
	},
	EnableTCPToUDP: true,
	EnableHttpDNS:  true,
	EnableTFO:      true,
}

func main() {

}
