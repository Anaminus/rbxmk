package main

import (
	"flag"

	"github.com/anaminus/snek"
)

func DocumentCommands() {
	for _, def := range Program.List() {
		Program.SetDoc(def.Name, snek.Doc{
			Summary:     Doc("commands/" + def.Name + "#Summary"),
			Arguments:   Doc("commands/" + def.Name + "#Arguments"),
			Description: Doc("commands/" + def.Name + "#Description"),
		})
		if def, ok := def.New().(snek.FlagSetter); ok {
			def.SetFlags(flag.NewFlagSet("", 0))
		}
	}
}
