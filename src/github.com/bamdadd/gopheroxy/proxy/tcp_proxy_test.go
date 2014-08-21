package proxy

import (
	"fmt"
	"testing"
	"github.com/bamdadd/gopheroxy/configuration"
)



func TestShouldStartTCPProxyWithThePassedConfiguration(t *testing.T) {


	var configReader = func(p string) configuration.Configuration {
		fmt.Println(p)
		return configuration.Configuration{"localhost:8888", "localhost:8080", 25, 100}
	}

	config := configuration.ParseConfig(configReader, "fakepath")

	go ProxyTCP(&config)


}

