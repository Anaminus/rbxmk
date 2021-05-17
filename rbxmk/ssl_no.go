// +build !sslkeylog

package main

import (
	"io"

	"github.com/anaminus/rbxmk"
)

func injectSSLKeyLogFile(world *rbxmk.World, werr io.Writer) {}
