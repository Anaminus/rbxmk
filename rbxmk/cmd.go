package main

import (
	"flag"

	"github.com/anaminus/snek"
)

func DocumentCommands() {
	for _, def := range Program.List() {
		Program.SetDoc(def.Name, snek.Doc{
			Summary:     Doc("Commands/" + def.Name + ":Summary"),
			Arguments:   Doc("Commands/" + def.Name + ":Arguments"),
			Description: Doc("Commands/" + def.Name + ":Description"),
		})
		if def, ok := def.New().(snek.FlagSetter); ok {
			def.SetFlags(flag.NewFlagSet("", 0))
		}
	}
}
