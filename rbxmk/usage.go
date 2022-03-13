package main

const usageTemplate = `{{frag "usage:Usage"}}{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} {{frag "usage:Arguments"}}{{end}}{{if gt (len .Aliases) 0}}

{{frag "usage:Aliases"}}
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

{{frag "usage:Examples"}}
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

{{frag "usage:Commands"}}{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

{{frag "usage:Flags"}}
{{width | .LocalFlags.FlagUsagesWrapped | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

{{frag "usage:GlobalFlags"}}
{{width | .InheritedFlags.FlagUsagesWrapped | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

{{frag "usage:HelpTopics"}}{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

{{frag "usage:MoreHelp" | expand "COMMAND" (print .CommandPath " " (frag "usage:Arguments") " --help") }}{{end}}
`
