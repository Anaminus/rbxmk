local game = DataModel.new()
PASS(function() game.ExplicitAutoJoints = true end, "DataModel must set properties")
PASS(game.ExplicitAutoJoints == true              , "DataModel must get properties")
FAIL(function() game.ClassName = "BoolValue"   end, "cannot set ClassName of DataModel")

local workspace = game:GetService("Workspace")
PASS(workspace.ClassName == "Workspace"        , "GetService must set ClassName to given value")
PASS(workspace.Name == "Workspace"             , "GetService must set Name to given value")
PASS(workspace.Parent == game                  , "GetService must set Parent to DataModel")
PASS(game:GetService("Workspace") == workspace , "GetService must return singleton")

FAIL(function() Instance.new("BoolValue"):GetService("Workspace") end, "non-DataModel instance must not have GetService")
