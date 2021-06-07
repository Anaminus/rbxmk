local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
T.Equal("FindFirstChild", desc:Member("Instance","FindFirstChild"), {
	MemberType = "Function",
	Name = "FindFirstChild",
	ReturnType = {Category="Class", Name="Instance"},
	Security = "None",
	Parameters = {
		{Type={Category="Primitive",Name="string"}, Name="name"},
		{Type={Category="Primitive",Name="bool"}, Name="recursive", Default="false"},
	},
	Tags = {},
})
