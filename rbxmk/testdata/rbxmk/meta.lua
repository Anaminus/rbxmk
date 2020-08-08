-- meta
local v = Instance.new("BoolValue")
T.Pass("meta sets with 3 arguments",
	select("#", rbxmk.meta(v, "Reference", "foobar")) == 0)
T.Pass("meta gets with 2 arguments",
	select("#", rbxmk.meta(v, "Reference")) == 1)

-- Reference
local v = Instance.new("BoolValue")
T.Pass("Reference initializes with empty string",
	rbxmk.meta(v, "Reference") == "")
T.Pass("Reference is set to string",
	function() rbxmk.meta(v, "Reference", "foobar") end)
T.Fail("Reference errors with non-string",
	function() rbxmk.meta(v, "Reference", true) end)
T.Pass("Reference returns set value",
	rbxmk.meta(v, "Reference") == "foobar")

-- IsService
local v = Instance.new("BoolValue")
T.Pass("IsService initializes with false",
	rbxmk.meta(v, "IsService") == false)
T.Pass("IsService is set to boolean",
	function() rbxmk.meta(v, "IsService", true) end)
T.Fail("IsService errors with non-boolean",
	function() rbxmk.meta(v, "IsService", "foobar") end)
T.Pass("IsService returns set value",
	rbxmk.meta(v, "IsService") == true)

-- Desc
local desc = file.read(os.expand("$sd/../dump.desc.json"))
local v = Instance.new("BoolValue")
T.Pass("Desc initializes with nil",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc initializes with nil",
	rbxmk.meta(v, "RawDesc") == nil)
T.Pass("Desc can be set to RootDesc",
	function() rbxmk.meta(v, "Desc", desc) end)
T.Pass("Desc returns RootDesc after being set to RootDesc",
	rbxmk.meta(v, "Desc") == desc)
T.Pass("RawDesc returns RootDesc after Desc is set to RootDesc",
	rbxmk.meta(v, "RawDesc") == desc)
T.Pass("Desc can be set to nil",
	function() rbxmk.meta(v, "Desc", nil) end)
T.Pass("Desc returns nil after being set to nil",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc returns nil after Desc is set to nil",
	rbxmk.meta(v, "RawDesc") == nil)
T.Pass("Desc can be set to false",
	function() rbxmk.meta(v, "Desc", false) end)
T.Fail("Desc cannot be set to true",
	function() rbxmk.meta(v, "Desc", true) end)
T.Pass("Desc returns nil after being set to false",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc returns false after Desc is set to false",
	rbxmk.meta(v, "RawDesc") == false)
T.Fail("Desc errors without RootDesc, false, or nil",
	function() rbxmk.meta(v, "Desc", "foobar") end)

-- RawDesc
local desc = file.read(os.expand("$sd/../dump.desc.json"))
local v = Instance.new("BoolValue")
T.Pass("RawDesc initializes with nil",
	rbxmk.meta(v, "RawDesc") == nil)
T.Pass("Desc initializes with nil",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc can be set to RootDesc",
	function() rbxmk.meta(v, "RawDesc", desc) end)
T.Pass("RawDesc returns RootDesc after being set to RootDesc",
	rbxmk.meta(v, "RawDesc") == desc)
T.Pass("Desc returns RootDesc after RawDesc is set to RootDesc",
	rbxmk.meta(v, "Desc") == desc)
T.Pass("RawDesc can be set to nil",
	function() rbxmk.meta(v, "RawDesc", nil) end)
T.Pass("RawDesc returns nil after being set to nil",
	rbxmk.meta(v, "RawDesc") == nil)
T.Pass("Desc returns nil after RawDesc is set to nil",
	rbxmk.meta(v, "Desc") == nil)
T.Pass("RawDesc can be set to false",
	function() rbxmk.meta(v, "RawDesc", false) end)
T.Fail("RawDesc cannot be set to true",
	function() rbxmk.meta(v, "RawDesc", true) end)
T.Pass("RawDesc returns false after being set to false",
	rbxmk.meta(v, "RawDesc") == false)
T.Pass("Desc returns nil after RawDesc is set to false",
	rbxmk.meta(v, "Desc") == nil)
T.Fail("RawDesc errors without RootDesc, false, or nil",
	function() rbxmk.meta(v, "RawDesc", "foobar") end)
