T.Fail(function() rbxmk.newDesc(42) end                         , "expects a string for its first argument")
T.Pass(typeof(rbxmk.newDesc("RootDesc")) == "RootDesc"          , "with 'Root', returns RootDesc")
T.Pass(typeof(rbxmk.newDesc("ClassDesc")) == "ClassDesc"        , "with 'Class', returns ClassDesc")
T.Pass(typeof(rbxmk.newDesc("PropertyDesc")) == "PropertyDesc"  , "with 'Property', returns PropertyDesc")
T.Pass(typeof(rbxmk.newDesc("FunctionDesc")) == "FunctionDesc"  , "with 'Function', returns FunctionDesc")
T.Pass(typeof(rbxmk.newDesc("EventDesc")) == "EventDesc"        , "with 'Event', returns EventDesc")
T.Pass(typeof(rbxmk.newDesc("CallbackDesc")) == "CallbackDesc"  , "with 'Callback', returns CallbackDesc")
T.Pass(typeof(rbxmk.newDesc("ParameterDesc")) == "ParameterDesc", "with 'Parameter', returns ParameterDesc")
T.Pass(typeof(rbxmk.newDesc("TypeDesc")) == "TypeDesc"          , "with 'Type', returns TypeDesc")
T.Pass(typeof(rbxmk.newDesc("EnumDesc")) == "EnumDesc"          , "with 'Enum', returns EnumDesc")
T.Pass(typeof(rbxmk.newDesc("EnumItemDesc")) == "EnumItemDesc"  , "with 'EnumItem', returns EnumItemDesc")
T.Fail(function() rbxmk.newDesc("NonextantDesc") end            , "with unknown type, throws an error")

T.Pass(rbxmk.newDesc("TypeDesc", "FooCategory").Category == "FooCategory", "with 'Type', has optional second string argument that sets Category field")
T.Fail(function() rbxmk.newDesc("TypeDesc", 42) end                      , "with 'Type', with second non-string argument throws an error")
T.Pass(rbxmk.newDesc("TypeDesc", nil, "FooName").Name == "FooName"       , "with 'Type', has optional third string argument that sets Name field")
T.Fail(function() rbxmk.newDesc("TypeDesc", nil, 42) end                 , "with 'Type', with third non-string argument throws an error")
T.Pass(function()
	local t = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")
	return t.Category == "FooCategory" and t.Name == "FooName"
end, "with 'Type', with each argument sets each component")

T.Pass(function()
	local t = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")
	return rbxmk.newDesc("ParameterDesc", t).Type == t
end, "with 'Parameter', has optional second TypeDesc argument that sets Type field")
T.Fail(function() rbxmk.newDesc("ParameterDesc", 42) end                             , "with 'Parameter', with second non-TypeDesc argument throws an error")
T.Pass(rbxmk.newDesc("ParameterDesc", nil, "fooName").Name == "fooName"              , "with 'Parameter', has optional third string argument that sets Name field")
T.Fail(function() rbxmk.newDesc("ParameterDesc", nil, 42) end                        , "with 'Parameter', with third non-string argument throws an error")
T.Pass(rbxmk.newDesc("ParameterDesc").Default == nil                                 , "with 'Parameter', has optional fourth argument that, when nil, sets Default to nil")
T.Pass(rbxmk.newDesc("ParameterDesc", nil, nil, "FooDefault").Default == "FooDefault", "with 'Parameter', has optional fourth argument that, when a string, sets Default")
T.Fail(function() rbxmk.newDesc("ParameterDesc", nil, nil, 42) end                   , "with 'Parameter', with fourth non-string argument throws an error")
T.Pass(function()
	local t = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName")
	local p = rbxmk.newDesc("ParameterDesc", t, "fooName", "FooDefault")
	return p.Type == t and p.Name == "fooName" and p.Default == "FooDefault"
end, "with 'Parameter', with each argument sets each component")
