local desc = rbxmk.newDesc("Function")

-- Metamethod tests
T.Pass("type of descriptor is FunctionDesc",
	typeof(desc) == "FunctionDesc")
T.Pass("metatable of descriptor is locked",
	type(getmetatable(desc)) == "string")
T.Pass("descriptor converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("descriptor is equal to itself",
	desc == desc)
T.Pass("descriptor is not equal to another descriptor of the same type",
	desc ~= rbxmk.newDesc("Function"))

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

-- Parameters
T.Pass("can call Parameters method",
	desc:Parameters())
T.Pass("Parameters method returns a table",
	type(desc:Parameters()) == "table")
T.Pass("Parameters method initializes with empty table",
	#desc:Parameters() == 0)
T.Pass("can call SetParameters method with table",
	function() desc:SetParameters({}) end)
T.Pass("SetParameters returns no values",
	select("#", desc:SetParameters({})) == 0)
T.Fail("cannot call SetParameters method with non-table",
	function() desc:SetParameters(42) end)
T.Fail("cannot call SetParameters method with nil",
	function() desc:SetParameters() end)
T.Fail("cannot call SetParameters method with ParameterDesc",
	function() desc:SetParameters(rbxmk.newDesc("Parameter")) end)
T.Pass("can call SetParameters method with table of ParameterDescs",
	function() desc:SetParameters({
		rbxmk.newDesc("Parameter", rbxmk.newDesc("Type", "FooCatA", "FooTypeA"), "fooNameA"),
		rbxmk.newDesc("Parameter", rbxmk.newDesc("Type", "FooCatB", "FooTypeB"), "fooNameB", "FooDefault"),
		rbxmk.newDesc("Parameter", rbxmk.newDesc("Type", "FooCatC", "FooTypeC"), "fooNameC"),
	}) end)
T.Fail("cannot call SetParameters method with table of non-ParameterDescs",
	function() desc:SetParameters({
		rbxmk.newDesc("Parameter", rbxmk.newDesc("Type", "FooCatA", "FooTypeA"), "fooNameA"),
		"Foobar",
		rbxmk.newDesc("Parameter", rbxmk.newDesc("Type", "FooCatC", "FooTypeC"), "fooNameC"),
	}) end)
T.Pass("first set parameter persists",
	function()
		return desc:Parameters()[1] == rbxmk.newDesc(
			"Parameter",
			rbxmk.newDesc("Type", "FooCatA", "FooTypeA"),
			"fooNameA"
		)
	end)
T.Pass("second set parameter persists",
	function()
		return desc:Parameters()[2] == rbxmk.newDesc(
			"Parameter",
			rbxmk.newDesc("Type", "FooCatB", "FooTypeB"),
			"fooNameB",
			"FooDefault"
		)
	end)
T.Pass("third set parameter persists",
	function()
		return desc:Parameters()[3] == rbxmk.newDesc(
			"Parameter",
			rbxmk.newDesc("Type", "FooCatC", "FooTypeC"),
			"fooNameC"
		)
	end)

-- ReturnType
T.Pass("can get ReturnType field",
	function() return desc.ReturnType end)
T.Pass("ReturnType field is a TypeDesc",
	function() return typeof(desc.ReturnType) == "TypeDesc" end)
T.Pass("ReturnType field initializes to empty TypeDesc",
	function() return desc.ReturnType == rbxmk.newDesc("Type") end)
T.Pass("can set ReturnType field to TypeDesc",
	function() desc.ReturnType = rbxmk.newDesc("Type", "FooCategory", "FooName") end)
T.Fail("cannot set ReturnType field to non-string",
	function() desc.ReturnType = 42 end)
T.Pass("set ReturnType field persists",
	function() return desc.ReturnType.Category == "FooCategory" and desc.ReturnType.Name == "FooName" end)

-- Security
T.Pass("can get Security field",
	function() return desc.Security end)
T.Pass("Security field is a string",
	function() return type(desc.Security) == "string" end)
T.Pass("Security field initializes to empty string",
	function() return desc.Security == "" end)
T.Pass("can set Security field to string",
	function() desc.Security = "Foobar" end)
T.Fail("cannot set Security field to non-string",
	function() desc.Security = 42 end)
T.Pass("set Security field persists",
	function() return desc.Security == "Foobar" end)

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
