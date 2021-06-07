local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
T.Equal("GearType", desc:Enum("GearType"), {
	Name = "GearType",
	Tags = {"Deprecated"},
})
T.Equal("NormalId", desc:Enum("NormalId"), {
	Name = "NormalId",
	Tags = {},
})
