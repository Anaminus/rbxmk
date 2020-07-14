local game = DataModel.new()
game.ExplicitAutoJoints = true

local workspace = game:GetService("Workspace")
assert(workspace.ClassName == "Workspace", "GetService must set ClassName to given value")
assert(workspace.Name == "Workspace", "GetService must set Name to given value")
assert(workspace.Parent == game, "GetService must set Parent to DataModel")
assert(game:GetService("Workspace") == workspace, "GetService must return singleton")

assert(not pcall(function() Instance.new("BoolValue"):GetService("Workspace") end))
