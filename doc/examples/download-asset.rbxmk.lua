-- Download a Roblox asset to a local file.
--
-- e.g. rbxmk run download-asset.rbxmk.lua 1818 Crossroads.rbxl

local id, output = ...
local asset = rbxassetid.read({
	AssetId = id,
	Format = "bin",
})
fs.write(output, asset, "bin")
