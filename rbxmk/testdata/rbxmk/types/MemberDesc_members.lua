-- Test members common to member descriptors.
local prop = rbxmk.newDesc("PropertyDesc")
local func = rbxmk.newDesc("FunctionDesc")
local event = rbxmk.newDesc("EventDesc")
local callback = rbxmk.newDesc("CallbackDesc")

-- Name
for _, desc in ipairs({prop,func,event,callback}) do
	local t = typeof(desc)
	T.Pass(t..": can get Name field",
		function() return desc.Name end)
	T.Pass(t..": Name field is a string",
		function() return type(desc.Name) == "string" end)
	T.Pass(t..": Name field initializes to empty string",
		function() return desc.Name == "" end)
	T.Pass(t..": can set Name field to string",
		function() desc.Name = "Foobar" end)
	T.Fail(t..": cannot set Name field to non-string",
		function() desc.Name = 42 end)
	T.Pass(t..": set Name field persists",
		function() return desc.Name == "Foobar" end)
end

-- Parameters
for _, desc in ipairs({func,event,callback}) do
	local t = typeof(desc)
	T.Pass(t..": can call Parameters method",
		desc:Parameters())
	T.Pass(t..": Parameters method returns a table",
		type(desc:Parameters()) == "table")
	T.Pass(t..": Parameters method initializes with empty table",
		#desc:Parameters() == 0)
	T.Pass(t..": can call SetParameters method with table",
		function() desc:SetParameters({}) end)
	T.Pass(t..": SetParameters returns no values",
		select("#", desc:SetParameters({})) == 0)
	T.Fail(t..": cannot call SetParameters method with non-table",
		function() desc:SetParameters(42) end)
	T.Fail(t..": cannot call SetParameters method with nil",
		function() desc:SetParameters() end)
	T.Fail(t..": cannot call SetParameters method with ParameterDesc",
		function() desc:SetParameters(rbxmk.newDesc("ParameterDesc")) end)
	T.Pass(t..": can call SetParameters method with table of ParameterDescs",
		function() desc:SetParameters({
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatA", "FooTypeA"), "fooNameA"),
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatB", "FooTypeB"), "fooNameB", "FooDefault"),
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatC", "FooTypeC"), "fooNameC"),
		}) end)
	T.Fail(t..": cannot call SetParameters method with table of non-ParameterDescs",
		function() desc:SetParameters({
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatA", "FooTypeA"), "fooNameA"),
			"Foobar",
			rbxmk.newDesc("ParameterDesc", rbxmk.newDesc("TypeDesc", "FooCatC", "FooTypeC"), "fooNameC"),
		}) end)
	T.Pass(t..": first set parameter persists",
		function()
			return desc:Parameters()[1] == rbxmk.newDesc(
				"ParameterDesc",
				rbxmk.newDesc("TypeDesc", "FooCatA", "FooTypeA"),
				"fooNameA"
			)
		end)
	T.Pass(t..": second set parameter persists",
		function()
			return desc:Parameters()[2] == rbxmk.newDesc(
				"ParameterDesc",
				rbxmk.newDesc("TypeDesc", "FooCatB", "FooTypeB"),
				"fooNameB",
				"FooDefault"
			)
		end)
	T.Pass(t..": third set parameter persists",
		function()
			return desc:Parameters()[3] == rbxmk.newDesc(
				"ParameterDesc",
				rbxmk.newDesc("TypeDesc", "FooCatC", "FooTypeC"),
				"fooNameC"
			)
		end)
end

-- ReturnType
for _, desc in ipairs({func,callback}) do
	local t = typeof(desc)
	T.Pass(t..": can get ReturnType field",
		function() return desc.ReturnType end)
	T.Pass(t..": ReturnType field is a TypeDesc",
		function() return typeof(desc.ReturnType) == "TypeDesc" end)
	T.Pass(t..": ReturnType field initializes to empty TypeDesc",
		function() return desc.ReturnType == rbxmk.newDesc("TypeDesc") end)
	T.Pass(t..": can set ReturnType field to TypeDesc",
		function() desc.ReturnType = rbxmk.newDesc("TypeDesc", "FooCategory", "FooName") end)
	T.Fail(t..": cannot set ReturnType field to non-string",
		function() desc.ReturnType = 42 end)
	T.Pass(t..": set ReturnType field persists",
		function() return desc.ReturnType.Category == "FooCategory" and desc.ReturnType.Name == "FooName" end)
end

-- Security
for _, desc in ipairs({func,event,callback}) do
	local t = typeof(desc)
	T.Pass(t..": can get Security field",
		function() return desc.Security end)
	T.Pass(t..": Security field is a string",
		function() return type(desc.Security) == "string" end)
	T.Pass(t..": Security field initializes to 'None'",
		function() return desc.Security == "None" end)
	T.Pass(t..": can set Security field to string",
		function() desc.Security = "Foobar" end)
	T.Fail(t..": cannot set Security field to non-string",
		function() desc.Security = 42 end)
	T.Pass(t..": set Security field persists",
		function() return desc.Security == "Foobar" end)
end
