local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
T.Equal("Archivable", desc:Member("Instance","Archivable"), {
	MemberType = "Property",
	Name = "Archivable",
	ValueType = {Category="Primitive", Name="bool"},
	Category = "Behavior",
	ReadSecurity = "None",
	WriteSecurity = "None",
	CanLoad = false,
	CanSave = false,
	Tags = {},
})
T.Equal("ClassName", desc:Member("Instance","ClassName"), {
	MemberType = "Property",
	Name = "ClassName",
	ValueType = {Category="Primitive", Name="string"},
	Category = "Data",
	ReadSecurity = "None",
	WriteSecurity = "None",
	CanLoad = false,
	CanSave = false,
	Tags = {"ReadOnly", "NotReplicated"},
})
