//go:build !no_desc_test

package main

import (
	"io"
	"io/fs"
	"testing"
	"time"

	"github.com/anaminus/cobra"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/reflect"
)

type nopFile struct{}

func (nopFile) Stat() (fs.FileInfo, error) { return nopFileInfo{}, nil }
func (nopFile) Read([]byte) (int, error)   { return 0, io.EOF }
func (nopFile) Close() error               { return nil }

type nopFileInfo struct{}

func (nopFileInfo) Name() string       { return "-" }
func (nopFileInfo) Size() int64        { return 0 }
func (nopFileInfo) Mode() fs.FileMode  { return 0 }
func (nopFileInfo) ModTime() time.Time { return time.Time{} }
func (nopFileInfo) IsDir() bool        { return false }
func (nopFileInfo) Sys() interface{}   { return nil }

// TestLatestDesc fetches the latest Roblox API dump and decodes it.
func TestLatestDesc(t *testing.T) {
	program := &cobra.Command{}
	program.SetArgs([]string{"run", "-"})
	program.SetIn(nopFile{})

	c := RunCommand{
		DescFlags: DescFlags{Latest: true},
		Init: func(c *RunCommand, s rbxmk.State) {
			c.DescFlags.Latest = true
			desc, err := c.Resolve(s.Client)
			if err != nil {
				t.Errorf("fetch latest descriptor: %s", err.Error())
				return
			}
			t.Log("Classes", len(desc.Classes))
			t.Log("Enums", len(desc.Enums))
			if _, err := reflect.RootDesc().PushTo(s.Context(), desc); err != nil {
				t.Errorf("reflect latest descriptor: %s", err.Error())
			}
		},
	}
	cmd := &cobra.Command{
		Use:  "run",
		RunE: c.Run,
	}
	c.SetFlags(cmd.PersistentFlags())
	program.AddCommand(cmd)

	if err := program.Execute(); err != nil {
		t.Error(err)
	}
}
