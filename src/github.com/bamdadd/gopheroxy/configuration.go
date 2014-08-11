package gopheroxy

import (
	"io/ioutil"
	"gopkg.in/yaml.v1"
)



type Configuration struct {
	Backend string
	Frontend string
}



func (c *Configuration) Read(f string) {
	var file, err = ioutil.ReadFile(f)
	if err != nil { panic(err) }

	yaml.Unmarshal([]byte(file), &c)
	if err != nil { panic(err) }

}

func ReadConfig(path string) *Configuration {
	config := &Configuration{}
	config.Read(path)
	return config
}
