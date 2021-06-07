local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
T.Equal("ValueType", desc:Member("Instance","Name").ValueType, {
	Category="Primitive",
	Name="string",
})
