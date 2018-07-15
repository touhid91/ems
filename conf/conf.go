package conf

import (
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Ports struct {
		App   int8 `yml:"app"`
		Redis int8 `yml:"redis"`
		Mongo int8 `yml:"mongo"`
	}
}

func New() *Conf {
	env := os.Getenv("EMS_ENV")
	fn, _ := filepath.Abs(fmt.Sprintf("./conf/%s.yml", env))
	f, _ := ioutil.ReadFile(fn)

	var c Conf
	yaml.Unmarshal(f, &c)

	return &c
}
