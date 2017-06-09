package scheme

import (
	"github.com/anaminus/rbxmk"
)

var registryInput = map[string]rbxmk.InputScheme{}
var registryOutput = map[string]rbxmk.OutputScheme{}

func registerInput(name string, scheme rbxmk.InputScheme) {
	registryInput[name] = scheme
}

func registerOutput(name string, scheme rbxmk.OutputScheme) {
	registryOutput[name] = scheme
}

// Register registers the schemes implemented by this package to a given
// rbxmk.Schemes.
func Register(schemes *rbxmk.Schemes) {
	for name, scheme := range registryInput {
		schemes.RegisterInput(name, scheme)
	}
	for name, scheme := range registryOutput {
		schemes.RegisterOutput(name, scheme)
	}
}
