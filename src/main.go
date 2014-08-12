package main

import (
	"fmt"
	"github.com/bamdadd/gopheroxy/configuration"
	"github.com/bamdadd/gopheroxy/proxy"
)


func main() {
	config := configuration.ReadConfig("config/config.yml")
	fmt.Printf("Proxying %s->%s.\r\n", *&config.Frontend, *&config.Backend)

	proxy.ProxyTCP(config)

}
