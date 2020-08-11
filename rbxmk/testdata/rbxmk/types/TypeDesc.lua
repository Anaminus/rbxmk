local desc = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")

-- Metamethod tests
T.Pass("type of descriptor is TypeDesc",
	typeof(desc) == "TypeDesc")
T.Pass("metatable of descriptor is locked",
	type(getmetatable(desc)) == "string")
T.Pass("descriptor converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("descriptor can be compared with itself",
	desc == desc)
T.Pass("descriptor can be compared with a matching TypeDesc",
	desc == rbxmk.newDesc("TypeDesc", "FooCategory", "FooName"))
T.Pass("descriptor can be compared with a non-matching TypeDesc",
	desc ~= rbxmk.newDesc("TypeDesc", "BarCategory", "BarName"))

-- Member tests

-- Category
T.Fail("descriptor cannot set Category field",
	function() desc.Category = "Foobar" end)
T.Pass("descriptor can get Category field",
	function() return desc.Category == "FooCategory" end)

-- Name
T.Fail("descriptor cannot set Name field",
	function() desc.Name = "Foobar" end)
T.Pass("descriptor can get Name field",
	function() return desc.Name == "FooName" end)
