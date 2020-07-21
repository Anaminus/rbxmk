package sources

import (
	"io/ioutil"

	"github.com/anaminus/rbxmk"
)

func File() rbxmk.Source {
	return rbxmk.Source{
		Name: "file",
		Read: func(options ...interface{}) (p []byte, err error) {
			return ioutil.ReadFile(options[0].(string))
		},
		Write: func(p []byte, options ...interface{}) (err error) {
			return ioutil.WriteFile(options[0].(string), p, 0666)
		},
	}
}
