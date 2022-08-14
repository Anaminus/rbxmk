package main

import (
	"sort"
	"testing"

	"github.com/anaminus/rbxmk/dump"
	"github.com/anaminus/rbxmk/library"
)

func walkDumpEnvRef(ref *dump.EnvRef, visit func(ref *dump.EnvRef, path []string), path ...string) {
	visit(ref, path)
	l := make([]string, 0, len(ref.Fields))
	for k := range ref.Fields {
		l = append(l, k)
	}
	sort.Strings(l)
	for _, k := range l {
		v := ref.Fields[k]
		path = append(path, k)
		walkDumpEnvRef(v, visit, path...)
		path = path[:len(path)-1]
	}
}

func TestDumpEnv(t *testing.T) {
	root, err := GenerateDump(WorldOpt{
		IncludeLibraries: library.All(),
		ExcludeRoots:     true,
	})
	if err != nil {
		t.Fatal(err)
	}
	walkDumpEnvRef(root.Environment, func(ref *dump.EnvRef, path []string) {
		if root.Resolve(ref.Path...) == nil {
			t.Errorf("ref %v with path %v resolved to nil", path, ref.Path)
		}
	})
}
