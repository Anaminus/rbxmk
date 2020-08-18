-- Test members common to all descriptors.
local class = rbxmk.newDesc("ClassDesc")
local prop = rbxmk.newDesc("PropertyDesc")
local func = rbxmk.newDesc("FunctionDesc")
local event = rbxmk.newDesc("EventDesc")
local callback = rbxmk.newDesc("CallbackDesc")
local enum = rbxmk.newDesc("EnumDesc")
local item = rbxmk.newDesc("EnumItemDesc")

-- Name
for _, desc in ipairs({class,prop,func,event,callback,enum,item}) do
	local t = typeof(desc)
	T.Pass(function() return desc.Name end,                   t..": can get Name field")
	T.Pass(function() return type(desc.Name) == "string" end, t..": Name field is a string")
	T.Pass(function() return desc.Name == "" end,             t..": Name field initializes to empty string")
	T.Pass(function() desc.Name = "Foobar" end,               t..": can set Name field to string")
	T.Fail(function() desc.Name = 42 end,                     t..": cannot set Name field to non-string")
	T.Pass(function() return desc.Name == "Foobar" end,       t..": set Name field persists")
end

-- Tags
for _, desc in ipairs({class,prop,func,event,callback,enum,item}) do
	local t = typeof(desc)
	T.Pass(type(desc:Tag("")) == "boolean",                    t..": Tag method returns a boolean")
	T.Pass(type(desc:Tags()) == "table",                       t..": Tags method returns a table")
	T.Pass(select("#", desc:SetTag()) == 0,                    t..": SetTag method returns no values")
	T.Pass(select("#", desc:UnsetTag()) == 0,                  t..": UnsetTag method returns no values")
	T.Pass(#desc:Tags() == 0,                                  t..": descriptor initializes with no tags")
	T.Pass(function() desc:Tag("Foobar") end,                  t..": Tag can receive string as first argument")
	T.Fail(function() desc:Tag(42) end,                        t..": Tag cannot receive non-string as first argument")
	T.Pass(desc:Tag("TagA") == false,                          t..": getting unset tag returns false")
	T.Pass(function() desc:SetTag("TagA") end,                 t..": SetTag receives strings as arguments")
	T.Pass(function() desc:SetTag() end,                       t..": SetTag can receive no arguments")
	T.Pass(desc:Tag("TagA") == true,                           t..": getting set tag returns true")
	T.Pass(function() desc:SetTag("TagA", "TagB", "TagC") end, t..": SetTag can receive multiple arguments")
	T.Fail(function() desc:SetTag("TagA", 42, "TagC") end,     t..": SetTag cannot receive non-string argument")
	T.Pass(desc:Tag("TagA") == true,                           t..": first set tag persists")
	T.Pass(desc:Tag("TagB") == true,                           t..": second set tag persists")
	T.Pass(desc:Tag("TagC") == true,                           t..": third set tag persists")
	T.Pass(function()
		local tags = desc:Tags()
		return #tags == 3 and
		tags[1] == "TagA" and
		tags[2] == "TagB" and
		tags[3] == "TagC"
	end, t..": Tags returns all three set tags")

	T.Pass(function() desc:UnsetTag("TagA") end, t..": UnsetTag receives strings as arguments")
	T.Pass(function() desc:UnsetTag() end,       t..": UnsetTag can receive no arguments")
	T.Pass(desc:Tag("TagA") == false,            t..": unset tag persists")
	T.Pass(function()
		local tags = desc:Tags()
		return #tags == 2 and
		tags[1] == "TagB" and
		tags[2] == "TagC"
	end, t..": Tags returns all two set tags")
	T.Pass(function() desc:UnsetTag("TagA", "TagB", "TagC") end, t..": UnsetTag can receive multiple arguments")
	T.Fail(function() desc:UnsetTag("TagA", 42, "TagC") end,     t..": UnsetTag cannot receive non-string argument")
	T.Pass(desc:Tag("TagA") == false,                            t..": first unset tag persists")
	T.Pass(desc:Tag("TagB") == false,                            t..": second unset tag persists")
	T.Pass(desc:Tag("TagC") == false,                            t..": third unset tag persists")
	T.Pass(#desc:Tags() == 0,                                    t..": Tags returns no tags from all tags being unset")
end
