local desc = fs.read(path.expand("$sd/../../../dump.desc.json"))
desc:Patch(fs.read(path.expand("$sd/../../../dump.desc-patch.json")))

rbxmk.globalDesc = nil

-- Read xml, no desc
local model = fs.read(path.expand("$sd/model.rbxmx"))
local some = model:Descend("ModelSome")
local none = model:Descend("ModelNone")
T.Pass(typeof(some.WorldPivotData) == "Optional", "read xml, some is optional")
T.Pass(typeof(some.WorldPivotData.Value) == "CFrame", "read xml, some type is CFrame")
T.Pass(typeof(none.WorldPivotData) == "Optional", "read xml, none is optional")
T.Pass(typeof(none.WorldPivotData.Value) == "nil", "read xml, nome type is nil")

-- Write xml, no desc
local model = Instance.new("Model")
model.WorldPivotData = types.some(CFrame.new(1,2,3))
model = rbxmk.decodeFormat("rbxmx", rbxmk.encodeFormat("rbxmx", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "Optional", "write xml some, value is optional")
T.Pass(typeof(model.WorldPivotData.Value) == "CFrame", "write xml some, type is CFrame")
T.Pass(model.WorldPivotData.Value == CFrame.new(1,2,3), "write xml some, CFrame has value")

local model = Instance.new("Model")
model.WorldPivotData = types.none("CFrame")
model = rbxmk.decodeFormat("rbxmx", rbxmk.encodeFormat("rbxmx", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "Optional", "write xml none, value is optional")
T.Pass(typeof(model.WorldPivotData.Value) == "nil", "write xml none, type is nil")

--------------------------------------------------------------------------------
--------------------------------------------------------------------------------

rbxmk.globalDesc = desc

-- Read xml, with desc
local model = fs.read(path.expand("$sd/model.rbxmx"))
local some = model:Descend("ModelSome")
local none = model:Descend("ModelNone")
T.Pass(typeof(some.WorldPivotData) == "CFrame", "read xml with desc, some is CFrame")
T.Pass(some.WorldPivotData == CFrame.new(10,20,30), "read xml with desc, some CFrame has value")
T.Pass(typeof(none.WorldPivotData) == "nil", "read xml with desc, none is nil")

-- Write xml, with desc
local model = Instance.new("Model")
model.WorldPivotData = types.some(CFrame.new(1,2,3))
model = rbxmk.decodeFormat("rbxmx", rbxmk.encodeFormat("rbxmx", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "CFrame", "write xml some with desc, value is CFrame")
T.Pass(model.WorldPivotData == CFrame.new(1,2,3), "write xml some with desc, CFrame has value")

local model = Instance.new("Model")
model.WorldPivotData = types.none("CFrame")
model = rbxmk.decodeFormat("rbxmx", rbxmk.encodeFormat("rbxmx", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "nil", "write xml none with desc, value is nil")

--------------------------------------------------------------------------------
--------------------------------------------------------------------------------

rbxmk.globalDesc = nil

-- Read binary, no desc
local model = fs.read(path.expand("$sd/model.rbxm"))
local some = model:Descend("ModelSome")
local none = model:Descend("ModelNone")
T.Pass(typeof(some.WorldPivotData) == "Optional", "read binary, some is optional")
T.Pass(typeof(some.WorldPivotData.Value) == "CFrame", "read binary, some type is CFrame")
T.Pass(typeof(none.WorldPivotData) == "Optional", "read binary, none is optional")
T.Pass(typeof(none.WorldPivotData.Value) == "nil", "read binary, nome type is nil")

-- Write binary, no desc
local model = Instance.new("Model")
model.WorldPivotData = types.some(CFrame.new(1,2,3))
model = rbxmk.decodeFormat("rbxm", rbxmk.encodeFormat("rbxm", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "Optional", "write binary some, value is optional")
T.Pass(typeof(model.WorldPivotData.Value) == "CFrame", "write binary some, type is CFrame")
T.Pass(model.WorldPivotData.Value == CFrame.new(1,2,3), "write binary some, CFrame has value")

local model = Instance.new("Model")
model.WorldPivotData = types.none("CFrame")
model = rbxmk.decodeFormat("rbxm", rbxmk.encodeFormat("rbxm", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "Optional", "write binary none, value is optional")
T.Pass(typeof(model.WorldPivotData.Value) == "nil", "write binary none, type is nil")

--------------------------------------------------------------------------------
--------------------------------------------------------------------------------

rbxmk.globalDesc = desc

-- Read binary, with desc
local model = fs.read(path.expand("$sd/model.rbxm"))
local some = model:Descend("ModelSome")
local none = model:Descend("ModelNone")
T.Pass(typeof(some.WorldPivotData) == "CFrame", "read binary with desc, some is CFrame")
T.Pass(some.WorldPivotData == CFrame.new(10,20,30), "read binary with desc, some CFrame has value")
T.Pass(typeof(none.WorldPivotData) == "nil", "read binary with desc, none is nil")

-- Write binary, with desc
local model = Instance.new("Model")
model.WorldPivotData = types.some(CFrame.new(1,2,3))
model = rbxmk.decodeFormat("rbxm", rbxmk.encodeFormat("rbxm", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "CFrame", "write binary some with desc, value is CFrame")
T.Pass(model.WorldPivotData == CFrame.new(1,2,3), "write binary some with desc, CFrame has value")

local model = Instance.new("Model")
model.WorldPivotData = types.none("CFrame")
model = rbxmk.decodeFormat("rbxm", rbxmk.encodeFormat("rbxm", model)):GetChildren()[1]
T.Pass(typeof(model.WorldPivotData) == "nil", "write binary none with desc, value is nil")
