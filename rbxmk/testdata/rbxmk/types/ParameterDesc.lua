local typeDesc = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")
local desc = rbxmk.newDesc("ParameterDesc", typeDesc, "fooName")
local descDefault = rbxmk.newDesc("ParameterDesc", typeDesc, "fooName", "FooDefault")

-- Metamethod tests
T.Pass("type of descriptor is ParameterDesc",
	typeof(desc) == "ParameterDesc")
T.Pass("metatable of descriptor is locked",
	type(getmetatable(desc)) == "string")
T.Pass("descriptor converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("descriptor can be compared with itself",
	desc == desc)
T.Pass("descriptor can be compared with a matching ParameterDesc",
	desc == desc)
T.Pass("descriptor can be compared with a non-matching ParameterDesc",
	desc ~= descDefault)
T.Pass("descriptor with default can be compared with a matching ParameterDesc",
	descDefault == descDefault)
T.Pass("descriptor with default can be compared with a non-matching ParameterDesc",
	descDefault ~= desc)

-- Member tests

-- Type
T.Fail("descriptor cannot set Type field",
	function() desc.Type = typeDesc end)
T.Pass("descriptor can get Type field",
	function() return desc.Type == typeDesc end)

-- Name
T.Fail("descriptor cannot set Name field",
	function() desc.Name = "Foobar" end)
T.Pass("descriptor can get Name field",
	function() return desc.Name == "fooName" end)

-- Default
T.Fail("descriptor cannot set Default field",
	function() desc.Default = "Foobar" end)
T.Pass("descriptor can get Default field",
	function() return desc.Default == nil end)
T.Fail("descriptor with default cannot set Default field",
	function() descDefault.Default = nil end)
T.Pass("descriptor with default can get Default field",
	function() return descDefault.Default == "FooDefault" end)
