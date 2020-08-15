local desc = rbxmk.newDesc("ClassDesc")

-- Metamethod tests
T.Pass("type of descriptor is ClassDesc",
	typeof(desc) == "ClassDesc")
T.Pass("metatable of descriptor is locked",
	type(getmetatable(desc)) == "string")
T.Pass("descriptor converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("descriptor is equal to itself",
	desc == desc)
T.Pass("descriptor is not equal to another descriptor of the same type",
	desc ~= rbxmk.newDesc("ClassDesc"))

-- Superclass
T.Pass("can get Superclass field",
	function() return desc.Superclass end)
T.Pass("Superclass field is a string",
	function() return type(desc.Superclass) == "string" end)
T.Pass("Superclass field initializes to empty string",
	function() return desc.Superclass == "" end)
T.Pass("can set Superclass field to string",
	function() desc.Superclass = "Foobar" end)
T.Fail("cannot set Superclass field to non-string",
	function() desc.Superclass = 42 end)
T.Pass("set Superclass field persists",
	function() return desc.Superclass == "Foobar" end)

-- MemoryCategory
T.Pass("can get MemoryCategory field",
	function() return desc.MemoryCategory end)
T.Pass("MemoryCategory field is a string",
	function() return type(desc.MemoryCategory) == "string" end)
T.Pass("MemoryCategory field initializes to empty string",
	function() return desc.MemoryCategory == "" end)
T.Pass("can set MemoryCategory field to string",
	function() desc.MemoryCategory = "Foobar" end)
T.Fail("cannot set MemoryCategory field to non-string",
	function() desc.MemoryCategory = 42 end)
T.Pass("set MemoryCategory field persists",
	function() return desc.MemoryCategory == "Foobar" end)

-- Members
local prop = rbxmk.newDesc("PropertyDesc")
prop.Name = "Property"
local func = rbxmk.newDesc("FunctionDesc")
func.Name = "Function"
local event = rbxmk.newDesc("EventDesc")
event.Name = "Event"
local callback = rbxmk.newDesc("CallbackDesc")
callback.Name = "Callback"

T.Pass("can call Member method with string",
	function() desc:Member("") end)
T.Fail("cannot call Member method with non-string",
	function() desc:Member(42) end)
T.Pass("can call AddMember method with PropertyDesc",
	desc:AddMember(prop) == true)
T.Pass("can call AddMember method with FunctionDesc",
	desc:AddMember(func) == true)
T.Pass("can call AddMember method with EventDesc",
	desc:AddMember(event) == true)
T.Pass("can call AddMember method with CallbackDesc",
	desc:AddMember(callback) == true)
T.Pass("AddMember returns false for existing member",
	desc:AddMember(prop) == false)
T.Fail("cannot call AddMember method with non-member descriptor",
	function() desc:AddMember(42) end)
T.Pass("Member returns nil for nonextant member",
	desc:Member("Nonextant") == nil)
T.Pass("Member returns member for existing member",
	desc:Member("Property") == prop)
T.Pass("can call RemoveMember method with string",
	function() desc:RemoveMember("") end)
T.Fail("cannot call RemoveMember method with non-string",
	function() desc:RemoveMember(42) end)
T.Pass("RemoveMember returns true for existing member",
	desc:RemoveMember("Property") == true)
T.Pass("RemoveMember returns false for nonextant member",
	desc:RemoveMember("Property") == false)
T.Pass("removal of member persists",
	desc:Member("Property") == nil)
T.Pass("Members method returns a table",
	type(desc:Members()) == "table")
T.Pass("has three members",
	#desc:Members() == 3)
T.Pass("first member is Callback",
	desc:Members()[1] == callback)
T.Pass("second member is Event",
	desc:Members()[2] == event)
T.Pass("third member is Function",
	desc:Members()[3] == func)
