local desc = rbxmk.newDesc("PropertyDesc")

-- Metamethod tests
T.Pass("type of descriptor is PropertyDesc",
	typeof(desc) == "PropertyDesc")
T.Pass("metatable of descriptor is locked",
	type(getmetatable(desc)) == "string")
T.Pass("descriptor converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("descriptor is equal to itself",
	desc == desc)
T.Pass("descriptor is not equal to another descriptor of the same type",
	desc ~= rbxmk.newDesc("PropertyDesc"))

-- Member tests

-- ValueType
T.Pass("can get ValueType field",
	function() return desc.ValueType end)
T.Pass("ValueType field is a TypeDesc",
	function() return typeof(desc.ValueType) == "TypeDesc" end)
T.Pass("ValueType field initializes to empty TypeDesc",
	function() return desc.ValueType == rbxmk.newDesc("TypeDesc") end)
T.Pass("can set ValueType field to TypeDesc",
	function() desc.ValueType = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName") end)
T.Fail("cannot set ValueType field to non-string",
	function() desc.ValueType = 42 end)
T.Pass("set ValueType field persists",
	function() return desc.ValueType.Category == "FooCategory" and desc.ValueType.Name == "FooName" end)

-- ReadSecurity
T.Pass("can get ReadSecurity field",
	function() return desc.ReadSecurity end)
T.Pass("ReadSecurity field is a string",
	function() return type(desc.ReadSecurity) == "string" end)
T.Pass("ReadSecurity field initializes to 'None'",
	function() return desc.ReadSecurity == "None" end)
T.Pass("can set ReadSecurity field to string",
	function() desc.ReadSecurity = "Foobar" end)
T.Fail("cannot set ReadSecurity field to non-string",
	function() desc.ReadSecurity = 42 end)
T.Pass("set ReadSecurity field persists",
	function() return desc.ReadSecurity == "Foobar" end)

-- WriteSecurity
T.Pass("can get WriteSecurity field",
	function() return desc.WriteSecurity end)
T.Pass("WriteSecurity field is a string",
	function() return type(desc.WriteSecurity) == "string" end)
T.Pass("WriteSecurity field initializes to 'None'",
	function() return desc.WriteSecurity == "None" end)
T.Pass("can set WriteSecurity field to string",
	function() desc.WriteSecurity = "Foobar" end)
T.Fail("cannot set WriteSecurity field to non-string",
	function() desc.WriteSecurity = 42 end)
T.Pass("set WriteSecurity field persists",
	function() return desc.WriteSecurity == "Foobar" end)

-- CanLoad
T.Pass("can get CanLoad field",
	function() local _ = desc.CanLoad end)
T.Pass("CanLoad field is a boolean",
	function() return type(desc.CanLoad) == "boolean" end)
T.Pass("CanLoad field initializes to false",
	function() return desc.CanLoad == false end)
T.Pass("can set CanLoad field to boolean",
	function() desc.CanLoad = true end)
T.Fail("cannot set CanLoad field to non-boolean",
	function() desc.CanLoad = 42 end)
T.Pass("set CanLoad field persists",
	function() return desc.CanLoad == true end)

-- CanSave
T.Pass("can get CanSave field",
	function() local _ = desc.CanSave end)
T.Pass("CanSave field is a boolean",
	function() return type(desc.CanSave) == "boolean" end)
T.Pass("CanSave field initializes to false",
	function() return desc.CanSave == false end)
T.Pass("can set CanSave field to boolean",
	function() desc.CanSave = true end)
T.Fail("cannot set CanSave field to non-boolean",
	function() desc.CanSave = 42 end)
T.Pass("set CanSave field persists",
	function() return desc.CanSave == true end)
