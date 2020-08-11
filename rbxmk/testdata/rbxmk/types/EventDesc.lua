local desc = rbxmk.newDesc("Event")

-- Metamethod tests
T.Pass("type of descriptor is EventDesc",
	typeof(desc) == "EventDesc")
T.Pass("metatable of descriptor is locked",
	type(getmetatable(desc)) == "string")
T.Pass("descriptor converts to a string",
	not string.match(tostring(desc), "^userdata"))
T.Pass("descriptor is equal to itself",
	desc == desc)
T.Pass("descriptor is not equal to another descriptor of the same type",
	desc ~= rbxmk.newDesc("Event"))
