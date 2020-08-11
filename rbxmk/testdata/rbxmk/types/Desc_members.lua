-- Test members common to all descriptors.
local class = rbxmk.newDesc("Class")
local prop = rbxmk.newDesc("Property")
local func = rbxmk.newDesc("Function")
local event = rbxmk.newDesc("Event")
local callback = rbxmk.newDesc("Callback")
local enum = rbxmk.newDesc("Enum")
local item = rbxmk.newDesc("EnumItem")

-- Tags
for _, desc in ipairs({class,prop,func,event,callback,enum,item}) do
	local t = typeof(desc)
	T.Pass(t..": Tag method returns a boolean",
		type(desc:Tag("")) == "boolean")
	T.Pass(t..": Tags method returns a table",
		type(desc:Tags()) == "table")
	T.Pass(t..": SetTag method returns no values",
		select("#", desc:SetTag()) == 0)
	T.Pass(t..": UnsetTag method returns no values",
		select("#", desc:UnsetTag()) == 0)
	T.Pass(t..": descriptor initializes with no tags",
		#desc:Tags() == 0)
	T.Pass(t..": Tag can receive string as first argument",
		function() desc:Tag("Foobar") end)
	T.Fail(t..": Tag cannot receive non-string as first argument",
		function() desc:Tag(42) end)
	T.Pass(t..": getting unset tag returns false",
		desc:Tag("TagA") == false)
	T.Pass(t..": SetTag receives strings as arguments",
		function() desc:SetTag("TagA") end)
	T.Pass(t..": SetTag can receive no arguments",
		function() desc:SetTag() end)
	T.Pass(t..": getting set tag returns true",
		desc:Tag("TagA") == true)
	T.Pass(t..": SetTag can receive multiple arguments",
		function() desc:SetTag("TagA", "TagB", "TagC") end)
	T.Fail(t..": SetTag cannot receive non-string argument",
		function() desc:SetTag("TagA", 42, "TagC") end)
	T.Pass(t..": first set tag persists",
		desc:Tag("TagA") == true)
	T.Pass(t..": second set tag persists",
		desc:Tag("TagB") == true)
	T.Pass(t..": third set tag persists",
		desc:Tag("TagC") == true)
	T.Pass(t..": Tags returns all three set tags",
		function()
			local tags = desc:Tags()
			return #tags == 3 and
			tags[1] == "TagA" and
			tags[2] == "TagB" and
			tags[3] == "TagC"
		end)

	T.Pass(t..": UnsetTag receives strings as arguments",
		function() desc:UnsetTag("TagA") end)
	T.Pass(t..": UnsetTag can receive no arguments",
		function() desc:UnsetTag() end)
	T.Pass(t..": unset tag persists",
		desc:Tag("TagA") == false)
	T.Pass(t..": Tags returns all two set tags",
		function()
			local tags = desc:Tags()
			return #tags == 2 and
			tags[1] == "TagB" and
			tags[2] == "TagC"
		end)
	T.Pass(t..": UnsetTag can receive multiple arguments",
		function() desc:UnsetTag("TagA", "TagB", "TagC") end)
	T.Fail(t..": UnsetTag cannot receive non-string argument",
		function() desc:UnsetTag("TagA", 42, "TagC") end)
	T.Pass(t..": first unset tag persists",
		desc:Tag("TagA") == false)
	T.Pass(t..": second unset tag persists",
		desc:Tag("TagB") == false)
	T.Pass(t..": third unset tag persists",
		desc:Tag("TagC") == false)
	T.Pass(t..": Tags returns no tags from all tags being unset",
		#desc:Tags() == 0)
end
