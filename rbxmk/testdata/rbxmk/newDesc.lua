T.Fail("expects a string for its first argument",
	function() rbxmk.newDesc(42) end)
T.Pass("with 'Root', returns RootDesc",
	typeof(rbxmk.newDesc("Root")) == "RootDesc")
T.Pass("with 'Class', returns ClassDesc",
	typeof(rbxmk.newDesc("Class")) == "ClassDesc")
T.Pass("with 'Property', returns PropertyDesc",
	typeof(rbxmk.newDesc("Property")) == "PropertyDesc")
T.Pass("with 'Function', returns FunctionDesc",
	typeof(rbxmk.newDesc("Function")) == "FunctionDesc")
T.Pass("with 'Event', returns EventDesc",
	typeof(rbxmk.newDesc("Event")) == "EventDesc")
T.Pass("with 'Callback', returns CallbackDesc",
	typeof(rbxmk.newDesc("Callback")) == "CallbackDesc")
T.Pass("with 'Parameter', returns ParameterDesc",
	typeof(rbxmk.newDesc("Parameter")) == "ParameterDesc")
T.Pass("with 'Type', returns TypeDesc",
	typeof(rbxmk.newDesc("Type")) == "TypeDesc")
T.Pass("with 'Enum', returns EnumDesc",
	typeof(rbxmk.newDesc("Enum")) == "EnumDesc")
T.Pass("with 'EnumItem', returns EnumItemDesc",
	typeof(rbxmk.newDesc("EnumItem")) == "EnumItemDesc")
T.Fail("with unknown type, throws an error",
	function() rbxmk.newDesc("NonextantType") end)

T.Pass("with 'Type', has optional second string argument that sets Category field",
	rbxmk.newDesc("Type", "FooCategory").Category == "FooCategory")
T.Fail("with 'Type', with second non-string argument throws an error",
	function() rbxmk.newDesc("Type", 42) end)
T.Pass("with 'Type', has optional third string argument that sets Name field",
	rbxmk.newDesc("Type", nil, "FooName").Name == "FooName")
T.Fail("with 'Type', with third non-string argument throws an error",
	function() rbxmk.newDesc("Type", nil, 42) end)
T.Pass("with 'Type', with each argument sets each component",
	function()
		local t = rbxmk.newDesc("Type", "FooCategory", "FooName")
		return t.Category == "FooCategory" and t.Name == "FooName"
	end)

T.Pass("with 'Parameter', has optional second TypeDesc argument that sets Type field",
	function()
		local t = rbxmk.newDesc("Type", "FooCategory", "FooName")
		return rbxmk.newDesc("Parameter", t).Type == t
	end)
T.Fail("with 'Parameter', with second non-TypeDesc argument throws an error",
	function() rbxmk.newDesc("Parameter", 42) end)
T.Pass("with 'Parameter', has optional third string argument that sets Name field",
	rbxmk.newDesc("Parameter", nil, "fooName").Name == "fooName")
T.Fail("with 'Parameter', with third non-string argument throws an error",
	function() rbxmk.newDesc("Parameter", nil, 42) end)
T.Pass("with 'Parameter', has optional fourth argument that, when nil, sets Default to nil",
	rbxmk.newDesc("Parameter").Default == nil)
T.Pass("with 'Parameter', has optional fourth argument that, when a string, sets Default",
	rbxmk.newDesc("Parameter", nil, nil, "FooDefault").Default == "FooDefault")
T.Fail("with 'Parameter', with fourth non-string argument throws an error",
	function() rbxmk.newDesc("Parameter", nil, nil, 42) end)
T.Pass("with 'Parameter', with each argument sets each component",
	function()
		local t = rbxmk.newDesc("Type", "FooCategory", "FooName")
		local p = rbxmk.newDesc("Parameter", t, "fooName", "FooDefault")
		return p.Type == t and p.Name == "fooName" and p.Default == "FooDefault"
	end)
