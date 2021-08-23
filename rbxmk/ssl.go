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

func injectSSLKeyLogFile(world *rbxmk.World, werr io.Writer) {
	log, ok := os.LookupEnv("SSLKEYLOGFILE")
	if !ok {
		return
	}
	w, err := os.OpenFile(log, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Fprintf(werr, "cannot create file %q: %s", log, err)
		return
	}
	world.Client.Client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			KeyLogWriter: w,
		},
	}
}
