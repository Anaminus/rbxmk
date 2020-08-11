T.Fail("expects a string for its first argument",
	function() rbxmk.newDesc(42) end)
T.Pass("with 'Root', returns RootDesc",
	typeof(rbxmk.newDesc("RootDesc")) == "RootDesc")
T.Pass("with 'Class', returns ClassDesc",
	typeof(rbxmk.newDesc("ClassDesc")) == "ClassDesc")
T.Pass("with 'Property', returns PropertyDesc",
	typeof(rbxmk.newDesc("PropertyDesc")) == "PropertyDesc")
T.Pass("with 'Function', returns FunctionDesc",
	typeof(rbxmk.newDesc("FunctionDesc")) == "FunctionDesc")
T.Pass("with 'Event', returns EventDesc",
	typeof(rbxmk.newDesc("EventDesc")) == "EventDesc")
T.Pass("with 'Callback', returns CallbackDesc",
	typeof(rbxmk.newDesc("CallbackDesc")) == "CallbackDesc")
T.Pass("with 'Parameter', returns ParameterDesc",
	typeof(rbxmk.newDesc("ParameterDesc")) == "ParameterDesc")
T.Pass("with 'Type', returns TypeDesc",
	typeof(rbxmk.newDesc("TypeDesc")) == "TypeDesc")
T.Pass("with 'Enum', returns EnumDesc",
	typeof(rbxmk.newDesc("EnumDesc")) == "EnumDesc")
T.Pass("with 'EnumItem', returns EnumItemDesc",
	typeof(rbxmk.newDesc("EnumItemDesc")) == "EnumItemDesc")
T.Fail("with unknown type, throws an error",
	function() rbxmk.newDesc("NonextantDesc") end)

T.Pass("with 'Type', has optional second string argument that sets Category field",
	rbxmk.newDesc("TypeDesc", "FooCategory").Category == "FooCategory")
T.Fail("with 'Type', with second non-string argument throws an error",
	function() rbxmk.newDesc("TypeDesc", 42) end)
T.Pass("with 'Type', has optional third string argument that sets Name field",
	rbxmk.newDesc("TypeDesc", nil, "FooName").Name == "FooName")
T.Fail("with 'Type', with third non-string argument throws an error",
	function() rbxmk.newDesc("TypeDesc", nil, 42) end)
T.Pass("with 'Type', with each argument sets each component",
	function()
		local t = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")
		return t.Category == "FooCategory" and t.Name == "FooName"
	end)

T.Pass("with 'Parameter', has optional second TypeDesc argument that sets Type field",
	function()
		local t = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")
		return rbxmk.newDesc("ParameterDesc", t).Type == t
	end)
T.Fail("with 'Parameter', with second non-TypeDesc argument throws an error",
	function() rbxmk.newDesc("ParameterDesc", 42) end)
T.Pass("with 'Parameter', has optional third string argument that sets Name field",
	rbxmk.newDesc("ParameterDesc", nil, "fooName").Name == "fooName")
T.Fail("with 'Parameter', with third non-string argument throws an error",
	function() rbxmk.newDesc("ParameterDesc", nil, 42) end)
T.Pass("with 'Parameter', has optional fourth argument that, when nil, sets Default to nil",
	rbxmk.newDesc("ParameterDesc").Default == nil)
T.Pass("with 'Parameter', has optional fourth argument that, when a string, sets Default",
	rbxmk.newDesc("ParameterDesc", nil, nil, "FooDefault").Default == "FooDefault")
T.Fail("with 'Parameter', with fourth non-string argument throws an error",
	function() rbxmk.newDesc("ParameterDesc", nil, nil, 42) end)
T.Pass("with 'Parameter', with each argument sets each component",
	function()
		local t = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")
		local p = rbxmk.newDesc("ParameterDesc", t, "fooName", "FooDefault")
		return p.Type == t and p.Name == "fooName" and p.Default == "FooDefault"
	end)
