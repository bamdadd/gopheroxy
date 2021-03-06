package main

import (
	"fmt"
	"github.com/bamdadd/gopheroxy/configuration"
	"github.com/bamdadd/gopheroxy/proxy"
)


func main() {
	c := configuration.ParseConfig(configuration.ReadConfig, "config/config.yml")
	fmt.Printf("Proxying %s->%s.\r\n", c.Frontend, c.Backend)


	proxy.ProxyTCP(&c)

}
