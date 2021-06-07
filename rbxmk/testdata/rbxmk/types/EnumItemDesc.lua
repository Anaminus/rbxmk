local desc = fs.read(path.expand("$sd/../../dump.desc.json"))
T.Equal("GearType.MeleeWeapons", desc:EnumItem("GearType", "MeleeWeapons"), {
	Name = "MeleeWeapons",
	Value = 0,
	Index = 0,
	Tags = {"Deprecated"},
})
T.Equal("NormalId.Front", desc:EnumItem("NormalId", "Front"), {
	Name = "Front",
	Value = 5,
	Index = 3,
	Tags = {},
})
