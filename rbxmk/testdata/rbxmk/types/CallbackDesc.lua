local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
T.Equal("OnServerInvoke", desc:Member("RemoteFunction","OnServerInvoke"), {
	MemberType = "Callback",
	Name = "OnServerInvoke",
	ReturnType = {Category="Group", Name="Tuple"},
	Security = "None",
	ThreadSafety = "Unsafe",
	Parameters = {
		{Type={Category="Class",Name="Instance"}, Name="player"},
		{Type={Category="Group",Name="Tuple"}, Name="arguments"},
	},
	Tags = {},
})
