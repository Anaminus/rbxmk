local game = DataModel.new()
T.Pass("DataModel can set properties",
	function() game.ExplicitAutoJoints = true end)
T.Pass("DataModel can get properties",
	game.ExplicitAutoJoints == true)
T.Fail("cannot set ClassName of DataModel",
	function() game.ClassName = "BoolValue" end)

local workspace = game:GetService("Workspace")
T.Pass("GetService sets ClassName to given value",
	workspace.ClassName == "Workspace")
T.Pass("GetService sets Name to given value",
	workspace.Name == "Workspace")
T.Pass("GetService sets Parent to DataModel",
	workspace.Parent == game)
T.Pass("GetService returns singleton",
	game:GetService("Workspace") == workspace)

T.Fail("non-DataModel instance does not have GetService",
	function() Instance.new("BoolValue"):GetService("Workspace") end)
