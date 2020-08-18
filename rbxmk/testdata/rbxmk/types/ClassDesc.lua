local desc = rbxmk.newDesc("ClassDesc")

-- Metamethod tests
T.Pass(typeof(desc) == "ClassDesc"                  , "type of descriptor is ClassDesc")
T.Pass(type(getmetatable(desc)) == "string"         , "metatable of descriptor is locked")
T.Pass(not string.match(tostring(desc), "^userdata"), "descriptor converts to a string")
T.Pass(desc == desc                                 , "descriptor is equal to itself")
T.Pass(desc ~= rbxmk.newDesc("ClassDesc")           , "descriptor is not equal to another descriptor of the same type")

-- Superclass
T.Pass(function() return desc.Superclass end                  , "can get Superclass field")
T.Pass(function() return type(desc.Superclass) == "string" end, "Superclass field is a string")
T.Pass(function() return desc.Superclass == "" end            , "Superclass field initializes to empty string")
T.Pass(function() desc.Superclass = "Foobar" end              , "can set Superclass field to string")
T.Fail(function() desc.Superclass = 42 end                    , "cannot set Superclass field to non-string")
T.Pass(function() return desc.Superclass == "Foobar" end      , "set Superclass field persists")

-- MemoryCategory
T.Pass(function() return desc.MemoryCategory end                  , "can get MemoryCategory field")
T.Pass(function() return type(desc.MemoryCategory) == "string" end, "MemoryCategory field is a string")
T.Pass(function() return desc.MemoryCategory == "" end            , "MemoryCategory field initializes to empty string")
T.Pass(function() desc.MemoryCategory = "Foobar" end              , "can set MemoryCategory field to string")
T.Fail(function() desc.MemoryCategory = 42 end                    , "cannot set MemoryCategory field to non-string")
T.Pass(function() return desc.MemoryCategory == "Foobar" end      , "set MemoryCategory field persists")

-- Members
local prop = rbxmk.newDesc("PropertyDesc")
prop.Name = "Property"
local func = rbxmk.newDesc("FunctionDesc")
func.Name = "Function"
local event = rbxmk.newDesc("EventDesc")
event.Name = "Event"
local callback = rbxmk.newDesc("CallbackDesc")
callback.Name = "Callback"

T.Pass(function() desc:Member("") end        , "can call Member method with string")
T.Fail(function() desc:Member(42) end        , "cannot call Member method with non-string")
T.Pass(desc:AddMember(prop) == true          , "can call AddMember method with PropertyDesc")
T.Pass(desc:AddMember(func) == true          , "can call AddMember method with FunctionDesc")
T.Pass(desc:AddMember(event) == true         , "can call AddMember method with EventDesc")
T.Pass(desc:AddMember(callback) == true      , "can call AddMember method with CallbackDesc")
T.Pass(desc:AddMember(prop) == false         , "AddMember returns false for existing member")
T.Fail(function() desc:AddMember(42) end     , "cannot call AddMember method with non-member descriptor")
T.Pass(desc:Member("Nonextant") == nil       , "Member returns nil for nonextant member")
T.Pass(desc:Member("Property") == prop       , "Member returns member for existing member")
T.Pass(function() desc:RemoveMember("") end  , "can call RemoveMember method with string")
T.Fail(function() desc:RemoveMember(42) end  , "cannot call RemoveMember method with non-string")
T.Pass(desc:RemoveMember("Property") == true , "RemoveMember returns true for existing member")
T.Pass(desc:RemoveMember("Property") == false, "RemoveMember returns false for nonextant member")
T.Pass(desc:Member("Property") == nil        , "Removal of member persists")
T.Pass(type(desc:Members()) == "table"       , "Members method returns a table")
T.Pass(#desc:Members() == 3                  , "has three members")
T.Pass(desc:Members()[1] == callback         , "first member is Callback")
T.Pass(desc:Members()[2] == event            , "second member is Event")
T.Pass(desc:Members()[3] == func             , "third member is Function")
