local desc = rbxmk.newDesc("Property")

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
	desc ~= rbxmk.newDesc("Property"))

-- Member tests

-- Name
T.Pass("can get Name field",
	function() return desc.Name end)
T.Pass("Name field is a string",
	function() return type(desc.Name) == "string" end)
T.Pass("Name field initializes to empty string",
	function() return desc.Name == "" end)
T.Pass("can set Name field to string",
	function() desc.Name = "Foobar" end)
T.Fail("cannot set Name field to non-string",
	function() desc.Name = 42 end)
T.Pass("set Name field persists",
	function() return desc.Name == "Foobar" end)

-- ValueType
T.Pass("can get ValueType field",
	function() return desc.ValueType end)
T.Pass("ValueType field is a TypeDesc",
	function() return typeof(desc.ValueType) == "TypeDesc" end)
T.Pass("ValueType field initializes to empty TypeDesc",
	function() return desc.ValueType == rbxmk.newDesc("Type") end)
T.Pass("can set ValueType field to TypeDesc",
	function() desc.ValueType = rbxmk.newDesc("Type", "FooCategory", "FooName") end)
T.Fail("cannot set ValueType field to non-string",
	function() desc.ValueType = 42 end)
T.Pass("set ValueType field persists",
	function() return desc.ValueType.Category == "FooCategory" and desc.ValueType.Name == "FooName" end)

-- ReadSecurity
T.Pass("can get ReadSecurity field",
	function() return desc.ReadSecurity end)
T.Pass("ReadSecurity field is a string",
	function() return type(desc.ReadSecurity) == "string" end)
T.Pass("ReadSecurity field initializes to empty string",
	function() return desc.ReadSecurity == "" end)
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
T.Pass("WriteSecurity field initializes to empty string",
	function() return desc.WriteSecurity == "" end)
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

--Tags
T.Pass("Tag method returns a boolean",
	type(desc:Tag("")) == "boolean")
T.Pass("Tags method returns a table",
	type(desc:Tags()) == "table")
T.Pass("SetTag method returns no values",
	select("#", desc:SetTag()) == 0)
T.Pass("UnsetTag method returns no values",
	select("#", desc:UnsetTag()) == 0)
T.Pass("descriptor initializes with no tags",
	#desc:Tags() == 0)
T.Pass("Tag can receive string as first argument",
	function() desc:Tag("Foobar") end)
T.Fail("Tag cannot receive non-string as first argument",
	function() desc:Tag(42) end)
T.Pass("getting unset tag returns false",
	desc:Tag("TagA") == false)
T.Pass("SetTag receives strings as arguments",
	function() desc:SetTag("TagA") end)
T.Pass("SetTag can receive no arguments",
	function() desc:SetTag() end)
T.Pass("getting set tag returns true",
	desc:Tag("TagA") == true)
T.Pass("SetTag can receive multiple arguments",
	function() desc:SetTag("TagA", "TagB", "TagC") end)
T.Fail("SetTag cannot receive non-string argument",
	function() desc:SetTag("TagA", 42, "TagC") end)
T.Pass("first set tag persists",
	desc:Tag("TagA") == true)
T.Pass("second set tag persists",
	desc:Tag("TagB") == true)
T.Pass("third set tag persists",
	desc:Tag("TagC") == true)
T.Pass("Tags returns all three set tags",
	function()
		local tags = desc:Tags()
		return #tags == 3 and
		tags[1] == "TagA" and
		tags[2] == "TagB" and
		tags[3] == "TagC"
	end)

T.Pass("UnsetTag receives strings as arguments",
	function() desc:UnsetTag("TagA") end)
T.Pass("UnsetTag can receive no arguments",
	function() desc:UnsetTag() end)
T.Pass("unset tag persists",
	desc:Tag("TagA") == false)
T.Pass("Tags returns all two set tags",
	function()
		local tags = desc:Tags()
		return #tags == 2 and
		tags[1] == "TagB" and
		tags[2] == "TagC"
	end)
T.Pass("UnsetTag can receive multiple arguments",
	function() desc:UnsetTag("TagA", "TagB", "TagC") end)
T.Fail("UnsetTag cannot receive non-string argument",
	function() desc:UnsetTag("TagA", 42, "TagC") end)
T.Pass("first unset tag persists",
	desc:Tag("TagA") == false)
T.Pass("second unset tag persists",
	desc:Tag("TagB") == false)
T.Pass("third unset tag persists",
	desc:Tag("TagC") == false)
T.Pass("Tags returns no tags from all tags being unset",
	#desc:Tags() == 0)
