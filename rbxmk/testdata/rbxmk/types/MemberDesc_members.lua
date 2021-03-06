-- Test members common to member descriptors.
local prop = rbxmk.newDesc("PropertyDesc")
local func = rbxmk.newDesc("FunctionDesc")
local event = rbxmk.newDesc("EventDesc")
local callback = rbxmk.newDesc("CallbackDesc")

-- Parameters
for _, desc in ipairs({func,event,callback}) do
	local t = typeof(desc)
	T.Pass(desc:Parameters()                        , t..": can call Parameters method")
	T.Pass(type(desc:Parameters()) == "table"       , t..": Parameters method returns a table")
	T.Pass(#desc:Parameters() == 0                  , t..": Parameters method initializes with empty table")
	T.Pass(function() desc:SetParameters({}) end    , t..": can call SetParameters method with table")
	T.Pass(select("#", desc:SetParameters({})) == 0 , t..": SetParameters returns no values")
	T.Fail(function() desc:SetParameters(42) end    , t..": cannot call SetParameters method with non-table")
	T.Fail(function() desc:SetParameters() end      , t..": cannot call SetParameters method with nil")
	T.Fail(function()
		desc:SetParameters(rbxmk.newDesc("ParameterDesc"))
	end, t..": cannot call SetParameters method with ParameterDesc")
	T.Pass(function()
		desc:SetParameters({
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatA", "FooTypeA"), "fooNameA"),
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatB", "FooTypeB"), "fooNameB", "FooDefault"),
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatC", "FooTypeC"), "fooNameC"),
		})
	end, t..": can call SetParameters method with table of ParameterDescs")
	T.Fail(function()
		desc:SetParameters({
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatA", "FooTypeA"), "fooNameA"),
			"Foobar",
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatC", "FooTypeC"), "fooNameC"),
		})
	end, t..": cannot call SetParameters method with table of non-ParameterDescs")
	T.Pass(function()
		return desc:Parameters()[1] == rbxmk.newDesc(
			"ParameterDesc",
			rbxmk.newDesc("TypeDesc", "FooCatA", "FooTypeA"),
			"fooNameA"
		)
	end, t..": first set parameter persists")
	T.Pass(function()
		return desc:Parameters()[2] == rbxmk.newDesc(
			"ParameterDesc",
			rbxmk.newDesc("TypeDesc", "FooCatB", "FooTypeB"),
			"fooNameB",
			"FooDefault"
		)
	end, t..": second set parameter persists")
	T.Pass(function()
		return desc:Parameters()[3] == rbxmk.newDesc(
			"ParameterDesc",
			rbxmk.newDesc("TypeDesc", "FooCatC", "FooTypeC"),
			"fooNameC"
		)
	end, t..": third set parameter persists")
end

-- ReturnType
for _, desc in ipairs({func,callback}) do
	local t = typeof(desc)
	T.Pass(function() return desc.ReturnType end,
		t..": can get ReturnType field")
	T.Pass(function() return typeof(desc.ReturnType) == "TypeDesc" end,
		t..": ReturnType field is a TypeDesc")
	T.Pass(function() return desc.ReturnType == rbxmk.newDesc("TypeDesc") end,
		t..": ReturnType field initializes to empty TypeDesc")
	T.Pass(function() desc.ReturnType = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName") end,
		t..": can set ReturnType field to TypeDesc")
	T.Fail(function() desc.ReturnType = 42 end,
		t..": cannot set ReturnType field to non-string")
	T.Pass(function() return desc.ReturnType.Category == "FooCategory" and desc.ReturnType.Name == "FooName" end,
		t..": set ReturnType field persists")
end

-- Security
for _, desc in ipairs({func,event,callback}) do
	local t = typeof(desc)
	T.Pass(function() return desc.Security end                   , t..": can get Security field")
	T.Pass(function() return type(desc.Security) == "string" end , t..": Security field is a string")
	T.Pass(function() return desc.Security == "None" end         , t..": Security field initializes to 'None'")
	T.Pass(function() desc.Security = "Foobar" end               , t..": can set Security field to string")
	T.Fail(function() desc.Security = 42 end                     , t..": cannot set Security field to non-string")
	T.Pass(function() return desc.Security == "Foobar" end       , t..": set Security field persists")
end
