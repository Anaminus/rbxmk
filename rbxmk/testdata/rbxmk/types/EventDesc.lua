local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
T.Equal("AncestryChanged", desc:Member("Instance","AncestryChanged"), {
	MemberType = "Event",
	Name = "AncestryChanged",
	Security = "None",
	Parameters = {
		{Type={Category="Class",Name="Instance"}, Name="child"},
		{Type={Category="Class",Name="Instance"}, Name="parent"},
	},
	Tags = {},
})
