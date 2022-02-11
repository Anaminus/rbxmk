package main

func DocumentCommands() {
	for _, cmd := range Program.Commands() {
		name := cmd.Name()
		if cmd.Use == name {
			cmd.Use += " " + Doc("Commands/"+name+":Arguments")
		}
		cmd.Short = Doc("Commands/" + name + ":Summary")
		cmd.Long = Doc("Commands/" + name + ":Description")
	}
}
