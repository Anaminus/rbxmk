//go:build !sslkeylog

package main

import (
	"io"

	"github.com/anaminus/rbxmk"
)

// Do not enable for production builds! See ssl.go.
func injectSSLKeyLogFile(world *rbxmk.World, werr io.Writer) {}
