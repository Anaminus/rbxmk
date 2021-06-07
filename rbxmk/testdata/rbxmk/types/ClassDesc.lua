local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
T.Equal("Instance", desc:Class("Instance"), {
	Name = "Instance",
	Superclass = "<<<ROOT>>>",
	MemoryCategory = "Instances",
	Tags = {"NotCreatable", "NotBrowsable"},
})
