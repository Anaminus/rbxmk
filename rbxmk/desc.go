package main

import (
	"fmt"
	"io"
	"os"

	"github.com/anaminus/pflag"
	"github.com/anaminus/rbxmk"
	"github.com/anaminus/rbxmk/formats"
	"github.com/anaminus/rbxmk/rtypes"
	"github.com/robloxapi/rbxdump"
	"github.com/robloxapi/rbxdump/diff"
)

const versionURL = `https://setup.rbxcdn.com/versionQTStudio`
const latestDumpURL = `https://setup.rbxcdn.com/%s-API-Dump.json`

type DescFlags struct {
	Latest  bool
	Files   []string
	Patches []string
}

// mergeDump merges two dumps by transforming prev to be the union of prev and
// next. That is, any changes that cause an item to be removed are excluded.
func mergeDump(prev, next *rbxdump.Root) *rbxdump.Root {
	actions := diff.Diff{Prev: prev, Next: next}.Diff()
	a := actions[:0]
	for _, action := range actions {
		if action.Type != diff.Remove {
			a = append(a, action)
		}
	}
	actions = a
	if prev == nil {
		prev = &rbxdump.Root{}
	}
	diff.Patch{Root: prev}.Patch(actions)
	return prev
}

func (d DescFlags) Resolve(client *rbxmk.Client) (desc *rtypes.RootDesc, err error) {
	var prev *rbxdump.Root
	if d.Latest {
		// Fetch version GUID.
		resp, err := client.Get(versionURL)
		if err != nil {
			return nil, fmt.Errorf("include latest descriptor: fetch version GUID: %w", err)
		}
		version, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if 200 < resp.StatusCode || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("include latest descriptor: fetch version GUID: %s", resp.Status)
		}
		if err != nil {
			return nil, fmt.Errorf("include latest descriptor: read version GUID: %w", err)
		}

		// Fetch dump.
		resp, err = client.Get(fmt.Sprintf(latestDumpURL, string(version)))
		if err != nil {
			return nil, fmt.Errorf("include latest descriptor: fetch dump: %w", err)
		}
		if 200 < resp.StatusCode || resp.StatusCode >= 300 {
			resp.Body.Close()
			return nil, fmt.Errorf("include latest descriptor: fetch version GUID: %s", resp.Status)
		}
		v, err := formats.Desc().Decode(rbxmk.Global{}, nil, resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("include latest descriptor: decode dump: %w", err)
		}
		prev = v.(*rtypes.RootDesc).Root
	}
	for _, path := range d.Files {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("include descriptor from file: open %s: %w", path, err)
		}
		v, err := formats.Desc().Decode(rbxmk.Global{}, nil, f)
		f.Close()
		if err != nil {
			return nil, fmt.Errorf("include descriptor from file: decode dump %s: %w", path, err)
		}
		next := v.(*rtypes.RootDesc).Root
		prev = mergeDump(prev, next)
	}
	for _, path := range d.Patches {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("include patch from file: open %s: %w", path, err)
		}
		v, err := formats.DescPatch().Decode(rbxmk.Global{}, nil, f)
		f.Close()
		if err != nil {
			return nil, fmt.Errorf("include patch from file: decode patches %s: %w", path, err)
		}
		descActions := v.(rtypes.DescActions)
		actions := make([]diff.Action, len(descActions))
		for i, action := range descActions {
			actions[i] = action.Action
		}
		if prev == nil {
			prev = &rbxdump.Root{}
		}
		diff.Patch{Root: prev}.Patch(actions)
	}
	if prev == nil {
		return nil, nil
	}
	return &rtypes.RootDesc{Root: prev}, nil
}

func (d *DescFlags) SetFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&d.Latest, "desc-latest", false, DocFlag("Flags/desc:desc-latest"))
	flags.Var(funcFlag(func(v string) error {
		d.Files = append(d.Files, v)
		return nil
	}), "desc-file", DocFlag("Flags/desc:desc-file"))
	flags.Var(funcFlag(func(v string) error {
		d.Patches = append(d.Patches, v)
		return nil
	}), "desc-patch", DocFlag("Flags/desc:desc-patch"))
}
