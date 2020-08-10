T.Fail("newDesc expects a string for its first argument",
	function() rbxmk.newDesc(42) end)
T.Pass("newDesc with 'Root' returns RootDesc",
	typeof(rbxmk.newDesc("Root")) == "RootDesc")
T.Pass("newDesc with 'Class' returns ClassDesc",
	typeof(rbxmk.newDesc("Class")) == "ClassDesc")
T.Pass("newDesc with 'Property' returns PropertyDesc",
	typeof(rbxmk.newDesc("Property")) == "PropertyDesc")
T.Pass("newDesc with 'Function' returns FunctionDesc",
	typeof(rbxmk.newDesc("Function")) == "FunctionDesc")
T.Pass("newDesc with 'Event' returns EventDesc",
	typeof(rbxmk.newDesc("Event")) == "EventDesc")
T.Pass("newDesc with 'Callback' returns CallbackDesc",
	typeof(rbxmk.newDesc("Callback")) == "CallbackDesc")
T.Pass("newDesc with 'Parameter' returns ParameterDesc",
	typeof(rbxmk.newDesc("Parameter")) == "ParameterDesc")
T.Pass("newDesc with 'Type' returns TypeDesc",
	typeof(rbxmk.newDesc("Type")) == "TypeDesc")
T.Pass("newDesc with 'Enum' returns EnumDesc",
	typeof(rbxmk.newDesc("Enum")) == "EnumDesc")
T.Pass("newDesc with 'EnumItem' returns EnumItemDesc",
	typeof(rbxmk.newDesc("EnumItem")) == "EnumItemDesc")
T.Fail("newDesc with unknown type throws an error",
	function() rbxmk.newDesc("NonextantType") end)
