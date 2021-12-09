local desc = fs.read(path.expand("$sd/../../../dump.desc.json"))
desc:Patch(fs.read(path.expand("$sd/../../../dump.desc-patch.json")))

-- Read, no desc
local model = fs.read(path.expand("$sd/model.rbxmx"))
local some = model:Descend("ModelSome")
local none = model:Descend("ModelNone")
T.Pass(typeof(some.WorldPivotData) == "Optional", "read, some is optional")
T.Pass(typeof(some.WorldPivotData.Value) == "CFrame", "read, some type is CFrame")
T.Pass(typeof(none.WorldPivotData) == "Optional", "read, none is optional")
T.Pass(typeof(none.WorldPivotData.Value) == "nil", "read, nome type is nil")

-- Write, no desc
local model = Instance.new("Model")
model.WorldPivotData = types.some(CFrame.new(1,2,3))
model = rbxmk.decodeFormat("rbxmx", rbxmk.encodeFormat("rbxmx", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "Optional", "write some, value is optional")
T.Pass(typeof(model.WorldPivotData.Value) == "CFrame", "write some, type is CFrame")
T.Pass(model.WorldPivotData.Value == CFrame.new(1,2,3), "write some, CFrame has value")

local model = Instance.new("Model")
model.WorldPivotData = types.none("CFrame")
model = rbxmk.decodeFormat("rbxmx", rbxmk.encodeFormat("rbxmx", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "Optional", "write none, value is optional")
T.Pass(typeof(model.WorldPivotData.Value) == "nil", "write none, type is nil")

--------------------------------------------------------------------------------
--------------------------------------------------------------------------------

rbxmk.globalDesc = desc

-- Read, with desc
local model = fs.read(path.expand("$sd/model.rbxmx"))
local some = model:Descend("ModelSome")
local none = model:Descend("ModelNone")
T.Pass(typeof(some.WorldPivotData) == "CFrame", "read with desc, some is CFrame")
T.Pass(some.WorldPivotData == CFrame.new(10,20,30), "read with desc, some CFrame has value")
T.Pass(typeof(none.WorldPivotData) == "nil", "read with desc, none is nil")

-- Write, with desc
local model = Instance.new("Model")
model.WorldPivotData = types.some(CFrame.new(1,2,3))
model = rbxmk.decodeFormat("rbxmx", rbxmk.encodeFormat("rbxmx", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "CFrame", "write some with desc, value is CFrame")
T.Pass(model.WorldPivotData == CFrame.new(1,2,3), "write some with desc, CFrame has value")

local model = Instance.new("Model")
model.WorldPivotData = types.none("CFrame")
model = rbxmk.decodeFormat("rbxmx", rbxmk.encodeFormat("rbxmx", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "nil", "write none with desc, value is nil")
