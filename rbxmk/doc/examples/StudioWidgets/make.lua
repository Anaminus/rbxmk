-- https://github.com/Roblox/StudioWidgets
-- Compile scripts into StudioWidgets.rbxmx file.

-- Command:
-- rbxmk -a 'path/to/repo' make.lua

-- Get location of repository.
local root = ...
if type(root) ~= "string" then
	-- Assume working directory.
	root = "."
end

-- Find Lua files.
local scriptsDir = rbxmk.path{root, "src"}
local scripts = {}
for _, file in pairs(rbxmk.readdir{scriptsDir}) do
	if rbxmk.filename{"fext", file} == ".lua" then
		-- Read Lua file as ModuleScript.
		scripts[#scripts+1] = rbxmk.input{format="modulescript.lua", rbxmk.path{scriptsDir, file}}
	end
end

-- Clear existing file.
rbxmk.delete{rbxmk.output{"StudioWidgets.rbxmx"}}
-- Create StudioWidgets folder.
rbxmk.map{
	rbxmk.output{"StudioWidgets.rbxmx"},
	rbxmk.input{"generate://Instance", [[Folder{Name:string="StudioWidgets"}]]},
}
-- Write scripts to folder.
rbxmk.map{rbxmk.output{"StudioWidgets.rbxmx", "StudioWidgets"}, unpack(scripts)}
