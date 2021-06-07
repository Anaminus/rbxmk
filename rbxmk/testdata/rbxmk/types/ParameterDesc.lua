local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
T.Equal("buttonId", desc:Member("PluginToolbar","CreateButton").Parameters[1], {
	Type = {Category="Primitive",Name="string"},
	Name = "buttonId",
})
T.Equal("tooltip", desc:Member("PluginToolbar","CreateButton").Parameters[2], {
	Type = {Category="Primitive",Name="string"},
	Name = "tooltip",
})
T.Equal("iconname", desc:Member("PluginToolbar","CreateButton").Parameters[3], {
	Type = {Category="Primitive",Name="string"},
	Name = "iconname",
})
T.Equal("text", desc:Member("PluginToolbar","CreateButton").Parameters[4], {
	Type = {Category="Primitive",Name="string"},
	Name = "text",
	Default = "",
})
