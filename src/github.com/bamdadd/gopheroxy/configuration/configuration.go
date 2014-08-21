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

func (c *Configuration) GetBackend() string {
	return c.Backend
}

func (c *Configuration) GetFrontend() string {
	return c.Frontend
}

func (c *Configuration) GetMaxConn() int {
	return c.MaxConn
}

func (c *Configuration) GetMaxWaitConn() int {
	return c.MaxWaitConn
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
