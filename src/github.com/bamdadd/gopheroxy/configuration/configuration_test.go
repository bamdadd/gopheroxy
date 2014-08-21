package configuration

import (
	"fmt"
	"testing"
)

func TestShouldInitConfiguration(t *testing.T) {


	var mockConfigReader = func(p string) Configuration {
		fmt.Println(p)
		return Configuration{"localhost:8888", "localhost:8080", 25, 100}
	}

	config := ParseConfig(mockConfigReader, "fakepath")

	if config.Frontend != "localhost:8080" ||
		config.Backend != "localhost:8888" ||
		config.MaxConn != 25 ||
		config.MaxWaitConn != 100 {
		t.Fail()
	}
}

