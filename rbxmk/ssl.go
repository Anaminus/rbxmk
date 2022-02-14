//go:build sslkeylog

package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/anaminus/rbxmk"
)

// Override with -ldflags="-X main.sslKeyLogFileEnvVar=ENV_VAR_NAME"
var sslKeyLogFileEnvVar string

// If rbxmk's environment contains the variable indicated by
// sslKeyLogFileEnvVar, external programs such as Wireshark will be able to
// decrypt HTTPS connections made by rbxmk, which is useful for debugging.
// Because this also compromises security, it is not compiled by default; rbxmk
// must be built with the above tag in order for this to be enabled.
func injectSSLKeyLogFile(world *rbxmk.World, werr io.Writer) {
	// Expect environment variable name to be set.
	if sslKeyLogFileEnvVar == "" {
		return
	}
	// Get value of environment variable, which is expected to be a file path.
	log, ok := os.LookupEnv(sslKeyLogFileEnvVar)
	if !ok {
		return
	}
	// Open the file for writing.
	w, err := os.OpenFile(log, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Fprintf(werr, "cannot create file %q: %s", log, err)
		return
	}
	// Set HTTP transport to write keys to the opened file. See
	// crypto/tls.Config.KeyLogWriter for more information:
	//
	// https://pkg.go.dev/crypto/tls@latest#Config
	world.Client.Client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			KeyLogWriter: w,
		},
	}
}
