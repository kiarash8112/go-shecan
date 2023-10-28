package resolv

import config "github.com/go-shecan/configs"

var resolvInstance Resolve

func Set(r Resolve) {
	resolvInstance = r
}
func Get() Resolve {
	return resolvInstance
}

func New(c config.Config) Resolve {
	return Resolve{DnsServers: c.DnsServers, resolvPath: c.ResolvPath}
}
