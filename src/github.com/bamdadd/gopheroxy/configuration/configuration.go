package configuration

import (
	"io/ioutil"
	"gopkg.in/yaml.v1"
)



type Configuration struct {
	Backend string
	Frontend string
	MaxConn int
	MaxWaitConn int
}

func ReadConfig(f string) Configuration {
	var file, err = ioutil.ReadFile(f)
	var c Configuration
	if err != nil { panic(err) }

	yaml.Unmarshal([]byte(file), &c)
	if err != nil { panic(err) }

	return c

}

type ConfigParser func(path string) Configuration

func ParseConfig(parserMethod ConfigParser, path string) Configuration{
	return parserMethod(path)
}
