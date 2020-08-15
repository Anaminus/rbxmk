local desc = rbxmk.newDesc("EnumItemDesc")

-- Metamethod tests
T.Pass("type of descriptor is EnumItemDesc",
	typeof(desc) == "EnumItemDesc")
T.Pass("metatable of descriptor is locked",
	type(getmetatable(desc)) == "string")
T.Pass("descriptor converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("descriptor is equal to itself",
	desc == desc)
T.Pass("descriptor is not equal to another descriptor of the same type",
	desc ~= rbxmk.newDesc("EnumItemDesc"))

-- Value
T.Pass("can get Value field",
	function() return desc.Value end)
T.Pass("Value field is an int",
	function() return type(desc.Value) == "number" end)
T.Pass("Value field initializes to 0",
	function() return desc.Value == 0 end)
T.Pass("can set Value field to int",
	function() desc.Value = 42.5 end)
T.Fail("cannot set Value field to non-int",
	function() desc.Value = "Foobar" end)
T.Pass("set Value field persists",
	function() return desc.Value == 42 end)

-- Index
T.Pass("can get Index field",
	function() return desc.Index end)
T.Pass("Index field is an int",
	function() return type(desc.Index) == "number" end)
T.Pass("Index field initializes to 0",
	function() return desc.Index == 0 end)
T.Pass("can set Index field to int",
	function() desc.Index = 42.5 end)
T.Fail("cannot set Index field to non-int",
	function() desc.Index = "Foobar" end)
T.Pass("set Index field persists",
	function() return desc.Index == 42 end)
