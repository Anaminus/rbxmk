package main

import (
	"path"

	"github.com/anaminus/snek"
)

func DocumentCommands() {
	for _, def := range Program.List() {
		Program.SetDoc(def.Name, snek.Doc{
			Summary:     Doc(path.Join("commands", def.Name+".md", "Summary")),
			Arguments:   Doc(path.Join("commands", def.Name+".md", "Arguments")),
			Description: Doc(path.Join("commands", def.Name+".md", "Description")),
		})
	}
}
