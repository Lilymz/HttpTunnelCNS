package main

import "time"

type RedisConfig struct {
	Ip       string `json:"Ip"`
	Port     int    `json:"Port"`
	Password string `json:"Password"`
}
type CnsConfig struct {
	UdpTimeOut time.Duration `yaml:"UdpTimeOut"`
	TcpTimeOut time.Duration `yaml:"TcpTimeOut"`
	ListenAddr []string      `yaml:"ListenAddr"`
	KeepAlive  bool          `yaml:"KeepAlive"`
	Account    string        `yaml:"Account"`
	*RedisConfig
	EnableTCPToUDP bool `yaml:"EnableTCPToUDP"`
	EnableHttpDNS  bool `yaml:"EnableHttpDNS"`
	EnableTFO      bool `yaml:"EnableTFO"`
}

var config = &CnsConfig{
	UdpTimeOut: 30,
	TcpTimeOut: 300,
	ListenAddr: []string{":8889", ":443"},
	KeepAlive:  true,
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
